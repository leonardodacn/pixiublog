package utils

// NumEq eq int and interface{}, interface is type of float64
func NumEq(a int, b interface{}) bool {
	if b == nil {
		return false
	}
	c := int(b.(float64))
	if a == c {
		return true
	}
	return false
}
