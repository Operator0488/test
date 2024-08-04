package main

import (
	"fmt"
)

func main() {

	fmt.Println("Введите выражение:")

	arabToArab := map[string]int{
		"1": 1, "2": 2, "3": 3, "4": 4, "5": 5,
		"6": 6, "7": 7, "8": 8, "9": 9, "10": 10,
	}

	romToArab := map[string]int{
		"I": 1, "II": 2, "III": 3, "IV": 4, "V": 5,
		"VI": 6, "VII": 7, "VIII": 8, "IX": 9, "X": 10,
	}

	arabToRom := map[int]string{
		0: " ", 1: "I", 2: "II", 3: "III", 4: "IV", 5: "V",
		6: "VI", 7: "VII", 8: "VIII", 9: "IX", 10: "X",
		11: "XI", 12: "XII", 13: "XIII", 14: "IXV", 15: "XV",
		16: "XVI", 17: "XVII", 18: "XVIII", 19: "XIX", 20: "XX", 30: "XXX",
		40: "XL", 50: "L", 60: "LX", 70: "LXX", 80: "LXXX", 90: "XC", 100: "C",
	}

	var b, a, d, c string
	var a1, c1 int
	fmt.Scanln(&a, &b, &c, &d)

	z1 := arabToArab[a]
	z2 := romToArab[a]

	if z1 != 0 && z2 == 0 {
		a1 = arabToArab[a]
		c1 = arabToArab[c]
	} else if z2 != 0 && z1 == 0 {
		a1 = romToArab[a]
		c1 = romToArab[c]
	} else if z2 != 0 && z1 != 0 {
		panic("паника0")
	} else if z2 == 0 && z1 == 0 {
		panic("паника паника")
	}

	if a1 == 0 || c1 == 0 {
		panic("паника1")
	} else if d != "" {
		panic("паника2")
	} else {
		if b == "+" {
			sum := a1 + c1
			switch {
			case z2 == 0:
				fmt.Println(sum)
			case z1 == 0:
				fmt.Println(arabToRom[sum])
			default:
				panic("паника sum")
			}
		} else if b == "-" {
			switch {
			case z2 == 0:
				minu := a1 - c1
				fmt.Println(minu)
			case z1 == 0:
				minu := a1 - c1
				if minu < 1 {
					panic("римская паника")
				} else {
					fmt.Println(arabToRom[minu])
				}
			default:
				panic("паника minu")
			}
		} else if b == "*" {
			umn := a1 * c1
			switch {
			case z2 == 0:
				fmt.Println(umn)
			case z1 == 0:
				switch {
				case umn <= 20:
					fmt.Println(arabToRom[umn])
				case umn >= 30 && umn < 40:
					fmt.Print(arabToRom[30], arabToRom[umn%30])
				case umn >= 40 && umn < 50:
					fmt.Print(arabToRom[40], arabToRom[umn%40])
				case umn >= 50 && umn < 60:
					fmt.Print(arabToRom[50], arabToRom[umn%50])
				case umn >= 60 && umn < 70:
					fmt.Print(arabToRom[60], arabToRom[umn%60])
				case umn >= 70 && umn < 80:
					fmt.Print(arabToRom[70], arabToRom[umn%70])
				case umn >= 80 && umn < 90:
					fmt.Print(arabToRom[80], arabToRom[umn%80])
				case umn >= 90 && umn < 100:
					fmt.Print(arabToRom[90], arabToRom[umn%90])
				case umn == 100:
					fmt.Print(arabToRom[100])
				}
			default:
				panic("паника umn")
			}
		} else if b == "/" {
			del := a1 / c1
			switch {
			case z2 == 0:
				fmt.Println(del)
			case z1 == 0:
				fmt.Println(arabToRom[del])
			default:
				panic("паника del")
			}
		} else {
			panic("паника3")
		}
	}

}
