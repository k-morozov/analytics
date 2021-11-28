package db

import (
	"encoding/json"
	"log"
)

// CollectRequest @TBD thinks about types
type CollectRequest struct {
	AppName    string `json:"an"`
	AppVersion string `json:"av"`
	ClientId   string `json:"cid"`
	Action     string `json:"ea"` // action?
	Category   string `json:"ec"` // Category
	Label      string `json:"el"` // label?
	Value      string `json:"ev"` // value?
	// @TBD adds all
	// t:event tid: v:1 z:-38677
}

func Convert(args map[string]string) (r *CollectRequest, err error) {
	marsh, err := json.Marshal(args)
	//log.Printf("marsh = %v\n", string(marsh))
	err = json.Unmarshal(marsh, &r)
	if err != nil {
		log.Fatalf("error: %v", err)
		return
	}
	return
}
