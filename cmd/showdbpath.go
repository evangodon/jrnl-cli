package cmd

import (
	"fmt"
	"jrnl/pkg/sqldb"
	"jrnl/pkg/ui"

	lg "github.com/charmbracelet/lipgloss"
	"github.com/urfave/cli/v2"
)

var ShowDbPathCmd = &cli.Command{
	Name:    "showdbpath",
	Aliases: []string{"sdp"},
	Usage:   "Show the path to the database",
	Action: func(c *cli.Context) error {

		path := lg.NewStyle().
			Background(ui.Color.Primary).
			Padding(0, 2).
			Render(sqldb.GetDbPath())

		fmt.Println("\nDb path: ", path)

		return nil
	},
}
