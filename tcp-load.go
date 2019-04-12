package main

import (
	"fmt"
	"net"
	"os"
	"strconv"
	"time"
)

var addr string
var TTimes int

func main() {
	if len(os.Args) != 3 {
		fmt.Println("usage: %s TIMES IP:PORT", os.Args[0])
		return
	}

	var err error
	TTimes, err = strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Printf("invalid Times %s - %s\n", os.Args[1], err)
		return
	}
	addr = os.Args[2]

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

	fmt.Println("=====begin=====")
	begin := time.Now()
	for i := 0; i < TTimes; i++ {
		if _, err := c.Write(wrBuf); err != nil {
			fmt.Println("client write failed: ", err)
			return
		}
		_, err := c.Read(rdBuf)
		if err != nil {
			fmt.Println("client read failed: ", err)
			return
		}
	}
	sub := time.Now().Sub(begin)
	fmt.Printf("time cost is %f s\n", sub.Seconds())
	fmt.Println("=====end=====")

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
		_, err := conn.Read(rdBuf)
		if err != nil {
			fmt.Println("server read failed: ", err)
			return
		}
		if _, err := conn.Write(wrBuf); err != nil {
			fmt.Println("server write failed: ", err)
			return
		}
	}

	return
}
