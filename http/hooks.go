package http

import "github.com/qbhy/goal/contracts"

type RequestBefore struct {
	request contracts.HttpRequest
}

func (this *RequestBefore) Event() string {
	return "REQUEST_BEFORE"
}

func (this *RequestBefore) Sync() bool {
	return true
}
func (this *RequestBefore) Request() contracts.HttpRequest {
	return this.request
}

type RequestAfter struct {
	request contracts.HttpRequest
}

func (this *RequestAfter) Event() string {
	return "REQUEST_AFTER"
}

func (this *RequestAfter) Request() contracts.HttpRequest {
	return this.request
}

func (this *RequestAfter) Sync() bool {
	return true
}
