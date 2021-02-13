package pipe

import (
	"errors"
	"fmt"
	"reflect"
	"runtime"
)

var (
	errPrepareArgsDifferentType                   = errors.New("given argument vs actual argument has different type")
	errPrepareArgsLimitExceeded                   = errors.New("number of arguments exceeded")
	errPreparePassArgsLimitExceeded               = errors.New("calling pipe.Pass() more than once")
	errPreparePassArgsAcceptDifferentArgumentType = errors.New("pipe.Pass(...) function accepts different compound return value in the argument")
)

type prepare struct {
	applyFn        *applyFn
	compoundResult reflect.Value

	passArgs            *passArgs
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
		prepare.checkArgumentsNumber(args)
		prepare.checkArgumentsType(args)
		return args
	}

	for _, arg := range prepare.applyFn.args {
		if argValue, ok := arg.(*passArgs); ok {
			prepare.passArgsFlagCounter++
			prepare.passArgs = argValue
			args = append(args, prepare.prepareCompoundResult(prepare.compoundResult))
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

	prepare.checkArgumentsNumber(args)
	prepare.checkArgumentsType(args)

	return args
}

func (prepare *prepare) prepareCompoundResult(compoundResult reflect.Value) reflect.Value {
	if _, isEmptyFnCandidate := prepare.passArgs.fnCandidate.Interface().(emptyFnCandidate); isEmptyFnCandidate {
		return compoundResult
	}

	prepare.checkPassArgs()

	return prepare.passArgs.fnCandidate.Call([]reflect.Value{compoundResult})[0]
}

func (prepare *prepare) checkPassArgs() {
	panicFn := func(err error) {
		fnName := runtime.FuncForPC(prepare.applyFn.fnCandidateValue.Pointer()).Name()
		panic(fmt.Sprintf("sequence:%d | fnName:%s | err:%v", prepare.sequence+1, fnName, err))
	}

	if err := prepare.passArgs.validateDeclaration(); err != nil {
		panicFn(err)
	}

	if prepare.compoundResult.Kind() != prepare.passArgs.fnCandidate.Type().In(0).Kind() {
		panicFn(errPreparePassArgsAcceptDifferentArgumentType)
	}
}

func (prepare *prepare) checkArgumentsNumber(receivedArgs []reflect.Value) {
	if len(receivedArgs) > prepare.applyFn.fnCandidateValue.Type().NumIn() {
		fnName := runtime.FuncForPC(prepare.applyFn.fnCandidateValue.Pointer()).Name()
		panic(fmt.Sprintf("sequence:%d | fnName:%s | err:%v", prepare.sequence+1, fnName, errPrepareArgsLimitExceeded))
	}
}

func (prepare *prepare) checkArgumentsType(receivedArgs []reflect.Value) {
	for i := 0; i < len(receivedArgs); i++ {
		if receivedArgs[i].Kind() != prepare.applyFn.fnCandidateValue.Type().In(i).Kind() {
			fnName := runtime.FuncForPC(prepare.applyFn.fnCandidateValue.Pointer()).Name()
			panic(fmt.Sprintf("arg-sequence:%d | fnName:%s | err:%v", i+1, fnName, errPrepareArgsDifferentType))
		}
	}
}
