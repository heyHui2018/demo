package main

import (
	"github.com/go-ini/ini"
	"strings"
	"reflect"
	"errors"
	"fmt"
	"os"
)

var ConfigInstance *Config

type FtpConfig struct {
	Host string
	Port string
}

type LogConfig struct {
	Path  string
	Level string
}

type Config struct {
	Ftp FtpConfig
	Log LogConfig
}

func main() {
	ConfigInit()
	fmt.Println("done-----------------------------------------")
}

func ConfigInit() {
	separator := "/"
	if os.IsPathSeparator('\\') {
		separator = "\\"
	} else {
		separator = "/"
	}
	config_file := "config.conf"
	config_file = strings.Replace(config_file, "/", separator, -1)
	c, err := ini.Load(config_file)
	if err != nil {
		panic(err)
	}
	ConfigInstance = new(Config)
	//获取结构体中元素一般通过反射，如下所示
	configType := reflect.TypeOf(*ConfigInstance)//ConfigInstance为指针，没有NumField方法，必须*ConfigInstance
	configValue := reflect.ValueOf(ConfigInstance).Elem()

	for i := 0; i < configType.NumField(); i++ {
		fieldType := configType.Field(i)
		if fieldType.Type.Kind() == reflect.Struct {
			fieldValue := configValue.Field(i)
			sectionKey := strings.ToLower(fieldType.Name)
			section := c.Section(sectionKey)

			for k := 0; k < fieldType.Type.NumField(); k ++ {
				innerField := fieldType.Type.Field(k)
				innerValue := fieldValue.Field(k)
				key := strings.ToLower(string(innerField.Name[0])) + innerField.Name[1:]
				kv, err := section.GetKey(key)
				if err != nil {
					panic(err)
				}
				switch innerField.Type.Kind() {
				case reflect.String:
					innerValue.SetString(kv.String())
				case reflect.Int, reflect.Int8, reflect.Int32, reflect.Int64:
					intValue, err := kv.Int64()
					if err != nil {
						panic(err)
					}
					innerValue.SetInt(intValue)
				default:
					panic(errors.New("未支持的类型：" + innerField.Type.Kind().String()))
				}
			}
		}
	}
	fmt.Printf("config = %+v", *ConfigInstance)
}
