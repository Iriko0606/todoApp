package controllers

import (
	"log"
	"net/http"
	"todo_app_heroku/app/models"
)

func top(w http.ResponseWriter, r *http.Request) {
	_, err := session(w, r)
	if err != nil {
		generateHTML(w, "Hello", "layout", "public_navbar", "top")
	} else {
		http.Redirect(w, r, "/todos", 302)
	}

}

func index(w http.ResponseWriter, r *http.Request) {
	session, err := session(w, r)
	if err != nil {
		http.Redirect(w, r, "/", 302)
	} else {
		user, err := session.GetUserBySession()
		if err != nil {
			log.Println(err)
		}
		todos, _ := user.GetTodosByUser()
		user.Todos = todos
		generateHTML(w, user, "layout", "private_navbar", "index")
	}
}

func todoEdit(w http.ResponseWriter, r *http.Request, id int) {
	session, err := session(w, r)
	if err != nil {
		http.Redirect(w, r, "/login", 302)
	} else {
		_, err := session.GetUserBySession()
		if err != nil {
			log.Println(err)
		}
		t, err := models.GetTodo(id)
		if err != nil {
			log.Println(err)
		}
		generateHTML(w, t, "layout", "private_navbar", "todo_edit")
	}
}

func todoUpdate(w http.ResponseWriter, r *http.Request, id int) {
	session, err := session(w, r)
	if err != nil {
		http.Redirect(w, r, "/login", 302)
	} else {
		if err != nil {
			log.Println(err)
		}
		user, err := session.GetUserBySession()
		if err != nil {
			log.Println(err)
		}
		content := r.PostFormValue("content")

		t := &models.Todo{ID: id, Content: content, UserID: user.ID}
		if err := t.UpdateTodo(); err != nil {
			log.Println(err)
		}
		http.Redirect(w, r, "/login", 302)
	}
}
