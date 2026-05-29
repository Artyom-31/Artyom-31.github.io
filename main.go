package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	now := time.Now().Format("15:04:05")
	content := "<h1 style=\"color:#8B0000;font-size:100px\">X</h1><p>" + now + "</p>"
	os.WriteFile("index.html", []byte(content), 0644)
	fmt.Println("done")
}
