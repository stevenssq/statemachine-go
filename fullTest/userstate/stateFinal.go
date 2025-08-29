package userstate

import (
	"fmt"
	. "fsm/statemachine"
)

type StateFinal struct {
	State
}

func (s StateFinal) OnEntry(para interface{}) {
	fmt.Println("StateFinal::OnEntry")
}

func (s StateFinal) OnExit() {
	fmt.Println("StateFinal::OnExit")
}
