package binary

type Package struct {
	Id     byte
	Length int
	Data   []byte
}

type People struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}
