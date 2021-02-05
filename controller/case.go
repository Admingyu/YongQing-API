package controller

import (
	"YongQing-API/database"
	"YongQing-API/errors"
	"YongQing-API/model"
	"YongQing-API/schema"
	"YongQing-API/serialization"

	"github.com/gin-gonic/gin"
)

func RegisterCase(rg *gin.RouterGroup) {
	r := rg.Group("/case")
	r.GET("/category", Category)
	r.GET("", CaseList)
	r.GET("/slides", CaseSlides)
	r.GET("/detail", CaseDetail)
}

func Category(c *gin.Context) {

	var categories []serialization.CategorySer
	err := database.DB.Model(model.Category{}).Select("id,icon,text,page").Order("id").Scan(&categories).Error
	errors.HandleError(err)

	Success(c, categories)
}

func CaseList(c *gin.Context) {
	var caseData []serialization.CaseListSer
	err := database.DB.Model(model.Case{}).Select("id, image, `desc`, price, location").Order("id desc").Scan(&caseData).Error
	errors.HandleError(err)

	Success(c, caseData)
}

func CaseSlides(c *gin.Context) {
	var slides []serialization.SlideSer
	err := database.DB.Model(model.Slides{}).Order("id").Select("id, name, image, link").Scan(&slides).Error
	errors.HandleError(err)
	Success(c, slides)
}

func CaseDetail(c *gin.Context) {
	var params schema.IDSchema
	err := c.ShouldBindQuery(&params)
	errors.HandleError(err)

	var data serialization.CaseDetail
	err = database.DB.Model(model.Case{}).Where("id=?", params.ID).Select("id, image,video, content").Scan(&data).Error
	errors.HandleError(err)
	Success(c, data)
}
