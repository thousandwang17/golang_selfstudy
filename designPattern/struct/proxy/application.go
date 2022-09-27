package main

type application struct{}

func NewApplication() server {
	return &application{}
}

func (a *application) HandleRequest(url, request string) (httpCode int, respond string) {
	if url == "/app/status" && request == "GET" {
		return 200, "ok"
	}

	if url == "/create/user" && request == "post" {
		return 201, "created"
	}

	return 404, "not found"
}
