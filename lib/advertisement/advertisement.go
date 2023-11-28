package advertisement

import (
	message "github.com/adrianguyareach/FIXlib/lib/message"
	"github.com/adrianguyareach/FIXlib/lib/utils"
	"github.com/quickfixgo/enum"
	"github.com/quickfixgo/field"
	adv "github.com/quickfixgo/fix44/advertisement"
	"github.com/quickfixgo/quickfix"
	"github.com/shopspring/decimal"
)

func FIX44Advertisement(appSettings *quickfix.Settings) {

	// var orderType field.OrdTypeField

	// orderType.FIXString = quickfix.FIXString("Market")
	// newOrderSingle := fix44nos.New(
	// 	field.NewClOrdID("12345"),
	// 	field.NewSide(enum.Side_BUY),
	// 	field.NewTransactTime(time.Now()),
	// 	orderType,
	// )

	// refID := field.NewAdvRefID("12345678")

	n, err := decimal.NewFromString("100.0")
	if err != nil {
		utils.PrintBad("Failure extracting decimal")
	}

	var scale int32 = 3

	advertisement := adv.New(
		field.NewAdvId("adver123"),
		field.NewAdvTransType(enum.AdvTransType_NEW), //Values can be NEW, CANCEL OR REPLACE
		field.NewAdvSide(enum.AdvSide_BUY),
		field.NewQuantity(n, scale),
	)
	m := message.Message{
		Content: advertisement.ToMessage(),
	}
	m.SendMessage(appSettings)
}
