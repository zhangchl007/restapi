package db
import (
	"fmt"
	"log"
	"os"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)
//define DBConfig struct for DB connection
type DBConfig struct {
	Host     string
	Port     string
	User     string
	DBName   string
	Password string
	SSLMode string
}
//update DBConfig as the parameter
func UpdateDBConfig(Host, User, DBName, Password, SSLMode string, Port int) {
	ConfigTemplate := "dbconfig"
	GOPATH := os.Getenv("GOPATH")
	ConfigPath := GOPATH + "/src/github.com/zhangchl007/restapi/config/"
	viper.SetConfigName(ConfigTemplate)
	viper.AddConfigPath(ConfigPath)
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal(err)
	}

	viper.Set("Host", Host)
	viper.Set("User", User)
	viper.Set("DBName", DBName)
	viper.Set("Password", Password)
	viper.Set("SSLMode", SSLMode)
	viper.Set("Port", Port)
	viper.WriteConfig()
	fmt.Printf("The DBinfo file: %s had been changed successfully!\n", ConfigPath+ConfigTemplate+".yaml")
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("配置发生变更：", e.Name)
	})
}

//Get DBConfig  configuration
func GetDBConfig() (*DBConfig) {
	GOPATH := os.Getenv("GOPATH")
	ConfigTemplate := "dbconfig"
	ConfigPath := GOPATH + "/src/github.com/zhangchl007/restapi/config/"
	viper.SetConfigName(ConfigTemplate)
	viper.AddConfigPath(ConfigPath)
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal(err)
	}
	var tc DBConfig
	tc.Host = viper.GetString("Host")
	tc.User = viper.GetString("User")
	tc.DBName = viper.GetString("DBName")
	tc.Password = viper.GetString("Password")
	tc.SSLMode = viper.GetString("SSLMode")
	tc.Port = viper.GetString("Port")
	return &tc
}
