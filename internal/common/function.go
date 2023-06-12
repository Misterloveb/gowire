package common

import (
	"crypto/sha256"
	"errors"
	"fmt"
	"gindemo/internal/model"
	"os"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

func HashKey(src *model.WorkDatasV3) string {
	stringbuf := strings.Builder{}
	stringbuf.Grow(100)
	stringbuf.WriteString(src.Zhuansu)
	stringbuf.WriteString(src.Qingjiao)
	stringbuf.WriteString(src.Hxfxll)
	stringbuf.WriteString(src.Zsfxll)
	stringbuf.WriteString(src.Plfxll)
	stringbuf.WriteString(src.Xlwd)
	stringbuf.WriteString(src.Ltkcxl)
	stringbuf.WriteString(src.Qtcxl)
	stringbuf.WriteString(src.Cxwd)
	stringbuf.WriteString(src.Lcckbz)
	stringbuf.WriteString(src.Qwsrgl)
	// h := sha256.New()
	res := sha256.Sum256([]byte(stringbuf.String()))
	return fmt.Sprintf("%x", res)
}
func DelFile(filepath string) bool {
	if FileExists(filepath) {
		err := os.Remove(filepath)
		if err != nil {
			return false
		}
	}
	return true
}
func FileExists(filepath string) bool {
	_, err := os.Stat(filepath)
	return !errors.Is(err, os.ErrNotExist)
}
func PostInt(key string, ctx *gin.Context) (int, error) {
	return strconv.Atoi(ctx.PostForm(key))
}
func PostInt8(key string, ctx *gin.Context) (int8, error) {
	res, err := strconv.Atoi(ctx.PostForm(key))
	if err != nil {
		return 0, err
	}
	return int8(res), nil
}
func GetExcelColName(index int) string {
	res := ""
	for index > 0 {
		index--
		res = string(rune(index%26)+'A') + res
		index /= 26
	}
	return res
}
