package view

import (
	"os"
)

func GetHomePage() string {
	ret := `<!DOCTYPE html>
<html lang="ru">
<head>
<title>Chat Example</title>`

	if s, err := os.ReadFile("view/form.js"); err == nil {
		ret = ret + `<script type="text/javascript">` + string(s) + `</script>`
	}

	if s, err := os.ReadFile("view/init.js"); err == nil {
		ret = ret + `<script type="text/javascript">` + string(s) + `</script>`
	}

	if s, err := os.ReadFile("view/form.css"); err == nil {
		ret = ret + `<style type="text/css">` + string(s) + `</style>`
	}

	ret = ret + `</head><body>`

	if s, err := os.ReadFile("view/form.html"); err == nil {
		ret = ret + string(s)
	}

	ret = ret + `<div id="connState" ></div></body></html>`
	return ret

}
