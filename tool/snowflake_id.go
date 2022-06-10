package tool

import (
	"github.com/bwmarrin/snowflake"
	"github.com/spf13/viper"
	"time"
)

var node *snowflake.Node

// InitSnowflake 使用雪花算法生成分布式ID
func InitSnowflake() error {
	var st time.Time
	st, err := time.Parse("2006-01-02", GetViper().GetString("app.start_time"))
	if err != nil {
		return err
	}
	//转换为毫秒
	snowflake.Epoch = st.UnixNano() / 1000000
	n, err := snowflake.NewNode(viper.GetInt64("machine_id"))
	if err != nil {
		return err
	}
	node = n
	return nil
}

//GetId 生成分布式id
func GetId() int64 {
	return node.Generate().Int64()
}
