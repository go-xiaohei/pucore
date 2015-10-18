package app

import (
	"fmt"
	"gopkg.in/inconshreveable/log15.v2"
	"os"
	"os/signal"
	"reflect"
	"runtime"
	"strings"
	"sync"
	"syscall"
	"time"
)

// ctrl+c service interface.
// It supports ctrl+c to stop.
type Service interface {
	Start()
	Stop()
}

// start ctrl+c service.
func Start(s Service) {
	runtime.GOMAXPROCS(runtime.NumCPU())

	signalChan := make(chan os.Signal)
	sName := reflect.TypeOf(s).String()

	// notify system signal terminal
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)
	log15.Info(fmt.Sprintf("%s.Start", strings.Title(sName)))

	// start service, notify signal
	s.Start()
	<-signalChan
	s.Stop()
	WrapWait() // wait for all global goroutines

	log15.Info(fmt.Sprintf("%s.Stop", strings.Title(sName)))
}

var (
	wg          WaitGroup
	wgLogFormat = "goroutine.%s"
)

// add global goroutine with its'name.
func Wrap(funcName string, fn func()) {
	wg.Wrap(funcName, fn)
}

func WrapWait() {
	wg.Wait()
}

type WaitGroup struct {
	sync.WaitGroup
}

func (w *WaitGroup) Wrap(funName string, fn func()) {
	w.Add(1)
	go func() {
		t := time.Now()
		fn()
		w.Done()
		log15.Debug(fmt.Sprintf(wgLogFormat, funName), "goroutine", runtime.NumGoroutine(), "duration", time.Since(t).Seconds()*1000)
		// 强制退出
		runtime.Gosched()
	}()
}
