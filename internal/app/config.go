package app

import (
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/webtoor/go-rest-api/internal/consts"
	"github.com/webtoor/go-rest-api/pkg/util"
)

var (
	once sync.Once
	_cfg *Config
)

type Config struct {
	App    *App      `json:"app"`
	Logger Logging   `json:"logger"`
	DB     *Database `json:"database"`
}

type App struct {
	Name     string `json:"name"`
	Debug    bool   `json:"debug"`
	Timezone string `json:"timezone"`
	Env      string `json:"env"`
	Port     int    `json:"port"`
}

type Logging struct {
	Level string `json:"level"`
}

type Database struct {
	Name         string        `json:"name"`
	User         string        `json:"user"`
	Pass         string        `json:"pass"`
	Host         string        `json:"host"`
	Port         int           `json:"port"`
	MaxOpen      int           `json:"max_open"`
	MaxIdle      int           `json:"max_idle"`
	DialTimeout  time.Duration `json:"dial_timeout"`
	MaxLifeTime  time.Duration `json:"max_life_time"`
	ReadTimeout  time.Duration `json:"read_timeout"`
	WriteTimeout time.Duration `json:"write_timeout"`
	Charset      string        `json:"charset"`
	Driver       string        `json:"driver"`
	Timezone     string        `json:"timezone"`
}

func NewConfig() *Config {
	fpath := []string{consts.ConfigPath}
	once.Do(func() {
		c, err := readCfg("app.json", fpath...)
		if err != nil {
			log.Fatal(err)
		}

		_cfg = c
	})

	return _cfg
}

func readCfg(fname string, ps ...string) (*Config, error) {
	var cfg *Config
	var errs []error

	for _, p := range ps {
		f := fmt.Sprint(p, fname)

		err := util.ReadFromJSON(f, &cfg)
		if err != nil {
			errs = append(errs, fmt.Errorf("file %s error %s", f, err.Error()))
			continue
		}
		break
	}

	if cfg == nil {
		return nil, fmt.Errorf("file config parse error %v", errs)
	}

	return cfg, nil
}
