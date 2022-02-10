package repositories

import (
	"crypto/md5"
	"fmt"
	"io"
)

func EncryptPassword(password string) string {
	tPass := md5.New()
	io.WriteString(tPass, password)
	ePassword := fmt.Sprintf("%x", tPass.Sum(nil))

	return ePassword
}
