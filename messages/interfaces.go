package messages

type Commitment struct {
	IDF       []byte `json:"idf"`
	Contract  []byte `json:"contract"`
	Wallet    string `json:"wallet"`
	Signature []byte `json:"signature"`
}
