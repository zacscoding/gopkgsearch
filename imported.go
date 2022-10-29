package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"regexp"
	"sort"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/pkg/errors"
)

var numericRegex = regexp.MustCompile(`[^0-9 ]+`)

type importedOpts struct {
	Number  uint
	Package string
}

func (opt *importedOpts) Validate() error {
	if opt.Package == "" {
		return fmt.Errorf("required package")
	}
	return nil
}

func RunImportedCmd(ctx context.Context, opts *importedOpts) error {
	// Search imported repositories from "https://pkg.go.dev/:opts.Package?tab=importedby"
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
		count, err := CountStargazers(ctx, repo)
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

	limit := int(opts.Number)
	if limit < len(repositories) {
		repositories = repositories[:limit]
	}
	return PrintGithubRepositories(ctx, repositories)
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
		repositories = append(repositories, &GithubRepository{Owner: owner, Name: name, FullName: fullName})
	})
	return repositories, nil
}

func PrintGithubRepositories(ctx context.Context, repositories []*GithubRepository) error {
	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
	}

	// TODO: support multiple output options such as html, markdown.
	tb := table.NewWriter()
	tb.AppendHeader(table.Row{"No", "User", "Repository", "Stargazers", "URL"})
	for i, repo := range repositories {
		tb.AppendRow(table.Row{
			i + 1, repo.Owner, repo.Name, repo.StargazersCount, fmt.Sprintf("https://github.com/%s", repo.FullName),
		})
	}
	tb.AppendSeparator()
	fmt.Println(tb.Render())
	return nil
}
