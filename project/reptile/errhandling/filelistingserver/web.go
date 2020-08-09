package main

import (
	"log"
	"net/http"
	_ "net/http/pprof"
	"os"
	"reptile/learning.fufeng.com/project/reptile/errhandling/filelistingserver/filelisting"
)

type appHandler func(writer http.ResponseWriter,
	request *http.Request) error

func errWrapper(
	handler appHandler) func(
	http.ResponseWriter, *http.Request) {
	return func(writer http.ResponseWriter,
		request *http.Request) {
		// panic
		defer func() {
			if r := recover(); r != nil {
				log.Printf("Panic: %v", r)
				http.Error(writer,
					http.StatusText(http.StatusInternalServerError),
					http.StatusInternalServerError)
			}
		}()

		err := handler(writer, request)

		if err != nil {
			log.Printf("Error occurred "+
				"handling request: %s",
				err.Error())

			// user error
			if userErr, ok := err.(userError); ok {
				http.Error(writer,
					userErr.Message(),
					http.StatusBadRequest)
				return
			}

			// system error
			code := http.StatusOK
			switch {
			case os.IsNotExist(err):
				code = http.StatusNotFound
			case os.IsPermission(err):
				code = http.StatusForbidden
			default:
				code = http.StatusInternalServerError
			}
			http.Error(writer,
				http.StatusText(code), code)
		}
	}
}

type userError interface {
	error
	Message() string
}

/**
  开启_ "net/http/pprof" 可以使用下面的链接访问
  http://localhost:8888/debug/pprof/

  go tool pprof http://localhost:8888/debug/pprof/profile?seconds=30
	Type: cpu
	Time: Feb 15, 2020 at 12:59pm (CST)
	Duration: 30s, Total samples = 0
	No samples were found with the default sample value type.
	Try "sample_index" command to analyze different sample values.
	Entering interactive mode (type "help" for commands, "o" for options)
	(pprof) web
	(pprof) exit

    godoc -http :8888
*/
func main() {
	http.HandleFunc("/",
		errWrapper(filelisting.HandleFileList))

	err := http.ListenAndServe(":8888", nil)
	if err != nil {
		panic(err)
	}
}
