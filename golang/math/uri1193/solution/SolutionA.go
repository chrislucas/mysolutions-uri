package main

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"
)

// https://www.urionlinejudge.com.br/judge/en/problems/view/1193 DONE

var TableDecToHex = map[int64]string {
	10 : "a",
	11 : "b",
	12 : "c",
	13 : "d",
	14 : "e",
	15 : "f",
}

var TableHexToDec = map[string]int64 {
	"a": 10,
	"b": 11,
	"c": 12,
	"d": 13,
	"e": 14,
	"f": 15,
}

// https://golang.org/pkg/strconv/
var parseInt = strconv.ParseInt

var formatInt = strconv.FormatInt

func reverse(target string) string{
	arr := [] rune(target)
	for i, j :=0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
	return string(arr)
}

func fromDecToSomeBase(value string, base int64) string {
	decValue, _ := parseInt(value, 10, 64)
	var buffer bytes.Buffer
	for decValue > 0 {
		r := decValue % base
		decValue /= base
		if base == 16 {
			if r >= 10 && r <= 15 {
				fmt.Fprintf(&buffer, "%s", TableDecToHex[r])
			} else {
				fmt.Fprintf(&buffer, "%d", r)
			}
		} else {
			fmt.Fprintf(&buffer, "%d", r)
		}
	}
	return reverse(buffer.String())
}

func fromSomeBaseToDec(value string, base int64) string {
	reversed := reverse(value)
	var acc int64 = 0
	var accBase int64 = 1
	for _, v := range reversed {
		if base == 16 {
			if newValue, ok := TableHexToDec[string(v)]; ok  {
				acc +=  newValue * accBase
				accBase *= base
			} else {
				v, _ := parseInt(string(v), 10, 64)
				acc +=  v * accBase
				accBase *= base
			}
		} else {
			v, _ := parseInt(string(v), 10, 64)
			acc +=  v * accBase
			accBase *= base
		}
	}
	return formatInt(acc, 10)
}

func fromBinToHex(value string) string {
	return fromDecToSomeBase(fromSomeBaseToDec(value, 2), 16)
}

func fromBinToDec(value string) string {
	return fromSomeBaseToDec(value, 2)
}

func fromHexToBin(value string) string {
	return fromDecToSomeBase(fromSomeBaseToDec(value, 16), 2)
}

func fromHexToDec(value string) string {
	return fromSomeBaseToDec(value, 16)
}

func fromDecToBin(value string) string {
	return fromDecToSomeBase(value, 2)
}

func fromDecToHex(value string) string {
	return fromDecToSomeBase(value, 16)
}


func solver() {
	var cases int
	fmt.Scanf("%d", &cases)
	for i :=1; i<=cases; i++ {
		var number string
		var base string
		fmt.Scanf("%s %s", &number, &base)
		fmt.Printf("Case %d:\n", i)
		number = strings.ToLower(number)
		switch base {
			case "bin" :
				fmt.Printf("%s dec\n", fromBinToDec(number))
				fmt.Printf("%s hex\n", fromBinToHex(number))
			case "hex":
				fmt.Printf("%s dec\n", fromHexToDec(number))
				fmt.Printf("%s bin\n", fromHexToBin(number))
			default:
				fmt.Printf("%s hex\n", fromDecToHex(number))
				fmt.Printf("%s bin\n", fromDecToBin(number))
		}
		fmt.Println("")
	}
}

func main() {
	solver()
}
