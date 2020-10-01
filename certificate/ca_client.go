package certificate

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/gogo/protobuf/jsonpb"
	"github.com/gogo/protobuf/types"
	grpc "google.golang.org/grpc"
)

const (
	address     = "localhost:50051"
	defaultName = "world"
)


func main(port int, csr string) {
	// Set up a connection to the server.
	conn, err := grpc.Dial(fmt.Sprintf("localhost:%d", port), grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	c := NewIstioCertificateServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	metadata := map[string]interface{}{
		"TrustDomain":"mh.cluser.local",
		"ClusterID":    "MH Mock ClusterID",
		"WorkloadName": "MH Mock WorkloadName",
		"WorkloadNamespace": "MH Mock WorkloadNamespace",
		"WorkloadIPs":  []string{"192.168.10.69"},
		"PodNamespace": "MH Mock PodNamespace",
		"PodIP": "10.2.10.12",
		"ServiceName": "MH Mock ServiceName",
	}
	metaStruct, _ := mapToStruct(metadata)

	r, err := c.CreateCertificate(ctx, &IstioCertificateRequest{
		Csr: csr,
		Metadata: metaStruct,
	})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r.GetCertChain())
}

func Run(port int, csr string) {
	main(port, csr)
}

func mapToStruct(msg map[string]interface{}) (*types.Struct, error) {
	jb, err := json.Marshal(msg)
	if err != nil {
		return nil, fmt.Errorf("parse Metadata Map[string]interface{} to Json failed: %v", err)
	}
	jbu := jsonpb.Unmarshaler{AllowUnknownFields: true}
	pb := &types.Struct{}
	err = jbu.Unmarshal(bytes.NewReader(jb), pb)
	if err != nil {
		return nil, fmt.Errorf("parse Metadata JSON to proto.struct failed: %v", err)
	}
	return pb, nil
}
