package main
import (
	"fmt"
	"net"
	"log"
	"os"
	"bytes"
)
func main() {
	//建立socket，监听端口，简单的golang socket只需要三步即可建立一个简单的socket server
	netListen, err := net.Listen("tcp", ":1024")//1. 利用net包中的listen方法监听ip：port
	CheckError(err)
	defer netListen.Close()
	Log("Waiting for clients")
	for {
		conn, err := netListen.Accept()//2. 一直阻塞，知道有一个新的client连接才往下操作
		if err != nil {
			continue
		}
		Log(conn.RemoteAddr().String(), " tcp connect success")
		go handleConnection(conn)//3. 利用conn 的read write方法即可对socket进行读和写操作
	}
}
//处理连接
func handleConnection(conn net.Conn) {
	buffer := make([]byte, 2048)
	//for {
		n, err := conn.Read(buffer)

		if err != nil {
			Log(conn.RemoteAddr().String(), " connection error: ", err)
			return
		}
	        s := string(buffer[:n])
		Log(conn.RemoteAddr().String(), "receive data string:\n",s)
	        if s == "hello world!" {
			b := bytes.Buffer{}
			b.WriteString(s)
			b.WriteString("client !!!")
			conn.Write([]byte(b.String()))
		}

	//}
}
func Log(v ...interface{}) {
	log.Println(v...)
}
func CheckError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}
