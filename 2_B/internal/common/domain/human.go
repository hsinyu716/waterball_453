package domain

import (
	"bufio"
	"cosmos.big2/internal/common/pattern"
	"cosmos.big2/internal/common/patternHandler"
	"cosmos.big2/internal/common/poker"
	"cosmos.big2/internal/utils"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Human struct {
	Player
}

func NewHuman() IPlayer {
	return &Human{
		Player{},
	}
}

func (h *Human) NameSelf(name string) {
	if name == "" {
		fmt.Print("Enter text: ")
		_, _ = fmt.Scanf("%s", &name)
	}
	h.name = name
}

func (h *Human) TakeTurn(input string) int {
	change := false
	h.GetHand().CardList()
	//fmt.Print("輸入要打的牌: ")
	reader := bufio.NewReader(os.Stdin)
	if text, _ := reader.ReadString('\n'); true {
		text = strings.Trim(text, " \n")
		var (
			keyIndex []string
			cards    []*poker.Card
		)

		if len(input) > 0 {
			text = input
		}
		keyIndex = strings.Split(text, " ")

		//validator := NewPassValidator(NewClubValidator(NewSameValidator(NewPatternValidator(nil))))
		//
		//validateEntity := ValidatorEntity{
		//	Hand:     h.GetHand(),
		//	KeyIndex: keyIndex,
		//	Cards:    nil,
		//	Player:   h,
		//	TopPlay:  h.GetBig2().TopPlay,
		//}
		//return validator.Validation(validateEntity)

		// todo: 驗可否pass
		keyIndex, noErr, i, done := pass(keyIndex, h)
		if done {
			return i
		}

		if len(text) > 0 && noErr {
			// todo: 驗梅花3
			noErr = club(h.GetHand(), keyIndex)
			if noErr {
				// todo: 驗打牌正確性
				cards, noErr, change = checkSame(keyIndex, cards, h.GetHand().cards, noErr, change)
				if noErr {
					// todo: 驗牌型合法
					cardPattern := pattern.NewPatternSingle(cards, pattern.NewPatternPair(cards, pattern.NewPatternStraight(cards, pattern.NewPatternFullHouse(cards, nil))))
					if cardPattern == nil {
						fmt.Println("此牌型不合法，請再嘗試一次。")
						noErr = false
						change = false
					} else {
						fmt.Print(fmt.Sprintf("玩家 %s ", h.GetName()))
						handler :=
							patternHandler.NewSingleHandler(
								patternHandler.NewPairHandler(
									patternHandler.NewStraightHandler(
										patternHandler.NewFullHouseHandler(nil))))
						handle := handler.Handle(cardPattern, h.big2.TopPlay)
						if handle == 1 {
							h.hand.PlayCard(cards)
							if h.hand.Size() == 0 {
								h.big2.Winner = h
								return 1
							}
							h.big2.TopPlayer = h
							h.big2.TopPlay = cardPattern
							fmt.Println(fmt.Sprintf("目前頂牌玩家為 %s, 頂牌為 %s ", h.big2.TopPlayer.GetName(), h.big2.TopPlay.ShowCard()))
							return 1
						} else if handle == 0 {
							fmt.Println(fmt.Sprintf("牌型比頂牌小, 頂牌為%s", h.big2.TopPlay.ShowCard()))
						}
					}
				}
			}
		}
	}
	return 0
}

func checkSame(keyIndex []string, cards []*poker.Card, handCard []*poker.Card, noErr bool, change bool) ([]*poker.Card, bool, bool) {
	sameMap := make(map[int]bool)
	same := false
	for _, s := range keyIndex {
		if s == "" {
			break
		}
		atoi, _ := strconv.Atoi(s)
		if sameMap[atoi] {
			same = true
			break
		}
		sameMap[atoi] = true
		cards = append(cards, handCard[atoi])
	}
	// 不可出同一張牌
	if same {
		fmt.Println("不可出同一張牌")
		noErr = false
		change = false
	}
	return cards, noErr, change
}

func club(hand *Hand, index []string) bool {
	clubIndex, hasClub := hand.HasClub()
	if hasClub {
		exists, _ := utils.InArray(strconv.Itoa(clubIndex), index)
		// todo: 判斷出梅花
		if !exists {
			fmt.Println("第一回合請輸入包含梅花3的牌型。")
			return false
		}
	}
	return true
}

func pass(index []string, h *Human) ([]string, bool, int, bool) {
	noErr := true
	if len(index) == 1 && index[0] == "-1" {
		if h.big2.TopPlayer != h {
			fmt.Println(fmt.Sprintf("玩家 %s PASS, 頂牌為 %s", h.GetName(), h.big2.TopPlay.ShowCard()))
			return nil, false, -1, true
		} else {
			fmt.Println(fmt.Sprintf("你不能在新的回合中喊 PASS"))
			index = []string{}
			noErr = false
		}
	}
	return index, noErr, 0, false
}
