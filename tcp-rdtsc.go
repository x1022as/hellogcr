package main

import (
	"fmt"
	"net"
	//"time"
	"os"

	"github.com/knative/docs/helloworld/rdtsc"
)

//const addr = "10.113.190.31:9999"
var addr string

const TTimes = 3

func main() {
	if len(os.Args) != 2 {
		fmt.Println("usage: %s IP:PORT", os.Args[0])
		return
	}
	/*
		if err := dd(); err != nil {
			fmt.Printf("dd failed: %s", err)
			return
		}*/

	addr = os.Args[1]

	done := make(chan bool)
	go server(done)

	<-done

	// client here
	rdBuf := make([]byte, 1024)
	wrBuf := []byte("ksdflsldfjdknkvckljodsifu0909809sdfsodfjsldjflksjdfdflkjsdlkfjsdfoiu09823042joidjfosdufsdfjsdfoiuwerw0e98sdfjasodifuosdfuoasudfosdufosufsjdiuwoer98sdfsid")
	c, err := net.Dial("tcp", addr)
	if err != nil {
		fmt.Println("client dial failed: ", err)
		return
	}

	// begin := time.Now()
	for i := 0; i < TTimes; i++ {
		fmt.Println("=====begin=====")
		start := rdtsc.Rdtsc()
		if _, err := c.Write(wrBuf); err != nil {
			fmt.Println("client write failed: ", err)
			return
		}
		end := rdtsc.Rdtsc()
		fmt.Printf("client write takes %d\n", end-start)
		start = rdtsc.Rdtsc()
		_, err := c.Read(rdBuf)
		end = rdtsc.Rdtsc()
		if err != nil {
			fmt.Println("client read failed: ", err)
			return
		}
		fmt.Printf("client read takes %d\n", end-start)
		fmt.Println("=====end=====\n")
	}
	// sub := time.Now().Sub(begin)
	// fmt.Printf("time cost is %f s\n", sub.Seconds())

	return
	// wait for server done

}

func server(begin chan bool) {
	l, err := net.Listen("tcp", addr)
	if err != nil {
		fmt.Printf("listen failed: %s\n", err)
		return
	}
	defer l.Close()

	close(begin)

	rdBuf := make([]byte, 1024)
	wrBuf := []byte("sdfsdflsdkweorikjldfjlasjdf90923498279347230429340283471923ijjdjkjdfljasldfjsdiuoiuer09w83479287340203490830928492347123103810238102odsfjdljldjsdfjsdfosudfsdiufosifuosdfu")
	conn, err := l.Accept()
	if err != nil {
		fmt.Println("accept failed: ", err)
		return
	}

	for i := 0; i < TTimes; i++ {
		start := rdtsc.Rdtsc()
		_, err := conn.Read(rdBuf)
		end := rdtsc.Rdtsc()
		if err != nil {
			fmt.Println("server read failed: ", err)
			return
		}
		fmt.Printf("server read takes %d\n", end-start)
		start = rdtsc.Rdtsc()
		if _, err := conn.Write(wrBuf); err != nil {
			fmt.Println("server write failed: ", err)
			return
		}
		end = rdtsc.Rdtsc()
		fmt.Printf("server write takes %d\n", end-start)
	}

	return
}
