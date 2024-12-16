package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

func FirstTask(numDecimal, numOctal, numHexadecimal int, float float64, str string, boolean bool, complexNum complex64) {

	fmt.Printf("Десятичное число: %v, тип: %v\n", numDecimal, reflect.TypeOf(numDecimal))
	fmt.Printf("Восьмиричное число: %v, тип: %v\n", numOctal, reflect.TypeOf(numOctal))
	fmt.Printf("Шестнадцатиричное число: %v, тип: %v\n", numHexadecimal, reflect.TypeOf(numHexadecimal))
	fmt.Printf("Дробное число: %v, тип: %v\n", float, reflect.TypeOf(float))
	fmt.Printf("Строка: %v, тип: %v\n", str, reflect.TypeOf(str))
	fmt.Printf("Логическое значение: %v, тип: %v\n", boolean, reflect.TypeOf(boolean))
	fmt.Printf("Комплексное число: %v, тип: %v\n", complexNum, reflect.TypeOf(complexNum))

	var builder strings.Builder

	builder.WriteString(strconv.Itoa(numDecimal))
	builder.WriteString(" ")
	builder.WriteString(strconv.Itoa(numOctal))
	builder.WriteString(" ")
	builder.WriteString(strconv.Itoa(numHexadecimal))
	builder.WriteString(" ")
	builder.WriteString(strconv.FormatFloat(float, 'f', -1, 64))
	builder.WriteString(" ")
	builder.WriteString(str)
	builder.WriteString(" ")
	builder.WriteString(strconv.FormatBool(boolean))
	builder.WriteString(" ")
	builder.WriteString(fmt.Sprintf("%v", complexNum))

	res := builder.String()

	fmt.Println("Объединенная строка:", res)

	runeSlice := []rune(res)

	salt := "go-2024"
	middle := len(runeSlice) / 2
	runesWithSalt := append(runeSlice[:middle], append([]rune(salt), runeSlice[middle:]...)...)

	hasher := sha256.New()
	hasher.Write([]byte(string(runesWithSalt)))
	hashString := hex.EncodeToString(hasher.Sum(nil))

	fmt.Println("Хэшированный результат:", hashString)
}

func main() {
	FirstTask(42, 052, 0x42, 3.14, "Hello", true, 1+2i)
}
