package database

type Config struct {
	Type, Charset, Database, TimeZone string
	Host                              string `flag:"name:db_host;usage:数据库host"`
	User, Password                    string
	TimeFormat                        string
	MaxIdleConns, MaxOpenConns        int
	Port                              int32
}
