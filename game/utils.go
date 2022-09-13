package game

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

func logErrorAndExit(msg string, a ...interface{}) {
	_, _ = fmt.Fprintf(os.Stderr, "ERROR: "+strings.TrimSpace(msg)+"\n", a...)
	cleanup()
	os.Exit(-1)
}

func moveCursorTo(x int, y int, interval time.Duration) {
	for cursor.X < x {
		cursor.X++
		time.Sleep(interval)
	}
	for cursor.X > x {
		cursor.X--
		time.Sleep(interval)
	}
	for cursor.Y < y {
		cursor.Y++
		time.Sleep(interval)
	}
	for cursor.Y > y {
		cursor.Y--
		time.Sleep(interval)
	}
}

func roundToNearest10th(n int) int {
	if n == 0 {
		return 1
	}

	rsp := (n / 10) * 10
	if rsp <= 0 {
		rsp = rand.Intn(n)
	}
	return rsp
}
