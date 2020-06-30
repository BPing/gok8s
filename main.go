package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Config struct {
	Addr    string
	Version string
}

var config *Config

func init() {
	loadConfig()
}

func loadConfig() {
	config = &Config{}
	data, err := ioutil.ReadFile("./conf/config.json")
	if err != nil {
		return
	}
	//读取的数据为json格式，需要进行解码
	err = json.Unmarshal(data, config)
	if err != nil {
		return
	}
}

func statusOKHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "success"})
}

func versionHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"version": config.Version})
}

func main() {
	router := gin.New()
	router.Use(gin.Recovery())
	router.GET("/ping", statusOKHandler)
	router.GET("/version", versionHandler)
	router.Run(config.Addr)
}
