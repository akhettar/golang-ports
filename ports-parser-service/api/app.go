package api

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"

	pb "ports-parser-service/grpc"
	m "ports-parser-service/model"
	"ports-parser-service/services"
)

// PortParserService type
type PortParserService struct {
	router      *mux.Router
	PortService services.PortService
}

// NewPortParserService create an instance of teh PortParserService
func NewPortParserService(addr string) *PortParserService {
	ps := &PortParserService{mux.NewRouter(), services.NewPortService(addr)}
	ps.routes()
	return ps
}

func (srv *PortParserService) routes() {
	srv.router.HandleFunc("/ports", srv.AddPorts).Methods("POST")
	srv.router.HandleFunc("/ports/{code}", srv.GetPort).Methods("GET")
}

// Run the application
func (ps *PortParserService) Run(addr string) {
	srv := &http.Server{
		Handler: ps.router,
		Addr:    addr,
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Fatal(srv.ListenAndServe())
}

// AddPorts send the add port request to Ports Writer Service
func (srv *PortParserService) AddPorts(w http.ResponseWriter, req *http.Request) {
	log.Println("received request to process ports")
	var pr m.ProcessPortsRequest
	if err := json.NewDecoder(req.Body).Decode(&pr); err != nil {
		returnError(w, http.StatusBadRequest, "Inavlid payload")
		return
	}

	log.Printf("Processing file: %s with size: %d", pr.File, pr.Size)
	f, err := os.Open(pr.File)

	if err != nil {
		returnError(w, http.StatusBadRequest, "failed to download the file")
		return
	}

	dec := json.NewDecoder(f)

	// read open bracket
	_, err = dec.Token()
	if err != nil {
		log.Fatal(err)
	}
	counter := 0
	for {
		t, err := dec.Token()
		_, ok := t.(json.Delim)

		// reached the end of the file
		if err != nil || ok {
			break
		}

		// extract the port details

		p := m.Port{PortCode: fmt.Sprintf("%v", t)}
		dec.More()

		if err := dec.Decode(&p); err != nil {
			log.Printf("failed to deserialise Port with code %s", p.PortCode)
			continue
		}

		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()

		// send add port request to the port manager
		log.Printf("sendnig port with code %s to port manager\n", p.PortCode)
		r, err := srv.PortService.Add(ctx, m.ToPbRequest(p))

		if err != nil {
			log.Printf("could not add port: %v", err)
			returnError(w, http.StatusInternalServerError, "Failed to create Port")
			return
		}
		log.Printf("Port %v added: %s", p, r.GetMessage())
		counter++
	}
	w.WriteHeader(http.StatusCreated)
	w.Header().Add("Content-Type", "applicaiton/json")
	res, _ := json.Marshal(m.ProcessPortsResponse{File: pr.File, NumberOfRecords: counter})
	w.Write(res)
}

// GetPort fetches the port details for given code
func (srv *PortParserService) GetPort(w http.ResponseWriter, req *http.Request) {

	code := mux.Vars(req)["code"]
	log.Printf("received request to fetch port by code: %s", code)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// fire add port request
	r, err := srv.PortService.Get(ctx, &pb.PortCode{Code: code})

	if err != nil {
		returnError(w, http.StatusNotFound, fmt.Sprintf("Port with code %s not found", code))
		return
	}
	log.Printf("Port found for code")

	w.WriteHeader(http.StatusOK)
	w.Header().Add("Content-Type", "applicaiton/json")
	res, err := json.Marshal(m.ToPortResponse(r))

	if err != nil {
		returnError(w, http.StatusInternalServerError, "failed to deserialise the resposne")
		return
	}
	w.Write(res)

}

func returnError(w http.ResponseWriter, code int, msg string) {
	body, _ := json.Marshal(map[string]string{"error": msg})
	w.WriteHeader(code)
	w.Header().Add("Content-Type", "appplication/json")
	w.Write(body)
}
