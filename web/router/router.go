package router

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/redis"
	"github.com/gin-gonic/gin"
	"microHome/web/controller"
	"microHome/web/utils"
)

func loginFilter() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		username := session.Get("username")
		if username == "" {
			c.Abort() //从这里返回，不继续执行
			controller.ResponseError(c, utils.RecodeSessionErr)
		} else {
			c.Next()
		}
	}
}

func SetUp(mode string) *gin.Engine {
	if mode == gin.ReleaseMode {
		gin.SetMode(gin.ReleaseMode)
	}

	r := gin.Default()
	store, _ := redis.NewStore(10, "tcp", "localhost:6379", "", []byte("secret"))
	r.Use(sessions.Sessions("mysession", store))
	r.Static("/home", "web/view")
	v1 := r.Group("/api/v1.0")
	{
		v1.GET("/session", controller.GetSession)
		v1.GET("/imagecode/:uuid", controller.GetImageCd)
		v1.GET("/smscode/:phone", controller.GetSmsCd)
		v1.POST("/users", controller.PostRet)
		v1.GET("/areas", controller.GetArea)
		v1.POST("/sessions", controller.PostLogin)

		v1.Use(loginFilter()) //以后的路由不需要校验session
		v1.DELETE("/session", controller.DeleteSession)
		v1.GET("/user", controller.GetUserInfo)
		v1.PUT("/user/name", controller.PutUserInfo)
		v1.POST("/user/avatar", controller.PostAvatar)
		v1.GET("/user/auth", controller.GetUserInfo)
		v1.POST("/user/auth", controller.PostAuth)
		//获取已发布房源信息
		v1.GET("/user/houses", controller.GetUserHouses)
		//发布房源
		v1.POST("/houses", controller.PostHouses)
		//添加房源图片
		v1.POST("/houses/:id/images", controller.PostHousesImage)
		//展示房屋详情
		v1.GET("/houses/:id", controller.GetHouseInfo)
		//展示首页轮播图
		v1.GET("/house/index", controller.GetIndex)
		//搜索房屋
		v1.GET("/houses", controller.GetHouses)
		//下订单
		v1.POST("/orders", controller.PostOrders)
		//获取订单
		v1.GET("/user/orders", controller.GetUserOrder)
		//同意/拒绝订单
		v1.PUT("/orders/:id/status", controller.PutOrders)
	}
	return r
}
