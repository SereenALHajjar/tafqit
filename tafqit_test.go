package tafqit

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMakeTens(t *testing.T) {
	opt := Options{
		Feminine: false,
		AG:       false,
		Miah:     false,
		Billions: false,
	}
	cnv := NumberConverter{
		Opt: opt,
		Num: 1,
	}
	require.Equal(t, "ثلاثون", cnv.makeTens(3))
	require.Equal(t, "أربعون", cnv.makeTens(4))
	require.Equal(t, "عشرة", cnv.makeTens(1))
	require.Equal(t, "عشرون", cnv.makeTens(2))
	require.Equal(t, "تسعون", cnv.makeTens(9))
	require.Equal(t, " ", cnv.makeTens(0))
	require.Equal(t, "خمسون", cnv.makeTens(5))
	cnv.Opt.AG = true
	require.Equal(t, "ثلاثين", cnv.makeTens(3))
	require.Equal(t, "أربعين", cnv.makeTens(4))
	require.Equal(t, "عشرة", cnv.makeTens(1))
	require.Equal(t, "عشرين", cnv.makeTens(2))
	require.Equal(t, "تسعين", cnv.makeTens(9))
	require.Equal(t, "خمسين", cnv.makeTens(5))
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
		Feminine: true,
		AG:       false,
		Miah:     false,
		Billions: false,
	}
	cnv := NumberConverter{
		Opt: opt,
		Num: 1,
	}
	require.Equal(t, " ", cnv.makeOneDigitNum(0))
	require.Equal(t, "اثنتان", cnv.makeOneDigitNum(2))
	require.Equal(t, "واحدة", cnv.makeOneDigitNum(1))
	require.Equal(t, "ثلاث", cnv.makeOneDigitNum(3))
	require.Equal(t, "أربع", cnv.makeOneDigitNum(4))
	cnv.Opt.Feminine = false
	require.Equal(t, "ثلاثة", cnv.makeOneDigitNum(3))
	require.Equal(t, "أربعة", cnv.makeOneDigitNum(4))
	require.Equal(t, "اثنان", cnv.makeOneDigitNum(2))
	require.Equal(t, "واحد", cnv.makeOneDigitNum(1))

	cnv.Opt.AG = true
	require.Equal(t, "اثنين", cnv.makeOneDigitNum(2))
	require.Equal(t, "واحد", cnv.makeOneDigitNum(1))
	require.Equal(t, "ثلاثة", cnv.makeOneDigitNum(3))
	require.Equal(t, "أربعة", cnv.makeOneDigitNum(4))
	cnv.Opt.Feminine = true
	require.Equal(t, "اثنتين", cnv.makeOneDigitNum(2))
	require.Equal(t, "واحدة", cnv.makeOneDigitNum(1))
	require.Equal(t, "ثلاث", cnv.makeOneDigitNum(3))
	require.Equal(t, "أربع", cnv.makeOneDigitNum(4))
}

