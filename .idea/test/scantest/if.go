package main

import "fmt"

func test_method() int {
	return 2
}

func UploadFile() (fileName string, success bool) {

	fileName = "test_upload"
	success = true
	return
}

func main() {
	if test_method() > 3 {
		fmt.Println("test")
	} else {
		fmt.Println("test fail")
	}

	if file, success := UploadFile(); success {
		if success == true {
			fmt.Println(file)
		} else {
			fmt.Println("fail")
		}
	}
}
