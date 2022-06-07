package operator

func Math(expr string) (interface{}, error) {
	return Evaluate(expr, nil)
}
