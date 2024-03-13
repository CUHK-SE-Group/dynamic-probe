package executor

import (
	"fmt"
	"io"
	"text/template"

	"github.com/pkg/errors"
)

type TemplateMgr struct {
	functions *template.FuncMap
	tmpl      *template.Template
	names     []string
}

func (t *TemplateMgr) LoadTemplates(tmpls ...string) error {
	t.names = append(t.names, tmpls...)

	var err error
	t.tmpl = template.New("").Funcs(*t.functions)
	for _, v := range tmpls {
		fmt.Println(v)
		t.tmpl, err = t.tmpl.ParseGlob(v)
		if err != nil {
			return errors.Wrap(err, "parse template dir failed")
		}
	}
	return nil
}

func (t *TemplateMgr) LoadFunc(name string, f any) {
	(*t.functions)[name] = f
}

func (t *TemplateMgr) LoadFuncMap(m *template.FuncMap) {
	t.functions = m
}

func (t *TemplateMgr) Generate(wr io.Writer, name string, data any) error {
	return t.tmpl.ExecuteTemplate(wr, name, data)
}

func (t *TemplateMgr) GetNames() []string {
	names := []string{}
	for _, v := range t.tmpl.Templates() {
		names = append(names, v.Name())
	}
	return names
}
