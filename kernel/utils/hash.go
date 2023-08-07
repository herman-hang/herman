package utils

import (
	"golang.org/x/crypto/bcrypt"
)

// HashEncode 加密密码
// @param string pwd 待加密的明文密码
// @return string error 返回一个哈希加密后的字符串
func HashEncode(pwd string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.DefaultCost)
	if err != nil {
		panic(err.Error())
	}
	return string(hash)
}

// ComparePasswords 验证 hash 密码
// @param string hashedPwd 已加密的hash密码
// @param string sourcePwd 待确认密码
// @return bool 返回判断输入密码的布尔值
func ComparePasswords(hashedPwd string, sourcePwd string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPwd), []byte(sourcePwd))
	if err != nil {
		return false
	}
	return true
}
