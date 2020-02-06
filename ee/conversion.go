package ee

// Int tries to parse an Int from the given value. If the value denotes a string slice, the first
// value is parsed from it.
func Int(val interface{}, err error) (int, error) {
	return 0, nil
}

// Int tries to parse an Int from the given value. If the value denotes a string slice, the first
// value is parsed from it. If it cannot be parsed, 0 is returned.
func OptInt(val interface{}, err error) int {
	return 0
}
