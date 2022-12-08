package database

import (
	"B/adapters/mapper"
	"B/pb"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
)

type PostServiceImpl struct {
	postCollection *mongo.Collection
	ctx            context.Context
}
func NewMonGoDb(collection *mongo.Collection)*PostServiceImpl{
	return &PostServiceImpl{
		postCollection: collection,
	}
}
func (p *PostServiceImpl)CreatePost(ctx context.Context,Request *pb.InsertPostRequest)(error){
	postModel := &pb.InsertPostRequest{
		Title:   Request.Title,
		Email:   Request.Email,
		Code:    Request.Code,
		Content: Request.Content,
		Image:   Request.Image,
		User:    Request.User,
	}
	_,err := p.postCollection.InsertOne(p.ctx,postModel)
	if err != nil {
		return err
	}
	return nil
}

func (p *PostServiceImpl)SearchPost(ctx context.Context,Title string,post []*pb.Posts)([]*pb.Posts,error){
	filter,err := p.postCollection.Find(ctx,bson.M{"title":Title})
	if err != nil{
		log.Fatal(err)
	}
	if err=filter.All(context.TODO(),&post);err != nil{
		log.Fatal(err)
	}
	return post,nil
}
func (p *PostServiceImpl)FindPost(ctx context.Context,post []*pb.Posts)([]*pb.Posts,error){
	cursor,err := p.postCollection.Find(p.ctx,bson.M{})
	if err != nil{
		log.Fatalf("err ",err)
		return nil,err
	}
	if err = cursor.All(ctx,&post); err != nil{
		fmt.Println("err find ",err)
		return nil,err
	}
	return post,nil
}
func (p *PostServiceImpl)DeletePost(ctx context.Context,Title string)error{
	filter := bson.M{"title":Title}
	_,err :=p.postCollection.DeleteOne(ctx,filter)
	if err != nil{
		log.Fatal(err)
	}
	return nil
}
func (p *PostServiceImpl)UpdateBook(ctx context.Context,Title string,data *pb.UpdateNew)error{
	var req ,_ = mapper.ToDoc(data.UpdatePost)
	_,err := p.postCollection.UpdateOne(ctx,bson.M{"title":Title}, bson.M{"$set": req})
	if err != nil{
		return err
	}
	return nil
}

















