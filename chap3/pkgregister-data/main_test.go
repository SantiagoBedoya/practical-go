package pkgregister

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"
)

func startTestPackageServer() *httptest.Server {
	ts := httptest.NewServer(http.HandlerFunc(packageRegHandler))
	return ts
}

func Test_registerPackageData(t *testing.T) {
	ts := startTestPackageServer()
	defer ts.Close()
	p := pkgData{
		Name:     "mypackage",
		Version:  "0.1",
		Filename: "mypackage-0.1.tar.gz",
		Bytes:    strings.NewReader("data"),
	}
	pResult, err := registerPackageData(createHTTPClientWithTimeout(5*time.Second), ts.URL, p)
	if err != nil {
		t.Fatal(err)
	}
	expectedID := fmt.Sprintf("%s-%s", p.Name, p.Version)
	if pResult.ID != expectedID {
		t.Errorf("expected package ID %v, got %v", expectedID, pResult.ID)
	}
	if pResult.Filename != p.Filename {
		t.Errorf("expected package filename %v, got %v", p.Filename, pResult.Filename)
	}
	if pResult.Size != 4 {
		t.Errorf("expected package size %v, got %v", 4, pResult.Size)
	}
}
