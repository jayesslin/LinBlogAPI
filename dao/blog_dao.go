package dao

import (
	"context"
	"github.com/gpmgo/gopm/modules/log"
	"golangDemo/db"
	"golangDemo/model"
	"sync"
)

type BlogDao struct{}

var blogdao *BlogDao
var BlogdaoOnce sync.Once

func BlogdaoOnceInstance() *BlogDao {
	BlogdaoOnce.Do(func() {
		blogdao = &BlogDao{}
	})
	return blogdao
}
func (*BlogDao) AddBlog(ctx context.Context, req model.AdditionBlog) (model.AdditionBlog, error) {
	NewBlog := &model.AdditionBlog{
		BlogTitle:   req.BlogTitle,
		BlogContent: req.BlogContent,
	}
	err := db.WriteDB(ctx).Table("blog").Create(&NewBlog).Error
	return *NewBlog, err
}
func (*BlogDao) GetBlogs(ctx context.Context, pageNum, pageSize int) ([]*model.Blog, error) {
	handle := db.ReadDB(ctx)
	handle = handle.Table("blog")
	handle = handle.Select("id,blog_title,blog_content,blog_author,create_time,update_time")
	handle = handle.Order("create_time desc")
	//var total int64
	//handle.Count(&total)
	//handle = handle.Offset((pageNum - 1) * 10)
	//handle = handle.Limit(pageSize)
	rows, err := handle.Rows()
	if err != nil {
		log.Fatal("err%s", err)
		return nil, err
	}
	defer rows.Close()
	var blogs []*model.Blog
	for rows.Next() {
		var tmpBlog = model.Blog{}
		if err := handle.ScanRows(rows, &tmpBlog); err != nil {
			log.Error(err.Error())
			continue
		}
		blogs = append(blogs, &tmpBlog)
	}
	return blogs, nil
}
func (*BlogDao) GetSpecifyBlog(ctx context.Context, id int) (*model.Blog, error) {
	handle := db.ReadDB(ctx)
	handle = handle.Table("blog")
	handle = handle.Select("id,blog_title,blog_content,blog_author,create_time,update_time")
	handle = handle.Where("id = ?", id)
	rows, err := handle.Rows()
	if err != nil {
		log.Fatal("err%s", err)
		return nil, err
	}
	defer rows.Close()
	var blogs *model.Blog
	for rows.Next() {
		var tmpBlog = model.Blog{}
		if err := handle.ScanRows(rows, &tmpBlog); err != nil {
			log.Error(err.Error())
			continue
		}
		blogs = &tmpBlog
	}
	return blogs, nil
}
