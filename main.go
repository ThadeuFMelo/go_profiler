package main

import (
	"context"
	"encoding/json"
	"fmt"
	"go-profiler/database/clickhouse"
	"go-profiler/gopsutil"
	"go-profiler/models"
	prometheusutil "go-profiler/prometheusutils"
	"io"
	"log"
	"net/http"
	"sort"
	"time"
)

type Vehicle struct {
	Type          string
	Wheels        int
	NumberOfDoors int `json:"number_of_doors"` // JSON tag is required if the JSON string is different than the Field name
}

type My_first_table struct {
	user_id   uint64
	message   string
	timestamp int64
	metric    float64
}

const prometheusEndpoint string = "localhost:2112"

func main() {

	prometheusutil.Register(prometheusEndpoint)
	fmt.Println("Hello, World!")
	url := "https://gorest.co.in/public/v2/users"
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	response := `[{"type": "Car","wheels": 4, "number_of_doors": 4},{"type": "Motorcycle","wheels": 2, "number_of_doors": 0}]`
	var vehicles []Vehicle

	// And we unmarshal our JSON, assigning the result to vehicles
	json.Unmarshal([]byte(response), &vehicles)
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	// body to array or users
	var users []models.User

	// var my_first_tables []Span

	err = json.Unmarshal(body, &users)
	if err != nil {
		log.Fatal(err)
	}

	for _, user := range users {
		fmt.Println(user.Name)
		fmt.Println(user.Gender)
	}

	db := clickhouse.Connect()

	//Kafka consumer
	// kafka := repository.NewKafka()
	// kafka.Consume("quickstart-events", true)
	// for 10 seconds, every 0,1 seconds get the process info and send to kafka

	ctx := context.Background()

	res, err := db.ExecContext(ctx, "SELECT * FROM helloworld.my_first_table")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("------------------------------------------------")
	fmt.Println(res)
	fmt.Println("------------------------------------------------")
	//get each line from sql.result in res

	// my_first_table := &My_first_table{
	// 	user_id:   1,
	// 	message:   "hello",
	// 	timestamp: time.Now().Unix(),
	// 	metric:    0.5,
	// }
	resI, err := db.NewCreateTable().Model((*models.ProcessMessage)(nil)).Exec(ctx)

	fmt.Println("************************************************")
	fmt.Println(resI)
	fmt.Println("------------------------------------------------")

	for i := 0; i < 2000000; i++ {
		//get process info
		process, _ := gopsutil.GetProcessesInfo()
		timestamp := time.Now()
		for _, p := range process {

			//add timestamp to process message struct
			processMessage := &models.ProcessMessage{
				Pid:       p.ProcessId,
				Cpu:       p.CPUUsage,
				Mem:       p.Memory,
				Name:      p.Name,
				TimeStamp: timestamp,
				Ctime:     p.CreateTime,
			}
			//convert process message struct to json
			// message, _ := json.Marshal(processMessage)
			// //send to kafka
			// repository.Produce("process-events", string(message))
			resP, err := db.NewInsert().Model(processMessage).Exec(ctx)
			if err != nil {
				fmt.Println("Error inserting process message")
				fmt.Println(err)
			}
			prometheusutil.ProcessCPUUsage.WithLabelValues(p.Name).Set(p.CPUUsage)
			prometheusutil.ProcessMemoryUsage.WithLabelValues(p.Name).Set(float64(p.Memory))

			fmt.Println("------------------------------------------------")
			fmt.Println(resP)
			fmt.Println("------------------------------------------------")
		}
		fmt.Printf("%d\n", i)
		time.Sleep(100 * time.Millisecond)
	}
}

func sortProcessByCPU(process []models.Process) []models.Process {
	sort.Slice(process, func(i, j int) bool {
		return process[i].CPUUsage > process[j].CPUUsage
	})
	return process
}
