package znet

import (
	"fmt"
	"github.com/Menah3m/zinx/ziface"
	"net"
)

type Server struct {
	Name string
	IP string
	Port int
	IPVersion string

}

//Start一个server
func (s *Server) Start()  {
	fmt.Printf("[start] Server listenner at ip：%v,port %d is starting\n", s.IP, s.Port)

	go func() {
		//1.获取TCP addr
		addr, err := net.ResolveTCPAddr(s.IPVersion, fmt.Sprintf("%s:%d", s.IP, s.Port))
		if err != nil {
			fmt.Println("tcp addr resolve error:", err)
			return
		}

		//2.监听服务器的地址
		listenner, err := net.ListenTCP(s.IPVersion, addr)
		if err != nil {
			fmt.Println("listen", s.IPVersion, "err", err)
			return
		}

		fmt.Println("start zinx server success.", s.Name, "Listenning...")

		//3.阻塞等待客户端连接，处理客户端请求（读写）
		for {
			conn, err := listenner.AcceptTCP()
			if err != nil {
				fmt.Println("accept err ", err)
				continue
			}

			// 已经与客户端建立连接，就可以做业务了，这里以回显最大为512字节长度的内容为例
			go func() {
				for {
					buf := make([]byte, 512)
					cnt, err := conn.Read(buf)
					if err != nil {
						fmt.Println("recv buff err ", err)
						continue
					}
					
					_, err = conn.Write(buf[:cnt])
					if err != nil {
						fmt.Println("write back buf err", err)
						continue
					}
				}
			}()
		}
	}()



}

func (s *Server) Stop()  {
	//Todo 将服务器的资源 状态以及开启的连接信息关闭和释放
}

func (s *Server) Serve()  {
	//启动server的服务功能
	s.Start()

	//Todo 做一些启动服务器之外的业务
	//阻塞状态
	select {

	}


}
//初始化server的方法
func New(name string) ziface.IServer {
	s := &Server{
		Name: name,
		IPVersion: "tcp4",
		IP: "0.0.0.0",
		Port: 8999,
	}
	return s
}