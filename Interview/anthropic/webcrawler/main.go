package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"
	"sync"

	"golang.org/x/net/html"
)

// Crawler represents a web crawler.
type Crawler struct {
	visited  map[string]bool
	mu       sync.Mutex
	wg       sync.WaitGroup
	urlQueue chan string
	baseURL  *url.URL
}

// NewCrawler creates a new Crawler.
func NewCrawler(startURL string) (*Crawler, error) {
	u, err := url.Parse(startURL)
	if err != nil {
		return nil, err
	}
	return &Crawler{
		visited:  make(map[string]bool),
		urlQueue: make(chan string, 100),
		baseURL:  u,
	}, nil
}

// Crawl starts the web crawling process.
func (c *Crawler) Crawl(startURL string, workers int) {
	c.wg.Add(1)
	c.urlQueue <- startURL

	for i := 0; i < workers; i++ {
		go c.worker()
	}

	c.wg.Wait()
	close(c.urlQueue)
}

func (c *Crawler) worker() {
	for rawURL := range c.urlQueue {
		c.mu.Lock()
		if c.visited[rawURL] {
			c.mu.Unlock()
			// This URL was already queued by another worker, but the WaitGroup
			// was not incremented for it, so we must not call Done().
			continue
		}
		c.visited[rawURL] = true
		c.mu.Unlock()

		fmt.Println("Crawling:", rawURL)

		body, err := fetch(rawURL)
		if err != nil {
			fmt.Fprintf(os.Stderr, "failed to fetch %s: %v", rawURL, err)
			c.wg.Done()
			continue
		}

		links := parse(body, c.baseURL)
		body.Close()

		for _, link := range links {
			c.mu.Lock()
			if !c.visited[link] {
				c.visited[link] = true
				c.wg.Add(1)
				go func(l string) {
					c.urlQueue <- l
				}(link)
			}
			c.mu.Unlock()
		}
		c.wg.Done()
	}
}

func fetch(rawURL string) (io.ReadCloser, error) {
	resp, err := http.Get(rawURL)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("status code %d", resp.StatusCode)
	}

	return resp.Body, nil
}

func parse(body io.Reader, base *url.URL) []string {
	doc, err := html.Parse(body)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to parse html: %v", err)
		return nil
	}

	var links []string
	var f func(*html.Node)
	f = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			for _, a := range n.Attr {
				if a.Key == "href" {
					link, err := base.Parse(a.Val)
					if err != nil {
						continue
					}
					if strings.HasPrefix(link.String(), base.String()) {
						links = append(links, link.String())
					}
				}
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}
	f(doc)
	return links
}

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "usage: %s <url>", os.Args[0])
		os.Exit(1)
	}

	crawler, err := NewCrawler(os.Args[1])
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to create crawler: %v", err)
		os.Exit(1)
	}
	crawler.Crawl(os.Args[1], 5)
}
