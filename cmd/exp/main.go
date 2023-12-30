package main

import (
	"html/template"
	"os"
)

type User struct {
	Name string
	Bio  string
	Age  int
	Meta UserMeta
}
type UserMeta struct {
	Visits int
}

func main() {
	t, err := template.ParseFiles("hello.gohtml")
	if err != nil {
		panic(err)
	}
	user := User{
		Name: "Joshua Mawuli Agbeku",
		Bio:  `<script>alert(haha, you have been h4x0r3d!);</script>`,
		Age:  33,
		Meta: UserMeta{
			Visits: 5,
		},
	}
	// OR
	// user := struct {
	// 	Name string
	// }{
	// 	Name: "Joshua Agbeku",
	// }

	err = t.Execute(os.Stdout, user)
	if err != nil {
		panic(err)
	}
}
