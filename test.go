package main

import (
	"fmt"
	"gin-demo/app/models"
	"gin-demo/app/services"
	"gin-demo/utils/setting"
)

func main() {
	setting.Setup()
	models.Setup()
	data,_:= services.GetAllWhitelistUserCode(map[string]interface{}{},"code asc",-1)
	fmt.Println(data)
}
