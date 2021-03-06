package models

import (
	"fmt"
	"gBlog/common/db"
	"gBlog/common/log"
	"math"
	"time"
)

type Article struct {
	Id         uint `gorm:"primary_key;column:id"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  *time.Time `sql:"index"`
	UserId     int64      `gorm:"column:user_id"`
	Title      string     `gorm:"column:title"`
	CategoryId int        `gorm:"column:category_id"`
	Tag        string     `gorm:"column:tag"`
	Content    string     `gorm:"column:content"`
	Status     int        `gorm:"column:status"`
	Pv         int        `gorm:"column:pv"` //每篇文章的点击量
	Review     int        `gorm:"column:review"`
	Recommend  int        `gorm:"column:recommend"`
	Like       int        `gorm:"column:like"`
	User       *User
	Category   *Category
}

func (Article) TableName() string {
	return "article"
}

type Paginator struct {
	CurrentPage int //当前页
	PageSize    int //每页数量
	TotalPage   int //总页数
	TotalCount  int //总数量
}

func GenPaginator(page, limit, count int) Paginator {
	var paginator Paginator
	paginator.TotalCount = count
	paginator.TotalPage = int(math.Ceil(float64(count) / float64(limit)))
	paginator.PageSize = limit
	paginator.CurrentPage = page
	return paginator
}

func GetSomeArticle(offset, limit int) []*Article {
	some := make([]*Article, limit)
	err := db.GetMySQL().Raw("SELECT * FROM article as b1 inner join (select id from article limit ?,?) as b2 on b1.id = b2.id", offset, limit).Scan(&some).Error
	if err != nil {
		log.GetLog().Error(fmt.Sprintf("GetSomeArtcle offset:%v limit:%v error:%v", offset, limit, err))
	}
	return some
}

func GetSomeArticleByTag(offset, limit int, tag string) []*Article {
	some := make([]*Article, limit)
	err := db.GetMySQL().Raw("SELECT * FROM article as b1 inner join (select id from article limit ?,?) as b2 on b1.id = b2.id where b1.tag = ?", offset, limit, tag).Scan(&some).Error
	if err != nil {
		log.GetLog().Error(fmt.Sprintf("GetSomeArtcle offset:%v limit:%v error:%v", offset, limit, err))
	}
	return some
}

func GetSomeArticleByDate(offset, limit int, date string) []*Article {
	some := make([]*Article, limit)
	err := db.GetMySQL().Raw("SELECT * FROM article as b1 inner join (select id from article limit ?,?) as b2 on b1.id = b2.id where date_format(b1.created_at, '%Y-%m') = ?", offset, limit, date).Scan(&some).Error
	if err != nil {
		log.GetLog().Error(fmt.Sprintf("GetSomeArtcle offset:%v limit:%v error:%v", offset, limit, err))
	}
	return some
}

func GetAllArticle() []*Article {
	all := make([]*Article, 0)
	err := db.GetMySQL().Find(&all).Error
	if err != nil {
		log.GetLog().Error(fmt.Sprintf("GetAllArticle error:%v", err))
	}
	return all
}

func GetArticle(id int) []*Article {
	//var article []*Article
	article := new(Article)
	err := db.GetMySQL().Where("id = ?", id).First(article).Error
	if err != nil {
		log.GetLog().Error(fmt.Sprintf("GetArticle error:%v", err))
	}
	a := make([]*Article, 0)
	a = append(a, article)
	return a
}

func CreateArticle(userId int64, title string, categoryId int, tag string, content string) {
	article := Article{
		UserId:     userId,
		Title:      title,
		CategoryId: categoryId,
		Tag:        tag,
		Content:    content,
	}

	result := db.GetMySQL().Create(&article)
	if result.Error != nil {
		log.GetLog().Error(fmt.Sprintf("CreateArticle error:%v", result.Error))
	}
}

func AddArticleByPV(id uint, pv int) {
	article := new(Article)
	article.Id = id
	err := db.GetMySQL().Model(article).Update("pv", pv).Error
	if err != nil {
		log.GetLog().Error(fmt.Sprintf("AddArticleByPV error:%v", err))
	}
}
