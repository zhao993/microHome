package redis

// SaveImgCode 存储图片id到redis数据库
func SaveImgCode(code, uuid string) (err error) {
	//1.连接数据库
	conn := RedisPool.Get()
	defer conn.Close()
	//2.操作数据库,有效时间5分钟
	_, err = conn.Do("setex", uuid, 60*5, code)
	return
}
