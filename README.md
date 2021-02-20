# csvlogger

loging struct to csv splited by day.

## usage

### sample code.
```
func main() {
	logger, _ := NewLogger("./", "sample")

	type sample struct {
		// export Field.
		Hoge string
		Fuga int
	}

	for i := 0; i < 10; i++ {
		data := sample{
			Hoge: "hoge_" + fmt.Sprint(i),
			Fuga: i,
		}
		logger.Add(&data)
	}
}
```

### result

`sample.20210220.csv`
```
Hoge,Fuga
hoge_0,0
hoge_1,1
hoge_2,2
hoge_3,3
hoge_4,4
hoge_5,5
hoge_6,6
hoge_7,7
hoge_8,8
hoge_9,9
```
