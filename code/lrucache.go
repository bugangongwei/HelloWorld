package code

import (
	"container/list"
	"fmt"
	"sync"
	"time"
)

/*
# 实现支持过期时间的LRU缓存（考察数据结构设计+边界处理）
class TimedLRUCache:
    def __init__(self, capacity: int, ttl: int):
        # 要求实现 get(key) 和 put(key, value)

# 附加要求：
# 1. 缓存命中时刷新过期时间
# 2. 后台线程自动清理过期项
*/

// 缓存项结构
type cacheItem struct {
	key      string // TODO: 为什么必须保留key
	value    interface{}
	expireAt time.Time // 过期时间戳，单位为秒
}

// LRU缓存结构
type TimedLRUCache struct {
	// 存储机制
	capacity int                      // 最大容量
	list     *list.List               // 双向链表，用于维护LRU顺序
	items    map[string]*list.Element // 哈希表，用于快速查找缓存项
	// 过期机制
	ttl           time.Duration // 缓存项的过期时间
	cleanupTicker *time.Ticker  // 定时清理过期项的定时器
	cleanupStop   chan struct{} // 用于停止清理后台线程
	// 读写锁保护map
	mu sync.RWMutex // 读写锁，保护map的并发访问
}

// NewTimedLRUCache 创建一个新的TimedLRUCache实例
func NewTimedLRUCache(capacity int, ttl time.Duration) *TimedLRUCache {
	cache := &TimedLRUCache{
		capacity:    capacity,
		list:        list.New(),
		items:       make(map[string]*list.Element),
		ttl:         ttl,
		cleanupStop: make(chan struct{}),
	}
	cache.startCleanup() // 启动清理后台线程
	return cache
}

func (c *TimedLRUCache) startCleanup() {
	c.cleanupTicker = time.NewTicker(c.ttl / 2) // 每个TTL的一半进行清理
	go func() {
		for {
			select {
			case <-c.cleanupTicker.C:
				c.cleanupExpiredItems()
			case <-c.cleanupStop:
				c.cleanupTicker.Stop()
				return
			}
		}
	}()
}

func (c *TimedLRUCache) cleanupExpiredItems() {
	// 清理过期项
	now := time.Now()

	c.mu.Lock()
	defer c.mu.Unlock()

	for k, v := range c.items {
		item := v.Value.(*cacheItem)
		if item.expireAt.Before(now) {
			fmt.Println("remove key:", k)
			c.list.Remove(v)
			delete(c.items, k)
		}
	}
}

// Get 获取缓存项
func (c *TimedLRUCache) Get(key string) (interface{}, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	if elem, found := c.items[key]; found {
		fmt.Println("get key:", key)
		item := elem.Value.(*cacheItem)
		fmt.Println("ttl:", c.ttl, "expireAt:", item.expireAt, "now:", time.Now())
		if item.expireAt.After(time.Now()) { // 检查是否过期
			fmt.Println("hit key:", key, "value:", item.value, "expireAt:", item.expireAt)
			// 刷新过期时间
			item.expireAt = time.Now().Add(c.ttl)
			c.list.MoveToFront(elem) // 将该项移动到链表头部
			return item.value, true
		}
		fmt.Println("miss key:", key, "expired at:", item.expireAt)
		c.list.Remove(elem) // 如果过期，删除该项
		delete(c.items, key)
	}
	return nil, false
}

func (c *TimedLRUCache) Put(key string, value interface{}) {
	c.mu.Lock()
	defer c.mu.Unlock()

	if elem, found := c.items[key]; found {
		// 如果已存在，更新值和过期时间
		item := elem.Value.(*cacheItem)
		item.value = value
		item.expireAt = time.Now().Add(c.ttl)
		c.list.MoveToFront(elem) // 将该项移动到链表头部
		return
	}

	// 如果不存在，创建新的缓存项
	if c.list.Len() >= c.capacity {
		// 如果达到容量限制，删除最旧的项
		oldest := c.list.Back()
		if oldest != nil {
			c.list.Remove(oldest)
			delete(c.items, oldest.Value.(*cacheItem).key)
		}
	}

	newItem := &cacheItem{
		key:      key,
		value:    value,
		expireAt: time.Now().Add(c.ttl*time.Second), // 设置过期时间
	}
	elem := c.list.PushFront(newItem)
	c.items[key] = elem // 添加到哈希表中
	fmt.Println("put key:", key, "value:", value, "expireAt:", newItem.expireAt)
}
