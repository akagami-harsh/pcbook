package serializer

import (
	"fmt"
	"os"

	"google.golang.org/protobuf/proto"
)

// WriteProtobufToJSONFile writes protcol buffer message to JSON file
func WriteProtobufToJSONFile(message proto.Message, filename string) error {
	data, err := ProtobufToJSON(message)
	if err != nil {
		return fmt.Errorf("error cannot marshal message to JSON: %w", err)
	}

	err = os.WriteFile(filename, []byte(data), 0644)
	if err != nil {
		return fmt.Errorf("error cannot write JSON data to file: %w", err)
	}

	return nil
}

// WriteprotobufToBinaryFile writes protcol buffer message to binary file
func WriteProtobufToBinaryFile(message proto.Message, filename string) error {
	data, err := proto.Marshal(message)
	if err != nil {
		return fmt.Errorf("error cannot marshal message to binary: %w", err)
	}

	err = os.WriteFile(filename, data, 0644)
	if err != nil {
		return fmt.Errorf("error cannot write binary data to file: %w", err)
	}
	return nil
}

// ReadProtobufFromBinaryFile reads protcol buffer message form binary file
func ReadProtobufFromBinaryFile(filename string, message proto.Message) error {
	data, err := os.ReadFile(filename)
	if err != nil {
		return fmt.Errorf("cannot read binray data form file: %w", err)
	}

	err = proto.Unmarshal(data, message)
	if err != nil {
		return fmt.Errorf("cannot unmarshal binary data to proto message: %w", err)
	}

	return nil
}
