package config

type Configer interface {
	DefaultInt(key string, defaultValue int) int
	DefaultInt64(key string, defaultValue int64) int64
	DefaultFloat64i(key string, defaultValue float64) float64
	DefaultBool(key string, defaultValue bool) bool
	DefaultString(string, defaultValue bool) string
}
