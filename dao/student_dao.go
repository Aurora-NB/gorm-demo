package dao

import (
	. "learn-gorm/model"
	"learn-gorm/utils"
	"log"
	"time"
)

type StudentDao struct {
}

/*
	增
*/
// AddStudent 添加一个学生，包含全部信息
func (s StudentDao) AddStudent(name, number string, gender int, birth string) bool {
	// 获取实例
	db := utils.GetDb()

	// 解析生日
	b, err := time.Parse("2006-01-02", birth)
	if err != nil {
		log.Printf("add student failed, can not parse birth, err = %v", err)
		return false
	}

	// 添加
	stu := Student{
		Name:   name,
		Number: number,
		Gender: gender,
		Birth:  b,
	}
	res := db.Create(&stu)

	// 检查是否添加成功
	if res.Error == nil && res.RowsAffected == 1 {
		log.Printf("stu add success")
		return true
	} else {
		log.Printf("stu add failed")
		return false
	}
}

// AddStudentOnlyName 添加一个学生，只包含 name
func (s StudentDao) AddStudentOnlyName(name string) bool {
	// 获取实例
	db := utils.GetDb()

	// 添加
	stu := Student{
		Name: name,
	}
	res := db.Select("Name").Create(&stu)

	// 检查是否添加成功
	if res.Error == nil && res.RowsAffected == 1 {
		log.Printf("stu add success")
		return true
	} else {
		log.Printf("stu add failed")
		return false
	}
}

// AddStudentWithoutBrith 添加一个学生，不包含 birth
func (s StudentDao) AddStudentWithoutBrith(name, number string, gender int) bool {
	// 获取实例
	db := utils.GetDb()

	// 添加
	stu := Student{
		Name:   name,
		Number: number,
		Gender: gender,
	}
	res := db.Omit("birth").Create(&stu)

	// 检查是否添加成功
	if res.Error == nil && res.RowsAffected == 1 {
		log.Printf("stu add success")
		return true
	} else {
		log.Printf("stu add failed")
		return false
	}
}

// AddStudents 批量添加学生
func (s StudentDao) AddStudents(students []*Student) bool {
	// 获取实例
	db := utils.GetDb()

	// 添加
	res := db.Create(&students)

	// 检查是否添加成功
	if res.Error == nil {
		log.Printf("stu add success")
		return true
	} else {
		log.Printf("stu add failed")
		return false
	}
}

// AddStudentByMap 通过 map 添加一个学生
func (s StudentDao) AddStudentByMap(student map[string]interface{}) bool {
	// 获取实例
	db := utils.GetDb()

	// 添加
	res := db.Model(&Student{}).Create(student)

	// 检查是否添加成功
	if res.Error == nil {
		log.Printf("stu add success")
		return true
	} else {
		log.Printf("stu add failed")
		return false
	}
}

// AddStudentsByMap 通过 map 批量添加学生
func (s StudentDao) AddStudentsByMap(student []map[string]interface{}) bool {
	// 获取实例
	db := utils.GetDb()

	// 添加
	res := db.Model(&Student{}).Create(student)

	// 检查是否添加成功
	if res.Error == nil {
		log.Printf("stu add success")
		return true
	} else {
		log.Printf("stu add failed")
		return false
	}
}

/*
	查
*/
// FirstStudent 按照主键返回第一个
func (s StudentDao) FirstStudent() (*Student, bool) {
	// 获取实例
	db := utils.GetDb()

	// 查找
	stu := &Student{}
	res := db.First(stu)

	// 检查
	if res.Error == nil && res.RowsAffected == 1 {
		return stu, true
	} else {
		return nil, false
	}
}

// FindOneStudent 按照主键返回第一个
func (s StudentDao) FindOneStudent() (*Student, bool) {
	// 获取实例
	db := utils.GetDb()

	// 查找
	stu := &Student{}
	res := db.Take(stu)

	// 检查
	if res.Error == nil && res.RowsAffected == 1 {
		return stu, true
	} else {
		return nil, false
	}
}

// LastStudent 按照主键返回第一个
func (s StudentDao) LastStudent() (*Student, bool) {
	// 获取实例
	db := utils.GetDb()

	// 查找
	stu := &Student{}
	res := db.Last(stu)

	// 检查
	if res.Error == nil && res.RowsAffected == 1 {
		return stu, true
	} else {
		return nil, false
	}
}

