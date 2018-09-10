package user

import (
	"github.com/echo-gorm/database"
	"github.com/labstack/echo"
	"github.com/uuidcode/coreutil"
	"net/http"
	"time"
)

type User struct {
	UserId      int64     `gorm:"PRIMARY_KEY" json:"userId" form:"userId" query:"userId"`
	Name        string    `json:"name" form:"name" query:"name"`
	RegDatetime time.Time `json:"regDatetime" form:"regDatetime" query:"regDatetime"`
	ModDatetime time.Time `json:"modDatetime" form:"modDatetime" query:"modDatetime"`
}

func (User) TableName() string {
	return "user"
}

func Index(c echo.Context) error {
	var userList []User

	database.MainDB.Find(&userList)

	c.Logger().Debug(coreutil.ToJson(userList))

	return c.Render(http.StatusOK, "user/index", echo.Map{
		"userList": userList,
	})
}

func Form(c echo.Context) error {
	user := new(User)
	err := c.Bind(user)
	coreutil.CheckErr(err)

	result := echo.Map{}

	if user.UserId != 0 {
		database.MainDB.First(user, User{
			UserId: user.UserId,
		})

		result["user"] = user
	}

	return c.Render(http.StatusOK, "user/form.html", result)
}

func Post(c echo.Context) error {
	user := new(User)
	err := c.Bind(user)
	coreutil.CheckErr(err)

	user.RegDatetime = time.Now()
	user.ModDatetime = time.Now()
	database.MainDB.Create(&user)

	return c.JSON(http.StatusOK, user)
}

func Put(c echo.Context) error {
	user := new(User)
	err := c.Bind(user)
	coreutil.CheckErr(err)

	newUser := new(User)

	database.MainDB.First(newUser, User{
		UserId: user.UserId,
	})

	newUser.Name = user.Name

	database.MainDB.Save(newUser)
	return c.JSON(http.StatusOK, user)
}

func Delete(c echo.Context) error {
	user := new(User)
	err := c.Bind(user)
	coreutil.CheckErr(err)

	database.MainDB.Delete(&user)

	return c.JSON(http.StatusOK, user.UserId)
}

func Get(c echo.Context) error {
	userIdValue := c.Param("userId")

	user := new(User)
	userId, err := coreutil.ParseInt(userIdValue)
	coreutil.CheckErr(err)

	database.MainDB.First(user, User{
		UserId: userId,
	})

	return c.JSON(http.StatusOK, user)
}
