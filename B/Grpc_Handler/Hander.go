package Grpc_Handler

import (
	"B/core/Services"
	"B/pb"
	"context"
	"fmt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
)

type ServerPost struct {
	postService *Services.PostService
	pb.UnimplementedPostServicesServer
}

func NewGrpcServer(PostService *Services.PostService)*ServerPost  {
	return &ServerPost{
		postService:                     PostService,
	}
}
func (s *ServerPost)InsertPosts(ctx context.Context,req *pb.InsertPostRequest) (*pb.InsertPostRequest, error){
	fmt.Println("Insert post")
	err := s.postService.Create_Post(ctx,req)
	if err != nil{
		log.Fatalf("err : ",err)
		return nil, status.Errorf(codes.Unimplemented, "method InsertBook not implemented")
	}
	return nil,nil
}
func (s *ServerPost)SearchPosts(ctx context.Context,in *pb.SearchPostRequest) (*pb.SearchPostResponse, error){
	fmt.Println("Search Title")
	arr:= make([]*pb.Posts,0)
	arr,err := s.postService.FindPostTitle(ctx,in.Title,arr)
	if err != nil{
		log.Fatalf("err : ",err)
		return nil, status.Errorf(codes.Unimplemented, "method InsertBook not implemented")
	}
	return &pb.SearchPostResponse{Post: arr},nil
}
func (s *ServerPost)FindAllPosts(ctx context.Context,in *pb.FindAllRequest) (*pb.FindAllResponse, error){
	fmt.Println("Find ALL")
	arr:= make([]*pb.Posts,0)
	arr,err := s.postService.Find_All(ctx,arr)
	if err != nil{
		log.Fatalf("err : ",err)
		return nil, status.Errorf(codes.Unimplemented, "method InsertBook not implemented")
	}
	return &pb.FindAllResponse{Post: arr},nil
}
///func (u *PostService)UpdatePostTitle(ctx context.Context,Title string,posts *pb.UpdateNew)
func (s *ServerPost)UpdatePosts(ctx context.Context,in *pb.UpdateNew) (*pb.UpdatePostResponse, error){
	fmt.Println("Update ")
//	var posts *pb.UpdatePost
	err := s.postService.UpdatePostTitle(ctx,in.UpdateRequest.Title,in)
	if err != nil{
		log.Fatalf("err : ",err)
		return nil, status.Errorf(codes.Unimplemented, "method InsertBook not implemented")
	}
	return &pb.UpdatePostResponse{Title: "Update success "},nil
}
//////
func (s *ServerPost)DeletePosts(ctx context.Context,in *pb.DeletePostRequest) (*pb.DeletePostResponse, error){
	fmt.Println("Delete")
	err := s.postService.DeletePostTitle(ctx,in.Title)
	if err != nil{
		return nil, status.Errorf(codes.Unimplemented, "method DeleteBook not implemented")
	}
	return nil,nil
}