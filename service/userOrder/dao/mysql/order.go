package mysql

import (
	"fmt"
	"strconv"
	"time"
	"userOrder/model"
	userOrder "userOrder/proto"
)

type UserData struct {
	Id int
}

func InsertOrder(houseId, beginDate, endDate, userName string) (int, error) {
	//获取插入对象
	var order model.OrderHouse
	//给对象赋值
	hid, _ := strconv.Atoi(houseId)
	order.HouseId = uint(hid)

	//把string类型的时间转换为time类型
	bDate, _ := time.Parse("2006-01-02", beginDate)
	order.BeginDate = bDate

	eDate, _ := time.Parse("2006-01-02", endDate)
	order.EndDate = eDate

	//需要userId
	/*var user User
	GlobalDB.Where("name = ?",userName).Find(&user)*/
	//select id form user where name = userName

	var userData UserData
	if err := DB.Raw("select id from user where name = ?", userName).Scan(&userData).Error; err != nil {
		fmt.Println("获取用户数据错误", err)
		return 0, err
	}

	//获取days
	dur := eDate.Sub(bDate)
	order.Days = int(dur.Hours()) / 24
	order.Status = "WAIT_ACCEPT"

	//房屋的单价和总价
	var house model.House
	DB.Where("id = ?", hid).Find(&house).Select("price")
	order.HousePrice = house.Price
	order.Amount = house.Price * order.Days

	order.UserId = uint(userData.Id)
	if err := DB.Create(&order).Error; err != nil {
		fmt.Println("插入订单失败", err)
		return 0, err
	}
	return int(order.ID), nil
}

// GetOrderInfo 获取房东订单如何实现?
func GetOrderInfo(userName, role string) ([]*userOrder.OrdersData, error) {
	//最终需要的数据
	var orderResp []*userOrder.OrdersData
	//获取当前用户的所有订单
	var orders []model.OrderHouse

	var userData UserData
	//用原生查询的时候,查询的字段必须跟数据库中的字段保持一直
	DB.Raw("select id from user where name = ?", userName).Scan(&userData)

	//查询租户的所有的订单
	if role == "custom" {
		if err := DB.Where("user_id = ?", userData.Id).Find(&orders).Error; err != nil {
			fmt.Println("获取当前用户所有订单失败")
			return nil, err
		}
	} else {
		//查询房东的订单  以房东视角来查看订单
		var houses []model.House
		DB.Where("user_id = ?", userData.Id).Find(&houses)

		for _, v := range houses {
			var tempOrders []model.OrderHouse
			DB.Model(&v).Preload("Orders").Find(&tempOrders)
			orders = append(orders, tempOrders...)
		}
	}

	//循环遍历一下orders
	for _, v := range orders {
		var orderTemp userOrder.OrdersData
		orderTemp.OrderId = int32(v.ID)
		orderTemp.EndDate = v.EndDate.Format("2006-01-02")
		orderTemp.StartDate = v.BeginDate.Format("2006-01-02")
		orderTemp.Ctime = v.CreatedAt.Format("2006-01-02")
		orderTemp.Amount = int32(v.Amount)
		orderTemp.Comment = v.Comment
		orderTemp.Days = int32(v.Days)
		orderTemp.Status = v.Status

		//关联house表
		var house model.House
		DB.Model(&v).Preload("houseId").Select("index_image_url", "title")
		orderTemp.ImgUrl = "http://127.0.0.1:8080/" + house.IndexImageUrl
		orderTemp.Title = house.Title

		orderResp = append(orderResp, &orderTemp)
	}
	return orderResp, nil
}

// UpdateStatus 更新订单状态
func UpdateStatus(action, id, reason string) error {
	db := DB.Model(model.OrderHouse{}).Where("id = ?", id)

	if action == "accept" {
		//标示房东同意订单
		return db.Update("status", "WAIT_COMMENT").Error
	} else {
		//表示房东不同意订单  如果拒单把拒绝的原因写到comment中
		return db.Updates(map[string]interface{}{"status": "REJECTED", "comment": reason}).Error
	}
}
