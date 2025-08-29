package userstate

import (
	"fmt"
	. "fsm/statemachine"
	"time"
)

type StateGetJob struct {
	Stater
}

func (s StateGetJob) OnEntry(para interface{}) {
	fmt.Println("StateGetJob::onEntry()")
	s.PostEvent("get job")
	time.Sleep(1 * time.Second)
}

func (s StateGetJob) OnLoop() {
	fmt.Println("StateGetJob::OnLoop()")
	s.PostEvent("do job")
	time.Sleep(1 * time.Second)
}

func (s StateGetJob) OnExit() {
	fmt.Println("StateGetJob::OnExit()")
}
