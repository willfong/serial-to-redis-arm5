package main

import (
	"bufio"
	"log"
	"strings"

	"github.com/gomodule/redigo/redis"
	"github.com/tarm/serial"
)

func main() {
	conn, err := redis.Dial("tcp", "redis:6379")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	c := &serial.Config{Name: "/dev/ttyACM0", Baud: 9600}
	s, err := serial.OpenPort(c)
	if err != nil {
		log.Fatal(err)
	}

	for {

		// When attaching to the Arduino's serial port, sometimes we will connect
		// in the middle of the output. This gives us corrupted data, so we need
		// to ensure we only read complete data
		// Arduino's println uses "\r\n" as a line delimiter
		// So we read until "\r" and then ensure the next line starts with "\n"

		reader := bufio.NewReader(s)
		reply, err := reader.ReadBytes('\x0d')
		if err != nil {
			panic(err)
		}
		r := string(reply)

		if !(strings.HasPrefix(r, "\x0a")) {
			continue
		}

		s := strings.TrimSpace(r)
		kv := strings.Split(s, ":")

		//fmt.Println(kv[0])
		//fmt.Println(kv[1])
		//fmt.Println(s)

		_, err = conn.Do("SET", kv[0], kv[1], "EX", 3600)
		if err != nil {
			log.Fatal(err)
		}

	}

}
