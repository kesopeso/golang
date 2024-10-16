package main

import (
	"context"
	"fmt"
)

type ChannelScoreboardManager chan func(map[string]int)

func (csm ChannelScoreboardManager) Update(name string, value int) {
	csm <- func(sb map[string]int) {
		sb[name] = value
	}
}

func (csm ChannelScoreboardManager) Read(name string) (int, bool) {
	type Result struct {
		value int
		ok    bool
	}

	ch := make(chan Result)

	csm <- func(sb map[string]int) {
		value, ok := sb[name]
		ch <- Result{value: value, ok: ok}
	}

	result := <-ch
	return result.value, result.ok
}

func NewChannelScoreboardManager(ctx context.Context) ChannelScoreboardManager {
	ch := make(ChannelScoreboardManager)
	scoreboard := map[string]int{} // HOW DO I READ THIS VALUE ?!?!?!?!
	go func() {
		for {
			select {
			case f := <-ch:
				f(scoreboard)
			case <-ctx.Done():
				return
			}
		}
	}()
	return ch
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	sbManager := NewChannelScoreboardManager(ctx)

	teamName := "dallas mavericks"
	sbManager.Update(teamName, 20)

	teamScore, ok := sbManager.Read(teamName)
	if !ok {
		fmt.Printf("No score found for team '%s'.\n", teamName)
		return
	}
	fmt.Printf("Team '%s' has total score %d\n", teamName, teamScore)
}
