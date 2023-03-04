package main

import "fmt"

type student struct {
	name       string
	namaEncode string
	score      int
}

type Chiper interface {
	Encode() string
	Decode() string
}

func (s *student) Encode() string {
	var namaEncode = ""
	namaEncode = SubstitusiChiper(s.name)
	return namaEncode
}

func (s *student) Decode() string {
	var namaDecode = ""
	if s.namaEncode == "" {
		s.namaEncode = s.Encode()
	}
	namaDecode = SubstitusiChiper(s.namaEncode)
	return namaDecode
}

func SubstitusiChiper(s string) string {
	HasilAkhir := []rune{}
	for _, nilai := range s {
		HasilAkhir = append(HasilAkhir, 'z'-((nilai)-'a'))
	}
	return string(HasilAkhir)
}

func main() {
	var menu int
	var a student = student{}
	var c Chiper = &a

	fmt.Print("[1] Encrypt \n[2] Decrypt \nChoose your menu? ")
	fmt.Scan(&menu)

	if menu == 1 {
		fmt.Print("\nInput Student Name: ")
		fmt.Scan(&a.name)
		fmt.Print("\nEncode of student’s name " + a.name + " is : " + c.Encode())
	} else if menu == 2 {
		fmt.Print("\nInput Student Name: ")
		fmt.Scan(&a.name)
		fmt.Print("\nDecode of student’s name " + a.name + " is : " + c.Decode())
	}
}
