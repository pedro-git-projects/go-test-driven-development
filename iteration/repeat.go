package iteration

// Repeat takes a string s and an int n and repeats the string n times
func Repeat(toRepeat string, nTimes int) string {
	var repeated string
	for i := 0; i < nTimes; i++ {
		repeated += toRepeat
	}
	return repeated
}
