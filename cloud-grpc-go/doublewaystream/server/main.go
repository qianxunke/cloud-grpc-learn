package main

import (
	"bufio"
	"cloud_grpc_go/pbfs/double_way_stream"
	"google.golang.org/grpc"
	"io"
	"log"
	"net"
	"os"
	"strconv"
)

type DWS struct {
}

func (s *DWS) DoubleWayStreamFun(stream double_way_stream.DoubleWayStreamService_DoubleWayStreamFunServer) error {
	ctx := stream.Context()
	//定义一个向客户端发送消息的go rutum
	go func() {
		input := bufio.NewReader(os.Stdin)
		for {
			// 获取 命令行输入的字符串， 以回车 \n 作为结束标志
			inStr, _ := input.ReadString('\n')

			// 向服务端发送 指令
			if err := stream.Send(&double_way_stream.ResponseMessage{RspMsg:inStr}); err != nil {
				return
			}
		}

	}()

	for {
		select {
		case <-ctx.Done():
			log.Println("收到客户端通过context发出的终止信号")
			return ctx.Err()
		default:
			// 接收从客户端发来的消息
			ReqStream, err := stream.Recv()
			if err == io.EOF {
				log.Println("客户端发送的数据流结束")
				return nil
			}
			if err != nil {
				log.Println("接收数据出错:", err)
				return err
			}
			switch ReqStream.ReqMsg {

			case "结束通话\n":
				log.Println("收到'结束对话'指令")
				if err := stream.Send(&double_way_stream.ResponseMessage{RspMsg: "收到结束指令"}); err != nil {
					return err
				}
				// 收到结束指令时，通过 return nil 终止双向数据流
				return nil
			case "返回数据流\n":
				log.Println("收到'返回数据流'指令")
				// 收到 收到'返回数据流'指令， 连续返回 10 条数据
				for i := 0; i < 10; i++ {
					if err := stream.Send(&double_way_stream.ResponseMessage{RspMsg: "数据流 #" + strconv.Itoa(i)}); err != nil {
						return err
					}
				}
				break
			default:
				// 缺省情况下， 返回 '服务端返回: ' + 输入信息
				log.Printf("[收到消息]: %s", ReqStream.ReqMsg)
				if err := stream.Send(&double_way_stream.ResponseMessage{RspMsg: "服务端返回: " + ReqStream.ReqMsg}); err != nil {
					return err
				}
			}
		}

	}
}



func main() {
	log.Println("启动服务端...")
	server := grpc.NewServer()

	// 注册 ChatServer
	double_way_stream.RegisterDoubleWayStreamServiceServer(server, &DWS{})

	address, err := net.Listen("tcp", ":3000")
	if err != nil {
		panic(err)
	}

	if err := server.Serve(address); err != nil {
		panic(err)
	}
}
