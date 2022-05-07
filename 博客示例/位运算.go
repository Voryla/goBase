package main

import (
	"fmt"
	"strings"
)

type KeySet byte

const (
	Copper  KeySet = 1 << iota // 铜钥匙
	Jade                       // 玉钥匙
	Crystal                    // 水晶钥匙
	maxKey
)

// Player 玩家
type Player struct {
	Name string
	Keys KeySet
}

func (p *Player) addKey(key KeySet) {
	p.Keys |= key
}

func (p *Player) hasKey(key KeySet) bool {
	return p.Keys&key != 0
}

func (p *Player) removeKey(key KeySet) {
	p.Keys &= ^key
}

func (k KeySet) String() string {
	if k > maxKey {
		return fmt.Sprintf("<unkonwn key: %d>", k)
	}

	switch k {
	case Copper:
		return "copper"
	case Jade:
		return "jade"
	case Crystal:
		return "crystal"
	}

	var names []string
	for key := Copper; key < maxKey; key <<= 1 {
		if k&key != 0 {
			names = append(names, key.String())
		}
	}
	return strings.Join(names, "|")
}
func main() {
	p := Player{Name: "zwk", Keys: Copper}
	p.addKey(Copper)
	fmt.Println(p.Keys)
	p.addKey(Jade)
	p.removeKey(Copper)
	fmt.Println(p.Keys)
	fmt.Println(p.hasKey(Jade))
}
