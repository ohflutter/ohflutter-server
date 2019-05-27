package e

var MsgFlags = map[int]string{
	SUCCESS:               "ok",
	ERROR:                 "fail",
	INVALID_PARAMS:        "请求参数错误",
	ERROR_ADD_USER_FAILED: "新增用户失败",
	ERROR_ADD_ARTICLE_FAILED: "新增文章失败",
}

// GetMsg get error information based on Code
func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}

	return MsgFlags[ERROR]
}
