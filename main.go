package main

import (
	"fmt"
	"learn-gorm/dao"
	"learn-gorm/model"
	"learn-gorm/utils"
	"time"
)

func main() {
	//testAdd()

	//testFind()

	//testFind2()

	//testFind3()

	//testUpdate()

	//testOntToMany()

	testMany2many()
}

// 测试添加数据
func testAdd() {
	stuDao := dao.StudentDao{}

	//单条添加
	stuDao.AddStudent("金伟民", "1914010805", 0, "2001-06-24")
	stuDao.AddStudentOnlyName("King")
	stuDao.AddStudentWithoutBrith("金伟民", "1914010805", 0)
	stuDao.AddStudentByMap(map[string]interface{}{
		"CreatedAt": time.Now(),
		"UpdatedAt": time.Now(),
		"Name":      "ByMap",
		"Number":    "1914010000",
		"Gender":    0,
		"Birth":     time.Now(),
	})

	//批量添加
	students := []*model.Student{
		{Name: "test1",
			Number: "1914010801",
			Gender: 0,
			Birth:  time.Date(2001, 1, 2, 0, 0, 0, 0, time.Local),
		},
		{Name: "test2",
			Number: "1914010802",
			Gender: 1,
			Birth:  time.Date(2002, 1, 2, 0, 0, 0, 0, time.Local),
		},
	}
	stuDao.AddStudents(students)
	stuDao.AddStudentsByMap([]map[string]interface{}{
		{
			"CreatedAt": time.Now(),
			"UpdatedAt": time.Now(),
			"Name":      "ByMapS1",
			"Number":    "1914010001",
			"Gender":    0,
			"Birth":     time.Now(),
		},
		{
			"CreatedAt": time.Now(),
			"UpdatedAt": time.Now(),
			"Name":      "ByMapS2",
			"Number":    "1914010002",
			"Gender":    0,
			"Birth":     time.Now(),
		},
	})

}

// 测试查找数据
func testFind() {
	stuDao := dao.StudentDao{}

	// 查找第一条数据
	stu, ok := stuDao.FirstStudent()
	if ok {
		fmt.Println(stu)
	}

	// 查找一条数据
	stu, ok = stuDao.FindOneStudent()
	if ok {
		fmt.Println(stu)
	}

	// 查找最后一条数据
	stu, ok = stuDao.LastStudent()
	if ok {
		fmt.Println(stu)
	}

	// 通过 primary key 查找
	stu, ok = stuDao.FindByPk(1)
	if ok {
		fmt.Println(stu)
	}

	// 通过 primary keys 查找
	stus, ok := stuDao.FindByPks([]int{1, 2, 3})
	if ok {
		fmt.Println("=======================================")
		for _, t := range *stus {
			fmt.Println(t)
		}
	}

	// 通过 primary keys 查找
	stus, ok = stuDao.FindAll()
	if ok {
		fmt.Println("=======================================")
		for _, t := range *stus {
			fmt.Println(t)
		}
	}
}

// 测试通过条件 where 查询数据
func testFind2() {
	stuDao := dao.StudentDao{}

	// 通过字符串查找
	stus, ok := stuDao.FindByString("id <> ?", 3)
	if ok {
		for _, t := range *stus {
			fmt.Println(t)
		}
	}
	stus, ok = stuDao.FindByString("name like ?", "%in%")
	if ok {
		fmt.Println("=======================================")
		for _, t := range *stus {
			fmt.Println(t)
		}
	}

	// 通过 map 查找
	stus, ok = stuDao.FindByMap(map[string]interface{}{"gender": 0, "name": "金伟民"})
	if ok {
		fmt.Println("=======================================")
		for _, t := range *stus {
			fmt.Println(t)
		}
	}

	// 通过主键切片查询
	stus, ok = stuDao.FindByPksWhere([]int{5, 6})
	if ok {
		fmt.Println("=======================================")
		for _, t := range *stus {
			fmt.Println(t)
		}
	}

	// 通过结构体并指定设置的属性
	stus, ok = stuDao.FindByStruct(&model.Student{
		Name:   "金伟民",
		Number: "",
		Gender: 0,
		Birth:  time.Time{},
	}, "gender", "name")
	if ok {
		fmt.Println("=======================================")
		for _, t := range *stus {
			fmt.Println(t)
		}
	}

}

