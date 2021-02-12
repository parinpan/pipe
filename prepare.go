package pipe

import (
	"errors"
	"reflect"
)

var (
	errPreparePassArgsLimitExceeded = errors.New("calling pipe.Pass() more than once")
)

type prepare struct {
	applyFn        *applyFn
	compoundResult reflect.Value

	sequence            int
	passArgsFlagCounter int
}

func (prepare *prepare) fnArgs() []reflect.Value {
	var args []reflect.Value

	defer func() {
		prepare.passArgsFlagCounter = 0
	}()

	if prepare.sequence > 0 && len(prepare.applyFn.args) == 0 {
		args = append(args, prepare.compoundResult)
		return args
	}

	for _, arg := range prepare.applyFn.args {
		if _, ok := arg.(PassArgs); ok {
			prepare.passArgsFlagCounter++
			args = append(args, prepare.compoundResult)
		} else {
			args = append(args, reflect.ValueOf(arg))
		}

		if prepare.passArgsFlagCounter > 1 {
			panic(errPreparePassArgsLimitExceeded)
		}
	}

	if prepare.sequence > 0 && prepare.passArgsFlagCounter == 0 {
		args = append(args, prepare.compoundResult)
	}

	return args
}
