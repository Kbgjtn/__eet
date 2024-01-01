package cases

// @case: 1672 Richest Customer Wealth
func RichestCustomerWealth(accounts [][]int) int {
	theRichest := 0
	for i := 0; i < len(accounts); i++ {
		currentWealth := 0
		for j := 0; j < len(accounts[i]); j++ {
			currentWealth += accounts[i][j]
		}

		if currentWealth > theRichest {
			theRichest = currentWealth
		}

	}
	return theRichest
}

func MaximumWealth(accounts [][]int) int {
	var maxI, sumJ int
	for i := range accounts {
		for j := range accounts[i] {
			sumJ += accounts[i][j]
		}
		if sumJ >= maxI {
			maxI = sumJ
		}
		sumJ = 0
	}
	return maxI
}
