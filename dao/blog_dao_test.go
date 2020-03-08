package dao

import (
	"context"
	"github.com/gpmgo/gopm/modules/log"
	"testing"
)

func TestBlogDao_GetBlogs(t *testing.T) {
	res, err := BlogdaoOnceInstance().GetBlogs(context.Background(), 1, 1)
	if err != nil {
		log.Error("get BLog failed%s", err)
	}
	print("rs:%s", res)
}
