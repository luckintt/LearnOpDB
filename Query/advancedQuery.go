package Query

import (
	"LearnOpDB/Config"
	"LearnOpDB/Model"
	"fmt"
)

func SubQuery()  {
	db := Config.DB

	// 用 *gorm.expr 进行子查询
	expr := db.Table("users").Select("AVG(age)").Where("name like ?", "%jinzhu%").SubQuery()
	fmt.Print("expr:")
	fmt.Println(expr)

	var subUsers []Model.User
	db.Where("age >= ?", db.Table("users").Select("AVG(age)").Where("name like ?", "%jinzhu%").SubQuery()).Find(&subUsers)
	// SELECT * FROM `users`  WHERE `users`.`deleted_at` IS NULL AND ((age >= (SELECT AVG(age) FROM `users`  WHERE (name like '%jinzhu%'))));
	fmt.Println("sub query:")
	fmt.Println(subUsers)
}

func SelectQuery()  {
	db := Config.DB

	// 类似投影----指定要从数据库检索的字段，默认是选择所有的字段
	// 投影一个字段，不能直接并列选择多个字段
	var selectUsers []Model.User
	db.Select("name").Find(&selectUsers) // SELECT name FROM users;
	fmt.Println("select:")
	fmt.Println(selectUsers) // 0 jinzhu 0 : 其实所有字段都存在，但是只有Name字段为真实值，其他的都是默认值

	// 投影多个字段
	var selectSliceUsers []Model.User
	db.Select([]string{"name", "age"}).Find(&selectSliceUsers) // SELECT name, age FROM users;
	fmt.Println("select slice:")
	fmt.Println(selectSliceUsers) // 0 jinzhu 18 : 其实所有字段都存在，但是只有Name和Age字段为真实值，其他的都是默认值

	// 选出的结果是空，跟mysql里面不一样
	var coalesceUsers []Model.User
	db.Table("users").Select("COALESCE('age', ?)", 18).Rows() // SELECT COALESCE(age, 18) FROM `users`;
	fmt.Println("coalesce:")
	fmt.Println(coalesceUsers)

}

func OrderQuery()  {
	db := Config.DB

	// Order指定查询数据时的排序方式，当第二个参数reorder置为true时会根据指定条件重新排序

	var multiOrderAnimals []Model.Animal
	db.Order("age desc, name").Find(&multiOrderAnimals) // 按age降序，name升序排序
	fmt.Println("multi order:")
	fmt.Println(multiOrderAnimals)

	var multiSingleOrderAnimals []Model.Animal
	db.Order("age desc").Order("name").Find(&multiSingleOrderAnimals) // 按age降序，name升序排序
	fmt.Println("multi single order:")
	fmt.Println(multiSingleOrderAnimals)

	// 第二个参数置为true时根据指定条件重新排序
	var reorderAnimals1, reorderAnimals2 []Model.Animal
	db.Order("age desc").Find(&reorderAnimals1).Order("age", true).Find(&reorderAnimals2)
	fmt.Println("reorder:")
	fmt.Println(reorderAnimals1) // 按age降序   SELECT * FROM `animals`   ORDER BY age desc
	fmt.Println(reorderAnimals2) // 按age升序   SELECT * FROM `animals`   ORDER BY `age`
}

func LimitQuery()  {
	db := Config.DB

	// 指定筛选记录的最大条数
	var limitAnimals []Model.Animal
	db.Limit(3).Find(&limitAnimals)
	fmt.Println("limit")
	fmt.Println(limitAnimals)

	// 当limit的值为-1时，表示取出所有记录
	var multiLimitUsers1, multiLimitUsers2 []Model.User
	db.Limit(10).Find(&multiLimitUsers1).Limit(-1).Find(&multiLimitUsers2)
	fmt.Println("multi limit:")
	fmt.Println(multiLimitUsers1) // SELECT * FROM users LIMIT 10;
	fmt.Println(multiLimitUsers2) // SELECT * FROM users;
}

func OffsetQuery()  {
	db := Config.DB

	// Offset要与Limit一起使用，单独用无效,⚠️并且需要保证limit的条数不大于跳过指定条数以后剩下的条数------指定跳过的记录数
	var offsetAnimals []Model.Animal
	db.Limit(2).Offset(2).Find(&offsetAnimals)
	fmt.Println("offset:")
	fmt.Println(offsetAnimals) // SELECT * FROM `animals`   LIMIT 2 OFFSET 2

	// 当offset的值为-1时，表示不跳过任何记录
	var cancelOffsetAnimals1, cancelOffsetAnimals2 []Model.Animal
	db.Limit(2).Offset(3).Find(&cancelOffsetAnimals1).Offset(-1).Find(&cancelOffsetAnimals2)
	fmt.Println("multi offset:")
	fmt.Println(cancelOffsetAnimals1) // SELECT * FROM `animals`   LIMIT 2 OFFSET 3
	fmt.Println(cancelOffsetAnimals2) // SELECT * FROM `animals`   LIMIT 2
}

