package pipe

import (
	"errors"
	"reflect"
)

var (
	ErrApplyNotAcceptFn   = errors.New("pipe.Apply(...) not accept function at the first arg")
	ErrApplyNotAcceptArgs = errors.New("pipe.Apply(...) given function could not accept any arg")
	ErrApplyReturnVoid    = errors.New("pipe.Apply(...) given function has void return value")
	ErrApplyMultiReturn   = errors.New("pipe.Apply(...) given function has multiple return values")
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

func (applyFn *applyFn) validateDeclaration() error {
	if applyFn.fnCandidateValue.Kind() != reflect.Func {
		return ErrApplyNotAcceptFn
	}

	if applyFn.fnCandidateValue.Type().NumIn() == 0 {
		return ErrApplyNotAcceptArgs
	}

	if applyFn.fnCandidateValue.Type().NumOut() == 0 {
		return ErrApplyReturnVoid
	}

	if applyFn.fnCandidateValue.Type().NumOut() > 1 {
		return ErrApplyMultiReturn
	}

	return nil
}
