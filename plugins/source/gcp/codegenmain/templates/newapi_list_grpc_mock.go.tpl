// Code generated by codegen; DO NOT EDIT.

package {{.Service}}

import (
	"context"
	"net"
	"fmt"
	"testing"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"github.com/cloudquery/plugin-sdk/faker"
	"github.com/cloudquery/plugins/source/gcp/client"
  {{range .MockImports}}
  "{{.}}"
  {{end}}
	{{if .ProtobufImport}}
  pb "{{.ProtobufImport}}"
  {{end}}
	"google.golang.org/api/option"
)


func create{{.SubService | ToCamel}}() (*client.Services, error) {
	fakeServer := &fake{{.SubService | ToCamel}}Server{}
	l, err := net.Listen("tcp", "localhost:0")
	if err != nil {
		return nil, fmt.Errorf("failed to listen: %w", err)
	}
	gsrv := grpc.NewServer()
	pb.{{.RegisterServerName}}(gsrv, fakeServer)
	fakeServerAddr := l.Addr().String()
	go func() {
		if err := gsrv.Serve(l); err != nil {
			panic(err)
		}
	}()

	// Create a client.
	svc, err := {{.Service}}.{{.NewFunctionName}}(context.Background(),
		option.WithEndpoint(fakeServerAddr),
		option.WithoutAuthentication(),
		option.WithGRPCDialOption(grpc.WithTransportCredentials(insecure.NewCredentials())),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create grpc client: %w", err)
	}

	return &client.Services{
		{{.Service | ToCamel}}{{.ClientName}}: svc,
	}, nil
}

type fake{{.SubService | ToCamel}}Server struct {
	pb.{{.UnimplementedServerName}}
}

func (f *fake{{.SubService | ToCamel}}Server) {{.ListFunctionName}}(context.Context, *pb.{{.RequestStructName}}) (*pb.{{.ResponseStructName}}, error) {
	resp := pb.{{.ResponseStructName}}{}
	if err := faker.FakeObject(&resp); err != nil {
		return nil, fmt.Errorf("failed to fake data: %w", err)
	}
	resp.NextPageToken = ""
	return &resp, nil
}

func Test{{.SubService | ToCamel}}(t *testing.T) {
	client.MockTestHelper(t, {{.SubService | ToCamel}}(), create{{.SubService | ToCamel}}, client.TestOptions{})
}
