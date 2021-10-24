# 格式转换
常用格式转换包括以下几种：
- int32 类型转 string
- int64 类型转 string
- float32 类型转 string
- float64 类型转 string
- string 类型转 int32
- string 类型转 int64
- string 类型转 float32
- string 类型转 float64

## 案例
```go
package main

import (
	"log"
	"strconv"
)

const (
	BaseOfTen = 10
	BaseOfSixtyFour = 64
	BaseOfThirtyTwo = 32
	FloatFormatPrecision = -1
	FloatFormatFmt = 'E'
	Int32Num = 32
	Int64Num = 1000011110064
	Float32Num = 3.1415926535
	Float64Num = 3.14159265351415926535
	StringInt32 = "3200"
	StringInt64 = "100001111006400"
	StringFloat32 = "3.1415926535"
	StringFloat64 = "3.14159265351415926535"
)

func main(){
	ConvertFormat()
}

func ConvertFormat() {
	/* --- converting of int and string --- */
	// int32 to string
	res := strconv.Itoa(Int32Num)
	log.Println("int32 to string: ", res)
	// int64 to string (无符号整型 FormatUint(i uint64, base int))
	res = strconv.FormatInt(Int64Num, BaseOfTen)
	log.Println("int64 to string: ", res)

	// string to int32
	if res, err := strconv.Atoi(StringInt32); err != nil {
		log.Fatalln(res, " is not an integer.")
	} else {
		log.Printf("string to int32: %d of type %T\n", res, res)
	}
	// string to int64
	if res, err := strconv.ParseInt(StringInt64, BaseOfTen, BaseOfSixtyFour); err != nil {
		log.Fatalln(res, " is not an integer64.")
	} else {
		log.Printf("string to int32: %d of type %T\n", res, res)
	}

	/* --- converting of float and string --- */
	// float32 to string (仅baseSize不同)
	res = strconv.FormatFloat(Float32Num, FloatFormatFmt, FloatFormatPrecision, BaseOfThirtyTwo)
	log.Printf("float32 to string: %v\n", res)
	// float64 to string
	res = strconv.FormatFloat(Float32Num, FloatFormatFmt, FloatFormatPrecision, BaseOfSixtyFour)
	log.Printf("float64 to string: %v\n", res)

	// string to float32 (仅baseSize不同)
	if res, err := strconv.ParseFloat(StringFloat32,BaseOfThirtyTwo); err != nil {
		log.Fatalln(res, " is not an float32.")
	} else {
		log.Printf("string to float32: %v of type %T\n", res, res)
	}
	// string to float64
	if res, err := strconv.ParseFloat(StringFloat32,BaseOfSixtyFour); err != nil {
		log.Fatalln(res, " is not an float64.")
	} else {
		log.Printf("string to float64: %v of type %T\n", res, res)
	}
}
```

## float 转 string 的格式控制
```md
'b' (-ddddp±ddd，二进制指数)
'e' (-d.dddde±dd，十进制指数)
'E' (-d.ddddE±dd，十进制指数)
'f' (-ddd.dddd，没有指数)
'g' ('e':大指数，'f':其它情况)
'G' ('E':大指数，'f':其它情况)
```

## 参考
- [strconv](https://pkg.go.dev/strconv#FormatFloat)
- [Go语言string，int，int64 ,float之间类型转换方法](https://www.cnblogs.com/yaowen/p/8353444.html)