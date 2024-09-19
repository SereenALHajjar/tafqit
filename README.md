# Tafqit - Number to Text Conversion (تفقيط)

**Tafqit** is a Go package that converts numbers into their equivalent text representation, commonly used in check writing and financial documents. This process is also known as **tafqit** in Arabic, where the numbers are transformed into text to avoid ambiguity or fraud.

## Features
- Converts numbers to their textual form (in Arabic).
- Supports integer numbers.
- Easy to use, with a simple API.

## Installation

To use this package in your Go project, first install it using `go get`:

```bash
go get github.com/SereenALHajjar/tafqit
```
Then, import the package in your code:

```go
import "github.com/SereenALHajjar/tafqit"
``` 

## Usage
### Example: Convert Number to Text
Here’s an example of how to use the MakeNumber function with the NumberConverter struct to convert an integer to its Arabic text equivalent, with options:
```go
package main

import (
    "fmt"
    "github.com/SereenALHajjar/tafqit"
)

func main() {
    // Example number to convert
    number := 123456
    
    // Set up options for the conversion
    opts := tafqit.Options{
        Feminine: false, // masculine form
        Miah:     true,  // use "مائة" instead of "مئة"
        Billions: true,  // use "بليون" instead of "مليار"
        AG:       false, // mean the statue is nominative 
    }

    // Initialize the converter with the number and options
    converter := tafqit.NumberConverter{
        Num: number,
        Opt: opts,
    }

    // Convert the number to its textual form
    result := converter.MakeNumber()

    // Print the result
    fmt.Println(result)
}
``` 
## output :
```text
مائة وثلاثة وعشرون ألف وأربعة مائة وستة وخمسون
```
In this example, the number 123456 is converted to the Arabic text representation "مائة و ثلاثة و عشرون ألفا و أربعة مائة و ستة و خمسون" with the specified options.

## API
### ``func (cnv *NumberConverter) MakeNumber() string``
- **Description** : Converts the number in the NumberConverter struct to its equivalent Arabic text representation based on the options provided.
- **Structs**: NumberConverter: Contains the number to be converted and options for how the conversion should be handled.
Options: Contains flags to customize the conversion (feminine, miah, billions, and Arabic grammar).

### ``Options:``
| Option | Type | Description |
| ------ | ---- | ----------- |
| **Feminine** | **bool** |	If **true**, converts the form to the feminine form according to Arabic grammar.
| **Miah** | **bool** | If **true**, use "مائة" instead of "مئة".
| **Billions** | **bool** | If **true**, use "بليون" instead of "مليار".
| **AG** | **bool** | If **true**, mean the status is Accusative/Genitive the default is nominative 

## Example:
```go
opts := tafqit.Options{
    Feminine: true,
    Miah:     false,
    Billions: false,
    AG:       true,
}

converter := tafqit.NumberConverter{
    Num: 1450,
    Opt: opts,
}

result := converter.MakeNumber()
fmt.Println(result) 
// Output: "ألف و أربع مئة و خمسين"

```
## Running Tests
To run the unit tests, use the following command:
```bash 
go test ./...
```
This will execute all the tests in the package to ensure the correctness of the tafqit functionality.

## Contributing
If you'd like to contribute to this project, feel free to fork the repository and submit a pull request. Issues and feature requests are also welcome!