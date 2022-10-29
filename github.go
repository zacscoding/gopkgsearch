package main

import (
	"context"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/pkg/errors"
)

type GithubRepository struct {
	Owner           string
	Name            string
	FullName        string
	StargazersCount int
}

func CountStargazers(ctx context.Context, repo *GithubRepository) (int, error) {
	endpoint := fmt.Sprintf("https://github.com/%s", repo.FullName)
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
