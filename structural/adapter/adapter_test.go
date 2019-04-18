package adapter

import "testing"

func TestAdapter(t *testing.T) {
	msg := "Hi, Septem"

	adapter := PrinterAdapter{
		OldPrinter: &MyLegacyPrinter{},
		Msg:        msg,
	}

	res := adapter.PrintStored()

	if res != "Legacy Printer: Adapter: Hi, Septem\n" {
		t.Errorf("Message didn't match: %s\n", res)
	}

	adapter = PrinterAdapter{
		OldPrinter: nil,
		Msg:        msg,
	}

	res = adapter.PrintStored()

	if res != "Hi, Septem" {
		t.Errorf("Message didn't match: %s\n", res)
	}
}
