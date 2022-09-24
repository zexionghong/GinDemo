package main

import (
	"fmt"
	"ipData/core"
	"ipData/models"
)

func ss(d *models.CoreUsr) {

}

func hello() {

	//var users []models.CoreUsr
	var count int64
	var users = []models.CoreUsr{
		{Oid: "name1"},
		{Oid: "name2"},
		{Oid: "name3"},
		{Oid: "name3"},
	}

	core.DB.Model(&models.CoreUsr{}).Group("name").Count(&count)
	//ss := db.Model(&models.CoreUsr{}).Where("createdon>?", "2020-01-01").Group("oid").Count()
	fmt.Println(count)
	for _, user := range users {
		fmt.Println(user.GetString())

	}

}

func main() {
	core.StartApp()

}
