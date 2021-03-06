// ------------------------------------------------------------
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.
// ------------------------------------------------------------

package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

// Operands Values to add
type Operands struct {
	OperandOne float32 `json:"operandOne,string"`
	OperandTwo float32 `json:"operandTwo,string"`
}

func add(w http.ResponseWriter, r *http.Request) {
	log.Infoln("########## IN add...")
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	var operands Operands
	err := json.NewDecoder(r.Body).Decode(&operands)
	if err != nil {
		os.Exit(1)
	}
	log.Infoln(fmt.Sprintf("%s%f%s%f", "########## Adding ", operands.OperandOne, " to ", operands.OperandTwo))
	err = json.NewEncoder(w).Encode(operands.OperandOne + operands.OperandTwo)
	if err != nil {
		os.Exit(1)
	}
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/add", add).Methods("POST", "OPTIONS")
	log.Fatal(http.ListenAndServe(":6000", router))
}
