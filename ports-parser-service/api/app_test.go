package api

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	mocks "ports-parser-service/mocks"
	"ports-parser-service/model"

	"github.com/golang/mock/gomock"
	"github.com/gorilla/mux"

	gc "ports-parser-service/grpc"
)

func TestPortParserService_AddPortsSuccess(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// setting mock expectations
	mock := mocks.NewMockPortService(ctrl)

	mock.EXPECT().Add(gomock.Any(), gomock.Any()).Return(&gc.PortResponse{Message: "success"}, nil).AnyTimes()

	body, _ := json.Marshal(model.ProcessPortsRequest{File: "../ports_small.json"})
	req := httptest.NewRequest("POST", "/ports", bytes.NewReader(body))

	app := &PortParserService{mux.NewRouter(), mock}
	app.routes()
	rr := httptest.NewRecorder()

	// send the request
	app.router.ServeHTTP(rr, req)

	if rr.Result().StatusCode != http.StatusCreated {
		t.Errorf("server responded with the wrong error code: %d", rr.Result().StatusCode)
	}

}

func TestPortParserService_AddPortsFailedToDownlaodTheFile(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// setting mock expectations
	mock := mocks.NewMockPortService(ctrl)

	mock.EXPECT().Add(gomock.Any(), gomock.Any()).Return(&gc.PortResponse{Message: "success"}, nil).AnyTimes()

	body, _ := json.Marshal(model.ProcessPortsRequest{File: "dummy_file.json"})
	req := httptest.NewRequest("POST", "/ports", bytes.NewReader(body))

	app := &PortParserService{mux.NewRouter(), mock}
	app.routes()
	rr := httptest.NewRecorder()

	// send the request
	app.router.ServeHTTP(rr, req)

	if rr.Result().StatusCode != http.StatusBadRequest {
		t.Errorf("server responded with the wrong error code: %d", rr.Result().StatusCode)
	}

}

func TestPortParserService_AddPortsGotErrorFromPortManagerService(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// setting mock expectations
	mock := mocks.NewMockPortService(ctrl)

	mock.EXPECT().Add(gomock.Any(), gomock.Any()).Return(nil, errors.New("failed to persist port into the database")).AnyTimes()

	body, _ := json.Marshal(model.ProcessPortsRequest{File: "dummy_file.json"})
	req := httptest.NewRequest("POST", "/ports", bytes.NewReader(body))

	app := &PortParserService{mux.NewRouter(), mock}
	app.routes()
	rr := httptest.NewRecorder()

	// send the request
	app.router.ServeHTTP(rr, req)

	if rr.Result().StatusCode != http.StatusBadRequest {
		t.Errorf("server responded with the wrong error code: %d", rr.Result().StatusCode)
	}

}
