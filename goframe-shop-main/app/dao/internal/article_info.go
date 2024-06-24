// ==========================================================================
// This is auto-generated by gf cli tool. DO NOT EDIT THIS FILE MANUALLY.
// ==========================================================================

package internal

import (
	"github.com/gogf/gf/database/gdb"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/frame/gmvc"
)

// ArticleInfoDao is the manager for logic model data accessing
// and custom defined data operations functions management.
type ArticleInfoDao struct {
	gmvc.M                     // M is the core and embedded struct that inherits all chaining operations from gdb.Model.
	DB      gdb.DB             // DB is the raw underlying database management object.
	Table   string             // Table is the table name of the DAO.
	Columns articleInfoColumns // Columns contains all the columns of Table that for convenient usage.
}

// ArticleInfoColumns defines and stores column names for table article_info.
type articleInfoColumns struct {
	Id        string //
	UserId    string // 作者id
	Title     string // 标题
	Desc      string // 摘要
	PicUrl    string // 封面图
	IsAdmin   string // 1后台管理员发布 2前台用户发布
	Praise    string // 点赞数
	Detail    string // 文章详情
	CreatedAt string //
	UpdatedAt string //
	DeletedAt string //
}

func NewArticleInfoDao() *ArticleInfoDao {
	return &ArticleInfoDao{
		M:     g.DB("default").Model("article_info").Safe(),
		DB:    g.DB("default"),
		Table: "article_info",
		Columns: articleInfoColumns{
			Id:        "id",
			UserId:    "user_id",
			Title:     "title",
			Desc:      "desc",
			PicUrl:    "pic_url",
			IsAdmin:   "is_admin",
			Praise:    "praise",
			Detail:    "detail",
			CreatedAt: "created_at",
			UpdatedAt: "updated_at",
			DeletedAt: "deleted_at",
		},
	}
}
