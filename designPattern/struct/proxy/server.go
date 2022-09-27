package main

type server interface {
	HandleRequest(url string, request string) (status int, respond string)
}
