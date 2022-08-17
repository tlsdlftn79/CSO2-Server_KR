package report

import (
	"bufio"
	"os"
	"time"
)

var (
	ReportPath string
)

func SaveReport(name, reporttype string, msg []byte) {
	fd_time := time.Now().Format("2006-01-02 15:04:05")

	fl, err := os.OpenFile(ReportPath+name+".txt", os.O_APPEND|os.O_CREATE, 0644)
	if err != nil {
		return
	}
	defer fl.Close()

	//写入文件时，使用带缓存的 *Writer
	write := bufio.NewWriter(fl)
	write.WriteString(fd_time + " [" + reporttype + "] " + string(msg))
	//Flush将缓存的文件真正写入到文件中
	write.Flush()
}
