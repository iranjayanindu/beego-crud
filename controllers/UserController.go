package controllers

import (
	"curd/models"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

type UserConteroller struct {
	beego.Controller
}

func (c *UserConteroller) Edit() {
	//获得id
	id, _ := c.GetInt64("Id", 0)
	fmt.Println("Edit eke Id :",id)
	user, err := models.GetUserById(id)

	logs.Info("User by id :", user)

	if err == nil {
		c.Data["User"] = user
	} else {
		fmt.Println("else :")
		tmpUser := &models.User{}
		c.Data["User"] = tmpUser
	}
	c.Data["json"]=map[string]interface{}{"Id":user.Id,"Name":user.UserName,"Address":user.Address};
	c.ServeJSON()




}

func (c *UserConteroller) Save(){
	user := models.User{}
	if err := c.ParseForm(&user); err == nil {
		if err := models.SaveUser(&user); err == nil{
			c.Data["json"] = ""
		}else {
			c.Data["json"] = "error"
		}
	}else {
		c.Data["json"] = "error"
	}
	c.ServeJSON()
}
func (c *UserConteroller) Delete() {
	//获得id
	fmt.Println("methanta enawa")
	id, _ := c.GetInt64("Id", 0)
	fmt.Println("Id : ",id)
	if err := models.DeleteUser(id); err == nil {
		c.Data["json"] = "ok"
	} else {
		c.Data["json"] = "error"
	}
	c.ServeJSON()
}

func (c *UserConteroller) List() {

	dataList, err := models.AllUsers()
	if err == nil {
		c.Data["List"] = dataList
	}
	logs.Info("dataList :", dataList)
	c.TplName = "home.html"

}