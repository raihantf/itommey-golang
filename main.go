package main

import "fmt"

func main() {
	segitigaSamaSisi(5)
}

func segitigaSamaSisi(baris int) {
	jarak := baris - 1
	for i := 1; i <= baris; i++ {
		for j := 1; j <= jarak; j++ {
			fmt.Print(" ")
		}
		for k := 1; k <= 2*i-1; k++ {
			if k%2 == 1 {
				fmt.Print("1")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
		jarak--
	}
}
