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
	App: App{Name: "Poulo", Mode: "release", Version: "v0.0.1", FileServerAddr: "http://localhost:51730"},
	Dir: Dir{},
}

type App struct {
	Name           string `json:"name" yaml:"name"`
	Mode           string `json:"mode" yaml:"mode"`
	Version        string `json:"version" yaml:"version"`
	FileServerAddr string `json:"file_server_addr" yaml:"file_server_addr"`
}

type Dir struct {
	ConfigDir string `json:"config_dir" yaml:"config_dir"`
	BaseDir   string `json:"base_dir" yaml:"base_dir"`
}

type Config struct {
	App App `json:"app" yaml:"app"`
	Dir Dir `json:"dir" yaml:"dir"`
}

func (d *Dir) makeDirAll() error {
	err := os.MkdirAll(d.ConfigDir, os.ModePerm)
	if err != nil {
		return err
	}

	var dirs = []string{"cache", "log", "data", "tmp"}
	for _, dir := range dirs {
		err := os.MkdirAll(d.BaseDir+Sep+dir, os.ModePerm)
		if err != nil {
			return err
		}
	}
	return nil
}

// Init 初始化配置
func Init() (cfg *Config, err error) {
	dir := getDefaultDir()

	_, err = os.Stat(dir.ConfigDir + Sep + CfgFileName) // ~/Library/Preferences/Poulo/config.yml
	if os.IsNotExist(err) {                             //配置文件不存在,创建默认配置文件
		err := dir.makeDirAll()
		if err != nil {
			return nil, err
		}

		defaultCfg.Dir = dir
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

func getDefaultDir() (dir Dir) {
	// export poulo_mode=debug
	if os.Getenv("poulo_mode") == "debug" {
		dir.ConfigDir = "./config"
		dir.BaseDir = "./dir"
		return
	}

	switch runtime.GOOS {
	case "darwin":
		home := os.Getenv("HOME") ///Users/xxx
		dir.ConfigDir = home + "/Library/Preferences/Poulo"
		dir.BaseDir = home + "/Library/Caches/Poulo"
	case "windows":
		dir.ConfigDir = os.Getenv("APPDATA") + "\\Poulo"    //C:\Users\{UserName}\AppData\Roaming
		dir.BaseDir = os.Getenv("LOCALAPPDATA") + "\\Poulo" //C:\Users\{UserName}\AppData\Local
	}
	return
}

func (c *Config) load() error {
	f, err := os.Open(c.Dir.ConfigDir + Sep + CfgFileName)
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
	f, err := os.Create(c.Dir.ConfigDir + Sep + CfgFileName)
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
