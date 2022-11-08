package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"

	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/pkg/errors"
	"gopkg.in/yaml.v3"
)

type Category struct {
	Category string     `json:"category" yaml:"category"`
	Packages []*Package `json:"packages" yaml:"packages"`
}

type Package struct {
	Name          string        `json:"name" yaml:"name"`
	URL           string        `json:"url" yaml:"url"`
	SearchOptions SearchOptions `json:"options" yaml:"options"`

	Repositories []*GithubRepository
	Err          error
}

type SearchOptions struct {
	Limit int    `json:"limit" yaml:"limit"`
	Stars string `json:"stars" yaml:"stars"`
}

func RunReadmeGen(ctx context.Context) error {
	var categories []Category
	data, err := os.ReadFile("packages.yaml")
	if err != nil {
		return errors.Wrap(err, "read packages.yml")
	}
	if err := yaml.Unmarshal(data, &categories); err != nil {
		return errors.Wrap(err, "unmarshal packages.yml")
	}

	for _, category := range categories {
		log.Printf("Try to search %s", category.Category)
		for _, pkg := range category.Packages {
			repositories, err := SearchImportedRepositories(ctx, pkg.URL)
			if err != nil {
				pkg.Err = err
				continue
			}
			log.Printf("%s(%s) imported by repositories: %d", pkg.Name, pkg.URL, len(repositories))

			for idx, repo := range repositories {
				if idx%100 == 0 {
					log.Printf("Search %s imported repositories.. %d", pkg.Name, idx)
				}
				count, err := CountStargazers(ctx, "", repo.Owner, repo.Name)
				if err == nil {
					repo.StargazersCount = count
				}
			}

			// Sort by stars
			sort.Slice(repositories, func(i, j int) bool {
				return repositories[i].StargazersCount > repositories[j].StargazersCount
			})

			var filtered []*GithubRepository
			limit := pkg.SearchOptions.Limit
			if limit < len(repositories) {
				repositories = repositories[:limit]
			}

			filter, err := parseStarsFilter(pkg.SearchOptions.Stars)
			if err != nil {
				filtered = repositories[:]
			} else {
				for _, repo := range repositories {
					if filter(repo) {
						filtered = append(filtered, repo)
					}
				}
			}
			pkg.Repositories = filtered
		}
	}
	return writeReadme(categories)
}

func writeReadme(categories []Category) error {
	f, err := os.OpenFile("README_PREPARE.md", os.O_RDWR|os.O_TRUNC, 0644)
	if err != nil {
		return errors.Wrap(err, "open read file")
	}
	head, err := os.ReadFile("head.md")
	if err != nil {
		return errors.Wrap(err, "read head.md")
	}
	f.Write(head)

	// Write Index
	f.WriteString("# Popular Projects  \n")
	for _, category := range categories {
		f.WriteString(fmt.Sprintf("- [%s](#%s)\n", category.Category, strings.ReplaceAll(category.Category, " ", "-")))
		for _, pkg := range category.Packages {
			f.WriteString(fmt.Sprintf("  - [%s](#%s)\n", pkg.Name, strings.ReplaceAll(pkg.Name, " ", "-")))
		}
	}
	f.WriteString("\n")
	f.WriteString("---  \n")
	f.WriteString("\n")

	for _, category := range categories {
		f.WriteString(fmt.Sprintf("## %s\n\n", category.Category))
		for _, pkg := range category.Packages {
			f.WriteString(fmt.Sprintf("### %s\n", pkg.Name))
			if pkg.Err != nil {
				f.WriteString(fmt.Sprintf("fetch error: %v\n\n", pkg.Err))
				continue
			}
			tb := table.NewWriter()
			tb.AppendHeader(table.Row{"No", "User", "Repository", "Stargazers", "URL"})
			for i, repo := range pkg.Repositories {
				tb.AppendRow(table.Row{
					i + 1, repo.Owner, repo.Name, repo.StargazersCount, fmt.Sprintf("https://github.com/%s", repo.FullName),
				})
			}
			f.WriteString(tb.RenderMarkdown())
			f.WriteString("\n\n")
			f.WriteString(fmt.Sprintf("**[⬆ top](#Popular-Projects)**  \n"))
			f.WriteString("\n\n")
		}
		f.WriteString("\n\n")
	}
	return nil
}
