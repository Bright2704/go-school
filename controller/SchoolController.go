package controller

import (
	"example/BatteryTracking/entity"
	"example/BatteryTracking/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type SchoolController struct {
	SchoolService service.SchoolService
}

func New(schoolservice service.SchoolService) SchoolController {
	return SchoolController{
		SchoolService: schoolservice,
	}
}

func (uc *SchoolController) CreateSchool(ctx *gin.Context) {
	var user entity.UserSchool
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	err := uc.SchoolService.CreateSchool(&user)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "success"})
}

func (uc *SchoolController) GetSchool(ctx *gin.Context) {
	var username string = ctx.Param("name")
	UserSchool, err := uc.SchoolService.GetSchool(&username)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, UserSchool)
}

func (uc *SchoolController) GetAll(ctx *gin.Context) {
	users, err := uc.SchoolService.GetAll()
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	} 
	ctx.JSON(http.StatusOK, users)
}

func (uc *SchoolController) UpdateSchool(ctx *gin.Context) {
	var user entity.UserSchool
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	err := uc.SchoolService.UpdateSchool(&user)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "success"})
}

func (uc *SchoolController) DeleteSchool(ctx *gin.Context) {
	var username string = ctx.Param("name")
	err := uc.SchoolService.DeleteSchool(&username)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "success"})
}

func (uc *SchoolController) RegisterUserRouts(rg *gin.RouterGroup) {
	userroute := rg.Group("/school")
	userroute.POST("/create", uc.CreateSchool)
	userroute.GET("/get/:name", uc.GetSchool)
	userroute.GET("/getall", uc.GetAll)
	userroute.PATCH("/update", uc.UpdateSchool)
	userroute.DELETE("/delete/:name", uc.DeleteSchool)
}



