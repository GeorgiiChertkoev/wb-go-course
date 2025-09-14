package main

import (
	"fmt"
	"os"
	"time"

	"github.com/beevik/ntp"
)

func main() {
	ntpTime, err := ntp.Time("0.beevik-ntp.pool.ntp.org")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Ошибка получения времени: %s\n", err)
		os.Exit(1)
	}

	fmt.Printf("Нынешнее время: %s", ntpTime.Format(time.RFC3339))
}
