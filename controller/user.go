package controller

import  "command-line-argumentsC:\\Users\\Erimbetov\\go\\homework\\strucks\\strucks.go"

func GenerateFunction(limit int) []strucks.User {

	var users = []strucks.User{}

	for i := 0; i < limit; i++ {

		users = append(users, []strucks.User{

			{
				FirstName: "Nodirek",
				LastName:  "Erimbetov",
				Age:       22,
			},
		})

	}
	return users
}