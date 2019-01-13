package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/hashicorp/consul/api"
	"log"
	"micro-gin/consul_api"
)

const (
	Port = 80
	Host = "localhost"
)

func main() {

	r := gin.Default()
	//r.GET("/check", Check)
	r.GET("/register", consul_api.RegisterConsul)
	//r.GET("/discovery", consul_api.Discovery)

	r.Run("0.0.0.0:80")
}

func Check(ctx *gin.Context) {
	c := api.DefaultConfig()
	c.Address = "localhost:8500"
	fmt.Println("default config : ", c)

	client, err := api.NewClient(c)

	if err != nil {
		fmt.Println("consul client err: ", err)
	}

	reg := new(api.AgentServiceRegistration)
	//reg.ID = "text"
	//reg.Port = 80
	reg.Name = "demo"
	//reg.Tags = []string{"test"}
	//reg.Address = "localhost"

	check := new(api.AgentServiceCheck)
	check.HTTP = fmt.Sprintf("http://%s:%d%s", "localhost", 80, "/check")
	check.Timeout = "5s"
	check.Interval = "5s"
	//注册check服务。
	reg.Check = check
	log.Println("get check.HTTP:", check)

	err = client.Agent().ServiceRegister(reg)

	if err != nil {
		log.Fatal("register server error : ", err)
	}

}
