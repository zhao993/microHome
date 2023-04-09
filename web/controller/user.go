package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/afocus/captcha"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"image/png"
	"microHome/web/model"
	getAreaMicro "microHome/web/proto/getArea"
	getCaptcha "microHome/web/proto/getCaptcha"
	registerMicro "microHome/web/proto/register"
	userMicro "microHome/web/proto/user"
	"microHome/web/utils"
	"path"
)

// GetSession 获取session信息
func GetSession(c *gin.Context) {
	session := sessions.Default(c)
	userName := session.Get("userName")
	if userName == nil {
		ResponseError(c, utils.RecodeSessionErr)
		return
	}
	data := make(map[string]interface{}, 1)
	data["name"] = userName
	ResponseSuccess(c, data)

}

// GetImageCd 微服务获取验证码图片信息
func GetImageCd(c *gin.Context) {
	uuid := c.Param("uuid")
	//初始化客户端
	microClient := getCaptcha.NewGetCaptchaService("getCaptcha", utils.GetMicroClient())
	//用客户端调用远程函数
	resp, err := microClient.Call(context.TODO(), &getCaptcha.Request{Uuid: uuid})
	if err != nil {
		zap.L().Error("call error:", zap.Error(err))
	}
	//将得到的数据反序列化
	var img captcha.Image
	json.Unmarshal(resp.Img, &img)
	//图片写出去
	png.Encode(c.Writer, img)
}

// GetSmsCd 校验图片验证码，获取短信验证码
func GetSmsCd(c *gin.Context) {
	phone := c.Param("phone")
	//拆分Get请求中的URL==格式：   ?k=value&k=value
	imgCode := c.Query("text")
	uuid := c.Query("id")
	// 初始化客户端
	microClient := registerMicro.NewRegisterService("register", utils.GetMicroClient())
	// 调用远程函数:
	resp, err := microClient.SendSms(context.TODO(), &registerMicro.CallRequest{Phone: phone, ImgCode: imgCode, Uuid: uuid})

	if err != nil {
		fmt.Println("调用远程函数 SendSms 失败:", err)
		zap.L().Error("err", zap.Error(err))
		return
	}
	errCode := utils.MyCode(resp.ErrCode)
	if errCode == utils.RecodeOk {
		ResponseSuccess(c, "success")
	} else {
		ResponseError(c, errCode)
	}
}

// PostRet 注册用户，存入数据库
func PostRet(c *gin.Context) {
	/*mobile := c.PostForm("mobile")
	pwd := c.PostForm("password")
	smsCode := c.PostForm("sms_code")*/
	//获取数据
	var ParamUser model.ParamUser
	if err := c.ShouldBindJSON(&ParamUser); err != nil {
		zap.L().Error("获取参数错误", zap.Error(err))
		return
	}
	//初始化客户端
	microClient := registerMicro.NewRegisterService("register", utils.GetMicroClient())
	//用客户端调用远程函数
	resp, err := microClient.Register(context.TODO(), &registerMicro.RegReq{Mobile: ParamUser.Mobile, Password: ParamUser.Password, SmsCode: ParamUser.SmsCode})
	if err != nil {
		fmt.Println("调用远程函数 SendSms 失败:", err)
		zap.L().Error("err", zap.Error(err))
		return
	}
	errCode := utils.MyCode(resp.ErrCode)
	if errCode == utils.RecodeOk {
		ResponseSuccess(c, "success")
	} else {
		ResponseError(c, errCode)
	}
}

// GetArea 获得地区信息
func GetArea(c *gin.Context) {
	//初始化客户端
	microClient := getAreaMicro.NewGetAreaService("getarea", utils.GetMicroClient())
	//用客户端调用远程函数
	resp, err := microClient.MicroGetArea(context.TODO(), &getAreaMicro.Request{})
	if err != nil {
		fmt.Println("调用远程函数 SendSms 失败:", err)
		zap.L().Error("err", zap.Error(err))
		return
	}
	errCode := utils.MyCode(resp.ErrCode)
	if errCode == utils.RecodeOk {
		ResponseSuccess(c, resp.Data)
	} else {
		ResponseError(c, errCode)
	}
}

// PostLogin 用户登录设置Session
func PostLogin(c *gin.Context) {
	var paramLogin model.ParamLogin
	if err := c.ShouldBindJSON(&paramLogin); err != nil {
		ResponseError(c, utils.RecodeParamErr)
		return
	}
	//初始化客户端
	microClient := registerMicro.NewRegisterService("register", utils.GetMicroClient())
	//用客户端调用远程函数
	resp, err := microClient.Login(context.TODO(), &registerMicro.LoginReq{Mobile: paramLogin.Mobile, Password: paramLogin.Password})
	if err != nil {
		fmt.Println("调用远程函数 Login 失败:", err)
		zap.L().Error("err", zap.Error(err))
		return
	}
	errCode := utils.MyCode(resp.ErrCode)
	if errCode == utils.RecodeOk {
		session := sessions.Default(c)
		//将登录状态写到session
		//设置session
		session.Set("userName", resp.Name) //将用户名存储到session中
		session.Save()
		ResponseSuccess(c, "success")
	} else {
		ResponseError(c, errCode)
	}
}

