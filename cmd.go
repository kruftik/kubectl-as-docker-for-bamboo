package main

import (
	"github.com/jessevdk/go-flags"
	"log"
	"os"
)

// Opts with all cli commands and flags
type Opts struct {
	RunCmd  RunCommand  `command:"run" description:"Run a command in a new POD"`
	CopyCmd  CopyCommand  `command:"cp" description:"Copy files/folders between a POD and the local filesystem"`
	ExecCmd  ExecCommand  `command:"exec" description:"Run a command in a running POD"`
}

var (
	//command Command
)

func ParseAndRunCommand(){
	var opts Opts
	p := flags.NewParser(&opts, flags.Default)

	p.CommandHandler = func(command flags.Commander, args []string) error {
		err := command.Execute(args)
		if err != nil {
			log.Printf("[ERROR] failed with %+v", err)
		}

		return err
	}

	if _, err := p.Parse(); err != nil {
		if flagsErr, ok := err.(*flags.Error); ok && flagsErr.Type == flags.ErrHelp {
			os.Exit(0)
		} else {
			os.Exit(1)
		}
	}
}
