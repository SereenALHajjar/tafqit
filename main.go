package main

import (
	"fmt"
	"os"
)

var manazl = []string{" ", "ألف", "مليون", "مليار", "تريليون", "كوادريليون"}
var manazlInPlural = []string{" ", "آلاف", "ملايين", "مليارات", "تريليونات", "كوادريليونات"}
var numbers = []string{"صفر", "واحد", "اثنان", "ثلاث", "أربع", "خمس", "ست", "سبع", "ثمان", "تسع", "عشر"}

func MakeTens(index int) string {
	if index == 0 {
		return " "
	}
	if index == 1 {
		return "عشرة"
	}
	if index == 2 {
		return "عشرون"
	}
	number := numbers[index]
	return number + string("ون")
}
func MakeOneDigitNum(num int) string {
	if num == 0 {
		return " "
	}
	number := numbers[num]
	return number
}
func MakeTwoDigitNum(num int) string {
	if num%10 == 0 {
		return MakeTens(num / 10)
	}
	if num/10 == 1 {
		if num%10 == 1 {
			return "احدى عشر"
		}
		if num%10 == 2 {
			return "اثنا عشر"
		}
		return numbers[num%10] + string(" ") + "عشر"
	}
	return numbers[num%10] + string(" ") + string("و") + MakeTens(num/10)
}
func MakeThreeDigitNum(num int) string {
	var hundred string
	if num/100 == 1 {
		hundred = "مئة"
	} else if num/100 == 2 {
		hundred = "مئتان"
	} else {
		hundred = numbers[num/100] + string(" ") + string("مئة")
	}
	if CountDigits(num%100) == 2 {
		return hundred + string(" ") + string("و") + MakeTwoDigitNum(num%100)
	} else if CountDigits(num%100) == 1 && num%100 != 0 {
		return hundred + string(" ") + string("و") + MakeOneDigitNum(num%100)
	} else {
		return hundred
	}

}
func CountDigits(num int) int {
	if num == 0 {
		return 1
	}
	counts := 0
	for {
		if num == 0 {
			break
		}
		counts++
		num /= 10

	}
	return counts
}
func extractLastNDigit(num, n int) int {
	var digits []int
	ans := 0
	for i := 0; i < n; i++ {
		digits = append(digits, num%10)
		num /= 10
	}
	for i := n - 1; i >= 0; i-- {
		ans *= 10
		ans += digits[i]
	}
	return ans
}
func ReturnBase(num int) string {
	digits := CountDigits(num)
	if digits == 1 {
		return MakeOneDigitNum(num)
	}
	if digits == 2 {
		return MakeTwoDigitNum(num)
	}
	if digits == 3 {
		return MakeThreeDigitNum(num)
	}
	return ""
}
func searchInNumbers(num string) bool {
	for i := 1; i <= 10; i++ {
		if num == numbers[i] {
			return true
		}
	}
	return false
}
func MakeNumber(num int) string {
	if num == 0 {
		return numbers[0]
	}
	var numberStr []string
	for {
		if CountDigits(num) >= 3 {
			lastThreeDigitNumber := extractLastNDigit(num, 3)
			numberStr = append(numberStr, ReturnBase(lastThreeDigitNumber))
			num /= 1000
		} else if digits := CountDigits(num); digits < 3 && digits != 0 {
			lastNDigitNumber := extractLastNDigit(num, digits)
			numberStr = append(numberStr, ReturnBase(lastNDigitNumber))
			num = 0
		}
		if num == 0 {
			break
		}
	}
	var final string
	for i := 0; i < len(numberStr); i++ {
		if numberStr[i] != " " && manazl[i] != " " {
			if numberStr[i] == numbers[1] {
				numberStr[i] = manazl[i]
			} else if numberStr[i] == numbers[2] {
				numberStr[i] = manazl[i] + "ان"
			} else if searchInNumbers(numberStr[i]) {
				numberStr[i] += string(" ")
				numberStr[i] += manazlInPlural[i]
			} else {
				numberStr[i] += string(" ")
				numberStr[i] += manazl[i]
			}
		}
	}
	for i := len(numberStr) - 1; i >= 0; i-- {
		if numberStr[i] == " " {
			continue
		}
		if i+1 < len(numberStr) && numberStr[i+1] != " " {
			numberStr[i] = "و" + numberStr[i]
		}
		final += " "
		final += numberStr[i]
	}
	return final
}

func main() {
	var index int
	file, err := os.Create("output.txt")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close() // Ensure the file is closed at the end

	// Write a string to the file
	for {
		fmt.Scan(&index)

		_, err = file.WriteString(MakeNumber(index) + "\n")
		if err != nil {
			fmt.Println("Error writing to file:", MakeNumber(index))
			return
		}

		fmt.Println("Data written to file successfully.")
		if index == 999 {
			break
		}
	}

}
