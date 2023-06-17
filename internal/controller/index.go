package controller

import (
	"errors"
	"github.com/Misterloveb/gowire/internal/common"
	"github.com/Misterloveb/gowire/internal/model"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"sync"
	"unicode/utf8"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/golang-module/carbon/v2"
	"github.com/google/uuid"
	"github.com/xuri/excelize/v2"
	"gorm.io/gorm"
)

const (
	xlsxfilename = "数据导入模板.xlsx"
)

type IndexController struct {
	*common.Server
}

func (ctl *IndexController) Index(ctx *gin.Context) {
	session := sessions.Default(ctx)
	ctl.Logger.Info("hello")
	ctx.HTML(http.StatusOK, "home.html", gin.H{
		"user": session.Get("userinfo"),
	})
}

func (ctl *IndexController) Handinsert(ctx *gin.Context) {
	res := ctl.WorkParamsDao.GetData()
	res2 := ctl.WorkResultDao.GetData()
	ctx.HTML(http.StatusOK, "Handinsertindex.html", gin.H{
		"param":  res,
		"result": res2,
	})
}
func (ctl *IndexController) SaveDatas(ctx *gin.Context) {
	insert_data := &model.DatasV3{}
	insert_data2 := &model.Dataresult{}
	res2 := ctl.WorkResultDao.GetData()
	resdata := gin.H{
		"status": 0,
		"info":   "保存失败！",
	}
	if err := ctx.ShouldBindBodyWith(insert_data, binding.JSON); err != nil {
		resdata["info"] = "json1 parse faild"
		ctx.JSON(http.StatusBadRequest, resdata)
	}
	if err := ctx.ShouldBindBodyWith(insert_data2, binding.JSON); err != nil {
		resdata["info"] = "json2 parse faild"
		ctx.JSON(http.StatusBadRequest, resdata)
	}
	res_id, err := strconv.Atoi(insert_data2.ResultID)
	if err != nil {
		resdata["info"] = "参数不合法！"
		ctx.JSON(http.StatusBadRequest, resdata)
	}
	insert_data.Kid = common.HashKey(insert_data)
	insert_data.InsertType = 2
	insert_data.InsertTime = carbon.DateTime{Carbon: carbon.Now()}
	for _, v := range res2 {
		if v.ID == res_id {
			insert_data2.ResultType = uint8(v.Type)
			insert_data2.ResultInfo = v.Info
			break
		}
	}
	insert_data2.Pkid = insert_data.Kid
	err = ctl.WorkParamsDao.Db.Transaction(func(tx *gorm.DB) error {
		if err := ctl.WorkDatasV3Dao.Insert(tx, []*model.DatasV3{insert_data}); err != nil {
			resdata["info"] = err.Error() + "1"
			ctx.JSON(http.StatusBadRequest, resdata)
			return err
		}
		if err := ctl.WorkDataResultDao.Insert(tx, []*model.Dataresult{insert_data2}); err != nil {
			resdata["info"] = err.Error() + "2"
			ctx.JSON(http.StatusBadRequest, resdata)
			return err
		}
		return nil
	})
	if err != nil {
		resdata["info"] = "数据保存失败！"
		ctx.JSON(http.StatusBadRequest, resdata)
		return
	}
	resdata["status"] = 1
	resdata["pid"] = insert_data.Kid
	resdata["dataid"] = insert_data.ID
	resdata["info"] = "数据保存成功！"
	ctx.JSON(http.StatusOK, resdata)
}
func (ctl *IndexController) SaveImages(ctx *gin.Context) {
	file, err := ctx.FormFile("file")
	imgobj := &model.RecordV3{}
	pid := strings.Trim(ctx.PostForm("pid"), " ")
	if len(pid) != 64 {
		ctx.JSON(200, gin.H{
			"status": 0,
			"info":   "pid格式错误",
		})
		return
	}
	result_id, err := strconv.Atoi(ctx.PostForm("result_id"))
	if err != nil {
		ctx.JSON(200, gin.H{
			"status": 0,
			"info":   err.Error(),
		})
		return
	}
	imgobj.ResultID = result_id
	imgobj.Pkid = pid

	if err != nil {
		ctx.JSON(200, gin.H{
			"status": 0,
			"info":   err.Error(),
		})
		return
	}
	dst_root := filepath.Join(ctl.Viper.GetString("upload.upload_path"), "img")
	imgobj.Filetype = filepath.Ext(file.Filename)
	imgobj.Filesize = int(file.Size)
	imgobj.Filename = uuid.New().String() + imgobj.Filetype
	dst := filepath.Join(dst_root, imgobj.Filename)
	imgobj.Filepath = dst
	if err := ctx.SaveUploadedFile(file, dst); err != nil {
		ctx.JSON(200, gin.H{
			"status": 0,
			"info":   err.Error(),
		})
		return
	}
	if err := ctl.WorkRecordV3Dao.Insert(nil, []*model.RecordV3{imgobj}); err != nil {
		ctx.JSON(200, gin.H{
			"status": 0,
			"info":   err.Error(),
		})
		return
	}
	ctx.JSON(200, gin.H{
		"status": 1,
	})
}

