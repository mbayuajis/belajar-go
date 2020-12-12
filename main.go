package main

import "fmt"
import "net/http"
import "html/template"
// import "path"

// func handlerIndex(w http.ResponseWriter, r *http.Request) {
//     var filepath = path.Join("views", "index.html")
// 	var tmpl, err = template.ParseFiles(filepath)
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		return
// 	}

// 	var data = map[string]interface{}{
// 		"title": "Learning Golang Web",
// 		"name":  "Batman",
// 	}

// 	err = tmpl.Execute(w, data)
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 	}
// }

// func handlerHello(w http.ResponseWriter, r *http.Request) {
//     var message = "Hello world!"
//     w.Write([]byte(message))
// }

// func main() {
// 	http.Handle("/static/",
// 		http.StripPrefix("/static/",
// 			http.FileServer(http.Dir("assets"))))

//     http.HandleFunc("/", handlerIndex)
//     http.HandleFunc("/index", handlerIndex)
//     http.HandleFunc("/hello", handlerHello)

//     var address = "localhost:9000"
//     fmt.Printf("server started at %s\n", address)
//     err := http.ListenAndServe(address, nil)
//     if err != nil {
//         fmt.Println(err.Error())
//     }
// }

type M  map[string]interface{}

func main() {
    // var tmpl, err = template.ParseGlob("views/*")
    // if err != nil {
    //     panic(err.Error())
    //     return
	// }

	// http.HandleFunc("/index", func(w http.ResponseWriter, r *http.Request) {
	// 	var data = M{"name": "Batman"}
	// 	err = tmpl.ExecuteTemplate(w, "index", data)
	// 	if err != nil {
	// 		http.Error(w, err.Error(), http.StatusInternalServerError)
	// 	}
	// })

	// http.HandleFunc("/about", func(w http.ResponseWriter, r *http.Request) {
	// 	var data = M{"name": "Batman"}
	// 	err = tmpl.ExecuteTemplate(w, "about", data)
	// 	if err != nil {
	// 		http.Error(w, err.Error(), http.StatusInternalServerError)
	// 	}
	// })

	http.HandleFunc("/index", func(w http.ResponseWriter, r *http.Request) {
		var data = M{"name": "Batman"}
		var tmpl = template.Must(template.ParseFiles(
			"views/index.html",
			"views/_header.html",
		))

		var err = tmpl.ExecuteTemplate(w, "index", data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	fmt.Println("server started at localhost:9000")
	http.ListenAndServe(":9000", nil)
}
