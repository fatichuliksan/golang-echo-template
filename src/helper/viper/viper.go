package viper

import (
	"strings"

	"github.com/spf13/viper"
)

// Method ...
type Method interface {
	GetString(key string) string
	GetInt(key string) int
	GetBool(key string) bool
	Init()
}

type viperHelper struct{}

// NewViper ...
func NewViper() Method {
	v := &viperHelper{}
	v.Init()
	return v
}

// Init ...
func (t *viperHelper) Init() {
	viper.SetEnvPrefix(`test`)
	viper.AutomaticEnv()

	replacer := strings.NewReplacer(`.`, `_`)
	viper.SetEnvKeyReplacer(replacer)
	viper.SetConfigType(`json`)
	viper.SetConfigFile(`config.json`)

	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}

// GetString ...
func (t *viperHelper) GetString(key string) string {
	return viper.GetString(key)
}

// GetInt ...
func (t *viperHelper) GetInt(key string) int {
	return viper.GetInt(key)
}

// GetBool ...
func (t *viperHelper) GetBool(key string) bool {
	return viper.GetBool(key)
}
