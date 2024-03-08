package config

import (
	"gopkg.in/yaml.v3"
	"io"
	"os"
	"path/filepath"
	"runtime"
)

const (
	CfgFileName = "config.yml"

	Sep = string(filepath.Separator)
)

var defaultCfg = Config{
	Name:           "Poulo",
	Mode:           "release",
	Version:        "v0.0.1",
	FileServerPort: 51730,
}

type Config struct {
	Name           string `json:"name" yaml:"name"`
	Mode           string `json:"mode" yaml:"mode"`
	Version        string `json:"version" yaml:"version"`
	ConfigPath     string `json:"config_path" yaml:"config_path"`
	BaseDir        string `json:"base_dir" yaml:"base_dir"`
	FileServerPort int    `json:"file_server_port" yaml:"file_server_port"`
}

func (c *Config) makeDirAll() error {
	err := os.MkdirAll(c.ConfigPath, os.ModePerm)
	if err != nil {
		return err
	}

	var dirs = []string{"cache", "log", "data", "tmp"}
	for _, dir := range dirs {
		err := os.MkdirAll(c.BaseDir+Sep+dir, os.ModePerm)
		if err != nil {
			return err
		}
	}
	return nil
}

// Init 初始化配置
func Init() (cfg *Config, err error) {
	defaultCfg.loadPath()
	_, err = os.Stat(defaultCfg.ConfigPath + Sep + CfgFileName) // ~/Library/Preferences/Poulo/config.yml
	if os.IsNotExist(err) {                                     //配置文件不存在,创建默认配置文件
		err := defaultCfg.makeDirAll()
		if err != nil {
			return nil, err
		}

		err = defaultCfg.save()
		if err != nil {
			return nil, err
		}
	} else if err == nil { //配置文件存在,加载配置文件
		err = defaultCfg.load()
		if err != nil {
			return nil, err
		}
	} else {
		return nil, err //其他错误
	}

	return &defaultCfg, nil
}

func (c *Config) loadPath() {
	// export poulo_mode=debug
	if os.Getenv("poulo_mode") == "debug" {
		c.Mode = "debug"
		c.ConfigPath = "./config"
		c.BaseDir = "./dir"
		return
	}

	switch runtime.GOOS {
	case "darwin":
		c.ConfigPath = "~/Library/Preferences/Poulo"
		c.BaseDir = "~/Library/Caches/Poulo"
	case "windows":
		c.ConfigPath = os.Getenv("APPDATA") + "\\Poulo"   //C:\Users\{UserName}\AppData\Roaming
		c.BaseDir = os.Getenv("LOCALAPPDATA") + "\\Poulo" //C:\Users\{UserName}\AppData\Local
	}
	return
}

func (c *Config) load() error {
	f, err := os.Open(c.ConfigPath + Sep + CfgFileName)
	if err != nil {
		return err
	}
	defer f.Close()

	bytes, err := io.ReadAll(f)
	if err != nil {
		return err
	}

	var cfg Config
	err = yaml.Unmarshal(bytes, &cfg)
	if err != nil {
		panic(err)
	}

	return nil
}

func (c *Config) save() error {
	f, err := os.Create(c.ConfigPath + Sep + CfgFileName)
	if err != nil {
		return err
	}
	defer f.Close()

	bytes, err := yaml.Marshal(c)
	if err != nil {
		return err
	}

	_, err = f.Write(bytes)
	if err != nil {
		return err
	}

	return nil
}
