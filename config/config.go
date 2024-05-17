package config

type Config struct {
	StripeSecretKey string
}

func LoadConfig() Config {
	return Config{
		StripeSecretKey: "sk_test_51OoozPJ279ibsqYKmw3sRoNbwYmdl2Yb4XHuIQoTpDr8cok2jdbDtHXCNllPvm8giZltURZPr44hpxiOCM4hmsvM00bCK09nPg",
	}
}
