package constants

var MessageFlags = map[int]string{
	Success:   "操作成功",
	Error:     "操作失败",
	MethodBan: "HTTP请求方法被禁止",
}

// GetMessage 根据状态码返回响应信息
// @param int code 状态码
// @return string 返回指定状态码的信息
func GetMessage(code int) string {
	msg, ok := MessageFlags[code]
	if ok {
		return msg
	}

	return MessageFlags[Error]
}
