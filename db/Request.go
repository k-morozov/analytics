package db

type CollectRequest struct {
	ApplicationName string `json:"an"`
	Version         string `json:"av"`
	ClientId        string `json:"cid"`
	Command         string `json:"ea"`
	Type            string `json:"ec"`
	Count           string `json:"el"`
	// @TBD adds all
}
