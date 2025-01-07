package grpc_server

import (
	"context"

	parserv1 "github.com/acronix0/REST-API-Go-protos/gen/go/parser"
	srv "github.com/acronix0/XML-Parser-Golang/internal/service"
	"google.golang.org/grpc"
)
type serverApi struct{
	parserv1.UnimplementedParserServer
	parser srv.Parser
}

func RegisterServer(grpc *grpc.Server, parser srv.Parser){
	parserv1.RegisterParserServer(grpc, &serverApi{parser: parser})
}

func (s *serverApi) Parse(
	ctx context.Context,
	req *parserv1.ParseRequest,
)(*parserv1.ParseResponse, error){
		err := s.parser.Parse(ctx, req.GetFilePath())
		if err != nil {
			return nil, err
		}
		return &parserv1.ParseResponse{Status: "200", Message: "Ok"}, nil
}