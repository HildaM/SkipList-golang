package core

import (
	"fmt"
	"reflect"
	"sync"
)

/*
	SkipList：跳表的具体实现
*/

// SkipList Define
type SkipList struct {
	// 跳表最大深度
	max_level int
	// 当前跳表的层数
	skip_list_level int
	// 头节点指针
	header *Node

	// 文件操作
	// TODO

	// 当前跳表的节点数
	element_count int
}

/*
	公共方法：
		Init(): 初始化变量方法
		GetRandomLevel()：随机指定节点的索引层数
		CreateNode()：创建节点
		InsertElement()：插入节点
		DisplayList()：查看所有节点
		SearchElement()：搜索指定节点
		DeleteElement()：删除指定节点
		DumpFile()：持久化数据
		LoadFile()：加载本地存储数据
		Size()：获取跳表长度

	私有方法：
		1. getKeyValueFromString()：从字符串中获取KV（？？？） // TODO
		2. isValidString()：检查字符串是否合法（？？？） // TODO

*/

// CreateNode
func (sklipList *SkipList) CreateNode(k string, v any, level int) *Node {
	node := new(Node)
	node.Init(k, v, level)
	return node
}

// InsertElement
// return 1: 元素存在		return 0: 元素插入成功
func (skipList *SkipList) InsertElement(key string, value any) int {
	maxLevel := skipList.max_level
	listLevel := skipList.skip_list_level

	// 上锁
	lock := sync.Mutex{}
	lock.Lock()

	//key = reflect.ValueOf(key)
	value = reflect.ValueOf(value)

	current := skipList.header

	/*
		更新数组update：
			倒序记录每层级需要修改的节点
			为插入节点和创建索引服务
	*/
	update := make([]*Node, maxLevel+1)

	/*
		1. 从头节点的最顶层开始遍历
		2. 一直到最接近key的最底层
	*/
	for i := listLevel; i >= 0; i-- {
		for current.Forward[i] != nil && current.Forward[i].GetKey() < key {
			current = current.Forward[i]
		}
		update[i] = current
	}

	// 回到底层
	current = current.Forward[0]

	// 如果存在key，则返回0
	if current != nil && current.GetKey() == key {
		fmt.Printf("key: %s alreadly exists!\n", key)
		lock.Unlock()
		return 1
	}

	/*
		插入操作：
			由上可知：如果cur为空，说明我们已经到达跳表底层
			我们需要在 update[0]和cur之前插入数值
	*/
	if current == nil || current.GetKey() != key {
		// 随机生成索引层数
		randomLevel := skipList.GetRandomLevel()

		// 如果生成的层数超过当前最高层数的话，那么剩下的空间就用header填充，并且更新当前层数
		if randomLevel > listLevel {
			for i := listLevel + 1; i < randomLevel+1; i++ {
				update[i] = skipList.header
			}
			// 这里需要更新原值
			skipList.skip_list_level = randomLevel
		}

		insertNode := skipList.CreateNode(key, value, randomLevel)

		// 插入元素
		for i := 0; i <= randomLevel; i++ {
			// 使用update数组实现逐层插入
			insertNode.Forward[i] = update[i].Forward[i]
			update[i].Forward[i] = insertNode
		}
		fmt.Printf("Successfully inserted key: %s", key)
		skipList.element_count++
	}

	lock.Unlock()
	return 0
}

// SearchElement
func (skipList *SkipList) SearchElement(key string) bool {
	fmt.Println("search element.....................")

	current := skipList.header

	// 从最高处开始寻找
	for i := skipList.skip_list_level; i >= 0; i-- {
		for current.Forward[i] != nil && current.Forward[i].GetKey() < key {
			current = current.Forward[i]
		}
	}

	// 根据上面找到的最近的元素，此刻可以从第0层开始寻找key
	current = current.Forward[0]

	if current != nil && current.GetKey() == key {
		fmt.Printf("Found key: %s, value: %d", current.GetKey(), current.GetValue())
		return true
	}

	fmt.Printf("Not found key: %s", key)
	return false
}

// GetRandomLevel
func (skipList *SkipList) GetRandomLevel() int {
	return 0
}
