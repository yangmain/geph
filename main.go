package main

import (
	"flag"
	"math/rand"
	"os"

	"golang.org/x/net/context"

	"github.com/bunsim/geph/binder"
	"github.com/bunsim/geph/client"
	"github.com/bunsim/geph/entry"
	"github.com/bunsim/geph/exit"
	"github.com/bunsim/geph/proxbinder"
	"github.com/google/subcommands"
	"gopkg.in/bunsim/natrium.v1"
)

func main() {
	// seed the insecure RNG from the secure one to prevent repeats
	rand.Seed(int64(natrium.RandUint32()))
	subcommands.Register(subcommands.HelpCommand(), "")
	subcommands.Register(subcommands.FlagsCommand(), "")
	subcommands.Register(subcommands.CommandsCommand(), "")
	subcommands.Register(&binder.Command{}, "")
	subcommands.Register(&exit.Command{}, "")
	subcommands.Register(&entry.Command{}, "")
	subcommands.Register(&client.Command{}, "")
	subcommands.Register(&proxbinder.Command{}, "")
	flag.Parse()
	ctx := context.Background()
	os.Exit(int(subcommands.Execute(ctx)))
}
