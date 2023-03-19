package main

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
)

/**
 * Структура для задачи
 *
 */
type Task struct {
	Id          string
	Title       string `validate:"required"`
	Description string `validate:"required"`
	Date        string `validate:"required"`
}

/**
 * Отображение страницы со списком задач
 *
 * w http.ResponseWriter
 * r http.Request
 */
func viewTasks(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("assets/templates/wrapper.tmpl", "assets/templates/pages/tasks.tmpl")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	db := initDateBase()
	allTasks, err := db.Query("SELECT * from `tasks`")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	defer allTasks.Close()

	tasks := []Task{}
	for allTasks.Next() {
		var task Task
		err := allTasks.Scan(&task.Id, &task.Title, &task.Description, &task.Date)
		if err != nil {
			panic(err)
		}
		tasks = append(tasks, task)
	}

	t.ExecuteTemplate(w, "wrapper", tasks)
}

/**
 * Отображение страницы добавления задачи
 *
 * w http.ResponseWriter
 * r http.Request
 */
func addTask(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("assets/templates/wrapper.tmpl", "assets/templates/pages/add.tmpl")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	t.ExecuteTemplate(w, "wrapper", nil)
}

/**
 * Добавление задачи
 *
 * w http.ResponseWriter
 * r http.Request
 */
func postTask(w http.ResponseWriter, r *http.Request) {
	title := r.FormValue("title")
	description := r.FormValue("description")
	date := r.FormValue("date")

	newTask := &Task{
		Title:       title,
		Description: description,
		Date:        date,
	}
	validate := validator.New()
	err := validate.Struct(newTask)
	if err != nil {
		http.Redirect(w, r, "/add?error=empty", http.StatusSeeOther)
	} else {
		db := initDateBase()
		insert, err := db.Query(fmt.Sprintf("INSERT INTO `tasks` (`title`, `description`, `date`) VALUES ('%s', '%s', '%s')", title, description, date))
		if err != nil {
			panic(err)
		}
		defer db.Close()
		defer insert.Close()

		http.Redirect(w, r, "/?msg=add", http.StatusSeeOther)
	}
}

/**
 * Отображение страницы просмотра задачи
 *
 * w http.ResponseWriter
 * r http.Request
 */
func viewTask(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("assets/templates/wrapper.tmpl", "assets/templates/pages/view.tmpl")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	vars := mux.Vars(r)

	db := initDateBase()
	res, err := db.Query(fmt.Sprintf("SELECT * FROM `tasks` WHERE `id` =  '%s'", vars["id"]))
	if err != nil {
		panic(err)
	}
	defer db.Close()
	defer res.Close()

	for res.Next() {
		var task Task
		err := res.Scan(&task.Id, &task.Title, &task.Description, &task.Date)
		if err != nil {
			panic(err)
		}
		t.ExecuteTemplate(w, "wrapper", task)
	}
}

/**
 * Отображение страницы изменения задачи
 *
 * w http.ResponseWriter
 * r http.Request
 */
func editTask(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("assets/templates/wrapper.tmpl", "assets/templates/pages/edit.tmpl")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	vars := mux.Vars(r)

	db := initDateBase()
	res, err := db.Query(fmt.Sprintf("SELECT * FROM `tasks` WHERE `id` =  '%s'", vars["id"]))
	if err != nil {
		panic(err)
	}
	defer db.Close()
	defer res.Close()

	for res.Next() {
		var task Task
		err := res.Scan(&task.Id, &task.Title, &task.Description, &task.Date)
		if err != nil {
			panic(err)
		}
		t.ExecuteTemplate(w, "wrapper", task)
	}
}

/**
 * Сохранение изменения задачи
 *
 * w http.ResponseWriter
 * r http.Request
 */
func putTask(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id := vars["id"]
	title := r.FormValue("title")
	description := r.FormValue("description")
	date := r.FormValue("date")

	editTask := &Task{
		Id:          id,
		Title:       title,
		Description: description,
		Date:        date,
	}

	validate := validator.New()
	err := validate.Struct(editTask)
	if err != nil {
		http.Redirect(w, r, r.Header.Get("Referer")+"?msg=empty", http.StatusSeeOther)
	} else {
		db := initDateBase()

		prepaireTask, err := db.Prepare("UPDATE tasks SET title = ?, description = ?, date = ? WHERE id = ?")
		if err != nil {
			panic(err)
		}

		prepaireTask.Exec(title, description, date, id)

		http.Redirect(w, r, "/?msg=edit", http.StatusSeeOther)
	}
}

/**
 * Удаление задачи
 *
 * w http.ResponseWriter
 * r http.Request
 */
func deleteTask(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	db := initDateBase()
	task, err := db.Query(fmt.Sprintf("DELETE FROM `tasks` WHERE `id` = '%s'", vars["id"]))
	if err != nil {
		panic(err)
	}

	defer db.Close()
	defer task.Close()
}
