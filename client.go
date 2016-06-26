package main
import (
    "fmt"
    "bufio"
    "net"
)
var print = fmt.Println
var scan = fmt.Scanln

func main() {
    var tcpAddr, _ = net.ResolveTCPAddr("tcp", "127.0.0.1:9999")
    
    var conn, _ = net.DialTCP("tcp", nil, tcpAddr)

    defer conn.Close()

    print("connected to", conn.RemoteAddr().String())

    go onMessageReceived(conn)

    for {
        var msg string
        scan(&msg)
        if msg == "quit" {
            break
        }
        b := []byte(msg + "\n")
        conn.Write(b)
    }
}

func onMessageReceived(conn *net.TCPConn) {
    var reader = bufio.NewReader(conn)
    for {
        var msg, err = reader.ReadString('\n')
        print("from server:", msg)
        if err != nil {
            break
        }
    }
}
