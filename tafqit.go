package tafqit

import "strings"

var manazl = []string{" ", "ألف", "مليون", "مليار", "تريليون", "كوادريليون"}
var manazlInPlural = []string{" ", "آلاف", "ملايين", "مليارات", "تريليونات", "كوادريليونات"}
var numbers = [][]string{
	{"صفر", "واحد", "اثنان", "ثلاث", "أربع", "خمس", "ست", "سبع", "ثمان", "تسع", "عشر"},
	{"صفر", "واحدة", "اثنتان", "ثلاثة", "أربعة", "خمسة", "ستة", "سبعة", "ثمانية", "تسعة", "عشرة"},
}

type Options struct {
	Feminine  bool
	Miah      bool
	SplitHund bool
	Billions  bool
	AG        bool
}

type NumberConverter struct {
	Num int
	Opt Options
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
func (cnv *NumberConverter) MakeTens(index int) string {
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
func (cnv *NumberConverter) MakeOneDigitNum(num int) string {
	if num == 0 {
		return " "
	}
	return handleFeminine(num, cnv.Opt.Feminine, cnv.Opt.AG)
}
func (cnv *NumberConverter) MakeTwoDigitNum(num int) string {
	if num%10 == 0 {
		return cnv.MakeTens(num / 10)
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
	return handleFeminine(num%10, cnv.Opt.Feminine, cnv.Opt.AG) + string(" ") + string("و") + cnv.MakeTens(num/10)
}

func (cnv *NumberConverter) MakeThreeDigitNum(num int) string {
	var hundred string
	if num/100 == 1 {
		hundred = "مئة"
	} else if num/100 == 2 {
		hundred = "مئتان"
		if cnv.Opt.AG {
			hundred = "مئتين"
		}
	} else {
		hundred = handleFeminine(num/100, cnv.Opt.Feminine, cnv.Opt.AG) + string(" ") + string("مئة")
	}
	if countsDigits(num%100) == 2 {
		return hundred + string(" ") + string("و") + cnv.MakeTwoDigitNum(num%100)
	} else if countsDigits(num%100) == 1 && num%100 != 0 {
		return hundred + string(" ") + string("و") + cnv.MakeOneDigitNum(num%100)
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

func (cnv *NumberConverter) ReturnBase(num int) string {
	digits := countsDigits(num)
	if digits == 1 {
		return cnv.MakeOneDigitNum(num)
	}
	if digits == 2 {
		return cnv.MakeTwoDigitNum(num)
	}
	if digits == 3 {
		return cnv.MakeThreeDigitNum(num)
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

	var numberStr []string
	for {
		if countsDigits(cnv.Num) >= 3 {
			lastThreeDigitNumber := extractLastNDigit(cnv.Num, 3)
			numberStr = append(numberStr, cnv.ReturnBase(lastThreeDigitNumber))
			cnv.Num /= 1000
		} else if digits := countsDigits(cnv.Num); digits < 3 && digits != 0 {
			lastNDigitNumber := extractLastNDigit(cnv.Num, digits)
			numberStr = append(numberStr, cnv.ReturnBase(lastNDigitNumber))
			cnv.Num = 0
		}
		if cnv.Num == 0 {
			break
		}
	}
	var final string
	for i := 0; i < len(numberStr); i++ {
		if numberStr[i] != " " && manazl[i] != " " {
			currentManzlah := manazl[i]
			currentPluralManzlah := manazl[i]
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

	return strings.TrimSpace(final)
}

func main() {
	// var index int
	// var cnv NumberConverter
	// file, err := os.Create("output.txt")
	// if err != nil {
	// 	fmt.Println("Error creating file:", err)
	// 	return
	// }
	// defer file.Close() // Ensure the file is closed at the end

	// // Write a string to the file
	// for {
	// 	fmt.Scan(&index)

	// 	_, err = file.WriteString(cnv.MakeNumber(index) + "\n")
	// 	if err != nil {
	// 		fmt.Println("Error writing to file:", cnv.MakeNumber(index))
	// 		return
	// 	}

	// 	fmt.Println("Data written to file successfully.")
	// 	if index == 999 {
	// 		break
	// 	}
	// }

}
