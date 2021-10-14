package database
import (
	"database/sql"
	"fmt"
	_"github.com/go-sql-driver/mysql"
	"test/_02_basic/_01_escape_char/_99_webapp/_05_database/utils"
)

var(
	Db *sql.DB
	err error
)

func init()  {
	//用户名：密码@tcp(端口号)/数据库名
	Db,err=sql.Open("mysql","root:root@tcp(localhost:3306)/test")
	if err!=nil{
		panic(err.Error())
	}
	fmt.Println("success")
}

type Topic struct {
	ID       int
	Title  string
	Userid string
}

func (topic *Topic)AddTopic() error {
	sqlStr:="insert into topics(title,userid) values(?,?)"
	_,err2:=Db.Exec(sqlStr,topic.Title,topic.Userid)
	if err2!=nil{
		fmt.Println("exec error",err2)
		return err2
	}
	return nil
}

func (topic *Topic)DeleteTopic() error {
	sqlStr:="delete from topics where title = ?"
	_,err2:=utils.Db.Exec(sqlStr,topic.Title)
	if err2!=nil{
		fmt.Println("exec error",err2)
		return err2
	}
	return nil
}


func (topic *Topic)UpdateTopic(newTitle string) error {
	sqlStr:="update topics set title=?,userid=? where title = ?"
	_,err2:=utils.Db.Exec(sqlStr,newTitle,topic.Userid,topic.Title)
	if err2!=nil{
		fmt.Println("exec error",err2)
		return err2
	}
	return nil
}

func (topic *Topic) GetTopicById() (*Topic,error) {
	sqlStr:="select id,title,userid from topics where id = ?"
	row:=utils.Db.QueryRow(sqlStr,topic.ID)
	var id int
	var title string
	var userid string
	err:=row.Scan(&id,&title,&userid)
	if err!=nil{
		return nil,err
	}
	u:=&Topic{
		ID:     id,
		Title:  title,
		Userid: userid,
	}
	return u,nil
}

func (topic *Topic) GetTopics()  ([]*Topic,error){
	sqlStr:="select id,title,userid from topics"
	rows,err:=utils.Db.Query(sqlStr)
	if err!=nil{
		return nil,err
	}
	//创建切片
	var topics []*Topic
	for rows.Next() {
		var id int
		var title string
		var userid string
		err:=rows.Scan(&id,&title,&userid)
		if err!=nil{
			return nil,err
		}
		u:=&Topic{
			ID:     id,
			Title:  title,
			Userid: userid,
		}
		topics= append(topics, u)
	}
	return topics,nil
}