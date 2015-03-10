package main

import (
    "github.com/ant0ine/go-json-rest/rest"
    "log"
    "net/http"
    "blog/db"
    "fmt"
)

type Message struct {
    Body string
}

func main() {
    api := rest.NewApi()
    api.Use(rest.DefaultDevStack...)
    router, err := rest.MakeRouter(
        &rest.Route{"GET", "/simpleInsert", simpleInsert},
        &rest.Route{"POST", "/postBlog", post},
        &rest.Route{"POST", "/comment", comment},
        &rest.Route{"GET", "/all", all},
        // &rest.Route{"GET", "/qcomment", qcomment},
    )
    if err != nil {
        log.Fatal(err)
    }
    api.SetApp(router)
    log.Fatal(http.ListenAndServe(":8080", api.MakeHandler()))
}

func simpleInsert(w rest.ResponseWriter, r *rest.Request) {

	w.WriteJson(db.Insertdb())
}

func post(w rest.ResponseWriter, r *rest.Request) {
	var input db.Post
	r.DecodeJsonPayload(&input)
	// fmt.Printf("%#v",input)
	// w.WriteJson(input)
	w.WriteJson(db.InsertPost(input.Title,input.Author,input.Category,input.Content))
	// w.WriteJson(db.InsertPost("Test","Pallat","Show","Test1"))
}

type commentStruct struct {
	PostId string `json:"postId"`
	Text string `json:"text"`
}

func comment(w rest.ResponseWriter, r *rest.Request) {
	var input commentStruct
	r.DecodeJsonPayload(&input)
	fmt.Printf("%#v",input)
	// w.WriteJson(input)
	w.WriteJson(db.InsertComment(input.PostId,input.Text))
}


func all(w rest.ResponseWriter, r *rest.Request) {
	responses := db.SelectAll()
	var comments []string
	for i:= range responses {
		comments = []string{}
		x := db.SelectComment(responses[i].CommentsId)
		for j:= range x {
			comments = append(comments, x[j].Text)
		}
		responses[i].Comments = comments
	}
	fmt.Println(responses)
	w.WriteJson(responses)
}