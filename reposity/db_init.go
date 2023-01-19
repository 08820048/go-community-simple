package reposity

import (
	"bufio"
	"encoding/json"
	"os"
)

var (
	topicIndexMap map[int64]*Topic
	postIndexMap  map[int64][]*Post
)

// Init 整体初始化方法
func Init(filePath string) error {
	//初始化话题数据
	if err := initTopicIndexMap(filePath); err != nil {
		return err
	}
	if err := initPostIndexMap(filePath); err != nil {
		return err
	}
	return nil
}

// 初始化数据话题数据
// 基本逻辑:打开文件，基于file初始化scanner，通过迭代器方式遍历数据行，转为结构体存储至内存map
func initTopicIndexMap(filePath string) error {
	open, err := os.Open(filePath + "topic")
	if err != nil {
		return err
	}
	scanner := bufio.NewScanner(open)
	topicTmpMap := make(map[int64]*Topic)
	for scanner.Scan() {
		text := scanner.Text()
		var topic Topic
		if err := json.Unmarshal([]byte(text), &topic); err != nil {
			return err
		}
		topicTmpMap[topic.Id] = &topic
	}
	topicIndexMap = topicTmpMap
	return nil
}

func initPostIndexMap(filePath string) error {
	open, err := os.Open(filePath + "post")
	if err != nil {
		return err
	}
	scanner := bufio.NewScanner(open)
	postTmpMap := make(map[int64][]*Post)
	for scanner.Scan() {
		text := scanner.Text()
		var post Post
		if err := json.Unmarshal([]byte(text), &post); err != nil {
			return err
		}
		posts, ok := postTmpMap[post.ParentId]
		if !ok {
			postTmpMap[post.ParentId] = []*Post{&post}
			continue
		}
		posts = append(posts, &post)
		postTmpMap[post.ParentId] = posts
	}
	postIndexMap = postTmpMap
	return nil
}
