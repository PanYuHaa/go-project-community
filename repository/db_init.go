package repository

import (
	"bufio"
	"encoding/json"
	"os"
	"sync"
)

// 话题 = Topic；评论 = Post

var (
	topicIndexMap map[int64]*Topic
	postIndexMap  map[int64][]*Post // key是int64，value是[]*Post，Post结构体的数组
	rwMutex       sync.RWMutex
)

func Init(filePath string) error {
	if err := initTopicIndexMap(filePath); err != nil {
		return err
	}
	if err := initPostIndexMap(filePath); err != nil {
		return err
	}
	return nil
}

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
		topicTmpMap[topic.Id] = &topic // 用topic里面的id来当索引
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
		posts, ok := postTmpMap[post.ParentId] // 用话题id来当作postMap的一维索引，这个话题里面有评论了吗
		if !ok {
			postTmpMap[post.ParentId] = []*Post{&post} // 如果没有评论呢就新建一个[]*Post{&post}放入当前post的话题数组
			continue
		}
		posts = append(posts, &post) // 为posts切片来增加元素，posts本身是post数组，相当于temp一样建立一个数组，然后将新来post接上
		postTmpMap[post.ParentId] = posts // 将更新后的posts存入对应的话题id里面
	}
	postIndexMap = postTmpMap
	return nil
}