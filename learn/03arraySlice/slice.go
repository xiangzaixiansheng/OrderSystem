package learn

import "fmt"

/*
		数组array：
			存储一组相同数据类型的数据结构。
				特点：定长

		切片slice：
			同数组类似，也叫做变长数组或者动态数组。
				特点：变长

			是一个引用类型的容器，指向了一个底层数组。

		make()
			func make(t Type, size ...IntegerType) Type

			第一个参数：类型
				slice，map，chan
			第二个参数：长度len
				实际存储元素的数量
			第三个参数：容量cap
				最多能够存储的元素的数量

		append()，专门用于向切片的尾部追加元素
			slice = append(slice, elem1, elem2)
			slice = append(slice, anotherSlice...)

		切片Slice：
			1.每一个切片引用了一个底层数组
			2.切片本身不存储任何数据，都是这个底层数组存储，所以修改切片也就是修改这个数组中的数据
			3.当向切片中添加数据时，如果没有超过容量，直接添加，如果超过容量，自动扩容(成倍增长)
			4.切片一旦扩容，就是重新指向一个新的底层数组

		s1:3--->6--->12--->24

		s2:4--->8--->16--->32....

		slice := arr[start:end]
		 	切片中的数据：[start,end)
		 	arr[:end],从头到end
		 	arr[start:]从start到末尾

		 从已有的数组上，直接创建切片，该切片的底层数组就是当前的数组。
		 	长度是从start到end切割的数据量。
			但是容量从start到数组的末尾。

**/

func slice_test() {
	//1.数组
	arr := [4]int{1, 2, 3, 4} //定长

	s2 := []int{1, 2, 3, 4} //变长
	fmt.Println(s2)
	fmt.Printf("%T,%T\n", arr, s2) //[4]int,[]int

	s3 := make([]int, 3, 8)
	fmt.Println(s3)
	fmt.Printf("容量：%d,长度：%d\n", cap(s3), len(s3))
	s3[0] = 1
	s3[1] = 2
	s3[2] = 3
	fmt.Println(s3)

}
