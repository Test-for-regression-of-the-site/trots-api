package api

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"sort"
	"sync"
	"time"

	"github.com/Test-for-regression-of-the-site/trots-api/lighthouse"
	"github.com/google/uuid"
)

type createTaskRequest struct {
	Time     int64    `json:"time"`
	Links    []string `json:"links"`
	Parallel int      `json:"parallel"`
	Type     string   `json:"typetest"`
}

func (req *createTaskRequest) Timestamp() time.Time {
	return time.Unix(req.Time, 0)
}

type lighthouseTasks struct {
	cfg   lighthouse.Config
	mu    sync.RWMutex
	tasks map[string]*task
}

func newLighthouseTasks() *lighthouseTasks {
	return &lighthouseTasks{
		tasks: make(map[string]*task),
	}
}

type task struct {
	sync.RWMutex
	done chan struct{}

	running bool
	err     error
	page    string
	report  *bytes.Buffer
}

func (lt *lighthouseTasks) AddTask(rw http.ResponseWriter, req *http.Request) {
	lt.mu.Lock()
	defer lt.mu.Unlock()
	var id = uuid.New().String()
	var taskRequest createTaskRequest
	if err := readObject(req.Body, &taskRequest); err != nil {
		http.Error(rw, err.Error(), http.StatusBadRequest)
		return
	}
	for _, page := range taskRequest.Links {
		var task = &task{
			done:   make(chan struct{}),
			page:   page,
			report: &bytes.Buffer{},
		}
		lt.tasks[id] = task
		_, _ = rw.Write([]byte(id))
		go func() {
			defer close(task.done)

			task.Lock()
			defer task.Unlock()

			task.running = true
			defer func() { task.running = false }()
			task.err = lt.cfg.Run(task.page, task.report)
		}()
	}
}

func readObject(re io.Reader, dst interface{}) error {
	var data, errRead = ioutil.ReadAll(re)
	if errRead != nil {
		return errRead
	}
	return json.Unmarshal(data, dst)
}

func (lt *lighthouseTasks) Status(rw http.ResponseWriter, req *http.Request) {
	var id = req.URL.Query().Get("id")
	if id == "" {
		rw.Header().Set("Content-Type", "application/json; charset=utf-8")
		var ids = lt.statuses()
		_ = json.NewEncoder(rw).Encode(ids)
		return
	}

	lt.mu.RLock()
	defer lt.mu.RUnlock()
	var task, ok = lt.tasks[id]
	if !ok {
		rw.WriteHeader(http.StatusNotFound)
		return
	}
	select {
	case <-req.Context().Done():
		http.Error(rw, "zombi task", http.StatusInternalServerError)
		return
	case <-task.done:
		rw.Header().Set("Content-Type", "application/json; charset=utf-8")
		_, _ = task.report.WriteTo(rw)
	}
}

type taskStatus struct {
	ID      string `json:"id"`
	Page    string `json:"page"`
	Running bool   `json:"running"`
	Error   string `json:"error,omitempty"`
}

func (lt *lighthouseTasks) statuses() []taskStatus {
	lt.mu.RLock()
	defer lt.mu.RUnlock()
	var tasks = make([]taskStatus, 0, len(lt.tasks))
	for id, task := range lt.tasks {
		var status = taskStatus{
			ID:      id,
			Running: true,
			Page:    task.page,
		}
		select {
		case <-task.done:
			status.Running = false
			if task.err != nil {
				status.Error = task.err.Error()
			}
		default:
		}
		tasks = append(tasks, status)
	}
	sort.Slice(tasks, func(i, j int) bool {
		var a, b = tasks[i], tasks[j]
		return a.ID < b.ID
	})
	return tasks
}
