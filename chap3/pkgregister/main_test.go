package pkgregister

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func startTestPackageServer() *httptest.Server {
	ts := httptest.NewServer(http.HandlerFunc(packageRegHandler))
	return ts
}

func Test_registerPackageData(t *testing.T) {
	ts := startTestPackageServer()
	defer ts.Close()
	p := pkgData{
		Name:    "mypackage",
		Version: "0.1",
	}
	resp, err := registerPackageData(ts.URL, p)
	if err != nil {
		t.Fatal(err)
	}
	expected := "mypackage-0.1"
	if resp.ID != expected {
		t.Errorf("expected package ID %v, got %v", expected, resp.ID)
	}
}

func Test_registerPackageDataEmtpy(t *testing.T) {
	ts := startTestPackageServer()
	defer ts.Close()
	p := pkgData{}
	resp, err := registerPackageData(ts.URL, p)
	if err == nil {
		t.Errorf("expected error to be non-nil, got %v", err)
	}
	if len(resp.ID) != 0 {
		t.Errorf("expected empty package ID, got %v", resp.ID)
	}
}
