// 设置全局变量
package initapp

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

// set environment variable
func SetEnv() {
	// check if .env exit
	if _, err := os.Stat(".env"); os.IsNotExist(err) {
		return
	}

	file, err := ioutil.ReadFile(".env")
	if err != nil {
		panic(err)
	}
	lines := strings.Split(string(file), "\n")
	for _, line := range lines {
		pair := strings.Split(line, "=")
		if len(pair) != 2 {
			continue
		}
		os.Setenv(pair[0], pair[1])
	}

	fmt.Println(os.Getenv("HOST"))
}
