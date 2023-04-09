package controller

import (
	"context"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"microHome/web/model"
	orderMicro "microHome/web/proto/userOrder"
	"microHome/web/utils"
)

// PostOrders 下订单
func PostOrders(c *gin.Context) {
	//获取数据
	var order model.OrderStu
	err := c.Bind(&order)
	//校验数据
	if err != nil {
		ResponseError(c, utils.RecodeParamErr)
		return
	}
	//获取用户名
	userName := sessions.Default(c).Get("userName")

	//处理数据  服务端处理业务
	microClient := orderMicro.NewUserOrderService("userOrder", utils.GetMicroClient())
	//调用服务
	resp, err := microClient.CreateOrder(context.TODO(), &orderMicro.Request{
		StartDate: order.StartDate,
		EndDate:   order.EndDate,
		HouseId:   order.HouseId,
		UserName:  userName.(string),
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

// GetUserOrder 获取订单信息
func GetUserOrder(c *gin.Context) {
	//获取get请求传参
	role := c.Query("role")
	//校验数据
	if role == "" {
		ResponseError(c, utils.RecodeParamErr)
		return
	}

	//处理数据  服务端
	microClient := orderMicro.NewUserOrderService("userOrder", utils.GetMicroClient())
	//调用远程服务
	resp, err := microClient.GetOrderInfo(context.TODO(), &orderMicro.GetReq{
		Role:     role,
		UserName: sessions.Default(c).Get("userName").(string),
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

// PutOrders 更新订单状态
func PutOrders(c *gin.Context) {
	//获取数据
	id := c.Param("id")
	var statusStu model.StatusStu
	err := c.Bind(&statusStu)

	//校验数据
	if err != nil || id == "" {
		ResponseError(c, utils.RecodeParamErr)
		return
	}

	//处理数据   更新订单状态
	microClient := orderMicro.NewUserOrderService("userOrder", utils.GetMicroClient())
	//调用元和产能服务
	resp, err := microClient.UpdateStatus(context.TODO(), &orderMicro.UpdateReq{
		Action: statusStu.Action,
		Reason: statusStu.Reason,
		Id:     id,
	})

	//返回数据
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
