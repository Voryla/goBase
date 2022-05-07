package aboutArray

type LRUCache struct {
	capacity int
	data     map[int]int
}

func Constructor(capacity int) LRUCache {
	return LRUCache{capacity: capacity, data: make(map[int]int)}
}

func (this *LRUCache) Get(key int) int {
	return 1
}

func (this *LRUCache) Put(key int, value int) {

}
func main() {

}