package executor

import (
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

	t.tmpl = template.New("").Funcs(*t.functions)
	_, err := t.tmpl.ParseFiles(tmpls...)
	if err != nil {
		err = errors.Wrap(err, "parse template failed")
	}
	return err
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
	return t.names
}
