package api

type Earning struct {
	Id         string    `json:"id"`
	Type       string `json:"type"`
	Attributes struct {
		Amount float64 `json:"amount"`
		CreatedAt string `json:"created_at"`
	} `json:"attributes"`
	Relationships struct {
		Team *struct {
			Data Program `json:"data"`
		} `json:"team"`
		Bounty *struct {
			Data Bounty `json:"data"`
		} `json:"bounty"`
		Pentester *struct {
			Data User `json:"data"`
		} `json:"pentester"`
		ReportRetestUser *struct {
			Data User `json:"data"`
		} `json:"report_retest_user"`
	} `json:"relationships"`
}
