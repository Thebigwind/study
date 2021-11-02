package errors

var (
	/*
		// ErrorServerInternal -
		ErrorServerInternal = errors.ServerError
		// ErrorUserInfo -
		ErrorUserId     = errors.NewError(10001, "缺少userId")
		ErrorDeviceId   = errors.NewError(10002, "缺少deviceId")
		ErrorUniqueId   = errors.NewError(10003, "缺少uniqueId")
		ErrorDevicePram = errors.NewError(10004, "设备参数缺失")
		// ErrorLuInfo -
		ErrorPram     = errors.NewError(10007, "请求参数错误")
		ErrorUserInfo = errors.NewError(10008, "获取用户信息失败")
		ErrorOperate  = errors.NewError(10009, "operate参数错误")
		ErrorNotExist = errors.NewError(10010, "数据不存在")
	*/
	// ErrorServerInternal -
	ErrorServerInternal = ServerError
	// ErrorUserInfo -
	ErrorUserId     = NewError(10001, "缺少userId")
	ErrorDeviceId   = NewError(10002, "缺少deviceId")
	ErrorUniqueId   = NewError(10003, "缺少uniqueId")
	ErrorDevicePram = NewError(10004, "设备参数缺失")
	// ErrorLuInfo -
	ErrorPram     = NewError(10007, "请求参数错误")
	ErrorUserInfo = NewError(10008, "获取用户信息失败")
	ErrorOperate  = NewError(10009, "operate参数错误")
	ErrorNotExist = NewError(10010, "数据不存在")
)
