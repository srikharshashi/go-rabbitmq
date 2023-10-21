package main

import (
	"fmt"

	"github.com/streadway/amqp"
)

func main()  {
	fmt.Println("Consumer Application")
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

	msgs,err:= ch.Consume(
		"TestQueue",
		"",
		true,
		false,
		false,
		false,nil,
	)
	// this channel was just used to block the main routine as the other function runs indefinetly
	forever := make(chan bool)
	go func(){
		for d := range msgs{
			fmt.Printf("Recieved Message :%s in Queue\n",d.Body)
		}
	}()

	fmt.Println("Sucessfully connected to RabbitMQ instance")
	fmt.Println("[*] - waiting for messages")
	<- forever
}