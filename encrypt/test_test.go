package encrypt

import (
	"fmt"
	"testing"
)

func TestEncryptPassword(t *testing.T) {
	password := EncryptPassword("hhr12138", "qwer")
	fmt.Println(password)
}
