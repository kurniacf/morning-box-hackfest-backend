package config

import (
	"os"

	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/snap"
)

var MidtransClient snap.Client

func InitMidtrans() {
	midtrans.ServerKey = os.Getenv("MIDTRANS_SERVER_KEY")
	midtrans.ClientKey = os.Getenv("MIDTRANS_CLIENT_KEY")
	midtrans.Environment = midtrans.Sandbox // Ganti dengan midtrans.Production untuk environment produksi

	MidtransClient = snap.Client{}
	MidtransClient.New(midtrans.ServerKey, midtrans.Environment)
}
