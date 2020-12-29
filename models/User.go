package models

import (
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
)

type User struct {
	Id         int64  `orm:"auto"`
	UserName   string `orm:"size(255)"`
	Address	   string  `orm:"size(255)"`
}

func init() {
	orm.RegisterModel(new(User))
}

func SaveUser (m *User) (err error){
	o := orm.NewOrm()
	var num int64
	if m.Id == 0 {
		if num, err = o.Insert(m); err == nil {
			logs.Info("Number of records insert in database:", num)
		}
	}else {
		var tmp *User
		tmp, err = GetUserById(m.Id)

		if err == nil {
			tmp.UserName = m.UserName
			tmp.Address = m.Address

			if num, err = o.Update(tmp); err == nil {
				logs.Info("Number of records updated in database:", num)
			}
		}
	}
	return
}

func GetUserById(id int64) (v *User, err error) {
	o := orm.NewOrm()
	v = &User{Id: id}
	if err = o.Read(v, "Id"); err == nil {
		return v, nil
	}
	return nil, err
}

func DeleteUser(id int64) (err error) {
	o :=orm.NewOrm()
	v :=User{Id: id}
	if err = o.Read(&v, "Id"); err == nil {
		if num, err := o.Delete(&User{Id: id}); err == nil {
			logs.Info("Number of records deleted in database:", num)
		}
	}
	return
}

func AllUsers() (dataList []interface{},err error) {
	var list []User
	o := orm.NewOrm()
	qs := o.QueryTable(new(User))

	if _, err = qs.All(&list); err == nil {
		for _, v := range list {
			dataList = append(dataList, v)
		}
		return dataList, nil
	}
	return nil, err
}