package main

func main() {

}

func lemonadeChange(bills []int) bool {
	billCount := make(map[int]int)

	for _, bill := range bills {
		billCount[bill] += 1
		changeNeeded := bill - 5
		if changeNeeded == 0 {
			continue
		}

		for _, b := range []int{20, 10, 5} {
			if b > changeNeeded {
				continue
			}

			currBillCount := billCount[b]
			if currBillCount == 0 {
				continue
			}

			countNeeded := changeNeeded / b

			if countNeeded >= currBillCount {
				billCount[b] = 0
				changeNeeded -= currBillCount * b
			} else {
				billCount[b] = currBillCount - countNeeded
				changeNeeded -= countNeeded * b
			}

		}

		if changeNeeded != 0 {
			return false
		}
	}

	return true

}
