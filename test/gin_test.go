package test

import (
	"fmt"
	"net/http"
	"sync"
	"testing"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type my interface {
	get() error
	set() (int, error)
}

type my2 interface {
	get()
	set()
}
type my3 struct {
	my
	my2
}

func test() []string {
a:
	goto a
}
func TestGin(t *testing.T) {
	r := gin.Default()

	// gin.H 是 map[string]interface{} 的一种快捷方式
	r.GET("/someJSON", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "hey", "status": http.StatusOK})
	})

	r.GET("/moreJSON", func(c *gin.Context) {
		// 你也可以使用一个结构体 omitempty
		type msg struct {
			Name    string `json:"name,omitempty"`
			Message string `json:"message,omitempty"`
			Number  int    `json:"number,omitempty"`
		}
		msg_arr := []msg{}
		msg1 := msg{Name: "张三", Message: "法国", Number: 22}
		msg2 := msg{Name: "李四", Message: "德国", Number: 0}
		msg_arr = append(msg_arr, msg1, msg2)
		// 注意 msg.Name 在 JSON 中变成了 "user"
		// 将输出：{"user": "Lena", "Message": "hey", "Number": 123}
		c.JSON(http.StatusOK, msg_arr)
	})

	r.GET("/someXML", func(c *gin.Context) {
		c.XML(http.StatusOK, gin.H{"message": "hey", "status": http.StatusOK})
	})

	r.GET("/someYAML", func(c *gin.Context) {
		c.YAML(http.StatusOK, gin.H{"message": "hey", "status": http.StatusOK})
	})

	// 监听并在 0.0.0.0:8080 上启动服务
	r.Run(":8080")
}

type db struct {
	DB   *gorm.DB
	once sync.Once
}

func TestSome(t *testing.T) {
	var ss *db
	if ss == nil {
		fmt.Println("空")
	}
	fmt.Println("ss")
}
