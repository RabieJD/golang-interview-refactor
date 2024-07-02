package utils

import (
	"bytes"
	"io"
)

type TTForm struct {
	Field1 string `form:"field1"`
	Field2 int    `form:"field2"`
}

type TTGetRequestForm struct {
	Name        string
	Body        io.ReadCloser
	Expected    TTForm
	ExpectError bool
}

func GenerateTTGetRequestForm() []TTGetRequestForm {
	return []TTGetRequestForm{
		{
			Name:        "Valid form data",
			Body:        io.NopCloser(bytes.NewBufferString("field1=test&field2=123")),
			Expected:    TTForm{Field1: "test", Field2: 123},
			ExpectError: false,
		},
		{
			Name:        "Missing field data",
			Body:        io.NopCloser(bytes.NewBufferString("field1=test")),
			Expected:    TTForm{Field1: "test", Field2: 0},
			ExpectError: false,
		},
		{
			Name:        "Empty Body",
			Body:        nil,
			Expected:    TTForm{},
			ExpectError: true,
		},
	}
}
