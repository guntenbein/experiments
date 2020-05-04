package main

import (
	"bitbucket.org/inturnco/go-sdk/errors"
	"bitbucket.org/inturnco/go-sdk/logger"
	"time"
)

func main() {
	sugar := logger.NewZapLogAdapter(false, []logger.ErrorOption{
		logger.KeyValueErrOption(),
		logger.TraceErrOption(),
	}).
		With("meta", map[string]interface{}{
			"service": "experiments",
		})
	for {
		sugar.Error(errors.New("generic with stacktrace"))
		time.Sleep(time.Second)
	}
	//Flush logger before exit
	defer func() {
		err := sugar.Sync()
		if err != nil {
			sugar.Error(errors.Wrap(err, "sync sugar"))
		}
	}()
}
