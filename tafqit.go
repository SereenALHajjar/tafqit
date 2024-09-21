package tafqit

import (
	"strings"
	"unicode"
)

var manazl = []string{" ", "ألف", "مليون", "مليار", "تريليون", "كوادريليون"}
var manazlInPlural = []string{" ", "آلاف", "ملايين", "مليارات", "تريليونات", "كوادريليونات"}
var numbers = [][]string{
	{"صفر", "واحد", "اثنان", "ثلاث", "أربع", "خمس", "ست", "سبع", "ثمان", "تسع", "عشر"},
	{"صفر", "واحدة", "اثنتان", "ثلاثة", "أربعة", "خمسة", "ستة", "سبعة", "ثمانية", "تسعة", "عشرة"},
}

type Options struct {
	Feminine bool
	Miah     bool
	Billions bool
	AG       bool
}

type NumberConverter struct {
	Num int
	Opt Options
}

func handleMiah(Miah bool) string {
	if Miah {
		return "مائة"

	}
	return "مئة"

}
func handleAG(index int, AG bool) string {
	if AG {
		return numbers[0][index] + "ين"
	}
	return numbers[0][index] + "ون"
}
func handleFeminine(index int, feminine bool, AG bool) string {
	if feminine {
		if index == 1 {
			return numbers[1][index]
		}
		if index == 2 {
			if AG {
				return "اثنتين"
			}
			return "اثنتان"
		}
		return numbers[0][index]
	}

	if index == 1 {
		return numbers[0][index]
	}
	if index == 2 {
		if AG {
			return "اثنين"
		}
		return "اثنان"
	}
	return numbers[1][index]
}
func (cnv *NumberConverter) makeTens(index int) string {
	if index == 0 {
		return " "
	}
	if index == 1 {
		return "عشرة"
	}
	if index == 2 {
		// that's because 20 is عشرين or عشرون it's have a 10 prefix
		return handleAG(10, cnv.Opt.AG)
	}
	return handleAG(index, cnv.Opt.AG)
}
func (cnv *NumberConverter) makeOneDigitNum(num int) string {
	if num == 0 {
		return " "
	}
	return handleFeminine(num, cnv.Opt.Feminine, cnv.Opt.AG)
}
func (cnv *NumberConverter) makeTwoDigitNum(num int) string {
	if num%10 == 0 {
		return cnv.makeTens(num / 10)
	}
	if num/10 == 1 {
		if num%10 == 1 {
			if cnv.Opt.Feminine {
				return "احدى عشرة"
			}
			return "أحد عشر"
		}
		if num%10 == 2 {
			if cnv.Opt.Feminine {
				if cnv.Opt.AG {
					return "اثنتي عشرة"
				}
				return "اثنتا عشرة"
			} else {
				if cnv.Opt.AG {
					return "اثني عشر"
				}
				return "اثنا عشر"
			}
		}
		ten := numbers[0][10]
		if cnv.Opt.Feminine {
			ten = numbers[1][10]
		}
		return handleFeminine(num%10, cnv.Opt.Feminine, cnv.Opt.AG) + string(" ") + ten
	}
	return handleFeminine(num%10, cnv.Opt.Feminine, cnv.Opt.AG) + string(" ") + string("و") + cnv.makeTens(num/10)
}
func handleTwoHaundred(AG, Miah bool) string {
	if AG {
		if Miah {
			return "مائتين"
		}
		return "مئتين"
	}
	if Miah {
		return "مائتان"
	}
	return "مئتان"
}
func (cnv *NumberConverter) makeThreeDigitNum(num int) string {
	var hundred string
	if num/100 == 1 {
		hundred = handleMiah(cnv.Opt.Miah)
		// hundred = "مئة"
	} else if num/100 == 2 {
		// hundred = "مئتان"
		// if cnv.Opt.AG {
		// 	hundred = "مئتين"
		// }
		hundred = handleTwoHaundred(cnv.Opt.AG, cnv.Opt.Miah)
	} else {
		hundred = handleFeminine(num/100, cnv.Opt.Feminine, cnv.Opt.AG) + string(" ") + handleMiah(cnv.Opt.Miah)
	}
	if countsDigits(num%100) == 2 {
		return hundred + string(" ") + string("و") + cnv.makeTwoDigitNum(num%100)
	} else if countsDigits(num%100) == 1 && num%100 != 0 {
		return hundred + string(" ") + string("و") + cnv.makeOneDigitNum(num%100)
	} else {
		return hundred
	}

}
func countsDigits(num int) int {
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

func (cnv *NumberConverter) returnBase(num int) string {
	digits := countsDigits(num)
	if digits == 1 {
		return cnv.makeOneDigitNum(num)
	}
	if digits == 2 {
		return cnv.makeTwoDigitNum(num)
	}
	if digits == 3 {
		return cnv.makeThreeDigitNum(num)
	}
	return " "
}

func searchInNumbers(num string) bool {
	for i := 1; i <= 10; i++ {
		if num == numbers[0][i] || num == numbers[1][i] {
			return true
		}
	}
	return false
}

func (cnv *NumberConverter) MakeNumber() string {
	if cnv.Num == 0 {
		return numbers[0][0]
	}
	var final string
	var numberStr []string
	if cnv.Num < 0 {
		final = "سالب "
		cnv.Num *= -1
	}
	for {
		if countsDigits(cnv.Num) >= 3 {
			lastThreeDigitNumber := extractLastNDigit(cnv.Num, 3)
			numberStr = append(numberStr, cnv.returnBase(lastThreeDigitNumber))
			cnv.Num /= 1000
		} else if digits := countsDigits(cnv.Num); digits < 3 && digits != 0 {
			lastNDigitNumber := extractLastNDigit(cnv.Num, digits)
			numberStr = append(numberStr, cnv.returnBase(lastNDigitNumber))
			cnv.Num = 0
		}
		if cnv.Num == 0 {
			break
		}
	}
	for i := 0; i < len(numberStr); i++ {
		if numberStr[i] != " " && manazl[i] != " " {
			currentManzlah := manazl[i]
			currentPluralManzlah := manazlInPlural[i]
			if manazl[i] == "مليار" && cnv.Opt.Billions {
				currentManzlah = "بليون"
				currentPluralManzlah = "بليونات"
			}
			if numberStr[i] == numbers[0][1] || numberStr[i] == numbers[1][1] {
				// مليار , مليون
				numberStr[i] = currentManzlah
			} else if numberStr[i] == "اثنان" || numberStr[i] == "اثنين" || numberStr[i] == "اثنتين" || numberStr[i] == "اثنتان" {
				// ملياران , مليونان في حالة الرفع أو  مليونين في حالة النصب والجر
				numberStr[i] = currentManzlah + "ان"
				if cnv.Opt.AG {
					numberStr[i] = currentManzlah + "ين"
				}
			} else if searchInNumbers(numberStr[i]) {
				// من 3 الى 9 يجب ان تكون المنزلة جمع اي ثلاثة الاف
				numberStr[i] += string(" ")
				numberStr[i] += currentPluralManzlah
			} else {
				// default
				numberStr[i] += string(" ")
				numberStr[i] += currentManzlah
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
	return strings.TrimSpace(removeConsecutiveSpaces(final))

}

func removeConsecutiveSpaces(input string) string {
	var result strings.Builder
	spaceFound := false

	for _, char := range input {
		if unicode.IsSpace(char) {
			if !spaceFound {
				result.WriteRune(' ')
				spaceFound = true
			}
		} else {
			result.WriteRune(char)
			spaceFound = false
		}
	}

	return result.String()
}
