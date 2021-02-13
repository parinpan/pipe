package pipe

import "errors"

var (
	ErrPassArgsOnDoApplyFnFirstSeqNotAllowed = errors.New("pipe.Pass() on pipe.Do(...) first sequence is not allowed")
)

type do struct {
	applyFns []*applyFn
	errors   []error
}

func Do(applyFns ...*applyFn) interface{} {
	d := do{applyFns: applyFns}
	return d.Result()
}

func (do *do) validate() error {
	if len(do.applyFns) == 0 {
		return nil
	}

	if isApplyFnHasPassArgs(do.applyFns[0].args) {
		return ErrPassArgsOnDoApplyFnFirstSeqNotAllowed
	}

	for applyFnSeq, applyFn := range do.applyFns {
		if err := applyFn.validateDeclaration(applyFnSeq); err != nil {
			return err
		}
	}

	return nil
}

func (do *do) Result() interface{} {
	if err := do.validate(); err != nil {
		panic(err)
	}

	if len(do.applyFns) == 0 {
		return nil
	}

	var prepare prepare

	for sequence, applyFn := range do.applyFns {
		prepare.sequence = sequence
		prepare.applyFn = applyFn
		prepare.compoundResult = applyFn.fnCandidateValue.Call(prepare.fnArgs())[0]
	}

	return prepare.compoundResult.Interface()
}
