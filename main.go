package main

import (
	"fmt"

	"github.com/daichi-sato-design/todoApp_go_sqlite/app/models"
)

func main(){
fmt.Println(models.Db)

	u := &models.User{}
	u.Name = "test"
	u.Email = "test@example.com"
	u.Password = "testtest"
	fmt.Println(u)

	u.CreateUser()
}
