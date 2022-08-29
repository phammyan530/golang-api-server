package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"github.com/golang/glog"
	"github.com/labstack/echo/v4"
	"github.com/spf13/cast"
)

func init() {
	orm.RegisterModel(new(User))
}

type User struct {
	Id    int64  `orm:"auto" json:"id"`
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Phone string `json:"phone"`
}

var listUsers = []User{
	{
		Name: "aaaaaaaa",
		Age:  18,
	},
	{
		Name: "bbbbbbbbb",
		Age:  19,
	},
	{
		Name: "cccccccccccc",
		Age:  20,
	},
}

func GetUser(c echo.Context) error {
	id := cast.ToInt64(c.QueryParam("id"))
	// name := c.QueryParam("name")
	o := orm.NewOrm()

	user := &User{
		Id: id,
	}

	err := o.Read(user)
	// err := o.Read(user, "Name")

	if err != nil {
		glog.Errorf("Error get user: %v", err)
		return err
	}
	return c.JSON(http.StatusOK, user)
}

func AddUser(c echo.Context) error {
	user := &User{}
	err := c.Bind(user)
	if err != nil {
		glog.Errorf("Error bind user: %v", err)
		return err
	}
	o := orm.NewOrm()
	id, err := o.Insert(user)
	if err != nil {
		glog.Errorf("Error insert user: %v", err)
		return err
	}
	glog.Infof("insert at row %d", id)
	return c.JSON(http.StatusOK, user)
}

func UpdateUser(c echo.Context) error {
	user := &User{}
	if err := c.Bind(user); err != nil {
		glog.Errorf("bind error user %v", err)
		return err
	}
	glog.Infof("req update user %+v", user)
	o := orm.NewOrm()
	_, err := o.Update(user, "Name", "Phone")
	if err != nil {
		glog.Errorf("Error update user %v: %v", user.Name, err)
		return err
	}

	userUpdate := &User{
		Name: user.Name,
	}

	o.Read(userUpdate, "Name")
	return c.JSON(http.StatusOK, userUpdate)
}

func DeleteUser(c echo.Context) error {
	id := cast.ToInt64(c.Param("id"))
	glog.Infof("Deleting user %v", id)
	user := &User{
		Id: id,
	}

	o := orm.NewOrm()
	row, err := o.Delete(user)
	if err != nil {
		glog.Errorf("Error deleting user %v: %v", id, err)
		return err
	}
	return c.String(http.StatusOK, fmt.Sprintf("User ID %d deleted successfully at %d", id, row))
}

func GetAllUsers_2(c echo.Context) error {
	c.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	c.Response().WriteHeader(http.StatusOK)

	enc := json.NewEncoder(c.Response())
	for _, user := range listUsers {
		if err := enc.Encode(user); err != nil {
			return err
		}
		c.Response().Flush()
		time.Sleep(1 * time.Second)
	}
	return nil
}

func GetAllUsers(c echo.Context) error {

	o := orm.NewOrm()

	var users []User
	num, err := o.QueryTable("user").All(&users)

	if err != nil {
		glog.Errorf("Error get user %v: %v", num, err)
		return err
	}
	return c.JSON(http.StatusOK, users)
}
