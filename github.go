package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"regexp"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/pkg/errors"
)

type GithubRepository struct {
	Owner           string `json:"owner" yaml:"owner"`
	Name            string `json:"name" yaml:"name"`
	FullName        string `json:"fullName" yaml:"fullName"`
	URL             string `json:"url" yaml:"url"`
	StargazersCount int    `json:"stargazersCount" yaml:"stargazersCount"`
}

func NewGithubRepository(owner, name string) *GithubRepository {
	return &GithubRepository{
		Owner:           owner,
		Name:            name,
		FullName:        fmt.Sprintf("%s/%s", owner, name),
		URL:             fmt.Sprintf("https://github.com/%s/%s", owner, name),
		StargazersCount: 0,
	}
}

type GithubRepositoryFilter func(r *GithubRepository) bool

type GithubRepositoryFilters []GithubRepositoryFilter

func (filters GithubRepositoryFilters) Filter(repository *GithubRepository) bool {
	for _, filter := range filters {
		if !filter(repository) {
			return false
		}
	}
	return true
}

var numericRegex = regexp.MustCompile(`[^0-9 ]+`)

func parseStarsFilter(opt string) (GithubRepositoryFilter, error) {
	if opt == "" {
		return nil, nil
	}

	var op, val []rune
	for i, r := range opt {
		if r >= '0' && r <= '9' {
			val = append(val, r)
			continue
		}
		if strings.ContainsRune("=><", r) {
			op = append(op, r)
			continue
		}
		return nil, fmt.Errorf("'%c' invalid character at %d", r, i)
	}
	v, _ := strconv.Atoi(string(val))

	switch string(op) {
	case "", "=":
		return func(r *GithubRepository) bool {
			return r.StargazersCount == v
		}, nil
	case ">":
		return func(r *GithubRepository) bool {
			return r.StargazersCount > v
		}, nil
	case ">=":
		return func(r *GithubRepository) bool {
			return r.StargazersCount >= v
		}, nil
	case "<":
		return func(r *GithubRepository) bool {
			return r.StargazersCount < v
		}, nil
	case "<=":
		return func(r *GithubRepository) bool {
			return r.StargazersCount <= v
		}, nil
	}
	return nil, fmt.Errorf("unreachable code. option: %s, op: %s, val: %d", opt, string(op), val)
}

var countCache = make(map[string]int)

func CountStargazers(ctx context.Context, accessToken, owner, repo string) (int, error) {
	fullName := fmt.Sprintf("%s/%s", owner, repo)
	if count, ok := countCache[fullName]; ok {
		return count, nil
	}

	var (
		count int
		err   error
	)
	if accessToken != "" {
		count, err = countStargazersByGithubAPI(ctx, accessToken, owner, repo)
	} else {
		count, err = countStargazersByGithubPublic(ctx, owner, repo)
	}
	if err != nil {
		return 0, err
	}
	countCache[fullName] = count
	return count, nil
}

func countStargazersByGithubAPI(ctx context.Context, accessToken, owner, repo string) (int, error) {
	endpoint := fmt.Sprintf("https://api.github.com/repos/%s/%s", owner, repo)
	resp, err := DoHttpCallWithRetry(ctx, 3, func(ctx context.Context) (*http.Response, error) {
		req, err := http.NewRequestWithContext(ctx, http.MethodGet, endpoint, nil)
		if err != nil {
			return nil, err
		}
		req.Header.Set("Authorization", "Bearer "+accessToken)
		return http.DefaultClient.Do(req)
	})
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	r := struct {
		StargazersCount int `json:"stargazers_count"`
	}{}
	if err := json.NewDecoder(resp.Body).Decode(&r); err != nil {
		return 0, err
	}
	return r.StargazersCount, nil
}

func countStargazersByGithubPublic(ctx context.Context, owner, repo string) (int, error) {
	endpoint := fmt.Sprintf("https://github.com/%s/%s", owner, repo)
	resp, err := DoHttpCallWithRetry(ctx, 3, func(ctx context.Context) (*http.Response, error) {
		req, err := http.NewRequestWithContext(ctx, http.MethodGet, endpoint, nil)
		if err != nil {
			return nil, errors.Wrapf(err, "new request. url: %s", endpoint)
		}
		return http.DefaultClient.Do(req)
	})
	if err != nil {
		return 0, errors.Wrap(err, "get github stars")
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return 0, fmt.Errorf("get github stars. status: %d", resp.StatusCode)
	}

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return 0, errors.Wrap(err, "new document")
	}

	stargazers := -1
	sel := doc.Find(".Counter")
	for i := range sel.Nodes {
		s := sel.Eq(i)
		if !strings.Contains(strings.TrimSpace(s.Parent().Text()), "Star") {
			continue
		}
		title, ok := s.Attr("title")
		if !ok {
			continue
		}
		countStr := strings.TrimSpace(numericRegex.ReplaceAllString(title, ""))
		count, err := strconv.Atoi(countStr)
		if err != nil {
			return 0, errors.Wrapf(err, "parse text to star count. text: %s", title)
		}
		stargazers = count
		break
	}
	if stargazers < 0 {
		return 0, errors.New("cannot find stargazers count")
	}
	return stargazers, nil
}
