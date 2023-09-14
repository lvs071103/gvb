package config

import "strconv"

type MySQL struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	DB       string `yaml:"db"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	LogLevel string `yaml:"log_level"` // 日志等级，debug就是输出全部sql
	MaxIdleConns int `yaml:"max_idle"`
	MaxOpenConns int `yaml:"max_open"`
}


func (m *MySQL)Dsn() string {
	return m.User + ":" + m.Password + "@tcp(" + m.Host + ":" + strconv.Itoa(m.Port) + ")/" + m.DB
}
