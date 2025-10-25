package kimi

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"sync"
	"sync/atomic"
	"time"
)

const maxBatchSize = 128

type Image struct {
	Data []byte
}

type Result struct {
	Faces []Rect `json:"faces"`
}

type Rect struct {
	// omit fields
}

// FaceDetectInBatch is a singleton method, we use a mutex to simulate it.
var mu sync.Mutex

// FaceDetectInBatch is a mock function that simulates face detection on GPU
func FaceDetectInBatch(images []Image) []Result {
	if len(images) > maxBatchSize {
		panic("batch size exceeded")
	}
	mu.Lock()
	defer mu.Unlock()
	// Simulate processing time
	time.Sleep(100 * time.Millisecond)

	results := make([]Result, len(images))
	for i := range results {
		results[i] = Result{
			Faces: []Rect{{}, {}},
		}
	}
	return results
}

var idGen atomic.Int64

type Request struct {
	img   Image
	id    int64
	notiQ chan Result
}

func NewRequest(img Image, id int64) *Request {
	return &Request{
		img:   img,
		id:    id,
		notiQ: make(chan Result, 1),
	}
}

type Processer struct {
	ctx    context.Context
	cancel context.CancelCauseFunc
	reqQ   chan *Request
}

func newProcesser() Processer {
	ctx, cancel := context.WithCancelCause(context.Background())
	p := Processer{
		ctx:    ctx,
		cancel: cancel,
		reqQ:   make(chan *Request, maxBatchSize),
	}
	go p.run()

	return p
}

func (p *Processer) run() {
	images := make([]Image, 0)
	m := make(map[*Image]*Request)

	for {
		select {
		case req, ok := <-p.reqQ:
			if !ok {
				return
			}
			images = append(images, req.img)
			imgPtr := &req.img
			m[imgPtr] = req

			if len(images) == maxBatchSize {
				results := FaceDetectInBatch(images)
				for i := range images {
					imgPtr := &images[i]
					request := m[imgPtr]
					request.notiQ <- results[i]
				}

				images = make([]Image, 0)
				m = make(map[*Image]*Request)
			}
		case <-p.ctx.Done():
			return
		}
	}
}

var processer = newProcesser()

func detectHandler(w http.ResponseWriter, req *http.Request) {
	data, _ := io.ReadAll(req.Body)
	img := Image{Data: data}
	id := idGen.Add(1)
	r := NewRequest(img, id)
	processer.reqQ <- r
	result := <-r.notiQ

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}

func main() {
	http.HandleFunc("/faces:detect", detectHandler)

	fmt.Println("Server is running on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
