package main

import (
	"fmt"

	uuid "github.com/satori/go.uuid"
)

func main() {
	uid := uuid.Must(uuid.NewV4()).String()
	fmt.Println(uid)

}
