package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
)

func main() {
	p := []byte{2, 3, 5}
	p = AppendByte(p, 7, 11, 13)
	fmt.Println(p)
	mu := []int{1, 2, 4, 4, 5, 6, 7}
	nu := mu[1:3]
	zu := mu[1:3:3]
	fmt.Printf("length of %s is %d and its capacity is %d\n", "mu", len(mu), cap(mu))
	fmt.Printf("length of %s is %d and its capacity is %d\n", "nu", len(nu), cap(nu))
	fmt.Printf("length of %s is %d and its capacity is %d\n", "zu", len(zu), cap(zu))
	for i := range mu {
		fmt.Printf("%d,", mu[i])
	}
	fmt.Println("")
	for i := range nu {
		fmt.Printf("%d,", nu[i])
	}
	fmt.Println("")
	FindDigits("./sieve123.go")
}

func AppendByte(slice []byte, data ...byte) []byte {
	m := len(slice)
	n := m + len(data)
	if n > cap(slice) {
		newSlice := make([]byte, (n+1)*2)
		copy(newSlice, slice)
		slice = newSlice
	}
	slice = slice[0:n]
	copy(slice[m:n], data)
	return slice
}

var digitRegexp = regexp.MustCompile("[0-9]+")

func FindDigits(filename string) []byte {
	b, _ := ioutil.ReadFile(filename)
	return digitRegexp.Find(b)
}

func CopyDigits(filename string) []byte {
	b, _ := ioutil.ReadFile(filename)
	b = digitRegexp.Find(b)
	c := make([]byte, len(b))
	copy(c, b)
	return c
}
