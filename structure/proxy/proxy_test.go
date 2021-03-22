package structure

import (
	"fmt"
	"testing"
)

func TestProxy(t *testing.T) {
	ngx := newNginx()
	appStatusURL := "/app/status"
	createuserURL := "/create/user"
	const (
		MethodGET  = "GET"
		MethodPost = "POST"
	)
	code, body := ngx.handleRequest(appStatusURL, MethodGET)
	fmt.Printf("\nUrl: %s\nHttpCode: %d\nBody: %s\n", appStatusURL, code, body)
	code, body = ngx.handleRequest(appStatusURL, MethodGET)
	fmt.Printf("\nUrl: %s\nHttpCode: %d\nBody: %s\n", appStatusURL, code, body)
	code, body = ngx.handleRequest(appStatusURL, MethodGET)
	fmt.Printf("\nUrl: %s\nHttpCode: %d\nBody: %s\n", appStatusURL, code, body)
	code, body = ngx.handleRequest(appStatusURL, MethodGET)
	fmt.Printf("\nUrl: %s\nHttpCode: %d\nBody: %s\n", appStatusURL, code, body)
	code, body = ngx.handleRequest(createuserURL, MethodPost)
	fmt.Printf("\nUrl: %s\nHttpCode: %d\nBody: %s\n", createuserURL, code, body)
	code, body = ngx.handleRequest(createuserURL, MethodPost)
	fmt.Printf("\nUrl: %s\nHttpCode: %d\nBody: %s\n", createuserURL, code, body)
	code, body = ngx.handleRequest(createuserURL, MethodPost)
	fmt.Printf("\nUrl: %s\nHttpCode: %d\nBody: %s\n", createuserURL, code, body)
	code, body = ngx.handleRequest(createuserURL, MethodPost)
	fmt.Printf("\nUrl: %s\nHttpCode: %d\nBody: %s\n", createuserURL, code, body)
	code, body = ngx.handleRequest(createuserURL, MethodPost)
	fmt.Printf("\nUrl: %s\nHttpCode: %d\nBody: %s\n", createuserURL, code, body)
}
