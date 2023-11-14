package tradeclient

import (
	"fmt"
	"time"

	"github.com/adrianguyareach/FIXlib/lib/utils"
	"github.com/quickfixgo/enum"
	"github.com/quickfixgo/field"
	fix44nos "github.com/quickfixgo/fix44/newordersingle"
	"github.com/quickfixgo/quickfix"
)

type header interface {
	Set(f quickfix.FieldWriter) *quickfix.FieldMap
}

func Fix44newordersingle() (err error) {

	var orderType field.OrdTypeField

	orderType.FIXString = quickfix.FIXString("Market")
	newOrderSingle := fix44nos.New(
		field.NewClOrdID("12345"),
		field.NewSide(enum.Side_BUY),
		field.NewTransactTime(time.Now()),
		orderType,
	)

	newOrderSingle.SetSenderCompID("TW")
	newOrderSingle.SetTargetCompID("ISLD")
	msg := newOrderSingle.ToMessage()

	utils.PrintGood(msg.String())
	senderr := quickfix.Send(msg)

	if senderr != nil {

		message := fmt.Sprintf("Failure sending nos message: %s", senderr)
		utils.PrintBad(message)
	}
	return nil

}
