package constants

var MessageFlags = map[int]string{
	SUCCESS: "ok",
	ERROR:   "fail",
}

// GetMessage 根据状态码返回响应信息
func GetMessage(code int) string {
	msg, ok := MessageFlags[code]
	if ok {
		return msg
	}

	return MessageFlags[ERROR]
}
