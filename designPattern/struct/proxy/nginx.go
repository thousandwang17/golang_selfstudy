package main

type Nginx struct {
	application     server
	maxAllowRequest int
	rateLimiter     map[string]int
}

func NewNginx() server {
	return &Nginx{
		application:     NewApplication(),
		maxAllowRequest: 2,
		rateLimiter:     map[string]int{},
	}
}

func (n *Nginx) HandleRequest(url, request string) (httpCode int, respond string) {

	n.rateLimiter[url]++

	if n.rateLimiter[url] > n.maxAllowRequest {
		return 429, "api limited "
	}

	status, respod := n.application.HandleRequest(url, request)
	n.rateLimiter[url]--

	return status, respod
}
