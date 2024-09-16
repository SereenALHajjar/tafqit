package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMakeTens(t *testing.T) {
	require.Equal(t, "ثلاثون", MakeTens(3))
	require.Equal(t, "أربعون", MakeTens(4))
	require.Equal(t, "عشرة", MakeTens(1))
	require.Equal(t, "عشرون", MakeTens(2))
	require.Equal(t, "تسعون", MakeTens(9))
	require.Equal(t, "خمسون", MakeTens(5))
}

func TestMakeTwoDigitNum(t *testing.T) {
	require.Equal(t, "واحد وخمسون", MakeTwoDigitNum(51))
	require.Equal(t, "اثنان وثلاثون", MakeTwoDigitNum(32))
	require.Equal(t, "خمس وسبعون", MakeTwoDigitNum(75))

}
func TestCountDigits(t *testing.T) {
	require.Equal(t, 1, CountDigits(0))
	require.Equal(t, 1, CountDigits(1))
	require.Equal(t, 1, CountDigits(2))
	require.Equal(t, 2, CountDigits(10))
	require.Equal(t, 2, CountDigits(89))
	require.Equal(t, 3, CountDigits(243))
	require.Equal(t, 3, CountDigits(364))
	require.Equal(t, 4, CountDigits(1234))
	require.Equal(t, 5, CountDigits(12345))
	require.Equal(t, 7, CountDigits(1234567))

}

func TestMakeThreeDigitNum(t *testing.T) {
	require.Equal(t, "مئة وثلاث وأربعون", MakeThreeDigitNum(143))
	require.Equal(t, "مئة", MakeThreeDigitNum(100))
	require.Equal(t, "مئتان", MakeThreeDigitNum(200))
	require.Equal(t, "ثلاث مئة وأربع", MakeThreeDigitNum(304))
	require.Equal(t, "خمس مئة", MakeThreeDigitNum(500))
	require.Equal(t, "مئتان وست وخمسون", MakeThreeDigitNum(256))
	require.Equal(t, "ثمان مئة وخمس وتسعون", MakeThreeDigitNum(895))

}
func TestReturnLastNDigit(t *testing.T) {
	require.Equal(t, 765, extractLastNDigit(1234765, 3))
	require.Equal(t, 7659, extractLastNDigit(12347659, 4))

}
