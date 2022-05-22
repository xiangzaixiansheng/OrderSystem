package learn

import "fmt"

/*
	map：映射，是一种专门用于存储键值对的集合。属于引用类型

	存储特点：
		A：存储的是无序的键值对
		B：键不能重复，并且和value值一一对应的。
				map中的key不能重复，如果重复，那么新的value会覆盖原来的，程序不会报错。

	语法结构：
		1.创建map
			var map1 map[key类型]value类型
				nil map，无法直接使用

			var map2 = make(map[key类型])value类型

			var map3 = map[key类型]value类型{key:value,key:value,key:value...}

		2.添加/修改
			map[key]=value
				如果key不存在，就是添加数据
				如果key存在，就是修改数据

		3.获取
			map[key]-->value

			value,ok := map[key]
				根据key获取对应的value
					如果key存在，value就是对应的数据，ok为true
					如果key不存在，value就是值类型的默认值，ok为false

		4.删除数据：
			delete(map，key)
				如果key存在，就可以直接删除
				如果key不存在，删除失败

		5.长度：
			len()
**/
func mapStudy() {
	map1 := make(map[string]string)
	map1["hello"] = "world"
	map1["hello1"] = "world1"
	map1["hello2"] = "world2"

	v1, ok := map1["hello"]
	if ok {
		fmt.Println("对应的数值是：", v1)
	} else {
		fmt.Println("操作的key不存在，获取到的是零值：", v1)
	}
	delete(map1, "hello")
	fmt.Println(map1)

	//1.遍历map
	for k, v := range map1 {
		fmt.Println(k, v)
	}

	//1.创建map存储第一个人的信息
	map2 := make(map[string]string)
	map2["name"] = "王二狗"
	map2["age"] = "30"
	map2["sex"] = "男性"
	map2["address"] = "北京市XX路XX号"
	fmt.Println(map2)

	//2.第二个人
	map3 := make(map[string]string)
	map3["name"] = "李小花"
	map3["age"] = "20"
	map3["sex"] = "女性"
	map3["address"] = "上海市。。。"
	fmt.Println(map3)

	//3.
	map4 := map[string]string{"name": "ruby", "age": "30", "sex": "女性", "address": "杭州市"}
	fmt.Println(map4)

	//将map存入到slice中
	s1 := make([]map[string]string, 0, 3)
	s1 = append(s1, map1)
	s1 = append(s1, map2)
	s1 = append(s1, map3)

	//遍历切片
	for i, val := range s1 {
		//val ：map1，map2，map3
		fmt.Printf("第 %d 个人的信息是：\n", i+1)
		fmt.Printf("\t姓名：%s\n", val["name"])
		fmt.Printf("\t年龄：%s\n", val["age"])
		fmt.Printf("\t性别：%s\n", val["sex"])
		fmt.Printf("\t地址：%s\n", val["address"])
	}

}
