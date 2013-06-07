package fibonacci

import (
	"github.com/andymoon/go-print"
	"strconv"
)

func Calculate(n int, toConsole bool) (result int) {
	if n < 2 {
		result = 1
	} else {
		result = Calculate(n-2, toConsole) + Calculate(n-1, toConsole)
	}
	if toConsole == true {
		print.Print(strconv.FormatInt(int64(result), 10))
	}
	return result
}
