package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"test/_02_basic/_01_escape_char/xingchen/main_04_01/database"
)

//// Topic 帖子
//type Topic struct {
//	ID int
//	Title string
//}
////数据源，类似MySQL中的数据

var topics []database.Topic

func handleGet(w http.ResponseWriter, r *http.Request){
	switch r.Method {
	case "GET":
		topic:=&database.Topic{}
		tc,_:=topic.GetTopics()
		topics,err:=json.Marshal(tc)
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

	//获取指定请求参数
	topicParas, ok := query["id"]
	if !ok || len(topicParas[0]) < 1 {
		log.Println("Url Param 'topic' is missing")
		return
	}
	id, _ := strconv.Atoi(topicParas[0])
	topic:=database.Topic{
		ID: id,
	}
	u,_:=topic.GetTopicById()
	w.WriteHeader(200)
	fmt.Fprintf(w, "id="+topicParas[0]+"\n")
	topics,_:=json.Marshal(u)
	w.Write(topics)

}

func handlePost(w http.ResponseWriter, r *http.Request){
	switch r.Method {
	case "POST":
		topic:=&database.Topic{
			ID:    len(topics)+1,
			Title: r.FormValue("title"),
			Userid: r.FormValue("userid"),
		}
		topic.AddTopic()
		w.WriteHeader(200)
		res,_:=topic.GetTopics()
		topics,_:=json.Marshal(res)
		w.Write(topics)
	default:
		w.WriteHeader(404)
		fmt.Fprint(w,"{\"message\": \"not found\"}")
	}
}

func handlePut(w http.ResponseWriter, r *http.Request){
	title:=r.FormValue("title")
	userid:=r.FormValue("userid")
	newTitle:=r.FormValue("newTitle")
	topic:=&database.Topic{
		Title: title,
		Userid: userid,
	}
	w.WriteHeader(200)
	topic.UpdateTopic(newTitle)
	res,_:=topic.GetTopics()
	topics,_:=json.Marshal(res)
	w.Write(topics)
}

func handleDelete(w http.ResponseWriter, r *http.Request){
	title:=r.FormValue("title")
	topic:=&database.Topic{
		Title: title,
	}
	topic.DeleteTopic()
	w.WriteHeader(200)
	res,_:=topic.GetTopics()
	topics,_:=json.Marshal(res)
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

