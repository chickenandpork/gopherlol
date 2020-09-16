package commands

import (
	"fmt"
	"net/http"
	"net/url"
	"reflect"
	"strings"
)

var commandsObject = new(Commands)

// Handler is the http dispatcher and responder: it either derives a list of available commands
// (sending back directly), or dispatches the search query to that set of commands (wrapping the
// response as a HTTP Redirect which should result in a HTTP/1 302)
func Handler(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query().Get("q")

	arr := strings.SplitN(q, " ", 2)
	cmdName := strings.Title(arr[0])
	cmdArg := ""
	if len(arr) > 1 {
		cmdArg = arr[1]
	}

	if cmdName == "List" || cmdName == "Help" {
		var html strings.Builder
		html.WriteString("<h1>gopherlol command list</h1>")
		html.WriteString("<ul>")
		for _, source := range GetCommands() {
			commandsType := reflect.TypeOf(source)

			for i := 0; i < commandsType.NumMethod(); i++ {
				method := commandsType.Method(i)

				takesArgs := ""
				if method.Type.NumIn() == 2 {
					takesArgs = ", takes args"
				}
				html.WriteString(fmt.Sprintf(
					"<li><strong>%s</strong>%s</li>",
					strings.ToLower(method.Name),
					takesArgs,
				))
			}
		}
		html.WriteString("</ul>")

		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		w.WriteHeader(http.StatusOK)
		_, _ = fmt.Fprint(w, html.String())
		return
	}

	/* else */
	// TODO: Still need to special-case the :Author() function so that a list is presented
	// similar to the "list" or "help" options.  THis allows us to leverage the signature
	// method of the CommandSource to provide a list of contributing sources.
	for _, source := range GetCommands() {
		cmdMethod := reflect.ValueOf(source).MethodByName(cmdName)

		if cmdMethod != reflect.ValueOf(nil) {
			url := ""
			cmdMethodNumIn := cmdMethod.Type().NumIn()
			if cmdMethodNumIn == 0 {
				res := cmdMethod.Call([]reflect.Value{})
				url = res[0].String()
			} else if cmdMethodNumIn == 1 {
				in := []reflect.Value{reflect.ValueOf(cmdArg)}
				res := cmdMethod.Call(in)
				url = res[0].String()
			} else {
				// cmdMethod was wrongly defined.
				// We currently only support cmdMethods with 0 or 1 parameters
				http.Error(
					w,
					http.StatusText(http.StatusInternalServerError),
					http.StatusInternalServerError,
				)
				return
			}

			http.Redirect(w, r, url, http.StatusSeeOther)
			logsink.Printf("found command: %+v -> redirecting to: %s\n", cmdMethod, url)
			return
		}
	}
	/* else */
	logsink.Printf("not found (forwarding to google): cmd: %s arg: %s args: %+v\n", cmdName, cmdArg, arr)

	// cmdMethod not found => fall back to google
	url := fmt.Sprintf("https://www.google.com/#q=%s", url.QueryEscape(q))
	http.Redirect(w, r, url, http.StatusSeeOther)
	return
}
