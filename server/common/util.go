package common

import (
	"fmt"
	"log"
)

func OnError(err error, format string, v ...any) {
	if err != nil {
		if err != nil {
			if format != "" {
				log.Panicf("err = %v, reason: %v", err, fmt.Sprintf(format, v...))
			} else {
				log.Panic(err)
			}
		} else {
			log.Panicf(format, v...)
		}
	}
}
