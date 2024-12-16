package app

import (
	"DC/FnO/messages"
	"DC/FnO/pkg/config"
	"DC/FnO/pkg/db"
	"DC/FnO/pkg/logger"
	"DC/FnO/pkg/models"
	"context"
	"fmt"
	"log"
	"net"
	"strconv"
	"sync"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gorm.io/gorm"
)

func Start(c context.Context) error {
	var conn net.Conn
	cnf := config.GetConfig()
	logLevel, err := strconv.Atoi(cnf.GetString("log.Level"))
	if err != nil {
		log.Panicln("Invalid log config: ", err)
		return err
	}
	logger.LoggerInit(cnf.GetString("log.path"), zapcore.Level(logLevel))

	// Convert viper config to the custom Config type
	configData, err := config.LoadConfigFromFile()
	if configData == nil {
		logger.Log().Error("Failed to convert Viper config to Config struct", zap.Error(err))
	}

	// Initialize the database using the config
	db, err := db.NewPostgresDB(configData, logger.Log())
	if err != nil {
		logger.Log().Error("Failed to connect to the PostgreSQL database", zap.Error(err))
		return err
	}

	conn, err = net.Dial("tcp", fmt.Sprintf("%s:%s", cnf.GetString("gateway.host"), cnf.GetString("gateway.port")))
	logger.Log().Info("connection with gateway established")
	if err != nil {
		logger.Log(c).Error("Failed to connect to gateway", zap.Error(err))
		return err
	}
	defer conn.Close()

	streamCount, err := messages.UserLogin(c, conn)
	if err != nil {
		logger.Log(c).Error("User login failed", zap.Error(err))
		return err
	}

	if err := processTrades(c, conn, streamCount, db.DB); err != nil {
		logger.Log(c).Error("Failed to process trades", zap.Error(err))
		return err
	}

	return nil
}

func processTrades(c context.Context, conn net.Conn, streamCount int, db *gorm.DB) error {
	// Channels for trade data and errors
	tradeChannel := make(chan models.ExtractTrade, 1000)
	errorChannel := make(chan error)

	// WaitGroup to wait for all worker goroutines to finish
	var wg sync.WaitGroup

	// Start worker goroutines to process trade data
	numWorkers := 5
	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go messages.ProcessTrade(db, &tradeChannel, &errorChannel, &wg)
	}

	// Start error handling goroutine
	go messages.ProcessError(&errorChannel)

	var initialSeqNumber uint32
	// Loop to run streamCount times
	for i := 0; i < streamCount; i++ {
		// logger.Log(c).Info("preapering to download trade for trade no. of ", zap.Int("iteration", i+1))

		if err := messages.TradeRequest(c, conn, &tradeChannel, i+1, &initialSeqNumber); err != nil {
			logger.Log(c).Error("Trade request failed", zap.Error(err), zap.Int("iteration", i+1))
			return err
		}

		// Optional: you can also add any logic here to handle different behaviors based on the iteration count
	}

	// Close the trade channel to signal the workers to stop
	close(tradeChannel)

	// Wait for all worker goroutines to complete
	wg.Wait()

	return nil
}
