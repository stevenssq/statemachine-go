package statemachine

import (
	"fmt"
	"time"
)

type Stater interface {
	OnEntry(para interface{})
	OnExit()
	OnLoop()
	StateEntry()
	StateExit()
	ExecuteState()
	PostEvent(event string, paras ...interface{})
}

type S_EVENT struct {
	nextState string
	para      interface{}
}

type State struct {
	initState    Stater
	currentState Stater
	lastState    Stater
	currentEvent S_EVENT
	lastEvent    S_EVENT
	eventList    []S_EVENT
	StateMap     map[string]Stater
}

func NewState() State {
	return State{StateMap: make(map[string]Stater)}
}

func (s *State) stateInit() {
	s.currentState = nil
	s.lastState = nil
	s.currentEvent = S_EVENT{nextState: "", para: ""}
	s.lastEvent = S_EVENT{nextState: "", para: ""}
	s.eventList = s.eventList[0:0]
}

func (s State) OnEntry(para interface{}) {

}

func (s State) OnLoop() {

}

func (s State) OnExit() {

}

func (s *State) StateEntry() {
	s.stateInit()
	s.currentState = s.initState
}

func (s *State) StateExit() {
	if s.lastState != nil {
		s.lastState.OnExit()
	}
}

func (s *State) SetInitState(initState string) {
	if value, ok := s.StateMap[initState]; ok {
		s.initState = value
	}
}

func (s *State) ExecuteState() {
	if s.currentState == nil {
		fmt.Println("currentState is NULL")
		time.Sleep(time.Duration(1) * time.Second)
		return
	}

	if s.lastState != s.currentState && s.lastState != nil {
		s.lastState.OnExit()
	}

	if s.lastState == s.currentState {
		s.currentState.OnLoop()
	} else {
		s.currentState.OnEntry(s.currentEvent.para)
	}

	s.lastState = s.currentState
	s.lastEvent = s.currentEvent

	event := s.getNextEvent()
	value, ok := s.StateMap[event.nextState]
	if ok {
		s.currentState = value
		s.currentEvent = event
	} else {
		s.currentState = nil
		time.Sleep(time.Duration(1) * time.Second)
	}
}

func (s *State) AddState(transEvent string, nextState Stater) {
	s.StateMap[transEvent] = nextState
}

func (s *State) PostEvent(event string, paras ...interface{}) {
	tmpEvent := S_EVENT{nextState: event}
	if len(paras) > 0 {
		tmpEvent.para = paras[0]
	}

	s.eventList = append(s.eventList, tmpEvent)
}

func (s *State) getNextEvent() S_EVENT {
	event := S_EVENT{"NULL event", ""}

	if len(s.eventList) == 0 {
		return event
	}

	for {
		if len(s.eventList) > 0 {
			event = s.eventList[0]
			s.eventList = s.eventList[1:]
		} else {
			break
		}
	}

	return event
}
