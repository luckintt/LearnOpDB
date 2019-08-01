package Query

import (
	"LearnOpDB/Config"
	"LearnOpDB/Model"
	"fmt"
)

func DirectQueryUser()  {
	var firstUser Model.User
	Config.DB.First(&firstUser) // 按照主键排序，获取第一个记录, 类似：SELECT * FROM users ORDER BY id LIMIT 1;
	fmt.Print("First:")
	fmt.Println(firstUser)

	var takeUser Model.User
	Config.DB.Take(&takeUser) // 不排序，随机取出一条记录, 类似：SELECT * FROM users LIMIT 1;
	fmt.Print("Take:")
	fmt.Println(takeUser)

	var lastUser Model.User
	Config.DB.Last(&lastUser) // 按照主键排序，取出最后一条记录，类似：SELECT * FROM users ORDER BY id DESC LIMIT 1;
	fmt.Print("Last:")
	fmt.Println(lastUser)

	var count int
	var users []Model.User // 获取所有记录，类似：SELECT * FROM users;
	Config.DB.Find(&users).Count(&count)
	fmt.Print("Find:\n")
	for i := 0; i < count; i++ {
		fmt.Println(users[i])
	}

	var firstIdUser Model.User
	Config.DB.First(&firstIdUser, 10) // 当主键是int型时，取出ID=10的记录，类似：SELECT * FROM users WHERE id = 10;
	fmt.Print("First:")
	fmt.Println(firstIdUser)
}

func WhereQueryUser() {
	db := Config.DB

	var user Model.User
	db.Where("name = ?", "jinzhu").First(&user) // SELECT * FROM users WHERE name = 'jinzhu' limit 1;
	fmt.Print("first matched record:")
	fmt.Println(user)

	var users []Model.User
	db.Where("name = ?", "jinzhu").Find(&users) // SELECT * FROM users WHERE name = 'jinzhu';
	fmt.Print("all matched record:\n")
	fmt.Println(users)

	var notEqualUsers []Model.User
	db.Where("name <> ?", "jinzhu").Find(&notEqualUsers)
	fmt.Print("<>:\n")
	fmt.Println(notEqualUsers)

	var inUsers []Model.User
	db.Where("name in (?)", []string{"jinzhu", "jinzhu1"}).Find(&inUsers)
	fmt.Print("in:\n")
	fmt.Println(inUsers)

	var likeUsers []Model.User
	db.Where("name like ?", "%jinzhu1%").Find(&likeUsers)
	fmt.Print("like:\n")
	fmt.Println(likeUsers)

	var andUsers []Model.User
	db.Where("name = ? and age >= ?", "jinzhu", 18).Find(&andUsers)
	fmt.Println("and:")
	fmt.Println(andUsers)

	var timeUsers []Model.User
	db.Where("updated_at > ?", "2019-07-29 18:31:36").Find(&timeUsers)
	fmt.Println("time:")
	fmt.Println(timeUsers)

	var betweenUsers []Model.User
	db.Where("created_at between ? and ?", "2019-07-29 18:17:45",  "2019-07-29 18:31:36").Find(&betweenUsers)
	fmt.Println("between(左右都包含):")
	fmt.Println(betweenUsers)
}

func StructAndMapQueryUser()  {
	db := Config.DB

	var structUser Model.User
	db.Where(&Model.User{Name:"jinzhu", Age:18}).First(&structUser) // SELECT * FROM users WHERE name = "jinzhu" AND age = 18 LIMIT 1;
	fmt.Println("struct:")
	fmt.Println(structUser)

	var mapUsers []Model.User
	db.Where(map[string]interface{}{"name":"jinzhu", "age":18}).Find(&mapUsers) // SELECT * FROM users WHERE name = "jinzhu" AND age = 18;
	fmt.Println("map:")
	fmt.Println(mapUsers)

	var slicePrimaryKey []Model.User
	db.Where([]int{1,5,15}).Find(&slicePrimaryKey) // SELECT * FROM users WHERE id IN (1, 5, 15);
	fmt.Println("slice of map keys:")
	fmt.Println(slicePrimaryKey)

	// ⚠️当使用结构体查询时，gorm只查询非零值(0, "", false)的字段, 零值字段不作为查询条件----可以使用指针类型或者scanner/valuer（sql.NullInt64）来避免这种情况
	var noteStructUsers []Model.User
	db.Where(&Model.User{Name:"jinzhu1", Age:0}).First(&noteStructUsers) // SELECT * FROM users WHERE name = "jinzhu";
	fmt.Println("struct with zero-value:")
	fmt.Println(noteStructUsers)
}

func NotQueryUser()  {
	db := Config.DB

	var firstUser Model.User
	db.Not("name", "jinzhu").First(&firstUser) // SELECT * FROM users WHERE name <> "jinzhu" LIMIT 1;
	fmt.Println("first:")
	fmt.Println(firstUser)

	var notInUsers []Model.User
	db.Not("name", []string{"jinzhu", "jinzhu1"}).Find(&notInUsers) // // SELECT * FROM users WHERE name NOT IN ("jinzhu", "jinzhu1");
	fmt.Println("not in:")
	fmt.Println(notInUsers)

	var notInSlicePrimary Model.User
	db.Not([]int{1,3,5}).First(&notInSlicePrimary) // SELECT * FROM users WHERE id NOT IN (1,3,5);
	fmt.Println("not in slice of primary keys:")
	fmt.Println(notInSlicePrimary)

	var allUsers []Model.User
	db.Not([]int{}).Find(&allUsers) // SELECT * FROM users;
	fmt.Println("all users:")
	fmt.Println(allUsers)

	var plainSQLUser Model.User
	db.Not("name = ?", "jinzhu").First(&plainSQLUser) // SELECT * FROM users WHERE NOT(name = "jinzhu");
	fmt.Println("plain SQL:")
	fmt.Println(plainSQLUser)

	var structUser Model.User
	db.Not(Model.User{Name:"jinzhu"}).First(&structUser) // SELECT * FROM users WHERE name <> "jinzhu"  LIMIT 1;
	fmt.Println("struct:")
	fmt.Println(structUser)
}

