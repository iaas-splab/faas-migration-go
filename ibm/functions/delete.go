package main

import (
	"encoding/json"
	"fmt"
	"github.com/c-mueller/faas-migration-go/core"
	"github.com/c-mueller/faas-migration-go/ibm"
	"os"
)

func main() {
	arg := os.Args[1]

	fmt.Println(arg)
	var obj core.IDRequest
	json.Unmarshal([]byte(arg), &obj)

	repo := ibm.NewCloudantRepository()

	err := core.Delete(obj,repo)
	if err != nil {
		fmt.Printf("Execution Failed: Error %s\n",err.Error())
		return
	}
	fmt.Println("{}")
}