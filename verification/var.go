package verification

import (
	"errors"
	"fmt"
	"strings"
)

// 返回值提前赋值
func Return_err(i int) (err error) {
	if i == 0 {
		err = errors.New("i == 0")
		return
	}

	if i == 1 {
		re, err := get()
		if err != nil {
			fmt.Println(re)
			return err
		}
	}

	return
}

func get() (int, error) {
	return 1, errors.New("i == 1")
}

// strings.Contains
func Contains() {
	f := "Error 1146: Table 'serah.month_shares_202001' doesn't exist"
	s := "error 1146"

	strings.Contains(f, s)
}


