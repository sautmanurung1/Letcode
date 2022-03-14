package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"io"
)

type Manager struct {
    FullName string
	Position string
	Age      int32
	YearsInCompany int32
}

func EncodeManager(manager *Manager) (io.Reader, error) {
	// make output like this : {"full_name":"John Doe","position":"CEO","age":42,"years_in_company":10}


	// create a buffer to write to
	var buffer bytes.Buffer

	// create a writer to write to the buffer

	writer := bufio.NewWriter(&buffer)

	// create a json encoder to write to the writer
	encoder := json.NewEncoder(writer)
	
}


func main(){
}