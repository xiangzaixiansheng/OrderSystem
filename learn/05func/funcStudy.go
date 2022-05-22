package learn

import "fmt"

func funcStudy() {
	//2.切片
	s1 := []int{1, 2, 3, 4, 5}
	getSum(s1...)

	//匿名函数
	func() {
		fmt.Println("我是一个匿名函数。。")
	}()

	fun3 := func() {
		fmt.Println("我也是一个匿名函数。。")
	}
	fun3()

	//定义带参数的匿名函数
	func(a, b int) {
		fmt.Println(a, b)
	}(1, 2)

}

func getSum(nums ...int) int {
	//fmt.Printf("%T\n",nums) //[]int
	sum := 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
	}
	fmt.Println("总和是：", sum)
	return sum
}
