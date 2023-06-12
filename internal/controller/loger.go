package controller

import (
	"gindemo/internal/common"
	"gindemo/internal/model"
	"log"
	"net/http"
	"os/exec"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/goccy/go-json"
	carbon "github.com/golang-module/carbon/v2"
	"gorm.io/gorm"
)

type LogerController struct {
}
type htmlAttribute struct {
	Sort     bool   `json:"sort,omitempty"`
	Hide     bool   `json:"hide,omitempty"`
	MinWidth int    `json:"minWidth,omitempty"`
	Field    string `json:"field,omitempty"`
	Style    string `json:"style,omitempty"`
	Width    string `json:"width,omitempty"`
	Title    string `json:"title,omitempty"`
	Align    string `json:"align,omitempty"`
	Type     string `json:"type,omitempty"`
	Templet  string `json:"templet,omitempty"`
	Toolbar  string `json:"toolbar,omitempty"`
	Fixed    string `json:"fixed,omitempty"`
}
type htmlAttrArr []*htmlAttribute

var (
	tpl_barDemo = `
		<script type="text/html">
			{{#  if(d.result_type == '文件夹地址' || d.result_type == '3'){ }}
				<a class="layui-btn layui-btn-xs layui-btn-warm" lay-event="openfile">打开文件夹</a>
			{{#  } else if(d.result_type == '图片' || d.result_type == '1') { }}
				<a class="layui-btn layui-btn-xs layui-btn-normal" lay-event="watchimg">查看图片</a>
			{{#  } else if(d.result_type == '数值' || d.result_type == '2') { }}
				<a class="layui-btn layui-btn-xs" lay-event="watchresult">查看数值</a>
			{{#  } }}
  		</script>`
	tpl_insert_type = `
		<script type="text/html">
			{{#  if(d.insert_type == '1'){ }}
				批量录入
			{{#  } else if(d.insert_type == '2') { }}
				手动录入
			{{#  } }}
  		</script>`
)

