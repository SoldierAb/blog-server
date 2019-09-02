package define


const (
	CODE_SUCC = iota       //成功
	CODE_FAIL              //失败
	CODE_SERVER_ERROR       //服务器错误
	CODE_UPLOAD_ERROR       //上传错误
	CODE_CONTENT_NOT_EXISTED //资源未找到
	CODE_PASS_WRONG           //密码错误
	CODE_USER_NOT_EXISTED     //用户不存在
	CODE_TOKEN_CREATE_ERROR    //TOKEN生成错误
	CODE_NOT_LOGIN            //未登录
	CODE_OVERTIME             //过期
	CODE_SIGN_IN_OTHER_PLACE    //别处登录
)

type BaseRes struct {
	Code int `json:"code"`
	Data interface{} `json:"data"`
	Msg string `json:"msg"`
}


func Msg(code int) string{
	MsgMap:=map[int]string{
		CODE_SUCC:"成功",
		CODE_FAIL:"失败",
		CODE_SERVER_ERROR:"服务器错误",
		CODE_UPLOAD_ERROR:"上传错误",
		CODE_CONTENT_NOT_EXISTED:"资源未找到",
		CODE_PASS_WRONG:"密码错误",
		CODE_USER_NOT_EXISTED:"用户不存在",
		CODE_TOKEN_CREATE_ERROR:"TOKEN生成错误",
		CODE_NOT_LOGIN:"请登录",
		CODE_OVERTIME:"token过期",
		CODE_SIGN_IN_OTHER_PLACE:"用户已在别处登录,请重新登录或者修改密码",
	}

	msg,ok := MsgMap[code]

	if ok {
		return msg
	}else{
		return ""
	}
}


func Res(code int,data ...interface{}) *BaseRes{
	return &BaseRes{
		Code:code,
		Data:data,
		Msg:Msg(code),
	}
}
















