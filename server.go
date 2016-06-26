package main
import (
    "fmt"
    "bufio"
    "net"
    "time"
)
var print = fmt.Println
var scan = fmt.Scanln

var connMap map[string] *net.TCPConn

func main() {
    var tcpAddr, _ = net.ResolveTCPAddr("tcp", "127.0.0.1:9999")
    var tcpListener, _ = net.ListenTCP("tcp", tcpAddr)
    connMap = make(map[string] *net.TCPConn)
    defer tcpListener.Close()

    for {
        var tcpConn, err = tcpListener.AcceptTCP()
        if err != nil {
            print("continue", err)
            continue
        }
        var clientIP = tcpConn.RemoteAddr().String()
        print(clientIP, "connected")
        connMap[clientIP] = tcpConn
        go tcpPipe(tcpConn)
    }

}

func tcpPipe(conn *net.TCPConn) {
    var ip string = conn.RemoteAddr().String()
    defer func() {
        print(ip, "disconnected")
        conn.Close()
    }()

    var reader = bufio.NewReader(conn)

    for {
        msg, err := reader.ReadString('\n')
        if err != nil {
            return
        }

        print(ip, ":", msg)
        msgstr := time.Now().String() + "\n"
        msgstr = ip + ":" + msg
        broadcast(msgstr)
    }
}


func broadcast(msg string) {
    var b = []byte(msg)
    for _, conn := range connMap {
        conn.Write(b)
    }
}
