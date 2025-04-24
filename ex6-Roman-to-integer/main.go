package main

func romantoint(s string) int {

	roman := map[string]int{
		"I": 1,
		"V": 5,
		"X": 10,
		"L": 100,
		"D": 500,
		"M": 1000,
	}
	result := 0
	for i, v := range s {
		result := roman[string(v)]

	}

}
