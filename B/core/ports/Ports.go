package ports

import (
	"B/adapters/database"
	"B/pb"
	"context"
)

type PostsServices interface {
	CreatePost(ctx context.Context,Request *pb.InsertPostRequest)(error)
	SearchPost(ctx context.Context,Title string,post []*pb.Posts)([]*pb.Posts,error)
	FindPost(ctx context.Context,post []*pb.Posts)([]*pb.Posts,error)
	DeletePost(ctx context.Context,Title string)error
	UpdateBook(ctx context.Context,Title string, data *pb.UpdateNew)error
}

func InitPostRepositoryPort(mongoDb *database.PostServiceImpl)PostsServices  {
	return mongoDb
}