func OrQueryUser()  {
	db := Config.DB

	// SELECT * FROM users WHERE name = 'jinzhuuu' OR name = 'jinzhu1';

	var plainSQLUser []Model.User
	db.Where("name = ?", "jinzhu").Or("name = ?", "jinzhu1").Find(&plainSQLUser)
	fmt.Println("plain sql:")
	fmt.Println(plainSQLUser)

	var structUsers []Model.User
	db.Where("name = 'jinzhu'").Or(Model.User{Name:"jinzhu1"}).Find(&structUsers)
	fmt.Println("struct:")
	fmt.Println(structUsers)

	var mapUsers []Model.User
	db.Where("name = 'jinzhu'").Or(map[string]interface{}{"name":"jinzhu1"}).Find(&mapUsers)
	fmt.Println("map:")
	fmt.Println(mapUsers)
}

func FirstOrInitQueryUser()  {
	// FirstOrInit只有struct/map时可用, 当记录不存在时，使用参数初始化结构体,但是不会存储到数据库中；若存在，则将查询的结果赋值给结构体
	db := Config.DB

	// 没有找到时，利用参数初始化结构体，但是并不操作数据库
	var unfoundUser Model.User
	db.Where(Model.User{Name:"non_existing"}).FirstOrInit(&unfoundUser)
	fmt.Println("unfound:")
	fmt.Println(unfoundUser) // unfoundUser -> User{Name: "non_existing"},初始化结构体中的Name字段，但是数据库中并没有这个记录

	// 找到时，直接利用记录来初始化结构体
	var foundStruct Model.User
	db.Where(Model.User{Name:"jinzhu"}).FirstOrInit(&foundStruct)
	fmt.Println("found struct:")
	fmt.Println(foundStruct) // foundStruct -> User{Id: 1, Name: "Jinzhu", Age: 18}

	// 找到时，直接利用记录来初始化map
	var foundMap Model.User
	db.Where(map[string]interface{}{"name":"jinzhu"}).FirstOrInit(&foundMap)
	fmt.Println("found map:")
	fmt.Println(foundMap) // foundMap -> User{Id: 1, Name: "Jinzhu", Age: 18}

	// attrs----当记录不存在时，使用参数初始化结构体, 但是不会更改数据库；若存在，则将查询的结果赋值给结构体
	// SELECT * FROM users WHERE name = 'non_existing';
	// user -> User{Name: "non_existing", Age: 20}
	var unfoundAttrStructUser Model.User
	db.Where(Model.User{Name:"non_existing"}).Attrs(Model.User{Age:20}).FirstOrInit(&unfoundAttrStructUser)
	fmt.Println("unfound attrs struct:")
	fmt.Println(unfoundAttrStructUser) // unfoundAttrStructUser -> User{Name: "non_existing", Age: 20}

	// SELECT * FROM users WHERE name = 'non_existing';
	// user -> User{Name: "non_existing", Age: 20}
	var unfoundAttrUser Model.User
	db.Where(Model.User{Name:"non_existing"}).Attrs("age",30).FirstOrInit(&unfoundAttrUser)
	fmt.Println("unfound attrs:")
	fmt.Println(unfoundAttrUser) // unfoundAttrUser -> User{Name: "non_existing", Age: 30}

	// SELECT * FROM users WHERE name = jinzhu';
	// user -> User{Id: 1, Name: "jinzhu", Age: 18}
	var foundAttrUser Model.User
	db.Where(Model.User{Name:"jinzhu"}).Attrs(Model.User{Age:20}).FirstOrInit(&foundAttrUser)
	fmt.Println("found attrs:")
	fmt.Println(foundAttrUser) // user -> User{Id: 1, Name: "jinzhu", Age: 18}

	// assign----无论是struct还是map:无论记录是否存在，都会将where与assign参数的值赋值给结构体, 但是不会操作数据库
	var unfoundAssignStructUser Model.User
	db.Where(Model.User{Name:"non_existing"}).Assign(Model.User{Age:20}).FirstOrInit(&unfoundAssignStructUser)
	fmt.Println("unfound assign struct:")
	fmt.Println(unfoundAssignStructUser) // unfoundAssignStructUser -> User{Name: "non_existing", Age: 20}

	var foundAssignStructUser Model.User
	db.Where(Model.User{Name:"jinzhu"}).Assign(Model.User{Age:20}).FirstOrInit(&foundAssignStructUser)
	fmt.Println("found assign struct:")
	fmt.Println(foundAssignStructUser) // foundAssignStructUser -> User{Id: 1, Name: "jinzhu", Age: 20}

	var unfoundAssignMapUser Model.User
	db.Where(Model.User{Name:"non_existing"}).Assign(Model.User{Age:20}).FirstOrInit(&unfoundAssignMapUser)
	fmt.Println("unfound assign map:")
	fmt.Println(unfoundAssignMapUser) // unfoundAssignMapUser -> User{Name: "non_existing", Age: 20}

	var foundAssignMapUser Model.User
	db.Where(Model.User{Name:"jinzhu"}).Assign(Model.User{Age:20}).FirstOrInit(&foundAssignMapUser)
	fmt.Println("found assign map:")
	fmt.Println(foundAssignMapUser) // unfoundAssignMapUser -> User{Id: 1, Name: "jinzhu", Age: 20}
}
