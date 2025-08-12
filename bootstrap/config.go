package bootstrap

import (
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
)

type Config struct {
	AppPort string

	DBHost, DBPort, DBUser, DBPass, DBName, DBSSL, DBTZ string

	JWTSecret       string
	AccessTokenTTL  time.Duration
	RefreshTokenTTL time.Duration
}

func LoadConfig() Config {
	_ = godotenv.Load(".env")

	accessTTL, _ := time.ParseDuration(get("ACCESS_TOKEN_TTL", "15m"))
	refreshTTL, _ := time.ParseDuration(get("REFRESH_TOKEN_TTL", "720h")) // 30 ngày mặc định

	cfg := Config{
		AppPort: get("APP_PORT", "8080"),

		DBHost: get("DB_HOST", "localhost"),
		DBPort: get("DB_PORT", "5432"),
		DBUser: get("DB_USER", "postgres"),
		DBPass: get("DB_PASSWORD", "postgres"),
		DBName: get("DB_NAME", "myapp"),
		DBSSL:  get("DB_SSLMODE", "disable"),
		DBTZ:   get("DB_TIMEZONE", "UTC"),

		JWTSecret:       get("JWT_SECRET", "dev_secret"),
		AccessTokenTTL:  accessTTL,
		RefreshTokenTTL: refreshTTL, // <-- thêm dòng này
	}

	log.Printf("config loaded: port=%s db=%s accessTTL=%s refreshTTL=%s",
		cfg.AppPort, cfg.DBName, cfg.AccessTokenTTL, cfg.RefreshTokenTTL)

	return cfg
}

// DSN: chuỗi kết nối dạng "key=value"
func (c Config) DSN() string {
	return "host=" + c.DBHost +
		" user=" + c.DBUser +
		" password=" + c.DBPass +
		" dbname=" + c.DBName +
		" port=" + c.DBPort +
		" sslmode=" + c.DBSSL +
		" TimeZone=" + c.DBTZ
}

// get: lấy env, nếu trống trả về default
func get(k, def string) string {
	if v := os.Getenv(k); v != "" {
		return v
	}
	return def
}
