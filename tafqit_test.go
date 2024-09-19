package tafqit

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMakeTens(t *testing.T) {
	opt := Options{
		Feminine:  false,
		AG:        false,
		Miah:      false,
		SplitHund: false,
		Billions:  false,
	}
	cnv := NumberConverter{
		Opt: opt,
		Num: 1,
	}
	require.Equal(t, "ثلاثون", cnv.MakeTens(3))
	require.Equal(t, "أربعون", cnv.MakeTens(4))
	require.Equal(t, "عشرة", cnv.MakeTens(1))
	require.Equal(t, "عشرون", cnv.MakeTens(2))
	require.Equal(t, "تسعون", cnv.MakeTens(9))
	require.Equal(t, " ", cnv.MakeTens(0))
	require.Equal(t, "خمسون", cnv.MakeTens(5))
	cnv.Opt.AG = true
	require.Equal(t, "ثلاثين", cnv.MakeTens(3))
	require.Equal(t, "أربعين", cnv.MakeTens(4))
	require.Equal(t, "عشرة", cnv.MakeTens(1))
	require.Equal(t, "عشرين", cnv.MakeTens(2))
	require.Equal(t, "تسعين", cnv.MakeTens(9))
	require.Equal(t, "خمسين", cnv.MakeTens(5))
}

func TestHandleAG(t *testing.T) {
	require.Equal(t, "عشرون", handleAG(10, false))
	require.Equal(t, "عشرين", handleAG(10, true))
	require.Equal(t, "ثلاثون", handleAG(3, false))
	require.Equal(t, "ثلاثين", handleAG(3, true))
	require.Equal(t, "أربعون", handleAG(4, false))
	require.Equal(t, "أربعين", handleAG(4, true))
	require.Equal(t, "خمسون", handleAG(5, false))
	require.Equal(t, "خمسين", handleAG(5, true))
}

func TestHandleFeminine(t *testing.T) {
	require.Equal(t, "اثنتان", handleFeminine(2, true, false))
	require.Equal(t, "اثنان", handleFeminine(2, false, false))
	require.Equal(t, "اثنتين", handleFeminine(2, true, true))
	require.Equal(t, "اثنين", handleFeminine(2, false, true))
	require.Equal(t, "واحدة", handleFeminine(1, true, false))
	require.Equal(t, "واحد", handleFeminine(1, false, false))
	require.Equal(t, "ثلاث", handleFeminine(3, true, false))
	require.Equal(t, "ثلاثة", handleFeminine(3, false, false))
	require.Equal(t, "أربع", handleFeminine(4, true, false))
	require.Equal(t, "أربعة", handleFeminine(4, false, false))
}

func TestMakeOneDigit(t *testing.T) {
	opt := Options{
		Feminine:  true,
		AG:        false,
		Miah:      false,
		SplitHund: false,
		Billions:  false,
	}
	cnv := NumberConverter{
		Opt: opt,
		Num: 1,
	}
	require.Equal(t, " ", cnv.MakeOneDigitNum(0))
	require.Equal(t, "اثنتان", cnv.MakeOneDigitNum(2))
	require.Equal(t, "واحدة", cnv.MakeOneDigitNum(1))
	require.Equal(t, "ثلاث", cnv.MakeOneDigitNum(3))
	require.Equal(t, "أربع", cnv.MakeOneDigitNum(4))
	cnv.Opt.Feminine = false
	require.Equal(t, "ثلاثة", cnv.MakeOneDigitNum(3))
	require.Equal(t, "أربعة", cnv.MakeOneDigitNum(4))
	require.Equal(t, "اثنان", cnv.MakeOneDigitNum(2))
	require.Equal(t, "واحد", cnv.MakeOneDigitNum(1))

	cnv.Opt.AG = true
	require.Equal(t, "اثنين", cnv.MakeOneDigitNum(2))
	require.Equal(t, "واحد", cnv.MakeOneDigitNum(1))
	require.Equal(t, "ثلاثة", cnv.MakeOneDigitNum(3))
	require.Equal(t, "أربعة", cnv.MakeOneDigitNum(4))
	cnv.Opt.Feminine = true
	require.Equal(t, "اثنتين", cnv.MakeOneDigitNum(2))
	require.Equal(t, "واحدة", cnv.MakeOneDigitNum(1))
	require.Equal(t, "ثلاث", cnv.MakeOneDigitNum(3))
	require.Equal(t, "أربع", cnv.MakeOneDigitNum(4))
}

