package client

import (
	"errors"
	"fmt"
	"log"
	"math"
	"os"
	"reflect"
	"strconv"
	"strings"
	"time"

	"github.com/go-routeros/routeros"
	"github.com/go-routeros/routeros/proto"
)

type Mikrotik struct {
	Host     string
	Username string
	Password string
}

func Unmarshal(reply routeros.Reply, v interface{}) error {
	rv := reflect.ValueOf(v)
	elem := rv.Elem()

	if rv.Kind() != reflect.Ptr {
		panic("Unmarshal cannot work without a pointer")
	}

	switch elem.Kind() {
	case reflect.Slice:
		l := len(reply.Re)
		if l <= 1 {
			panic(fmt.Sprintf("Cannot Unmarshal %d sentence(s) into a slice", l))
		}

		t := elem.Type()
		d := reflect.MakeSlice(t, l, l)

		for i := 0; i < l; i++ {
			item := d.Index(i)
			sentence := reply.Re[i]

			parseStruct(&item, *sentence)
		}
		elem.Set(d)

	case reflect.Struct:
		if len(reply.Re) < 1 {
			// This is an empty message
			return nil
		}
		if len(reply.Re) > 1 {
			msg := fmt.Sprintf("Failed to decode reply: %v", reply)
			return errors.New(msg)
		}

		parseStruct(&elem, *reply.Re[0])
	}

	return nil
}
func parseStruct(v *reflect.Value, sentence proto.Sentence) {
	elem := *v
	for i := 0; i < elem.NumField(); i++ {
		field := elem.Field(i)
		fieldType := elem.Type().Field(i)
		tags := strings.Split(fieldType.Tag.Get("mikrotik"), ",")

		path := strings.ToLower(fieldType.Name)
		fieldName := tags[0]

		for _, pair := range sentence.List {
			if strings.Compare(pair.Key, path) == 0 || strings.Compare(pair.Key, fieldName) == 0 {
				switch fieldType.Type.Kind() {
				case reflect.String:
					field.SetString(pair.Value)
				case reflect.Bool:
					b, _ := strconv.ParseBool(pair.Value)
					field.SetBool(b)
				case reflect.Int:
					if contains(tags, "ttlToSeconds") {
						field.SetInt(int64(ttlToSeconds(pair.Value)))
					}
				}

			}
		}
	}
}

func ttlToSeconds(ttl string) int {
	parts := strings.Split(ttl, "d")

	idx := 0
	days := 0
	var err error
	if len(parts) == 2 {
		idx = 1
		days, err = strconv.Atoi(parts[0])

		// We should be parsing an ascii number
		// if this fails we should fail loudly
		if err != nil {
			panic(err)
		}

		// In the event we just get days parts[1] will be an
		// empty string. Just coerce that into 0 seconds.
		if parts[1] == "" {
			parts[1] = "0s"
		}
	}
	d, err := time.ParseDuration(parts[idx])

	// We should never receive a duration greater than
	// 23h59m59s. So this should always parse.
	if err != nil {
		panic(err)
	}
	return 86400*days + int(d)/int(math.Pow10(9))

}

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func NewClient(host, username, password string) Mikrotik {
	return Mikrotik{
		Host:     host,
		Username: username,
		Password: password,
	}
}

func GetConfigFromEnv() (host, username, password string) {
	host = os.Getenv("MIKROTIK_HOST")
	username = os.Getenv("MIKROTIK_USER")
	password = os.Getenv("MIKROTIK_PASSWORD")
	if host == "" || username == "" || password == "" {
		// panic("Unable to find the MIKROTIK_HOST, MIKROTIK_USER or MIKROTIK_PASSWORD environment variable")
	}
	return host, username, password
}

func (client Mikrotik) getMikrotikClient() (c *routeros.Client, err error) {
	address := client.Host
	username := client.Username
	password := client.Password
	c, err = routeros.Dial(address, username, password)

	if err != nil {
		log.Printf("[ERROR] Failed to login to routerOS with error: %v", err)
	}

	return
}
