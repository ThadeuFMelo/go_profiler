package main

import (
	"fmt"
	"go-profiler/database"
	scyllaDB "go-profiler/database/scylla"
	"go-profiler/gopsutil"
	prometheusutil "go-profiler/prometheusutils"
	"sort"
	"time"

	"go.uber.org/zap"
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

	// body to array or users
	var users []database.User

	db := scyllaDB.Connect()
	defer db.Close()

	//Kafka consumer
	// kafka := repository.NewKafka()
	// kafka.Consume("quickstart-events", true)
	// for 10 seconds, every 0,1 seconds get the process info and send to kafka

	// ctx := context.Background()
	logger := gopsutil.CreateLogger("info")

	results := scyllaDB.SelectQuery(db, &database.ScyllaUser{}, []string{"first_name", "last_name", "picture_location"}, logger)

	for _, datarow := range results {
		columns := datarow.Columns
		values := datarow.Values
		user := database.User{}
		for i := 0; i < len(columns); i++ {
			fmt.Println(columns[i], values[i])
			switch columns[i] {
			case "first_name":
				user.FirstName = values[i].(string)
			case "last_name":
				user.LastName = values[i].(string)
			case "address":
				user.Address = values[i].(string)
			case "picture_location":
				user.PictureLocation = values[i].(string)
			}
		}
		users = append(users, user)
	}
	//get each line from sql.result in res

	// my_first_table := &My_first_table{
	// 	user_id:   1,
	// 	message:   "hello",
	// 	timestamp: time.Now().Unix(),
	// 	metric:    0.5,
	// }
	//resI, err := db.NewCreateTable().Model((*models.ProcessMessage)(nil)).Exec(ctx)

	fmt.Println("************************************************")
	fmt.Println(users)
	fmt.Println("------------------------------------------------")

	for i := 0; i < 2000000; i++ {
		//get process info
		process, _ := gopsutil.GetProcessesInfo()
		for _, p := range process {

			p.Timestamp = time.Now().UnixNano()

			err := scyllaDB.InsertRow(db, &p, logger)
			if err != nil {
				logger.Error("insert row", zap.Error(err))
			}

			prometheusutil.ProcessCPUUsage.WithLabelValues(p.Name).Set(p.CPUUsage)
			prometheusutil.ProcessMemoryUsage.WithLabelValues(p.Name).Set(float64(p.Memory))

			fmt.Println("------------------------------------------------")
			fmt.Println(p.Name)
			fmt.Println(p.CPUUsage)
			fmt.Println(p.Memory)
			fmt.Println("------------------------------------------------")
		}
		fmt.Printf("%d\n", i)
		time.Sleep(100 * time.Millisecond)
	}
}

func sortProcessByCPU(process []database.Process) []database.Process {
	sort.Slice(process, func(i, j int) bool {
		return process[i].CPUUsage > process[j].CPUUsage
	})
	return process
}
