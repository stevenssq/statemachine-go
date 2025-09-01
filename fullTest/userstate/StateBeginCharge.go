package userstate

import (
	"fmt"
	"time"

	. "github.com/stevenssq/statemachine-go/statemachine"
)

type StateBeginCharge struct {
	Stater
}

func (s StateBeginCharge) OnEntry(para interface{}) {
	fmt.Println("StateBeginCharge::onEntry()")
	s.PostEvent("begin charge")
	time.Sleep(1 * time.Second)
}

func (s StateBeginCharge) OnLoop() {
	fmt.Println("StateBeginCharge::OnLoop()")
	time.Sleep(1 * time.Second)
	s.PostEvent("finish charge")
}

func (s StateBeginCharge) OnExit() {
	fmt.Println("StateBeginCharge::OnExit()")
}
