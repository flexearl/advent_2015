package Day_4

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
)

func Part_1(data string) {
	data = "ckczppom"
	validNumb := 0
	for i := 1; true; i++ {
		hash := getMD5Hash(data + fmt.Sprint(i))
		if hasLeadingZeros(hash) {
			validNumb = i
			break
		}
	}
	fmt.Println(validNumb)
}

func getMD5Hash(text string) string {
	hash := md5.Sum([]byte(text))
	return hex.EncodeToString(hash[:])
}

func hasLeadingZeros(text string) bool {
	hasLeadingZeros := true
	for i := 0; i < 6 && i < len(text); i++ {
		if text[i] != '0' {
			hasLeadingZeros = false
			break
		}
	}
	return hasLeadingZeros
}
