package dao

import (
	"blogapi/db"
	"blogapi/model"
	"context"
	log "github.com/sirupsen/logrus"
	"strings"
	"sync"
)

type BlogTypeDao struct{}

var blogtypedao *BlogTypeDao
var BlogTypedaoOnce sync.Once

const (
	blogTypeTableName = "blog_type"
)

func BlogTypedaoOnceInstance() *BlogTypeDao {
	BlogTypedaoOnce.Do(func() {
		blogtypedao = &BlogTypeDao{}
	})
	return blogtypedao
}
func (*BlogTypeDao) AddBlogType(ctx context.Context, req model.AdditionBlog) error {
	types := strings.Split(req.BlogType, ",")
	var err error = nil
	for _, v := range types {
		err = db.WriteDB(ctx).Table(blogTypeTableName).Create(model.BlogType{BlogType: v}).Error
		if err != nil {
			break
		}
	}
	return err
}
func (*BlogTypeDao) GetBlogTypeCategory(ctx context.Context) ([]*model.BlogTypeCategory, error) {
	handle := db.ReadDB(ctx)
	handle = handle.Table(blogTypeTableName)
	handle = handle.Select("count(*),blog_type")
	handle = handle.Group("blog_type")
	handle = handle.Order("count(*) desc")
	rows, err := handle.Rows()
	if err != nil {
		log.Fatal("err%s", err)
		return nil, err
	}
	defer rows.Close()
	var tmpBlogTypeCategorys []*model.BlogTypeCategory
	for rows.Next() {
		var tmpBlogTypeCategory = model.BlogTypeCategory{}
		if err := handle.ScanRows(rows, &tmpBlogTypeCategory); err != nil {
			log.Error(err.Error())
			continue
		}
		tmpBlogTypeCategorys = append(tmpBlogTypeCategorys, &tmpBlogTypeCategory)
	}
	return tmpBlogTypeCategorys, nil
}
