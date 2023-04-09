package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"getCaptcha/dao/redis"
	"github.com/afocus/captcha"
	"image/color"

	pb "getCaptcha/proto"
)

type GetCaptcha struct{}

func (e *GetCaptcha) Call(ctx context.Context, req *pb.Request, rsp *pb.Response) error {
	//生成图片验证码
	//初始化对象
	cap := captcha.New()
	//设置字体
	cap.SetFont("./utils/comic.ttf")
	//设置验证码大小
	cap.SetSize(128, 64)
	//设置干扰强度
	cap.SetDisturbance(captcha.MEDIUM)
	//设置前景色
	cap.SetFrontColor(color.RGBA{0, 0, 0, 255})
	//设置背景色
	cap.SetBkgColor(color.RGBA{0, 128, 128, 128}, color.RGBA{255, 255, 10, 128})
	// 生成字体
	img, str := cap.Create(4, captcha.ALL)
	//存储图片验证码到redis中
	err := redis.SaveImgCode(str, req.Uuid)
	if err != nil {
		fmt.Println("save img code error:", err)
		return err
	}
	//将生成的图片进行序列化
	imgBuf, err := json.Marshal(img)
	if err != nil {
		return err
	}
	//使用参数rsp传出
	rsp.Img = imgBuf
	return nil
}
