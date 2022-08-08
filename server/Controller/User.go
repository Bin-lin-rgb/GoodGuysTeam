package Controller

import (
	"github.com/gin-gonic/gin"
	"server/Dao"
	"server/Model"
	"server/response"
)

func AddUser(c *gin.Context) {
	var user Model.User
	if err := c.ShouldBind(&user); err != nil {
		response.Failed("参数错误！", c)
	}
	Dao.Mgr.AddUser(&user)
	response.Success("添加成功！", user, c)
}

func ListUser(c *gin.Context) {
	users, err := Dao.Mgr.GetUser()
	if err != nil {
		response.Failed("返回失败！", c)
	}
	response.Success("数据返回成功！", users, c)
}
