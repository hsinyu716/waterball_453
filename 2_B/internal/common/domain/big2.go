package domain

import (
	"cosmos.big2/internal/common"
	"cosmos.big2/internal/common/pattern"
	"fmt"
)

type Big2 struct {
	Players    []IPlayer
	Deck       *common.Deck
	TopPlayer  IPlayer
	TopPlay    pattern.ICardPattern
	nowIndex   int
	passCount  int
	playStatus []bool
	Winner     IPlayer
}

func NewBig2(players []IPlayer, deck *common.Deck) *Big2 {
	return &Big2{
		Players:    players,
		Deck:       deck,
		passCount:  0,
		playStatus: []bool{true, true, true, true},
	}
}

func (b *Big2) Start() {
	b.Deck.Shuffle()
	b.nameSelf()
	b.DrawHand()
	b.PlayRound([]string{})
}

func (b *Big2) Start2() {
	b.DrawHand()
	b.PlayRound([]string{})
}

func (b *Big2) PlayRound(input []string) {
	// for 某玩家手牌為0  結束
	// 最初回合為梅花3玩家先出牌
	// 連續三位pass後頂牌玩家出牌
	for i, iPlayer := range b.Players {
		if iPlayer == b.TopPlayer {
			b.nowIndex = i
			break
		}
	}

	nextOne := true
	fmt.Println("新的回合開始了。")
	for {
		inputText := ""
		if len(input) > 0 {
			inputText = input[0]
			input = input[1:]
		}
		if b.playStatus[b.nowIndex%4] {
			iPlayer := b.Players[b.nowIndex%4]
			iPlayer.SetBig2(b)

			if nextOne {
				fmt.Println(fmt.Sprintf("輪到%s了", iPlayer.GetName()))
			}
			turn := iPlayer.TakeTurn(inputText)
			nextOne = true
			if turn == -1 {
				b.playStatus[b.nowIndex%4] = false
				b.passCount++
				b.nowIndex++
			} else if turn == 1 {
				b.playStatus = []bool{true, true, true, true}
				b.passCount = 0
				b.nowIndex++
			} else {
				nextOne = false
			}
			if b.passCount == 3 {
				b.passCount = 0
				b.TopPlay = nil
				b.TopPlayer = b.Players[b.nowIndex%4]
				fmt.Println("新的回合開始了。")
				b.playStatus = []bool{true, true, true, true}
			}
		} else {
			b.nowIndex++
		}
		if b.Winner != nil {
			fmt.Println(fmt.Sprintf("遊戲結束，遊戲的勝利者為 %s ", b.Winner.GetName()))
			break
		}
	}
}

func (b *Big2) nameSelf() {
	for _, iPlayer := range b.Players {
		iPlayer.NameSelf("")
	}
}

func (b *Big2) DrawHand() {
	// 因測資原因所以反向給牌
	for i := 51; i >= 0; i-- {
		card := b.Deck.DrawCard()
		player := b.Players[i%4]
		if card.GetRank() == 0 && card.GetSuit() == 0 {
			b.TopPlayer = player
		}
		if player.GetHand() == nil {
			player.SetHand(NewHand(player.GetName()))
		}
		player.GetHand().AddCard(card)
	}
}
