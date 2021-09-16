package main

import (
	"fmt"
	"strings"
)

type list []string

func (l list) has(item string) bool {
	item = strings.ToLower(item)
	for _, i := range l {
		if i == item {
			return true
		}
	}
	return false
}

func showMenu() {
	fmt.Println("Toko Makanan Indonesia")
	fmt.Println("======================")
	fmt.Println("Tahu")
	fmt.Println("Tempe")
	fmt.Println("Sambel")
	fmt.Println("Nasi")
	fmt.Println("Lele")
	fmt.Println("Ayam")
	fmt.Println("======================")
}

func showData(l list) {
	for _, item := range l {
		fmt.Println("Pesanan anda :", item)
	}
}

func getInput(data list, l list) {
	var item, choice string
	for cond1 := true; cond1; cond1 = strings.ToLower(choice) == "y" {
		for cond2 := true; cond2; cond2 = !l.has(item) {
			fmt.Print("Masukan menu pesanan anda dalam huruf (eg: Tahu) : ")
			fmt.Scanln(&item)
			if l.has(item) {
				data = append(data, strings.ToLower(item))
			}
		}
		fmt.Print("Lanjutkan memesan ? (Y/T) : ")
		fmt.Scanln(&choice)
	}
	showData(data)
}
func main() {
	var data list
	items := list{"tahu", "tempe", "sambel", "nasi", "lele", "ayam"}
	showMenu()
	getInput(data, items)

}
