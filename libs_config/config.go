package libs_config

// Listener default config for go struct is created for automatic yaml generation so that it can be implemented in new projects.
type Listener struct {
	Listen string `yaml:"listen" default:"0.0.0.0"`
	Port   int    `yaml:"port" default:"8081"`
}

// SQLConfig SQL config is used to configure RDMS databases such as MySQL, PostgreSQL, and oraclesql struct is created for automatic YAML creation so that it can be implemented in new projects.
type SQLConfig struct {
	Enable          bool   `yaml:"enable" default:"false" desc:"config:sql:enable"`
	Driver          string `yaml:"driver" default:"" desc:"config:sql:driver"`
	Host            string `yaml:"host" default:"127.0.0.1" desc:"config:sql:host"`
	Port            int    `yaml:"port" default:"3306" desc:"config:sql:port"`
	Username        string `yaml:"username" default:"root"  desc:"config:sql:username"`
	Password        string `yaml:"password" default:"root" desc:"config:sql:password"`
	Database        string `yaml:"database" default:"mydb" desc:"config:sql:database"`
	Options         string `yaml:"options" default:"" desc:"config:sql:options"`
	Connection      string `yaml:"connection" default:"" desc:"config:sql:connection"`
	AutoReconnect   bool   `yaml:"autoreconnect" default:"false"  desc:"config:sql:autoreconnect"`
	StartInterval   int    `yaml:"startinterval" default:"2"  desc:"config:sql:startinterval"`
	MaxError        int    `yaml:"maxerror" default:"5"  desc:"config:sql:maxerror"`
	CustomPool      bool   `yaml:"customPool" default:"5"  desc:"config:sql:customPool"`
	MaxConn         int    `yaml:"maxConn" default:"5"  desc:"config:sql:maxConn"`
	MaxIdle         int    `yaml:"maxIdle" default:"5"  desc:"config:sql:maxIdle"`
	LifeTime        int    `yaml:"lifeTime" default:"5"  desc:"config:sql:lifeTime"`
	UseMock         bool   `yaml:"useMock" default:"5"  desc:"config:sql:useMock"`
	MultiStatements bool   `yaml:"multiStatements" default:"false"  desc:"config:sql:multiStatements"`
}

// RabbitMQConfig Rabbit config is used to configure structs created for automatic YAML generation so that they can be implemented in new projects.
type RabbitMQConfig struct {
	Enable              bool   `yaml:"enable" default:"false" desc:"config:rabbitmq:enable"`
	Host                string `yaml:"host" default:"127.0.0.1" desc:"config:rabbitmq:host"`
	Port                int    `yaml:"port" default:"5672" desc:"config:rabbitmq:port"`
	Username            string `yaml:"username" default:"guest"  desc:"config:rabbitmq:username"`
	Password            string `yaml:"password" default:"guest" desc:"config:rabbitmq:password"`
	ReconnectDuration   int    `yaml:"reconnectDuration" default:"5" desc:"config:rabbitmq:reconnectDuration"`
	DedicatedConnection bool   `yaml:"dedicatedConnection" default:"false" desc:"config:rabbitmq:dedicatedConnection"`
}

// RedisConfig REDIS config is used to configure structs created for automatic YAML generation so that they can be implemented in new projects.
type RedisConfig struct {
	Enable        bool   `yaml:"enable" default:"false" desc:"config:redis:enable"`
	Host          string `yaml:"host" default:"127.0.0.1" desc:"config:redis:host"`
	Port          int    `yaml:"port" default:"6379" desc:"config:redis:port"`
	Password      string `yaml:"password" default:"" desc:"config:redis:password"`
	Pool          int    `yaml:"pool" default:"10" desc:"config:redis:pool"`
	AutoReconnect bool   `yaml:"autoreconnect" default:"false"  desc:"config:autoreconnect"`
	StartInterval int    `yaml:"startinterval" default:"2"  desc:"config:startinterval"`
	MaxError      int    `yaml:"maxerror" default:"5"  desc:"config:maxerror"`
	PoolSize      int    `yaml:"poolize" default:"30" desc:"config:poolize"`
	PoolTimeout   int    `yaml:"pooltimeout" default:"30" desc:"config:pooltimeout"`
	MinIdleConns  int    `yaml:"minidleconns" default:"7" desc:"config:minidleconns"`
	MaxIdleConns  int    `yaml:"maxidleconns" default:"15" desc:"config:maxidleconns"`
	ConMaxLife    int    `yaml:"conmaxlife" default:"600" desc:"config:conmaxlife"`
	ConMaxIdle    int    `yaml:"conmaxidle" default:"600" desc:"config:conmaxidle"`
}

// KafkaConfig Kafka config is used to configure structs created for automatic YAML generation so that they can be implemented in new projects.
type KafkaConfig struct {
	Enable           bool   `yaml:"enable" default:"false" desc:"config:kafka:enable"`
	Host             string `yaml:"host" default:"127.0.0.1:9092" desc:"config:kafka:host"`
	Registry         string `yaml:"registry" default:"" desc:"config:kafka:registry"`
	Username         string `yaml:"username" default:""  desc:"config:kafka:username"`
	Password         string `yaml:"password" default:"" desc:"config:kafka:password"`
	SecurityProtocol string `yaml:"securityProtocol" default:"SASL_SSL"  desc:"config:kafka:securityProtocol"`
	Mechanisms       string `yaml:"mechanisms" default:"PLAIN"  desc:"config:kafka:mechanisms"`
	Debug            string `yaml:"debug" default:"consumer"  desc:"config:kafka:debug"`
}

// JWTConfig JWT config is used to configure structs created for automatic YAML generation so that they can be implemented in new projects.
type JWTConfig struct {
	Access         string `yaml:"access" default:"random"`
	Refresh        string `yaml:"refresh" default:"random"`
	ExpiredAccess  int    `yaml:"expiredAccess" default:"30"`
	ExpiredRefresh int    `yaml:"expiredRefresh" default:"24"`
}
