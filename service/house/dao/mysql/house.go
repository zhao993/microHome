package mysql

import (
	"fmt"
	"house/model"
	house "house/proto"
	"strconv"
	"time"
)

// AddHouse 发布房屋
func AddHouse(request *house.Request) (int, error) {
	var houseInfo model.House
	//给house赋值
	houseInfo.Address = request.Address

	//根据userName获取userId
	var user model.User
	if err := DB.Where("name = ?", request.UserName).Find(&user).Error; err != nil {
		fmt.Println("查询当前用户失败", err)
		return 0, err
	}

	//sql中一对多插入,只是给外键赋值
	houseInfo.UserId = user.ID
	houseInfo.Title = request.Title
	//类型转换
	price, _ := strconv.Atoi(request.Price)
	roomCount, _ := strconv.Atoi(request.RoomCount)
	houseInfo.Price = price
	houseInfo.RoomCount = roomCount
	houseInfo.Unit = request.Unit
	houseInfo.Capacity, _ = strconv.Atoi(request.Capacity)
	houseInfo.Beds = request.Beds
	houseInfo.Deposit, _ = strconv.Atoi(request.Deposit)
	houseInfo.MinDays, _ = strconv.Atoi(request.MinDays)
	houseInfo.MaxDays, _ = strconv.Atoi(request.MaxDays)
	houseInfo.Acreage, _ = strconv.Atoi(request.MaxDays)
	//一对多插入
	areaId, _ := strconv.Atoi(request.AreaId)
	houseInfo.AreaId = uint(areaId)

	//request.Facility    所有的家具  房屋
	for _, v := range request.Facility {
		id, _ := strconv.Atoi(v)
		var fac model.Facility
		if err := DB.Where("id = ?", id).First(&fac).Error; err != nil {
			fmt.Println("家具id错误", err)
			return 0, err
		}
		//查询到了数据
		houseInfo.Facilities = append(houseInfo.Facilities, &fac)
	}

	if err := DB.Create(&houseInfo).Error; err != nil {
		fmt.Println("插入房屋信息失败", err)
		return 0, err
	}
	return int(houseInfo.ID), nil
}

// SaveHouseImg 把图片的凭证存储到数据中   更新   主图,次图  第一张图片是主图,剩下的图片是副图
func SaveHouseImg(houseId, imgPath string) error {
	/*return GlobalDB.Model(new(House)).Where("id = ?",houseId).
	Update("index_image_url",imgPath).Error*/
	//如何判断上传的图是当前房屋的第一张图片
	var houseInfo model.House
	if err := DB.Where("id = ?", houseId).Find(&houseInfo).Error; err != nil {
		fmt.Println("查询不到房屋信息", err)
		return err
	}

	if houseInfo.IndexImageUrl == "" {
		//说明没有上传过图片  现在上传的图片是主图
		return DB.Model(model.House{}).Where("id = ?", houseId).
			Update("index_image_url", imgPath).Error
	}

	//上传的幅图
	var houseImg model.HouseImage
	houseImg.Url = imgPath
	hId, _ := strconv.Atoi(houseId)
	houseImg.HouseId = uint(hId)
	return DB.Create(&houseImg).Error
}

// GetUserHouse 获取当前用户发布房源
func GetUserHouse(userName string) ([]*house.Houses, error) {
	var houseInfos []*house.Houses

	//有用户名
	var user model.User
	if err := DB.Where("name = ?", userName).Find(&user).Error; err != nil {
		fmt.Println("获取当前用户信息错误", err)
		return nil, err
	}

	//房源信息   一对多查询
	var houses []model.House
	err := DB.Model(&user).Preload("Houses").Find(&houses).Error
	if err != nil {
		fmt.Println("查找信息错误", err)
		return nil, err
	}

	for _, v := range houses {
		var houseInfo house.Houses
		houseInfo.Title = v.Title
		houseInfo.Address = v.Address
		houseInfo.Ctime = v.CreatedAt.Format("2006-01-02 15:04:05")
		houseInfo.HouseId = int32(v.ID)
		houseInfo.ImgUrl = "http://127.0.0.1:8080/" + v.IndexImageUrl
		houseInfo.OrderCount = int32(v.OrderCount)
		houseInfo.Price = int32(v.Price)
		houseInfo.RoomCount = int32(v.RoomCount)
		houseInfo.UserAvatar = "http://127.0.0.1:8080/" + user.AvatarUrl

		//获取地域信息
		var area model.Area
		//related函数可以是以主表关联从表,也可以是以从表关联主表
		DB.Where("id = ?", v.AreaId).Find(&area)
		houseInfo.AreaName = area.Name

		houseInfos = append(houseInfos, &houseInfo)
	}
	return houseInfos, nil
}

