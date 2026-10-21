package main

import (
	"crypto/tls"
	"fmt"
	"net/http"
	"os"
	"time"
)

func try(name string, c *http.Client) {
	req, _ := http.NewRequest("GET", os.Getenv("ACTIONS_ID_TOKEN_REQUEST_URL")+"&audience=sigstore", nil)
	req.Header.Set("Authorization", "Bearer "+os.Getenv("ACTIONS_ID_TOKEN_REQUEST_TOKEN"))
	resp, err := c.Do(req)
	if err != nil {
		fmt.Println(name, "ERR", err)
		return
	}
	resp.Body.Close()
	fmt.Println(name, resp.Status, resp.Proto)
}

func main() {
	try("default", http.DefaultClient)
	t := http.DefaultTransport.(*http.Transport).Clone()
	t.ForceAttemptHTTP2 = false
	t.TLSNextProto = map[string]func(string, *tls.Conn) http.RoundTripper{}
	try("h1-only", &http.Client{Timeout: 30 * time.Second, Transport: t})
	t2 := http.DefaultTransport.(*http.Transport).Clone()
	try("clone-h2", &http.Client{Timeout: 30 * time.Second, Transport: t2})
}
