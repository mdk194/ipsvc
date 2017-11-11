package main

import (
	"html/template"
	"log"
	"net"
	"net/http"
	"strings"
)

// Page ...
type Page struct {
	IP      string
	Color   string
	Version string
}

var (
	doc = `
		<html>
			<body bgcolor="#{{.Color}}">
				<h1>IP: {{.IP}}</h1>
				<h1>Version: {{.Version}}</h1>
			</body>
		</html>
	`

	ip      string
	t       *template.Template
	color   string
	version string
)

func init() {
	localAddr := getLocalAddr()
	ip = strings.Split(localAddr, ":")[0]

	version = "1.0"
	color = "FFFFFF"

	t = template.New("site")
	t.Parse(doc)
}

func getLocalAddr() string {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	localAddr := conn.LocalAddr().(*net.UDPAddr).String()
	return localAddr
}

func renderMain(w http.ResponseWriter, r *http.Request) {
	page := &Page{
		IP:      ip,
		Color:   color,
		Version: version,
	}
	t.Execute(w, page)
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", renderMain)
	http.ListenAndServe(":8000", mux)
}