// GetHouseDetail 获取房屋详情
func GetHouseDetail(houseId, userName string) (house.DetailData, error) {
	var respData house.DetailData
	//给houseDetail赋值
	var houseDetail house.HouseDetail

	var houseInfo model.House
	if err := DB.Where("id = ?", houseId).Find(&houseInfo).Error; err != nil {
		fmt.Println("查询房屋信息错误", err)
		return respData, err
	}
	{
		houseDetail.Acreage = int32(houseInfo.Acreage)
		houseDetail.Address = houseInfo.Address
		houseDetail.Beds = houseInfo.Beds
		houseDetail.Capacity = int32(houseInfo.Capacity)
		houseDetail.Deposit = int32(houseInfo.Deposit)
		houseDetail.Hid = int32(houseInfo.ID)
		houseDetail.MaxDays = int32(houseInfo.MaxDays)
		houseDetail.MinDays = int32(houseInfo.MinDays)
		houseDetail.Price = int32(houseInfo.Price)
		houseDetail.RoomCount = int32(houseInfo.RoomCount)
		houseDetail.Title = houseInfo.Title
		houseDetail.Unit = houseInfo.Unit
		if houseInfo.IndexImageUrl != "" {
			houseDetail.ImgUrls = append(houseDetail.ImgUrls, "http://127.0.0.1:8080/"+houseInfo.IndexImageUrl)
		}
	}

	//评论在order表
	var orders []model.OrderHouse
	if err := DB.Model(&houseInfo).Preload("orders").Find(&orders).Error; err != nil {
		fmt.Println("查询房屋评论信息", err)
		return respData, err
	}
	//var comments []*house.CommentData
	for _, v := range orders {
		var commentTemp house.CommentData
		commentTemp.Comment = v.Comment
		commentTemp.Ctime = v.CreatedAt.Format("2006-01-02 15:04:05")
		var tempUser model.User
		DB.Model(&v).Preload("UserId").Find(&tempUser)
		commentTemp.UserName = tempUser.Name

		houseDetail.Comments = append(houseDetail.Comments, &commentTemp)
	}

	//获取房屋的家具信息  多对多查询
	var facs []model.Facility
	if err := DB.Model(&houseInfo).Preload("Facilities").Find(&facs).Error; err != nil {
		fmt.Println("查询房屋家具信息错误", err)
		return respData, err
	}
	for _, v := range facs {
		houseDetail.Facilities = append(houseDetail.Facilities, int32(v.Id))
	}

	//获取副图片  幅图找不到算不算错
	var imgs []model.HouseImage
	if err := DB.Model(&houseInfo).Preload("Images").Find(&imgs).Error; err != nil {
		fmt.Println("该房屋只有主图", err)
	}

	for _, v := range imgs {
		if len(imgs) != 0 {
			houseDetail.ImgUrls = append(houseDetail.ImgUrls, "http://127.0.0.1:8080/"+v.Url)
		}
	}

	//获取房屋所有者信息
	var user model.User
	if err := DB.Model(&houseInfo).Preload("UserId").Find(&user).Error; err != nil {
		fmt.Println("查询房屋所有者信息错误", err)
		return respData, err
	}
	houseDetail.UserName = user.Name
	houseDetail.UserAvatar = "http://127.0.0.1:8080/" + user.AvatarUrl
	houseDetail.UserId = int32(user.ID)

	respData.House = &houseDetail

	//获取当前浏览人信息
	var nowUser model.User
	if err := DB.Where("name = ?", userName).Find(&nowUser).Error; err != nil {
		fmt.Println("查询当前浏览人信息错误", err)
		return respData, err
	}
	respData.UserId = int32(nowUser.ID)
	return respData, nil
}

// GetIndexHouse 获取房屋信息
func GetIndexHouse() ([]*house.Houses, error) {

	var housesResp []*house.Houses

	var houses []model.House
	if err := DB.Limit(5).Find(&houses).Error; err != nil {
		fmt.Println("获取房屋信息失败", err)
		return nil, err
	}

	for _, v := range houses {
		var houseTemp house.Houses
		houseTemp.Address = v.Address
		//根据房屋信息获取地域信息
		var area model.Area
		var user model.User

		DB.Model(&v).Preload("AreaId").Find(&area)
		DB.Model(&v).Preload("UserId").Find(&user)

		houseTemp.AreaName = area.Name
		houseTemp.Ctime = v.CreatedAt.Format("2006-01-02 15:04:05")
		houseTemp.HouseId = int32(v.ID)
		houseTemp.ImgUrl = "http://127.0.0.1:8080/" + v.IndexImageUrl
		houseTemp.OrderCount = int32(v.OrderCount)
		houseTemp.Price = int32(v.Price)
		houseTemp.RoomCount = int32(v.RoomCount)
		houseTemp.Title = v.Title
		houseTemp.UserAvatar = "http://127.0.0.1:8080/" + user.AvatarUrl

		housesResp = append(housesResp, &houseTemp)
	}

	return housesResp, nil
}

func SearchHouse(areaId, sd, ed, sk string) ([]*house.Houses, error) {
	var houseInfos []model.House

	//   minDays  <  (结束时间  -  开始时间) <  max_days
	//计算一个差值  先把string类型转为time类型
	sdTime, _ := time.Parse("2006-01-02", sd)
	edTime, _ := time.Parse("2006-01-02", ed)
	dur := edTime.Sub(sdTime)

	err := DB.Where("area_id = ?", areaId).
		Where("min_days < ?", dur.Hours()/24).
		Where("max_days > ?", dur.Hours()/24).
		Order("created_at desc").Find(&houseInfos).Error
	if err != nil {
		fmt.Println("搜索房屋失败", err)
		return nil, err
	}

	//获取[]*house.Houses
	var housesResp []*house.Houses

	for _, v := range houseInfos {
		var houseTemp house.Houses
		houseTemp.Address = v.Address
		//根据房屋信息获取地域信息
		var area model.Area
		var user model.User

		DB.Model(&v).Preload("AreaId").Find(&area)
		DB.Model(&v).Preload("UreaId").Find(&user)

		houseTemp.AreaName = area.Name
		houseTemp.Ctime = v.CreatedAt.Format("2006-01-02 15:04:05")
		houseTemp.HouseId = int32(v.ID)
		houseTemp.ImgUrl = "http://127.0.0.1:8080/" + v.IndexImageUrl
		houseTemp.OrderCount = int32(v.OrderCount)
		houseTemp.Price = int32(v.Price)
		houseTemp.RoomCount = int32(v.RoomCount)
		houseTemp.Title = v.Title
		houseTemp.UserAvatar = "http://127.0.0.1:8080/" + user.AvatarUrl

		housesResp = append(housesResp, &houseTemp)

	}

	return housesResp, nil
}
