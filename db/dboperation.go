package db

import (
	"log"
        "gopkg.in/mgo.v2"
        "gopkg.in/mgo.v2/bson"
        "time"
        "fmt"
)

type Person struct {
        Name string
        Phone string
}

func Insertdb() Person {
        session, err := mgo.Dial("localhost")
        if err != nil {
                panic(err)
        }
        defer session.Close()

        // Optional. Switch the session to a monotonic behavior.
        session.SetMode(mgo.Monotonic, true)

        c := session.DB("test").C("people")
        err = c.Insert(&Person{"Ale", "+55 53 8116 9639"},
	               &Person{"Cla", "+55 53 8402 8510"})
        if err != nil {
                log.Fatal(err)
        }

        result := Person{}
        err = c.Find(bson.M{"name": "Ale"}).One(&result)
        if err != nil {
                log.Fatal(err)
        }

        return result
}

type Post struct {
        Id bson.ObjectId `json:"_id"`
        Title string `json:"title"`
        Author string `json:"author"`
        Date time.Time `json:"-"`
        Category string `json:"category"`
        Content string `json:"content"`
        CommentsId []bson.ObjectId `json:"-"`
        Comments []string `json:"comments"`
}

func InsertPost(title,author,category,content string) Post {
        session, err := mgo.Dial("localhost")
        if err != nil {
                panic(err)
        }
        defer session.Close()

        // Optional. Switch the session to a monotonic behavior.
        session.SetMode(mgo.Monotonic, true)

        c := session.DB("blog").C("post")

        id := bson.NewObjectId()

        err = c.Insert(bson.M{
                "_id": id,
                "id":id,
                "title": title,
                "author": author,
                "date": time.Now(),
                "category": category,
                "content": content,    
        })

        
        if err != nil {
                log.Fatal(err)
        }

        result := Post{}
        err = c.Find(bson.M{"_id": id}).One(&result)
        if err != nil {
                log.Fatal(err)
        }

        return result
}

func SelectAll() []Post {
        session, err := mgo.Dial("localhost")
        if err != nil {
                panic(err)
        }
        defer session.Close()

        // Optional. Switch the session to a monotonic behavior.
        session.SetMode(mgo.Monotonic, true)

        c := session.DB("blog").C("post")

        result := []Post{}
        err = c.Find(bson.M{}).All(&result)
        if err != nil {
                log.Fatal(err)
        }

        return result
}

type commentStruct struct {
        PostId string `json:"postId"`
        Text string `json:"text"`
}

func SelectComment(comments []bson.ObjectId) []commentStruct {
        session, err := mgo.Dial("localhost")
        if err != nil {
                panic(err)
        }
        defer session.Close()

        // Optional. Switch the session to a monotonic behavior.
        session.SetMode(mgo.Monotonic, true)

        c := session.DB("blog").C("comment")

        
        results := []commentStruct{}

        for id := range comments {
                result := commentStruct{}
                err = c.Find(bson.M{"_id": comments[id]}).One(&result)
                if err != nil {
                        // log.Fatal(err)
                }
                results = append(results,result)
        }
        

        return results
}

func InsertComment(postId string, text string) Post {
        session, err := mgo.Dial("localhost")
        if err != nil {
                panic(err)
        }
        defer session.Close()

        // Optional. Switch the session to a monotonic behavior.
        session.SetMode(mgo.Monotonic, true)

        c := session.DB("blog").C("comment")

        id := bson.NewObjectId()

        err = c.Insert(bson.M{
                "_id": id,
                "text": text,
        })

        
        if err != nil {
                log.Fatal(err)
        }

        result := Post{}
        err = c.Find(bson.M{"_id": id}).One(&result)
        if err != nil {
                log.Fatal(err)
        }
        
        insertCommentIdToPost(postId, id)

        return result
}

func insertCommentIdToPost(postId string, commentId bson.ObjectId) (info *mgo.ChangeInfo, err error){
        session, err := mgo.Dial("localhost")
        if err != nil {
                panic(err)
        }
        defer session.Close()

        // Optional. Switch the session to a monotonic behavior.
        session.SetMode(mgo.Monotonic, true)

        c := session.DB("blog").C("post")


        result := Post{}
        err = c.Find(bson.M{"_id": bson.ObjectIdHex(postId)}).One(&result)
        result.CommentsId = append(result.CommentsId,commentId)
        if err != nil {
                fmt.Println(">>>>>",postId,"<<<",commentId,"++++",err)
                log.Fatal(err)
        }

        return c.Upsert(bson.M{"_id": bson.ObjectIdHex(postId)}, result)

}