func TestMakeTwoDigitNum(t *testing.T) {
	opt := Options{
		Feminine: false,
		AG:       false,
		Miah:     false,
		Billions: false,
	}
	cnv := NumberConverter{
		Opt: opt,
		Num: 1,
	}
	require.Equal(t, "واحد وخمسون", cnv.makeTwoDigitNum(51))
	require.Equal(t, "اثنان وثلاثون", cnv.makeTwoDigitNum(32))
	require.Equal(t, "خمسة وسبعون", cnv.makeTwoDigitNum(75))
	require.Equal(t, "اثنا عشر", cnv.makeTwoDigitNum(12))
	require.Equal(t, "أحد عشر", cnv.makeTwoDigitNum(11))
	require.Equal(t, "خمسة عشر", cnv.makeTwoDigitNum(15))
	require.Equal(t, "تسعة عشر", cnv.makeTwoDigitNum(19))
	require.Equal(t, "ثمانية وثمانون", cnv.makeTwoDigitNum(88))
	require.Equal(t, "تسعة وأربعون", cnv.makeTwoDigitNum(49))
	cnv.Opt.AG = true
	require.Equal(t, "اثني عشر", cnv.makeTwoDigitNum(12))
	cnv.Opt.AG = false
	cnv.Opt.Feminine = true
	require.Equal(t, "اثنتا عشرة", cnv.makeTwoDigitNum(12))
	require.Equal(t, "احدى عشرة", cnv.makeTwoDigitNum(11))
	cnv.Opt.AG = true

	require.Equal(t, "واحدة وخمسين", cnv.makeTwoDigitNum(51))
	require.Equal(t, "اثنتين وثلاثين", cnv.makeTwoDigitNum(32))
	require.Equal(t, "خمس وسبعين", cnv.makeTwoDigitNum(75))
	require.Equal(t, "اثنتي عشرة", cnv.makeTwoDigitNum(12))
	require.Equal(t, "احدى عشرة", cnv.makeTwoDigitNum(11))
	require.Equal(t, "خمس عشرة", cnv.makeTwoDigitNum(15))
	require.Equal(t, "تسع عشرة", cnv.makeTwoDigitNum(19))
	require.Equal(t, "ثمان وثمانين", cnv.makeTwoDigitNum(88))
	require.Equal(t, "تسع وأربعين", cnv.makeTwoDigitNum(49))
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
		Feminine: false,
		AG:       false,
		Miah:     false,
		Billions: false,
	}
	cnv := NumberConverter{
		Opt: opt,
		Num: 1,
	}
	require.Equal(t, "مئة وثلاثة وأربعون", cnv.makeThreeDigitNum(143))
	require.Equal(t, "مئة", cnv.makeThreeDigitNum(100))
	require.Equal(t, "مئتان", cnv.makeThreeDigitNum(200))
	require.Equal(t, "ثلاثة مئة وأربعة", cnv.makeThreeDigitNum(304))
	require.Equal(t, "خمسة مئة", cnv.makeThreeDigitNum(500))
	require.Equal(t, "مئتان وستة وخمسون", cnv.makeThreeDigitNum(256))
	require.Equal(t, "ثمانية مئة وخمسة وتسعون", cnv.makeThreeDigitNum(895))
	require.Equal(t, "ثمانية مئة واثنا عشر", cnv.makeThreeDigitNum(812))
	require.Equal(t, "ثمانية مئة وأحد عشر", cnv.makeThreeDigitNum(811))
	cnv.Opt.Feminine = true
	require.Equal(t, "مئة وثلاث وأربعون", cnv.makeThreeDigitNum(143))
	require.Equal(t, "مئة", cnv.makeThreeDigitNum(100))
	require.Equal(t, "مئتان", cnv.makeThreeDigitNum(200))
	require.Equal(t, "ثلاث مئة وأربع", cnv.makeThreeDigitNum(304))
	require.Equal(t, "خمس مئة", cnv.makeThreeDigitNum(500))
	require.Equal(t, "مئتان وست وخمسون", cnv.makeThreeDigitNum(256))
	require.Equal(t, "ثمان مئة وخمس وتسعون", cnv.makeThreeDigitNum(895))
	require.Equal(t, "ثمان مئة واثنتا عشرة", cnv.makeThreeDigitNum(812))
	require.Equal(t, "ثمان مئة واحدى عشرة", cnv.makeThreeDigitNum(811))
	cnv.Opt.AG = true
	require.Equal(t, "مئة وثلاث وأربعين", cnv.makeThreeDigitNum(143))
	require.Equal(t, "مئة", cnv.makeThreeDigitNum(100))
	require.Equal(t, "مئتين", cnv.makeThreeDigitNum(200))
	require.Equal(t, "ثلاث مئة وأربع", cnv.makeThreeDigitNum(304))
	require.Equal(t, "خمس مئة", cnv.makeThreeDigitNum(500))
	require.Equal(t, "مئتين وست وخمسين", cnv.makeThreeDigitNum(256))
	require.Equal(t, "ثمان مئة وخمس وتسعين", cnv.makeThreeDigitNum(895))
	require.Equal(t, "ثمان مئة واثنتي عشرة", cnv.makeThreeDigitNum(812))
	require.Equal(t, "ثمان مئة واحدى عشرة", cnv.makeThreeDigitNum(811))
	cnv.Opt.Miah = true
	require.Equal(t, "مائة وثلاث وأربعين", cnv.makeThreeDigitNum(143))
	require.Equal(t, "مائة", cnv.makeThreeDigitNum(100))
	require.Equal(t, "مائتين", cnv.makeThreeDigitNum(200))
	require.Equal(t, "ثلاث مائة وأربع", cnv.makeThreeDigitNum(304))
	require.Equal(t, "خمس مائة", cnv.makeThreeDigitNum(500))
	require.Equal(t, "مائتين وست وخمسين", cnv.makeThreeDigitNum(256))
	require.Equal(t, "ثمان مائة وخمس وتسعين", cnv.makeThreeDigitNum(895))
	require.Equal(t, "ثمان مائة واثنتي عشرة", cnv.makeThreeDigitNum(812))
	require.Equal(t, "ثمان مائة واحدى عشرة", cnv.makeThreeDigitNum(811))

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
		Feminine: false,
		AG:       false,
		Miah:     false,
		Billions: false,
	}
	cnv := NumberConverter{
		Opt: opt,
		Num: 1,
	}
	require.Equal(t, "مئة وثلاثة وأربعون", cnv.returnBase(143))
	require.Equal(t, "واحد وخمسون", cnv.returnBase(51))
	require.Equal(t, "ثلاثة", cnv.returnBase(3))
	require.Equal(t, " ", cnv.returnBase(1234))

}

func TestMakeNumber(t *testing.T) {
	opt := Options{
		Feminine: false,
		AG:       false,
		Miah:     false,
		Billions: false,
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
	require.Equal(t, "أربعة بلايين وتسعة مئة وثمانية وسبعين مليون وستة مئة وأربعة وخمسين ألف ومئة وعشرين", cnv.MakeNumber())
	cnv.Num = 2978654120
	require.Equal(t, "بليونين وتسعة مئة وثمانية وسبعين مليون وستة مئة وأربعة وخمسين ألف ومئة وعشرين", cnv.MakeNumber())
	cnv.Num = -5
	require.Equal(t, "سالب خمسة", cnv.MakeNumber())
	cnv.Num = 2000033
	require.Equal(t, "مليونين وثلاثة وثلاثين", cnv.MakeNumber())

}
func TestHandleTwoHandred(t *testing.T) {
	require.Equal(t, "مائتين", handleTwoHaundred(true, true))
	require.Equal(t, "مئتين", handleTwoHaundred(true, false))
	require.Equal(t, "مائتان", handleTwoHaundred(false, true))
	require.Equal(t, "مئتان", handleTwoHaundred(false, false))
}
func TestHandleMiah(t *testing.T) {
	require.Equal(t, "مئة", handleMiah(false))
	require.Equal(t, "مائة", handleMiah(true))

}
