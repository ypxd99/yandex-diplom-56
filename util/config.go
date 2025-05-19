package util

import (
	"flag"
	"fmt"
	"log"
	"os"
	"sync"

	"github.com/pkg/errors"
	"gopkg.in/yaml.v3"
)

var (
	onceCFG sync.Once
	config  *Config
	cfgPath = "configuration/config.yaml"
)

type Config struct {
	Logger   LoggerCfg `yaml:"Logger"`
	Server   Server    `yaml:"Server"`
	Postgres Postgres  `yaml:"Postgres"`
	Auth     Auth      `yaml:"Auth"`
}

type Auth struct {
	SecretKey  string `yaml:"SecretKey"`
	CookieName string `yaml:"CookieName"`
}

type Server struct {
	ServerAddress string `yaml:"-"`
	Address       string `yaml:"Address"`
	Port          uint   `yaml:"Port"`
	RTimeout      int64  `yaml:"RTimeout"`
	WTimeout      int64  `yaml:"WTimeout"`
}

type Postgres struct {
	ConnString      string   `yaml:"-"`
	DriverName      string   `yaml:"DriverName"`
	Address         string   `yaml:"Address"`
	DBName          string   `yaml:"DBName"`
	User            string   `yaml:"User"`
	Password        string   `yaml:"Password"`
	MaxConn         int      `yaml:"MaxConn"`
	MaxConnLifeTime int64    `yaml:"MaxConnLifeTime"`
	Trace           bool     `yaml:"Trace"`
	MakeMigration   bool     `yaml:"MakeMigration"`
	SQLKeyWords     []string `yaml:"SQLKeyWords"`
}

func parseConfig(st interface{}, cfgPath string) {
	f, err := os.Open(cfgPath)
	if err != nil {
		log.Fatal(errors.WithMessage(err, "error occurred while opening cfg file"))
	}

	fi, err := f.Stat()
	if err != nil {
		log.Fatal(errors.WithMessage(err, "error occurred while getting file stats"))
	}

	data := make([]byte, fi.Size())
	_, err = f.Read(data)
	if err != nil {
		log.Fatal(errors.WithMessage(err, "error occurred while reading data"))
	}

	err = yaml.Unmarshal(data, st)
	if err != nil {
		log.Fatal(errors.WithMessage(err, "error occurred while unmashaling data"))
	}
}

func GetConfig() *Config {
	onceCFG.Do(func() {
		var (
			conf Config
		)
		parseConfig(&conf, cfgPath)

		flag.StringVar(&conf.Server.ServerAddress, "a", fmt.Sprintf("%s:%d", conf.Server.Address, conf.Server.Port), "HTTP server address")
		flag.StringVar(&conf.Postgres.ConnString, "d", fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=disable", conf.Postgres.User, conf.Postgres.Password, conf.Postgres.Address, conf.Postgres.DBName), "Database connect string")
		flag.Parse()

		if envAddr, exists := os.LookupEnv("SERVER_ADDRESS"); exists {
			conf.Server.ServerAddress = envAddr
		}

		if envDB, exists := os.LookupEnv("DATABASE_DSN"); exists {
			conf.Postgres.ConnString = envDB
		}

		config = &conf
	})

	if config == nil {
		log.Fatal("nil config")
	}

	return config
}
