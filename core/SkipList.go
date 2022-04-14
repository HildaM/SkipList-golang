package core

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
		1. Init(): 初始化变量方法
		2. GetRandomLevel()：随机指定节点的索引层数
		3. InsertElement()：插入节点
		4. DisplayList()：查看所有节点
		5. SearchElement()：搜索指定节点
		6. DeleteElement()：删除指定节点
		7. DumpFile()：持久化数据
		8. LoadFile()：加载本地存储数据
		9. Size()：获取跳表长度

	私有方法：
		1. getKeyValueFromString()：从字符串中获取KV（？？？） // TODO
		2. isValidString()：检查字符串是否合法（？？？） // TODO

*/
