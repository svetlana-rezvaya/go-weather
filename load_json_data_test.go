package main

import (
	"net/http"
	"net/http/httptest"
	"reflect"
	"strings"
	"testing"
)

func Test_loadJSONData(test *testing.T) {
	type args struct {
		httpClient   *http.Client
		url          string
		responseData interface{}
	}
	type data struct {
		name               string
		testServerHandler  http.HandlerFunc
		args               args
		wantedResponseData interface{}
		wantedErrStr       string
	}
	type testPayload struct {
		FieldOne int
		FieldTwo string
	}

	tests := []data{
		data{
			name: "success",
			testServerHandler: func(writer http.ResponseWriter, request *http.Request) {
				writer.Write([]byte(`{"FieldOne": 23, "FieldTwo": "test"}`))
			},
			args: args{
				httpClient:   &http.Client{},
				url:          "http://example.com/",
				responseData: &testPayload{},
			},
			wantedResponseData: &testPayload{FieldOne: 23, FieldTwo: "test"},
			wantedErrStr:       "",
		},
		data{
			name: "error with request creating",
			testServerHandler: func(writer http.ResponseWriter, request *http.Request) {
			},
			args: args{
				httpClient:   &http.Client{},
				url:          ":",
				responseData: &testPayload{},
			},
			wantedResponseData: &testPayload{},
			wantedErrStr: "unable to create the request: " +
				`parse ":": missing protocol scheme`,
		},
		data{
			name: "error with request sending",
			testServerHandler: func(writer http.ResponseWriter, request *http.Request) {
				panic(http.ErrAbortHandler)
			},
			args: args{
				httpClient:   &http.Client{},
				url:          "http://example.com/",
				responseData: &testPayload{},
			},
			wantedResponseData: &testPayload{},
			wantedErrStr: "unable to send the request: " +
				`Get "http://example.com/": EOF`,
		},
		data{
			name: "error with reading of the response body",
			testServerHandler: func(writer http.ResponseWriter, request *http.Request) {
				writer.Header().Set("Content-Length", "1")
			},
			args: args{
				httpClient:   &http.Client{},
				url:          "http://example.com/",
				responseData: &testPayload{},
			},
			wantedResponseData: &testPayload{},
			wantedErrStr:       "unable to read the response body: unexpected EOF",
		},
		data{
			name: "error with the response status",
			testServerHandler: func(writer http.ResponseWriter, request *http.Request) {
				writer.WriteHeader(http.StatusInternalServerError)
				writer.Write([]byte("error"))
			},
			args: args{
				httpClient:   &http.Client{},
				url:          "http://example.com/",
				responseData: &testPayload{},
			},
			wantedResponseData: &testPayload{},
			wantedErrStr:       "request failed: 500 error",
		},
		data{
			name: "error with unmarshalling of the response body",
			testServerHandler: func(writer http.ResponseWriter, request *http.Request) {
				writer.Write([]byte("incorrect"))
			},
			args: args{
				httpClient:   &http.Client{},
				url:          "http://example.com/",
				responseData: &testPayload{},
			},
			wantedResponseData: &testPayload{},
			wantedErrStr: "unable to unmarshal the response body: " +
				"invalid character 'i' looking for beginning of value",
		},
	}
	for _, testData := range tests {
		testServer := httptest.NewServer(testData.testServerHandler)
		defer testServer.Close()

		url := strings.ReplaceAll(
			testData.args.url,
			"http://example.com",
			testServer.URL,
		)
		err := loadJSONData(testData.args.httpClient, url, testData.args.responseData)

		if !reflect.DeepEqual(
			testData.args.responseData,
			testData.wantedResponseData,
		) {
			test.Logf(
				"failed %q:\n  expected: %+v\n  actual: %+v",
				testData.name,
				testData.wantedResponseData,
				testData.args.responseData,
			)
			test.Fail()
		}

		wantedErrStr := strings.ReplaceAll(
			testData.wantedErrStr,
			"http://example.com",
			testServer.URL,
		)
		wantedErr := wantedErrStr != ""
		if !wantedErr && err != nil ||
			wantedErr && (err == nil || err.Error() != wantedErrStr) {
			test.Logf(
				"failed %q:\n  expected: %+v\n  actual: %+v",
				testData.name,
				testData.wantedErrStr,
				err,
			)
			test.Fail()
		}
	}
}
