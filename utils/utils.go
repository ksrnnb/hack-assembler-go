package utils

type StringSlice []string

func (ss StringSlice) Contains(str string) bool {
	for _, s := range ss {
		if s == str {
			return true
		}
	}

	return false
}
