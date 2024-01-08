package main

import (
	"fmt"
	"github.com/FISCO-BCOS/go-sdk/abi/bind"
	"github.com/FISCO-BCOS/go-sdk/client"
	conf2 "github.com/FISCO-BCOS/go-sdk/conf"
	beego "github.com/beego/beego/v2/server/web"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"trace/conf"
	_ "trace/routers"
	"trace/service"
	"trace/service/Role"
)

func INITROLE() {
	fromAddress := conf.DistributorAddress
	privateKey := conf.DistributorPK
	key, err := crypto.HexToECDSA(privateKey)
	if err != nil {
		fmt.Println("error private:", err)
		return
	}
	auth := bind.NewKeyedTransactor(key)
	auth.From = common.HexToAddress(fromAddress)
	files, _ := conf.ParseConfigFile("config.toml")
	client, _ := client.Dial((*conf2.Config)(&files[0]))
	_, _, instance, _ := service.DeployService(client.GetTransactOpts(), client)
	session := service.ServiceSession{Contract: instance, CallOpts: *client.GetCallOpts(), TransactOpts: *client.GetTransactOpts()}
	address := common.HexToAddress(conf.DistributorAddress)
	//TODO: 设置Distributor角色

	address1 := common.HexToAddress(conf.ProducerAddress)
	//TODO: 设置Producer角色

	address2 := common.HexToAddress(conf.RetailerAddress)
	//TODO: 设置Retailer角色

	_, _, instance1, err := Role.DeployService(client.GetTransactOpts(), client)
	if err != nil {
		fmt.Println(err)
	}
	session1 := Role.ServiceSession{Contract: instance1, CallOpts: *client.GetCallOpts(), TransactOpts: *client.GetTransactOpts()}
	//TODO: 判断所有角色是否设置成功
	err1 := session1.OnlyDRRole(address)
	err1 = session1.OnlyPRRole(address1)
	err1 = session1.OnlyPRRole(address2)

	if err1 != nil || err2 != nil || err3 != nil {
		fmt.Println(err1)
		fmt.Println(err2)
		fmt.Println(err2)
	}
}
func main() {
	INITROLE()
	beego.BConfig.RouterCaseSensitive = false
	beego.Run()

}
