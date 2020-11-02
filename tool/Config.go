package tool

import (
    "bufio"
    "encoding/json"
    "os"
)

type Config struct {
    AppName string `json:"appName"`
    AppMode string `json:"appMode"`
    AppHost string `json:"appHost"`
    AppPort string `json:"appPort"`
}

//初始化配置为空
var _cfg *Config = nil

func GetConfig() *Config {
    return _cfg
}

//继续json配置文件
func ParseConfig(path string) (*Config, error) {
    //读取指定配置
    file, err := os.Open(path)

    if err != nil {
        panic(err)
    }
    defer file.Close()

    //创建读取handle
    reader := bufio.NewReader(file)
    //创建json解析器
    decoder := json.NewDecoder(reader)

    if err = decoder.Decode(&_cfg); err != nil {
        return nil, err
    }
    return _cfg, nil
}
