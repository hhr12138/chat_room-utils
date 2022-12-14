package encrypt

import (
	"crypto/md5"
	"fmt"
)

func EncryptPassword(password string, salt string) string {
	return EncryptPasswordByMD5(password, salt)
}

func EncryptPasswordByMD5(password string, salt string) string {
	m := md5.New()
	m.Write([]byte(password + salt))
	result := m.Sum(nil)
	return fmt.Sprintf("%x", result)
}
