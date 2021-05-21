package main

import (
	"fmt"
	"gin-demo/utils"
	"gin-demo/utils/setting"
)

func main() {
	setting.Setup()
	fmt.Println(utils.EncodeMD5(setting.AppSetting.SecretSalt+"123456"))
}
