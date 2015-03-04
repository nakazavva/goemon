package main

import (
	"fmt"
	"github.com/mattn/goemon"
	"os"
	"sort"
)

var defaultConf = map[string]string{
	"web": `# Generated by goemon -g
livereload: :35730
tasks:
- match: './assets/*.js'
  commands:
  - minifyjs -m -i ${GOEMON_TARGET_FILE} > ${GOEMON_TARGET_DIR}/${GOEMON_TARGET_NAME}.min.js
  - :livereload /
- match: './assets/*.css'
  commands:
  - :livereload /
- match: './assets/*.html'
  commands:
  - :livereload /
- match: '*.go'
  commands:
  - go build
  - :restart
  - :livereload /
`,
	"c": `# Generated by goemon -g
tasks:
- match: '*.(c|cpp|cxx)'
  commands:
  - make
`,
	"md": `# Generated by goemon -g
tasks:
- match: '*.md'
  commands:
  - pandoc -f markdown -t html -o ${GOEMON_TARGET_DIR}/${GOEMON_TARGET_NAME}.html ${GOEMON_TARGET_FILE}
`,
}

func usage() {
	fmt.Printf("Usage of %s [options] [command] [args...]\n", os.Args[0])
	fmt.Println(" goemon -g [NAME]     : generate default configuration")
	fmt.Println(" goemon -c [FILE] ... : set configuration file")
	fmt.Println("")
	fmt.Println("* Examples:")
	fmt.Println("  Generate default configuration:")
	fmt.Println("    goemon -g > goemon.yml")
	fmt.Println("")
	fmt.Println("  Generate C configuration:")
	fmt.Println("    goemon -g c > goemon.yml")
	fmt.Println("")
	fmt.Println("  List default configurations:")
	fmt.Println("    goemon -g ?")
	fmt.Println("")
	fmt.Println("  Start standalone server:")
	fmt.Println("    goemon --")
	os.Exit(1)
}

func main() {
	file := ""
	args := []string{}

	switch len(os.Args) {
	case 1:
		usage()
	default:
		switch os.Args[1] {
		case "-h":
			usage()
		case "-g":
			if len(os.Args) == 2 {
				fmt.Print(defaultConf["web"])
			} else if os.Args[2] == "?" {
				var keys []string
				for k := range defaultConf {
					keys = append(keys, k)
				}
				sort.Strings(keys)
				for _, k := range keys {
					fmt.Println(k)
				}
			} else if t, ok := defaultConf[os.Args[2]]; ok {
				fmt.Print(t)
			} else {
				usage()
			}
			return
		case "-c":
			if len(os.Args) == 2 {
				usage()
				return
			}
			file = os.Args[2]
			args = os.Args[3:]
		case "--":
			args = os.Args[2:]
		default:
			args = os.Args[1:]
		}
	}

	g := goemon.NewWithArgs(args)
	if file != "" {
		g.File = file
	}
	g.Run()
	if len(args) == 0 {
		select {}
	}
}
