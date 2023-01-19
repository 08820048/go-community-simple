package service

import (
	"communityMini/repository"
	"errors"
	"sync"
)

// PageInfo service实体
type PageInfo struct {
	Topic    *repository.Topic
	PostList []*repository.Post
}

func QueryPageInfo(topicId int64) (*PageInfo, error) {
	return NewQueryPageInfoFlow(topicId).Do()
}

func NewQueryPageInfoFlow(topId int64) *QueryPageInfoFlow {
	return &QueryPageInfoFlow{
		topicId: topId,
	}
}

type QueryPageInfoFlow struct {
	topicId  int64
	pageInfo *PageInfo

	topic *repository.Topic
	posts []*repository.Post
}

func (f *QueryPageInfoFlow) Do() (*PageInfo, error) {
	if err := f.checkParam(); err != nil {
		return nil, err
	}
	if err := f.prepareInfo(); err != nil {
		return nil, err
	}
	if err := f.packPageInfo(); err != nil {
		return nil, err
	}
	return f.pageInfo, nil
}

func (f *QueryPageInfoFlow) checkParam() error {
	if f.topicId <= 0 {
		return errors.New("topic id must be larger than 0")
	}
	return nil
}

// 由于话题和回帖信息的获取都依赖topicId,适用等待组以及协程实现并行执行，压榨CPU，提高执行效率，降低接口耗时
func (f *QueryPageInfoFlow) prepareInfo() error {
	//获取topic的信息
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		topic := repository.NewTopicDaoInstance().QueryTopicById(f.topicId)
		f.topic = topic
	}()
	//获取post列表
	go func() {
		defer wg.Done()
		posts := repository.NewPostDaoInstance().QueryPostByParentId(f.topicId)
		f.posts = posts
	}()
	wg.Wait()
	return nil
}

func (f *QueryPageInfoFlow) packPageInfo() error {
	f.pageInfo = &PageInfo{
		Topic:    f.topic,
		PostList: f.posts,
	}
	return nil
}
