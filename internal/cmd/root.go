package cmd

import (
	"Telegraph/internal/cmd/serve"
)

func Exec() {
	e := serve.GetServer()
	e.Logger.Fatal(e.Start(":5000"))
}