func (l *LogerController) Querylog(ctx *gin.Context) {
	workparam := &model.WorkParams{}
	fieldsarr := workparam.GetData()
	paramsarr := make([]*model.WorkParams, 0, 20)
	htmlinit := make(htmlAttrArr, 3)
	columnsarr := make([]htmlAttrArr, 1)
	columnsarr2 := make([]htmlAttrArr, 1)
	htmlinit[0], htmlinit[1], htmlinit[2] =
		&htmlAttribute{
			Field:    "checknum",
			Width:    "5%",
			MinWidth: 60,
			Align:    "center",
			Type:     "checkbox",
		},
		&htmlAttribute{
			Field:    "id",
			Width:    "5%",
			MinWidth: 60,
			Align:    "center",
			Title:    "序号",
			Type:     "numbers",
		},
		&htmlAttribute{
			Field:    "kid",
			Width:    "20",
			MinWidth: 60,
			Hide:     true,
			Align:    "center",
			Title:    "kid",
			Type:     "space",
		}
	columnsarr[0], columnsarr2[0] = htmlinit, htmlinit
	for _, val := range fieldsarr {
		title := val.Info
		if val.Unit != "" {
			title = title + "(" + val.Unit + ")"
		}
		htmlattr := &htmlAttribute{
			Field:    val.Name,
			MinWidth: 60,
			Width:    strconv.Itoa(int(val.HTMLWidth)),
			Title:    title,
			Sort:     true,
			Align:    "center",
		}
		if val.Name == "zhuansu" || val.Name == "qingjiao" || val.Name == "lcckbz" {
			paramsarr = append(paramsarr, val)
			htmlattr.Hide = true
			htmlattr.Style = "color:#393D49;"
		} else {
			htmlattr.Hide = false
			htmlattr.Style = "color:#FF5722;"
		}
		columnsarr2[0] = append(columnsarr2[0], htmlattr)
		htmlattr2 := &htmlAttribute{
			Field:    val.Name,
			MinWidth: 60,
			Width:    strconv.Itoa(int(val.HTMLWidth)),
			Title:    title,
			Sort:     true,
			Align:    "center",
		}
		htmlattr2.Style = "color:#FF5722;"
		htmlattr2.Hide = true
		columnsarr[0] = append(columnsarr[0], htmlattr2)
	}
	columnsarr[0] = append(columnsarr[0],
		newHtmlAttr(gin.H{"Field": "result_type", "Width": "120", "Title": "数值类型", "Hide": true, "Align": "center"}),
		newHtmlAttr(gin.H{"Field": "result_value", "Width": "90", "Title": "结果值", "Hide": true, "Align": "center"}),
		newHtmlAttr(gin.H{"Field": "addtime", "Width": "15%", "Title": "查询时间", "Align": "center", "Sort": true}),
		newHtmlAttr(gin.H{"Field": "result_info", "Width": "15%", "Align": "center", "Title": "结果类型"}),
		newHtmlAttr(gin.H{"Field": "result_id", "Width": "150", "Title": "结果id", "Hide": true, "Align": "center", "Type": "space"}),
		newHtmlAttr(gin.H{"Field": "insert_type", "Width": "120", "Title": "录入方式", "Hide": true, "Align": "center", "Templet": tpl_insert_type}),
		newHtmlAttr(gin.H{"Width": "13%", "Title": "操作", "Toolbar": tpl_barDemo, "Type": "space", "Fixed": "right"}),
	)
	columnsarr2[0] = append(columnsarr2[0],
		newHtmlAttr(gin.H{"Field": "insert_type", "Width": "120", "Title": "录入方式", "Align": "center", "Templet": tpl_insert_type}),
		newHtmlAttr(gin.H{"Field": "insert_time", "Width": "180", "Title": "录入时间", "Align": "center", "Sort": true, "Templet": ""}),
		newHtmlAttr(gin.H{"Width": "90", "Title": "操作", "Toolbar": "#resultselect", "Type": "space", "Fixed": "right"}),
	)
	resultarr := [3]map[string]any{
		0: {
			"title": "图片",
			"id":    "parent_1",
			"child": []map[string]string{},
		},
		1: {
			"title": "数值",
			"id":    "parent_2",
			"child": []map[string]string{},
		},
		2: {
			"title": "文件夹地址",
			"id":    "parent_3",
			"child": []map[string]string{},
		},
	}
	workresult := model.WorkResult{}
	res2 := workresult.GetData()
	for _, val := range res2 {
		resultarr[val.Type-1]["child"] = append((resultarr[val.Type-1]["child"]).([]map[string]string), map[string]string{
			"title":  val.Info,
			"id":     strconv.Itoa(val.ID),
			"mytype": strconv.Itoa(int(val.Type)),
		})
	}
	columnsarr_json, _ := json.Marshal(columnsarr)
	columnsarr2_json, _ := json.Marshal(columnsarr2)
	resultarr_json, _ := json.Marshal(resultarr)
	res := gin.H{
		"columns":   string(columnsarr_json),
		"columns2":  string(columnsarr2_json),
		"resultarr": string(resultarr_json),
		"fields":    paramsarr,
	}
	ctx.HTML(http.StatusOK, "Querylogindex.html", res)
}
func (l *LogerController) SearchData(ctx *gin.Context) {
	res := gin.H{
		"code":  1,
		"msg":   "暂无相关数据",
		"count": 0,
		"data":  nil,
	}
	isclick, err := common.PostInt("isclick", ctx)
	if isclick == 0 || err != nil {
		ctx.JSON(200, res)
		return
	}
	search_obj := &model.WorkDatasV3{}
	page_num, _ := common.PostInt("page", ctx)
	limit_num, _ := common.PostInt("limit", ctx)
	search_obj.Zhuansu = ctx.PostForm("zhuansu")
	search_obj.Qingjiao = ctx.PostForm("qingjiao")
	search_obj.Lcckbz = ctx.PostForm("lcckbz")
	search_obj.InsertType, _ = common.PostInt8("insert_type", ctx)
	res_data := search_obj.GetDataByWhere((page_num-1)*limit_num, limit_num)
	if len(res_data) == 0 {
		ctx.JSON(200, res)
		return
	}
	count_num := search_obj.GetCount()
	res["code"] = 0
	res["msg"] = "查询数据成功！"
	res["count"] = count_num
	res["data"] = res_data
	ctx.JSON(200, res)
}
func (l *LogerController) SearchResult(ctx *gin.Context) {
	obj := &model.WorkDataresult{}
	res := gin.H{
		"status":  0,
		"resdata": nil,
	}
	pkid := ctx.PostForm("kid")
	result_id := ctx.PostForm("result_id")
	res_data := obj.GetDataByWhere(pkid, result_id)
	if len(res_data) == 0 {
		ctx.JSON(200, res)
		return
	}
	res["status"] = 1
	res["resdata"] = res_data[0]
	ctx.JSON(200, res)
}
func (l *LogerController) DeleteData(ctx *gin.Context) {
	record := &model.WorkRecordV3{}
	kids := ctx.PostForm("kid")
	kids_arr := strings.Split(kids, ",")
	img_files := record.GetDataByWhere("filepath", "pkid in ?", kids_arr)
	model.DB().Transaction(func(tx *gorm.DB) error {
		wdatares := &model.WorkDataresult{}
		wdatav3 := &model.WorkDatasV3{}
		record.Delete(tx, "pkid in ?", kids_arr)
		wdatav3.Delete(tx, "kid in ?", kids_arr)
		wdatares.Delete(tx, "kid in ?", kids_arr)
		return nil
	})
	if len(img_files) != 0 {
		for _, imgval := range img_files {
			_ = common.DelFile(imgval.Filepath)
		}
	}
	ctx.JSON(200, gin.H{"status": 1})
}
func (l *LogerController) ViewImage(ctx *gin.Context) {
	imgobj := &model.WorkRecordV3{}
	pkid := ctx.Query("pkid")
	result_id := ctx.Query("result_id")
	res := imgobj.GetDataByWhere("*", "pkid = ? AND result_id = ?", pkid, result_id)
	for _, v := range res {
		v.Filepath = strings.ReplaceAll(v.Filepath, "\\", "/")
	}
	ctx.HTML(200, "viewImage.html", gin.H{
		"datas": res,
	})
}
func (l *LogerController) DeleteImg(ctx *gin.Context) {
	imgobj := &model.WorkRecordV3{}
	resdata := gin.H{
		"status": 0,
	}
	id, _ := common.PostInt("id", ctx)
	res := imgobj.GetDataByWhere("*", "id = ?", id)
	if len(res) == 0 {
		resdata["status"] = 1
		ctx.JSON(200, resdata)
		return
	}
	if common.DelFile(res[0].Filepath) {
		resdata["status"] = 1
		imgobj.Delete(nil, "id = ?", id)
	}
	ctx.JSON(200, resdata)
	return
}
func (l *LogerController) CheckImg(ctx *gin.Context) {
	imgobj := &model.WorkRecordV3{}
	resdata := gin.H{"status": 0}
	pkid := ctx.PostForm("kid")
	result_id := ctx.PostForm("result_id")
	res := imgobj.GetDataByWhere("id", "pkid = ? AND result_id = ?", pkid, result_id)
	if len(res) != 0 {
		resdata["status"] = 1
	}
	ctx.JSON(200, resdata)
}
func (l *LogerController) ExportDatas(ctx *gin.Context) {

}
func (l *LogerController) OpenFileDir(ctx *gin.Context) {
	file_path := ctx.PostForm("path")
	argstr := "/c start explorer " + file_path
	if common.FileExists(file_path) {
		if err := exec.Command("cmd", argstr).Run(); err != nil {
			ctx.JSON(200, gin.H{"status": 0})
			return
		}
		ctx.JSON(200, gin.H{"status": 1})
		return
	}
	ctx.JSON(200, gin.H{"status": 0})
}
func (l *LogerController) SaveLogs(ctx *gin.Context) {
	res_data := gin.H{"status": 0}
	log_model := &model.WorkLog{}
	if err := ctx.ShouldBindJSON(log_model); err != nil {
		ctx.JSON(200, res_data)
		return
	}
	log_model.ID = 0
	log_model.Addtime = int(time.Now().Unix())
	if err := log_model.Insert(); err != nil {
		ctx.JSON(200, res_data)
		log.Default().Println(err)
		return
	}
	res_data["status"] = 1
	ctx.JSON(200, res_data)
}
func (l *LogerController) DeleteLog(ctx *gin.Context) {
	res_data := gin.H{"status": 0, "info": "删除失败"}
	log_model := &model.WorkLog{}
	ids := strings.Split(ctx.PostForm("ids"), ",")
	if err := log_model.Delete("id in ?", ids); err != nil {
		ctx.JSON(200, res_data)
		log.Default().Println(err)
		return
	}
	res_data["status"] = 1
	res_data["info"] = "删除失败"
	ctx.JSON(200, res_data)
}

