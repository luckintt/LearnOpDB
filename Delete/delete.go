package Delete

import (
	"LearnOpDB/Config"
	"LearnOpDB/Model"
	"fmt"
)

func DeleteRecord()  {
	db := Config.DB

	// 在删除记录时，需要确保主键有值，GORM使用主键去删除记录
	// 当主键为空时，会删除模型中的所有记录
	var user Model.User
	db.First(&user)

	// 删除一个存在的记录
    db.Delete(&user)
	// UPDATE `users` SET `deleted_at`='2019-08-01 14:09:50'  WHERE `users`.`deleted_at` IS NULL AND `users`.`id` = 1
}

func BatchDeleteRecord()  {
	db := Config.DB

	db.Debug().Where(" name like ?", "%jinzhu1%").Delete(Model.User{})
	//  UPDATE `users` SET `deleted_at`='2019-08-01 14:26:14'  WHERE `users`.`deleted_at` IS NULL AND (( name like '%jinzhu1%'))

	db.Debug().Delete(Model.User{}, " name like ?", "%jinzhu1%")
	// UPDATE `users` SET `deleted_at`='2019-08-01 14:26:14'  WHERE `users`.`deleted_at` IS NULL AND (( name like '%jinzhu1%'))
}

func SoftDelete()  {
	db := Config.DB

	// 当一个modle有DeletedAt域时，默认会使用软删除：也就是不从数据库中删除记录，只是将DeletedAt置为当前时间
	// 在查询时会忽略DeletedAt域不为空的记录，也就是被软删除的记录不会被查询出来
	var users []Model.User
	db.Where("name like ?", "%jinzhu1%").Find(&users)
	// SELECT * FROM `users`  WHERE `users`.`deleted_at` IS NULL AND ((name like '%jinzhu1%'))
	fmt.Println(users)  // []

	// Unscoped:忽略DeletedAt域，把符合条件的被软删除的记录也查询出来
	var unscopedUsers []Model.User
	db.Unscoped().Where("name like ?", "%jinzhu1%").Find(&unscopedUsers)
	// SELECT * FROM `users`  WHERE (name like '%jinzhu1%')
	fmt.Println(unscopedUsers)
}

func PermanentlyDelete()  {
	db := Config.DB
	// Unscoped:从数据库中永久删除记录
	var user Model.User
	db.First(&user)
	db.Debug().Unscoped().Delete(&user)
	//  DELETE FROM `users`  WHERE `users`.`id` = 3
}