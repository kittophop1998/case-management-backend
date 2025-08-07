package appcore_config

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

var Config *Configurations

type Configurations struct {
	// App
	AppName    string
	AppVersion string
	Mode       string
	SecretKey  string

	// Server
	GinMode    string
	ServerPort int
	LogLevel   string

	// Database
	PostgresHost     string
	PostgresPort     int
	PostgresUser     string
	PostgresPassword string
	PostgresDBName   string
	PostgresSSLMode  string
	PostgresTimezone string

	// Redis
	RedisHost     string
	RedisPort     int
	RedisPassword string

	// MinIO
	MinioURL        string
	MinioSSL        bool
	MinioAccessKey  string
	MinioSecretKey  string
	MinioBucketName string

	// Memphis
	MemphisHost     string
	MemphisUsername string
	MemphisPassword string

	// Email SMTP
	SMTPUser     string
	SMTPPassword string
	SMTPHost     string
	SMTPPort     string

	// System Integration
	SystemIWebsiteURL      string
	TreasureDataWebsiteURL string
	LdapURL                string

	// Optional (future extensions)
	RunCronScheduler     bool
	IsObserve            bool
	ProdPostgresURL      string
	PostgresRailwayURL   string
	RedisRailwayURL      string
	RedisRailwayPassword string
}

func InitConfigurations() {
	// Load .env if exists
	_ = godotenv.Load()

	// Set up Viper for config.yaml
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("appcore/appcore_config")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file: %v", err)
	}

	viper.AutomaticEnv() // also read from environment variables if set

	Config = &Configurations{
		AppName:    viper.GetString("app.name"),
		AppVersion: viper.GetString("app.version"),
		Mode:       viper.GetString("app.mode"),
		SecretKey:  viper.GetString("app.secret_key"),

		GinMode:    viper.GetString("server.gin_mode"),
		ServerPort: viper.GetInt("server.port"),
		LogLevel:   viper.GetString("server.log_level"),

		PostgresHost:     viper.GetString("database.host"),
		PostgresPort:     viper.GetInt("database.port"),
		PostgresUser:     viper.GetString("database.user"),
		PostgresPassword: viper.GetString("database.password"),
		PostgresDBName:   viper.GetString("database.dbname"),
		PostgresSSLMode:  viper.GetString("database.sslmode"),
		PostgresTimezone: viper.GetString("database.timezone"),

		RedisHost:     viper.GetString("redis.host"),
		RedisPort:     viper.GetInt("redis.port"),
		RedisPassword: viper.GetString("redis.password"),

		MinioURL:        viper.GetString("minio.url"),
		MinioSSL:        viper.GetBool("minio.ssl"),
		MinioAccessKey:  viper.GetString("minio.access_key"),
		MinioSecretKey:  viper.GetString("minio.secret_key"),
		MinioBucketName: viper.GetString("minio.bucket_name"),

		MemphisHost:     viper.GetString("memphis.host"),
		MemphisUsername: viper.GetString("memphis.username"),
		MemphisPassword: viper.GetString("memphis.password"),

		SMTPHost:     viper.GetString("smtp.host"),
		SMTPPort:     viper.GetString("smtp.port"),
		SMTPUser:     viper.GetString("smtp.user"),
		SMTPPassword: viper.GetString("smtp.password"),

		SystemIWebsiteURL:      viper.GetString("systems.system_i_url"),
		TreasureDataWebsiteURL: viper.GetString("systems.treasure_data_url"),
		LdapURL:                viper.GetString("systems.ldap_url"),

		PostgresRailwayURL:   viper.GetString("POSTGRES_RAILWAY_URL"),
		RedisRailwayURL:      viper.GetString("REDIS_RAILWAY_URL"),
		RedisRailwayPassword: viper.GetString("REDIS_RAILWAY_PASSWORD"),
	}
}
