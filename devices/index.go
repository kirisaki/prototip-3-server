package devices

import (
    "fmt"
    "net/http"
    "encoding/json"
)

type Res struct {
    Ids []string `json:"ids"`
}

func Handler(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    d := Res {
        []string{
            "a",
            "b",
            "c",
        },
    }
    bytes, _ := json.Marshal(d)
    fmt.Fprintf(w, string(bytes))
}

