package encrypt

import (
	"crypto/md5"
	"strconv"
)

func Token(username int64, passwd string) string {
	hash := md5.New()
	hash.Write([]byte(strconv.FormatInt(username, 16) + passwd))
	sum := hash.Sum(nil)
	return string(sum)
}
