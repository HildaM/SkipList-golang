package core

/*
	Node:跳表的节点
*/

// 自定义类型
type K interface {
	string | int
}

// Node
type Node[K, V interface{}] struct {
	key   K
	value V

	Forward    []Node // 线性数组：保存不同层级的当前节点
	Node_level int    // 当前节点所在层数
}

/*
	对外提供的方法：
		0. Init()：初始化函数
		1. GetKey()：获取key
		2. GetValue(): 获取value
		3. SetValue(): 设置value
*/

// Init
func (node *Node[K, V]) Init(key K, value V, level int) {
	node.key = key
	node.value = value
	node.Node_level = level
	node.Forward = make([]Node, level+1, level+1)
}

// GetKey
func (node Node[K, V]) GetKey() K {
	return node.key
}

// GetValue
func (node Node[K, V]) GetValue() V {
	return node.value
}

// SetValue
func (node Node[K, V]) SetValue(value V) {
	node.value = value
}
