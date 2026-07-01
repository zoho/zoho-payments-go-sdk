package internal

const (
	SDKName = "zoho-payments-go-sdk"
	Version = "1.0.0"
)

func UserAgent() string { return SDKName + "/" + Version }
