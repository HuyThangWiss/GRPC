package main

import (
	"A/Interceptor"
	"A/pb"
	"context"
	"fmt"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	"google.golang.org/grpc"
	"log"
)

func main() {
	conn, err := grpc.Dial(
		":1111",
		grpc.WithInsecure(),
		grpc.WithUnaryInterceptor(
			grpc_middleware.ChainUnaryClient(
				Interceptor.DateLogClientInterceptor,
				Interceptor.MethodLogClientInterceptor,
			),
		),
	)
	if err != nil {
		fmt.Println(err)
	}
	defer conn.Close()

	BookServ := pb.NewPostServicesClient(conn)
	fmt.Println("start call gRpc")
/*
	_,err1 := BookServ.InsertPosts(context.Background(),&pb.InsertPostRequest{
		Title:   "New Title Db",
		Email:   "Buu@gmail.com",
		Code:    "M4A1",
		Content: "Your Content",
		Image:   "Super.png",
		User:    "Vegetate",
	})
	if err1 != nil{
		log.Fatal("err : ",err1)
		return
	}
	fmt.Println("Insert successfully")
*/
	/*
	post,err1 := BookServ.FindAllPosts(context.Background(),&pb.FindAllRequest{})
	if err1 != nil{
		log.Fatal("err : ",err1)
		return
	}
	fmt.Println("Data : ",post)
	 */
	/*
	post,err1 := BookServ.SearchPosts(context.Background(),&pb.SearchPostRequest{Title: "Hello 500 ae"})
	if err1 != nil{
		log.Fatal("err : ",err1)
		return
	}
	fmt.Println("Data Search : ",post)
	*/
/*
	_,err1 := BookServ.DeletePosts(context.Background(),&pb.DeletePostRequest{Title: "New Title Db "})
	if err1 != nil{
		log.Fatal("err : ",err1)
		return
	}
	fmt.Println("Delete successfully")

*/

	_,err1 := BookServ.UpdatePosts(context.Background(),&pb.UpdateNew{
		UpdatePost:    &pb.UpdatePost{
			Email:   "New",
			Code:    "New",
			Content: "New",
			Image:   "New",
			User:    "New",
		},
		UpdateRequest: &pb.UpdatePostRequest{Title: "New Title Db"},
	})
	if err1 != nil{
		log.Fatal("err : ",err1)
		return
	}
	fmt.Println("Update successfully")


}
















