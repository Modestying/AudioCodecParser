package opus

import (
	"bufio"
	"fmt"
	"os"
	"testing"
)

func TestOpusData_Parse(t *testing.T) {
	printFunc := []InfoOpt{
		WithOpusVersion,
		WithNumberPageSegments,
		WithSegmentTable,
	}
	file := "./test.opus"
	fd, err := os.Open(file)
	if err != nil {
		panic(err)
	}
	defer func() {
		_ = fd.Close()
	}()
	reader := bufio.NewReader(fd)
	for i := 0; i < 5; i++ {
		d := &OpusData{}
		d.Parse(reader)
		fmt.Println(fmt.Sprintf("Pkt(%d)\n", i+1) +
			d.Info(printFunc...))
	}
}
