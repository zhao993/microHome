package handler

import (
	"context"
	"fmt"
	"github.com/tedcy/fdfs_client"
	"user/dao/mysql"
	pb "user/proto"
	"user/utils"
)

type User struct{}

func (e *User) UserInfo(ctx context.Context, req *pb.NameData, rsp *pb.Response) error {
	user, err := mysql.GetUserInfo(req.Name)
	if err != nil {
		rsp.ErrCode = string(utils.RecodeDbErr)
		return err
	}
	rsp.ErrCode = string(utils.RecodeOk)
	var userInfo pb.UserInfo
	userInfo.UserId = user.ID
	userInfo.Name = user.Name
	userInfo.Mobile = user.Mobile
	userInfo.RealName = user.RealName
	userInfo.IdCard = user.IDCard
	userInfo.AvatarUrl = user.AvatarUrl
	rsp.Data = &userInfo
	return nil
}
func (e *User) UpdateUserName(ctx context.Context, req *pb.UpdateReq, rsp *pb.UpdateResp) error {
	err := mysql.UpdateUserName(req.OldName, req.NewName)
	if err != nil {
		rsp.ErrCode = string(utils.RecodeDbErr)
		return err
	}
	var nameData pb.NameData
	nameData.Name = req.NewName
	rsp.Data = &nameData
	rsp.ErrCode = string(utils.RecodeOk)
	return nil
}
func (e *User) UploadAvatar(ctx context.Context, req *pb.UploadReq, rsp *pb.UploadResp) error {
	fClient, _ := fdfs_client.NewClientWithConfig("/etc/fdfs/client.conf")
	//上传文件到fdfs
	fdfsResp, err := fClient.UploadByBuffer(req.Avatar, req.FileExt[1:])
	if err != nil {
		fmt.Println("上传文件错误", err)
		rsp.ErrCode = string(utils.RecodeDbErr)
		return nil
	}

	//把存储凭证写入数据库
	err = mysql.SaveUserAvatar(req.UserName, fdfsResp)
	if err != nil {
		fmt.Println("存储用户头像错误", err)
		rsp.ErrCode = string(utils.RecodeDbErr)
		return nil
	}

	rsp.ErrCode = string(utils.RecodeOk)
	var uploadData pb.UploadData
	uploadData.AvatarUrl = "http://192.168.17.192:8888/" + fdfsResp
	rsp.Data = &uploadData
	return nil
}
func (e *User) AuthUpdate(ctx context.Context, req *pb.AuthReq, rsp *pb.CallResponse) error {
	err := mysql.UpdateAuth(req.Name, req.RealName, req.IdCard)
	if err != nil {
		rsp.ErrCode = string(utils.RecodeDbErr)
		return err
	}
	rsp.ErrCode = string(utils.RecodeOk)
	return nil
}
