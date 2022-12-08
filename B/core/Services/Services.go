package Services

import (
	"B/core/ports"
	"B/pb"
	"context"
	"log"
)

type PostService struct {
	PostRepository ports.PostsServices
}

func NewUserService(userRepositoryPort ports.PostsServices) *PostService {
	return &PostService{
		PostRepository: userRepositoryPort,
	}
}

func (u *PostService)Create_Post(ctx context.Context,req *pb.InsertPostRequest)error  {
	err := u.PostRepository.CreatePost(ctx,&pb.InsertPostRequest{
		Title:   req.Title,
		Email:   req.Email,
		Code:    req.Code,
		Content: req.Content,
		Image:   req.Image,
		User:    req.User,
	})
	if err != nil {
		log.Fatalf("err : ", err)
		return  err
	}
	return nil
}

func (u *PostService)Find_All(ctx context.Context,posts []*pb.Posts)([]*pb.Posts,error)  {
	
	posts,err := u.PostRepository.FindPost(ctx,posts)
	if err != nil{
		log.Fatalf("err : ",err)
		return nil,err
	}
	return posts,nil
}

func (u *PostService)FindPostTitle(ctx context.Context,Title string,posts []*pb.Posts)([]*pb.Posts,error)  {
	posts,err := u.PostRepository.SearchPost(ctx,Title,posts)
	if err != nil{
		log.Fatalf("err : ",err)
		return nil,err
	}
	return posts,nil
}
func (u *PostService)UpdatePostTitle(ctx context.Context,Title string,posts *pb.UpdateNew)  error{
	err := u.PostRepository.UpdateBook(ctx,Title,posts)
	if err != nil {
		log.Fatalf("err ", err)
		return err
	}
	return nil
}

func (u *PostService)DeletePostTitle(ctx context.Context,Title string)error  {
	err := u.PostRepository.DeletePost(ctx,Title)
	if err != nil {
		log.Fatalf("err ", err)
		return err
	}
	return nil
}

















