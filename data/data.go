package data

type Patch struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}
type Film struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Year  int    `json:"year"`
	Votes int    `json:"votes"`
}
