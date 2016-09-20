package imago

// imago
var (
	G       = &iGenerate{}
	Convert = &iConvert{}
	Crypto  = &iCrypto{}
	File    = &iFile{}
)

type (
	iGenerate struct{}
	iConvert  struct{}
	iCrypto   struct{}
	iFile     struct{}
)
