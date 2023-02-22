package limiter

import (
	"os"
	"log"
	"time"
	"sync"
	"strconv"
	"github.com/redis/go-redis/v9"

	"github.com/transactional_outbox_pattern/main_service/database"
)

type RateLimiter struct {
	dbConnection *database.RedisDB
}

var rateLimiter *RateLimiter
var lock = &sync.Mutex{}

func NewRateLimiter() *RateLimiter {

	if rateLimiter == nil {
		lock.Lock()
		defer lock.Unlock()

		rateLimiter = &RateLimiter{
			database.NewRedisDBConnection(getRedisDBInfo()),
		}

	}

	return rateLimiter
}

func (r *RateLimiter) Run() {

	var ticker = time.NewTicker(6000 * time.Millisecond)

	for {
		select {
		case <-ticker.C:
			r.fillBucket()
		}
	}
}

func (r *RateLimiter) fillBucket() {
	r.dbConnection.SetKey("tokens", 3)
}

func (r *RateLimiter) IsBucketEmpty() bool {
	val, err := strconv.Atoi(r.dbConnection.GetVal("tokens"))

	if err != nil {
		log.Fatal("Error checking if bucket is empty")
	}

	return  val == 0 
}


func getRedisDBInfo() *redis.Options {
	return &redis.Options {
		Addr:	  os.Getenv("REDIS_ADDR"),
		Password: "", // no password set
		DB:	0,
	}
}