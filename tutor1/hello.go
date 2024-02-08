package main

import (
	"fmt"
)

var x int = 5
var y int

func Inti() {
	fmt.Println("x =", x)
}
func main() {

	var t bool
	var f bool = true

	fmt.Println("t is", t)
	fmt.Println("f is", f)

	var age int = 21
	var favNum float64 = 1.6180339

	var complex complex128 = 5 + 5i

	var r rune = 10

	fmt.Println("age is", age, "favNum is", favNum)
	fmt.Println("complex128 is", complex)
	fmt.Println("rune is", r)

	var myName string = "Rocky"
	fmt.Println(myName + "is a robot")
	fmt.Println(myName[3])
	fmt.Println("length of myName is", len(myName))

	var arry5 [5]float64 //array of 5 float64
	arry5[0] = 98.7
	arry5[1] = 93.2
	arry5[2] = 77.2
	arry5[3] = 83.7
	arry5[4] = 88.2

	fmt.Println(arry5)
	fmt.Println("length of arry is", len(arry5))
	fmt.Println("arry[3] is ", arry5[3])

	arry3 := [3]float64{98, 93, 77}
	fmt.Println(arry3)

	var s []float64 = arry5[0:2] //slce of an array
	fmt.Println(s)

	type Person struct {
		name string
		age  int
	}

	var p Person
	p.name = "Rocky"
	p.age = 40
	fmt.Println(p)

	//pointer

	var x int = 5
	var xPtr *int = &x
	fmt.Println(xPtr) //print the memory address of x

	//Map types
	// ประกาศและสร้าง map เก็บข้อมูลเกี่ยวกับอายุ ครอบครัวของ กัญญารัตน์
	person := map[string]int{
		"Rocky": 21,
		"Mymom":    50,
		"Mydad":    57,
	}

	// แสดงผล map
	fmt.Println("Map อายุของคนแต่ละคน:", person)

	// เข้าถึงค่าใน map
	ageKanyarat := person["Kanyarat"]
	fmt.Println("อายุของ Kanyarat:", ageKanyarat)

	// เพิ่มคู่ key-value ใน map
	person["Sister"] = 28
	fmt.Println("Map หลังเพิ่ม Sister:", person)

	// ลบคู่ key-value จาก map
	delete(person, "Mydad")
	fmt.Println("Map หลังลบ Mydad:", person)

	// ตรวจสอบว่า key มีใน map หรือไม่
	_, isMydadPresent := person["Mydad"]
	fmt.Println("Mydad มีใน map หรือไม่:", isMydadPresent)

	// วนลูปทุกรายการใน map
	fmt.Println("รายชื่อและอายุของทุกคน:")
	for name, age := range person {
		fmt.Println(name, ":", age)
	}
}