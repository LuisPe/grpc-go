package main

import (
	"context"
	"fmt"
	"io"
	"log"

	"github.com/luispe/grpc-go/blog/blogpb"

	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Blog client")
	opts := grpc.WithInsecure()
	cc, err := grpc.Dial("localhost:50051", opts)

	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}
	defer cc.Close()
	c := blogpb.NewBlogServiceClient(cc)

	// create a blog
	fmt.Println("Creating a blog")
	blog := &blogpb.Blog{
		AuthorId: "LuisPe",
		Title:    "My first blog",
		Content:  "Content of the fist blog",
	}

	cbr, err := c.CreateBlog(context.Background(), &blogpb.CreateBlogRequest{Blog: blog})

	if err != nil {
		log.Fatalf("Unexpected error: %v\n", err)
	}
	fmt.Printf("Blog has been created: %v\n", cbr)

	// Read Blog
	blogID := cbr.GetBlog().GetId()
	fmt.Println("Reading a blog")
	_, err2 := c.ReadBlog(context.Background(), &blogpb.ReadBlogRequest{BlogId: "5ce5694fe3bd5c193a88c8f6"})
	if err2 != nil {
		fmt.Printf("Error happened while reading: %v \n", err2)
	}

	readBlogReq := &blogpb.ReadBlogRequest{BlogId: blogID}
	readBlogRes, readBlogErr := c.ReadBlog(context.Background(), readBlogReq)
	if readBlogErr != nil {
		fmt.Printf("Error happened while reading: %v \n", readBlogErr)
	}

	fmt.Printf("Blog was read: %v \n", readBlogRes)

	// update Blog
	ub := &blogpb.Blog{
		Id:       blogID,
		AuthorId: "Changed Author",
		Title:    "My First Blog (edited)",
		Content:  "Content of the first blog, with some awesome additions!",
	}
	ubr, uErr := c.UpdateBlog(context.Background(), &blogpb.UpdateBlogRequest{Blog: ub})
	if uErr != nil {
		fmt.Printf("Error happened while updating: %v \n", uErr)
	}
	fmt.Printf("Blog was updated: %v\n", ubr)

	// delete Blog
	dr, dErr := c.DeleteBlog(context.Background(), &blogpb.DeleteBlogRequest{BlogId: blogID})

	if dErr != nil {
		fmt.Printf("Error happened while deleting: %v \n", dErr)
	}
	fmt.Printf("Blog was deleted: %v \n", dr)

	// list Blogs
	stream, err := c.ListBlog(context.Background(), &blogpb.ListBlogRequest{})
	if err != nil {
		log.Fatalf("error while calling ListBlog RPC: %v", err)
	}
	for {
		res, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Something happened: %v", err)
		}
		fmt.Println(res.GetBlog())
	}
}
