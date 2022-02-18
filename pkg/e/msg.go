package e

var MsgFlags = map[int]string{
	SUCCESS:        "成功",
	ERROR:          "错误",
	INVALID_PARAMS: "请求参数错误",

	ERROR_EXIST_TAG:         "已存在标签名",
	ERROR_NOT_EXIST_TAG:     "该标签不存在",
	ERROR_NOT_EXIST_ARTICLE: "该文章不存在",

	ERROR_AUTH_CHECK_TOKEN_FAIL:     "Token鉴权失败",
	ERROR_AUTH_CHECK_TOKEN_TIMEOUT:  "Token已超时",
	ERROR_AUTH_TOKEN:                "Token生成错误",
	ERROR_AUTH:                      "Token错误",
	ERROR_UPLOAD_SAVE_IMAGE_FAIL:    "保存图片失败",
	ERROR_UPLOAD_CHECK_IMAGE_FAIL:   "检查图片失败",
	ERROR_UPLOAD_CHECK_IMAGE_FORMAT: "校验文件错误,图片格式或大小有问题",
}

func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}
	return MsgFlags[ERROR]
}
