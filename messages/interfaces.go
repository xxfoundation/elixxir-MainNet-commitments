package messages

// Commitment struct defines the JSON structure used by the REST api
type Commitment struct {
	IDF           string `json:"idf"`
	Contract      string `json:"contract"`
	PaymentWallet string `json:"validator-wallet"`
	Signature     string `json:"signature"`
}
