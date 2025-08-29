package userstate

import (
	"fmt"
	. "fsm/statemachine"
	"time"
)

type StateFinishCharge struct {
	Stater
}

func (s StateFinishCharge) OnEntry(para interface{}) {
	fmt.Println("StateFinishCharge::onEntry()")
	s.PostEvent("finish charge")
	time.Sleep(1 * time.Second)
}

func (s StateFinishCharge) OnLoop() {
	fmt.Println("StateFinishCharge::OnLoop()")
	time.Sleep(1 * time.Second)
	s.PostEvent("finish charge")
}

func (s StateFinishCharge) OnExit() {
	fmt.Println("StateFinishCharge::OnExit()")
}
