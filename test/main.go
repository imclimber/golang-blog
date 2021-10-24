
package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

const (
	MaxID = 1000
	MinID = 100
)

func main(){
	number := GenerateRandomNumber()
	fmt.Println("GenerateRandomNumber: ", number)
}

func GenerateRandomNumber() string {
	// 种子函数，基于不同的源创建随机值
	rand.Seed(time.Now().UnixNano())
	// 创建范围随机值
	number := rand.Intn(MaxID - MinID) + MinID
	return strconv.Itoa(number)
}