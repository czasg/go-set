# go-set

## 1.背景
在 go 中，不存在集合类型，这对数据处理造成了极大的不便。

go-set 定义了简单的集合标准，提供了部分集合的能力。

## 2.目标
- [x] String Set - 字符串集合
- [ ] Interface Set

## 3.使用
String Set 字符串集合
- `Add(keys ...string)`：新增集合元素
- `Del(keys ...string)`：删除集合元素
- `Has(keys ...string)`：集合包含全部元素
- `Any(keys ...string)`：集合存在任意元素
- `Len()`：集合长度
- `Keys()`：集合元素列表
- `Intersect(ss StringSet)`：交集
- `Union(ss StringSet)`：并集
- `Expect(ss StringSet)`：差集
- `Equals(ss StringSet)`：集合全等判断
- `Contains(ss StringSet)`：集合包含判断
