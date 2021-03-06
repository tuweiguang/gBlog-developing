package admin

import (
	"fmt"
	"gBlog/models"
	"gBlog/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type ArticleCtl struct{}

var Recommend = map[int]string{0: "否", 1: "是"}

func (a *ArticleCtl) List(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1")) // 页数

	offset := (page - 1) * LIMIT // 偏移量
	start := c.Query("start")
	end := c.Query("end")
	status, _ := strconv.Atoi(c.Query("status"))
	title := c.Query("title")

	fmt.Printf("==========================> list, page:%v,start:%v,end:%v,status:%v,title:%v", page, start, end, status, title)
	res := gin.H{}
	res["Start"] = start
	res["End"] = end
	res["Status"] = status
	res["Title"] = title

	some := models.GetSomeArticle(offset, LIMIT)
	var users []int64
	var categorys []int
	for _, v := range some {
		if !utils.IsExistsElementInt64(users, v.UserId) {
			users = append(users, v.UserId)
		}

		if !utils.IsExistsElementInt(categorys, v.CategoryId) {
			categorys = append(categorys, v.CategoryId)
		}
	}
	someUser := models.GetSomeUserByIds(users)
	someCate := models.GetCategoryById(categorys)
	for i, v := range some {
		for _, vv := range someUser {
			if v.UserId == vv.Id {
				some[i].User = vv
				break
			}
		}

		for _, vv := range someCate {
			if v.CategoryId == vv.Id {
				some[i].Category = vv
				break
			}
		}
	}

	res["Data"] = some
	res["Paginator"] = models.GenPaginator(page, LIMIT, len(some))
	res["StatusText"] = Status
	res["RecommendText"] = Recommend
	c.HTML(http.StatusOK, "article-list.html", res)
}

func (a *ArticleCtl) Add(c *gin.Context) {
	cate := models.GetAllCategory()

	res := gin.H{}
	res["Category"] = cate
	c.HTML(http.StatusOK, "article-add.html", res)
}

func (a *ArticleCtl) Upload(c *gin.Context) {
	if c.Request.Method == "POST" {
		var info UploadArticle

		err := c.ShouldBind(&info)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"status":  "error",
				"message": "参数错误!",
			})
			return
		}

		// 单个文件
		file, err := c.FormFile("file")
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"status":  "error",
				"message": "获取文件失败！",
			})
			return
		}

		//dst := fmt.Sprintf("./%s", file.Filename)
		// 上传文件到指定的目录
		//err = c.SaveUploadedFile(file, dst)
		//if err != nil {
		//	c.JSON(http.StatusOK, gin.H{
		//		"status":  "error",
		//		"message": "上传文件失败！",
		//	})
		//	return
		//}
		content, err := utils.ReadUploaded(file)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"status":  "error",
				"message": "读取文件失败",
			})
		}
		models.CreateArticle(1, info.Title, info.Cate, info.Tag, string(content))
		c.JSON(http.StatusOK, gin.H{
			"status":  "ok",
			"message": "添加文章成功！",
		})
	}
}
