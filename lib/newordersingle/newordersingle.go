package newordersingle

import (
	"fmt"
	"time"

	"github.com/adrianguyareach/FIXlib/lib/utils"
	"github.com/quickfixgo/enum"
	"github.com/quickfixgo/field"
	fix44nos "github.com/quickfixgo/fix44/newordersingle"
	"github.com/quickfixgo/quickfix"
)

func Fix44newordersingle(initiator *quickfix.Initiator, appSettings *quickfix.Settings) {

	var orderType field.OrdTypeField

	orderType.FIXString = quickfix.FIXString("Market")
	newOrderSingle := fix44nos.New(
		field.NewClOrdID("12345"),
		field.NewSide(enum.Side_BUY),
		field.NewTransactTime(time.Now()),
		orderType,
	)

	// body := newOrderSingle.ToMessage()

	// SET THE FIX MESSAGE HEADER
	// header.SetHeader(newOrderSingle.ToMessage().Header, appSettings)

	// // Get sender compID
	// senderCompID, err := header.GetString(49)
	// if err != nil {
	// 	utils.PrintBad(fmt.Sprintf("unable to fetch senderCompID: %s", err))
	// }

	// targetCompID, err := header.GetString(56)
	// if err != nil {
	// 	utils.PrintBad(fmt.Sprintf("unable to fetch targetCompID: %s", err))
	// }

	// fmt.Printf("sender: %s target: %s\n", senderCompID, targetCompID)
	newOrderSingle.SetBeginString("FIX.4.4")
	newOrderSingle.SetTargetCompID("ISLD")
	newOrderSingle.SetSenderCompID("TW")
	// newOrderSingle.Set(field.NewSenderCompID("TW"))
	// newOrderSingle.SetField(56, field.NewTargetCompID("ISLD"))

	msg := newOrderSingle.ToMessage()

	for {

		utils.PrintInfo(msg.String())
		senderr := quickfix.Send(msg)
		time.Sleep(time.Second)

		if senderr != nil {

			message := fmt.Sprintf("Failure sending nos message: %s", senderr)
			utils.PrintBad(message)
		}

	}

}
