package bench

import (
	"context"
	"github.com/valyala/fasthttp"
	"sync"
	"time"
)

type Bench struct {
	url                 string
	context             context.Context
	contextCancel       context.CancelFunc
	startTime           time.Time
	reqCount            uint64
	errCount            uint64
	requestPerSecond    uint16
	responseTimePointer *[]float32
}
type Result struct {
	ReqCount            uint64  `json:"req_count"`
	ErrCount            uint64  `json:"err_count"`
	AverageResponseTime float32 `json:"average_response_time_ms"`
	MaxResponseTime     float32 `json:"max_response_time_ms"`
	MinResponseTime     float32 `json:"min_response_time_ms"`
	TimeOfBench         float32 `json:"time_of_bench_sec"`
	PercentOfErrors     float32 `json:"percent_of_errors"`
}

var (
	wg = sync.WaitGroup{}
)

// Start of bench
func (b *Bench) Start() {
	go func() {
		var (
			responseTimeChan = make(chan float32, b.requestPerSecond*10)
		)

		b.startTime = time.Now()
		b.responseTime(responseTimeChan)

		defer func() {
			b.contextCancel()
			close(responseTimeChan)
		}()

		var waitTime time.Duration
		if b.requestPerSecond != 0 {
			waitTime = time.Second / time.Duration(b.requestPerSecond)
		}
		if b.requestPerSecond == 0 {
			waitTime = time.Nanosecond
		}

		timer := time.NewTimer(waitTime)

		for {
			select {
			case <-b.context.Done():
				break
			default:
				select {
				case <-timer.C:
					wg.Add(1)
					go b.request(responseTimeChan)
					timer.Reset(waitTime)
				}
			}

		}
	}()
}

// Stop of bench
func (b *Bench) Stop() *Result {
	b.contextCancel()
	wg.Wait()
	return &Result{
		ReqCount:            b.reqCount,
		ErrCount:            b.errCount,
		PercentOfErrors:     percentOfErrors(&b.reqCount, &b.errCount),
		AverageResponseTime: averageResponseTime(*b.responseTimePointer),
		MaxResponseTime:     b.maxResponseTime(),
		TimeOfBench:         b.benchTime(),
		MinResponseTime:     b.minResponseTime(),
	}
}

// request is only 1 req in func for bench func
func (b *Bench) request(ch chan<- float32) {
	select {
	case <-b.context.Done():
		wg.Done()
		return
	default:
		sendTime := time.Now()
		b.reqCount += 1

		// TODO add status code in result

		_, _, err := fasthttp.Get(nil, b.url)
		if err != nil {
			b.errCount += 1
		}
		ch <- float32(time.Now().Sub(sendTime).Milliseconds())
	}
}

// responseTime was created to collect the response time from each
// request through the responseTimeChan channel,
// at which time is written from request
func (b *Bench) responseTime(ch <-chan float32) {
	go func() {
		for {
			select {
			case <-b.context.Done():
				wg.Done()
				return
			default:
				for data := range ch {
					*b.responseTimePointer = append(*b.responseTimePointer, data)
					wg.Done()
				}
			}
		}
	}()
}

// benchTime returns the time elapsed from the start
// of the benchmark to the execution of Stop
func (b *Bench) benchTime() float32 {
	return float32(time.Now().Sub(b.startTime).Seconds())
}
func (b *Bench) maxResponseTime() float32 {
	var max float32 = 0
	for _, item := range *b.responseTimePointer {
		if item > max {
			max = item
		}
	}
	return max
}
func (b *Bench) minResponseTime() float32 {
	var min float32 = 100000.0
	for _, item := range *b.responseTimePointer {
		if item < min {
			min = item
		}
	}
	return min
}

// New create Bench struct object
func New(url string, reqPerSec uint16) *Bench {
	ctx, cancel := context.WithCancel(context.Background())
	responseTime := make([]float32, 0, reqPerSec*10)
	return &Bench{
		url:                 url,
		context:             ctx,
		contextCancel:       cancel,
		reqCount:            0,
		errCount:            0,
		requestPerSecond:    reqPerSec,
		responseTimePointer: &responseTime,
	}
}

func percentOfErrors(req, err *uint64) float32 {
	if *req == 0 || *err == 0 {
		return 0
	}
	return (float32(*err) / float32(*req)) * 100.0
}
func averageResponseTime(slice []float32) float32 {
	var average float32 = 0
	for _, item := range slice {
		average += item
	}
	average = average / float32(len(slice))
	return average
}
