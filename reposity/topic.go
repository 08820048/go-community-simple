package reposity

import "sync"

// Topic 话题数据结构体
type Topic struct {
	Id         int64  `json:"id"`
	Title      string `json:"title"`
	Content    string `json:"content"`
	CreateTime int64  `json:"create_time"`
}

type TopicDao struct {
}

var (
	topicDao  *TopicDao
	topicOnce sync.Once
)

// NewTopicDaoInstance 有了内存索引，直接根据查询key获得对应的value就可以了，这里用到了一个sync.once，主要用于适用高并发场景下只执行一次
// 这里基于once的实现模式就是平常说的单例模式，减少存储的浪费。
func NewTopicDaoInstance() *TopicDao {
	topicOnce.Do(
		func() {
			topicDao = &TopicDao{}
		})
	return topicDao
}

func (*TopicDao) QueryTopicById(id int64) *Topic {
	return topicIndexMap[id]
}
