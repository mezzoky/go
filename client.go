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
    
    var conn, _ = net.DialTCP("tcp", nil, tcpAddr)

    defer conn.Close()

    print("connected!")

    var quitSemaphone = make(chan bool)
    go onMessageReceived(conn, quitSemaphone)

    var b = []byte("time\n")
    conn.Write(b)

    <- quitSemaphone
}

func onMessageReceived(conn *net.TCPConn, quitSemaphone chan bool) {
    var reader = bufio.NewReader(conn)
    for {
        var msg, err = reader.ReadString('\n')
        print(msg)
        if err != nil {
            quitSemaphone <- true
            break
        }
        time.Sleep(time.Second)
        var b []byte = []byte(msg)
        conn.Write(b)
    }
}
