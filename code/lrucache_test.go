package code

import (
	"testing"
	"time"
)

func TestTimedLruCache(t *testing.T) {
	cache := NewTimedLRUCache(2, 5) // 容量为2，过期时间为5秒

	// 测试put和get
	cache.Put("key1", "value1")
	if val, ok := cache.Get("key1"); !ok || val != "value1" {
		t.Errorf("Expected value1, got %v", val)
	}

	// 测试过期
	cache.Put("key2", "value2")
	cache.Put("key3", "value3") // 这将导致key1被淘汰
	if _, ok := cache.Get("key1"); ok {
		t.Error("Expected key1 to be evicted")
	}

	// 等待过期时间
	time.Sleep(6 * time.Second)
	cache.Put("key4", "value4")
	if val, ok := cache.Get("key2"); ok || val == "value2" {
		t.Errorf("Expected value2 exipired, got value2")
	}
}
