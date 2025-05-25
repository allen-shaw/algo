package sdk

import "net/http"

type LowLevelClient struct {
	url string
	c   *http.Client
}

func NewLowLevelClient(url string) *LowLevelClient {
	return &LowLevelClient{
		url: url,
		c:   http.DefaultClient, // TODO
	}
}

func (llc *LowLevelClient) init() error {
	// TODO: 鉴权
	return nil
}

type Request struct {
	messages []Message
	model    string
	// ...
}

func newRequest(messages []Message, model string) *Request {
	return &Request{messages: messages, model: model}
}

func (r *Request) ToJson() string {
	panic("unimplemented")
}

type Stream interface {
	Get() (Response, error)
	GetResponse() <-chan Response
}

type stream struct {
	stream *http.Response
}

func NewStream(res *http.Response) *stream {
	return &stream{stream: res}
}

// func (s *Stream[T]) get() T {
// 	// return s.stream.
// }

type Response struct {
}

func (c *LowLevelClient) send(request *Request) (Stream, error) {
	req := request.ToJson()
	resp, err := c.c.Post(c.url, "", req)
	if err != nil {
		// log.error()
		return nil, err
	}

	return NewStream(resp), nil
}
