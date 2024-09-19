func main() {
	opts := Options{
		Feminine: true,
		Miah:     false,
		Billions: false,
		AG:       true,
	}

	converter := NumberConverter{
		Num: 1450,
		Opt: opts,
	}

	result := converter.MakeNumber()
	fmt.Println(result)
	// Output: "ألف و أربع مئة و خمسين"

}