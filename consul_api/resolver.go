package consul_api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/hashicorp/consul/api"
	"log"
	"net"
	"strconv"
)

type consulWatcher struct {
	Client    *api.Client
	Server    string
	Address   map[string]struct{}
	lastIndex uint64
}

func Discovery(ctx *gin.Context) {
	var cw consulWatcher
	entries, meta, err := cw.Client.Health().Service(cw.Server, "", true, &api.QueryOptions{WaitIndex: cw.lastIndex})

	if err != nil {
		log.Fatalf("discovery error %s", err)
	}

	cw.lastIndex = meta.LastIndex

	addr := map[string]struct{}{}
	for _, s := range entries {
		addr[net.JoinHostPort(s.Service.Address, strconv.Itoa(s.Service.Port))] = struct{}{}
	}

	fmt.Println(entries)
}
