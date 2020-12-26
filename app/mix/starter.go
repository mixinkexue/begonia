package mix

import (
	"context"
	"fmt"
	"github.com/MashiroC/begonia/app/client"
	"github.com/MashiroC/begonia/dispatch"
	"github.com/MashiroC/begonia/internal"
	"github.com/MashiroC/begonia/logic"
	"log"
)

// BootStartByCenter 根据center cluster模式启动
func BootStart(optionMap map[string]interface{}) *MixNode {

	fmt.Println("  ____                              _        \n |  _ \\                            (_)       \n | |_) |  ___   __ _   ___   _ __   _   __ _ \n |  _ <  / _ \\ / _` | / _ \\ | '_ \\ | | / _` |\n | |_) ||  __/| (_| || (_) || | | || || (_| |\n |____/  \\___| \\__, | \\___/ |_| |_||_| \\__,_|\n                __/ |                        \n               |___/                         ")

	log.Printf("begonia client start with [%s] mode\n", internal.ServiceAppMode)

	ctx, cancel := context.WithCancel(context.Background())
	c := &rClient{
		ctx:    ctx,
		cancel: cancel,
	}

	// TODO:给dispatch初始化

	var addr string
	if addrIn, ok := optionMap["addr"]; ok {
		addr = addrIn.(string)
	} else {
		panic("addr must exist")
	}

	log.Printf("begonia client will link to [%s]", addr)

	var dp dispatch.Dispatcher
	dp = dispatch.NewLinkedByDefaultCluster()

	if err := dp.Link(addr); err != nil {
		panic(err)
	}

	lg:=logic.NewService(dp,logic.NewWaitChans())
	cli:=client.NewClient(&lg.Client)

	dp.Handle("frame", c.lg.DpHandler)

	//TODO: 发一个包，拉取配置

	/*

		先不去拉配置 后面再加

		// 假设这个getConfig是sub service的一个远程函数
		var getConfig = func(...interface{}) (interface{}, error) {
			return map[string]interface{}{}, nil
		}

		// 假设m就是拿到的远程配置
		m, err := getConfig()

		// TODO:根据拿到的远程配置来修改配置
		// do some thing
		// 修改配置之前的一系列调用全部都是按默认配置来的
	*/

	return c
}