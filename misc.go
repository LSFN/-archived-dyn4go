package dyn4go

func ReverseSlice(s []interface{}) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func SliceContains(container []interface{}, containee interface{}) bool {
	for _, v := range container {
		if v == containee {
			return true
		}
	}
	return false
}
