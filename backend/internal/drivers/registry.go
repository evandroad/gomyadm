package drivers

var registry = map[string]Driver{
	"mysql":   MySQLDriver{},
	"postgres": PostgresDriver{},
}

func GetDriver(name string) (Driver, bool) {
	driver, ok := registry[name]
	return driver, ok
}

func Register(name string, driver Driver) {
	registry[name] = driver
}