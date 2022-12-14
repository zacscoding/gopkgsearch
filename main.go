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
	readmeGenCmd := flag.NewFlagSet("readmegen", flag.ExitOnError)
	commands := strings.Join([]string{
		importedCmd.Name(), versionCmd.Name(), readmeGenCmd.Name(),
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
		importedCmd.StringVar(&opts.Stars, "stars", "", "Stars filter. e.g) '=10', '>10', '>=10'")
		importedCmd.StringVar(&opts.Output, "output", "table", "Output format. available: ['table', 'json', 'yaml', 'markdown'")
		importedCmd.StringVar(&opts.OutputPath, "o", "", "Output path. default: std")

		if err := importedCmd.Parse(os.Args[2:]); err != nil {
			log.Fatal(err)
		}
		opts.GithubAccessToken = os.Getenv("GITHUB_ACCESS_TOKEN")

		if err := opts.Init(); err != nil {
			log.Fatal(err)
		}

		running.Add(1)
		go func() {
			defer running.Done()
			useGithubAPI := opts.GithubAccessToken != ""
			log.Printf("Search imported github repositories. limit: %d, package: %s, enable github api: %v", opts.Number, opts.Package, useGithubAPI)
			done <- RunImportedCmd(ctx, &opts)
		}()
	case readmeGenCmd.Name():
		running.Add(1)
		go func() {
			defer running.Done()
			done <- RunReadmeGen(ctx)
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
