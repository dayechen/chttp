package main

import (
	"fmt"
	"testing"
)

func Test_log(t *testing.T) {
	// defer_call()
	// pase_student()
}

// 草
func waitGroup() {
	var counter int

}

func defer_call() {
	defer func() {
		fmt.Println("打印前")
	}()
	defer func() {
		fmt.Println("打印中")
	}()
	defer func() {
		fmt.Println("打印后")
	}()
	// panic("触发异常")
}

type student struct {
	Name string
	Age  int
}

func pase_student() {
	m := make(map[string]*student)
	stu := []student{
		{Name: "一号", Age: 10},
		{Name: "二号", Age: 20},
		{Name: "三号", Age: 30},
	}
	for _, stu := range stu {
		m[stu.Name] = &stu
	}
	println(m)
}
