package learn

import "fmt"

func array() {
	/*
		var 数组名 [长度] 数据类型
		var 数组名 = [长度] 数据类型{元素1，元素2.。。}
		数组名 := [...]数据类型{元素。。。}

		一维数组：存储的多个数据是数值本身
			a1 :=[3]int{1,2,3}

		二维数组：存储的是一维的一维
			a2 := [3][4]int{{},{},{}}

			该二维数组的长度，就是3。
			存储的元素是一维数组，一维数组的元素是数值，每个一维数组长度为4。

	**/
	//step1：创建数组
	var arr1 [4]int
	fmt.Printf("%p\n", &arr1)
	//step2：数组的访问
	arr1[0] = 1
	arr1[1] = 2
	arr1[2] = 3

	fmt.Println("数组的长度：", len(arr1)) //容器中实际存储的数据量
	fmt.Println("数组的容量：", cap(arr1)) //容器中能够存储的最大的数量

	var b = [4]int{1, 2, 3, 4}
	fmt.Println(b)

	f := [...]int{1, 2, 3, 4, 5}
	fmt.Println(f)
	fmt.Println(len(f))

}
