package main

import (
	"fmt"
	"os"

	"OrderSystem/learn/12protobuffer/ProtocDemo/example"

	"github.com/golang/protobuf/proto"
)

func main() {
	fmt.Println("Hello World. \n")

	fmt.Println(float64(3.4132323232322323232))

	msg_test := &example.Person{
		Name: proto.String("Davie"),
		Age:  proto.Int(18),
		From: proto.String("China"),
	}

	//序列化
	msgDataEncoding, err := proto.Marshal(msg_test)
	if err != nil {
		panic(err.Error())
		return
	}

	//反序列化：
	msgEntity := example.Person{}
	err = proto.Unmarshal(msgDataEncoding, &msgEntity)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
		return
	}
	fmt.Printf("姓名：%s\n\n", msgEntity.GetName())
	fmt.Printf("年龄：%d\n\n", msgEntity.GetAge())
	fmt.Printf("国籍：%s\n\n", msgEntity.GetFrom())

	//=============================
	msg_test2 := &example.Order{
		OrderId:   proto.Int(12),
		Timestamp: proto.String("345"),
		Basce:     proto.String("beijing"),
	}

	//序列化
	msgDataEncoding2, err := proto.Marshal(msg_test2)
	if err != nil {
		panic(err.Error())
		return
	}

	//反序列化：
	msgEntity2 := example.Order{}
	err = proto.Unmarshal(msgDataEncoding2, &msgEntity2)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
		return
	}
	fmt.Printf("orderid%d\n\n", msgEntity2.GetOrderId())
	fmt.Printf("Basce%s\n\n", msgEntity2.GetBasce())

}
