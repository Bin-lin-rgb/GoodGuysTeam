package Utils

import (
	"gopkg.in/ini.v1"
	"log"
)

var (
	AppMode  string
	HttpPort string
	JwtKey   string

	DbHost     string
	DbPort     string
	DbUser     string
	DbPassWord string
	DbName     string

	Level      string
	Filename   string
	MaxSize    int
	MaxAge     int
	MaxBackups int
)

func init() {
	file, err := ini.Load("config/config.ini")
	if err != nil {
		log.Fatal("Configuration file read error", err)
	}
	LoadServer(file)
	LoadData(file)
	LoadLogConfig(file)
}

func LoadServer(file *ini.File) {
	// MustString 默认值
	AppMode = file.Section("server").Key("AppMode").MustString("debug")
	HttpPort = file.Section("server").Key("HttpPort").MustString(":8080")
	JwtKey = file.Section("server").Key("JwtKey").MustString("89js82js72")
}

func LoadData(file *ini.File) {
	DbHost = file.Section("database").Key("DbHost").MustString("127.0.0.1")
	DbPort = file.Section("database").Key("DbPort").MustString("3306")
	DbUser = file.Section("database").Key("DbUser").MustString("blogserver")
	DbPassWord = file.Section("database").Key("DbPassWord").MustString("Admin666")
	DbName = file.Section("database").Key("DbName").MustString("root")
}

func LoadLogConfig(file *ini.File) {
	Level = file.Section("logConfig").Key("Level      ").MustString("debug")
	Filename = file.Section("logConfig").Key("Filename   ").MustString("default.log")
	MaxSize = file.Section("logConfig").Key("MaxSize    ").MustInt(200)
	MaxAge = file.Section("logConfig").Key("MaxAge     ").MustInt(7)
	MaxBackups = file.Section("logConfig").Key("MaxBackups ").MustInt(30)
}
