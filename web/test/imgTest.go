package test

import (
	"github.com/afocus/captcha"
	"image/color"
)

func imgTest() {
	//初始化对象
	cap := captcha.New()
	//设置字体
	cap.SetFont("comic.ttf")
	//设置验证码大小
	cap.SetSize(128, 64)
	//设置干扰强度
	cap.SetDisturbance(captcha.MEDIUM)
	//设置前景色
	cap.SetFrontColor(color.RGBA{0, 0, 0, 255})
	//设置背景色
	cap.SetBkgColor(color.RGBA{0, 128, 128, 128}, color.RGBA{255, 255, 10, 128})
	//生成字体，将图片验证码展示到页面
	cap.Create(4, captcha.ALL)
}
