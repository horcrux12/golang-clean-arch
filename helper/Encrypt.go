package helper

import (
	"crypto/md5"
	"crypto/sha256"
	"crypto/sha512"
	"github.com/gobuffalo/packr/v2/file/resolver/encoding/hex"
)

func CheckSumWithMD5(content []byte) (checksum string) {
	hash := md5.New()
	hash.Write(content)
	hashInBytes := hash.Sum(nil)[:16]
	return hex.EncodeToString(hashInBytes)
}

func CheckSumWithSha256(content []byte) string {
	result := sha256.Sum256(content)
	return hex.EncodeToString(result[:])
}

func CheckSumWithSha512(content []byte) string {
	result := sha512.Sum512(content)
	return hex.EncodeToString(result[:])
}

func EncryptPassword(password string, salt string) string {
	return CheckSumWithSha512([]byte(password + salt))
}

func CheckIsPasswordMatch(passwordInput string, passwordDB string, salt string) bool {
	return EncryptPassword(passwordInput, salt) == passwordDB
}
