package database

import "reflect"

type Process struct {
	Name       string  `json:"name"`
	CPUUsage   float64 `json:"cpu_usage"`
	Memory     float32 `json:"memory"`
	ProcessId  uint32  `json:"process_id"`
	CreateTime int64   `json:"create_time"`
}

// GetTableName implements ITable.
func (*Process) GetTableName() string {
	return "process"
}

// GetColumns implements ITable.
func (p *Process) GetColumns() []string {
	return getColumns(p)
}

// GetValues implements ITable.
func (p *Process) GetValues() map[string]string {
	return getValues(p)
}

type ScyllaProcess struct {
	Name       string  `json:"name"`
	CPUUsage   float64 `json:"cpu_usage"`
	Memory     float32 `json:"memory"`
	ProcessId  uint32  `json:"process_id"`
	CreateTime int64   `json:"create_time"`
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

func (p *ScyllaProcess) GetValues() map[string]string {
	return getValues(p)
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
