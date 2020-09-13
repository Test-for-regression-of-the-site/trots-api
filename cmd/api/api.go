package main

import (
	"context"
	"log"

	"github.com/Test-for-regression-of-the-site/trots-api/api"
)

func main() {
	var service, err = api.NewAPI(func(cfg *api.Config) error {
		cfg.Lighthouse.Exec = "/home/merlin/.yarn/bin/lighthouse"
		return nil
	})
	if err != nil {
		log.Fatal(err)
	}
	var ctx = context.Background()
	var errServe = service.ListenAndServe(ctx)
	if errServe != nil {
		log.Fatal(errServe)
	}
}
