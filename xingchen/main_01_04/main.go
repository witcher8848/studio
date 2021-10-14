package main

import (
	"fmt"
	"regexp"
)

func reUse(str string,re string)  string{
	//解析正则表达式，如果成功返回解释器
	reg1 := regexp.MustCompile(re)
	if reg1 == nil {
		fmt.Println("regexp err")
		return ""
	}
	//根据规则提取关键信息
	result1 := reg1.FindString(str)
	//fmt.Println("result1 = ", result1)
	return result1
}
func printResult(str string)  {
	if str!=""{
		fmt.Println("True",str[:4],str[4:6],str[6:8],str[8:10],str[10:])
	}else {
		fmt.Println("False")
	}
}

func judgeString(str string)  {
	reUndergraduate:=`^\d{11}ys|\d{13}$`
	reGraduate:=`^\d{12}$`
	switch len(str) {
	case 12:
		str1:=reUse(str,reGraduate)
		printResult(str1)
	case 13:
		str1:=reUse(str,reUndergraduate)
		printResult(str1)
	default:
		fmt.Println("False")
	}
}
func main()  {
	count:=0
	fmt.Scanln(&count)
	strArr:=make([]string,count)
	for i := 0; i < count; i++ {
		fmt.Scanln(&strArr[i])
	}
	for i := 0; i < count; i++ {
		judgeString(strArr[i])
	}
	//s:="2021020202002"
	//buf := "2021020202002m"
	//buf1 := "20210202020ys"
}
