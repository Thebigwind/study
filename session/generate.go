package main

import (
	"context"
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"
)

func (d *DeviceLoginSession) GenerateUserSession(ctx context.Context, req *deviceloginsessiontype.GenerateUserSessionRequest, resp *deviceloginsessiontype.GenerateUserSessionResponse, userInfo model.User) error {

	//sessionKey
	sessionKey := ""
	if req.Client != "" && tools.ContainString(variable.GetNotXZAppClient(), req.Client) {
		sessionKey = tools.GetSessId4Html5(strconv.FormatInt(req.UserID, 10), userInfo.Business)
		resp.SessId = sessionKey
	} else if req.SessID != "" && d.IsValidSessId(ctx, req.SessID, userInfo, req.Client, false) { //是否是有效的sessId
		sessionKey = req.SessID
		resp.SessId = sessionKey
	} else {
		if req.UniqueID == "" {
			return nil
		}
		if req.Client == "" {
			return nil
		}
		//
		deviceUniqueId := ""
		match, _ := regexp.MatchString("'/^0[0\\-]+0$/", req.UniqueID)

		if req.UniqueID == "" || match {
			deviceUniqueId = "unknown-device-" + strconv.FormatInt(req.UserID, 10)
		} else {
			deviceUniqueId = req.UniqueID
		}
		//goto 跳转
	REQUERY:
		//1.查询 device信息
		deviceSvc := &Device{}
		var deviceResp devicetype.GetDevice
		if err := deviceSvc.GetDeviceByUniqueId(ctx, deviceUniqueId, &deviceResp); err != nil {
			log.GetLogger(ctx).Infof("GetDeviceByUniqueId err:%s", err.Error())
		}

		//
		var deviceAddReq devicetype.AddDeviceRequest
		var deviceAddResp devicetype.AddDeviceResponse
		//device设备信息不存在则插入device信息
		if deviceResp.DeviceUniqueId == "" {
			//insert
			deviceAddReq.DeviceUniqueId = req.UniqueID
			deviceAddReq.UserId = req.UserID
			deviceAddReq.Client = req.Client
			deviceAddReq.DeviceBrand = req.Brand
			deviceAddReq.Idfa = ""

			if err := deviceSvc.Add(ctx, &deviceAddReq, &deviceAddResp); err != nil {
				if !strings.Contains(err.Error(), "Duplicate") { //如果错误类型不是 Duplicate key,返回错误
					return err
				}
				//
				goto REQUERY
			}
			deviceResp.ID = deviceAddResp.ID
		}

		deviceResp.DeviceBrand = req.Brand
		deviceResp.DeviceUniqueId = req.UniqueID
		deviceResp.Client = req.Client
		deviceResp.UserId = req.UserID

		//usersession
		//deviceloginsessionSvc := &service.DeviceLoginSession{}
		var addReq deviceloginsessiontype.AddDeviceLoginSessionRequest
		var addResq deviceloginsessiontype.AddDeviceLoginSessionResponse

		addReq.SessionType = req.SessionType
		addReq.EnvironmentInfo = req.Env
		addReq.DeviceType = req.Client

		//fmt.Printf("\n req.SessionType:%s\n", req.SessionType)
		//sessionKey
		err, sessKey := d.CreateSession(ctx, &addReq, &addResq, deviceResp, userInfo)
		if err != nil {
			return err
		}
		resp.SessId = sessKey
	}
	return nil
}


func (d *DeviceLoginSession) CreateSession(
	ctx context.Context,
	req *deviceloginsessiontype.AddDeviceLoginSessionRequest,
	resp *deviceloginsessiontype.AddDeviceLoginSessionResponse,
	device devicetype.GetDevice,
	userInfo model.User) (error, string) {
	var session model.DeviceLoginSession

	//userId, _ := strconv.ParseInt(userInfo.UserID, 10, 64)
	userId := userInfo.ID
	refreshTimeStamp := time.Now().Unix() + int64(variable.GetExpireTime(req.SessionType))

	//获取唯一id
	uniqueID := idmaker.GetID()
	if uniqueID == 0 {
		return nil, ""
	}

	dataStr := strconv.FormatInt(userInfo.ID, 10) + "|" + userInfo.Mobile + "|" + userInfo.Password + "|" + VerifyParams(req.SessionType, userInfo)

	var dd = tools.SessionData{
		SsId:   uniqueID,        //dataMap["ssId"]
		SsType: req.SessionType, //dataMap["ssType"]
		Data:   tools.Str2Md5(dataStr),
		Expire: refreshTimeStamp,
	}

	//生成 $sessionKey
	sessionKey := tools.StateGenerate(ctx, dd, variable.USER_SESSION_SALT)

	//异步删除旧session  为什么不直接设置状态？？
	// (1)查询
	deviceSession, err := session.GetLoginSessionByDeviceId(ctx, device.ID)
	if err != nil {
		fmt.Printf("GetLoginSessionByDeviceId error:%+v", err)
	}
	//（2）设置deleted状态
	session.EnvironmentInfo = deviceSession.EnvironmentInfo
	session.DeviceId = device.ID
	reason := "deviceId:" + strconv.FormatInt(device.ID, 10)
	if err := session.DestroyDeviceLoginSession(ctx, reason); err != nil {
		return err, ""
	}

	//存储新的session
	session.UserId = userId
	session.ExpireTime = orm.LocalTime{Time: time.Unix(refreshTimeStamp, 0)}
	session.RefreshTime = orm.LocalTime{Time: time.Unix(refreshTimeStamp, 0)}
	session.Ver = 1 //需要处理
	session.Estate = variable.STATE_VALID
	session.DeviceId = device.ID
	session.DeviceType = req.DeviceType
	session.EnvironmentInfo = req.EnvironmentInfo
	session.ID = uniqueID
	session.SessionKey = sessionKey
	session.SessionType = req.SessionType
	session.CreateTime = orm.LocalTime{Time: time.Now()}

	if err = session.AddDeviceLoginSession(ctx); err != nil {
		return err, ""
	}
	resp.ID = session.ID

	//同步redis
	//sessionCache := cache.NewDeviceLoginSessionCache()
	//sessionCache.Key = strconv.FormatInt(userInfo.ID, 10) + sessionCache.Key //和 php 保持一致
	//if err := sessionCache.SetCache(ctx, session); err != nil {
	//	log.GetLogger(ctx).Errorf("deviceloginsession update to db success, but sync redis error:%s", err.Error())
	//}
	go func() {
		key := strconv.FormatInt(userInfo.ID, 10) + cache.DeviceLoginSessionKeySuffix
		if err := cache.SetDeviceLoginSessionCache(ctx, key, session); err != nil {
			log.GetLogger(ctx).Infof("deviceloginsession update to db success, but sync redis error:%s", err.Error())
		}
	}()

	return nil, sessionKey
}