package conf

import (
	"encoding/json"
	"flag"
	"io/ioutil"
	"time"

	log "github.com/sirupsen/logrus"
)

var Mysql struct {
	Hostname       string
	Username       string
	Password       string
	MaxConnNum     int
	MainDatabase   string
}

var RedisConfig struct {
	Host       string
	Prefix     string
	Username   string
	DbName     string
	Pwd        string
	MaxConnNum int
}

var MongoConfig struct {
	DBUrl      string //mongodb://root:mypass@localhost:27017?maxPoolSize=10"
	MaxConnNum int
}

var Config struct {
	Host 			string
	LogPath       	string
	LogName  		string
}

//初始化
func init() {
	dir := flag.String("configPath","","configPath Dir")
	flag.Parse()
	cf := "bin/conf/config.json"
	mysql := "bin/conf/mysql.json"
	redis := "bin/conf/redis.json"
	if *dir != ""{
		cf = *dir+cf
		mysql = *dir+mysql
		redis = *dir+redis
	}


	loadConfig(cf, &Config)
	loadConfig(mysql, &Mysql)
	loadConfig(redis, &RedisConfig)
	LocalLogger(Config.LogPath,Config.LogName,time.Hour*240,time.Hour *24)
}

func loadConfig(configPath string, structConfig interface{}) {
	data, err := ioutil.ReadFile(configPath)
	if err != nil {
		log.Fatal(err)
	}
	err = json.Unmarshal(data, &structConfig)
	if err != nil {
		log.Fatal(err)
	}
}
