package bench

import (
	"context"
	"net/http"
	"time"
)

type Bench struct {
	url              string
	context          context.Context
	contextCancel    context.CancelFunc
	reqCount         uint64
	errCount         uint64
	requestPerSecond uint16
}

func (b *Bench) Start() {
	go b.bench()
}
func (b *Bench) Stop() {
	b.contextCancel()
	println("err:", b.errCount, "\nreq", b.reqCount)
}

func (b *Bench) bench() {
	defer b.contextCancel()

	var waitTime time.Duration
	if b.requestPerSecond != 0 {
		waitTime = time.Second / time.Duration(b.requestPerSecond)
	}
	if b.requestPerSecond == 0 {
		waitTime = time.Second / 100
	}

	timer := time.NewTimer(waitTime)

exit:
	for {
		select {
		case <-b.context.Done():
			break exit
		default:
			select {
			case <-timer.C:
				go b.request()
				timer.Reset(waitTime)
			}
		}

	}
}

func (b *Bench) request() {
	b.reqCount += 1
	_, err := http.Get(b.url)
	if err != nil {
		b.errCount += 1
	}
}

func New(url string, reqPerSec uint16) *Bench {
	ctx, cancel := context.WithCancel(context.Background())
	return &Bench{
		url:              url,
		context:          ctx,
		contextCancel:    cancel,
		reqCount:         0,
		errCount:         0,
		requestPerSecond: reqPerSec,
	}
}
