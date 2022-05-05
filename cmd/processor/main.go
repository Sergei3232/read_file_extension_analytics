package main

import (
	"github.com/Sergei3232/read_file_extension_analytics/config"
	"github.com/Sergei3232/read_file_extension_analytics/internal/app/counter_id"
	"github.com/Sergei3232/read_file_extension_analytics/internal/app/db"
	"github.com/sirupsen/logrus"
)

func main() {
	configs, err := config.NenConfig()
	if err != nil {
		logrus.Error(err.Error())
	}

	dbClient, errDb := db.NewDbConnectClient(configs.PathFileCountId)
	if errDb != nil {
		logrus.Error(errDb.Error())
	}

	counter, errCounter := counter_id.NewtextCounter(configs.PathFileCountId, configs.StartId)
	if errCounter != nil {
		logrus.Error(errCounter.Error())
	}

	id, err := counter.ReadFile(configs.PathFileCountId)
	logrus.Info(id, err)

	//for i:= 0; i < 100; i++{
	//	err := counter.SaveFile(i)
	//	if err != nil {
	//		logrus.Error(err.Error())
	//	}
	//}

	logrus.Info(dbClient, counter)
	logrus.Info("END SCRIPT")
}
