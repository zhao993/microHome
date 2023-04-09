package setting

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"time"
)

var Conf *Config

type Config struct {
	AppName      string `mapstructure:"name"`
	AppMode      string `mapstructure:"mode"`
	AppPort      string `mapstructure:"port"`
	AppVersion   string `mapstructure:"version"`
	*LogConfig   `mapstructure:"log"`
	*MysqlConfig `mapstructure:"mysql"`
	*RedisConfig `mapstructure:"redis"`
}
type LogConfig struct {
	LogLevel      string `mapstructure:"level"`
	LogPath       string `mapstructure:"filepath"`
	LogMaxSize    int    `mapstructure:"max_size"`
	LogMaxAge     int    `mapstructure:"max_age"`
	LogMaxBackups int    `mapstructure:"max_backups"`
}
type MysqlConfig struct {
	MysqlHost    string `mapstructure:"host"`
	MysqlPort    string `mapstructure:"port"`
	MysqlUser    string `mapstructure:"user"`
	MysqlPass    string `mapstructure:"password"`
	MysqlDB      string `mapstructure:"dbname"`
	MysqlMaxConn int    `mapstructure:"max_idle_conns"`
	MysqlMaxOpen int    `mapstructure:"max_open_conns"`
}
type RedisConfig struct {
	RedisHost       string        `mapstructure:"host"`
	RedisPort       string        `mapstructure:"port"`
	RedisPass       string        `mapstructure:"password"`
	RedisDb         int           `mapstructure:"db"`
	MaxIdle         int           `mapstructure:"max_idle"`
	MaxActive       int           `mapstructure:"max_active"`
	MaxConnLifetime time.Duration `mapstructure:"max_conn_lifetime"`
	IdleTimeout     time.Duration `mapstructure:"idle_timeout"`
}

func Init() (err error) {
	viper.SetConfigFile("web/conf/conf.yaml") // 读取配置文件
	if err != nil {
		fmt.Println("读取配置文件失败：", err)
		return
	}
	err = viper.ReadInConfig() // 加载配置文件
	if err != nil {
		fmt.Println("加载配置文件失败：", err)
		return
	}
	err = viper.Unmarshal(&Conf)
	if err != nil {
		fmt.Println("解析配置文件失败：", err)
		return
	}
	viper.WatchConfig() //监听配置文件变化
	viper.OnConfigChange(func(in fsnotify.Event) {
		fmt.Println("配置文件被修改了")
		viper.Unmarshal(&Conf) //重新解析配置文件
	})
	return
}
