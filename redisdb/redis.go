package redisdb

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/go-redis/redis/v8"
)

//var URL string = "redis://default:password@localhost:6379/0"

var URL string = os.Getenv("REDIS_URL")
var ctx = context.Background()

func getConnexion() (*redis.Client, error) {

	opt, err := redis.ParseURL(URL)
	if err != nil {
		return nil, fmt.Errorf("Problème de connexion : %v", err)
	}

	con := redis.NewClient(opt)

	if con == nil {
		return nil, fmt.Errorf("Impossible de se connecter")
	}

	return con, nil
}

func Set(key string, value any, timer time.Duration) error {
	conn, err := getConnexion()
	if err != nil {
		return err
	}
	defer conn.Close()

	err = conn.Set(ctx, key, value, timer).Err()

	if err != nil {
		return err
	}
	return nil
}

func Get(key string) (*string, error) {
	conn, err := getConnexion()

	if err != nil {
		return nil, err
	}
	defer conn.Close()

	res, err := conn.Get(ctx, key).Result()

	if err != nil {
		return nil, err
	}
	return &res, nil
}
