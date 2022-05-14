# go-event
a simple event system


Sample:

```Go

package main

import (
	"context"
	"fmt"
	"time"

	"github.com/benny502/go-event/event"
	"github.com/benny502/go-event/manager"
	"golang.org/x/exp/slices"
)

const (
	TestEvent event.EventType = iota
)

type MyListener struct {
}

func (*MyListener) Handler(e event.Event) {
	fmt.Println(e.Type)
}

func main() {

	manager.GetInstance().Register(manager.NewEvent(context.Background(), TestEvent), &MyListener{})
	manager.GetInstance().Start()

	manager.GetInstance().Send(manager.NewEvent(context.Background(), TestEvent))

	<-time.After(10 * time.Second)

}

```
