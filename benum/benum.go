package benum

import (
	"text/template"

	pgs "github.com/lyft/protoc-gen-star"
	pgsgo "github.com/lyft/protoc-gen-star/lang/go"
)

type module struct {
	*pgs.ModuleBase
	ctx pgsgo.Context
	tpl *template.Template
}

func Benum() *module { return &module{ModuleBase: &pgs.ModuleBase{}} }

func (p *module) Name() string {
	return "benum"
}

func (p *module) InitContext(c pgs.BuildContext) {
	p.ModuleBase.InitContext(c)
	p.ctx = pgsgo.InitContext(c.Parameters())

	tpl := template.New("benum").Funcs(map[string]interface{}{
		"package": p.ctx.PackageName,
		"name":    p.ctx.Name,
		"dbVal":   p.dbVal,
		"gqlVal":  p.gqlVal,
	})

	p.tpl = template.Must(tpl.Parse(benumTpl))
}

func (p *module) Execute(targets map[string]pgs.File, pkgs map[string]pgs.Package) []pgs.Artifact {
	for _, t := range targets {
		p.generate(t)
	}
	return p.Artifacts()
}

func (p *module) generate(f pgs.File) {
	if len(f.Enums()) == 0 {
		return
	}

	name := p.ctx.OutputPath(f).SetExt(".benum.go")
	p.AddGeneratorTemplateFile(name.String(), p.tpl, f)
}

func (p *module) dbVal(ev pgs.EnumValue) string {
	var val string
	ok, err := ev.Extension(E_Db, &val)
	if !ok || err != nil {
		return ev.Name().String()
	}
	return val
}

func (p *module) gqlVal(ev pgs.EnumValue) string {
	var val string
	ok, err := ev.Extension(E_Gql, &val)
	if !ok || err != nil {
		return ev.Name().String()
	}
	return val
}

const benumTpl = `package {{ package . }}

import (
	"fmt"
	"database/sql/driver"
	"io"
	"strconv"
)

{{ range .AllEnums }}

// -------------------------------------------------------------------
// {{ name . }}
// -------------------------------------------------------------------

func (e {{ name . }}) IsValid() bool {
	_, ok := {{ name . }}_name[int32(e)]
	return ok
}

// ------------------------- gqlgen ----------------------------------

var {{ name . }}_gql_name = map[int32]string{
{{ range .Values -}}
	{{ .Value }}: "{{ gqlVal . }}",
{{ end }}
}

var {{ name . }}_gql_value = map[string]int32{
{{ range .Values -}}
	"{{ gqlVal . }}": {{ .Value }},
{{ end }}
}

func (e {{ name . }}) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote({{ name . }}_gql_name[int32(e)]))
}

func (e *{{ name . }}) UnmarshalGQL(v interface{}) error {
	value, ok := v.(string)
	if !ok {
		return fmt.Errorf("%T is not a valid {{ name . }}", v)
	}
	res, ok := {{ name . }}_gql_value[value]
	if !ok {
		return fmt.Errorf("%T is not a valid {{ name . }}", v)
	}
	*e = {{ name . }}(res)
	return nil
}

// --------------------------- db ------------------------------------

var {{ name . }}_db_name = map[int32]string{
{{ range .Values -}}
	{{ .Value }}: "{{ dbVal . }}",
{{ end }}
}

var {{ name . }}_db_value = map[string]int32{
{{ range .Values -}}
	"{{ dbVal . }}": {{ .Value }},
{{ end }}
}

func (e {{ name . }}) Value() (driver.Value, error) {
	if !e.IsValid() {
		return nil, fmt.Errorf("invalid {{ name . }} '%s'", e)
	}
	return {{ name . }}_db_name[int32(e)], nil
}

func (e *{{ name . }}) Scan(value interface{}) error {
	sv, err := driver.String.ConvertValue(value)
	if err != nil {
		return fmt.Errorf("failed to scan %#v into {{ name . }}", value)
	}
	res, ok := int32(0), false
	switch v := sv.(type) {
	case string:
		res, ok = {{ name . }}_db_value[v]
	case []byte:
		res, ok = {{ name . }}_db_value[string(v)]
	default:
		panic("unexpected type from ConvertValue")
	}
	if !ok {
		panic(fmt.Errorf("invalid Enum1 '%s'", e))
	}
	*e = {{ name . }}(res)
	return nil
}

{{ end }}
`
