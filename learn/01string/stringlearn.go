package learn

import (
	"fmt"
	"strconv"
	"strings"
)

/*
	Go中的字符串是一个字节的切片。
			可以通过将其内容封装在“”中来创建字符串。Go中的字符串是Unicode兼容的，并且是UTF-8编码的。

**/
func main() {
	s1 := "orderSystem系统"
	fmt.Println(s1[2])
	fmt.Println(string(s1[2]))

	slice1 := []byte{65, 66, 67, 68, 69}
	s3 := string(slice1) //根据一个字节切片，构建字符串
	fmt.Println(s3)

	s4 := "abcdef"
	slice2 := []byte(s4) //根据字符串，获取对应的字节切片
	fmt.Println(slice2)

	/*
		strings包下的关于字符串的函数
	*/

	s5 := "helloworld"
	//1.是否包含指定的内容-->bool
	fmt.Println(strings.Contains(s5, "abc"))

	//3.统计substr在s中出现的次数
	fmt.Println(strings.Count(s5, "lloo"))

	//4.以xxx前缀开头，以xxx后缀结尾
	s6 := "20220501.txt"
	if strings.HasPrefix(s6, "20220501") {
		fmt.Println("2022年05月的文件。。")
	}
	if strings.HasSuffix(s6, ".txt") {
		fmt.Println("文本文档。。")
	}

	//索引
	//helloworld
	fmt.Println(strings.Index(s5, "llo"))
	//字符串的拼接
	ss1 := []string{"hello", "word", "study"}
	fmt.Println(strings.Join(ss1, "-"))

	//切割
	s7 := "123,4563,aaa,49595,45"
	ss2 := strings.Split(s7, ",")
	fmt.Println(ss2)

	/*
		截取子串：
		substring(start,end)-->substr
		str[start:end]-->substr
			包含start，不包含end下标
	*/

	s8 := "helloworld"
	fmt.Println(s8[:2]) //he

	/*
		strconv包：字符串和基本类型之前的转换
			string convert
	*/

	s9 := "true"
	b1, err := strconv.ParseBool(s9)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%T,%t\n", b1, b1)

	//2.整数
	s10 := "100"
	i2, err := strconv.ParseInt(s10, 2, 64)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%T,%d\n", i2, i2)

	ss3 := strconv.FormatInt(i2, 10)
	fmt.Printf("%T,%s\n", ss3, ss3)

	//itoa(),atoi()
	i3, err := strconv.Atoi("-42") //转为int类型
	fmt.Printf("%T,%d\n", i3, i3)
	ss4 := strconv.Itoa(-42)
	fmt.Printf("%T,%s\n", ss4, ss4)

}