type ViewWorkLog struct {
	*model.WorkLog
	ResultType string `json:"result_type,omitempty"`
	Addtime    string `json:"addtime,omitempty"`
}

func (l *LogerController) GetData(ctx *gin.Context) {
	res_data := gin.H{
		"code":  1,
		"msg":   "暂无相关数据",
		"count": 0,
		"data":  nil,
	}
	log_mod := &model.WorkLog{}
	page, _ := common.PostInt("page", ctx)
	limit, _ := common.PostInt("limit", ctx)
	count_num := log_mod.Count()
	if count_num == 0 {
		ctx.JSON(200, res_data)
		return
	}
	fieldsarr := log_mod.GetData((page-1)*limit, limit)
	typestr := []string{1: "图片", 2: "数值", 3: "文件夹地址"}
	res_data["code"] = 0
	res_data["count"] = count_num
	res_data["msg"] = "查询日志成功！"
	res_data["data"] = make([]*ViewWorkLog, 0, 64)
	for _, v := range fieldsarr {
		res := &ViewWorkLog{}
		res.WorkLog = v
		res.ResultType = typestr[v.ResultType]
		res.Addtime = carbon.CreateFromTimestamp(int64(v.Addtime)).Format("Y-m-d H:i:s")
		res_data["data"] = append(res_data["data"].([]*ViewWorkLog), res)
	}
	ctx.JSON(200, res_data)
}
func newHtmlAttr(data map[string]any) *htmlAttribute {
	res := &htmlAttribute{}
	for k, v := range data {
		switch k {
		case "Sort":
			res.Sort, _ = v.(bool)
		case "Hide":
			res.Hide, _ = v.(bool)
		case "MinWidth":
			res.MinWidth, _ = v.(int)
		case "Field":
			res.Field, _ = v.(string)
		case "Style":
			res.Style, _ = v.(string)
		case "Width":
			res.Width, _ = v.(string)
		case "Title":
			res.Title, _ = v.(string)
		case "Align":
			res.Align, _ = v.(string)
		case "Type":
			res.Type, _ = v.(string)
		case "Templet":
			res.Templet, _ = v.(string)
		case "Toolbar":
			res.Toolbar, _ = v.(string)
		case "Fixed":
			res.Fixed, _ = v.(string)
		}
	}
	return res
}
