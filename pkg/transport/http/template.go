package http

import (
	"html/template"
	"net/http"

	"github.com/sirupsen/logrus"
)

type Templates struct {
	Main *template.Template
	log  *logrus.Logger
}

func NewTemplates(log *logrus.Logger) (t Templates) {

	tpl, err := template.ParseFiles("web/html/header.html", "web/html/index.html", "web/html/footer.html")
	t.Main = tpl
	if err != nil {
		t.log.Error(err, "не удалось открыть главную страничку")
	}
	return t
}

func (t *Templates) indexPage(w http.ResponseWriter, r *http.Request) {
	err := t.Main.ExecuteTemplate(w, "index", nil)
	if err != nil {
		t.log.Error(err)
	}
}
