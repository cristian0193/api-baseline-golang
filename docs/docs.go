package docs

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"strings"
	"ws-baseline-golang-v2/utils"

	"github.com/alecthomas/template"
	"github.com/swaggo/swag"
)

type swaggerInfo struct {
	Version     string
	Host        string
	BasePath    string
	Schemes     []string
	Title       string
	Description string
}

var SwaggerInfo = swaggerInfo{
	Version:     "",
	Host:        "",
	BasePath:    "",
	Schemes:     []string{},
	Title:       "",
	Description: "",
}

type s struct{}

func (s *s) ReadDoc() string {
	sInfo := SwaggerInfo
	sInfo.Description = strings.Replace(sInfo.Description, "\n", "\\n", -1)

	doc, err := ioutil.ReadFile(utils.GetStrEnv("SWAGGER_JSON"))
	if err != nil {
		return err.Error()
	}

	t, err := template.New("swagger_info").Funcs(template.FuncMap{
		"marshal": func(v interface{}) string {
			a, _ := json.Marshal(v)
			return string(a)
		},
	}).Parse(string(doc))
	if err != nil {
		return string(doc)
	}

	var tpl bytes.Buffer
	if err := t.Execute(&tpl, sInfo); err != nil {
		return string(doc)
	}

	return tpl.String()
}

func init() {
	swag.Register(swag.Name, &s{})
}
