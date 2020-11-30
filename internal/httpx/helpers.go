package httpx

import (
	"encoding/json"
	"net/http"
	"sync"
)

type errMessage struct {
	Message string `json:"message"`
}

type header struct {
	key   string
	value string
}

type builder struct {
	err           error
	isJson        bool
	item          interface{}
	data          []byte
	status        int
	rw            http.ResponseWriter
	headers       []header
	serializeOnce sync.Once
}

func Response(rw http.ResponseWriter) *builder {
	return &builder{rw: rw}
}

func (b *builder) WithHeader(key, value string) *builder {
	b.headers = append(b.headers, header{key: key, value: value})
	return b
}

func (b *builder) AsJSON() *builder {
	b.isJson = true
	if b.item != nil {
		b.serializeOnce.Do(func() {
			b.data, b.err = json.Marshal(b.item)
		})
	}
	return b.WithHeader("Content-Type", "application/json; charset=utf-8")
}

func (b *builder) WithStatusCode(status int) *builder {
	b.status = status
	return b
}

func (b *builder) WithInterface(v interface{}) *builder {
	b.item = v
	if b.isJson && b.item != nil {
		b.serializeOnce.Do(func() {
			b.data, b.err = json.Marshal(b.item)
		})
	}
	return b
}

func (b *builder) Finish() error {
	_, err := b.Write(b.data)
	return err
}

func (b *builder) Write(data []byte) (int, error) {
	if b.err != nil {
		return -1, b.err
	}
	if b.status == 0 {
		b.status = http.StatusOK
	}
	for _, header := range b.headers {
		b.rw.Header().Add(header.key, header.value)
	}
	b.rw.WriteHeader(b.status)
	return b.rw.Write(data)
}

func Error(rw http.ResponseWriter, status int, err error) error {
	return Response(rw).AsJSON().WithInterface(&errMessage{
		Message: err.Error(),
	}).WithStatusCode(status).Finish()
}

func InternalServerError(rw http.ResponseWriter, err error) error {
	return Error(rw, http.StatusInternalServerError, err)
}
