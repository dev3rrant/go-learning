package main

import "testing"

func TestRomanNumerals(t *testing.T) {
	t.Run("Convert number to roman numeral", func(t *testing.T) {
		cases := []struct {
			Description string
			Arabic      int
			Want        string
		}{
			{"1 gets converted to I", 1, "I"},
			{"2 gets converted to II", 2, "II"},
			{"3 gets converted to III", 3, "III"},
			{"4 gets converted to IV", 4, "IV"},
			{"5 gets converted to V", 5, "V"},
			{"6 gets converted to VI", 6, "VI"},
			{"7 gets converted to VII", 7, "VII"},
			{"8 gets converted to VIII", 8, "VIII"},
			{"9 gets converted to IX", 9, "IX"},
			{"10 gets converted to X", 10, "X"},
			{"14 gets converted to XIV", 14, "XIV"},
			{"18 gets converted to XVIII", 18, "XVIII"},
			{"20 gets converted to XX", 20, "XX"},
			{"39 gets converted to XXXIX", 39, "XXXIX"},
			{"40 gets converted to XXXIX", 39, "XXXIX"},
			{"40 gets converted to XL", 40, "XL"},
			{"47 gets converted to XLVII", 47, "XLVII"},
			{"49 gets converted to XLIX", 49, "XLIX"},
			{"50 gets converted to L", 50, "L"},
			{"100 gets converted to C", 100, "C"},
			{"500 gets converted to L", 500, "D"},
			{"1000 gets converted to M", 1000, "M"},
			{"1984 gets converted to MCMLXXXIV", 1984, "MCMLXXXIV"},
		}

		for _, test := range cases {
			t.Run(test.Description, func(t *testing.T) {
				got := ConvertToRoman(test.Arabic)
				if got != test.Want {
					t.Errorf("got %v, want %v", got, test.Want)
				}
			})
		}
	})
	t.Run("Convert roman numeral to integer", func(t *testing.T) {
		cases := []struct {
			Description string
			Symbol      string
			Want        int
		}{
			{"I gets converted to 1", "I", 1},
			{"II gets converted to 2", "II", 2},
			{"III gets converted to 3", "III", 3},
			{"IV gets converted to 4", "IV", 4},
			{"V gets converted to 5", "V", 5},
			{"VI gets converted to 5", "VI", 6},
		}

		for _, test := range cases {
			t.Run(test.Description, func(t *testing.T) {
				got := ConvertToInt(test.Symbol)
				if got != test.Want {
					t.Errorf("got %v, want %v", got, test.Want)
				}
			})
		}
	})
}
