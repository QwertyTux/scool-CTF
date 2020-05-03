package main

import(
  "net/http"
  "log"
  "text/template"
  "fmt"
  "os"
  )

//Главная функция
func main()  {

  var p string
  fmt.Print("Введите порт, который будет исользоваться: ")
  fmt.Fscan(os.Stdin, &p)

  http.HandleFunc("/", mainPage)
  http.HandleFunc("/page_1", page_1)
  http.HandleFunc("/page_2", page_2)
  http.HandleFunc("/page_3", page_3)

  port := ":" + p
  println("Server listen on port:", port)
  err := http.ListenAndServe(port, nil)

  if err != nil {
    log.Fatal("ListenAndServe", err)
  }
}

//Функция главной страницы
func mainPage(w http.ResponseWriter, r *http.Request )  {
  tmpl, err := template.ParseFiles("templates/mainPage.html")

  if err != nil{
    http.Error(w, err.Error(), 400)
    return
  }

  if err := tmpl.Execute(w, nil); err != nil{
    http.Error(w, err.Error(), 400)
    return
  }
}

//Функция страницы 1
func  page_1(w http.ResponseWriter, r *http.Request )  {
  tmpl, err := template.ParseFiles("templates/page_1.html")

  if err != nil{
    http.Error(w, err.Error(), 400)
    return
  }

  if err := tmpl.Execute(w, nil); err != nil{
    http.Error(w, err.Error(), 400)
    return
  }
}

//Функция страницы 2
func  page_2(w http.ResponseWriter, r *http.Request )  {
  tmpl, err := template.ParseFiles("templates/page_2.html")

  if err != nil{
    http.Error(w, err.Error(), 400)
    return
  }

  if err := tmpl.Execute(w, nil); err != nil{
    http.Error(w, err.Error(), 400)
    return
  }
}

//Функция станицы 3
func page_3(w http.ResponseWriter, r *http.Request ){

  tmpl, err := template.ParseFiles("templates/page_3.html")
  if err != nil{
    http.Error(w, err.Error(), 400)
    return
  }

  if err := tmpl.Execute(w, nil); err != nil{
    http.Error(w, err.Error(), 400)
    return
  }

  if r.Method == "POST" {

    err := r.ParseForm()
  	if err != nil {
  		log.Println(err)
  	}

    vvod := r.PostFormValue("vvod")

    if vvod == "\"><script>alert(document.cookie)</script>" || vvod == "<script>alert(document.cookie)</script>" || vvod == "'><script>alert(document.cookie)</script>" {
      tmpl, err := template.ParseFiles("templates/page_3_flug_alert.html")
      if err != nil{
        http.Error(w, err.Error(), 400)
        return
      }

      if err := tmpl.Execute(w, nil); err != nil{
        http.Error(w, err.Error(), 400)
        return
      }
    }

    if vvod == "\"><script>document.write(document.cookie)</script>" || vvod == "<script>document.write(document.cookie)</script>" || vvod == "'><script>document.write(document.cookie)</script>" {
      tmpl, err := template.ParseFiles("templates/page_3_flug_document.html")
      if err != nil{
        http.Error(w, err.Error(), 400)
        return
      }

      if err := tmpl.Execute(w, nil); err != nil{
        http.Error(w, err.Error(), 400)
        return
      }
    }

    if vvod == "\"><script>console.log(document.cookie)</script>" || vvod == "<script>console.log(document.cookie)</script>" || vvod == "'><script>console.log(document.cookie)</script>" {
      tmpl, err := template.ParseFiles("templates/page_3_flug_console.html")
      if err != nil{
        http.Error(w, err.Error(), 400)
        return
      }

      if err := tmpl.Execute(w, nil); err != nil{
        http.Error(w, err.Error(), 400)
        return
      }
    }
  }
}
