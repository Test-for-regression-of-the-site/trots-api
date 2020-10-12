package main

import (
	"github.com/Test-for-regression-of-the-site/trots-api/server"
	"github.com/Test-for-regression-of-the-site/trots-api/storage"
)

func main() {
	server.Serve()
	storage.Disconnect()
}
