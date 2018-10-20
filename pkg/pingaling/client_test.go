package pingaling

import (
	"bytes"
	"context"
	"io"
	"log"

	"github.com/stretchr/testify/mock"
)

type ClientMock struct {
	mock.Mock
}

func (cm *ClientMock) Get(ctx context.Context, url string, ts interface{}) error {
	log.Println("Mocked client Get")
	return nil
}

func (cm *ClientMock) Delete(ctx context.Context, url string, ts interface{}) error {
	log.Println("Mocked client Delete")
	ts.(*DeleteMsg).Message = "Test Delete Message"
	return nil
}

func (cm *ClientMock) Post(ctx context.Context, url string, body io.Reader) (bytes.Buffer, error) {
	log.Println("Mocked client Post")
	return bytes.Buffer{}, nil
}
