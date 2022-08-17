module github.com/KouKouChan/CSO2-Server

go 1.15

require (
	github.com/KouKouChan/iconv-go v0.0.0-20201112084036-b281919f810b
	github.com/djimenez/iconv-go v0.0.0-20160305225143-8960e66bd3da
	github.com/garyburd/redigo v1.6.0
	github.com/mattn/go-sqlite3 v1.14.0
	github.com/qiniu/iconv v1.2.0
	github.com/willf/bitset v1.1.11
	golang.org/x/text v0.3.3
	gopkg.in/alexcesaro/quotedprintable.v3 v3.0.0-20150716171945-2caba252f4dc // indirect
	gopkg.in/gomail.v2 v2.0.0-20160411212932-81ebce5c23df
	gopkg.in/ini.v1 v1.57.0
)

replace golang.org/x/text v0.3.3 => github.com/golang/text v0.3.3
