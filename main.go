package main

import (
	"LearnOpDB/Config"
	"LearnOpDB/Delete"
)

func main()  {
	Config.Conn()

	//Create.CreateUserTable()
	//user := Model.User{Name:"jinzhu1111", Birthday:time.Now()}
	//Create.CreateUserRecord(&user)

	//Create.CreateAnimalTable()
	//animal := Model.Animal{Name:"dog1", Age:1}
	//Create.CreateAnimalRecord(&animal)



	//Query.DirectQueryUser()
	//Query.WhereQueryUser()
	//Query.StructAndMapQueryUser()
	//Query.NotQueryUser()
	//Query.OrQueryUser()
	//Query.FirstOrInitQueryUser()
	//Query.FirstOrCreateQueryAnimal()
	//Query.SubQuery()
	//Query.SelectQuery()
	//Query.OrderQuery()
	//Query.LimitQuery()
	//Query.OffsetQuery()
	//Query.CountQuery()
	//Query.GroupANDHavingQuery()
	//Query.JoinQuery()
	//Query.PluckQuery()
	//Query.ScanQuery()



	//Update.UpdateAllFields()
	//Update.UpdateChangedFields()
	//Update.UpdateSelectedFields()
	//Update.UpdateColumns()
	//Update.BatchUpdates()
	//Update.UpdateWithSQLExpr()



	//Delete.DeleteRecord()
	//Delete.BatchDeleteRecord()
	//Delete.SoftDelete()
	Delete.PermanentlyDelete()

	Config.DB.Close()
}
