package newordersingle

import (
	"time"

	message "github.com/adrianguyareach/FIXlib/lib/message"
	"github.com/quickfixgo/enum"
	"github.com/quickfixgo/field"
	fix44nos "github.com/quickfixgo/fix44/newordersingle"
	"github.com/quickfixgo/quickfix"
)

func Fix44newordersingle(appSettings *quickfix.Settings) {

	var orderType field.OrdTypeField

	orderType.FIXString = quickfix.FIXString("Market")
	newOrderSingle := fix44nos.New(
		field.NewClOrdID("12345"),
		field.NewSide(enum.Side_BUY),
		field.NewTransactTime(time.Now()),
		orderType,
	)

	m := message.Message{
		Content: newOrderSingle.ToMessage(),
	}

	m.SendMessage(appSettings)

}
