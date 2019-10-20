// Copyright (c) 2018 Dean Jackson <deanishe@deanishe.net>
// MIT Licence applies http://opensource.org/licenses/MIT
package env

import (
	"fmt"
	"net/url"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type Nested struct {
	NestedString string
	NestedNum    int
}

type TestNested struct {
	Nested
}

type TestInvalid struct {
	Map map[string]string
}

type BindTarget struct {
	Ignored    string `env:"-"`
	unexported string `env:"STRING"`

	String   string
	StringP  *string `env:"STRING"`
	Strings  []string
	StringsP []*string `env:"STRINGS"`

	Bool   bool
	BoolP  *bool `env:"BOOL"`
	Bools  []bool
	BoolsP []*bool `env:"BOOLS"`

	Int   int
	IntP  *int `env:"INT"`
	Ints  []int
	IntsP []*int `env:"INTS"`

	Int8   int8
	Int8P  *int8 `env:"INT8"`
	Ints8  []int8
	Ints8P []*int8 `env:"INTS8"`

	Int16   int16
	Int16P  *int16 `env:"INT16"`
	Ints16  []int16
	Ints16P []*int16 `env:"INTS16"`

	Int32   int32
	Int32P  *int32 `env:"INT32"`
	Ints32  []int32
	Ints32P []*int32 `env:"INTS32"`

	Int64   int64
	Int64P  *int64 `env:"INT64"`
	Ints64  []int64
	Ints64P []*int64 `env:"INTS64"`

	Uint   uint
	UintP  *uint `env:"UINT"`
	Uints  []uint
	UintsP []*uint `env:"UINTS"`

	Uint8   uint8
	Uint8P  *uint8 `env:"UINT8"`
	Uints8  []uint8
	Uints8P []*uint8 `env:"UINTS8"`

	Uint16   uint16
	Uint16P  *uint16 `env:"UINT16"`
	Uints16  []uint16
	Uints16P []*uint16 `env:"UINTS16"`

	Uint32   uint32
	Uint32P  *uint32 `env:"UINT32"`
	Uints32  []uint32
	Uints32P []*uint32 `env:"UINTS32"`

	Uint64   uint64
	Uint64P  *uint64 `env:"UINT64"`
	Uints64  []uint64
	Uints64P []*uint64 `env:"UINTS64"`

	Float32   float32
	Float32P  *float32 `env:"FLOAT32"`
	Floats32  []float32
	Floats32P []*float32 `env:"FLOATS32"`

	Float64   float64
	Float64P  *float64 `env:"FLOAT64"`
	Floats64  []float64
	Floats64P []*float64 `env:"FLOATS64"`

	Duration   time.Duration
	DurationP  *time.Duration `env:"DURATION"`
	Durations  []time.Duration
	DurationsP []*time.Duration `env:"DURATIONS"`

	URL   url.URL
	URLP  *url.URL `env:"URL"`
	URLS  []url.URL
	URLSP []*url.URL `env:"URLS"`

	// *time.Time implements encoding.TextMarshaler
	TimeP  *time.Time   `env:"TIME"`
	TimesP []*time.Time `env:"TIMES"`

	Nested *Nested

	Undefined struct {
		UndefinedField string
	}
}

func TestBind(t *testing.T) {
	var (
		str1 = "one"
		str2 = "two"

		bool1 = true
		bool2 = true

		int1          = 10
		int2          = 20
		int8_1  int8  = 1
		int8_2  int8  = 2
		int16_1 int16 = 12
		int16_2 int16 = 24
		int32_1 int32 = 64
		int32_2 int32 = 128
		int64_1 int64 = 99
		int64_2 int64 = 198

		uint1    uint   = 10
		uint2    uint   = 20
		uint8_1  uint8  = 1
		uint8_2  uint8  = 2
		uint16_1 uint16 = 12
		uint16_2 uint16 = 24
		uint32_1 uint32 = 64
		uint32_2 uint32 = 128
		uint64_1 uint64 = 99
		uint64_2 uint64 = 198

		float32_1 float32 = 1.1
		float32_2 float32 = 2.1

		float64_1 = 3.1
		float64_2 = 4.1

		duration1 = time.Hour
		duration2 = time.Minute

		url1, _ = url.Parse("http://www.example.com")
		url2, _ = url.Parse("http://www.example.org")

		// Time.MarshalText has 1-second resolution
		time1 = time.Now().Truncate(time.Second)
		time2 = time.Now().Add(-time.Hour).Truncate(time.Second)
	)

	str := func(v interface{}) string {
		return fmt.Sprintf("%v", v)
	}
	strs := func(v ...interface{}) string {
		s := make([]string, len(v))
		for i, val := range v {
			s[i] = fmt.Sprintf("%v", val)
		}
		return strings.Join(s, ",")
	}

	testEnv := MapEnv{
		"IGNORED": "not empty",
		"STRING":  str1,
		"STRINGS": strings.Join([]string{str1, str2}, ","),

		"BOOL":  str(bool1),
		"BOOLS": strs(bool1, bool2),

		"INT":    str(int1),
		"INTS":   strs(int1, int2),
		"INT8":   str(int8_1),
		"INTS8":  strs(int8_1, int8_2),
		"INT16":  str(int16_1),
		"INTS16": strs(int16_1, int16_2),
		"INT32":  str(int32_1),
		"INTS32": strs(int32_1, int32_2),
		"INT64":  str(int64_1),
		"INTS64": strs(int64_1, int64_2),

		"UINT":    str(uint1),
		"UINTS":   strs(uint1, uint2),
		"UINT8":   str(uint8_1),
		"UINTS8":  strs(uint8_1, uint8_2),
		"UINT16":  str(uint16_1),
		"UINTS16": strs(uint16_1, uint16_2),
		"UINT32":  str(uint32_1),
		"UINTS32": strs(uint32_1, uint32_2),
		"UINT64":  str(uint64_1),
		"UINTS64": strs(uint64_1, uint64_2),

		"FLOAT32":  str(float32_1),
		"FLOATS32": strs(float32_1, float32_2),

		"FLOAT64":  str(float64_1),
		"FLOATS64": strs(float64_1, float64_2),

		"DURATION":  str(duration1),
		"DURATIONS": strs(duration1, duration2),

		"URL":  str(url1),
		"URLS": strs(url1, url2),

		"TIME": time1.Format(time.RFC3339),
		"TIMES": strings.Join([]string{
			time1.Format(time.RFC3339),
			time2.Format(time.RFC3339),
		}, ","),

		"UNDEFINED_FIELD": str1,
	}

	// expected result
	x := BindTarget{
		Ignored:    "",
		unexported: "",
		String:     str1,
		StringP:    &str1,
		Strings:    []string{str1, str2},
		StringsP:   []*string{&str1, &str2},

		Bool:   bool1,
		BoolP:  &bool1,
		Bools:  []bool{bool1, bool2},
		BoolsP: []*bool{&bool1, &bool2},

		Int:   int1,
		IntP:  &int1,
		Ints:  []int{int1, int2},
		IntsP: []*int{&int1, &int2},

		Int8:   int8_1,
		Int8P:  &int8_1,
		Ints8:  []int8{int8_1, int8_2},
		Ints8P: []*int8{&int8_1, &int8_2},

		Int16:   int16_1,
		Int16P:  &int16_1,
		Ints16:  []int16{int16_1, int16_2},
		Ints16P: []*int16{&int16_1, &int16_2},

		Int32:   int32_1,
		Int32P:  &int32_1,
		Ints32:  []int32{int32_1, int32_2},
		Ints32P: []*int32{&int32_1, &int32_2},

		Int64:   int64_1,
		Int64P:  &int64_1,
		Ints64:  []int64{int64_1, int64_2},
		Ints64P: []*int64{&int64_1, &int64_2},

		Uint:   uint1,
		UintP:  &uint1,
		Uints:  []uint{uint1, uint2},
		UintsP: []*uint{&uint1, &uint2},

		Uint8:   uint8_1,
		Uint8P:  &uint8_1,
		Uints8:  []uint8{uint8_1, uint8_2},
		Uints8P: []*uint8{&uint8_1, &uint8_2},

		Uint16:   uint16_1,
		Uint16P:  &uint16_1,
		Uints16:  []uint16{uint16_1, uint16_2},
		Uints16P: []*uint16{&uint16_1, &uint16_2},

		Uint32:   uint32_1,
		Uint32P:  &uint32_1,
		Uints32:  []uint32{uint32_1, uint32_2},
		Uints32P: []*uint32{&uint32_1, &uint32_2},

		Uint64:   uint64_1,
		Uint64P:  &uint64_1,
		Uints64:  []uint64{uint64_1, uint64_2},
		Uints64P: []*uint64{&uint64_1, &uint64_2},

		Float32:   float32_1,
		Float32P:  &float32_1,
		Floats32:  []float32{float32_1, float32_2},
		Floats32P: []*float32{&float32_1, &float32_2},

		Float64:   float64_1,
		Float64P:  &float64_1,
		Floats64:  []float64{float64_1, float64_2},
		Floats64P: []*float64{&float64_1, &float64_2},

		Duration:   duration1,
		DurationP:  &duration1,
		Durations:  []time.Duration{duration1, duration2},
		DurationsP: []*time.Duration{&duration1, &duration2},

		URL:   *url1,
		URLP:  url1,
		URLS:  []url.URL{*url1, *url2},
		URLSP: []*url.URL{url1, url2},

		TimeP:  &time1,
		TimesP: []*time.Time{&time1, &time2},
	}
	x.Undefined.UndefinedField = str1

	target := BindTarget{}
	require.NoError(t, Bind(&target, testEnv), "bind failed")
	assert.Equal(t, x, target, "unexpected result")
}

func TestBind_empty(t *testing.T) {
	x := BindTarget{}
	env := MapEnv{}
	target := BindTarget{}
	require.NoError(t, Bind(&target, env), "bind failed")
	assert.Equal(t, x, target, "unexpected value")
}

func TestBind_nested(t *testing.T) {
	env := MapEnv{
		"NESTED_STRING": "Nested",
		"NESTED_NUM":    "10",
	}

	x := BindTarget{
		Nested: &Nested{
			NestedString: "Nested",
			NestedNum:    10,
		},
	}

	target := BindTarget{Nested: &Nested{}}
	require.NoError(t, Bind(&target, env), "bind failed")
	assert.Equal(t, x, target, "unexpected result")
}

func TestBind_embedded(t *testing.T) {
	env := MapEnv{
		"NESTED_STRING": "Nested",
		"NESTED_NUM":    "10",
	}

	x := TestNested{
		Nested{
			NestedString: "Nested",
			NestedNum:    10,
		},
	}

	target := TestNested{}
	require.NoError(t, Bind(&target, env), "bind failed")
	assert.Equal(t, x, target, "unexpected result")
}

func TestBind_invalidTypes(t *testing.T) {
	var i int
	env := MapEnv{
		"MAP":   "blah",
		"SLICE": "blah,blah",
	}
	invalid := []struct {
		name string
		v    interface{}
	}{
		// unsupported types
		{"string slice", []string{}},
		{"map", map[string]string{}},
		{"int", i},
		{"&int", &i},
		// not a pointer
		{"struct", struct{}{}},
	}
	for _, td := range invalid {
		td := td
		t.Run(td.name, func(t *testing.T) {
			assert.EqualError(t, Bind(td.v, env), "not a pointer to a struct", "unexpected result")
		})
	}

	type embedded struct {
		TestInvalid
	}
	unsupported := []struct {
		name string
		v    interface{}
		err  string
	}{
		{
			"map field",
			&TestInvalid{},
			"unsupported type: map[string]string",
		},
		{
			"Nested map field",
			&struct {
				Nested *TestInvalid
			}{Nested: &TestInvalid{}},
			"unsupported type: map[string]string",
		},
		{
			"embedded map field",
			&embedded{},
			"unsupported type: map[string]string",
		},
		{
			"embedded map slice field",
			&struct {
				Slice []map[string]string
			}{},
			"unsupported type: map[string]string",
		},
	}
	for _, td := range unsupported {
		td := td
		t.Run(td.name, func(t *testing.T) {
			assert.EqualError(t, Bind(td.v, env), td.err, "unexpected error")
		})
	}
}

func TestBind_invalidValues(t *testing.T) {
	tests := []struct {
		key, val string
	}{
		{"BOOL", "dave"},
		{"BOOLS", "siegfried,roy"},

		{"INT", "dave"},
		{"INTS", "siegfried,roy"},
		{"INT8", "dave"},
		{"INTS8", "siegfried,roy"},
		{"INT16", "dave"},
		{"INTS16", "siegfried,roy"},
		{"INT32", "dave"},
		{"INTS32", "siegfried,roy"},
		{"INT64", "dave"},
		{"INTS64", "siegfried,roy"},

		{"UINT", "dave"},
		{"UINTS", "siegfried,roy"},
		{"UINT8", "dave"},
		{"UINTS8", "siegfried,roy"},
		{"UINT16", "dave"},
		{"UINTS16", "siegfried,roy"},
		{"UINT32", "dave"},
		{"UINTS32", "siegfried,roy"},
		{"UINT64", "dave"},
		{"UINTS64", "siegfried,roy"},

		{"FLOAT32", "dave"},
		{"FLOATS32", "siegfried,roy"},
		{"FLOAT64", "dave"},
		{"FLOATS64", "siegfried,roy"},

		{"DURATION", "dave"},
		{"DURATIONS", "siegfried,roy"},

		{"URL", ":"},
		{"URLS", ":,:"},

		{"TIME", "dave"},
		{"TIMES", "siegfried,roy"},
	}

	for _, td := range tests {
		td := td
		t.Run(td.key, func(t *testing.T) {
			target := BindTarget{}
			env := MapEnv{
				td.key: td.val,
			}
			assert.Errorf(t, Bind(&target, env), "%s: invalid value accepted")
		})
	}

	target := struct {
		Oops badMarshaller
	}{
		badMarshaller("oops"),
	}

	_, err := Dump(target)
	assert.Error(t, err, "dump accepted bogus target")
}

// Populate a struct from environment variables.
func ExampleBind() {
	// Simple configuration struct
	type config struct {
		HostName     string        `env:"HOSTNAME"` // default: HOST_NAME
		UserName     string        `env:"USERNAME"` // default: USER_NAME
		SSL          bool          `env:"USE_SSL"`  // default: SSL
		Port         int           // leave as default (PORT)
		PingInterval time.Duration `env:"PING"` // default: PING_INTERVAL
		Online       bool          `env:"-"`    // ignore this field
	}

	// Set some values in the environment for test purposes
	_ = os.Setenv("HOSTNAME", "api.example.com")
	_ = os.Setenv("USERNAME", "") // empty
	_ = os.Setenv("PORT", "443")
	_ = os.Setenv("USE_SSL", "1")
	_ = os.Setenv("PING", "5m")
	_ = os.Setenv("ONLINE", "1") // will be ignored

	// Create a config and bind it to the environment
	c := &config{}
	if err := Bind(c); err != nil {
		// handle error...
	}

	// config struct now populated from the environment
	fmt.Println(c.HostName)
	fmt.Println(c.UserName)
	fmt.Printf("%d\n", c.Port)
	fmt.Printf("%v\n", c.SSL)
	fmt.Printf("%v\n", c.Online)
	fmt.Printf("%v\n", c.PingInterval*4) // it's not a string!
	// Output:
	// api.example.com
	//
	// 443
	// true
	// false
	// 20m0s

	os.Clearenv()
}

// In contrast to the Get* functions, Bind treats empty variables
// the same as unset ones and ignores them.
func ExampleBind_emptyVars() {
	type config struct {
		Username string
		Email    string
	}

	// Defaults
	c := &config{
		Username: "bob",
		Email:    "bob@aol.com",
	}

	_ = os.Setenv("USERNAME", "dave") // different value
	_ = os.Setenv("EMAIL", "")        // empty value, ignored by Bind()

	// Bind config to environment
	if err := Bind(c); err != nil {
		panic(err)
	}

	fmt.Println(c.Username)
	fmt.Println(c.Email)

	// Output:
	// dave
	// bob@aol.com

	os.Clearenv()
}

// TestVarName tests the envvar name algorithm.
func TestVarName(t *testing.T) {
	data := []struct {
		in, x string
	}{
		{"URL", "URL"},
		{"Name", "NAME"},
		{"LastName", "LAST_NAME"},
		{"URLEncoding", "URL_ENCODING"},
		{"LongBeard", "LONG_BEARD"},
		{"HTML", "HTML"},
		{"etc", "ETC"},
	}

	for _, td := range data {
		s := VarName(td.in)
		assert.Equal(t, td.x, s, "unexpected VarName")
	}
}

// Example output of VarName.
func ExampleVarName() {
	// single-case words are upper-cased
	fmt.Println(VarName("URL"))
	fmt.Println(VarName("name"))
	// words that start with fewer than 3 uppercase chars are
	// upper-cased
	fmt.Println(VarName("Folder"))
	fmt.Println(VarName("MTime"))
	// but with 3+ uppercase chars, the last is treated as the first
	// char of the next word
	fmt.Println(VarName("VIPath"))
	fmt.Println(VarName("URLEncoding"))
	fmt.Println(VarName("SSLPort"))
	// camel-case words are split on the case changes
	fmt.Println(VarName("LastName"))
	fmt.Println(VarName("LongHorse"))
	fmt.Println(VarName("loginURL"))
	fmt.Println(VarName("newHomeAddress"))
	fmt.Println(VarName("PointA"))
	// digits are considered as the end of a word, not the start
	fmt.Println(VarName("b2B"))
	// Output:
	// URL
	// NAME
	// FOLDER
	// MTIME
	// VI_PATH
	// URL_ENCODING
	// SSL_PORT
	// LAST_NAME
	// LONG_HORSE
	// LOGIN_URL
	// NEW_HOME_ADDRESS
	// POINT_A
	// B2_B
}

func TestIsCamelCase(t *testing.T) {
	data := []struct {
		s string
		x bool
	}{
		{"", false},
		{"URL", false},
		{"url", false},
		{"Url", false},
		{"HomeAddress", true},
		{"myHomeAddress", true},
		{"PlaceA", true},
		{"myPlaceB", true},
		{"myB", true},
		{"my2B", true},
		{"B2B", false},
		{"SSLPort", true},
	}

	for _, td := range data {
		v := isCamelCase(td.s)
		assert.Equal(t, td.x, v, "unexpected result")
	}
}

func TestSplitCamelCase(t *testing.T) {
	data := []struct {
		in string
		x  string
	}{
		{"", ""},
		{"HomeAddress", "HOME_ADDRESS"},
		{"homeAddress", "HOME_ADDRESS"},
		{"loginURL", "LOGIN_URL"},
		{"SSLPort", "SSL_PORT"},
		{"HomeAddress", "HOME_ADDRESS"},
		{"myHomeAddress", "MY_HOME_ADDRESS"},
		{"PlaceA", "PLACE_A"},
		{"myPlaceB", "MY_PLACE_B"},
		{"myB", "MY_B"},
		{"my2B", "MY2_B"},
		{"URLEncoding", "URL_ENCODING"},
	}

	for _, td := range data {
		s := splitCamelCase(td.in)
		assert.Equal(t, td.x, s, "unexpected result")
	}
}
