package database

import (
	"fmt"
	"testing"
)

func TestTopic_AddTopic(t *testing.T) {
	fmt.Println("test add topic")
	topic:=&Topic{}
	topic.AddTopic()
}

func TestTopic_GetTopicById(t *testing.T) {
	fmt.Println("test search")
	topic:=Topic{
		ID: 2,
	}
	u,_:=topic.GetTopicById()
	fmt.Println("get topic",u)
}

func TestTopic_GetTopics(t *testing.T) {
	fmt.Println("test all records")
	topic:=&Topic{}
	tc,_:=topic.GetTopics()
	for k, v := range tc {
		fmt.Printf("%v个帖子是%v\n",k+1,v)
	}
}

func TestTopic_UpdateTopic(t *testing.T) {
	fmt.Println("update topic")
	topic:=&Topic{
		ID: 2,
		Title: "world",
		Userid: "1234",
	}
	newTitle:="Hello"
	topic.UpdateTopic(newTitle)
}

func TestTopic_DeleteTopic(t *testing.T) {
	fmt.Println("delete topic")
	topic:=&Topic{
		Title: "Hello",
	}
	topic.DeleteTopic()
}