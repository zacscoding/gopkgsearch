package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/pkg/errors"
	"gopkg.in/yaml.v3"
)

type importedOpts struct {
	Number            uint
	Package           string
	Stars             string
	Output            string
	OutputPath        string
	GithubAccessToken string

	StartTime time.Time
	Elapsed   time.Duration
	Filters   GithubRepositoryFilters
}

func (opts *importedOpts) Init() error {
	if opts.Package == "" {
		return fmt.Errorf("required package")
	}
	if opts.Stars != "" {
		filter, err := parseStarsFilter(opts.Stars)
		if err != nil {
			return err
		}
		opts.Filters = append(opts.Filters, filter)
	}
	switch opts.Output {
	default:
		return fmt.Errorf("invalid output format: %s", opts.Output)
	case "table", "json", "yaml", "markdown":
	}
	return nil
}

func RunImportedCmd(ctx context.Context, opts *importedOpts) error {
	// Search imported repositories from "https://pkg.go.dev/:opts.Package?tab=importedby"
	opts.StartTime = time.Now()
	repositories, err := SearchImportedRepositories(ctx, opts.Package)
	if err != nil {
		return err
	}
	log.Printf("> imported repositories: %d", len(repositories))

	// Search stargazers count from "https://github.com/:owner/:name"
	for i, repo := range repositories {
		if i%100 == 0 {
			log.Printf("processed repositories: %d", i)
		}
		count, err := CountStargazers(ctx, opts.GithubAccessToken, repo.Owner, repo.Name)
		if err != nil {
			if errors.Is(err, ctx.Err()) {
				return err
			}
			if !strings.Contains(err.Error(), "not found") {
				log.Printf("failed to get %s stargazers count. err: %v", repo.FullName, err)
			}
			continue
		}
		repo.StargazersCount = count
	}

	// Sort by stars
	sort.Slice(repositories, func(i, j int) bool {
		return repositories[i].StargazersCount > repositories[j].StargazersCount
	})

	var filtered []*GithubRepository
	limit := int(opts.Number)
	if limit < len(repositories) {
		repositories = repositories[:limit]
	}
	for _, repository := range repositories {
		if opts.Filters.Filter(repository) {
			filtered = append(filtered, repository)
		}
	}
	opts.Elapsed = time.Since(opts.StartTime)
	return PrintGithubRepositories(ctx, opts, filtered)
}

func SearchImportedRepositories(ctx context.Context, pkg string) ([]*GithubRepository, error) {
	endpoint := fmt.Sprintf("https://pkg.go.dev/%s?tab=importedby", pkg)
	resp, err := DoHttpCallWithRetry(ctx, 3, func(ctx context.Context) (*http.Response, error) {
		req, err := http.NewRequestWithContext(ctx, http.MethodGet, endpoint, nil)
		if err != nil {
			return nil, errors.Wrapf(err, "new request. url: %s", endpoint)
		}
		return http.DefaultClient.Do(req)
	})
	if err != nil {
		return nil, errors.Wrap(err, "get go packages")
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("get go packages. status: %d", resp.StatusCode)
	}

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return nil, errors.Wrap(err, "new document")
	}

	var (
		repositories []*GithubRepository
		visited      = make(map[string]struct{})
	)
	doc.Find(".u-breakWord").Each(func(i int, s *goquery.Selection) {
		href, ok := s.Attr("href")
		if !ok || !strings.Contains(href, "/github.com") {
			return
		}

		values := strings.Split(href[len("/github.com/"):], "/")
		if len(values) < 2 {
			return
		}
		owner, name := values[0], values[1]
		fullName := owner + "/" + name
		if _, ok := visited[fullName]; ok {
			return
		}
		visited[fullName] = struct{}{}
		repositories = append(repositories, NewGithubRepository(owner, name))
	})
	return repositories, nil
}

func PrintGithubRepositories(ctx context.Context, opts *importedOpts, repositories []*GithubRepository) error {
	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
	}

	var writer io.Writer
	if opts.OutputPath == "" {
		writer = os.Stdout
	} else {
		splits := strings.Split(opts.Package, "/")
		filename := fmt.Sprintf("search_%s_%s.log", splits[len(splits)-1], time.Now().Format("2006-01-02_15:04:05"))
		f, err := os.Create(filepath.Join(opts.OutputPath, filename))
		if err != nil {
			return err
		}
		writer = f
	}

	s := struct {
		Package      string              `json:"package" yaml:"package"`
		Elapsed      time.Duration       `json:"elapsed" yaml:"elapsed"`
		UpdatedAt    time.Time           `json:"updatedAt" yaml:"updatedAt"`
		Repositories []*GithubRepository `json:"repositories" yaml:"repositories"`
	}{
		Package:      opts.Package,
		Elapsed:      opts.Elapsed,
		UpdatedAt:    time.Now(),
		Repositories: repositories,
	}

	switch opts.Output {
	case "table", "markdown":
		writer.Write([]byte(fmt.Sprintf("- Packages: %s\n", s.Package)))
		writer.Write([]byte(fmt.Sprintf("- UpdatedAt: %s\n", s.UpdatedAt.UTC().Format("2006-01-02 15:04:05"))))
		writer.Write([]byte(fmt.Sprintf("- Elapsed: %s\n", s.Elapsed)))
		writer.Write([]byte("\n"))

		tb := table.NewWriter()
		tb.AppendHeader(table.Row{"No", "User", "Repository", "Stargazers", "URL"})
		for i, repo := range repositories {
			tb.AppendRow(table.Row{
				i + 1, repo.Owner, repo.Name, repo.StargazersCount, fmt.Sprintf("https://github.com/%s", repo.FullName),
			})
		}
		var render string
		if opts.Output == "table" {
			render = tb.Render()
		} else {
			render = tb.RenderMarkdown()
		}
		_, err := writer.Write([]byte(render))
		return err
	case "json":
		return json.NewEncoder(writer).Encode(&s)
	case "yaml":
		return yaml.NewEncoder(writer).Encode(&s)
	}
	return nil
}
