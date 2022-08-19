package main

import (
	"os"
	"fmt"
	"net/http"
	"time"
	"flag"
	"strings"
)

type Config struct {
	url string
	agent string
}

var UserAgents = map[string]string{
	"chrome": "Mozilla/5.0 (Macintosh; Intel Mac OS X 12_5) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/104.0.0.0 Safari/537.36",
	"safari": "Mozilla/5.0 (Macintosh; Intel Mac OS X 12_5) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/15.6 Safari/605.1.15",
	"firefox": "Mozilla/5.0 (Macintosh; Intel Mac OS X 12.5; rv:103.0) Gecko/20100101 Firefox/103.0",
}

func add_schema(site string) string {
	schema_len := len("https")
	if len(site) > schema_len {
		site = "https://" + site
	}
	if site[:schema_len] != "https" {
		site = "https://" + site
	}
	return site
}

func createClient() *http.Client {
	client := &http.Client{
		Timeout: time.Second * 10,
	}
	return client
}

func main() {
	var cfg Config
	flag.StringVar(&cfg.url, "url", "", "URL to fetch headers from")
	flag.StringVar(&cfg.agent, "agent", "safari", "User-agent to test")
	flag.Parse()

	client := createClient()

	if cfg.url == "" {
		fmt.Println("URL cannot be empty - use --url to specify which url to request headers from")
		os.Exit(0)
	}

	site := add_schema(cfg.url)
	req, err := http.NewRequest("GET", site, nil)
	if err != nil {
		fmt.Println(err)
	}

	key := strings.ToLower(cfg.agent)
	usr_agent := UserAgents[key]
	req.Header.Set("user-agent", usr_agent)

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}

	if resp.StatusCode != 200 {
		fmt.Println("Status code not 200")
		os.Exit(1)
	}

	fmt.Println("_________________________________________________________")
	fmt.Printf("\t\t%s\n", site)
	fmt.Println("_________________________________________________________")
	fmt.Println("{");
	for key, val :=range resp.Header {
		fmt.Printf("\t%s: %s\n", key, val[0])
	}
	fmt.Println("}");
}
