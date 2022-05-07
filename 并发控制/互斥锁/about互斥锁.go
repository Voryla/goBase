package main

import "sync"

func main() {
	lock := sync.Mutex{}
	lock.Lock()
	lock.Unlock()
}
