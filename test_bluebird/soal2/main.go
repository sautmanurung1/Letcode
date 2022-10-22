package main

func solution(str string) bool {
	len := len(str)
	for i := 0; i < len/2; i++ {
		if str[i] != str[len-1-i] {
			return false
		}
	}
	return true
}
func main() {

}
