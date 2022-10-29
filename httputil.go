package main

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

type HTTPDoer func(ctx context.Context) (*http.Response, error)

const defaultBackoff = time.Millisecond * 500

func DoHttpCallWithRetry(ctx context.Context, maxAttempts int, doer HTTPDoer) (*http.Response, error) {
	var (
		attempts = 0
		lastErr  error
	)
	for {
		attempts++
		resp, err := doer(ctx)
		if err == nil && resp.StatusCode >= 200 && resp.StatusCode < 300 {
			return resp, err
		}
		if err != nil {
			lastErr = err
		} else {
			lastErr = fmt.Errorf("status: %d", resp.StatusCode)
			resp.Body.Close()
		}

		retryable, next := isRetryable(resp, err)
		if !retryable {
			break
		}
		if attempts >= maxAttempts {
			lastErr = fmt.Errorf("exceeded retry limit")
			break
		}

		timer := time.NewTimer(next)
		select {
		case <-ctx.Done():
			timer.Stop()
			return nil, ctx.Err()
		case <-timer.C:
			timer.Stop()
		}
	}
	return nil, lastErr
}

func isRetryable(resp *http.Response, err error) (retryable bool, wait time.Duration) {
	if err != nil {
		if e, ok := err.(*url.Error); ok {
			msg := e.Error()
			if strings.Contains(msg, "stopped after") && strings.Contains(msg, "redirects") {
				return false, 0
			}
			if strings.Contains(msg, "unsupported protocol scheme") {
				return false, 0
			}
			if strings.Contains(msg, "certificated is not trusted") {
				return false, 0
			}
		}
		return true, defaultBackoff
	}

	if resp.StatusCode == http.StatusTooManyRequests {
		next := getRetryAfterHeader(resp.Header)
		if next == 0 {
			next = defaultBackoff
		}
		return true, next
	}
	if resp.StatusCode >= http.StatusInternalServerError {
		return true, defaultBackoff
	}
	return false, 0
}

func getRetryAfterHeader(header http.Header) time.Duration {
	retryAfterStr := header.Get("Retry-After")
	if retryAfterStr == "" {
		return 0
	}
	retryAfter, parseErr := strconv.Atoi(retryAfterStr)
	if parseErr != nil {
		return 0
	}
	return time.Duration(retryAfter) * time.Second
}
