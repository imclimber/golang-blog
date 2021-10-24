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
	// float32 to string
	res = strconv.FormatFloat(Float32Num, FloatFormatFmt, FloatFormatPrecision, BaseOfThirtyTwo)
	log.Printf("float32 to string: %v\n", res)
	// float64 to string
	res = strconv.FormatFloat(Float32Num, FloatFormatFmt, FloatFormatPrecision, BaseOfSixtyFour)
	log.Printf("float64 to string: %v\n", res)

	// string to float32
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
