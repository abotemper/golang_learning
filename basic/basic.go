package main

import "fmt"

//函数外面不可以使用 bb := ....， 必须使用var

//集中定义
var (
	aa  = 3
	ss  = "kkk"
	asd = true
)

func variableZeroValue() {
	var a int
	var s string
	fmt.Println(a, s)
	//%q 会打印出引号
	fmt.Printf("%d %q\n", a, s)
}

func variableInitialValue() {
	//定义了变量必须用 不然报错
	var a, b int = 3, 4
	var s string = "abc"
	fmt.Println(a, s, b)
}

func variableTypeDeduction() {
	var a, b, c, s = 3, 4, true, "def"
	fmt.Println(a, b, c, s)
}

//下面是最好的形式，不想写var
func variableShorter() {
	a, b, c, s := 3, 4, true, "def"
	b = 5
	fmt.Println(a, b, c, s)
}

//bool string int int8 int16 in32 int64 uintptr byte , rune字符型 32位，不是char
// float32, float64, 复数complex64, complex128

func main() {
	fmt.Println("hello world")
	variableZeroValue()
	variableInitialValue()

}
