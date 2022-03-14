package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"reflect"
	"strconv"
	"strings"

	"github.com/hashicorp/hcl"
	"github.com/pelletier/go-toml"
	"gopkg.in/yaml.v3"
)

var (
	EnvDelimeter      = "_"
	EnvSliceDelimeter = ";"
)

type format string

const (
	YAML format = "yaml"
	JSON format = "json"
	TOML format = "toml"
	HCL  format = "hcl"
)

// AutoloadAndEnrichConfig takes a config file and a receiver and enriches the config with the value from env variables.
// @filePath: The path to the config file.
// @receiver: The receiver to parse the config file into.
// @prefix: The prefix to use for the env variables.
func AutoloadAndEnrichConfigWithEnvPrefix(filePath string, prefix string, receiver interface{}) error {
	fileFormat := detectFormat(filePath)
	err := loadAndParseFile(filePath, receiver, fileFormat)
	if err != nil {
		return err
	}
	readStructAndEnrichWithEnv(receiver, prefix)
	return nil
}

// AutoloadAndEnrichConfig takes a config file and a receiver and enriches the config with the value from env variables.
// @filePath: The path to the config file.
// @receiver: The receiver to parse the config file into.
//
// By default the prefix is set to CFG.
func AutoloadAndEnrichConfig(filePath string, receiver interface{}) error {
	return AutoloadAndEnrichConfigWithEnvPrefix(filePath, "CFG", receiver)
}

// detectFormat detects the format of the config file.
// @filePath: The path to the config file.
func detectFormat(filePath string) format {
	ext := path.Ext(filePath)
	switch ext {
	case ".yaml", ".yml":
		return YAML
	case ".json":
		return JSON
	case ".toml":
		return TOML
	case ".hcl":
		return HCL
	default:
		return ""
	}
}

// loadAndParseFile takes a config file and a receiver and parses the config file into the receiver.
// @filePath: The path to the config file.
// @receiver: The receiver to parse the config file into.
// @f: The format of the config file.
func loadAndParseFile(filePath string, receiver interface{}, f format) error {
	bts, err := ioutil.ReadFile(filePath)
	if err != nil {
		return err
	}
	switch f {
	case YAML:
		err := yaml.Unmarshal(bts, receiver)
		if err != nil {
			return err
		}
	case JSON:
		err := json.Unmarshal(bts, receiver)
		if err != nil {
			return err
		}
	case TOML:
		err := toml.Unmarshal(bts, receiver)
		if err != nil {
			return err
		}
	case HCL:
		err := hcl.Unmarshal(bts, receiver)
		if err != nil {
			return err
		}
	default:
		return fmt.Errorf("unsupported format: %s", f)
	}
	return nil
}

// prefixString returns the string s with prefix p.
func prefixString(prefix, fieldName string) string {
	if prefix == "" {
		return strings.ToUpper(fieldName)
	}
	return strings.ToUpper(fmt.Sprintf("%s%s%s", prefix, EnvDelimeter, fieldName))
}

func readStructAndEnrichWithEnv(st interface{}, prefix string) {
	val := reflect.ValueOf(st)
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}
	for i := 0; i < val.NumField(); i++ {
		// fmt.Println(val.Type().Field(i).Type.Kind())
		f := val.Field(i)
		prefixedFieldName := prefixString(prefix, val.Type().Field(i).Name)
		osEnv := os.Getenv(prefixedFieldName)
		switch f.Kind() {
		case reflect.Struct:
			readStructAndEnrichWithEnv(f.Addr().Interface(), prefixedFieldName)
		case reflect.Slice:
			_, ok := f.Interface().([]string)
			if !ok {
				// we only support []string from env
				continue
			}
			if osEnv == "" {
				// env var not set
				continue
			}
			f.Set(
				reflect.ValueOf(
					strings.Split(osEnv, EnvSliceDelimeter),
				),
			)
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			if osEnv == "" {
				continue
			}
			in, err := strconv.ParseInt(osEnv, 10, 64)
			if err != nil {
				// could not parse int
				// so we skip this field
				continue
			}
			if !f.IsValid() {
				f.SetInt(in)
				continue
			}
			if f.CanSet() {
				f.SetInt(in)
			}
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			if osEnv == "" {
				continue
			}
			uit, err := strconv.ParseUint(osEnv, 10, 64)
			if err != nil {
				// could not parse uint
				// so we skip this field
				continue
			}
			if !f.IsValid() {
				f.SetUint(uit)
				continue
			}
			if f.CanSet() {
				f.SetUint(uit)
			}
		case reflect.Float32, reflect.Float64:
			if osEnv == "" {
				continue
			}
			fl, err := strconv.ParseFloat(osEnv, 64)
			if err != nil {
				// could not parse float
				// so we skip this field
				continue
			}
			if !f.IsValid() {
				f.SetFloat(fl)
				continue
			}
			if f.CanSet() {
				f.SetFloat(fl)
			}
		case reflect.Bool:
			if osEnv == "" {
				continue
			}
			bl, err := strconv.ParseBool(osEnv)
			if err != nil {
				// could not parse bool
				// so we skip this field
				continue
			}
			if !f.IsValid() {
				f.SetBool(bl)
				continue
			}
			if f.CanSet() {
				f.SetBool(bl)
			}
		case reflect.String:
			if osEnv == "" {
				continue
			}
			if !f.IsValid() {
				f.SetString(osEnv)
				continue
			}
			if f.CanSet() {
				f.SetString(osEnv)
				continue
			}
		}
	}
}
