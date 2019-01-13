package consul_api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/hashicorp/consul/api"
	"log"
	"net/http"
)

type R struct {
	ID   string   `json:"id"`
	Port int      `json:"port"`
	Name string   `json:"name"`
	Tags []string `json:"tags"`
	Host string   `json:"host"`
}

func RegisterConsul(ctx *gin.Context) {
	var r R
	bind := ctx.BindJSON(&r)
	if bind != nil {
		log.Fatalf("bindJosn error %s", bind)
		return
	}

	c := api.DefaultConfig()
	c.Address = "localhost:8500"
	fmt.Println("default config : ", c)

	client, err := api.NewClient(c)

	if err != nil {
		fmt.Println("consul client err: ", err)
	}

	reg := &api.AgentServiceRegistration{
		ID: fmt.Sprintf("%v-%v-%v", r.Name, r.Host, r.Port), // 服务节点的名称
		//Name:    fmt.Sprintf("grpc.health.v1.%v", r.Name),        // 服务名称
		Name:    fmt.Sprintf("%v", r.Name),
		Tags:    r.Tags, // tag，可以为空
		Port:    r.Port, // 服务端口
		Address: r.Host, // 服务 IP
		Check: &api.AgentServiceCheck{ // 健康检查
			Interval:                       "5s",                                                    // 健康检查间隔
			HTTP:                           fmt.Sprintf("http://%s:%d%s", r.Host, r.Port, "/check"), // grpc 支持，执行健康检查的地址，service 会传到 Health.Check 函数中
			Timeout:                        "3s",
			DeregisterCriticalServiceAfter: "30s", // 注销时间，相当于过期时间
		},
	}

	err = client.Agent().ServiceRegister(reg)

	if err != nil {
		log.Fatal("register server error : ", err)
	}

	ctx.JSON(http.StatusOK, "ok")
}
