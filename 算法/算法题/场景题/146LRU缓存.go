package 场景题

type LRUCache struct {
	capacity  int
	values    map[int]int
	cacheLink *CacheLink
}

type CacheLink struct {
	next *CacheLink
	val  int
	pre  *CacheLink
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		capacity: capacity,
		values:   make(map[int]int),
		cacheLink: &CacheLink{
			next: nil,
			val:  0,
			pre:  nil,
		},
	}
}

func (this *LRUCache) Get(key int) int {
	v := this.values[key]
	// slow path
	if v == 0 {
		index := this.index(key)
		// 若值不存在
		if index == -1 {
			return -1
		}
	}
	// 若值存在，调整相应的缓存序列，返回值
	this.moveCache(key)
	return v
}

func (this *LRUCache) Put(key int, value int) {

	index := this.index(key)
	if len(this.values) >= this.capacity && index == -1 {
		delete(this.values, this.cacheLink.next.val)
		// 删除最久未使用
		if this.cacheLink.next.next == nil {
			this.cacheLink.next = nil
		} else {
			this.cacheLink.next = this.cacheLink.next.next
			this.cacheLink.next.pre = this.cacheLink
		}
	}
	this.values[key] = value
	if index == -1 { // 添加新的缓存至链表尾部
		temp := this.cacheLink
		for temp.next != nil {
			temp = temp.next
		}
		temp.next = &CacheLink{
			next: nil,
			val:  key,
			pre:  temp,
		}
	} else { // 由于对原有缓存进行了赋值，调整缓存
		this.moveCache(key)
	}
}

func (this *LRUCache) index(key int) int {
	temp := this.cacheLink.next
	for i := 0; i < len(this.values); i++ {
		if key == temp.val {
			return i
		}
		temp = temp.next
	}
	return -1
}

func (this *LRUCache) moveCache(key int) {
	temp := this.cacheLink
	var target *CacheLink
	for temp.next != nil {
		if temp.val == key && temp.pre != nil {
			temp.pre.next = temp.next
			temp.next.pre = temp.pre
			target = temp
		}
		temp = temp.next
	}
	if target == nil {
		return
	}
	temp.next = target
	target.pre = temp
	target.next = nil
}
