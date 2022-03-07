package main

func dnaComplement(s string) string {
	// first reverse string
	var r string = ""
	for i := len(s)-1; i >= 0; i-- {
		r += string(s[i])
	}
	// then complement string
	var c string = ""
	for i := 0; i < len(r); i++ {
		switch r[i] {
		case 'A':
			c += "T"
		case 'T':
			c += "A"
		case 'C':
			c += "G"
		case 'G':
			c += "C"
		}
	}
	return c
}
	// var complement string = ""
	// for _, v := range s {
	// 	switch v {
	// 	case 'A':
	// 		complement += "T"
	// 	case 'T':
	// 		complement += "A"
	// 	case 'C':
	// 		complement += "G"
	// 	case 'G':
	// 		complement += "C"
	// 	}
	// }
	// return complement

func main(){
	println(dnaComplement("ATCGTA"))
	println(dnaComplement("ACCGGGTTTT"))
}