// DeleteSession 用户退出登录
func DeleteSession(c *gin.Context) {
	session := sessions.Default(c)
	session.Delete("userName")
	//必须保存
	err := session.Save()
	if err != nil {
		ResponseError(c, utils.RecodeServerErr)
	}
	ResponseSuccess(c, "success")
}

// GetUserInfo 获得用户信息
func GetUserInfo(c *gin.Context) {
	username := sessions.Default(c).Get("userName")
	var user model.ParamUsers
	//初始化客户端
	microClient := userMicro.NewUserService("user", utils.GetMicroClient())
	//用客户端调用远程函数
	resp, err := microClient.UserInfo(context.TODO(), &userMicro.NameData{Name: username.(string)})
	if err != nil {
		fmt.Println("调用远程函数UserInfo失败:", err)
		zap.L().Error("err", zap.Error(err))
		return
	}
	errCode := utils.MyCode(resp.ErrCode)
	if errCode == utils.RecodeOk {
		user.Name = resp.Data.Name
		user.Mobile = resp.Data.Mobile
		user.RealName = resp.Data.RealName
		user.ID = resp.Data.UserId
		user.IDCard = resp.Data.IdCard
		user.AvatarUrl = "http://192.168.17.129:8888/" + resp.Data.AvatarUrl
		ResponseSuccess(c, user)
	} else {
		ResponseError(c, errCode)
	}
}

// PutUserInfo 用户修改信息
func PutUserInfo(c *gin.Context) {
	//获取当前用户名
	session := sessions.Default(c)
	username := session.Get("userName")
	//获取用户名
	var nameData struct {
		Name string `json:"name"`
	}
	err := c.ShouldBindJSON(&nameData)
	if err != nil {
		ResponseError(c, utils.RecodeParamErr)
		return
	}
	//初始化客户端
	microClient := userMicro.NewUserService("user", utils.GetMicroClient())
	//用客户端调用远程函数
	resp, err := microClient.UpdateUserName(context.TODO(), &userMicro.UpdateReq{OldName: username.(string), NewName: nameData.Name})
	if err != nil {
		fmt.Println("调用远程函数 SendSms 失败:", err)
		zap.L().Error("err", zap.Error(err))
		return
	}
	errCode := utils.MyCode(resp.ErrCode)
	//更新用户名
	if errCode == utils.RecodeOk {
		session.Set("userName", resp.Data.Name)
		session.Save()
		data := make(map[string]string, 1)
		data["name"] = resp.Data.Name
		ResponseSuccess(c, data)
	} else {
		ResponseError(c, errCode)
	}
}

// PostAvatar 用户上传头像
func PostAvatar(c *gin.Context) {
	//获取数据  获取图片  文件流  文件头  err
	fileHeader, err := c.FormFile("avatar")

	//检验数据
	if err != nil {
		fmt.Println("文件上传失败")
		return
	}

	//三种校验 大小,类型,防止重名  fastdfs
	if fileHeader.Size > 50000000 {
		fmt.Println("文件过大,请重新选择")
		return
	}

	fileExt := path.Ext(fileHeader.Filename)
	if fileExt != ".png" && fileExt != ".jpg" {
		fmt.Println("文件类型错误,请重新选择")
		return
	}
	//只读的文件指针
	file, _ := fileHeader.Open()
	buf := make([]byte, fileHeader.Size)
	file.Read(buf)

	/*
		fdfsClient,_ := fdfs_client.NewFdfsClient("/etc/fdfs/client.conf")
		//fdfsClient.UploadByFilename()
		fdfsResp,_ := fdfsClient.UploadByBuffer(buf,fileExt[1:])
		fmt.Println("上传文件到fastdfs的组名为",fdfsResp.GroupName," 凭证为",fdfsResp.RemoteFileId)*/

	//获取用户名
	session := sessions.Default(c)
	userName := session.Get("userName")

	//处理数据
	//初始化客户端
	microClient := userMicro.NewUserService("user", utils.GetMicroClient())
	//调用远程函数
	resp, err := microClient.UploadAvatar(context.TODO(), &userMicro.UploadReq{
		UserName: userName.(string),
		Avatar:   buf,
		FileExt:  fileExt,
	})
	if err != nil {
		fmt.Println("调用远程函数 SendSms 失败:", err)
		zap.L().Error("err", zap.Error(err))
		return
	}
	errCode := utils.MyCode(resp.ErrCode)
	if errCode == utils.RecodeOk {
		ResponseSuccess(c, resp)
	} else {
		ResponseError(c, errCode)
	}
}

// PostAuth 实名认证
func PostAuth(c *gin.Context) {
	var auth model.ParamAuth
	err := c.ShouldBindJSON(&auth)
	if err != nil {
		ResponseError(c, utils.RecodeParamErr)
		return
	}
	username := sessions.Default(c).Get("userName")
	//初始化客户端
	microClient := userMicro.NewUserService("user", utils.GetMicroClient())
	//用客户端调用远程函数
	resp, err := microClient.AuthUpdate(context.TODO(), &userMicro.AuthReq{Name: username.(string), RealName: auth.RealName, IdCard: auth.IDCard})
	if err != nil {
		fmt.Println("调用远程函数 SendSms 失败:", err)
		zap.L().Error("err", zap.Error(err))
		return
	}
	errCode := utils.MyCode(resp.ErrCode)
	if errCode == utils.RecodeOk {
		ResponseSuccess(c, "success")
	} else {
		ResponseError(c, errCode)
	}
}
