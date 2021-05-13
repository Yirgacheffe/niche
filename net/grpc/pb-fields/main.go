package main

import (
	"fmt"
	"log"

	"google.golang.org/protobuf/proto"
)

func main() {

	elliot := Person{
		Name: "elliot",
		Age:  40,
		SocialFollowers: &SocialFollowers{
			Youtube: 2409,
			Twitter: 3748,
		},
	}

	data, err := proto.Marshal(&elliot)
	if err != nil {
		log.Fatal("marshal error:", err)
	}

	fmt.Println(data)

	elliotAgain := Person{}
	err = proto.Unmarshal(data, &elliotAgain)
	if err != nil {
		log.Fatal("unmarshal error:", err)
	}

	fmt.Println(elliotAgain.GetName(), elliotAgain.GetAge())
	fmt.Println(elliotAgain.SocialFollowers.GetYoutube())
	fmt.Println(elliotAgain.SocialFollowers.GetTwitter())

}
