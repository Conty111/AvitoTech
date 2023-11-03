package main

import (
	"fmt"

	"github.com/Conty111/AvitoTech/storage"
)

func main() {
	db := storage.New("localhost", "somethinglongpassword123", 6379, 16)
	db.Connect()
	db.SetValue("key1", "1234")
	fmt.Println(db.GetValue("key1"))
	db.Disconnect()
}
