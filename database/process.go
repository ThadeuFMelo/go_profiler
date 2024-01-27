package database

import (
	"log"
	"reflect"
	"strconv"
	"time"
)

type Process struct {
	Name       string  `json:"name"`
	CPUUsage   float64 `json:"cpu_usage"`
	Memory     float32 `json:"mem_usage"`
	ProcessId  uint32  `json:"pid"`
	CreateTime int64   `json:"ctime"`
	Timestamp  int64   `json:"time"`
}

// GetTableName implements ITable.
func (*Process) GetTableName() string {
	return "process"
}

func (p *Process) Encode() IScyllaTable {
	process := &ScyllaProcess{
		Name:       p.Name,
		CPUUsage:   p.CPUUsage,
		Memory:     p.Memory,
		ProcessId:  p.ProcessId,
		CreateTime: p.CreateTime,
		Timestamp:  p.Timestamp,
	}
	return process
}

type ScyllaProcess struct {
	Name       string  `json:"name"`
	CPUUsage   float64 `json:"cpu_usage"`
	Memory     float32 `json:"mem_usage"`
	ProcessId  uint32  `json:"pid"`
	CreateTime int64   `json:"ctime"`
	Timestamp  int64   `json:"time"`
}

func (ScyllaProcess) GetTableName() string {
	return "process"
}

func (p ScyllaProcess) GetColumns() []string {
	val := reflect.ValueOf(p)
	typeOf := val.Type()
	var columns []string
	for i := 0; i < typeOf.NumField(); i++ {
		columns = append(columns, typeOf.Field(i).Tag.Get("json"))
	}
	return columns
}

func (p ScyllaProcess) GetValues() map[string]string {
	val := reflect.ValueOf(p)
	typeOf := val.Type()
	values := map[string]string{}
	for i := 0; i < typeOf.NumField(); i++ {
		fieldType := typeOf.Field(i).Type
		column := typeOf.Field(i).Tag.Get("json")
		switch fieldType {
		case reflect.TypeOf(""):
			values[column] = "'" + val.Field(i).String() + "'"
		case reflect.TypeOf(uint32(0)):
			values[column] = strconv.FormatUint(uint64(val.Field(i).Uint()), 10)
		case reflect.TypeOf(float64(0)):
			values[column] = strconv.FormatFloat(val.Field(i).Float(), 'f', -1, 64)
		case reflect.TypeOf(float32(0)):
			values[column] = strconv.FormatFloat(float64(val.Field(i).Float()), 'f', -1, 32)
		case reflect.TypeOf(int64(0)):
			values[column] = strconv.FormatInt(val.Field(i).Int(), 10)
		case reflect.TypeOf(time.Now().Unix()):
			values[column] = string(rune(val.Field(i).Int()))
		default:
			log.Fatal("type not supported")
		}
	}
	return values
}

func (p *ScyllaProcess) BuildSelectQuery(fields []string) string {
	return buildSelectQuery(p, fields)
}

func (p *ScyllaProcess) BuildInsertQuery() string {
	return buildInsertQuery(p)
}

func (p *ScyllaProcess) Decode() ITable {
	process := &Process{
		Name:       p.Name,
		CPUUsage:   p.CPUUsage,
		Memory:     p.Memory,
		ProcessId:  p.ProcessId,
		CreateTime: p.CreateTime,
	}
	return process
}
