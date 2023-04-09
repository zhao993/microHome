# microHome
Go-Micro微服务实战项目-iHome租房网项目。
功能:用户注册，用户登录，头像上传，用户详细信息获取，实名认证检测，房源发布，首页展示，搜索房源，订单管理，用户评价等模块。

注:个人学习微服务使用

## 技术栈

golang + consul服务发现 +zap日志库+ grpc + protobuf + gin + gorm + mysql + redis + fastDFS + nginx

### 功能模块

####  用户模块

 注册
 获取验证码图片服务
 获取短信验证码服务
 发送注册信息服务
 登录
 获取session信息服务
 获取登录信息服务
 退出
 个人信息获取
 获取用户基本信息服务
 更新用户名服务
 发送上传用户头像服务
 实名认证
 获取用户实名信息服务
 发送用户实名认证信息服务

#### 房屋模块

 首页展示
 获取首页轮播图服务
 房屋详情
 地区列表
 房屋搜索
 订单模块
 订单确认
 发布订单
 查看订单信息
 订单评论

consul启动：

开发测试过程中可以使用单机模式
consul agent -dev

FastDFS服务启动

sudo fdfs_trackerd /etc/fdfs/tracker.conf
sudo fdfs_storaged /etc/fdfs/storage.conf
nginx

启动nginx
sudo /usr/local/nginx/sbin/nginx

## 项目布局

分为服务端和web端
service
getArea     获取地区信息微服务
getCaptcha  获取验证码图片微服务
register   用户注册微服务，包括发送短信验证码，注册和登录业务
user       用户相关微服务，包括展示用户信息，更新用户名，上传用户头像和实名认证业务
house      房子相关的微服务，包括发布房源信息，上传房源图片，获取详细信息，按地区搜索房子等业务
userOrder  用户订单微服务，包括创建用户订单，获得订单详情等业务
web
conf        web端配置文件
controller   处理业务层
dao          数据库相关
log          日志文件
loggerm      日志配置相关业务
model        数据库模型
proto        与服务端建立连接的protoc文件
router       路由模块
setting      读取配置文件模块
test         测试文件
utils        一些工具类文件
view        前端页面展示
main.go
README.md
