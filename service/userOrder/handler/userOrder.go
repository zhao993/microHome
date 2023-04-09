package handler

import (
	"context"
	"fmt"
	"strconv"
	"userOrder/dao/mysql"
	pb "userOrder/proto"
	"userOrder/utils"
)

type UserOrder struct{}

func (e *UserOrder) CreateOrder(ctx context.Context, req *pb.Request, rsp *pb.Response) error {
	//获取到相关数据,插入到数据库
	orderId, err := mysql.InsertOrder(req.HouseId, req.StartDate, req.EndDate, req.UserName)
	if err != nil {
		rsp.ErrCode = string(utils.RecodeDbErr)
		return nil
	}
	rsp.ErrCode = string(utils.RecodeOk)
	var orderData pb.OrderData
	orderData.OrderId = strconv.Itoa(orderId)
	rsp.Data = &orderData
	return nil

}

func (e *UserOrder) GetOrderInfo(ctx context.Context, req *pb.GetReq, rsp *pb.GetResp) error {
	//要根据传入数据获取订单信息   mysql
	respData, err := mysql.GetOrderInfo(req.UserName, req.Role)
	if err != nil {
		rsp.ErrCode = string(utils.RecodeDbErr)
		return nil
	}

	rsp.ErrCode = string(utils.RecodeOk)
	var getData pb.GetData
	getData.Orders = respData
	rsp.Data = &getData
	return nil
}
func (e *UserOrder) UpdateStatus(ctx context.Context, req *pb.UpdateReq, rsp *pb.UpdateResp) error {
	//根据传入数据,更新订单状态
	err := mysql.UpdateStatus(req.Action, req.Id, req.Reason)
	if err != nil {
		fmt.Println("更新订单装填错误", err)
		rsp.ErrCode = string(utils.RecodeDbErr)
		return nil
	}
	rsp.ErrCode = string(utils.RecodeOk)
	return nil
}
