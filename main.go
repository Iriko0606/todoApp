package main

import "main/app/models"

func main() {

	// fmt.Println(models.Db)
	// fmt.Println(config.Config.Port)
	// fmt.Println(config.Config.SQLDriver)
	// fmt.Println(config.Config.DbName)
	// fmt.Println(config.Config.LogFile)

	// log.Println("test")

	// fmt.Println(models.Db)

	// u := &models.User{}
	// u.Name = "test4"
	// u.Email = "test4@example.com"
	// u.PassWord = "pass4"
	// fmt.Println(u)

	// u.CreateUser()

	// // fmt.Println(u)

	// u.Name = "test3"
	// u.Email = "test3@test.com"
	// u.UpdateUser()

	// u.DeleteUser()
	// u, _ := models.GetUser(1)
	// user, _ := models.GetUser(4)
	// user.CreateTodo("First Todo")

	// user, _ := models.GetUser(1)
	// user.CreateTodo("Todo")

	// todos, _ := models.GetTodos()
	// for _, v := range todos {
	// 	fmt.Println(v)
	// }

	// user2, _ := models.GetUser(4)
	// todos, _ := user2.GetTodosByUser()

	// for _, v := range todos {
	// 	fmt.Println(v)
	// }
	t, _ := models.GetTodo(3)
	t.DeleteTodo()
	// t.UpdateTodo()

}
