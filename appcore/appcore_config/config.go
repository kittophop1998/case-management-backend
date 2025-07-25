package appcore_config

import (
	"github.com/spf13/viper"
)

var Config *Configurations

// Configurations wraps all the config variables required by the service
type Configurations struct {
	// Develop or Production
	Mode string

	// Gin Mode
	GinIsReleaseMode bool

	// Database
	PostgresConnString string

	// Redis
	RedisUrl  string
	RedisPass string

	// Message Broker
	MemphisHost     string
	MemphisUsername string
	MemphisPassword string

	//Storage
	MinioURL        string
	MinioSSL        bool
	MinioAccessKey  string
	MinioSecretKey  string
	MinioBucketName string

	// JWT
	SecretKey string

	// CRON Scheduler
	RunCronScheduler bool

	// System-i
	SystemIWebsiteURL string

	// Treasure Data
	TreasureDataWebsiteURL string

	//Otel
	IsObserve bool
}

// NewConfigurations returns a new Configuration object
func InitConfigurations() {
	viper.AutomaticEnv()
	viper.SetDefault("mode", "")
	viper.SetDefault("GIN_IS_RELEASE_MODE", false)
	viper.SetDefault("POSTGRES_URL", "host=127.0.0.1 user=postgres password=postgres dbname=postgres port=5432 sslmode=disable TimeZone=Asia/Bangkok")
	viper.SetDefault("REDIS_URL", "localhost:6379")
	viper.SetDefault("REDIS_PASS", "password123")
	viper.SetDefault("MEMPHIS_HOST", "")
	viper.SetDefault("MEMPHIS_USERNAME", "memphis")
	viper.SetDefault("MEMPHIS_PASSWORD", "memphis")
	viper.SetDefault("MINIO_URL", "localhost:9010")
	viper.SetDefault("MINIO_SSL", false)
	viper.SetDefault("MINIO_ACCESS_KEY", "minioadmin")
	viper.SetDefault("MINIO_SECRET_KEY", "minioadmin")
	viper.SetDefault("MINIO_BUCKET_NAME", "miniobucket")
	viper.SetDefault("SECRET_KEY", "")
	viper.SetDefault("SYSTEM_I_URL", "")
	viper.SetDefault("TREASURE_DATA_URL", "")

	Config = &Configurations{
		Mode:                   viper.GetString("mode"),
		GinIsReleaseMode:       viper.GetBool("GIN_IS_RELEASE_MODE"),
		PostgresConnString:     viper.GetString("POSTGRES_URL"),
		RedisUrl:               viper.GetString("REDIS_URL"),
		RedisPass:              viper.GetString("REDIS_PASS"),
		MemphisHost:            viper.GetString("MEMPHIS_HOST"),
		MemphisUsername:        viper.GetString("MEMPHIS_USERNAME"),
		MemphisPassword:        viper.GetString("MEMPHIS_PASSWORD"),
		MinioURL:               viper.GetString("MINIO_URL"),
		MinioSSL:               viper.GetBool("MINIO_SSL"),
		MinioAccessKey:         viper.GetString("MINIO_ACCESS_KEY"),
		MinioSecretKey:         viper.GetString("MINIO_SECRET_KEY"),
		MinioBucketName:        viper.GetString("MINIO_BUCKET_NAME"),
		SecretKey:              viper.GetString("SECRET_KEY"),
		RunCronScheduler:       viper.GetBool("RUN_CRON_SCHEDULER"),
		IsObserve:              viper.GetBool("IS_OBSERVE"),
		SystemIWebsiteURL:      viper.GetString("SYSTEM_I_URL"),
		TreasureDataWebsiteURL: viper.GetString("TREASURE_DATA_URL"),
	}
}
