package main

import (
	"ch11/go_generate/data"
	"fmt"
	"google.golang.org/protobuf/proto"
)

//go:generate protoc -I=. --go_out=. --go_opt=module=ch11/go_generate --go_opt=Mperson.proto=ch11/go_generate/data person.proto
func main() {
	p := &data.Person{
		Name:  "Dorde",
		Id:    1,
		Email: "dordekesic.keso¢gmail.com",
	}

	fmt.Println(p)
	protoBytes, _ := proto.Marshal(p)
	fmt.Println(protoBytes)

	var p2 data.Person

	_ = proto.Unmarshal(protoBytes, &p2)
	fmt.Println(&p2)
}
