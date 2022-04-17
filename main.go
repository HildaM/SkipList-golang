package main

import (
	"SkipList-golang/core"
	"fmt"
)

func main() {
	test01() // 基础功能
	// test02() // 10万数据插入测试
}

func test02() {
	skipList := new(core.SkipList)
	skipList.Init(20)

	skipList.LoadFile()

	skipList.TestInfo()
}

func test01() {
	skipList := new(core.SkipList)
	skipList.Init(6)

	// 插入数据
	skipList.InsertElement("1", "Hilda")
	skipList.InsertElement("2", "Quan")
	skipList.InsertElement("3", "Learning")
	skipList.InsertElement("4", "Golang")
	skipList.InsertElement("5", "deleteTest")

	skipList.DumpFile()

	fmt.Println("容量：", skipList.Size())
	skipList.DisplayList()

	skipList.SearchElement("1")
	skipList.SearchElement("4")

	skipList.DeleteElement("5")
	fmt.Println("容量：", skipList.Size())
	skipList.DisplayList()
}
