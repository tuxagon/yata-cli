package cmd

import "github.com/urfave/cli"
import "fmt"

// Show will display a task based on its ID
func Show(ctx *cli.Context) error {
	fmt.Printf("%+v\n", ctx.Args())
	return nil
}
