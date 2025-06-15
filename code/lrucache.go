package code

import (
	"container/heap"
	"container/list"
	"fmt"
	"sync"
	"sync/atomic"
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
	key      string // 删除最后一个元素的时候需要获取key
	value    interface{}
	expireAt time.Time // 过期时间戳，单位为秒
}

// 带过期时间的LRU缓存结构
type TimedLRUCache struct {
	// 存储机制
	capacity int                      // 最大容量
	list     *list.List               // 双向链表，用于维护LRU顺序
	items    map[string]*list.Element // 哈希表，用于快速查找缓存项
	// 过期机制
	ttl           time.Duration // 缓存项的过期时间
	cleanupTicker *time.Ticker  // 定时清理过期项的定时器
	cleanupStop   chan struct{} // 用于停止清理后台线程
	// 读写锁保护map读写
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

	// 全表扫描，扫描期间不允许map的写操作
	// 对于多读少写的场景，使用读写锁可以提高并发性能
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
		expireAt: time.Now().Add(c.ttl * time.Second), // 设置过期时间
	}
	elem := c.list.PushFront(newItem)
	c.items[key] = elem // 添加到哈希表中
	fmt.Println("put key:", key, "value:", value, "expireAt:", newItem.expireAt)
}

/* ========== 高并发读写的过期lru缓存 ========== */
/* ========== 高并发读写的过期lru缓存 ========== */
/* ========== 高并发读写的过期lru缓存 ========== */

type cacheItemV2 struct {
	key      string
	value    interface{}
	expireAt time.Time // 过期时间戳，单位为秒

	elem    *list.Element // TODO(suping.liu)
	heapIdx int           // 堆索引，用于快速查找过期项
}

// 最小堆实现
type expirationHeap []*cacheItemV2

func (h expirationHeap) Len() int {
	return len(h)
}
func (h expirationHeap) Less(i, j int) bool {
	return h[i].expireAt.Before(h[j].expireAt)
}
func (h expirationHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
	h[i].heapIdx = i
	h[j].heapIdx = j
}
func (h *expirationHeap) Push(x interface{}) {
	item := x.(*cacheItemV2)
	item.heapIdx = len(*h)
	*h = append(*h, item)
}
func (h *expirationHeap) Pop() interface{} {
	if len(*h) == 0 {
		return nil
	}
	item := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	item.heapIdx = -1 // 标记为无效
	return item
}

type cacheShard struct {
	capacity int
	mu       sync.RWMutex            // 读写锁，保护map的并发访问
	items    map[string]*cacheItemV2 // 哈希表，用于快速查找缓存项
	list     *list.List              // 双向链表，用于维护LRU顺序
	// 过期机制
	ttl            time.Duration  // 缓存项的过期时间
	expirationHeap expirationHeap // 最小堆，用于快速查找过期项
	// 统计信息
	stats struct { // TODO(suping.liu)
		hits   uint64 // 命中次数
		misses uint64 // 未命中次数
		evicts uint64 // 淘汰次数
	}
}

type TimedLRUCacheV2 struct {
	shards     []*cacheShard // 分片数组
	shardCount int           // 分片数量
	// 过期机制
	ttl           time.Duration // 缓存项的过期时间
	cleanupTicker *time.Ticker  // 定时清理过期项的定时器
	cleanupStop   chan struct{} // 用于停止清理后台线程
	// 统计信息
	statsTicker *time.Ticker  // 定时打印统计信息
	statsStop   chan struct{} // 用于停止统计信息打印的定时器
}

func NewTimedLRUCacheV2(shardCount int, capacity int, ttl time.Duration, cleanupInterval, reportInterval time.Duration) *TimedLRUCacheV2 {
	if shardCount <= 0 {
		shardCount = 1 // 至少一个分片
	}

	cache := &TimedLRUCacheV2{
		shards:      make([]*cacheShard, shardCount),
		shardCount:  shardCount,
		ttl:         ttl,
		cleanupStop: make(chan struct{}),
		statsStop:   make(chan struct{}),
	}

	shardCap := capacity / shardCount
	if shardCap <= 0 {
		shardCap = 1 // 至少一个容量
	}

	for i := 0; i < shardCount; i++ {
		cache.shards[i] = &cacheShard{
			capacity:       shardCap,
			items:          make(map[string]*cacheItemV2),
			list:           list.New(),
			ttl:            ttl,
			expirationHeap: make(expirationHeap, 0, shardCap),
		}
	}

	// 后台清理
	cache.cleanupTicker = time.NewTicker(cleanupInterval * time.Second)
	go cache.cleanupV2()

	// 上报统计信息
	cache.statsTicker = time.NewTicker(reportInterval * time.Second)
	go cache.reportStatsPeriod()

	return cache
}

func fnv32(key string) uint32 {
	var hash uint32 = 2166136261
	for _, b := range key {
		hash ^= uint32(b)
		hash *= 16777619
	}
	return hash
}

func (c *TimedLRUCacheV2) getShard(key string) *cacheShard {
	hash := fnv32(key)
	return c.shards[hash%uint32(c.shardCount)]
}

