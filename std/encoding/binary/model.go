package binary

type Package struct {
	Id     byte
	Length int
	Data   []byte
}

type User struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

var user = User{
	Id:   1,
	Name: "zing",
	Age:  27,
}

var users = []User{
	user,
	{Id: 2, Name: "zrx", Age: 26},
	{Id: 3, Name: "trump", Age: 75},
}
