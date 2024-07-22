package configs

import (
	"fmt"
	"learning/grpc/pkg/service"
	"learning/grpc/util/protoc/pb"
	"log"
	"time"

	"github.com/fsnotify/fsnotify"
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
	grpc_handler "learning/grpc/pkg/handler/grpc"
	"google.golang.org/grpc"
)

type Config struct {
	App struct {
		Debug    bool   `mapstructure:"APP_DEBUG"`
		Host     string `mapstructure:"APP_HOST"`
		GRPCPort string `mapstructure:"APP_GRPC_PORT"`
		Version  string `mapstructure:"APP_VERSION"`
		ID       string `mapstructure:"APP_ID"`
		Language string `mapstructure:"APP_LANGUAGE"`
		TZ       string `mapstructure:"APP_TIMEZONE"`
	}
	configPath string // config path
	configName string // config name
	configType string // config type
}

func NewConfig(configPath, configName, configType string) *Config {
	return &Config{
		configPath: configPath,
		configName: configName,
		configType: configType,
	}
}

func init() {
	fmt.Println("load env")
	if err := godotenv.Load(); err != nil {
		log.Fatalf("failed to load env %v", err.Error())
	}

	viper.SetConfigFile(".env")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("failed read config file %v", err.Error())
	}
}

func (c *Config) bind() {
	c.App.Debug = viper.GetBool("APP_DEBUG")
	c.App.Host = viper.GetString("APP_HOST")
	c.App.GRPCPort = viper.GetString("APP_GRPC_PORT")
	c.App.Version = viper.GetString("APP_VERSION")
	c.App.ID = viper.GetString("APP_ID")
	c.App.Language = viper.GetString("APP_LANGUAGE")
	c.App.TZ = viper.GetString("APP_TIMEZONE")

	loc, err := time.LoadLocation(c.App.TZ)
	if err != nil {
		log.Fatalf("failed set timezone %v", err.Error())
	}
	time.Local = loc

}

func (c *Config) LoadConfig() {
	c.bind()
	viper.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("config file changed:", e.Name)
		if err := viper.ReadInConfig(); err != nil {
			log.Fatalf("failed read config file %v", err.Error())
		}
		c.bind()
	})

	viper.WatchConfig()
}

func (c *Config) InitConnection() {

}


func (c *Config) InitGRPCService(server grpc.ServiceRegistrar) {
	userSvc := service.NewUserService()
	pb.RegisterUserServer(server, grpc_handler.NewGRPCHandler(userSvc))
}

func (c *Config) Teardown() {
	
}