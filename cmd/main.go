package main

import (
	"github.com/Test-for-regression-of-the-site/trots-api/provider"
	"github.com/Test-for-regression-of-the-site/trots-api/server"
)

func main() {
	provider.LoadConfiguration()
	server.Serve()
}
