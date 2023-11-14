package tradeclient

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"os/signal"
	"path"
	"syscall"

	"github.com/adrianguyareach/FIXlib/lib/utils"
	"github.com/quickfixgo/quickfix"
)

// TradeClient implements the quickfix.Application interface
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

func Client(args []string) error {

	var cfgFileName string
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
			return fmt.Errorf("incorrect argument number")
		}
	}

	cfg, err := os.Open(cfgFileName)
	if err != nil {
		return fmt.Errorf("error opening %v, %v", cfgFileName, err)
	}
	defer cfg.Close()

	stringData, readErr := io.ReadAll(cfg)
	if readErr != nil {
		return fmt.Errorf("error reading cfg: %s,", readErr)
	}

	appSettings, err := quickfix.ParseSettings(bytes.NewReader(stringData))
	if err != nil {
		return fmt.Errorf("error reading cfg: %s,", err)
	}

	app := TradeClient{}
	fileLogFactory, err := quickfix.NewFileLogFactory(appSettings)

	if err != nil {
		return fmt.Errorf("error creating file log factory: %s,", err)
	}

	initiator, err := quickfix.NewInitiator(app, quickfix.NewMemoryStoreFactory(), appSettings, fileLogFactory)
	if err != nil {
		return fmt.Errorf("unable to create initiator: %s", err)
	}

	err = initiator.Start()
	if err != nil {
		return fmt.Errorf("unable to start initiator: %s", err)
	}

	utils.PrintConfig("initiator", bytes.NewReader(stringData))
	Fix44newordersingle()

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)
	<-interrupt

	utils.PrintInfo("stopping FIX initiator ..")
	defer initiator.Stop()
	utils.PrintInfo("stopped")

	return nil
}
