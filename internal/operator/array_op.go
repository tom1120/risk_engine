package operator

// jundge val in arr
func InArray(arr []interface{}, val interface{}) bool {
	if len(arr) == 0 {
		return false
	}
	for _, v := range arr {
		if ok, err := Compare("EQ", v, val); err == nil && ok {
			return true
		}
	}
	return false
}

// jundge array A in Array B
func AInB(a []interface{}, b []interface{}) bool {
	if len(b) == 0 {
		return false
	}
	if len(a) == 0 {
		return true
	}
	if len(a) > len(b) {
		return false
	}
	tmp := make(map[interface{}]struct{}, len(b))
	for _, v := range b {
		tmp[v] = struct{}{}
	}
	for _, v := range a {
		if _, ok := tmp[v]; !ok {
			return false
		}
	}
	return true
}
