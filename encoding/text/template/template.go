package templatei

import (
	"io"
	"strings"
	"text/template"

	"github.com/hopeio/utils/log"
)

var CommonTemp = template.New("all")

func init() {
	CommonTemp.Funcs(template.FuncMap{"join": strings.Join})
}
func Parse(tpl string) *template.Template {
	t, err := CommonTemp.Parse(tpl)
	if err != nil {
		log.Fatal(err)
	}
	return t
}

func Execute(wr io.Writer, name string, data interface{}) error {
	return CommonTemp.ExecuteTemplate(wr, name, data)
}
