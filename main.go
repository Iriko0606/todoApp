package main

import (
	"fmt"
	"main/app/models"
)

func main() {
	// fmt.Println(config.Config.Port)
	// fmt.Println(config.Config.SQLDriver)
	// fmt.Println(config.Config.DbName)
	// fmt.Println(config.Config.LogFile)

	// log.Println("test")

	// fmt.Println(models.Db)

	// u := &models.User{}
	// u.Name = "test"
	// u.Email = "test@example.com"
	// u.PassWord = "pass"
	// fmt.Println(u)

	// u.CreateUser()

	u, _ := models.GetUser(1)
	// fmt.Println(u)

	u.Name = "test3"
	u.Email = "test3@test.com"
	u.UpdateUser()
	u, _ = models.GetUser(1)

	fmt.Println(u)

	u.DeleteUser()
	u, _ = models.GetUser(1)
	fmt.Println(u)
}
