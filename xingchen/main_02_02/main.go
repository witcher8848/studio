package main

import (
	"fmt"
	"net/http"
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
func GetResult(str string)  string{
	var strRes string
	if str!=""{
		strRes=fmt.Sprintf("%s %s %s %s %s %s","True",str[:4],str[4:6],str[6:8],str[8:10],str[10:])
	}else {
		strRes="False"
	}
	return strRes
}

func judgeString(str string) string {
	var strRes string
	reUndergraduate:=`^\d{11}ys|\d{13}$`
	reGraduate:=`^\d{12}$`
	switch len(str) {
	case 12:
		str1:=reUse(str,reGraduate)
		strRes=GetResult(str1)
	case 13:
		str1:=reUse(str,reUndergraduate)
		strRes=GetResult(str1)
	default:
		strRes="False"
	}
	return strRes
}
func handler(w http.ResponseWriter,r *http.Request)  {
	w.WriteHeader(200)
	var strRes string
	reUndergraduate:=`^\d{11}ys|\d{13}$`
	reGraduate:=`^\d{12}$`
	str:=r.FormValue("num")
	switch len(str) {
	case 12:
		str1:=reUse(str,reGraduate)
		strRes=judgeString(str1)
	case 13:
		str1:=reUse(str,reUndergraduate)
		strRes=judgeString(str1)
	default:
		strRes="False"
	}
	w.Write([]byte(strRes))
}

func main()  {
	http.HandleFunc("/num",handler)
	http.ListenAndServe(":8080",nil)
}