package repository

import (
"sync"
)

type Topic struct {
	// 相当于TopicDao类里面的一个Topic属性，只不过他也是一个类
	Id         int64  `json:"id"`
	Title      string `json:"title"`
	Content    string `json:"content"`
	CreateTime int64  `json:"create_time"`
}
type TopicDao struct {
}

var (
	topicDao  *TopicDao
	topicOnce sync.Once // 只执行一次，单例模式减少内存浪费
)

func NewTopicDaoInstance() *TopicDao {
	// 使用单例模式创建topicDao
	topicOnce.Do(
		func() {
			topicDao = &TopicDao{} // &struct来相当于new了一个topicDao
//			topicDao = new(TopicDao)
		})
	return topicDao
}
func (*TopicDao) QueryTopicById(id int64) *Topic {
	return topicIndexMap[id]
}
