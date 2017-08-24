/*
 * Revision History:
 *     Initial: 2017/08/24        Yang Zhengtian
 */
package main

import "log"

type student struct {
	Name string
	Age  int
}

func pase_student() {
	m := make(map[string]*student)

	stus := []student{
		{Name: "zhou", Age: 24},
		{Name: "li", Age: 23},
		{Name: "wang", Age: 22},
	}

	for _, stu := range stus {
		log.Println(&stu)
		m[stu.Name] = &stu
		log.Println(m[stu.Name])
		log.Println(m)
	}
	log.Println(m)
	log.Println(m["li"], m["zhou"], m["wang"])
}

func main() {
	pase_student()
}
