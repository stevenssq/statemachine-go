package userstate

import (
	"fmt"

	. "github.com/stevenssq/statemachine-go/statemachine"
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
