package servers

import (
	"context"

	pb "algohub.dev/backend/proto"
	"algohub.dev/backend/structs"
	"gorm.io/gorm"
)

type ExampleServer struct {
	pb.ExampleServer
	DB  *gorm.DB
	Env *structs.Env
}

// Example echo method
func (s *ExampleServer) Echo(ctx context.Context, in *pb.EchoRequest) (*pb.EchoResponse, error) {
	return &pb.EchoResponse{Message: in.Message}, nil
}
