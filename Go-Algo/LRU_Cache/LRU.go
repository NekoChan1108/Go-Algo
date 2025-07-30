// 自己写的LRU缓存
package main

type LRUCache struct {
	CacheMap   map[int]int
	Capacity   int
	CacheOrder []int
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		CacheMap:   make(map[int]int),
		Capacity:   capacity,
		CacheOrder: make([]int, 0),
	}
}

func (lru *LRUCache) Index(key int) int {
	for idx, v := range lru.CacheOrder {
		if v == key {
			return idx
		}
	}
	return -1
}

func (lru *LRUCache) Get(key int) int {
	if val, ok := lru.CacheMap[key]; ok {
		//TODO 移动到最后面
		idx := lru.Index(key)
		if idx != len(lru.CacheOrder)-1 && idx != -1 {
			lru.CacheOrder = append(append(lru.CacheOrder[:idx], lru.CacheOrder[idx+1:]...), key)
		}
		return val
	}
	return -1
}

func (lru *LRUCache) Put(key int, value int) {
	len := len(lru.CacheOrder)
	if _, ok := lru.CacheMap[key]; ok {
		lru.CacheMap[key] = value
		//移动
		idx := lru.Index(key)
		if idx != len-1 {
			lru.CacheOrder = append(append(lru.CacheOrder[:idx], lru.CacheOrder[idx+1:]...), key)
		}
		return
	}
	if len == lru.Capacity {
		delete(lru.CacheMap, lru.CacheOrder[0])
		lru.CacheOrder = lru.CacheOrder[1:]
	}
	lru.CacheMap[key] = value
	lru.CacheOrder = append(lru.CacheOrder, key)
}

