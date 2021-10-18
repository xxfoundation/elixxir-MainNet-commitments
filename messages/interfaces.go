package messages

// Commitment struct defines the JSON structure used by the REST api
type Commitment struct {
	IDF       []byte `json:"idf"`
	Contract  []byte `json:"contract"`
	Wallet    string `json:"wallet"`
	Signature []byte `json:"signature"`
}
