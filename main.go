package main

import (
	"Go-grpc/server"
	"fmt"
	"google.golang.org/protobuf/proto"
)

func main(){

	user := server.User{
		Age: 18,
		Name : "shanghai",
	}

	ret,err := proto.Marshal(&user)
	if err != nil {
		panic(err)
	}

	userret := server.User{}

	err = proto.Unmarshal(ret, &userret)
	if err != nil{
		panic(err)
	}

	fmt.Println(userret.String())

}
