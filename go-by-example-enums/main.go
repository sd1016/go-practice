package main

import "fmt"

type ServerState int

const (
	StateIdle = iota
	StateConnected
	StateError
	StateRetrying
)

// By implementing the fmt.Stringer interface, values of ServerState
// can be printed out or converted to strings.
var stateName = map[ServerState]string{
	StateIdle:      "idle",
	StateConnected: "connected",
	StateError:     "error",
	StateRetrying:  "retrying",
}

func (ss ServerState) String() string {
	return stateName[ss]
}

// If we have a value of type int, we cannot pass it to transition -
// the compiler will complain about type mismatch.
// This provides some degree of compile-time type safety for enums.
// go run main.go
// # command-line-arguments
// ./main.go:31:20: cannot use value1 (variable of type int) as ServerState value in argument to transition
// func main() {
// 	var value1 int = 2
// 	ns := transition(StateIdle)
// 	fmt.Println(ns)
// 	ns2 := transition(value1)
// 	fmt.Println(ns2)
// }

func main() {
	ns := transition(StateIdle)
	fmt.Println(ns)
	ns2 := transition(ns)
	fmt.Println(ns2)
}

// transition emulates a state transition for a server;
// it takes the existing state and returns a new state.
func transition(s ServerState) ServerState {
	switch s {
	case StateIdle:
		return StateConnected
	case StateConnected, StateRetrying:
		return StateIdle
	case StateError:
		return StateError
	default:
		panic(fmt.Errorf("unknown state: %s", s))
	}

	return StateConnected
}
