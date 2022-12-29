CREATE TABLE helloworld.my_first_table
(
    user_id UInt32,
    message String,
    timestamp DateTime,
    metric Float32
)
ENGINE = MergeTree()
PRIMARY KEY (user_id, timestamp);

-- Path: repository/clickhouse.sql

CREATE TABLE helloworld.process_messages(
    pid UInt32,
    cpu Float32,
    mem Float32,
    name String,
    time_stamp DateTime,
    ctime DateTime
)
ENGINE = MergeTree()
PARTITION BY toYYYYMM(timestamp)
ORDER BY (pid, timestamp)
SAMPLE BY pid;


select 
cpu,
any(cpu) OVER(PARTITION by name order by time_stamp rows between 1 preceding and 0 preceding ) as pre_cpu,
cpu - any(cpu) 
OVER(PARTITION by name order by time_stamp rows between 1 preceding and 0 preceding ) as deleta_cpu,
time_stamp - any(time_stamp) OVER(PARTITION by name order by time_stamp rows between 1 preceding and 0 preceding ) as delta_time_stamp
 from helloworld.process_messages prewhere name = 'main'  order by time_stamp desc

 