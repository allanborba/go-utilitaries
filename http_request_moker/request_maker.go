package httprequestmoker

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
)

type ControllerHttp interface {
	Execute(w http.ResponseWriter, r *http.Request)
}

type RequestMocker[T any] struct {
	body         []byte
	controller   ControllerHttp
	req          *http.Request
	w            *httptest.ResponseRecorder
	headersToAdd map[string]string
}

func NewRequestMocker[T any](body []byte, controller ControllerHttp) *RequestMocker[T] {
	return &RequestMocker[T]{body, controller, nil, nil, nil}
}

func (this *RequestMocker[T]) AddHeader(headers map[string]string) *RequestMocker[T] {
	this.headersToAdd = headers
	return this
}

func (this *RequestMocker[T]) Request() {
	this.req = httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(this.body))
	this.w = httptest.NewRecorder()

	for key, value := range this.headersToAdd {
		this.req.Header.Add(key, value)
	}

	this.controller.Execute(this.w, this.req)
}

func (this *RequestMocker[T]) RequestWithResponse() T {
	this.Request()

	var result T

	err := json.NewDecoder(this.w.Body).Decode(&result)
	if err != nil {
		panic(err)
	}

	return result
}
