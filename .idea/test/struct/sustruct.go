package main

import "fmt"

type House struct {
	Address string
	Size    int
	Price   float64
	Type    string
}

type User struct {
	Name string
	ID   string
	Age  int
}

type VIPUser struct {
	UserINFO User
	VIPLevel int
	Price    int
}

func main() {
	//var house House;
	//house.Size = 28;
	//house.Price = 9.8;
	//house.Type = "아파트";
	//house.Address = "서울시 강동구";
	//
	//fmt.Println(house.Size);

	//inner struct
	user := User{"asdf", "siosio", 15}
	fmt.Println(user)
	vipUser := VIPUser{user, 3, 1000}
	fmt.Println(vipUser.UserINFO)
	fmt.Println(vipUser)

}
