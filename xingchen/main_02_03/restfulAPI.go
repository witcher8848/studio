package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// Topic 帖子
type Topic struct {
	ID int
	Title string
}
//数据源，类似MySQL中的数据

var topics = []Topic{
	{ID: 1,Title: "昨天"},
	{ID: 2,Title: "今天"},
	{ID: 3,Title: "明天"},
}

func handleGet(w http.ResponseWriter, r *http.Request){
	switch r.Method {
	case "GET":
		topics,err:=json.Marshal(topics)
		if err!=nil {
			w.WriteHeader(500)
			fmt.Fprint(w,"{\"message\": \""+err.Error()+"\"}")
		}else {
			w.WriteHeader(200)
			w.Write(topics)
		}
	default:
		w.WriteHeader(404)
		fmt.Fprint(w,"{\"message\": \"not found\"}")
	}
}

func handleGetOne(w http.ResponseWriter, r *http.Request) {
	//获取所有请求参数
	query := r.URL.Query()
	_, err := json.Marshal(query)
	if err != nil {
		log.Fatalf("%v\n", err)
	}
	//获取指定请求参数
	topicParas, ok := query["topic"]
	if !ok || len(topicParas[0]) < 1 {
		log.Println("Url Param 'topic' is missing")
		return
	}
	index:=-1
	for i, topic := range topics {
		if topicParas[0]==topic.Title{
			index=i
		}
	}
	w.WriteHeader(200)
	fmt.Fprintf(w, "topic="+topicParas[0]+"\n")
	topics,_:=json.Marshal(topics[index])
	w.Write(topics)

}

func handlePost(w http.ResponseWriter, r *http.Request){
	switch r.Method {
	case "POST":
		topic:=Topic{
			ID:    len(topics)+1,
			Title: r.FormValue("title"),
		}
		topics= append(topics, topic)
		w.WriteHeader(200)
		topics,_:=json.Marshal(topics)
		w.Write(topics)
	default:
		w.WriteHeader(404)
		fmt.Fprint(w,"{\"message\": \"not found\"}")
	}
}

func handlePut(w http.ResponseWriter, r *http.Request){
	title:=r.FormValue("title")
	newTitle:=r.FormValue("newTitle")
	for index, topic := range topics {
		if title==topic.Title{
			topics[index].Title=newTitle
		}
	}
	w.WriteHeader(200)
	topics,_:=json.Marshal(topics)
	w.Write(topics)
}

func handleDelete(w http.ResponseWriter, r *http.Request){
	title:=r.FormValue("title")
	index:=-1
	for i, topic := range topics {
		if title==topic.Title{
			index=i
		}
	}
	if index!=-1 {
		topics= append(topics[:index], topics[index+1:]...)
	}
	w.WriteHeader(200)
	topics,_:=json.Marshal(topics)
	w.Write(topics)
}
func main()  {
	http.HandleFunc("/get",handleGet)
	http.HandleFunc("/getOne",handleGetOne)
	http.HandleFunc("/post",handlePost)
	http.HandleFunc("/put",handlePut)
	http.HandleFunc("/delete",handleDelete)
	http.ListenAndServe(":8080",nil)
}

