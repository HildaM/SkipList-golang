package core

/*
	Node:跳表的节点
*/

// Node
type Node struct {
	key   string
	value any

	/*
		线性指针数组
			1. 保存不同层级的当前元素
			2. i: 表示层数  Forward[i]：表示当前层数下，该节点的下一个节点的指针值
	*/
	Forward []*Node

	Node_level int // 当前节点所在层数
}

/*
	对外提供的方法：
		0. Init()：初始化函数
		1. GetKey()：获取key
		2. GetValue(): 获取value
		3. SetValue(): 设置value
*/

// Init
func (node *Node) Init(key string, value any, level int) {
	node.key = key
	node.value = value
	node.Node_level = level
	node.Forward = make([]*Node, level+1, level+1)
}

// GetKey
func (node Node) GetKey() string {
	return node.key
}

// GetValue
func (node Node) GetValue() any {
	return node.value
}

// SetValue
func (node Node) SetValue(value any) {
	node.value = value
}
