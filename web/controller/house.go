package controller

import (
	"context"
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"microHome/web/model"
	houseMicro "microHome/web/proto/house"
	"microHome/web/utils"
	"path"
)

// PostHouses 添加房屋信息
func PostHouses(c *gin.Context) {
	userName := sessions.Default(c).Get("userName")
	var house model.ParamHouseStu
	if err := c.ShouldBind(&house); err != nil {
		ResponseError(c, utils.RecodeParamErr)
		return
	}
	//初始化客户端
	microCilent := houseMicro.NewHouseService("house", utils.GetMicroClient())
	//调用远程函数
	resp, err := microCilent.PubHouse(context.TODO(), &houseMicro.Request{
		Acreage:   house.Acreage,
		Address:   house.Address,
		AreaId:    house.AreaId,
		Beds:      house.Beds,
		Capacity:  house.Capacity,
		Deposit:   house.Deposit,
		Facility:  house.Facility,
		MaxDays:   house.MaxDays,
		MinDays:   house.MinDays,
		Price:     house.Price,
		RoomCount: house.RoomCount,
		Title:     house.Title,
		Unit:      house.Unit,
		UserName:  userName.(string),
	})
	if err != nil {
		zap.L().Error("调用远程函数失败", zap.Any("err", err))
		return
	}
	errCode := utils.MyCode(resp.ErrCode)
	if errCode == utils.RecodeOk {
		ResponseSuccess(c, resp.Data)
	} else {
		ResponseError(c, errCode)
	}

}

// PostHousesImage 添加房屋图片
func PostHousesImage(c *gin.Context) {
	//获取数据
	houseId := c.Param("id")
	fileHeader, err := c.FormFile("house_image")
	//校验数据
	if houseId == "" || err != nil {
		fmt.Println("传入数据不完整", err)
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

	//获取文件字节切片
	file, _ := fileHeader.Open()
	buf := make([]byte, fileHeader.Size)
	file.Read(buf)

	//处理数据  服务中实现
	microClient := houseMicro.NewHouseService("house", utils.GetMicroClient())
	//调用服务
	resp, err := microClient.UploadHouseImg(context.TODO(), &houseMicro.ImgReq{
		HouseId: houseId,
		ImgData: buf,
		FileExt: fileExt,
	})
	if err != nil {
		zap.L().Error("调用远程函数失败", zap.Any("err", err))
		return
	}
	errCode := utils.MyCode(resp.ErrCode)
	if errCode == utils.RecodeOk {
		ResponseSuccess(c, resp)
		return
	} else {
		ResponseError(c, errCode)
		return
	}
}

// GetUserHouses 获取房屋信息
func GetUserHouses(c *gin.Context) {
	userName := sessions.Default(c).Get("userName")
	//初始化客户端
	microClient := houseMicro.NewHouseService("house", utils.GetMicroClient())
	//用客户端调用远程函数
	resp, err := microClient.GetHouseInfo(context.TODO(), &houseMicro.GetReq{UserName: userName.(string)})
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

// GetHouseInfo 获取房屋详情
func GetHouseInfo(c *gin.Context) {
	//获取数据
	houseId := c.Param("id")
	//校验数据
	if houseId == "" {
		ResponseError(c, utils.RecodeParamErr)
		return
	}
	userName := sessions.Default(c).Get("userName")
	//处理数据
	microClient := houseMicro.NewHouseService("house", utils.GetMicroClient())
	//调用远程服务
	resp, err := microClient.GetHouseDetail(context.TODO(), &houseMicro.DetailReq{
		HouseId:  houseId,
		UserName: userName.(string),
	})
	if err != nil {
		zap.L().Error("远程函数 GetHouseDetail 失败:", zap.Error(err))
		return
	}
	errCode := utils.MyCode(resp.ErrCode)
	if errCode == utils.RecodeOk {
		ResponseSuccess(c, resp)
		return
	} else {
		ResponseError(c, errCode)
	}
}

// GetIndex 房屋轮播图
func GetIndex(c *gin.Context) {
	//处理数据
	microClient := houseMicro.NewHouseService("house", utils.GetMicroClient())
	//调用服务
	resp, err := microClient.GetIndexHouse(context.TODO(), &houseMicro.IndexReq{})

	if err != nil {
		zap.L().Error("远程函数 GetHouseDetail 失败:", zap.Error(err))
		return
	}
	errCode := utils.MyCode(resp.ErrCode)
	if errCode == utils.RecodeOk {
		ResponseSuccess(c, resp)
		return
	} else {
		ResponseError(c, errCode)
	}
}

// GetHouses 搜索房屋
func GetHouses(c *gin.Context) {
	//获取数据
	//areaId
	aid := c.Query("aid")
	//start day
	sd := c.Query("sd")
	//end day
	ed := c.Query("ed")
	//排序方式
	sk := c.Query("sk")
	//page  第几页
	//ctx.Query("p")
	//校验数据
	if aid == "" || sd == "" || ed == "" || sk == "" {
		ResponseError(c, utils.RecodeParamErr)
		return
	}

	//处理数据   服务端  把字符串转换为时间格式,使用函数time.Parse()  第一个参数是转换模板,需要转换的二字符串,两者格式一致
	/*sdTime ,_:=time.Parse("2006-01-02 15:04:05",sd+" 00:00:00")
	edTime,_ := time.Parse("2006-01-02",ed)*/

	/*sdTime,_ :=time.Parse("2006-01-02",sd)
	edTime,_ := time.Parse("2006-01-02",ed)
	d := edTime.Sub(sdTime)
	fmt.Println(d.Hours())*/

	microClient := houseMicro.NewHouseService("house", utils.GetMicroClient())
	//调用远程服务
	resp, err := microClient.SearchHouse(context.TODO(), &houseMicro.SearchReq{
		Aid: aid,
		Sd:  sd,
		Ed:  ed,
		Sk:  sk,
	})

	if err != nil {
		zap.L().Error("远程函数 GetHouseDetail 失败:", zap.Error(err))
		return
	}
	errCode := utils.MyCode(resp.ErrCode)
	if errCode == utils.RecodeOk {
		ResponseSuccess(c, resp)
		return
	} else {
		ResponseError(c, errCode)
	}
}
