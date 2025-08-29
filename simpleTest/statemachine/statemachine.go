package statemachine

import (
	"fmt"
	"sync"
	"time"
)

const (
	INIT    = 0
	RUNNING = 1
	STOPPED = 2
)

type S_TransEvent struct {
	cmd       string
	nextState Stater
}

type StateMachine struct {
	State
	finalState    Stater
	machineState  int
	transVec      []S_TransEvent
	transVecMutex sync.Mutex
}

func (s *StateMachine) AddFinalState(finalState Stater) {
	s.finalState = finalState
}

func (s *StateMachine) setTransEvent(transEvent S_TransEvent) {
	s.transVecMutex.Lock()
	s.transVec = append(s.transVec, transEvent)
	s.transVecMutex.Unlock()
}

func (s *StateMachine) getTransEvent() S_TransEvent {
	s.transVecMutex.Lock()
	transEvent := s.transVec[0]
	s.transVec = s.transVec[1:]
	s.transVecMutex.Unlock()

	return transEvent
}

func (s *StateMachine) executeState() {
	if s.currentState == nil {
		fmt.Println("null point currentState!")
		time.Sleep(time.Duration(1) * time.Second)
		return
	}

	if s.currentState != s.lastState {
		if s.lastState != nil {
			if s.lastState == s.finalState {
				s.lastState.OnExit()
			} else {
				s.lastState.StateExit()
			}
		}
		s.lastState = s.currentState
	}

	s.currentState.ExecuteState()
}

func (s *StateMachine) TransferState(transState string) {
	value, ok := s.StateMap[transState]

	if ok {
		transEvent := S_TransEvent{"transfer", value}
		s.setTransEvent(transEvent)
	} else {
		transEvent := S_TransEvent{"transfer", nil}
		s.setTransEvent(transEvent)
	}
}

func (s *StateMachine) executeFinalState() {
	if s.finalState != nil {
		s.finalState.OnEntry("")
		s.lastState = s.finalState
	}
}

func (s StateMachine) getMachineState() int {
	return s.machineState
}

func (s *StateMachine) setMachineState(state int) {
	s.machineState = state
}

func (s *StateMachine) transferOperate() {
	if 0 == len(s.transVec) {
		return
	}

	transEvent := s.getTransEvent()

	if transEvent.cmd == "start" {
		if RUNNING == s.getMachineState() {
			return
		}
		s.autoSetInitState()
		s.currentState = s.initState
		if s.currentState != nil {
			s.currentState.StateEntry()
		}

		s.setMachineState(RUNNING)
	} else if transEvent.cmd == "stop" {
		if STOPPED == s.getMachineState() {
			return
		}

		if s.lastState != nil {
			s.lastState.StateExit()
		}

		s.lastState = nil
		s.executeFinalState()
		s.setMachineState(STOPPED)
	} else {
		if STOPPED == s.getMachineState() {
			fmt.Println("state machine stopped, ignore transfer")
			return
		}

		if transEvent.nextState == s.currentState {
			return
		}

		s.lastState = s.currentState
		s.currentState = transEvent.nextState
		if s.currentState != nil {
			s.currentState.StateEntry()
		}
	}
}

func (s *StateMachine) RunningMachine() {
	for {
		time.Sleep(time.Duration(10) * time.Microsecond)
		s.transferOperate()

		switch s.machineState {
		case RUNNING:
			s.executeState()
		case STOPPED:
			time.Sleep(time.Duration(10) * time.Microsecond)
		}
	}
}

func (s *StateMachine) Start() {
	transEvent := S_TransEvent{"start", nil}
	s.setTransEvent(transEvent)
}

func (s *StateMachine) Stop() {
	transEvent := S_TransEvent{"stop", nil}
	s.setTransEvent(transEvent)
}

func (s *StateMachine) autoSetInitState() {
	if s.initState == nil && len(s.StateMap) == 1 {
		for _, v := range s.StateMap {
			s.initState = v
			break
		}
	}
}
