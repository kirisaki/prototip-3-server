package device

import (
    "fmt"
    "net/http"
    "encoding/json"
    "github.com/kirisaki/prototip-3-server/types"
)

type Res struct {
    Detail string `json:"detail"`
}

func Handler(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    id := r.URL.Query().Get("id")
    var bytes []byte

    switch id {
    case "a":
        bytes, _ = json.Marshal(Res{"a dayo"})
    case "b":
        bytes, _ = json.Marshal(Res{"b dayo"})
    case "c":
        bytes, _ = json.Marshal(Res{"c dayo"})
    default:
        w.WriteHeader(404)
        bytes, _ = json.Marshal(types.ResError{"not found jai!!"})
    }

    fmt.Fprintf(w, string(bytes))
}

