package main

import "fmt"

type Data struct {
	value int
	data  [200]int
}

// func ChangeData(arg Data) {
func ChangeData(arg *Data) {
	arg.value = 999
	arg.data[100] = 999
}

func main() {
	var data Data
	// ChangeData(data)
	ChangeData(&data)
	fmt.Println(data.value, data.data[100])

	// p1 only instance
	var p1 *Data = &Data{}
	var p2 *Data = p1
	var p3 *Data = p1

	// data1, data2, data3 all are instances
	var data1 Data
	var data2 Data = data1
	var data3 Data = data2

	var p4 = &Data{}
	var p5 = new(Data)
}
