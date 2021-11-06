package hw04lrucache

import "sync"

type Key string

type Cache interface {
	Set(key Key, value interface{}) bool
	Get(key Key) (interface{}, bool)
	Clear()
}

type lruCache struct {
	mu       sync.Mutex
	capacity int
	queue    List
	items    map[Key]*ListItem
}

type cacheItem struct {
	key   string
	value interface{}
}

func NewCache(capacity int) Cache {
	return &lruCache{
		capacity: capacity,
		queue:    NewList(),
		items:    make(map[Key]*ListItem, capacity),
	}
}

func (l *lruCache) Set(key Key, value interface{}) bool {
	l.mu.Lock()
	defer l.mu.Unlock()

	item, ok := l.items[key]
	cItem := cacheItem{key: key.string(), value: value}

	if ok {
		item.Value = cItem
		l.queue.MoveToFront(item)

		return true
	}

	l.queue.PushFront(cItem)

	if l.queue.Len() > l.capacity {
		last := l.queue.Back()
		cache := last.Value.(cacheItem)
		l.queue.Remove(last)

		delete(l.items, Key(cache.key))
	}

	l.items[key] = l.queue.Front()

	return false
}

func (l *lruCache) Get(key Key) (interface{}, bool) {
	l.mu.Lock()
	defer l.mu.Unlock()

	item, ok := l.items[key]
	if !ok {
		return nil, false
	}

	l.queue.MoveToFront(item)
	cItem := item.Value.(cacheItem)

	return cItem.value, true
}

func (l *lruCache) Clear() {
	l.items = make(map[Key]*ListItem, l.capacity)
	l.queue = NewList()
}

func (k Key) string() string {
	return string(k)
}
