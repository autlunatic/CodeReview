package businessLogicStuff

import "strconv"

func DoStuff(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		return -1
	}

	return i

}
