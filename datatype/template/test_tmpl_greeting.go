package main

import (
	"html/template"
	"os"
)

type Friend struct {
	Name string
}

type Person struct {
	Username string
	Emails   []string
	Friends  []*Friend
}

func main() {

	f1 := Friend{Name: "lance.qi"}
	f2 := Friend{Name: "Wangwei"}

	p := Person{
		Username: "Donald",
		Emails:   []string{"xyz@gmail.com", "kkk@outlook.com"},
		Friends:  []*Friend{&f1, &f2}}

	t := template.New("test_template_example")
	t, _ = t.Parse(`hello {{.Username}}!
            {{range .Emails}}
                an email {{.}}
            {{end}}
            {{with .Friends}}
            {{range .}}
                friend name is {{.Name}}
            {{end}}
            {{end}}
            `)

	// email filter function
	// t = t.Funcs(template.FuncMap{"email_func": emailFilter})

	t.Execute(os.Stdout, p)
}
