package main

import (
	"github.com/wjojf/go-htmx/transport/server"
)

func main() {
	srv := server.New()
	if err := srv.Start(":42069"); err != nil {
		srv.Logger.Fatal(err)
	}
}
