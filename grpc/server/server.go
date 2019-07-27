package server

import (
	"context"
	"errors"
	"log"
	"net"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/CyrivlClth/snowserver/config"
	pb "github.com/CyrivlClth/snowserver/grpc/pb"
)

type server struct{}

func (s *server) NextID(context.Context, *pb.IDRequest) (*pb.IDResponse, error) {
	id, err := config.IDGenerator().NextID()
	return &pb.IDResponse{Id: id}, err
}

func (s *server) GetIDs(ctx context.Context, req *pb.IDsRequest) (*pb.IDsResponse, error) {
	c := req.GetCount()
	if c <= 0 {
		return &pb.IDsResponse{}, errors.New("request parameter [count] must be greater than 0")
	}
	ids := make([]int64, 0)
	for i := int64(0); i < c; i++ {
		id, err := config.IDGenerator().NextID()
		if err != nil {
			return &pb.IDsResponse{Ids: ids}, err
		}
		ids = append(ids, id)
	}
	return &pb.IDsResponse{Ids: ids}, nil
}

func (s *server) Stats(context.Context, *pb.StatsRequest) (*pb.StatsResponse, error) {
	panic("implement me")
}

func RecoverInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		startTime := time.Now()
		defer func() {
			if r := recover(); r != nil {
				err = status.Errorf(codes.Internal, "%s", r)
			}
		}()
		resp, err = handler(ctx, req)
		log.Printf("%dms", time.Since(startTime).Nanoseconds()/10000)
		return
	}
}

func Run(addr string) (err error) {
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		return
	}
	srv := grpc.NewServer(grpc.UnaryInterceptor(
		RecoverInterceptor(),
	))
	pb.RegisterSnowflakeServer(srv, &server{})
	log.Printf("grpc sever listening at %v", lis.Addr())
	return srv.Serve(lis)
}
