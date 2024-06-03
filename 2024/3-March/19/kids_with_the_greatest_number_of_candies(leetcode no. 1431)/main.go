package main

func main() {

}

func kidsWithCandies(candies []int, extraCandies int) []bool {
	maxCandies := -1
	res := []bool{}

	for _, candy := range candies {
		if candy > maxCandies {
			maxCandies = candy
		}
	}

	for _, candy := range candies {
		res = append(res, candy+extraCandies >= maxCandies)
	}

	return res
}
