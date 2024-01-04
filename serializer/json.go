package serializer

import (
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

func ProtobufToJSON(message proto.Message) (string, error) {
	marshaler := protojson.MarshalOptions{
		UseProtoNames:     true,
		Indent:            " ",
		EmitDefaultValues: true,
		UseEnumNumbers:    false,
	}

	jsonBytes, err := marshaler.Marshal(message)
	if err != nil {
		return "", err
	}

	return string(jsonBytes), nil
}
