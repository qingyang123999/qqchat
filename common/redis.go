package common

import (
	"context"
	"fmt"
	"qqchat/common/qqlog"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/spf13/viper"
)

// 全局上下文和Redis客户端变量
var RedisDb *redis.Client // 全局Redis客户端实例

// InitRedis 初始化Redis连接池
func InitRedis(ctx context.Context) {
	options := &redis.Options{
		Addr:         viper.GetString("Redis.Addr"),                                  // Redis服务器地址
		Password:     viper.GetString("Redis.Password"),                              // 认证密码（若无则留空）
		DB:           viper.GetInt("Redis.DB"),                                       // 数据库索引（0-15）
		PoolSize:     viper.GetInt("Redis.PoolSize"),                                 // 最大连接数
		MaxIdleConns: viper.GetInt("Redis.MaxIdleConns"),                             // 最大空闲连接数
		MinIdleConns: viper.GetInt("Redis.MinIdleConns"),                             // 最小空闲连接数
		DialTimeout:  time.Duration(viper.GetInt("Redis.DialTimeout")) * time.Second, // 连接超时时间
	}

	// 创建Redis客户端实例
	RedisDb = redis.NewClient(options)

	// 测试连接有效性
	if err := pingRedisDb(ctx); err != nil {
		qqlog.Log.Errorf("Redis连接初始化失败: %v", err)
		panic(fmt.Errorf("Redis连接初始化失败: %v", err))
	}
	fmt.Println("Redis连接初始化成功")

	//err := RedisDb.Set(ctx, "key111", "value111", 0).Err()
	//if err != nil {
	//	panic(err)
	//}
	//val, err := RedisDb.Get(ctx, "key111").Result()
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Println("redis============key111===", val)
}

// pingRedis 发送PING命令测试连接
func pingRedisDb(ctx context.Context) error {
	// 设置5秒超时上下文
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	_, err := RedisDb.Ping(ctx).Result()
	return err
}

// RedisDbPublish  发布消息到channel  Redis的发布/订阅（Pub/Sub）模式
// param  channel 消息的目标频道名称
// content  待发送的消息内容
func RedisDbPublish(ctx context.Context, channel string, content interface{}) error {
	var err error
	err = RedisDb.Publish(ctx, channel, content).Err()
	return err
}

// RedisDbSubscribe  订阅channel  Redis的发布/订阅（Pub/Sub）模式、
// param  channel 消息的目标频道名称
// return 订阅对象
func RedisDbSubscribe(ctx context.Context, channel string) *redis.PubSub {
	return RedisDb.Subscribe(ctx, channel)
}

// RedisDbReceiveMessage  从channel接收消息  Redis的发布/订阅（Pub/Sub）模式
// param  订阅对象
func RedisDbReceiveMessage(ctx context.Context, pubSub *redis.PubSub) (message string, err error) {
	receiveMessage, err := pubSub.ReceiveMessage(ctx)
	if err != nil {
		return "", err
	}
	fmt.Println(receiveMessage.String())
	return receiveMessage.Payload, nil
}
