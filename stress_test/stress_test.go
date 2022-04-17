package stress_test

import (
	"SkipList-golang/core"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"sync"
	"testing"
	"time"
)

// 压力测试
/*
	golang实现并发
*/

var (
	NUM_THREAD int = 1
	TEST_COUNT int = 10000
	skipList       = new(core.SkipList)

	wg sync.WaitGroup
)

func insertElement() {
	rand.Seed(time.Now().UnixNano())

	tid := os.Getpid()
	fmt.Println(tid)

	tmp := TEST_COUNT / NUM_THREAD
	count := 0
	for i := tid * tmp; count < tmp; i++ {
		count++
		skipList.InsertElement(strconv.Itoa(rand.Int()%TEST_COUNT), "a")
	}

	wg.Done()
}

func getElement() {
	rand.Seed(time.Now().UnixNano())

	tid := os.Getpid()
	fmt.Println(tid)

	tmp := TEST_COUNT / NUM_THREAD
	count := 0
	for i := tid * tmp; count < tmp; i++ {
		count++
		skipList.SearchElement(strconv.Itoa(rand.Int() % TEST_COUNT))
	}

	wg.Done()
}

func storeFile() {
	skipList.DumpFile()
}

func TestStress(t *testing.T) {
	wg.Add(2)
	startTime := time.Now()

	skipList.Init(18)

	go insertElement()
	// time.Sleep(1000 * time.Millisecond)
	go getElement()

	wg.Wait()

	endTime := time.Now()

	// time.Sleep(100 * time.Millisecond)
	// storeFile()
	fmt.Printf("开始时间：%v；结束时间：%v\n", startTime, endTime)
	skipList.TestInfo()

}
