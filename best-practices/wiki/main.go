package main

import (
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
)

const rootPath = "/Users/wenzg/workspace/go/src/go-study/best-practices/wiki"

type Page struct {
	Title string // 首字母要大写，不然tpl读不出来
	Body  []byte
}

func savePage(page *Page) error {
	fileName := rootPath + "/tmp/" + page.Title + ".txt"
	return ioutil.WriteFile(fileName, page.Body, 0600)
}

func loadPage(title string) (*Page, error) {
	fileName := rootPath + "/tmp/" + title + ".txt"

	b, err := ioutil.ReadFile(fileName)
	if err != nil {
		return nil, err
	}

	return &Page{title, b}, nil
}

func main() {

	wikiHandler := func(fn func(http.ResponseWriter, *http.Request, string)) http.HandlerFunc {
		var validPath = regexp.MustCompile("^/(edit|save|view)/([\\w\\-]+)$")
		return func(w http.ResponseWriter, r *http.Request) {
			m := validPath.FindStringSubmatch(r.URL.Path)
			if m == nil {
				http.Error(w, "非法参数", http.StatusNotFound)
				return
			}
			fn(w, r, m[2])
		}
	}

	http.HandleFunc("/view/", wikiHandler(viewHandler))
	http.HandleFunc("/save/", wikiHandler(saveHandler))
	http.HandleFunc("/edit/", wikiHandler(editHandler))

	log.Fatal(http.ListenAndServe(":8888", nil))
}

func editHandler(w http.ResponseWriter, r *http.Request, wikiTitle string) {
	// load
	page, err := loadPage(wikiTitle)

	if err != nil {
		page = &Page{wikiTitle, []byte("")}
	}

	renderTemplate(w, r, "edit", page)
}

func saveHandler(w http.ResponseWriter, r *http.Request, wikiTitle string) {
	body := r.FormValue("body")

	err := savePage(&Page{wikiTitle, []byte(body)})

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/view/"+wikiTitle, http.StatusFound)
}

// 缓存
var templates = template.Must(template.ParseFiles(rootPath+"/tpl/view.html", rootPath+"/tpl/edit.html", ))

// wiki内容显示
func viewHandler(w http.ResponseWriter, r *http.Request, wikiTitle string) {

	// load
	page, err := loadPage(wikiTitle)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	renderTemplate(w, r, "view", page)
}

// 模版渲染
func renderTemplate(w http.ResponseWriter, r *http.Request, tpl string, page *Page) {
	err := templates.ExecuteTemplate(w, tpl+".html", page)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
