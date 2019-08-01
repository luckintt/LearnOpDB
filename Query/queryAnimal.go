package Query

import (
	"LearnOpDB/Config"
	"LearnOpDB/Model"
	"fmt"
)

func FirstOrCreateQueryAnimal() {
	// FirstOrCreate只有struct/map时可用——获取匹配的记录，或者根据给定条件创建一条新纪录，并存储到数据库中
	db := Config.DB

	// 没有找到时会创建一条记录，更新数据库，并且将对应的记录中的字段赋值给结构体
	var unfoundAnimal Model.Animal
	db.Where(Model.Animal{Name:"non_existing"}).FirstOrCreate(&unfoundAnimal)
	fmt.Println("unfound:")
	fmt.Println(unfoundAnimal) // unfoundAnimal -> Animal{Id:"f0e00e04-835d-4158-89e5-f645d3b9f7ec", Name:"non_existing", Age:0}

	// 找到时会直接将该记录赋值给结构体
	var foundAnimal Model.Animal
	db.Where(Model.Animal{Name:"dog"}).FirstOrCreate(&foundAnimal)
	fmt.Println("found")
	fmt.Println(foundAnimal) // foundAnimal -> Animal{Id:"cac03423-d31c-45c4-bbbe-73deffc2c06a", Name:"dog", Age:1}

	// attrs----没有找到记录时会利用参数创建一个记录，更新数据库，并将该记录赋值给结构体；找到记录时直接将该记录返回
	// SELECT * FROM animals WHERE name = 'non_existing1';
	// INSERT INTO "animals" (name, age) VALUES ("non_existing1", 2);
	var unfoundAttrsAnimal Model.Animal
	db.Where(Model.Animal{Name:"non_existing1"}).Attrs(Model.Animal{Age:2}).FirstOrCreate(&unfoundAttrsAnimal)
	fmt.Println("unfound attrs:")
	fmt.Println(unfoundAttrsAnimal) // unfoundAttrsAnimals -> Animal{Id:"b112b4dc-72cc-47cd-81ba-0b5a9aa467d5", Name:"non_existing1", Age:2}

	// SELECT * FROM animals WHERE name = 'dog';
	var foundAttrsAnimal Model.Animal
	db.Where(Model.Animal{Name:"dog"}).Attrs(Model.Animal{Age:2}).FirstOrCreate(&foundAttrsAnimal)
	fmt.Println("found attrs:")
	fmt.Println(foundAttrsAnimal) // foundAttrsAnimals -> Animal{Id:"cac03423-d31c-45c4-bbbe-73deffc2c06a", Name:"dog", Age:1}

	// assign----无论记录是否找到，都使用参数更新记录的值（没有找到时利用参数创建记录，更新数据库； 找到时利用参数更新记录再返回记录赋值给结构体）
	// SELECT * FROM animals WHERE name = 'non_existing2';
	// INSERT INTO "animals" (name, age) VALUES ("non_existing2", 3);
	var unfoundAssignAnimal Model.Animal
	db.Where(Model.Animal{Name:"non_existing2"}).Assign(Model.Animal{Age:3}).FirstOrCreate(&unfoundAssignAnimal)
	fmt.Println("unfound assign:")
	fmt.Println(unfoundAssignAnimal) // unfoundAssignAnimal -> Animal{Id:"fc936e3b-86e2-4bf5-9a68-ac209c7bcee1", Name:"non_existing2", Age:3}

	// SELECT * FROM animals WHERE name = 'dog';
	// UPDATE animals SET age=3 WHERE id = "cac03423-d31c-45c4-bbbe-73deffc2c06a";
	var foundAssignAnimal Model.Animal
	db.Where(Model.Animal{Name:"dog"}).Assign(Model.Animal{Age:3}).FirstOrCreate(&foundAssignAnimal)
	fmt.Println("found assign:")
	fmt.Println(foundAssignAnimal) // foundAssignAnimal -> Animal{Id:"cac03423-d31c-45c4-bbbe-73deffc2c06a", Name:"dog", Age:3}
}
