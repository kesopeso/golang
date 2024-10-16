package main

import (
	"fmt"
	"sync"
)

type ScoreboardManager struct {
	scoreboard map[string]int
	l          sync.RWMutex
}

func (sm *ScoreboardManager) Update(name string, val int) {
	sm.l.Lock()
	defer sm.l.Unlock()
	sm.scoreboard[name] = val
}

func (sm *ScoreboardManager) Read(name string) (int, bool) {
	sm.l.RLock()
	defer sm.l.RUnlock()
	result, ok := sm.scoreboard[name]
	return result, ok
}

func main() {
	sm := ScoreboardManager{
		scoreboard: map[string]int{},
	}
	sm.Update("dallas", 20)
	val, ok := sm.Read("dallas")
	if !ok {
		fmt.Println("dallas is not added")
		return
	}
	fmt.Printf("dallas score is %d\n", val)
}
