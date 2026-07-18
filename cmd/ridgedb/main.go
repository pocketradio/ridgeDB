package main

import (
	"fmt"
	"ridgeDB/internal/storage"
)

func main() {
	db := storage.NewStore()
	fmt.Println(db)
}
