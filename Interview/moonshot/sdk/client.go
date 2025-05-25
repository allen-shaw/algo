package sdk

const (
	url string = "..."
)

type Role string

type Model struct {
	name string
}

type Message struct {
	id      string
	Role    Role
	Content string
}

type Reply struct {
}

type Session interface {
	ID() string
	SendMessage(messages []Message) (Stream, error)
}

type session struct {
	id          string
	model       Model
	temperature float32
	client      *LowLevelClient
}

func newSession(model Model, temperature float32, client *LowLevelClient) (*session, error) {
	id := genUUID()
	return &session{
		id:          id,
		model:       model,
		temperature: temperature,
		client:      client,
	}, nil
}

func (s *session) ID() string {
	return s.id
}

func (s *session) SendMessage(messages []Message) (Stream, error) {
	req := newRequest(messages, s.model.name)
	return s.client.send(req)
}

type Client interface {
	NewSession() (Session, error)
}

type cilent struct {
	id  string
	llc *LowLevelClient
}

func (c *cilent) NewSession(model Model, temperature float32) (Session, error) {
	return newSession(model, temperature, c.llc)
}

func NewClient() (Client, error) {
	id := genUUID()
	llc := NewLowLevelClient(url)
	err = llc.init()
	c := &client{id: id, llc: llc}
	return c, nil
}
