package main

import (
	"fmt"

	"github.com/Conty111/AvitoTech/storage"
)

func main() {
	rdb := storage.NewRedis("localhost:6379", "", 0)
	rdb.Connect()
	rdb.SetValue("key", "123")
	fmt.Println(rdb.GetValue("key"))
}
