package handlers

import (
	"encoding/json"
	"fmt"
	"mongodb"
	"net/http"
	"redisdb"
	"time"

	"github.com/go-redis/redis/v8"
)

func StudentsRedis(w http.ResponseWriter, r *http.Request) {

	res, err := redisdb.Get("Students")

	if err == redis.Nil {

		fmt.Println("Not in redis")
		fmt.Println("Fetch data from mongo")

		students, err := mongodb.GetStudents()

		if err != nil {
			fmt.Fprintf(w, "%s", err.Error())
		}

		res, _ := json.Marshal(students)

		redisdb.Set("Students", res, time.Second*5)

		fmt.Fprintf(w, "%s", res)
		return
	}

	if err != nil {
		fmt.Fprintf(w, "%s", err.Error())
		return
	}

	fmt.Println("Data from redis")
	fmt.Fprintf(w, "%s", *res)
	return

}

func StudentsNotRedis(w http.ResponseWriter, r *http.Request) {

	students, err := mongodb.GetStudents()

	if err != nil {
		fmt.Println("Err :", err)
		fmt.Fprintf(w, "%s", err.Error())
	}

	res, _ := json.Marshal(students)

	fmt.Fprintf(w, "%s", res)
	return
}

func ErrorRoute() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		var response string = "This API provide GET method for /redis & /notredis."

		fmt.Fprintf(w, "%s", response)
	}
}
