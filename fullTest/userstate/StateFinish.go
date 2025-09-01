package userstate

import (
	"fmt"
	"time"

	. "github.com/stevenssq/statemachine-go/statemachine"
)

type StateFinish struct {
	Stater
}

func (s StateFinish) OnEntry(para interface{}) {
	fmt.Println("StateFinish::onEntry()")
	fmt.Println("job result:", para.(string))
	s.PostEvent("get job")
	time.Sleep(1 * time.Second)
}

func (s StateFinish) OnLoop() {
	fmt.Println("StateFinish::OnLoop()")
}

func (s StateFinish) OnExit() {
	fmt.Println("StateFinish::OnExit()")
}
