package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"strings"
	"sync"
	"syscall"
)

const version = "0.1.0"

var (
	Commit = ""
	Tag    = ""
)

func main() {
	importedCmd := flag.NewFlagSet("imported", flag.ExitOnError)
	versionCmd := flag.NewFlagSet("version", flag.ExitOnError)
	commands := strings.Join([]string{
		importedCmd.Name(), versionCmd.Name(),
	}, " | ")

	if len(os.Args) < 2 {
		log.Fatalf("require [%s] subcommand", commands)
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	done := make(chan error, 1)
	running := sync.WaitGroup{}

	switch os.Args[1] {
	case importedCmd.Name():
		opts := importedOpts{}

		importedCmd.UintVar(&opts.Number, "n", 10, "Number of retrieved repositories.")
		importedCmd.StringVar(&opts.Package, "p", "", "Golang Package. e.g) net/http, github.com/gin-gonic/gin")

		if err := importedCmd.Parse(os.Args[2:]); err != nil {
			log.Fatal(err)
		}
		if err := opts.Validate(); err != nil {
			log.Fatal(err)
		}

		running.Add(1)
		go func() {
			defer running.Done()
			log.Printf("Search imported github repositories. limit: %d, package: %s", opts.Number, opts.Package)
			done <- RunImportedCmd(ctx, &opts)
		}()
	case versionCmd.Name():
		fmt.Println("gopkgsearch version", getVersionInfo())
		close(done)
	default:
		log.Fatalf("invalid subcommand: %s", os.Args[1])
	}

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGQUIT, syscall.SIGTERM)

	select {
	case sig := <-quit:
		log.Printf("received termination signal: %s", sig.String())
		cancel()
		running.Wait()
	case err := <-done:
		if err != nil {
			log.Fatal(err)
		}
		os.Exit(0)
	}
}

func getVersionInfo() string {
	if Tag == "" || Tag == "undefined" {
		v := version
		if len(Commit) >= 7 {
			v += "+" + Commit[:7]
		}
		return v
	}
	return Tag
}
