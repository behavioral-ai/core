package fmtx

import "fmt"

func ExampleParseDuration() {
	s := ""
	duration, err := ParseDuration(s)
	fmt.Printf("test: ParseDuration(\"%v\") [err:%v] [duration:%v]\n", s, err, duration)

	s = "  "
	duration, err = ParseDuration(s)
	fmt.Printf("test: ParseDuration(\"%v\") [err:%v] [duration:%v]\n", s, err, duration)

	s = "12as"
	duration, err = ParseDuration(s)
	fmt.Printf("test: ParseDuration(\"%v\") [err:%v] [duration:%v]\n", s, err, duration)

	s = "1000"
	duration, err = ParseDuration(s)
	fmt.Printf("test: ParseDuration(\"%v\") [err:%v] [duration:%v]\n", s, err, duration)

	s = "1000s"
	duration, err = ParseDuration(s)
	fmt.Printf("test: ParseDuration(\"%v\") [err:%v] [duration:%v]\n", s, err, duration)

	s = "1000m"
	duration, err = ParseDuration(s)
	fmt.Printf("test: ParseDuration(\"%v\") [err:%v] [duration:%v]\n", s, err, duration)

	s = "1m"
	duration, err = ParseDuration(s)
	fmt.Printf("test: ParseDuration(\"%v\") [err:%v] [duration:%v]\n", s, err, duration)

	s = "10ms"
	duration, err = ParseDuration(s)
	fmt.Printf("test: ParseDuration(\"%v\") [err:%v] [duration:%v]\n", s, err, duration)

	//t := time.Microsecond * 100
	//fmt.Printf("test: time.String %v\n", t.String())

	s = "10µs"
	duration, err = ParseDuration(s)
	fmt.Printf("test: ParseDuration(\"%v\") [err:%v] [duration:%v]\n", s, err, duration)

	//Output:
	//test: ParseDuration("") [err:<nil>] [duration:0s]
	//test: ParseDuration("  ") [err:strconv.Atoi: parsing "  ": invalid syntax] [duration:0s]
	//test: ParseDuration("12as") [err:strconv.Atoi: parsing "12a": invalid syntax] [duration:0s]
	//test: ParseDuration("1000") [err:<nil>] [duration:16m40s]
	//test: ParseDuration("1000s") [err:<nil>] [duration:16m40s]
	//test: ParseDuration("1000m") [err:<nil>] [duration:16h40m0s]
	//test: ParseDuration("1m") [err:<nil>] [duration:1m0s]
	//test: ParseDuration("10ms") [err:<nil>] [duration:10ms]
	//test: ParseDuration("10µs") [err:<nil>] [duration:10µs]

}
