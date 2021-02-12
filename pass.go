package pipe

type PassArgs int

func Pass() PassArgs {
	return PassArgs(0)
}

func isPassArgs(arg interface{}) bool {
	_, ok := arg.(PassArgs)
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
