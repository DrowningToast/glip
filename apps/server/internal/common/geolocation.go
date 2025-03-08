package common

type Region struct {
	Name string `json:"name"`
	Id   string `json:"id"`
	Key  string `json:"code"`
}

var (
	RegionNA Region = Region{
		Name: "North America",
		Id:   "NA",
		Key:  "zoneNA",
	}
	RegionEU Region = Region{
		Name: "Europe",
		Id:   "EU",
		Key:  "zoneEU",
	}
	RegionAP Region = Region{
		Name: "Asia Pacific",
		Id:   "AP",
		Key:  "zoneAP",
	}
)

type Country struct {
	Name   string `json:"name"`
	Region Region `json:"region"`
}

var (
	CountryUSA Country = Country{
		Name:   "United States of America",
		Region: RegionNA,
	}
	CountryFrance Country = Country{
		Name:   "France",
		Region: RegionEU,
	}
	CountryItaly Country = Country{
		Name:   "Italy",
		Region: RegionEU,
	}
	CountrySwitzerland Country = Country{
		Name:   "Switzerland",
		Region: RegionEU,
	}
	CountrySingapore Country = Country{
		Name:   "Singapore",
		Region: RegionAP,
	}
	CountryChina Country = Country{
		Name:   "China",
		Region: RegionAP,
	}
	CountryJapan Country = Country{
		Name:   "Japan",
		Region: RegionAP,
	}
)