type colParam struct {
	info string
	name string
}

func (ctl *IndexController) makeColMap(pararms []*model.DemoParams, result []model.DemoResult) map[string]*colParam {
	cellarr := make(map[string]*colParam, len(pararms)+len(result))
	var strbuffer strings.Builder
	strbuffer.Grow(64)
	index := 1
	for _, v := range pararms {
		strbuffer.WriteString(v.Info)
		if v.Unit != "" {
			strbuffer.WriteRune('(')
			strbuffer.WriteString(v.Unit)
			strbuffer.WriteRune(')')
		}
		cellarr[common.GetExcelColName(index)] = &colParam{info: strbuffer.String(), name: v.Name}
		strbuffer.Reset()
		index++
	}
	for _, v := range result {
		strbuffer.WriteString(strconv.Itoa(v.ID))
		strbuffer.WriteRune('-')
		strbuffer.WriteString(strconv.Itoa(int(v.Type)))
		cellarr[common.GetExcelColName(index)] = &colParam{info: v.Info, name: strbuffer.String()}
		strbuffer.Reset()
		index++
	}
	return cellarr
}
func (ctl *IndexController) setDownloadHeader(ctx *gin.Context) {
	ctx.Header("Pragma", "public")
	ctx.Header("Expires", "0")
	ctx.Header("Cache-Control", "max-age=-1")
	ctx.Header("Content-Type", "application/octet-stream")
	ctx.Header("Content-Disposition", "attachment;filename="+xlsxfilename)
}
func (ctl *IndexController) ExplodeExcel(ctx *gin.Context) {
	paramsarr := ctl.WorkParamsDao.GetData()
	resultarr := ctl.WorkResultDao.GetData()
	col_map := ctl.makeColMap(paramsarr, resultarr)
	f := excelize.NewFile()
	defer func() {
		if err := f.Close(); err != nil {
			log.Println(err.Error())
		}
	}()
	defaultsheet := "Sheet1"
	rowHeight := 26.0
	CustomHeight := true
	err := f.SetSheetProps(defaultsheet, &excelize.SheetPropsOptions{
		DefaultRowHeight: &rowHeight,
		CustomHeight:     &CustomHeight,
	})
	if err != nil {
		ctl.Server.Error("导出附件失败：", err.Error())
		return
	}
	err = f.SetRowHeight(defaultsheet, 1, 20.0)
	if err != nil {
		ctl.Server.Error("导出附件失败：", err.Error())
		return
	}

	col_buf := strings.Builder{}
	col_buf.Grow(4)
	var wg sync.WaitGroup
	richtext := []excelize.RichTextRun{
		{
			Font: &excelize.Font{
				Bold:  true,
				Color: "e6071a",
			},
			Text: "",
		},
	}
	row_style, _ := f.NewStyle(&excelize.Style{
		Font: &excelize.Font{
			Bold: true,
		},
	})
	col_style, _ := f.NewStyle(&excelize.Style{
		Alignment: &excelize.Alignment{
			WrapText: true,
		},
		Font: &excelize.Font{
			Color: "e6071a",
		},
	})
	err = f.SetRowStyle(defaultsheet, 1, 1, row_style)
	if err != nil {
		ctl.Server.Error("导出附件失败：", err.Error())
		return
	}
	err = f.SetColStyle(defaultsheet, "A:K", col_style)
	if err != nil {
		ctl.Server.Error("导出附件失败：", err.Error())
		return
	}
	for key, colname := range col_map {
		col_buf.WriteString(key)
		col_buf.WriteString("1")
		colname_str := col_buf.String()
		colnum := 3.0
		if strings.ContainsRune(colname.info, '(') {
			colnum = 2.1
		}
		width := float64(utf8.RuneCountInString(colname.info)) * colnum
		wg.Add(2)
		go func(cstr, ckey string, colobj *colParam) {
			err2 := f.SetCellValue(defaultsheet, cstr, colobj.info)
			if err2 != nil {
				ctl.Server.Error("导出附件失败：", err2.Error())
				return
			}
			err2 = f.SetCellValue(defaultsheet, ckey+"2", colobj.name)
			if err2 != nil {
				ctl.Server.Error("导出附件失败：", err2.Error())
				return
			}
			wg.Done()
		}(colname_str, key, colname)
		go func(k string, w float64) {
			err2 := f.SetColWidth(defaultsheet, k, k, w)
			if err2 != nil {
				ctl.Server.Error("导出附件失败：", err2.Error())
				return
			}
			wg.Done()
		}(key, width)
		wg.Wait()
		richtext[0].Text = colname.info
		col_buf.Reset()
	}
	err = f.SetRowVisible(defaultsheet, 2, false)
	if err != nil {
		ctl.Server.Error("导出附件失败：", err.Error())
		return
	}
	ctl.setDownloadHeader(ctx)
	err = f.Write(ctx.Writer)
	if err != nil {
		ctl.Server.Error("导出附件失败：", err.Error())
		return
	}
}

