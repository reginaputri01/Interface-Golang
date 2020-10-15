package main

import "fmt"

func main() {
	var cetakSegitiga cetakSegitiga = &segitiga{5}

	fmt.Println(cetakSegitiga.rataKiri())
	fmt.Println(cetakSegitiga.rataKanan())
}

type cetakSegitiga interface {
	rataKiri() string
	rataKanan() string
}

type segitiga struct {
	value int
}

func (a *segitiga) rataKiri() string {
	cetak := ""
	for b := 0; b < a.value; b++ {
		cetakKiri := ""
		for c := 0; c <= b; c++ {
			cetakKiri += "*"
		}
		fmt.Println(cetakKiri)
	}
	return cetak
}

func (a *segitiga) rataKanan() string {
	cetak := ""
	for b := 0; b < a.value; b++ {
		cetakKanan := ""
		for c := 0; c < a.value; c++ {
			if c >= b {
				cetakKanan += "*"
			} else {
				cetakKanan += " "
			}
		}
		fmt.Println(cetakKanan)
	}
	return cetak
}