func TestMakeTwoDigitNum(t *testing.T) {
	opt := Options{
		Feminine:  false,
		AG:        false,
		Miah:      false,
		SplitHund: false,
		Billions:  false,
	}
	cnv := NumberConverter{
		Opt: opt,
		Num: 1,
	}
	require.Equal(t, "واحد وخمسون", cnv.MakeTwoDigitNum(51))
	require.Equal(t, "اثنان وثلاثون", cnv.MakeTwoDigitNum(32))
	require.Equal(t, "خمسة وسبعون", cnv.MakeTwoDigitNum(75))
	require.Equal(t, "اثنا عشر", cnv.MakeTwoDigitNum(12))
	require.Equal(t, "أحد عشر", cnv.MakeTwoDigitNum(11))
	require.Equal(t, "خمسة عشر", cnv.MakeTwoDigitNum(15))
	require.Equal(t, "تسعة عشر", cnv.MakeTwoDigitNum(19))
	require.Equal(t, "ثمانية وثمانون", cnv.MakeTwoDigitNum(88))
	require.Equal(t, "تسعة وأربعون", cnv.MakeTwoDigitNum(49))
	cnv.Opt.AG = true
	require.Equal(t, "اثني عشر", cnv.MakeTwoDigitNum(12))
	cnv.Opt.AG = false
	cnv.Opt.Feminine = true
	require.Equal(t, "اثنتا عشرة", cnv.MakeTwoDigitNum(12))
	require.Equal(t, "احدى عشرة", cnv.MakeTwoDigitNum(11))
	cnv.Opt.AG = true

	require.Equal(t, "واحدة وخمسين", cnv.MakeTwoDigitNum(51))
	require.Equal(t, "اثنتين وثلاثين", cnv.MakeTwoDigitNum(32))
	require.Equal(t, "خمس وسبعين", cnv.MakeTwoDigitNum(75))
	require.Equal(t, "اثنتي عشرة", cnv.MakeTwoDigitNum(12))
	require.Equal(t, "احدى عشرة", cnv.MakeTwoDigitNum(11))
	require.Equal(t, "خمس عشرة", cnv.MakeTwoDigitNum(15))
	require.Equal(t, "تسع عشرة", cnv.MakeTwoDigitNum(19))
	require.Equal(t, "ثمان وثمانين", cnv.MakeTwoDigitNum(88))
	require.Equal(t, "تسع وأربعين", cnv.MakeTwoDigitNum(49))
}

func TestCountDigits(t *testing.T) {
	require.Equal(t, 1, countsDigits(0))
	require.Equal(t, 1, countsDigits(1))
	require.Equal(t, 1, countsDigits(2))
	require.Equal(t, 2, countsDigits(10))
	require.Equal(t, 2, countsDigits(89))
	require.Equal(t, 3, countsDigits(243))
	require.Equal(t, 3, countsDigits(364))
	require.Equal(t, 4, countsDigits(1234))
	require.Equal(t, 5, countsDigits(12345))
	require.Equal(t, 7, countsDigits(1234567))
}

