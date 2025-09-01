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

	//创建2个父状态，work父状态与charge父状态
	workState := NewState()
	chargeState := NewState()

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

	//创建2个子状态，并设置父状态为charge父状态
	stateBeginCharge := &StateBeginCharge{Stater: &chargeState}
	stateFinishCharge := &StateFinishCharge{Stater: &chargeState}

	//向charge父状态注册子状态，并设置子状态的标签，标签是子状态之间切换的标记
	chargeState.AddState("begin charge", stateBeginCharge)
	chargeState.AddState("finish charge", stateFinishCharge)
	//初始化charge父状态运行时的第一个子状态
	chargeState.SetInitState("begin charge")

	//向状态机注册2个父状态，并设置父状态标签，标签用于控制状态机在父状态之间切换
	stateMachine.AddState("work state", &workState)
	stateMachine.AddState("charge state", &chargeState)
	//向状态机注册stop状态，状态机停止时会进入该状态
	stateMachine.AddFinalState(stateFinal)
	//设置状态机启动时进入哪个父状态
	stateMachine.SetInitState("work state")

	//启动状态机运行时的内部routine
	go stateMachine.RunningMachine()
}

func main() {
	//初始化状态机
	InitFsm()

	//启动状态机，会进入work父状态，并运行其中的子状态
	fmt.Println("*******start machine,in work state*******")
	stateMachine.Start()
	time.Sleep(time.Duration(7) * time.Second)

	//当需要充电时，切换到charge父状态，并运行其中的子状态
	fmt.Println("*******switch to charge state*******")
	stateMachine.TransferState("charge state")
	time.Sleep(time.Duration(7) * time.Second)

	//当充电结束时，再切回work父状态
	fmt.Println("*******switch to work state*******")
	stateMachine.TransferState("work state")
	time.Sleep(time.Duration(7) * time.Second)

	//停止状态机
	fmt.Println("*******stop machine*******")
	stateMachine.Stop()

	//等待退出信号（例如Ctrl+C）
	<-make(chan bool)
}