func (c *TimedLRUCacheV2) reportStatsPeriod() {
	for {
		select {
		case <-c.statsTicker.C:
			fmt.Println("report stats...")
			c.reportStats()
		case <-c.statsStop:
			c.statsTicker.Stop()
			return
		}
	}
}

func (c *TimedLRUCacheV2) reportStats() {
	var totalItems, totalHits, totalMisses, totalEvicts uint64
	for _, shard := range c.shards {
		shard.mu.RLock()
		totalItems += uint64(len(shard.items))
		totalHits += atomic.LoadUint64(&shard.stats.hits)
		totalMisses += atomic.LoadUint64(&shard.stats.misses)
		totalEvicts += atomic.LoadUint64(&shard.stats.evicts)
		shard.mu.RUnlock()
	}
	totalRequest := totalHits + totalMisses
	hitRate := 0.0
	if totalRequest > 0 {
		hitRate = float64(totalHits) / float64(totalRequest) * 100.0
	}
	fmt.Printf("totalitems=%d, hits=%d, misses=%d, evicts=%d, hitrate=%f\n", totalItems, totalHits, totalMisses, totalEvicts, hitRate)
}

func (c *TimedLRUCacheV2) Close() {
	close(c.cleanupStop) // 停止清理后台线程
	close(c.statsStop)   // 停止统计信息打印的定时器
}

func (c *TimedLRUCacheV2) cleanupV2() {
	for {
		select {
		case <-c.cleanupTicker.C:
			fmt.Println("cleanup expired items...")
			// 定时清理过期项
			for _, shard := range c.shards {
				shard.cleanupV2()
			}
		case <-c.cleanupStop:
			c.cleanupTicker.Stop() // 停止清理定时器
			return
		}
	}
}

func (s *cacheShard) cleanupV2() {
	for {
		s.mu.Lock()
		defer s.mu.Unlock()

		if len(s.expirationHeap) == 0 {
			fmt.Println("no expired items to cleanup")
			return // 没有过期项，直接返回
		}

		top := s.expirationHeap[0]
		if top.expireAt.After(time.Now()) {
			fmt.Println("no expired items to cleanup, next item:", top.key, "expireAt:%v", top.key, top.expireAt)
			return // 没有过期项，直接返回
		}

		fmt.Println("cleanup expired item:", top.key, "expireAt:", top.expireAt)
		s.list.Remove(top.elem)                     // 从链表中删除
		heap.Remove(&s.expirationHeap, top.heapIdx) // 从堆中删除
		delete(s.items, top.key)                    // 从哈希表中删除
	}
}

func (c *TimedLRUCacheV2) Get(key string) (interface{}, bool) {
	shard := c.getShard(key)

	shard.mu.RLock()
	defer shard.mu.RUnlock()

	if item, found := shard.items[key]; found {
		if item.expireAt.After(time.Now()) { // 检查是否过期
			fmt.Println("^^^ hit key:", key, "value:", item.value, "expireAt:", item.expireAt)
			atomic.AddUint64(&shard.stats.hits, 1)
			item.expireAt = time.Now().Add(shard.ttl) // 刷新过期时间
			shard.list.MoveToFront(item.elem)         // 将该项移动到链表头部
			return item.value, true
		}
		// 如果过期，删除该项
		atomic.AddUint64(&shard.stats.misses, 1)
		shard.list.Remove(item.elem)
		delete(shard.items, key)
		return nil, false
	}
	fmt.Println("=== miss key:", key, "not found")
	atomic.AddUint64(&shard.stats.misses, 1)
	return nil, false
}

func (c *TimedLRUCacheV2) Put(key string, value interface{}) {
	shard := c.getShard(key)

	shard.mu.Lock()
	defer shard.mu.Unlock()

	if item, found := shard.items[key]; found {
		// 如果已存在，更新值和过期时间
		item.value = value
		item.expireAt = time.Now().Add(shard.ttl)
		shard.list.MoveToFront(item.elem) // 将该项移动到链表头部
		return
	}

	// 如果不存在，创建新的缓存项
	if len(shard.items) >= shard.capacity {
		// 如果达到容量限制，删除最旧的项
		oldest := shard.list.Back()
		if oldest != nil {
			oldestItem := oldest.Value.(*cacheItemV2)

			fmt.Println("||| evict key:", oldestItem.key, "expireAt:", oldestItem.expireAt, "now:", time.Now())
			shard.list.Remove(oldest)
			delete(shard.items, oldestItem.key)
			heap.Remove(&shard.expirationHeap, oldestItem.heapIdx) // 从堆中删除
			atomic.AddUint64(&shard.stats.evicts, 1)
		}
	}

	newItem := &cacheItemV2{
		key:      key,
		value:    value,
		expireAt: time.Now().Add(shard.ttl * time.Second),
	}
	elem := shard.list.PushFront(newItem)
	newItem.elem = elem // 保存双向链表元素引用

	shard.items[key] = newItem                // 添加到哈希表中
	heap.Push(&shard.expirationHeap, newItem) // 添加到堆中

	fmt.Println(">>> put key:", key, "value:", value, "expireAt:", newItem.expireAt, "now:", time.Now())
}
