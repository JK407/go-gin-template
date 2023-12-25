package conf

import (
	"fmt"
	"gin-template/constants"
	"gin-template/utils/getenv"
	"github.com/spf13/viper"
	"github.com/yunduan16/micro-service-go-component-redisx"
	"gorm.io/gorm/logger"
	"time"
)

// RunMode 运行模式
const (
	DEV  = "dev"
	GRAY = "gray"
	TEST = "test"
	PROD = "prod"
)

type Config struct {
	App   *AppConf
	MySql *MySqlConfig            `yaml:"MySQL"`
	Redis map[string]*RedisConfig `yaml:"Redis"`
}

type AppConf struct {
	Host        string
	Port        int
	Name        string
	RunMode     string
	LogFile     string
	LogTraceKey string
	SourceDir   string //资源目录
}

type MySqlConfig struct {
	Test *MySqLDB `yaml:"test"`
}

type MySqLDB struct {
	Gorm            GormConfig    `yaml:"Gorm"`
	DSN             string        `yaml:"DSN"`
	MaxOpenConns    int           `yaml:"MaxOpenConns"`
	MaxIdleConns    int           `yaml:"MaxIdleConns"`
	ConnMaxLifetime time.Duration `yaml:"ConnMaxLifetime"`
}

type GormConfig struct {
	SlowThreshold             time.Duration   `yaml:"SlowThreshold"`
	Colorful                  bool            `yaml:"Colorful"`
	IgnoreRecordNotFoundError bool            `yaml:"IgnoreRecordNotFoundError"`
	LogLevel                  logger.LogLevel `yaml:"LogLevel"`
}

type RedisDbConfig struct {
	ExpireTime int
}
type RedisConfig struct {
	ConnConf *redisx.RedisConf `yaml:"ConnConf"`
	DbConf   *RedisDbConfig    `yaml:"DbConf"`
}

var (
	config      = new(Config)
	Name        = "YGL"
	Version     string
	BasePath    string
	TokenUserId int
	TokenAppId  int
)

func Get() Config {
	return *config
}

func NewConfig(configPath string) error {
	if configPath == "" {
		configPath = constants.DEFAULT_CONFIG_PATH
	}
	rootBase := getenv.GetRootDir(configPath)
	fmt.Println("root dir：", rootBase)
	v := viper.New()
	v.SetConfigFile(configPath)
	if err := v.ReadInConfig(); err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	if err := v.Unmarshal(config); err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	return nil
}
