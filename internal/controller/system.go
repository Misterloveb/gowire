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
	*common.Server
}

const (
	imgdirpath = "./upload/imgpath.txt"
)

func (ctl *SystemController) SystemSettings(ctx *gin.Context) {
	params_arr := ctl.WorkParamsDao.GetData()
	result_arr := ctl.WorkResultDao.GetData()
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

func (ctl *SystemController) SaveParamsDatas(ctx *gin.Context) {
	res_data := gin.H{"status": 0}
	ctx.Request.ParseForm()
	post_arr := ctx.Request.PostForm
	if err := ctl.WorkParamsDao.Update(post_arr); err != nil {
		ctx.JSON(200, res_data)
		return
	}
	res_data["status"] = 1
	ctx.JSON(200, res_data)
}
func (ctl *SystemController) SaveImgDirPath(ctx *gin.Context) {
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
func (ctl *SystemController) AddResult(ctx *gin.Context) {
	mod_workresult := &model.WorkResult{}
	_ = ctx.ShouldBindWith(mod_workresult, binding.Form)
	res_data := gin.H{"status": 0}
	if err := ctl.WorkResultDao.Insert([]*model.WorkResult{mod_workresult}); err != nil {
		ctx.JSON(200, res_data)
		return
	}
	res_data["status"] = 1
	res_data["id"] = mod_workresult.ID
	ctx.JSON(200, res_data)
}
func (ctl *SystemController) EditReuslt(ctx *gin.Context) {
	mod_workresult := &model.WorkResult{}
	_ = ctx.ShouldBindWith(mod_workresult, binding.Form)
	res_data := gin.H{"status": 0}
	if err := ctl.WorkResultDao.Update(mod_workresult); err != nil {
		ctx.JSON(200, res_data)
		return
	}
	res_data["status"] = 1
	ctx.JSON(200, res_data)
}
func (ctl *SystemController) DeleteResult(ctx *gin.Context) {
	mod_workresult := &model.WorkResult{}
	id, err := strconv.Atoi(ctx.PostForm("id"))
	if err != nil {
		ctx.JSON(200, gin.H{"status": 0})
		return
	}
	mod_workresult.ID = id
	if err := ctl.WorkResultDao.Delete(mod_workresult); err != nil {
		ctx.JSON(200, gin.H{"status": 0})
		return
	}
	ctx.JSON(200, gin.H{"status": 1})
}
