// +build testsa

package spamc

import (
	"context"
	"os"
	"strings"
	"testing"
)

var addr = os.Getenv("SPAMC_SA_ADDRESS") + ":783"

func TestSAPing(t *testing.T) {
	client := New(addr, 0)
	err := client.Ping(context.Background())
	if err != nil {
		t.Fatal(err)
	}
}

func TestSATell(t *testing.T) {
	client := New(addr, 0)
	message := strings.NewReader("Subject: Hello, world!\r\n\r\nTest message.\r\n")
	r, err := client.Tell(context.Background(), message, Header{
		"Message-class": "spam",
		"Set":           "remote",
	})
	if err != nil {
		t.Fatal(err)
	}
	if r == nil {
		t.Fatal("r is nil")
	}
}

func TestSANoTrailingNewline(t *testing.T) {
	client := New(addr, 0)

	r, err := client.Check(context.Background(), strings.NewReader("woot"), nil)
	if err != nil {
		t.Fatal(err)
	}
	if r == nil {
		t.Fatal("r is nil")
	}

	r, err = client.Check(context.Background(), strings.NewReader("Subject: woot\r\n\r\nwoot"), nil)
	if err != nil {
		t.Fatal(err)
	}
	if r == nil {
		t.Fatal("r is nil")
	}
}

func TestSACheck(t *testing.T) {
	client := New(addr, 0)
	r, err := client.Check(context.Background(), strings.NewReader("\r\nPenis viagra\r\n"), nil)
	if err != nil {
		t.Fatal(err)
	}
	if r == nil {
		t.Fatal("r is nil")
	}
}

func TestSASymbols(t *testing.T) {
	client := New(addr, 0)
	r, err := client.Symbols(context.Background(), strings.NewReader(""+
		"Date: now\r\n"+
		"From: a@example.com\r\n"+
		"Subject: Hello\r\n"+
		"Message-ID: <serverfoo2131645635@example.com>\r\n"+
		"\r\n\r\nthe body\r\n"+
		""), nil)
	if err != nil {
		t.Fatal(err)
	}
	if r == nil {
		t.Fatal("r is nil")
	}
}

func TestSAReport(t *testing.T) {
	client := New(addr, 0)
	r, err := client.Report(context.Background(), strings.NewReader(""+
		"Date: now\r\n"+
		"From: a@example.com\r\n"+
		"Subject: Hello\r\n"+
		"Message-ID: <serverfoo2131645635@example.com>\r\n"+
		"\r\n\r\nthe body"+
		""), nil)
	if err != nil {
		t.Fatal(err)
	}
	if r == nil {
		t.Fatal("r is nil")
	}

	//fmt.Printf("%#v\n", r)
}

func TestSAProcess(t *testing.T) {
	client := New(addr, 0)
	r, err := client.Process(context.Background(), strings.NewReader(""+
		"Date: now\r\n"+
		"From: a@example.com\r\n"+
		"Subject: Hello\r\n"+
		"Message-ID: <serverfoo2131645635@example.com>\r\n"+
		"\r\n\r\nthe body"+
		""), nil)
	if err != nil {
		t.Fatal(err)
	}
	if r == nil {
		t.Fatal("r is nil")
	}

	// Added headers
	//X-Spam-Checker-Version: SpamAssassin 3.4.1 (2015-04-28) on 6c9bb381daaf
	//X-Spam-Level: *
	//X-Spam-Status: No, score=1.6 required=5.0 tests=INVALID_DATE,MISSING_HEADERS,
	//        NO_RECEIVED,NO_RELAYS autolearn=no autolearn_force=no version=3.4.1
}
