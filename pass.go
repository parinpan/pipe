package pipe

import (
	"errors"
	"reflect"
)

var (
	errPassArgsFnCandidateIllegalType = errors.New("pipe.Pass(...) only accepts function in the argument")
	errPassArgsFnCandidateNoArgument  = errors.New("pipe.Pass(...) accepted function should accept at least one argument")
	errPassArgsFnCandidateVoidReturn  = errors.New("pipe.Pass(...) accepted function should not return void")
)

type emptyFnCandidate struct {
}

type passArgs struct {
	fnCandidate reflect.Value
}

func Pass(fn ...interface{}) *passArgs {
	if len(fn) > 0 {
		return &passArgs{fnCandidate: reflect.ValueOf(fn[0])}
	}

	return &passArgs{fnCandidate: reflect.ValueOf(emptyFnCandidate{})}
}

func (pa *passArgs) validateDeclaration() error {
	if pa.fnCandidate.Kind() != reflect.Func {
		return errPassArgsFnCandidateIllegalType
	}

	if pa.fnCandidate.Type().NumIn() == 0 {
		return errPassArgsFnCandidateNoArgument
	}

	if pa.fnCandidate.Type().NumOut() == 0 {
		return errPassArgsFnCandidateVoidReturn
	}

	return nil
}

func isPassArgs(arg interface{}) bool {
	_, ok := arg.(passArgs)
	return ok
}

func isApplyFnHasPassArgs(args []interface{}) bool {
	for _, arg := range args {
		if isPassArgs(arg) {
			return true
		}
	}

	return false
}
