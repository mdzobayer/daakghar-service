package token

import (
	"fmt"
	"testing"
	"time"
)

var testset = []struct {
	plaintext  string
	ciphertext string
	key        string
}{
	{"", "", ""},
	{"", "", ""},
	{"admin@Mon, 20 Feb 2016 15:04:05 -0700", "", "Hello World&*^^&KJNHJ&y7jnjbb((("},
}

func TestEncrypt(t *testing.T) {
	for _, val := range testset {
		user := User{}
		user.fromString(val.plaintext)
		fmt.Println(user.UserName)
		fmt.Println(user.Expire.Format(time.RFC1123Z))
		fmt.Println("token string: ", user.string())
		curToken, _ := Encrypt(user, val.key)
		if curToken != val.ciphertext {
			fmt.Println("error: ", curToken, " - ", val.ciphertext)
			fmt.Println("decryption: ", Decrypt(curToken, val.key))
		}
	}
}
