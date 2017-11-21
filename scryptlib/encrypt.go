package scryptlib

import (
	"fmt"
	"io"
)

func Encrypt(password []byte, in io.Reader, out io.Writer) error {
	i, err := cpuperf()
	fmt.Println(i, err)
	pickparams(1024*1024*1024, 2.0)
	return nil
}
