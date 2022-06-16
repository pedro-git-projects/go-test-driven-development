package iteration

const nTimes = 5

func Repeat(toRepeat string) string {
	var repeated string
	for i := 0; i < nTimes; i++ {
		repeated += toRepeat
	}
	return repeated
}
