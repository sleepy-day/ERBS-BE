package grpcserver

import (
	"errors"
	"github.com/sleepy-day/ERBS-BE/src/erbsdb"
	"github.com/sleepy-day/ERBS-BE/src/protos"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"
)

var serv *grpc.Server

type GrpcServer struct {
	pb.UnimplementedErbsApiServer
}

func (grpc *GrpcServer) GetCharacters(context.Context, *pb.NoParameters) (*pb.Characters, error) {
	characters, err := erbsdb.GetCharacters()
	if err != nil {
		return nil, err
	}

	return &pb.Characters{Characters: characters}, err
}

func (grpc *GrpcServer) GetCharacterDetails(_ context.Context, req *pb.GetCharDetails) (*pb.CharacterProfile, error) {
	if req.Code <= 0 {
		return nil, errors.New("Invalid character code.")
	}
	character, err := erbsdb.GetCharacterProfile(req.Code)
	if err != nil {
		return nil, err
	}

	return character, nil
}

func StartServer() {
	listen, err := net.Listen("tcp", ":8102")
	if err != nil {
		panic(err)
	}

	serv = grpc.NewServer()
	pb.RegisterErbsApiServer(serv, new(GrpcServer))
	reflection.Register(serv)
	err = serv.Serve(listen)
	if err != nil {
		panic(err)
	}
}
