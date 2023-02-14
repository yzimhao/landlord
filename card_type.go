package landlord

import (
	"sort"
)

// 斗地主有以下几种出牌类型：
//     单牌：一张牌。
//     对子：两张牌，权值相同。
//     三张牌：三张牌，权值相同。
//     三带一：三张牌加上一张单牌，权值相同。
//     三带二：三张牌加上两张牌，权值相同。
//     顺子：五张或五张以上的牌，按照顺序连续的排列。
//     连对：三对或三对以上的对子，按照顺序连续的排列。
//     飞机：两个或两个以上的三张牌，按照顺序连续的排列。
//     飞机带单：两个或两个以上的三张牌加上若干张单牌，按照顺序连续的排列。
//     飞机带对：两个或两个以上的三张牌加上若干对子，按照顺序连续的排列。
//     炸弹：四张牌，权值相同。
//     王炸：两张牌，分别是大王和小王。

type CardType int

const (
	unknown CardType = iota
	danpai
	duizi
	sanbudai
	sandaiyi
	sandaier
	shunzi
	liandui
	feiji
	feijidaidan
	feijidaidui
	zhadan
	wangzha
)

func (ct CardType) String() string {
	str := "未知"
	switch ct {
	case danpai:
		str = "单牌"
	case duizi:
		str = "对子"
	case sanbudai:
		str = "三不带"
	case sandaiyi:
		str = "三带一"
	case sandaier:
		str = "三带二"
	case shunzi:
		str = "顺子"
	case liandui:
		str = "连对"
	case feiji:
		str = "飞机"
	case feijidaidan:
		str = "飞机带单"
	case feijidaidui:
		str = "飞机带对"
	case zhadan:
		str = "炸弹"
	case wangzha:
		str = "王炸"
	}
	return str
}

func CheckCardsType(cards []Card) CardType {
	if len(cards) == 1 {
		return danpai
	}

	if len(cards) == 2 {
		if cards[0].Value == cards[1].Value {
			return duizi
		}
		if cards[0].Point() == Smalljoker || cards[0].Point() == Bigjoker {
			return wangzha
		}
	}
	if len(cards) == 3 {
		if cards[0].Value == cards[1].Value && cards[1].Value == cards[2].Value {
			return sanbudai
		}
	}
	if len(cards) == 4 {
		if cards[0].Value == cards[1].Value && cards[1].Value == cards[2].Value && cards[2].Value == cards[3].Value {
			return zhadan
		}
		if isThreeAndOne(cards) {
			return sandaiyi
		}
	}
	if len(cards) == 5 {
		if isStraight(cards) {
			return shunzi
		}
		if isThreeAndTwo(cards) {
			return sandaier
		}
	}
	if len(cards) >= 6 {
		if isStraight(cards) {
			return shunzi
		}
		if isDoubleStraight(cards) {
			return liandui
		}
		if isPlane(cards) {
			return feiji
		}
		if isPlaneAndOne(cards) {
			return feijidaidan
		}
		if isPlaneAndTwo(cards) {
			return feijidaidui
		}
	}
	return unknown
}

func isThreeAndOne(cards []Card) bool {
	countMap := make(map[int]int)
	for _, card := range cards {
		countMap[card.Value]++
	}

	threeCount := 0
	for _, count := range countMap {
		if count == 3 {
			threeCount++
		}
	}

	return threeCount == 1 && len(cards) == 4
}

func isStraight(cards []Card) bool {
	// 判断牌的数量是否不足 5 张
	if len(cards) < 5 {
		return false
	}

	// 对牌进行排序
	sort.Slice(cards, func(i, j int) bool { return cards[i].Value < cards[j].Value })

	// 遍历牌
	for i := 0; i < len(cards)-1; i++ {
		// 判断当前牌与下一张牌的权值差是否为 1
		if cards[i].Value-cards[i+1].Value != -1 {
			return false
		}
	}

	// 如果所有牌都满足条件，则返回 true
	return true
}

func isThreeAndTwo(cards []Card) bool {
	cardCounts := make(map[int]int)
	for _, card := range cards {
		cardCounts[card.Value]++
	}

	hasThree := false
	hasTwo := false
	for _, count := range cardCounts {
		if count == 3 {
			hasThree = true
		} else if count == 2 {
			hasTwo = true
		}
	}

	return hasThree && hasTwo
}

func isPlane(cards []Card) bool {
	count := make(map[int]int)
	for _, card := range cards {
		count[card.Value]++
	}
	var last, length int
	for value, c := range count {
		if c >= 3 {
			if last == 0 {
				last = value
				length = c
			} else if value == last+1 && c == length {
				last = value
			} else {
				return false
			}
		}
	}
	return length >= 2
}

func isDoubleStraight(cards []Card) bool {
	if len(cards) < 6 {
		return false
	}

	values := make([]int, 0, len(cards))
	for _, card := range cards {
		values = append(values, card.Value)
	}

	sort.Ints(values)
	countMap := make(map[int]int)
	for _, value := range values {
		countMap[value]++
	}

	// Check if all cards are pairs
	for _, count := range countMap {
		if count != 2 {
			return false
		}
	}

	// // Check if cards form a straight
	// for i := 0; i < len(values)-1; i++ {
	// 	if values[i+1] != string(int(values[i][0])+1) {
	// 		return false
	// 	}
	// }

	return true
}

func isPlaneAndOne(cards []Card) bool {
	if len(cards) < 6 {
		return false
	}

	var values []int
	for _, card := range cards {
		values = append(values, card.Value)
	}

	sort.Ints(values)

	count := make(map[int]int)
	for _, value := range values {
		count[value]++
	}

	var planeStart, planeLen int
	for i := 0; i < len(values)-2; i++ {
		if values[i]+1 == values[i+1] && values[i+1]+1 == values[i+2] {
			if planeLen == 0 {
				planeStart = values[i]
				planeLen = 3
			} else if values[i] == planeStart+planeLen {
				planeLen += 3
			} else {
				return false
			}
		}
	}

	if planeLen == 0 {
		return false
	}

	for i := planeStart; i < planeStart+planeLen; i++ {
		if count[i] < 3 {
			return false
		}
		count[i] -= 3
	}

	oneCount := 0
	for _, v := range count {
		if v >= 1 {
			oneCount++
		}
	}

	return oneCount >= planeLen/3
}

func isPlaneAndTwo(hand []Card) bool {
	if len(hand) < 8 {
		return false
	}

	countMap := make(map[int]int)
	for _, card := range hand {
		countMap[card.Value]++
	}

	var threeCount int
	for _, count := range countMap {
		if count == 3 {
			threeCount++
		}
	}

	if threeCount < 2 {
		return false
	}

	var lastValue int
	for value, count := range countMap {
		if count == 3 {
			if lastValue != 0 && value-lastValue != 1 {
				return false
			}
			lastValue = value
		}
	}

	var twoCount int
	for _, count := range countMap {
		if count == 2 {
			twoCount++
		}
	}

	return twoCount >= (len(hand)/5)*2
}
