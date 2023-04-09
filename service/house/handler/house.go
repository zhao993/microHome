package handler

import (
	"context"
	"github.com/tedcy/fdfs_client"
	"house/dao/mysql"
	pb "house/proto"
	"house/utils"
	"strconv"
)

type House struct{}

func (e *House) PubHouse(ctx context.Context, req *pb.Request, rsp *pb.Response) error {
	//上传房屋业务  把获取到的房屋数据插入数据库
	houseId, err := mysql.AddHouse(req)
	if err != nil {
		rsp.ErrCode = string(utils.RecodeDbErr)
		return nil
	}

	rsp.ErrCode = string(utils.RecodeOk)
	var h pb.HouseData
	h.HouseId = strconv.Itoa(houseId)
	rsp.Data = &h
	return nil
}

func (e *House) UploadHouseImg(ctx context.Context, req *pb.ImgReq, rsp *pb.ImgResp) error {
	//把图片存储到fastdfs中
	//初始化fdfs的客户端
	fClient, _ := fdfs_client.NewClientWithConfig("/etc/fdfs/client.conf")
	//上传图片到fdfs
	fdfsResp, err := fClient.UploadByBuffer(req.ImgData, req.FileExt[1:])
	if err != nil {
		rsp.ErrCode = string(utils.RecodeDbErr)
		return nil
	}

	//把凭证存储到数据库中
	err = mysql.SaveHouseImg(req.HouseId, fdfsResp)
	if err != nil {
		rsp.ErrCode = string(utils.RecodeDbErr)
		return nil
	}

	rsp.ErrCode = string(utils.RecodeOk)
	var img pb.ImgData
	img.Url = "http://192.168.17.129:8888/" + fdfsResp
	rsp.Data = &img
	return nil

}

func (e *House) GetHouseInfo(ctx context.Context, req *pb.GetReq, rsp *pb.GetResp) error {
	//根据用户名获取所有的房屋数据
	houseInfos, err := mysql.GetUserHouse(req.UserName)
	if err != nil {
		rsp.ErrCode = string(utils.RecodeDbErr)
		return nil
	}

	var getData pb.GetData
	getData.Houses = houseInfos
	rsp.Data = &getData
	rsp.ErrCode = string(utils.RecodeOk)
	return nil
}

func (e *House) GetHouseDetail(ctx context.Context, req *pb.DetailReq, rsp *pb.DetailResp) error {
	//根据houseId获取所有的返回数据
	respData, err := mysql.GetHouseDetail(req.HouseId, req.UserName)
	if err != nil {
		rsp.ErrCode = string(utils.RecodeDbErr)
		return nil
	}
	rsp.ErrCode = string(utils.RecodeOk)
	rsp.Data = &respData
	return nil
}

func (e *House) GetIndexHouse(ctx context.Context, req *pb.IndexReq, rsp *pb.GetResp) error {
	//获取房屋信息
	houseResp, err := mysql.GetIndexHouse()
	if err != nil {
		rsp.ErrCode = string(utils.RecodeDbErr)
		return nil
	}

	rsp.ErrCode = string(utils.RecodeOk)
	rsp.Data = &pb.GetData{Houses: houseResp}

	return nil
}

func (e *House) SearchHouse(ctx context.Context, req *pb.SearchReq, rsp *pb.GetResp) error {
	//根据传入的参数,查询符合条件的房屋信息
	houseResp, err := mysql.SearchHouse(req.Aid, req.Sd, req.Ed, req.Sk)
	if err != nil {
		rsp.ErrCode = string(utils.RecodeDbErr)
		return nil
	}
	rsp.ErrCode = string(utils.RecodeOk)
	rsp.Data = &pb.GetData{Houses: houseResp}
	return nil
}
