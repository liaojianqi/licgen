package utils

import (
    "strconv"
    "security"
)
type License struct {
    Client string
    Not_brfore string
    Not_after string
    Max_host string
}
func Hex_byte_to_string(data []byte) string {
    ans := ""
    for _, v := range(data) {
        ss := strconv.FormatInt(int64(v), 16)
        ans += ss
    }
    return ans
}
func Encrypt(plaintext []byte) []byte{
    key := "0\x8e{\x05\x7f\xb7IM\x8f\xc82\xce\xe2j\xf3\xed\x81\xf2U]Q=\xd3\x8f\xcf\xfe\xfe\x86\xd7b\xdb\xac"
	ciphertext := security.Encrypt(security.PKCS7Pad(plaintext), key)
    return ciphertext
}
