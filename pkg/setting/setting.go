package setting

import (
	"log"
	"time"

	"github.com/go-ini/ini"
)

type App struct {
	JwtSecret       string
	PageSize        int
	RuntimeRootPath string

	ImagePrefixUrl string
	ImageSavePath  string
	ImageMaxSize   int
	ImageAllowExts []string

	LogSavePath string
	LogSaveName string
	LogFileExt  string
	TimeFormat  string
}

var AppSetting = &App{}

type Server struct {
	RunMode      string
	HTTPPort     int
	ReadTimeout  time.Duration
	WriteTimeOut time.Duration
}

var ServerSetting = &Server{}

type Database struct {
	Type        string
	User        string
	Password    string
	Host        string
	Name        string
	TablePrefix string
}

var DatabaseSetting = &Database{}

type Redis struct {
	Host        string
	Password    string
	MaxIdle     int
	MaxActive   int
	IdelTimeout time.Duration
}

var RedisSetting = &Redis{}

func Setup() {
	Cfg, err := ini.Load("conf/app.ini")

	log.Println("开始解析配置文件conf/app.ini")
	defer log.Println("配置文件conf/app.ini解析完成")

	if err != nil {
		log.Fatalf("conf/app.ini配置文件有误,错误信息:%v", err)
	}

	err = Cfg.Section("app").MapTo(AppSetting)
	if err != nil {
		log.Fatalf("conf/app.ini配置文件[app]中内容有误,错误信息:%v", err)
	}
	//转换为MB
	AppSetting.ImageMaxSize = AppSetting.ImageMaxSize * 1024 * 1024

	err = Cfg.Section("server").MapTo(ServerSetting)
	if err != nil {
		log.Fatalf("conf/app.ini配置文件[server]中内容有误,错误信息:%v", err)
	}

	ServerSetting.ReadTimeout = ServerSetting.ReadTimeout * time.Second
	ServerSetting.WriteTimeOut = ServerSetting.WriteTimeOut * time.Second

	err = Cfg.Section("database").MapTo(DatabaseSetting)
	if err != nil {
		log.Fatalf("conf/app.ini配置文件[database]中内容有误,错误信息:%v", err)
	}

	err = Cfg.Section("redis").MapTo(RedisSetting)
	if err != nil {
		log.Fatalf("conf/app.ini配置文件[redis]中内容有误,错误信息:%v", err)
	}
}
