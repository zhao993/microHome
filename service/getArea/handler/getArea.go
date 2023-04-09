package handler

import (
	"context"
	"getArea/dao/redis"
	pb "getArea/proto"
	"getArea/utils"
)

type GetArea struct{}

func (e *GetArea) MicroGetArea(ctx context.Context, req *pb.Request, rsp *pb.Response) error {
	areas, err := redis.GetArea()
	if err != nil {
		rsp.ErrCode = string(utils.RecodeDataErr)
		return err
	}

	for _, v := range areas {
		var areaInfo pb.AreaInfo
		areaInfo.Aid = int32(v.Id)
		areaInfo.Aname = v.Name
		rsp.ErrCode = string(utils.RecodeOk)
		rsp.Data = append(rsp.Data, &areaInfo)
	}
	return nil
}
