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


### 遇到的问题
1. 泛型细节问题
   current.Forward[i].GetKey() < key 语句，即使泛型都正确，但是依然显示current的key与K的类型不同？


