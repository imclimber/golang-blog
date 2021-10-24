# 生成随机数函数
生成范围随机值是非常常用的功能，本文选取两种生成随机方案，第一是 `math/rand`，第二种是 `crypto/rand`。

## math/rand 伪随机数生成器
- 单独使用可能有问题，因为重新启动程序生成随机数不变。如果每次运行需要不同的结果，请使用种子函数初始化默认的源。
- 一般使用时都是获取随机范围值。
- 此包并不严谨，某些 golang check 规则可能通不过，请使用 `crypto/rand` 加密安全的随机数生成器。

```go
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
```

## crypto/rand 加密安全的随机数生成器

```go
package main

import (
	"crypto/rand"
	"log"
	"math/big"
	"strconv"
)

const (
	MaxID = 1000
	MinID = 100
	BaseOfTen = 10
)

func main(){
	number := GenerateRandomNumber()
	log.Println("GenerateRandomNumber: ", number)
}

func GenerateRandomNumber() string{
	number, err := rand.Int(rand.Reader, big.NewInt(MaxID-MinID))
	if err != nil {
		log.Fatalf("Generate new rand number error: %v", err)
	}
	needNum := number.Int64() + MinID
	return  strconv.FormatInt(needNum, BaseOfTen)
}
```

## 参考
- [Golang: math/rand 和 crypto/rand 区别](https://segmentfault.com/a/1190000021543499)
- [golang 随机数 math/rand包 crypto/rand包](https://blog.csdn.net/whatday/article/details/106617208)