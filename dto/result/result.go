package resultdto

type SuccesResult struct {
	Code int `json:"Code"`
	Data interface{} `json:"Data"`

}

type ErrorResult struct {
	Code int `json:"Code"`
	Message string `json:"Message"`
}
