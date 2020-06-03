package config

// Config struct def
type Config struct {
	Database        string
	CarColl         string
	CarMakeColl     string
	CarModelColl    string
	CarBodyTrimColl string
	CarStyleColl    string
	ProvinceColl    string
}

// New returns a new config
func New() Config {
	return Config{
		Database:        "car",
		CarColl:         "car",
		CarMakeColl:     "masterCarMake",
		CarModelColl:    "masterCarModel",
		CarBodyTrimColl: "masterCarBodyTrim",
		CarStyleColl:    "masterCarStyle",
		ProvinceColl:    "masterProvince",
	}
}
