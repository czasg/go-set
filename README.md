# go-set

## 1.背景
go-set 定义了简单的集合标准，提供了部分集合的能力。

## 2.支持
- [x] String Set - 字符串集合

## 3.使用
String Set 字符串集合
- `Add(keys ...string)`：新增集合元素
- `Del(keys ...string)`：删除集合元素
- `HasAll(keys ...string)`：集合包含全部元素
- `HasAny(keys ...string)`：集合存在任意元素
- `Len()`：集合长度
- `Keys()`：集合元素列表
- `Intersect(ss StringSet)`：交集
- `Union(ss StringSet)`：并集
- `Expect(ss StringSet)`：差集
- `Equals(ss StringSet)`：集合全等判断
- `Contains(ss StringSet)`：集合包含判断
