package query

type Queries struct {
	CaseID  int    `json:"case_id"`
	Queries string `json:"queries"`
}

type Query struct {
	CaseID int    `json:"case_id"`
	Query  string `json:"query"`
}

type QueryDTO struct {
	CaseID  int    `json:"case_id"`
	Query   string `json:"query"`
	Queries string `json:"queries"`
}
