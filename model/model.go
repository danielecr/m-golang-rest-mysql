package model

import (
	"fmt"

	"github.com/danielecr/nsa-golang-rest-mysql/config"
)

func Index() string {
	return "Welcome!"
}

func UserById(id int) string {

	rows, err := config.DB.Query("SELECT * FROM user WHERE id = ?", id)
	if err != nil {
		fmt.Println(err)
	}
	defer rows.Close()

	var name string
	for rows.Next() {
		err := rows.Scan(&name)
		if err != nil {
			fmt.Println(err)
		}
	}

	return name
}

func TodosByUserId(id int) interface{} {

	rows, err := config.DB.Query("SELECT * FROM todo WHERE user_id = ?", id)
	if err != nil {
		fmt.Println(err)
	}
	defer rows.Close()

	var todos []interface{}
	for rows.Next() {
		var id int
		var name string
		err := rows.Scan(&id, &name)
		if err != nil {
			fmt.Println(err)
		}
		todo := map[string]interface{}{"id": id, "name": name}
		todos = append(todos, todo)
	}
	return todos
}
