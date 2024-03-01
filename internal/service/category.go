package service

import (
	"context"

	"github.com/Cristuker/go-grpc/internal/database"
	"github.com/Cristuker/go-grpc/internal/pb"
)

type CategoryService struct {
	pb.UnimplementedCategoryServiceServer
	CategoryDB database.Category
}

func NewCategoryService(categoryDB database.Category) *CategoryService {
	return &CategoryService{
		CategoryDB: categoryDB,
	}
}

// ctx headers essas coisas, in body
func (c *CategoryService) CreateCategory(ctx context.Context, in *pb.CreateCategoryRequest) (*pb.CategoryResponse, error) {
	result, err := c.CategoryDB.Create(in.Name, in.Description)
	if err != nil {
		return nil, err
	}
	categoryResponse := &pb.Category{
		Id:          result.ID,
		Name:        result.Name,
		Description: result.Description,
	}

	return &pb.CategoryResponse{
		Category: categoryResponse,
	}, nil
}
