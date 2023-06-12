package controller

import (
	"gindemo/internal/common"
	"gindemo/internal/model"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

type SystemController struct {
}

const (
	imgdirpath = "./upload/imgpath.txt"
)

func (s *SystemController) SystemSettings(ctx *gin.Context) {
	mod_work_params := model.WorkParams{}
	mod_work_result := model.WorkResult{}
	params_arr := mod_work_params.GetData()
	result_arr := mod_work_result.GetData()
	file_obj := []byte{}
	if common.FileExists(imgdirpath) {
		var err error
		file_obj, err = os.ReadFile(imgdirpath)
		if err != nil {
			log.Println(err.Error())
		}
	}
	ctx.HTML(http.StatusOK, "saveimgpath.html", gin.H{
		"params":      params_arr,
		"reslutsdata": result_arr,
		"imgdirpath":  string(file_obj),
	})
}

func (s *SystemController) SaveParamsDatas(ctx *gin.Context) {
	res_data := gin.H{"status": 0}
	mod_work_params := model.WorkParams{}
	ctx.Request.ParseForm()
	post_arr := ctx.Request.PostForm
	if err := mod_work_params.Update(post_arr); err != nil {
		ctx.JSON(200, res_data)
		return
	}
	res_data["status"] = 1
	ctx.JSON(200, res_data)
}
func (s *SystemController) SaveImgDirPath(ctx *gin.Context) {
	imagpath := strings.Trim(ctx.PostForm("imgpath"), " ")
	if !common.FileExists(imagpath) {
		ctx.JSON(200, gin.H{"status": 0, "msg": "文件夹不存在,请重新输入!"})
		return
	}
	fileobj, err := os.OpenFile(imgdirpath, os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		ctx.JSON(200, gin.H{"status": 0, "msg": err.Error()})
		return
	}
	defer fileobj.Close()
	imagpath = strings.TrimRight(imagpath, "/")
	imagpath = strings.TrimRight(imagpath, "\\")
	_, err = fileobj.WriteString(imagpath)
	if err != nil {
		ctx.JSON(200, gin.H{"status": 0, "msg": err.Error()})
		return
	}
	ctx.JSON(200, gin.H{"status": 1})
}
func (s *SystemController) AddResult(ctx *gin.Context) {
	mod_workresult := &model.WorkResult{}
	_ = ctx.ShouldBindWith(mod_workresult, binding.Form)
	res_data := gin.H{"status": 0}
	if err := mod_workresult.Insert(); err != nil {
		ctx.JSON(200, res_data)
		return
	}
	res_data["status"] = 1
	res_data["id"] = mod_workresult.ID
	ctx.JSON(200, res_data)
}
func (s *SystemController) EditReuslt(ctx *gin.Context) {
	mod_workresult := &model.WorkResult{}
	_ = ctx.ShouldBindWith(mod_workresult, binding.Form)
	res_data := gin.H{"status": 0}
	if err := mod_workresult.Update(); err != nil {
		ctx.JSON(200, res_data)
		return
	}
	res_data["status"] = 1
	ctx.JSON(200, res_data)
}
func (s *SystemController) DeleteResult(ctx *gin.Context) {
	mod_workresult := &model.WorkResult{}
	id, err := strconv.Atoi(ctx.PostForm("id"))
	if err != nil {
		ctx.JSON(200, gin.H{"status": 0})
		return
	}
	mod_workresult.ID = id
	if err := mod_workresult.Delete(); err != nil {
		ctx.JSON(200, gin.H{"status": 0})
		return
	}
	ctx.JSON(200, gin.H{"status": 1})
}
