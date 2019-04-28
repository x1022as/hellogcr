package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/exec"
	"strings"
	"time"
)

var logFile *os.File

func handler(w http.ResponseWriter, r *http.Request) {
	log.Print("Hello world received a request.")
	n := time.Now()
	fmt.Fprintf(logFile, "new request at: %s\n", n.Format(time.RFC3339Nano))
	target := os.Getenv("TARGET")
	if target == "" {
		target = "World"
	}

	cLen := r.ContentLength
	body := make([]byte, cLen)
	r.Body.Read(body)

	ss := strings.Split(string(body), " ")
	if len(ss) == 0 {
		fmt.Fprintf(w, "no command specified\n")
		return
	}
	nss := []string{"-c"}
	jss := strings.Join(ss, " ")
	nss = append(nss, jss)

	cmd := exec.Command("bash", nss...)
	out, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Fprintf(w, "%s failed with %s\n%s\n", string(body), string(out), err)
		return
	}

	fmt.Fprintf(w, "%s output is :\n%s\n", string(body), string(out))

	/*
		cmds := strings.Split(string(body), "&&")
		for _, cmdstr := range cmds {
			cmdstr = strings.TrimSuffix(cmdstr, " ")
			cmdstr = strings.TrimPrefix(cmdstr, " ")
			ss := strings.Split(string(cmdstr), " ")
			var cmd *exec.Cmd
			if len(ss) == 0 {
				fmt.Fprintf(w, "no command specified\n")
				return
			} else if len(ss) == 1 {
				cmd = exec.Command(ss[0])
			} else {
				cmd = exec.Command(ss[0], ss[1:]...)
			}
			out, err := cmd.CombinedOutput()
			if err != nil {
				fmt.Fprintf(w, "%s failed with %s\n", string(cmdstr), err)
				return
			}
			fmt.Fprintf(w, "%s output is :\n%s\n", string(cmdstr), string(out))
		}
	*/
}

func main() {
	log.Print("Hello world sample started.")

	logFile, _ = os.OpenFile("/logfile", os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0666)
	defer logFile.Close()
	n := time.Now()
	fmt.Fprintf(logFile, "service started at: %s\n", n.Format(time.RFC3339Nano))

	http.HandleFunc("/", handler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	ds := os.Getenv("DEBUG_SLEEP")
	if ds == "1" {
		time.Sleep(10 * time.Second)
	}

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}
