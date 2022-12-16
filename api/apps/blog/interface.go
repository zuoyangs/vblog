package blog

import "context"

type Service interface {

	//创建文章
	CreateBlog(context.Context, *CreateBlogRequest) (*Blog, error)

	//更新文章
	UpdateBlog(context.Context, *CreateBlogRequest) (*Blog, error)

	//删除文章
	DeleteBlog(context.Context, *DeleteBlogRequest) (*Blog, error)

	//文章列表
	QueryBlog(context.Context, *QueryBlogRequest) (*Blog, error)

	//文章详情
	DescribeBlog(context.Context, *DescribeBlogRequest) (*Blog, error)

	//更新文章的状态
	UpdateBlogStatus(context.Context, *UpdateBlogStatusRequest)(*Blog, error)
}
