package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

/*
 * Complete the 'findLongestSubsequence' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts INTEGER_ARRAY arr as parameter.
 */

func findLongestSubsequence(arr []int32) int32 {
	var map_str = make(map[int32]int32)
	var str_len = int32(len(arr))
	var i int32
	for i = 0; i < str_len; i++{
		map_str[arr[i]] = i
	}
	for i = 0; i < str_len; i++{
		if map_str[arr[i]] == i{
			return str_len
		}
	}
	return -1
}
func main() {
    reader := bufio.NewReaderSize(os.Stdin, 16 * 1024 * 1024)
    stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
    checkError(err)
    defer stdout.Close()
    writer := bufio.NewWriterSize(stdout, 16 * 1024 * 1024)
    arrCount, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
    checkError(err)
    var arr []int32
    for i := 0; i < int(arrCount); i++ {
        arrItemTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
        checkError(err)
        arrItem := int32(arrItemTemp)
        arr = append(arr, arrItem)
    }
    result := findLongestSubsequence(arr)
    fmt.Fprintf(writer, "%d\n", result)
    writer.Flush()
}

func readLine(reader *bufio.Reader) string {
    str, _, err := reader.ReadLine()
    if err == io.EOF {
        return ""
    }
    return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
    if err != nil {
        panic(err)
    }
}