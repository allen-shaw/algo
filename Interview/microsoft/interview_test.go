package microsoft

// 假设我们有一个TTS服务的实时请求流，我们需要实时监控过去 k 毫秒内耗时最长的那个请求，以便触发报警。

type Request struct {
	id        int
	latency   int
	timestamp int
}

type SliceWindowAlter struct {
	k      int
	stepMs int
	window []*Request
}

func NewSliceWindowAlter(k int, step int) SliceWindowAlter {
	a := SliceWindowAlter{
		k:      k,
		stepMs: step,
		window: make([]*Request, k/step),
	}
	return a
}

func (swa *SliceWindowAlter) AddRequest(req *Request) {
	swa.clean(req.timestamp)
	for len(swa.window) > 0 && swa.window[len(swa.window)-1].latency <= req.latency {
		swa.window = swa.window[:len(swa.window)-1]
	}
	swa.window = append(swa.window, req)

	// if len(swa.window) > 0 && req.timestamp-swa.window[len(swa.window)-1].timestamp < swa.stepMs {
	// 	if req.latency > swa.window[len(swa.window)-1].latency {
	// 		swa.window[len(swa.window)-1] = req
	// 	}
	// } else {
	// 	swa.window = append(swa.window, req)
	// }

	// // O(k/step) -> O(log(k/step))
	// var alterReq *Request
	// for _, req := range swa.window {
	// 	if alterReq == nil {
	// 		alterReq = req
	// 	} else {
	// 		if req.latency > alterReq.latency {
	// 			alterReq = req
	// 		}
	// 	}
	// }

	if len(swa.window) > 0 {
		swa.SendAlert(swa.window[0])
	}
}

// O(n) -> O(1)
func (swa *SliceWindowAlter) clean(now int) {
	if len(swa.window) == 0 {
		return
	}
	idx := 0
	for id, req := range swa.window {
		if now-req.timestamp > swa.k {
			idx = id
		} else {
			break
		}
	}
	idx += 1
	if idx >= len(swa.window) {
		swa.window = make([]*Request, 0)
	} else {
		swa.window = swa.window[idx:]
	}
}

func (swa *SliceWindowAlter) SendAlert(req *Request) {
	// TODO
}
