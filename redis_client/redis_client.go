package redis_client

import (
	"fmt"
	"github.com/go-redis/redis/v8"
	"strings"
	"time"
)

// RedisPool type
type RedisPool struct {
	client redis.UniversalClient
	env    string

	Conf RedisConfig
}

// NewRedisUniversalClient func;
func NewRedisUniversalClient(conf *RedisConfig) *RedisPool {
	if conf.Host != "" {
		if conf.DialTimeout < 10 {
			conf.DialTimeout = 10
		}

		if conf.ReadTimeout < 10 {
			conf.ReadTimeout = 10
		}

		if conf.WriteTimeout < 10 {
			conf.WriteTimeout = 10
		}

		if conf.IdleTimeout < 60 {
			conf.IdleTimeout = 60
		}

		myConfig := &redis.Options{
			Addr:     conf.Host,
			DB:       conf.DBNum,
			Username: conf.Username,
			Password: conf.Password,

			MaxRetries: conf.MaxRetries,

			DialTimeout:  time.Duration(conf.DialTimeout) * time.Second,
			ReadTimeout:  time.Duration(conf.ReadTimeout) * time.Second,
			WriteTimeout: time.Duration(conf.WriteTimeout) * time.Second,
			IdleTimeout:  time.Duration(conf.IdleTimeout) * time.Second,

			MinIdleConns: conf.Idle,
			PoolSize:     conf.Active,
		}

		myClient := redis.NewClient(myConfig)

		myEnv := fmt.Sprintf("[%s]tcp@%s", conf.Name, strings.Join(conf.Hosts, ";"))

		pool := &RedisPool{env: myEnv, client: myClient}
		return pool
	}

	if len(conf.Hosts) > 0 {
		if len(conf.MasterName) > 0 {
			myConfig := &redis.FailoverOptions{
				SentinelAddrs:    conf.Hosts,
				DB:               conf.DBNum,
				Username:         conf.Username,
				Password:         conf.Password,
				SentinelPassword: conf.Password,

				MaxRetries: conf.MaxRetries,

				DialTimeout:  time.Duration(conf.DialTimeout) * time.Second,
				ReadTimeout:  time.Duration(conf.ReadTimeout) * time.Second,
				WriteTimeout: time.Duration(conf.WriteTimeout) * time.Second,
				IdleTimeout:  time.Duration(conf.IdleTimeout) * time.Second,

				MinIdleConns: conf.Idle,
				PoolSize:     conf.Active,

				// Only cluster clients.
				RouteByLatency: true,

				// The sentinel master name.
				// Only failover clients.
				MasterName: conf.MasterName,
			}

			myClient := redis.NewFailoverClient(myConfig)

			myEnv := fmt.Sprintf("[%s]tcp@%s", conf.Name, strings.Join(conf.Hosts, ";"))

			pool := &RedisPool{env: myEnv, client: myClient}
			return pool
		}

		myConfig := &redis.ClusterOptions{
			Addrs:      conf.Hosts,
			Username:   conf.Username,
			Password:   conf.Password,
			MaxRetries: conf.MaxRetries,

			DialTimeout:  time.Duration(conf.DialTimeout) * time.Second,
			ReadTimeout:  time.Duration(conf.ReadTimeout) * time.Second,
			WriteTimeout: time.Duration(conf.WriteTimeout) * time.Second,
			IdleTimeout:  time.Duration(conf.IdleTimeout) * time.Second,

			MinIdleConns: conf.Idle,
			PoolSize:     conf.Active,

			// Only cluster clients.
			RouteByLatency: true,
		}

		myClient := redis.NewClusterClient(myConfig)

		myEnv := fmt.Sprintf("[%s]tcp@%s", conf.Name, strings.Join(conf.Hosts, ";"))

		pool := &RedisPool{env: myEnv, client: myClient}
		return pool
	}

	return nil
}

// Get func
func (p *RedisPool) Get() redis.UniversalClient {
	return p.client
}

// Close func
func (p *RedisPool) Close() error {
	return p.client.Close()
}
