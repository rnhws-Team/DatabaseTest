package main

import (
	"fmt"
)

func main() {
	new := DataBase{}
	dataVolume := 10000 // dataVolume 描述测试的数据量

	err := new.OpenDatabase("/storage/emulated/0/Download/test")
	if err != nil {
		panic(err)
	}
	new.WriteTestData(dataVolume)
	err = new.CloseDatabase()
	if err != nil {
		panic(err)
	}

	err = new.OpenDatabase("/storage/emulated/0/Download/test")
	if err != nil {
		panic(err)
	}
	new.GetTestData()
	fmt.Println(new.DataReceived)
	new.WriteTestData(dataVolume)
	new.GetTestData()
	fmt.Println(new.DataReceived)
	new.CloseDatabase()

	err = new.OpenDatabase("/storage/emulated/0/Download/test")
	if err != nil {
		panic(err)
	}
	new.AppendTestData(dataVolume)
	new.GetTestData()
	fmt.Println(new.DataReceived)
	new.CloseDatabase()
}
