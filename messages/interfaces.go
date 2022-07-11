package messages

// Commitment struct defines the JSON structure used by the REST api
type Commitment struct {
	IDF                string  `json:"idf"`
	Contract           string  `json:"contract"`
	ValidatorWallet    string  `json:"validator-wallet"`
	NominatorWallet    string  `json:"nominator-wallet"`
	SelectedMultiplier float32 `json:"selected-multiplier"`
	MaxMultiplier      float32 `json:"max-multiplier"`
	Email              string  `json:"email"`
	Signature          string  `json:"signature"`
}

type CommitmentInfo struct {
	ValidatorWallet    string  `json:"validator-wallet"`
	NominatorWallet    string  `json:"nominator-wallet"`
	SelectedMultiplier float32 `json:"selected-multiplier"`
	MaxMultiplier      float32 `json:"max-multiplier"`
}
