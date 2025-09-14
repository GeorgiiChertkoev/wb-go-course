package main

import (
	"fmt"
	"os"
	"time"

	"github.com/beevik/ntp"
)

const StandartServerAddress = "0.beevik-ntp.pool.ntp.org"

func GetNtpTime(address string) (time.Time, error) {
	return ntp.Time(address)
}

func main() {
	for {
		ntpTime, err := GetNtpTime(StandartServerAddress)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Ошибка получения времени: %s\n", err)
			os.Exit(1)
		}

		fmt.Printf("Нынешнее время: %s\n", ntpTime.Format(time.RFC3339))
		<-time.After(1 * time.Second)
	}
}
