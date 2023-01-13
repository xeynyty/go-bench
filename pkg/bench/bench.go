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
type Result struct {
	ReqCount        uint64  `json:"req_count"`
	ErrCount        uint64  `json:"err_count"`
	PercentOfErrors float32 `json:"percent_of_errors"`
}

// Start of bench
func (b *Bench) Start() {
	go b.bench()
}

// Stop of bench
func (b *Bench) Stop() *Result {
	b.contextCancel()
	return &Result{
		ReqCount:        b.reqCount,
		ErrCount:        b.errCount,
		PercentOfErrors: percentOfErrors(b.reqCount, b.errCount),
	}
}

// New create Bench struct object
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

// bench is a main func in lib
func (b *Bench) bench() {
	defer b.contextCancel()

	var waitTime time.Duration
	if b.requestPerSecond != 0 {
		waitTime = time.Second / time.Duration(b.requestPerSecond)
	}
	if b.requestPerSecond == 0 {
		waitTime = time.Nanosecond
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

// request is only 1 req in func for bench func
func (b *Bench) request() {
	b.reqCount += 1
	_, err := http.Get(b.url)
	if err != nil {
		b.errCount += 1
	}
}

func percentOfErrors(req, err uint64) float32 {
	return (float32(err) / float32(req)) * 100.0
}
