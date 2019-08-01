package Update

import (
	"LearnOpDB/Config"
	"LearnOpDB/Model"
	"fmt"
	"github.com/jinzhu/gorm"
)

func UpdateAllFields()  {
	db := Config.DB

	// 使用Save时会更新所有的域，无论是否修改都会更新
	var user Model.User
	db.First(&user) // SELECT * FROM `users`  WHERE `users`.`deleted_at` IS NULL ORDER BY `users`.`id` ASC LIMIT
	user.Name = "jinzhu 2"
	user.Age = 22
	db.Debug().Save(&user) // UPDATE `users` SET `created_at` = '2019-07-29 15:53:10', `updated_at` = '2019-07-30 16:58:36', `deleted_at` = NULL, `name` = 'jinzhu 2', `age` = 22, `birthday` = '2019-07-29 15:53:10', `email` = '', `role` = 'stu', `member_number` = NULL, `num` = 0, `address` = ''  WHERE `users`.`deleted_at` IS NULL AND `users`.`id` = 1
}

func UpdateChangedFields()  {
	db := Config.DB

	//Update/Updates ---- 只更新改变了的字段，Update是更新单个字段，Updates是更新多个字段

	var user Model.User
	db.First(&user)
	db.Model(&user).Update("name", "hello")
	// UPDATE `users` SET `name` = 'hello', `updated_at` = '2019-07-30 17:15:28'  WHERE `users`.`deleted_at` IS NULL AND `users`.`id` = 1
	fmt.Println("after update:")
	fmt.Println(user) // 1 hello 22

	db.Model(&user).Where("age > ?", 18).Update("name", "hello")
	// UPDATE `users` SET `name` = 'hello', `updated_at` = '2019-07-30 17:15:28'  WHERE `users`.`deleted_at` IS NULL AND `users`.`id` = 1 AND ((age > 18))
	fmt.Println("update with conditions:")
	fmt.Println(user) // 1 hello 22

	// 使用map更新多个字段时----只会更新改变了的字段
	db.Model(&user).Updates(map[string]interface{}{"name":"hello", "age":20})
	fmt.Println("updates by map:")
	fmt.Println(user) // 1 hello 20

	// 使用struct更新多个字段时----只会更新改变了的非零值字段
	db.Model(&user).Updates(Model.User{Name:"hello", Age:21})
	fmt.Println("updates by struct:")
	fmt.Println(user) // 1 hello 21

	db.Model(&user).Updates(Model.User{Name:"", Age:0})
	fmt.Println("updates by struct with zero-value:")
	fmt.Println(user) // 1 hello 21
}

func UpdateSelectedFields()  {
	db := Config.DB

	// Select/Omit ---- 更新或者忽略指定字段
	var user Model.User
	db.First(&user)

	// select ---- 更新指定字段
	db.Model(&user).Select("name").Updates(map[string]interface{}{"name":"hello!", "age":30})
	// UPDATE `users` SET `name` = 'hello!', `updated_at` = '2019-07-30 17:23:33'  WHERE `users`.`deleted_at` IS NULL AND `users`.`id` = 1
	fmt.Println("select:")
	fmt.Println(user) // 1 hello! 21

	// omit ---- 去掉指定字段，更新剩余字段
	db.Model(&user).Omit("name").Updates(map[string]interface{}{"name":"hello!", "age":32})
	// UPDATE `users` SET `age` = 32, `updated_at` = '2019-07-30 17:23:33'  WHERE `users`.`deleted_at` IS NULL AND `users`.`id` = 1
	fmt.Println("omit:")
	fmt.Println(user) // 1 hello! 32
}

func UpdateColumns()  {
	db := Config.DB

	// 上面的更新操作都会执行模型的BeforeUpdate, AfterUpdate方法来更新记录的UpdatedAt时间戳
	// 使用UpdateColumn, UpdateColumns时，就不会调用BeforeUpdate, AfterUpdate方法
	var user Model.User
	db.First(&user)

	// 更新单个字段，类似update
	db.Model(&user).UpdateColumn("name", "hello")
	// UPDATE `users` SET `name` = 'hello'  WHERE `users`.`deleted_at` IS NULL AND `users`.`id` = 1
	fmt.Println("update column:")
	fmt.Println(user) // 1 hello 32

	// 更新多个字段，类似updates
	db.Model(&user).UpdateColumns(Model.User{Name:"hello", Age:40})
	// UPDATE `users` SET `age` = 40, `name` = 'hello'  WHERE `users`.`deleted_at` IS NULL AND `users`.`id` = 1
	fmt.Println("update columns:")
	fmt.Println(user) // 1 hello 40
}

func BatchUpdates()  {
	db := Config.DB

	// 官方文档中说：当做批更新时不会执行Hooks(也就是BeforeUpdate, AfterUpdate等方法)
	// 实际：还是会更新updated_at字段

	db.Table("users").Where("id IN (?)", []int{10, 11}).Updates(map[string]interface{}{"role": "new_role", "age": 18})
	// UPDATE `users` SET `age` = 18, `role` = 'new_role'  WHERE (id IN (10,11))

	// struct只会更新改变了的非零值，map是更新改变了的字段
	db.Model(Model.User{}).Updates(Model.User{Name: "", Age: 33})
	// UPDATE `users` SET `age` = 33, `updated_at` = '2019-07-30 17:39:05'  WHERE `users`.`deleted_at` IS NULL

	// RowsAffected ---- 获取更新的记录数
	count := db.Model(Model.User{}).Updates(Model.User{Role: "hellohello", Age: 0}).RowsAffected
	// UPDATE `users` SET `role` = 'hellohello', `updated_at` = '2019-07-30 17:40:50'  WHERE `users`.`deleted_at` IS NULL
	fmt.Println(count) // 13
}

func UpdateWithSQLExpr()  {
	db := Config.DB

	var user Model.User
	db.First(&user)
	fmt.Println(user) // 1 hello 33

	db.Model(&user).Update("age", gorm.Expr("age * ? + ?", 2, 10))
	// UPDATE `users` SET `age` = age * 2 + 10, `updated_at` = '2019-07-30 18:01:45'  WHERE `users`.`deleted_at` IS NULL AND `users`.`id` = 1
	db.Where("id = ?", user.ID).First(&user)
	fmt.Println(user) // 1 hello 76

	db.Model(&user).Updates(map[string]interface{}{"age": gorm.Expr("age * ? + ?", 2, 10)})
	// UPDATE `users` SET `age` = age * 2 + 10, `updated_at` = '2019-07-30 18:01:45'  WHERE `users`.`deleted_at` IS NULL AND `users`.`id` = 1
	db.Where("id = ?", user.ID).First(&user)
	fmt.Println(user) // 1 hello 162

	db.Model(&user).UpdateColumn("age", gorm.Expr("age - ?", 10))
	// UPDATE `users` SET `age` = age - 10  WHERE `users`.`deleted_at` IS NULL AND `users`.`id` = 1
	db.Where("id = ?", user.ID).First(&user)
	fmt.Println(user) // 1 hello 152

	db.Model(&user).Where("age > 50").UpdateColumn("age", gorm.Expr("age - ?", 1))
	// UPDATE `users` SET `age` = age - 1  WHERE `users`.`deleted_at` IS NULL AND ((age > 50))
	db.Where("id = ?", user.ID).First(&user)
	fmt.Println(user) // 1 hello 151
}