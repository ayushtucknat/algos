package algo_expert

var data = make(map[string]bool)

// https://www.algoexpert.io/questions/non-constructible-change

func NonConstructibleChange(coins []int) int {
	min := 1
	for {
		constructed := canChangeBeConstructed(coins, min)
		if !constructed {
			return min
		}
		min++
	}
}

func canChangeBeConstructed(coins []int, num int) bool {
	if len(coins) == 0 {
		return false
	}

	for _, coin := range coins {
		if coin == num {
			return true
		}
	}
	for index, coin := range coins {
		if coin < num {
			remainingCoins := make([]int, 0)
			copy(remainingCoins, coins[:index])
			if index < len(coins)-1 {
				remainingCoins = append(remainingCoins, coins[index+1:]...)
			}
			if canChangeBeConstructed(remainingCoins, num-coin) {
				return true
			}
		}
	}
	return false
}
