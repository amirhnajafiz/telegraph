package cmd

import (
	"github.com/amirhnajafiz/telegraph/internal/cmd/migrate"
	"github.com/amirhnajafiz/telegraph/internal/cmd/serve"
	"github.com/spf13/cobra"
)

func Exec() {
	rootCommand := cobra.Command{}

	rootCommand.AddCommand(
		migrate.GetCommand(),
		serve.GetCommand(),
	)

	if err := rootCommand.Execute(); err != nil {
		panic(err)
	}
}
