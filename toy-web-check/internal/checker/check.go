package checker

import (
	"net/http"
	"sync"
	"time"
)

type Result struct {
	URL        string
	StatusCode int
	Err        error
}

func CheckURLs(urls []string, timeout time.Duration) []Result {
	results := make([]Result, len(urls))
	var wg sync.WaitGroup
	ch := make(chan Result)

	client := &http.Client{Timeout: timeout}

	for _, url := range urls {
		wg.Add(1)
		go func(u string) {
			defer wg.Done()
			resp, err := client.Get(u)
			if err != nil {
				ch <- Result{URL: u, Err: err}
				return
			}
			defer resp.Body.Close()
			ch <- Result{URL: u, StatusCode: resp.StatusCode}
		}(url)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	i := 0
	for res := range ch {
		results[i] = res
		i++
	}
	return results
}
