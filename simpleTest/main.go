package main

import (
	"fmt"
	. "fsm/userstate"
	"time"

	. "github.com/stevenssq/statemachine-go/statemachine"
)

var stateMachine *StateMachine

func InitFsm() {
	//创建状态机
	stateMachine = &StateMachine{State: NewState()}

	//创建父状态
	workState := NewState()
	//创建stop状态，这是一个特殊的状态，当停止状态机时会进入该状态
	stateFinal := &StateFinal{State: NewState()}

	//创建3个子状态，并设置父状态为work父状态
	stateGetJob := &StateGetJob{Stater: &workState}
	stateDoJob := &StateDoJob{Stater: &workState}
	stateFinish := &StateFinish{Stater: &workState}

	//向work父状态注册子状态，并设置子状态的标签，标签是子状态之间切换的标记
	workState.AddState("get job", stateGetJob)
	workState.AddState("do job", stateDoJob)
	workState.AddState("job finish", stateFinish)
	//初始化work父状态运行时的第一个子状态
	workState.SetInitState("get job")

	//向状态机注册父状态与stop状态
	stateMachine.AddState("work state", &workState)
	stateMachine.AddFinalState(stateFinal)
	//初始化状态机启动时进入work父状态，当只有一个父状态时，该行可以省略
	//stateMachine.SetInitState("work state")

	//启动状态机运行时的内部routine
	go stateMachine.RunningMachine()
}

func main() {
	//初始化状态机
	InitFsm()

	//启动状态机
	fmt.Println("*******start machine*******")
	stateMachine.Start()
	time.Sleep(time.Duration(7) * time.Second)

	//停止状态机
	fmt.Println("*******stop machine*******")
	stateMachine.Stop()
	time.Sleep(time.Duration(2) * time.Second)

	//再次启动状态机
	fmt.Println("*******start machine*******")
	stateMachine.Start()

	//等待退出信号（例如Ctrl+C）
	<-make(chan bool)
}
