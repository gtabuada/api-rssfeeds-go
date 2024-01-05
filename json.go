package main

import ( 
  "net/http"
  "encoding/json"
  "log"
)

func respondWithError(w http.ResponseWriter, code int, msg string) {
  if code > 499 {
    log.Println("Internal Server Error: %v", msg)
  }

  type errResponse struct {
    Error string `json:"error"`
  }

  respondWithJson(w, code, errResponse {
    Error: msg,
  })
}

func respondWithJson(w http.ResponseWriter, code int, payload interface{}) {

  data, err := json.Marshal(payload)

  if err != nil {
    w.WriteHeader(500)
    log.Println("Failed to marshal JSON response: %v", payload)
    return
  }

  w.Header().Add("Content-Type", "application/json")
  w.WriteHeader(code)
  w.Write(data)
}
