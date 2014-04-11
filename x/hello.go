package x

import (
	"circuit/use/circuit"
	"fmt"
	"sync/atomic"
	"time"
)

type App struct{}

type RemoteWorker struct {
	counter int64
}

func (self *RemoteWorker) Bounce() int64 {
	return atomic.AddInt64(&self.counter, 1)
}

func (App) Main(suffix string) circuit.XPerm {
	circuit.RunInBack(func() {
		time.Sleep(5 * time.Second)
		fmt.Printf("Hello %s\n", suffix)
	})

	fmt.Printf(`about to return`)

	return circuit.PermRef(&RemoteWorker{})
}

func init() {
	circuit.RegisterFunc(App{})
	circuit.RegisterValue(&RemoteWorker{})
}