// 测试通过内联查询数据，基本同 where 用法相同
func testFind3() {
	stuDao := dao.StudentDao{}

	// 内联 map
	stus, ok := stuDao.FindByInline(map[string]interface{}{"name": "金伟民"})
	if ok {
		fmt.Println("=======================================")
		for _, t := range *stus {
			fmt.Println(t)
		}
	}

	// 内联 string
	stus, ok = stuDao.FindByInline("id > ?", 5)
	if ok {
		fmt.Println("=======================================")
		for _, t := range *stus {
			fmt.Println(t)
		}
	}
}

/*
	Not: 与 where 用法相同
	Or: 与 where 联用，用法同 where
	Select: 查询部分属性，在 Find 前使用
		db.Select("name", "age").Find(&users)
		db.Select([]string{"name", "age"}).Find(&users)
	Order:
		db.Order("age desc, name").Find(&users)
		db.Order("age desc").Order("name").Find(&users)
	Limit: 限制最大查询数量
	Offset: 偏移量
		db.Limit(10).Offset(5).Find(&users)
	Distinct: 查找不相同的值
	.....
*/

/*
	FirstOrInit: 在数据库中查找，未找到则初始化
	FirstOrCreat： 在数据库中查找，未找到则在数据库中创建
	Attrs： 在上述未找到情况下可以指定其他参数

*/

// 测试更新
func testUpdate() {
	stuDao := dao.StudentDao{}

	stu, ok := stuDao.FirstStudent()
	if ok {
		// 通过 Save 更新
		stu.Name = "jwm1"
		ok = stuDao.UpdateBySave(stu)
		if ok {
			fmt.Println("更新成功1")
		}

		// 更新单列
		ok = stuDao.UpdateSingleLine(stu, "name", "jwm2")
		if ok {
			fmt.Println("更新成功2")
		}

		// 通过 map 更新多列
		ok = stuDao.UpdateMultiLinesByMap(stu, map[string]interface{}{
			"name":   "jwm3",
			"number": "1914010805",
		})
		if ok {
			fmt.Println("更新成功3")
		}

	}

}

/*
	Delete: 删除
		db.Delete(&email) // 根据主键删除
		db.Where("name = ?", "stu").Delete(&email) // 根据条件删除
		db.Delete(&users, []int{1,2,3}) // 根据主键删除多个

	如果指定 DeleteAt 或 继承了 gorm.Model 默认是软删除
*/

// 一对多测试
func testOntToMany() {
	db := utils.GetDb()

	user := &model.User{}
	user.ID = 3

	// 不存在则关联创建
	db.FirstOrCreate(&model.User{
		Name: "userName",
		CreditCard: []*model.CreditCard{
			{
				Number: "1234567",
			},
			{
				Number: "7654321",
			},
		},
	})

	// 关联查询，Model 的主键不能为空
	creditCards := &[]*model.CreditCard{}
	err := db.Model(user).Association("CreditCard").Find(creditCards)
	if err != nil {
		fmt.Println(err)
		return
	}
	count := db.Model(&user).Association("CreditCard").Count()
	fmt.Printf("count: %v\n", count)
	fmt.Printf("user: %v\n", user)
	for _, cred := range *creditCards {
		fmt.Printf("creditCard: %v\n", cred)
	}

	// Preload 预加载
	db.Preload("CreditCard").First(user)
	utils.JsonOut(user)

	// Joins 预加载 (此处会报错, Joins 只支持一对一)
	//db.Joins("CreditCard").First(user)
	//utils.JsonOut(user)

	// 嵌套预加载
	//db.Preload("Orders.OrderItems.Product").Preload(clause.Associations).Find(&user)

}

// 多对多测试
func testMany2many() {
	db := utils.GetDb()

	user := &model.User{
		Name: "testMany2Many",
		Languages: []*model.Language{
			{
				Name: "英语",
			},
			{
				Name: "中文",
			},
		},
	}

	db.FirstOrCreate(user)

}
