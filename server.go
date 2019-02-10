package stocktables

import (
	"fmt"
	"github.com/skratchdot/open-golang/open"
	"net"
	"net/http"
)

func Serve(addr string, handler http.Handler) error {
	srv := &http.Server{Addr: addr, Handler: handler}
	ln, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}
	port := getPort(ln)
	listen := make(chan bool)
	go func() {
		fmt.Println("http://127.0.0.1:" + port + "/")
		<-listen
		open.Run("http://127.0.0.1:" + port + "/")
	}()
	listen <- true
	return srv.Serve(ln.(*net.TCPListener))

}

func getPort(listen net.Listener) string {
	addr := listen.Addr().String()
	_, port, err := net.SplitHostPort(addr)

	if err != nil {
		panic(err)
	}
	return port
}
