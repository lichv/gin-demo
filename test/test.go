package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

func main() {
	//setting.Setup()
	//fmt.Println(time.Now().UnixNano())
	now := time.Now().UnixNano()/1000
	fmt.Println(now)
	code := "wx_"+strconv.FormatUint(uint64(now),36)+strconv.Itoa(rand.New(rand.NewSource(now)).Intn(90)+10)
	fmt.Println(strconv.FormatUint(uint64(now),36))
	fmt.Println(code)
	//fmt.Println(utils.EncodeMD5(setting.AppSetting.SecretSalt+"123456"))
}
