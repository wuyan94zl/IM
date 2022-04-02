package utils

import (
	"crypto/md5"
	"fmt"
	"io"
)

/** 加密方式 **/

func Md5ByString(str string) string {
	m := md5.New()
	io.WriteString(m, str)
	arr := m.Sum(nil)
	return fmt.Sprintf("%x", arr)
}

func Md5ByBytes(b []byte) string {
	return fmt.Sprintf("%x", md5.Sum(b))
}
