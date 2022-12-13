package encrypt

import "crypto/md5"

func EncryptPassword(password string, salt string) string {
	return EncryptPasswordByMD5(password, salt)
}

func EncryptPasswordByMD5(password string, salt string) string {
	m := md5.New()
	m.Write([]byte(password))
	result := m.Sum([]byte(salt))
	return string(result)
}
