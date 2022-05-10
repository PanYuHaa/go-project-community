package repository

import (
	"encoding/json"
	"os"
	"sync"
)

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

func NewPostDaoInstance() *PostDao {
	postOnce.Do(
		func() {
			postDao = &PostDao{}
		})
	return postDao
}
func (*PostDao) QueryPostsByParentId(parentId int64) []*Post {
	return postIndexMap[parentId]
}

func (*PostDao) InsertPost(post *Post) error {
	f, err := os.OpenFile("./data/post", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		return err
	}

	defer f.Close() // 关闭文件
	marshal, _ := json.Marshal(post) // 使传入的那条post序列化
	if _, err = f.WriteString(string(marshal)+"\n"); err != nil {
		return err
	}

	rwMutex.Lock()
	postList, ok := postIndexMap[post.ParentId] // postMap中有postList数组吗？
	if !ok {
		postIndexMap[post.ParentId] = []*Post{post} // 如果没有就建立一个post数组
	} else {
		postList = append(postList, post)	// 如果存在的话就在对应的List里面添加新来的post
		postIndexMap[post.ParentId] = postList // 更新对应话题id下的postList
	}
	rwMutex.Unlock()
	return nil
}
