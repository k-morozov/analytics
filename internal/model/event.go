package model

// CollectRequest @TBD thinks about types
type CollectRequest struct {
	AppName    string `json:"an"`
	AppVersion string `json:"av"`
	ClientId   string `json:"cid"`
	Action     string `json:"ea"` // action?
	Category   string `json:"ec"` // Category?
	Label      string `json:"el"` // label?
	Value      string `json:"ev"` // value?
	// @TBD adds all
	// t:event tid: v:1 z:-38677
}