func TestMakeThreeDigitNum(t *testing.T) {
	opt := Options{
		Feminine:  false,
		AG:        false,
		Miah:      false,
		SplitHund: false,
		Billions:  false,
	}
	cnv := NumberConverter{
		Opt: opt,
		Num: 1,
	}
	require.Equal(t, "مئة وثلاثة وأربعون", cnv.MakeThreeDigitNum(143))
	require.Equal(t, "مئة", cnv.MakeThreeDigitNum(100))
	require.Equal(t, "مئتان", cnv.MakeThreeDigitNum(200))
	require.Equal(t, "ثلاثة مئة وأربعة", cnv.MakeThreeDigitNum(304))
	require.Equal(t, "خمسة مئة", cnv.MakeThreeDigitNum(500))
	require.Equal(t, "مئتان وستة وخمسون", cnv.MakeThreeDigitNum(256))
	require.Equal(t, "ثمانية مئة وخمسة وتسعون", cnv.MakeThreeDigitNum(895))
	require.Equal(t, "ثمانية مئة واثنا عشر", cnv.MakeThreeDigitNum(812))
	require.Equal(t, "ثمانية مئة وأحد عشر", cnv.MakeThreeDigitNum(811))
	cnv.Opt.Feminine = true
	require.Equal(t, "مئة وثلاث وأربعون", cnv.MakeThreeDigitNum(143))
	require.Equal(t, "مئة", cnv.MakeThreeDigitNum(100))
	require.Equal(t, "مئتان", cnv.MakeThreeDigitNum(200))
	require.Equal(t, "ثلاث مئة وأربع", cnv.MakeThreeDigitNum(304))
	require.Equal(t, "خمس مئة", cnv.MakeThreeDigitNum(500))
	require.Equal(t, "مئتان وست وخمسون", cnv.MakeThreeDigitNum(256))
	require.Equal(t, "ثمان مئة وخمس وتسعون", cnv.MakeThreeDigitNum(895))
	require.Equal(t, "ثمان مئة واثنتا عشرة", cnv.MakeThreeDigitNum(812))
	require.Equal(t, "ثمان مئة واحدى عشرة", cnv.MakeThreeDigitNum(811))
	cnv.Opt.AG = true
	require.Equal(t, "مئة وثلاث وأربعين", cnv.MakeThreeDigitNum(143))
	require.Equal(t, "مئة", cnv.MakeThreeDigitNum(100))
	require.Equal(t, "مئتين", cnv.MakeThreeDigitNum(200))
	require.Equal(t, "ثلاث مئة وأربع", cnv.MakeThreeDigitNum(304))
	require.Equal(t, "خمس مئة", cnv.MakeThreeDigitNum(500))
	require.Equal(t, "مئتين وست وخمسين", cnv.MakeThreeDigitNum(256))
	require.Equal(t, "ثمان مئة وخمس وتسعين", cnv.MakeThreeDigitNum(895))
	require.Equal(t, "ثمان مئة واثنتي عشرة", cnv.MakeThreeDigitNum(812))
	require.Equal(t, "ثمان مئة واحدى عشرة", cnv.MakeThreeDigitNum(811))

}

func TestReturnLastNDigit(t *testing.T) {
	require.Equal(t, 765, extractLastNDigit(1234765, 3))
	require.Equal(t, 7659, extractLastNDigit(12347659, 4))
}
func TestSearchInNumbers(t *testing.T) {
	require.Equal(t, true, searchInNumbers("واحد"))
	require.Equal(t, true, searchInNumbers("واحدة"))
	require.Equal(t, true, searchInNumbers("عشرة"))
	require.Equal(t, true, searchInNumbers("ثلاثة"))
	require.Equal(t, false, searchInNumbers("تين"))
	require.Equal(t, false, searchInNumbers("مربع"))
}

func TestReturnBase(t *testing.T) {
	opt := Options{
		Feminine:  false,
		AG:        false,
		Miah:      false,
		SplitHund: false,
		Billions:  false,
	}
	cnv := NumberConverter{
		Opt: opt,
		Num: 1,
	}
	require.Equal(t, "مئة وثلاثة وأربعون", cnv.ReturnBase(143))
	require.Equal(t, "واحد وخمسون", cnv.ReturnBase(51))
	require.Equal(t, "ثلاثة", cnv.ReturnBase(3))
	require.Equal(t, " ", cnv.ReturnBase(1234))

}

func TestMakeNumber(t *testing.T) {
	opt := Options{
		Feminine:  false,
		AG:        false,
		Miah:      false,
		SplitHund: false,
		Billions:  false,
	}
	cnv := NumberConverter{
		Opt: opt,
		Num: 1650212,
	}
	require.Equal(t, "مليون وستة مئة وخمسون ألف ومئتان واثنا عشر", cnv.MakeNumber())
	cnv.Num = 49800
	require.Equal(t, "تسعة وأربعون ألف وثمانية مئة", cnv.MakeNumber())
	cnv.Num = 0
	require.Equal(t, "صفر", cnv.MakeNumber())
	cnv.Num = 2000000
	require.Equal(t, "مليونان", cnv.MakeNumber())
	cnv.Opt.AG = true
	cnv.Num = 2000000
	require.Equal(t, "مليونين", cnv.MakeNumber())
	cnv.Num = 4978654120
	cnv.Opt.Billions = true
	require.Equal(t, "أربعة بليونات وتسعة مئة وثمانية وسبعين مليون وستة مئة وأربعة وخمسين ألف ومئة وعشرين", cnv.MakeNumber())
	cnv.Num = 2978654120
	require.Equal(t, "بليونين وتسعة مئة وثمانية وسبعين مليون وستة مئة وأربعة وخمسين ألف ومئة وعشرين", cnv.MakeNumber())

}
