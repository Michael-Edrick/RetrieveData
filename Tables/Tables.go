package Tables

type User struct {
	Id        int     `json:"id"`
	Uid       string  `json:"uid"`
	Password  string  `json:"password"`
	FirstName string  `json:"first_name"`
	LastName  string  `json:"last_name"`
	UserName  string  `json:"username"`
	Addresses Address `json:"address"`
}

type Address struct {
	City          string     `json:"city"`
	StreetName    string     `json:"street_name"`
	StreetAddress string     `json:"street_address"`
	ZipCode       string     `json:"zip_code"`
	State         string     `json:"state"`
	Country       string     `json:"country"`
	Coordinates   Coordinate `json:"coordinates"`
}

type Coordinate struct {
	Lat float64 `json:"lat"`
	Lng float64 `json:"lng"`
}