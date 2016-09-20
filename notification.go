package huaweipush

//
//type Notification struct {
//	tokens      string
//	appMethod   string
//	nspTS       int64
//	nspSVC      string
//	nspFmt      string
//	accessToken string
//	android     string
//}
//
//func NewNotification(token, access_token, android string) *Notification {
//	return &Notification{
//		tokens:      token,
//		nspTS:       time.Now().Second(),
//		nspSVC:      apiMethodPrefix + notificationSendURL,
//		nspFmt:      "JSON",
//		accessToken: access_token,
//		android:     android,
//	}
//}

type SingleNotification struct {
	deviceToken string // 32 字节长度，由系统分配的合法TMID
	message     string // 发送到设备上的消息，最长为4096 字节
	priority    int32  // 0：高优先级; 1：普通优先级, 缺省值为1
	cacheMode   int32  // 消息是否需要缓存 0：不缓存; 1：缓存, 缺省值为0
	msgType     int32  // 标识消息类型（缓存机制），由调用端赋值，取值范围（1~100）。当TMID+msgType的值一样时，仅缓存最新的一条消息
	requestID   string // 如果开发者填写了该字段，则需要保证该字段唯一
	expireTime  string // 消息过期删除时间 unix时间戳,格式：2013-08-29 19:55, 如果不填写，默认超时时间为当前时间后48小时
}

type BatchNotification struct {
	deviceTokenList []string // Device token列表，最多填1000个
	message         string   // 发送到设备上的消息，最长为4096 字节
	cacheMode       int32    // 消息是否需要缓存 0：不缓存; 1：缓存, 缺省值为0
	msgType         int32    // 标识消息类型（缓存机制），由调用端赋值，取值范围（1~100）。当TMID+msgType的值一样时，仅缓存最新的一条消息
	expireTime      string   // 消息过期删除时间 unix时间戳,格式：2013-08-29 19:55, 如果不填写，默认超时时间为当前时间后48小时
}

func NewSingleNotification(deviceToken, message, requestID string) *SingleNotification {
	return &SingleNotification{
		deviceToken: deviceToken,
		message:     message,
		priority:    1,
		cacheMode:   0,
		msgType:     0,
		requestID:   requestID,
		expireTime:  "",
	}
}

func (s *SingleNotification) SetHighPriority() *SingleNotification {
	s.priority = 0
	return s
}

func (s *SingleNotification) SetExpireTime(expireTime string) *SingleNotification {
	s.expireTime = expireTime
	return s
}

func (s *SingleNotification) SetCacheMode(cacheMode int32) *SingleNotification {
	s.cacheMode = cacheMode
	return s
}

func (s *SingleNotification) Map() map[string]interface{} {
	m := make(map[string]interface{})
	m["deviceToken"] = s.deviceToken
	m["message"] = s.message
	m["priority"] = s.priority
	m["cacheMode"] = s.cacheMode
	m["msgType"] = s.msgType
	m["requestID"] = s.requestID
	m["expireTime"] = s.expireTime
	return m
}

func NewBatchNotification(deviceTokenList []string, message string) *BatchNotification {
	return &BatchNotification{
		deviceTokenList: deviceTokenList,
		message:         message,
		cacheMode:       0,
		msgType:         0,
		expireTime:      "",
	}
}
