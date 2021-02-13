package pipe

import (
	"errors"
	"reflect"
)

var (
	errApplyNotAcceptFn   = errors.New("pipe.Apply(...) not accept function at the first arg")
	errApplyNotAcceptArgs = errors.New("pipe.Apply(...) given function could not accept any arg")
	errApplyReturnVoid    = errors.New("pipe.Apply(...) given function has void return value")
	errApplyMultiReturn   = errors.New("pipe.Apply(...) given function has multiple return values")
)

func Apply(fn interface{}, args ...interface{}) *applyFn {
	return &applyFn{
		fnCandidateValue: reflect.ValueOf(fn),
		args:             args,
	}
}

type applyFn struct {
	fnCandidateValue reflect.Value
	args             []interface{}
}

func (applyFn *applyFn) validateDeclaration(applyFnSequence int) error {
	if applyFn.fnCandidateValue.Kind() != reflect.Func {
		return errApplyNotAcceptFn
	}

	if applyFn.fnCandidateValue.Type().NumIn() == 0 && applyFnSequence > 0 {
		return errApplyNotAcceptArgs
	}

	if applyFn.fnCandidateValue.Type().NumOut() == 0 {
		return errApplyReturnVoid
	}

	if applyFn.fnCandidateValue.Type().NumOut() > 1 {
		return errApplyMultiReturn
	}

	return nil
}
