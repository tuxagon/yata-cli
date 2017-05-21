package cmd

import (
	"github.com/urfave/cli"
)

// Push TODO docs
func Push(ctx *cli.Context) error {
	// Looks like the golang library requires oauth, which
	// is a lot more work to set everything up, so it might
	// be worth writing the network code in a separate lang
	// Python or Ruby, depending on whichever allows api keys
	return nil
}
