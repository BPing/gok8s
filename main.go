package main

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"log"
	"net/http"
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
		log.Fatal(err)
		return
	}
}

func statusOKHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "success"})
}

func versionHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"version": config.Version})
}

func fileHandler(c *gin.Context) {
	data := c.GetString("content")
	if data == "" {
		data = "hello world [pv]"
	}
	err := ioutil.WriteFile("./pv/output.txt", []byte(data), 0666)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{"status": data})
	}
}

func main() {
	router := gin.New()
	router.Use(gin.Recovery())
	router.GET("/ping", statusOKHandler)
	router.GET("/version", versionHandler)
	router.GET("/file", fileHandler)
	router.Run(config.Addr)
}
