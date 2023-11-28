package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"os/signal"
	"path"
	"syscall"

	"github.com/adrianguyareach/FIXlib/lib/advertisement"
	"github.com/adrianguyareach/FIXlib/lib/utils"
	"github.com/quickfixgo/quickfix"
)

type TradeClient struct {
}

func (e TradeClient) OnCreate(sessionID quickfix.SessionID) {}
func (e TradeClient) OnLogon(sessionID quickfix.SessionID)  {}
func (e TradeClient) OnLogout(sessionID quickfix.SessionID) {}
func (e TradeClient) FromAdmin(msg *quickfix.Message, sessionID quickfix.SessionID) (reject quickfix.MessageRejectError) {
	return nil
}
func (e TradeClient) ToAdmin(msg *quickfix.Message, sessionID quickfix.SessionID) {}
func (e TradeClient) ToApp(msg *quickfix.Message, sessionID quickfix.SessionID) (err error) {
	utils.PrintInfo(fmt.Sprintf("Sending: %s", msg.String()))
	return
}
func (e TradeClient) FromApp(msg *quickfix.Message, sessionID quickfix.SessionID) (reject quickfix.MessageRejectError) {
	utils.PrintInfo(fmt.Sprintf("FromApp: %s", msg.String()))
	return
}

func main() {

	var cfgFileName string
	args := os.Args
	argLen := len(args)
	switch argLen {
	case 1:
		{
			utils.PrintInfo("FIX config file not provided...")
			utils.PrintInfo("attempting to use default location './config/tradeclient.cfg' ...")
			cfgFileName = path.Join("config", "tradeclient.cfg")
		}
	case 2:
		{
			cfgFileName = args[0]
		}
	default:
		{
			utils.PrintBad("incorrect argument number")
		}
	}

	cfg, err := os.Open(cfgFileName)
	if err != nil {
		utils.PrintBad(fmt.Sprintf("error opening %v, %v", cfgFileName, err))
	}
	defer cfg.Close()

	stringData, readErr := io.ReadAll(cfg)
	if readErr != nil {
		utils.PrintBad(fmt.Sprintf("error reading cfg: %s,", readErr))
	}

	appSettings, err := quickfix.ParseSettings(bytes.NewReader(stringData))
	if err != nil {
		utils.PrintBad(fmt.Sprintf("error reading cfg: %s,", err))
	}
	appSettings.SessionSettings()

	app := TradeClient{}
	// fileLogFactory, err := quickfix.NewFileLogFactory(appSettings)
	logger := utils.NewFancyLog()

	if err != nil {
		utils.PrintBad(fmt.Sprintf("error creating file log factory: %s,", err))
	}

	initiator, err := quickfix.NewInitiator(app, quickfix.NewMemoryStoreFactory(), appSettings, logger)
	if err != nil {
		utils.PrintBad(fmt.Sprintf("unable to create initiator: %s", err))
	}

	err = initiator.Start()
	if err != nil {
		utils.PrintBad(fmt.Sprintf("unable to start initiator: %s", err))
	}

	utils.PrintConfig("initiator", bytes.NewReader(stringData))
	// go
	// go newordersingle.Fix44newordersingle()

	// go func() {
	// 	newordersingle.Fix44newordersingle(appSettings)
	// }()
	go func() {
		advertisement.Advertisement(appSettings)
	}()
	// go func() {
	// 	newordersingle.Fix44newordersingle(3)
	// }()
	// go func() {
	// 	newordersingle.Fix44newordersingle(4)
	// }()
	// go func() {
	// 	newordersingle.Fix44newordersingle(5)
	// }()

	// Handle interrupt signals to gracefully shut down the application
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	select {
	case <-interrupt:
		utils.PrintInfo("Received interrupt signal. Shutting down...")
		// Perform cleanup and stop the initiator
		initiator.Stop()
		utils.PrintInfo("Stopped FIX initiator.")
	}

	// select {}

}
