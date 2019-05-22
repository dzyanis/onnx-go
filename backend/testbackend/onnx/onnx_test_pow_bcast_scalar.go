package onnxtest

// this file is auto-generated... DO NOT EDIT

import (
	"github.com/owulveryck/onnx-go/backend/testbackend"
	"gorgonia.org/tensor"
)

func init() {
	testbackend.Register("Pow", "TestPowBcastScalar", NewTestPowBcastScalar)
}

// NewTestPowBcastScalar version: 3.
func NewTestPowBcastScalar() *testbackend.TestCase {
	return &testbackend.TestCase{
		OpType: "Pow",
		Title:  "TestPowBcastScalar",
		ModelB: []byte{0x8, 0x3, 0x12, 0xc, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2d, 0x74, 0x65, 0x73, 0x74, 0x3a, 0x56, 0xa, 0xe, 0xa, 0x1, 0x78, 0xa, 0x1, 0x79, 0x12, 0x1, 0x7a, 0x22, 0x3, 0x50, 0x6f, 0x77, 0x12, 0x15, 0x74, 0x65, 0x73, 0x74, 0x5f, 0x70, 0x6f, 0x77, 0x5f, 0x62, 0x63, 0x61, 0x73, 0x74, 0x5f, 0x73, 0x63, 0x61, 0x6c, 0x61, 0x72, 0x5a, 0xf, 0xa, 0x1, 0x78, 0x12, 0xa, 0xa, 0x8, 0x8, 0x1, 0x12, 0x4, 0xa, 0x2, 0x8, 0x3, 0x5a, 0xb, 0xa, 0x1, 0x79, 0x12, 0x6, 0xa, 0x4, 0x8, 0x1, 0x12, 0x0, 0x62, 0xf, 0xa, 0x1, 0x7a, 0x12, 0xa, 0xa, 0x8, 0x8, 0x1, 0x12, 0x4, 0xa, 0x2, 0x8, 0x3, 0x42, 0x2, 0x10, 0x9},

		/*

		   &pb.NodeProto{
		     Input:     []string{"x", "y"},
		     Output:    []string{"z"},
		     Name:      "",
		     OpType:    "Pow",
		     Attributes: ([]*pb.AttributeProto) <nil>
		   ,
		   },


		*/

		Input: []tensor.Tensor{

			tensor.New(
				tensor.WithShape(3),
				tensor.WithBacking([]float32{1, 2, 3}),
			),

			tensor.New(
				tensor.WithShape(1),
				tensor.WithBacking([]float32{2}),
			),
		},
		ExpectedOutput: []tensor.Tensor{

			tensor.New(
				tensor.WithShape(3),
				tensor.WithBacking([]float32{1, 4, 9}),
			),
		},
	}
}
