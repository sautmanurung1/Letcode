package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

/*
 * Complete the 'getAnagramPeriod' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts STRING input_str as parameter.
 */

func getAnagramPeriod(input_str string) int32 {
    var map_str = make(map[string]int32)
    var str_len = int32(len(input_str))
    var i int32
    for i = 0; i < str_len; i++{
        map_str[input_str[i:i+1]] = i
    }
    for i = 0; i < str_len; i++{
        if map_str[input_str[i:i+1]] == i{
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
    input_str := readLine(reader)
    result := getAnagramPeriod(input_str)
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
