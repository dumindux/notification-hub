package utils

import (
	"log"
	"encoding/json"
	"fmt"
	"notification-hub/monitor"
	"time"
	"github.com/influxdata/influxdb/client/v2"
)

func DataExporter(config *InfluxDB, name string, dataChannel chan []byte) {
	dbClient := createConnection(config);
	for {
		data := <-dataChannel

		var status monitor.Status
		json.Unmarshal(data, &status)

		log.Println("receved data")
		fmt.Println(status)

		batchPoints, err := client.NewBatchPoints(client.BatchPointsConfig{
			Database:  config.DBName,
			Precision: "ns",
		})
		if err != nil {
			log.Fatal(err)
		}
		pt, err := client.NewPoint(
			"health_check",
			status.Tags,
			status.Fields,
			time.Unix(0, status.Time * int64(time.Millisecond)),
		)

		if err != nil {
			log.Fatal(err)
		}

		batchPoints.AddPoint(pt)

		if err := (*dbClient).Write(batchPoints); err != nil {
			log.Fatal(err)
			dbClient = createConnection(config);
		}
	}
}

func createConnection(config *InfluxDB) *client.Client {
	c, err := client.NewHTTPClient(client.HTTPConfig{
		Addr: config.Address,
		Username: config.Username,
		Password: config.Password,
	})
	if err != nil {
		log.Fatal(err)
	}
	defer c.Close()

	return &c
} 