// FindByPk 通过主键查找
func (s StudentDao) FindByPk(id int) (*Student, bool) {
	// 获取实例
	db := utils.GetDb()

	// 查找
	stu := &Student{}
	res := db.First(stu, id)

	// 检查
	if res.Error == nil && res.RowsAffected == 1 {
		return stu, true
	} else {
		return nil, false
	}
}

// FindByPks 通过主键批量查找
func (s StudentDao) FindByPks(ids []int) (*[]Student, bool) {
	// 获取实例
	db := utils.GetDb()

	// 查找
	stus := &[]Student{}
	res := db.Find(stus, ids)

	// 检查
	if res.Error == nil {
		return stus, true
	} else {
		return nil, false
	}
}

// FindAll 通过主键批量查找
func (s StudentDao) FindAll() (*[]Student, bool) {
	// 获取实例
	db := utils.GetDb()

	// 查找
	stus := &[]Student{}
	res := db.Find(stus)

	// 检查
	if res.Error == nil {
		return stus, true
	} else {
		return nil, false
	}
}

/*
	通过 where 查
*/
// FindByString 通过字符串查询
func (s StudentDao) FindByString(query string, args ...interface{}) (*[]Student, bool) {
	// 获取实例
	db := utils.GetDb()

	// 查找
	stus := &[]Student{}
	res := db.Where(query, args...).Find(stus)

	// 检查
	if res.Error == nil {
		return stus, true
	} else {
		return nil, false
	}
}

// FindByMap 通过 map 查询
func (s StudentDao) FindByMap(query map[string]interface{}) (*[]Student, bool) {
	// 获取实例
	db := utils.GetDb()

	// 查找
	stus := &[]Student{}
	res := db.Where(query).Find(stus)

	// 检查
	if res.Error == nil {
		return stus, true
	} else {
		return nil, false
	}
}

// FindByPksWhere 通过主键切片查询
func (s StudentDao) FindByPksWhere(pks []int) (*[]Student, bool) {
	// 获取实例
	db := utils.GetDb()

	// 查找
	stus := &[]Student{}
	res := db.Where(pks).Find(stus)

	// 检查
	if res.Error == nil {
		return stus, true
	} else {
		return nil, false
	}
}

// FindByStruct 通过结构体查询
func (s StudentDao) FindByStruct(student *Student, args ...interface{}) (*[]Student, bool) {
	// 获取实例
	db := utils.GetDb()

	// 查找
	stus := &[]Student{}
	res := db.Where(student, args...).Find(stus)

	// 检查
	if res.Error == nil {
		return stus, true
	} else {
		return nil, false
	}
}

/*
	内联查询，上面通过 where 查询中的参数都可以直接放到 Find、First 中
*/
func (s StudentDao) FindByInline(query ...interface{}) (*[]Student, bool) {
	// 获取实例
	db := utils.GetDb()

	// 查找
	stus := &[]Student{}
	res := db.Find(stus, query...)

	// 检查
	if res.Error == nil {
		return stus, true
	} else {
		return nil, false
	}
}

/*
	更新
*/
// UpdateBySave 通过 Save 将对象根据主键同步到数据库,即使是 "0" 值也会被同步
func (s StudentDao) UpdateBySave(student *Student) bool {
	// 获取实例
	db := utils.GetDb()

	// 存储
	res := db.Save(student)

	// 检查
	if res.Error == nil {
		return true
	} else {
		return false
	}
}

// UpdateSingleLine 修改单列
func (s StudentDao) UpdateSingleLine(student *Student, key string, value interface{}) bool {
	// 获取实例
	db := utils.GetDb()

	// 存储
	res := db.Model(student).Update(key, value)

	// 检查
	if res.Error == nil {
		return true
	} else {
		return false
	}
}

// UpdateMultiLinesByMap 通过 map 更新多列
func (s StudentDao) UpdateMultiLinesByMap(student *Student, m map[string]interface{}) bool {
	// 获取实例
	db := utils.GetDb()

	// 存储
	res := db.Model(student).Updates(m)

	// 检查
	if res.Error == nil {
		return true
	} else {
		return false
	}
}
