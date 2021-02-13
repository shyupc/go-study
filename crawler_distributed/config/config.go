package config

const (
	// Parser names
	ParseCityList = "ParseCityList"
	ParseCity     = "ParseCity"
	ParseProfile  = "ParseProfile"
	NilParser     = "NilParser"

	// ElasticSearch index
	ElasticIndex = "dating_profile"

	// RPC Endpoints
	ItemSaverRpc    = "ItemSaverService.Save"
	CrawlServiceRpc = "CrawlService.Process"

	// Rate Limiting
	Qps = 20
)
