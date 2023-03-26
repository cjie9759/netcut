package controller

import (
	"fmt"
	"net/http"
	"netcut/controller/param"
	"netcut/model"

	"github.com/gin-gonic/gin"
)

type CutCtr struct{}

func (T *CutCtr) Get(c *gin.Context) {
	var p param.Page
	c.BindQuery(&p)
	p.Default()
	var data []*model.Cut
	count := model.NewCut().WebList(&p, &data)
	d := make(map[string]interface{}, 5)
	d["data"] = data
	d["count"] = count
	c.JSON(http.StatusOK, d)
}
func (T *CutCtr) Creat(c *gin.Context) {
	var data model.Cut
	err := c.BindJSON(&data)
	fmt.Println(err)

	if data.Text == "" {
		c.JSON(http.StatusBadRequest, "参数错误")
		return
	}

	err = data.Creat()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	d := make(map[string]interface{}, 5)
	d["data"] = data
	c.JSON(http.StatusOK, d)

}
func (T *CutCtr) CreatText(c *gin.Context) {
	var data model.Cut
	data.Text = c.Query("d")

	if data.Text == "" {
		c.JSON(http.StatusBadRequest, "参数错误")
		return
	}

	err := data.Creat()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	d := make(map[string]interface{}, 5)
	d["data"] = data
	c.JSON(http.StatusOK, d)
}
