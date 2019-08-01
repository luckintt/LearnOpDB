package Create

import (
	"LearnOpDB/Config"
	"LearnOpDB/Model"
	"fmt"
)

func CreateAnimalTable()  {
	if !Config.DB.HasTable("animals") {
		fmt.Println("not exist")
		// 指定表名并用结构体User来创建该表
		Config.DB.Table("animals").CreateTable(&Model.Animal{})
	} else {
		fmt.Println("exist")
	}
}

func CreateAnimalRecord(animal *Model.Animal)  {
	// NewRecord----判断记录是否存在
	// 初始时记录不存在，所以返回值时true
	if Config.DB.NewRecord(animal){
		Config.DB.Create(&animal)
		// 创建来该记录来以后，会返回false
		fmt.Println(Config.DB.NewRecord(animal))
	}
}
