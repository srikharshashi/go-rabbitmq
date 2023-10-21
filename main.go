package main

import (
	"fmt"
	"github.com/streadway/amqp"
)
func main()  {
	fmt.Println("GO RabbitMQ Tutorial")
	conn,err:=amqp.Dial("amqp://guest:guest@localhost:5672/")
	if(err!=nil){
		fmt.Println("Error opening Connection")
		panic(err)
	}

	defer conn.Close()
	fmt.Println("SuccessFully Connected to RabbitMQ Instance")

	ch,err:= conn.Channel()
	if(err!=nil){
		fmt.Println("Error opening Channel")
		panic(err)
	}

	defer ch.Close()

	q, err:= ch.QueueDeclare("TestQueue",false,false,false,false,nil)

	if(err!=nil){
		fmt.Println("Error declaring Queue")
		panic(err)
	}

	fmt.Println(q)

	err=ch.Publish(
		"","TestQueue",false,false,amqp.Publishing{
			ContentType: "text/plain",
			Body: []byte("Hello World"),
		},
	)

	if(err!=nil){
		fmt.Println(err)
		panic(err)
	}

	fmt.Println("Sucessfully Published Message to Queue")
	



}