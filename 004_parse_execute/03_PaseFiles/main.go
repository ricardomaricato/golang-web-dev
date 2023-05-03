package main

import (
	"log"
	"os"
	"text/template"
)

func main() {
	tpl, err := template.ParseFiles("one.gmao")
	if err != nil {
		log.Fatalln(err)
	}

	if err = tpl.Execute(os.Stdout, nil); err != nil {
		log.Fatalln(err)
	}

	tpl, err = tpl.ParseFiles("two.gmao", "vespa.gmao")
	if err != nil {
		log.Fatalln(err)
	}

	if err = tpl.ExecuteTemplate(os.Stdout, "two.gmao", nil); err != nil {
		log.Fatalln(err)
	}

	if err = tpl.ExecuteTemplate(os.Stdout, "vespa.gmao", nil); err != nil {
		log.Fatalln(err)
	}

	if err = tpl.ExecuteTemplate(os.Stdout, "one.gmao", nil); err != nil {
		log.Fatalln(err)
	}

	if err = tpl.Execute(os.Stdout, nil); err != nil {
		log.Fatalln(err)
	}
}
