package config

import (
	"fmt"

	"github.com/ilyakaznacheev/cleanenv"
)

// IDHasher configurations
type IDHasher struct {
	MinLength int    `yaml:"min_length" env:"IDHASHER_MIN_LENGTH"`
	Salt      string `yaml:"salt" env:"IDHASHER_SALT"`
}

// MySQL config for MySQL server
type MySQL struct {
	Masters         string `yaml:"masters" env:"MYSQL_MASTERS"`
	Slaves          string `yaml:"slaves" env:"MYSQL_SLAVES"`
	User            string `yaml:"user" env:"MYSQL_USER"`
	Password        string `yaml:"password" env:"MYSQL_PASSWORD"`
	DB              string `yaml:"db" env:"MYSQL_DB"`
	MaxOpenConns    int    `yaml:"max_open_conns" env:"MYSQL_MAX_OPEN_Conns"`
	MaxIdleConns    int    `yaml:"max_idle_conns" env:"MYSQL_MAX_IDLE_Conns"`
	ConnMaxLifetime int    `yaml:"conn_max_lifetime" env:"MYSQL_CONN_MAX_LIFETIME"` // time in minute
	IsEnabledLog    bool   `yaml:"is_enabled_log" env:"MYSQL_IS_ENABLED_LOG"`
}

// Conn return connection string
func (m MySQL) Conn(host string) string {
	return fmt.Sprintf("%s:%s@tcp(%s)/%s?multiStatements=true&parseTime=true&charset=utf8mb4", m.User, m.Password, host, m.DB)
}

// Config is APP config information
type Config struct {
	Env       string   `yaml:"env" ENV:"ENV"`
	Port      string   `yaml:"port" env:"PORT"`
	IDHasher  IDHasher `yaml:"id_hasher" env:"-"`
	MySQL     MySQL    `yaml:"mysql" env:"-"`
	JWTSecret string   `yaml:"jwt_secret" env:"-JWT_SECRET"`
}

func LoadConfig() (Config, error) {
	cfg := Config{}

	err := cleanenv.ReadConfig("app/config/config.yml", &cfg)
	if err != nil {
		return Config{}, fmt.Errorf("read config from yml file: %w", err)
	}

	err = cleanenv.ReadEnv(&cfg)
	if err != nil {
		return Config{}, fmt.Errorf("read config from env: %w", err)
	}

	return cfg, nil
}
