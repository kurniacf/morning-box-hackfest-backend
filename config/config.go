package config

import (
	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/snap"
)

var MidtransClient snap.Client

func InitMidtrans() {
	midtrans.ServerKey = "SB-Mid-server-VoVtdr2Wycry4xHajdm5EX1d"
	midtrans.ClientKey = "SB-Mid-client-B2g-59qPVMbJnlG9"
	midtrans.Environment = midtrans.Sandbox // Ganti dengan midtrans.Production untuk environment produksi

	MidtransClient = snap.Client{}
	MidtransClient.New(midtrans.ServerKey, midtrans.Environment)
}
