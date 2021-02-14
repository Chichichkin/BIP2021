package responses

type Error struct {
	Error   string  `json:"error"`
	Code    int     `json:"code"`
	Message string  `json:"message"`
	Details details `json:"details"`
}

type details struct {
	TypeUrl string `json:"type_url"`
	Value   string `json:"value"`
}
