package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"text/template"
)

func main() {
	tmp1 := template.New("main")
	tmp1, _ = tmp1.Parse(
		`<div style="display; inline-block; border: 1px solid #aaa;
		border-radius: 3px; padding: 30px; margin: 20px;">
			{{if ne . "str"}}
				not str
			{{end}}
			<pre>{{.}}</pre>
		</div>
		<b>asda</b>
		<p id="demo"></p>
		<script>
		document.getElementById("demo").innerHTML = "My First JS";
		</script>
		`)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path

		c := http.Client{}
		resp, err := c.Get("http://artii.herokuapp.com/make?text=" + path)
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("error"))
			return
		}

		defer resp.Body.Close()

		body, _ := ioutil.ReadAll(resp.Body)

		tmp1.Execute(w, string(body))
	})

	http.ListenAndServe(":8001", nil)
}
