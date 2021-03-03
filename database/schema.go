package database

import "gihub.com/bnsngltn/go-fiber-blog-app/model"

//CreateTables func
func CreateTables() {
	if !DB.Migrator().HasTable(&model.User{}) {
		DB.Migrator().CreateTable(&model.User{})
	}

	if !DB.Migrator().HasTable(&model.Post{}) {
		DB.Migrator().CreateTable(&model.Post{})
	}
}
