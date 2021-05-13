package main

import "context"

func main() {

	ctx := context.Background()
	conf := Setup()

	token, err := GetToken(ctx, conf)
	if err != nil {
		panic(err)
	}

	client := conf.Client(ctx, token)
	if err = GetUsers(client); err != nil {
		panic(err)
	}

}
