package main

type stringSlice []string

func (ss stringSlice) contains(str string) bool {
	for _, s := range ss {
		if s == str {
			return true
		}
	}

	return false
}
