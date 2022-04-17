# 基于SkipList的KV存储引擎

### 简介
本项目基于 [SkipList-CPP](https://github.com/youngyangyang04/Skiplist-CPP) ，使用 Golang 重写。是学习golang的练手项目。

SkipList-golang是基于跳表实现的轻量级键值型存储引擎，使用golang重构。
提供：插入数据、删除数据、查询数据、数据展示、数据落盘、文件加载数据，以及数据库大小显示

### 提供接口
- InsertElement（插入数据）
- DeleteElement（删除数据）
- SearchElement（查询数据）
- DisplayList（展示已存数据）
- DumpFile（数据落盘）
- LoadFile（加载数据）
- Size（返回数据规模）

### 压力测试
执行stress_test文件夹下的测试文件即可

#### 表现
在插入1万次和搜索1万次并发下：
```
开始时间：2022-04-17 17:04:33.2174831 +0800 CST m=+0.011774401；结束        时间：2022-04-17 17:04:40.227864 +0800 CST m=+7.022155301
生成 3682 元素，索引层数为 11，最大的索引层数为 18
--- PASS: TestStress (7.09s)
```

### 不足
- 我对泛型的理解还不是很深，所以key只能用string表示。在之后会对项目再进行修改


### 遇到的问题
1. 泛型细节问题
   current.Forward[i].GetKey() < key 语句，即使泛型都正确，但是依然显示current的key与K的类型不同？


