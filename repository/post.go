package repository

import "sync"

type Post struct {
	Id         int64  `json:"id"`
	ParentId   int64  `json:"parent_id"`
	Content    string `json:"content"`
	CreateTime int64  `json:"create_time"`
}
type PostDao struct {
}

var (
	postDao  *PostDao
	postOnce sync.Once
)

// NewPostDaoInstance 并发模式下的单例
func NewPostDaoInstance() *PostDao {
	postOnce.Do(
		func() {
			postDao = &PostDao{}
		})
	return postDao
}

// QueryPostByParentId 查询函数
func (*PostDao) QueryPostByParentId(parentId int64) []*Post {
	return postIndexMap[parentId]
}
