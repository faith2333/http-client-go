package http_client_go

import (
	"context"
	"os"
	"testing"
)

var client Interface

func TestMain(m *testing.M) {
	client = NewDefaultClient()
	os.Exit(m.Run())
}

func TestDefaultHttpClient_Do(t *testing.T) {
	resp, err := client.WithContentTypeJson().WithHost("https://www.baidu.com").
		WithURL("/search").WithMethod(MethodGET).Do(context.Background()).RespBytes()
	if err != nil {
		t.Logf("test failed: %s", err.Error())
		t.FailNow()
	}

	t.Logf("response: %v", string(resp))
}
