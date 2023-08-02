package main

import "fmt"

func main() {

	var n1, operator, n2, operator2 string

	var a1, a2, result int

	var resultRoman string

	arab := map[string]int{
		"1": 1, "2": 2, "3": 3, "4": 4, "5": 5, "6": 6, "7": 7, "8": 8, "9": 9, "10": 10,
	}

	roman := map[string]int{
		"I": 1, "II": 2, "III": 3, "IV": 4, "V": 5, "VI": 6, "VII": 7, "VIII": 8, "IX": 9, "X": 10,
		"XI": 11, "XII": 12, "XIII": 13, "XIV": 14, "XV": 15, "XVI": 16, "XVII": 17, "XVIII": 18, "XIX": 19, "XX": 20,
		"XXI": 21, "XXIV": 24, "XXV": 25, "XXVII": 27, "XXVIII": 28, "XXX": 30,
		"XXXII": 32, "XXXV": 35, "XXXVI": 36, "XL": 40, "XLII": 42, "XLV": 45, "XLVIII": 48, "XLIX": 49, "L": 50,
		"LIV": 54, "LVI": 56, "LX": 60, "LXIII": 63, "LXIV": 64, "LXX": 70,
		"LXXII": 72, "LXXX": 80, "LXXXI": 81, "XC": 90, "C": 100,
	}

	fmt.Scanf("%s%s%s%s", &n1, &operator, &n2, &operator2)

	if operator == "" {

		fmt.Print("Вывод ошибки, так как строка не является математической операцией.")

	} else if operator2 != "" {

		fmt.Print("Вывод ошибки, так как формат математической операции не удовлетворяет заданию — два операнда и один оператор (+, -, /, *).")

	} else if (n1 == "1" || n1 == "2" || n1 == "3" || n1 == "4" || n1 == "5" || n1 == "6" || n1 == "7" || n1 == "8" || n1 == "9" || n1 == "10") && (n2 == "1" || n2 == "2" || n2 == "3" || n2 == "4" || n2 == "5" || n2 == "6" || n2 == "7" || n2 == "8" || n2 == "9" || n2 == "10") {

		for key, value := range arab {
			if n1 == key {

				a1 = value
			}

			if n2 == key {

				a2 = value
			}

		}

		switch operator {

		case "+":
			result = a1 + a2
		case "-":
			result = a1 - a2
		case "*":
			result = a1 * a2
		case "/":
			result = a1 / a2
		default:
			fmt.Print("Ошибка")

		}

		fmt.Print(result)

	} else if (n1 == "I" || n1 == "II" || n1 == "III" || n1 == "IV" || n1 == "V" || n1 == "VI" || n1 == "VII" || n1 == "VIII" || n1 == "IX" || n1 == "X") && (n2 == "I" || n2 == "II" || n2 == "III" || n2 == "IV" || n2 == "V" || n2 == "VI" || n2 == "VII" || n2 == "VIII" || n2 == "IX" || n2 == "X") {

		for key, value := range roman {
			if n1 == key {

				a1 = value
			}

			if n2 == key {

				a2 = value
			}

		}

		switch operator {

		case "+":
			result = a1 + a2
		case "-":
			result = a1 - a2

			if result <= 0 {

				fmt.Print("Вывод ошибки, так как в римской системе нет нуля и отрицательных чисел.")

			}

		case "*":
			result = a1 * a2
		case "/":
			result = a1 / a2
		default:
			fmt.Print("Ошибка")

		}

		for key, value := range roman {
			if result == value {

				resultRoman = key

			}
		}

		fmt.Print(resultRoman)

	} else if (n1 == "1" || n1 == "2" || n1 == "3" || n1 == "4" || n1 == "5" || n1 == "6" || n1 == "7" || n1 == "8" || n1 == "9" || n1 == "10") && (n2 == "I" || n2 == "II" || n2 == "III" || n2 == "IV" || n2 == "V" || n2 == "VI" || n2 == "VII" || n2 == "VIII" || n2 == "IX" || n2 == "X") {

		fmt.Print("Вывод ошибки, так как используются одновременно разные системы счислeния.")

	} else if (n1 == "I" || n1 == "II" || n1 == "III" || n1 == "IV" || n1 == "V" || n1 == "VI" || n1 == "VII" || n1 == "VIII" || n1 == "IX" || n1 == "X") && (n2 == "1" || n2 == "2" || n2 == "3" || n2 == "4" || n2 == "5" || n2 == "6" || n2 == "7" || n2 == "8" || n2 == "9" || n2 == "10") {

		fmt.Print("Вывод ошибки, так как используются одновременно разные системы счислeния.")

	} else {

		fmt.Print("Ошибка")
	}

}