func CountQuery()  {
	db := Config.DB

	// 统计查询出的记录条数，⚠️Count要放在查询语句的最后，因为它会重写select的字段
	var whereCountUser int
	db.Where("name = ?", "jinzhu").Or("name = ?", "jinzhuu").Find(&[]Model.User{}).Count(&whereCountUser)
	fmt.Print("count by where:")
	fmt.Println(whereCountUser) // 2

	var modelCountUser int
	db.Model(&[]Model.User{}).Where("name = ?", "jinzhu").Count(&modelCountUser)
	fmt.Print("count by model:")
	fmt.Println(modelCountUser) // 1

	var tableCountUser int
	db.Table("users").Count(&tableCountUser)
	fmt.Print("count by table:")
	fmt.Println(tableCountUser) // 13
}

func GroupANDHavingQuery()  {
	db := Config.DB

	// ⚠️⚠️⚠️⚠️⚠️⚠️⚠️⚠️
	// sql: expected 2 destination arguments in Scan, not 1
	// struct出错原因：1. scan的参数个数必须与投影的结果一致 2. 结构体必须声明一个对象，不能直接定义一个空结构 3. 结构体的属性没有加&
	// map必须按照下面的方式来才行
	// https://kylewbanks.com/blog/query-result-to-map-in-golang

	// SELECT role, sum(age) as totalage FROM `users`   GROUP BY role
	rows, _ := db.Table("users").Select("role, sum(age) as totalage").Group("role").Rows()
	fmt.Println("group struct:")
	cols, _ := rows.Columns()
	fmt.Println(cols)
	for rows.Next() {
		type Result struct {
			Role string
			TotalAge int
		}
		var groupResult Result
		rows.Scan(&groupResult.Role, &groupResult.TotalAge)
		fmt.Println(groupResult.Role, groupResult.TotalAge)
	}

	rows, _ = db.Table("users").Select("role, sum(age) as totalAge").Group("role").Having("sum(age) > ?", 50).Rows()
	cols, _ = rows.Columns() // 获取列名
	fmt.Println("having map:")
	for rows.Next() {
		columns := make([]interface{}, len(cols)) // 用来存储每一列的值
		columnPointers := make([]interface{}, len(cols)) // 指向每一列的指针
		for i, _ := range columns {
			columnPointers[i] = &columns[i]
		}

		// 将查询出的结果扫描到指针数组中
		if err := rows.Scan(columnPointers...); err != nil {
			fmt.Println(err)
			return
		}

		havingResult := make(map[string]interface{})
		for i, colName := range cols {
			val := columnPointers[i].(*interface{})
			havingResult[colName] = *val
		}

		fmt.Printf("%s:%s  %s:%s\n","role", havingResult["role"], "totalAge", havingResult["totalAge"])
	}

	// ⚠️这里只能用投影的字段的首字母的大写形式，否则扫描不出来， 并且别名中不能有大写字母
	type Result struct {
		Role  string
		Total int
	}
	var results []Result
	db.Table("users").Select("role, sum(age) as total").Group("role").Having("sum(age) > ?", 50).Scan(&results)
	fmt.Println("scan:") // SELECT role, sum(age) as totalAge FROM `users`   GROUP BY role HAVING (sum(age) > 50)
	fmt.Println(results) // [{stu 126} {tea 72}]
}

func JoinQuery()  {
	db := Config.DB

	type Result struct {
		MasterName string
		AnimalName string
	}

	// 指定连接条件
	rows, _ := db.Debug().Table("users").Select("users.name as master, animals.name as animal").Joins("left join animals on animals.master_id = users.id").Rows()
	fmt.Println("join:")
	for rows.Next() {
		var joinResult Result
		rows.Scan(&joinResult.MasterName, &joinResult.AnimalName)
		fmt.Println(joinResult)
	}

	var joinResults []Result
	db.Table("users").Select("users.name as master_name, animals.name as animal_name").Joins("left join animals on animals.master_id = users.id").Scan(&joinResults)
	fmt.Println("join results:")
	fmt.Println(joinResults)

	// multiple joins with parameter
	var multiJoinUsers []Model.User
	db.Joins("JOIN animals ON animals.master_id = users.id AND animals.name like ?", "%dog%").Joins("JOIN teachers ON teachers.stu_id = users.id").Where("teachers.name in (?)", []string{"bbb", "ddd"}).Find(&multiJoinUsers)
	fmt.Println("multi join:")
	fmt.Println(multiJoinUsers)
}

func PluckQuery() {
	db := Config.DB

	// 查询出数据库中的一列字段

	var ages []int
	db.Find(&[]Model.User{}).Pluck("age", &ages) // SELECT age FROM `users`  WHERE `users`.`deleted_at` IS NULL
	fmt.Println("model in find:")
	fmt.Println(ages) // [18 18 18 18 18 18 18 18 18 18 18 18 18]

	var names []string
	db.Model(Model.User{}).Pluck("name", &names)
	fmt.Println("use model directly:")
	fmt.Println(names)

	names = nil
	db.Table("users").Pluck("name", &names)
	fmt.Println("table:")
	fmt.Println(names)
}

func ScanQuery()  {
	db := Config.DB

	// 将结果扫描到结构体中
	type Result struct {
		Name string
		Age  int
	}

	var scanResults []Result
	db.Table("users").Select("name, age").Where("name like ?", "%jinzhu%").Scan(&scanResults)
	fmt.Println("scan:")
	fmt.Println(scanResults)

	// Raw SQL ---- 通过Raw()直接执行sql
	var rawResult Result
	db.Raw("SELECT name, age FROM users WHERE name = ?", "jinzhu").Scan(&rawResult)
	fmt.Println("raw:")
	fmt.Println(rawResult)
}