package main
import (
    "fmt"
    "bufio"
    "net"
    "time"
)
var print = fmt.Println
var scan = fmt.Scanln


func main() {
    var tcpAddr *net.TCPAddr
    tcpAddr, _ = net.ResolveTCPAddr("tcp", "127.0.0.1:9999")
    var tcpListener, _ = net.ListenTCP("tcp", tcpAddr)

    defer tcpListener.Close()

    for {
        var tcpConn, err = tcpListener.AcceptTCP()
        if err != nil {
            continue
        }
        print("A client connected: ", tcpConn.RemoteAddr().String())
        go tcpPipe(tcpConn)
    }

}

func tcpPipe(conn *net.TCPConn) {
    var ip string = conn.RemoteAddr().String()
    defer func() {
        print("disconnected: ", ip)
        conn.Close()
    }()

    var reader = bufio.NewReader(conn)

    for {
        msg, err := reader.ReadString('\n')
        if err != nil {
            return
        }

        print(string(msg))
        msgstr := time.Now().String() + "\n"
        var b []byte = []byte(msgstr)
        conn.Write(b)
    }
}