func (ctl *IndexController) ImportExcel(ctx *gin.Context) {
	res_data := gin.H{
		"status": 0,
		"msg":    "",
	}
	excelfile, err := ctx.FormFile("file")
	if err != nil {
		res_data["msg"] = "上传excel格式错误"
		ctx.JSON(400, res_data)
		return
	}
	filestream, err := excelfile.Open()
	if err != nil {
		res_data["msg"] = "打开excel失败"
		ctx.JSON(500, res_data)
		return
	}
	defer filestream.Close()
	excelize_obj, err := excelize.OpenReader(filestream, excelize.Options{})
	if err != nil {
		res_data["msg"] = "解析excel失败"
		ctx.JSON(500, res_data)
		return
	}
	defer excelize_obj.Close()
	sheet_name := excelize_obj.GetSheetName(excelize_obj.GetActiveSheetIndex())
	rows, err := excelize_obj.Rows(sheet_name)
	if err != nil {
		res_data["msg"] = "读取行数据失败"
		ctx.JSON(500, res_data)
		return
	}
	if err := ctl.batchInsert(rows); err != nil {
		res_data["msg"] = err.Error()
		ctx.JSON(200, res_data)
		return
	}
	res_data["status"] = 1
	ctx.JSON(200, res_data)
}
func (ctl *IndexController) batchInsert(rows *excelize.Rows) error {
	rows.Next()
	work_data_arr := make([]*model.DatasV3, 0, 200)
	work_dataresult_arr := make([]*model.Dataresult, 0, 200)
	work_record_arr := make([]*model.RecordV3, 0, 200)
	info_arr, err := rows.Columns()
	if err != nil {
		return err
	}
	rows.Next()
	name_arr, err := rows.Columns()
	if err != nil {
		return err
	}
	insert_time := carbon.DateTime{Carbon: carbon.Now()}
	for rows.Next() {
		col_arr, err := rows.Columns()
		if err != nil {
			return errors.New("获取行数据失败！")
		}
		wdv3 := ctl.makeWorkData(col_arr, insert_time)
		work_data_arr = append(work_data_arr, wdv3)
		wres, wrecord := ctl.makeWorkDateResult(wdv3.Kid, col_arr, name_arr, info_arr)
		work_dataresult_arr = append(work_dataresult_arr, wres...)
		work_record_arr = append(work_record_arr, wrecord...)
	}
	err = ctl.WorkParamsDao.Db.Transaction(func(tx *gorm.DB) error {
		if len(work_data_arr) > 0 {
			if err := ctl.WorkDatasV3Dao.Insert(tx, work_data_arr); err != nil {
				return err
			}
		}
		if len(work_dataresult_arr) > 0 {
			if err := ctl.WorkDataResultDao.Insert(tx, work_dataresult_arr); err != nil {
				return err
			}
		}
		if len(work_record_arr) > 0 {
			if err := ctl.WorkRecordV3Dao.Insert(tx, work_record_arr); err != nil {
				return err
			}
		}
		return nil
	})
	if err != nil {
		return err
	}
	return nil
}
func (ctl *IndexController) makeWorkDateResult(pkid string, colarr, namearr, infoarr []string) ([]*model.Dataresult, []*model.RecordV3) {
	wdatares := make([]*model.Dataresult, 0, 16)
	wrecord := make([]*model.RecordV3, 0, 16)
	for i := 11; i < 25; i++ {
		res_type := strings.Split(namearr[i], "-")
		if val := colarr[i]; val != "" {
			rtype, _ := strconv.Atoi(res_type[1])
			wdr := &model.Dataresult{
				Pkid:       pkid,
				ResultID:   res_type[0],
				ResultType: uint8(rtype),
				ResultInfo: infoarr[i],
			}
			switch res_type[1] {
			//图片
			case "1":
				rsid, _ := strconv.Atoi(res_type[0])
				wrarr := ctl.makeWorkRecord(val, pkid, rsid)
				wrecord = append(wrecord, wrarr...)
			//数值 ,结果文件
			case "2", "3":
				wdr.ResultValue = val
			default:
				log.Println("未知类型")
				return nil, nil
			}
			wdatares = append(wdatares, wdr)
		}
	}
	return wdatares, wrecord
}
func (ctl *IndexController) makeWorkRecord(imgstr string, pkid string, resultid int) []*model.RecordV3 {
	wrecord := make([]*model.RecordV3, 0, 32)
	root_img_path, err := os.ReadFile(imgdirpath)
	if err != nil {
		log.Println("图片地址读取失败")
		return nil
	}
	imgstr = strings.ReplaceAll(imgstr, "，", ",")
	imag_arr := strings.Split(imgstr, ",")
	for _, imgname := range imag_arr {
		img_src_path := string(root_img_path) + "/" + imgname
		if common.FileExists(img_src_path) {
			img_ext := filepath.Ext(imgname)
			img_name := uuid.New().String() + img_ext
			img_dst_path := ctl.Viper.GetString("upload.upload_path") + "/img/" + img_name
			wr := &model.RecordV3{
				Pkid:     pkid,
				ResultID: resultid,
				Status:   true,
				Filename: img_name,
				Filetype: img_ext,
				Filepath: img_dst_path,
			}
			imginfo, _ := os.Stat(img_src_path)
			wr.Filesize = int(imginfo.Size())
			//拷贝文件
			src_img, err := os.OpenFile(img_src_path, os.O_RDONLY, 0)
			if err != nil {
				log.Println("打开源图片失败:", err.Error())
			}
			defer src_img.Close()
			dst_img, err := os.OpenFile(img_dst_path, os.O_WRONLY|os.O_CREATE, 0)
			if err != nil {
				log.Println("生成目标图片失败:", err.Error())
			}
			defer dst_img.Close()
			io.Copy(dst_img, src_img)
			wrecord = append(wrecord, wr)
		}
	}
	return wrecord
}
func (ctl *IndexController) makeWorkData(colarr []string, insert_time carbon.DateTime) *model.DatasV3 {
	workdata := &model.DatasV3{
		Zhuansu:  colarr[0],
		Qingjiao: colarr[1],
		Hxfxll:   colarr[2],
		Zsfxll:   colarr[3],
		Plfxll:   colarr[4],
		Xlwd:     colarr[5],
		Ltkcxl:   colarr[6],
		Qtcxl:    colarr[7],
		Cxwd:     colarr[8],
		Lcckbz:   colarr[9],
		Qwsrgl:   colarr[10],
	}
	workdata.Kid = common.HashKey(workdata)
	workdata.InsertType = 1
	workdata.InsertTime = insert_time
	return workdata
}
