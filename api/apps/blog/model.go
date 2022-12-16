package blog

import (
	"encoding/json"
	"net/http"
	"strconv"
)

type Blog struct {
	// 文章Id
	Id int `json:"id"`
	// 创建时间
	CreateAt int64 `json:"create_at"`
	// 更新时间
	UpdateAt int64 `json:"update_at"`
	// 发布时间
	PublishAt int64 `json:"publish_at"`
	// 用户提交数据
	*CreateBlogRequest
	// 文章状态 草稿/发布
	Status Status `json:"status"`
	// 博客标签
	Tags []*tag.Tag `json:"tags"`
}

func (b *Blog) String() string {
	dj, _ := json.Marshal(b)
	return string(dj)
}

type BlogSet struct {
	// 总条目个数, 用于前端分页
	Total int64 `json:"total"`
	// 列表数据
	Items []*Blog `json:"items"`
}

func NewBlogSet() *BlogSet {
	return &BlogSet{
		Items: []*Blog{},
	}
}

func (b *BlogSet) String() string {
	dj, _ := json.Marshal(b)
	return string(dj)
}

func NewCreateBlogRequest() *CreateBlogRequest {
	return &CreateBlogRequest{}
}

type CreateBlogRequest struct {
	// 文章摘要信息,通过提前Content内容获取
	Summary string `json:"summary" validate:"required"`
	// 文章图片
	TitleImg string `json:"title_img"`
	// 文章标题
	TitleName string `json:"title_name" validate:"required"`
	// 文章副标题
	SubTitle string `json:"sub_title"`
	// 文章内容
	Content string `json:"content" validate:"required"`
	// 文章作者
	Author string `json:"author"`
}

// 创建对象的校验
func (req *CreateBlogRequest) Validate() error {
	return validate.Struct(req)
}

func NewPutUpdateBlogRequest(id int) *UpdateBlogRequest {
	return &UpdateBlogRequest{
		BlogId:            id,
		UpdateMode:        UPDATE_MODE_PUT,
		CreateBlogRequest: NewCreateBlogRequest(),
	}
}

func NewPatchUpdateBlogRequest(id int) *UpdateBlogRequest {
	return &UpdateBlogRequest{
		BlogId:            id,
		UpdateMode:        UPDATE_MODE_PATCH,
		CreateBlogRequest: NewCreateBlogRequest(),
	}
}

type UpdateBlogRequest struct {
	BlogId     int
	UpdateMode UpdateMode
	*CreateBlogRequest
}

func NewDeleteBlogRequest(id int) *DeleteBlogRequest {
	return &DeleteBlogRequest{Id: id}
}

type DeleteBlogRequest struct {
	Id int
}

func NewQueryBlogRequest() *QueryBlogRequest {
	return &QueryBlogRequest{
		PageSize:   20,
		PageNumber: 1,
	}
}

// http query string: ?keywords=b&page_size=20&page_number=1&status=published
func NewQueryBlogRequestFromHTTP(r *http.Request) (*QueryBlogRequest, error) {
	qs := r.URL.Query()

	req := NewQueryBlogRequest()
	req.Keywords = qs.Get("keywords")

	psStr := qs.Get("page_size")
	if psStr != "" {
		req.PageSize, _ = strconv.Atoi(psStr)
	}
	pnStr := qs.Get("page_number")
	if pnStr != "" {
		req.PageNumber, _ = strconv.Atoi(pnStr)
	}

	// 获取状态过滤参数
	status := qs.Get("status")
	if status != "" {
		status, err := ParseStatusFromString(status)
		if err != nil {
			return nil, err
		}
		req.Status = &status
	}

	return req, nil
}

type QueryBlogRequest struct {
	PageSize   int
	PageNumber int
	Keywords   string
	// 补充状态过滤参数, 用于web 前台 过滤已经发布的文章
	// 比如过滤 状态为发布的文章
	Status  *Status
	BlogIds []int
}

func (req *QueryBlogRequest) Offset() int {
	return (req.PageNumber - 1) * req.PageSize
}

func NewDescribeBlogRequest(id int) *DescribeBlogRequest {
	return &DescribeBlogRequest{Id: id}
}

type DescribeBlogRequest struct {
	Id int
}

func NewUpdateBlogStatusRequest(id int, status Status) *UpdateBlogStatusRequest {
	return &UpdateBlogStatusRequest{
		Id:     id,
		Status: status,
	}
}

func NewDefaultUpdateBlogStatusRequest() *UpdateBlogStatusRequest {
	return &UpdateBlogStatusRequest{}
}

type UpdateBlogStatusRequest struct {
	// 文章Id
	Id int `json:"id"`
	// 文章状态 草稿/发布
	Status Status `json:"status"`
}
