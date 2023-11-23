package newordersingle

import (
	"fmt"
	"time"

	message "github.com/adrianguyareach/FIXlib/lib/message"
	"github.com/adrianguyareach/FIXlib/lib/utils"
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

	nos := message.Message{
		Content: newOrderSingle.ToMessage(),
	}

	newmsg := func(constructor message.MessageConstructor) *quickfix.Message {
		return constructor.ConstructMessage(appSettings)
	}(nos)

	for {
		senderr := quickfix.Send(newmsg)
		time.Sleep(time.Second)

		if senderr != nil {

			message := fmt.Sprintf("Failure sending nos message: %s", senderr)
			utils.PrintBad(message)
		}

	}

}
