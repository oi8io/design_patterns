package structure

import "fmt"

type server interface {
	handleRequest(url string, method string) (code int, body string)
}

type nginx struct {
	application       *application
	maxAllowedRequest int
	rateLimiter       map[string]int
}

func (n *nginx) handleRequest(url string, method string) (code int, body string) {
	if !n.checkRateLimiting(url) {
		return 429, "Too many requests"
	}
	return n.application.handleRequest(url, method)
}
func (n *nginx) checkRateLimiting(url string) bool {
	if n.rateLimiter[url] < n.maxAllowedRequest {
		n.rateLimiter[url] += 1
		return true
	}
	return false
}

func newNginx() *nginx {
	return &nginx{application: new(application), maxAllowedRequest: 2, rateLimiter: make(map[string]int)}
}

type application struct {
}

func (a application) handleRequest(url string, method string) (code int, body string) {
	fmt.Println(url,method)
	if url == "/app/status" && method == "GET" {
		return 200, "Ok"
	}

	if url == "/create/user" && method == "POST" {
		return 201, "User Created"
	}
	return 404, "NOT FOUND"
}
