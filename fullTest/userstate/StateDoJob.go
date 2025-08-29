package userstate

import (
	"fmt"
	. "fsm/statemachine"
	"time"
)

type StateDoJob struct {
	Stater
}

func (s StateDoJob) OnEntry(para interface{}) {
	fmt.Println("StateDoJob::onEntry()")
	s.PostEvent("do job")
	time.Sleep(1 * time.Second)
}

func (s StateDoJob) OnLoop() {
	fmt.Println("StateDoJob::OnLoop()")
	time.Sleep(1 * time.Second)
	s.PostEvent("job finish", "success")
}

func (s StateDoJob) OnExit() {
	fmt.Println("StateDoJob::OnExit()")
}
