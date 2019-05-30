package main

import (
	"context"
	"log"
	"math/rand"
	"time"

	"github.com/quipo/statsd"
)

func main() {
	prefix := "myproject."
	statsdclient := statsd.NewStatsdClient("127.0.0.1:8125", prefix)
	if err := statsdclient.CreateSocket(); err != nil {
		log.Fatal(err)
	}

	// Aggregate stats and flush every 2 seconds.
	interval := 2 * time.Second
	stats := statsd.NewStatsdBuffer(interval, statsdclient)
	defer stats.Close()

	// Not buffered, sent immediately.
	statsdclient.Incr("mymetric", 4)

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Minute)
	defer cancel()
	for {

		select {
		case <-ctx.Done():
			return
		default:
			n := int64(rand.Intn(10))
			// Gauge.
			stats.Gauge("mygauge", rand.Int63n(100))

			// Counter.
			stats.Incr("mymetric", n)

			// Total.
			stats.Total("mytotal", n)

			// Timer.
			stats.Timing("time", rand.Int63n(1000000))

			time.Sleep(time.Duration(rand.Intn(500)+500) * time.Millisecond)
		}
	}
}
