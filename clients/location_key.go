package clients

var LocationKey = map[string]string{
	"1":  "Aceh",
	"2":  "Bali",
	"3":  "BangkaBelitung",
	"4":  "Banten",
	"5":  "Bengkulu",
	"6":  "DIYogyakarta",
	"7":  "DKIJakarta",
	"8":  "Gorontalo",
	"9":  "Jambi",
	"10": "JawaBarat",
	"11": "JawaTengah",
	"12": "JawaTimur",
	"13": "KalimantanBarat",
	"14": "KalimantanSelatan",
	"15": "KalimantanTengah",
	"16": "KalimantanTimur",
	"17": "KalimantanUtara",
	"18": "KepulauanRiau",
	"19": "Lampung",
	"20": "Maluku",
	"21": "MalukuUtara",
	"22": "NusaTenggaraBarat",
	"23": "NusaTenggaraTimur",
	"24": "Papua",
	"25": "PapuaBarat",
	"26": "Riau",
	"27": "SulawesiBarat",
	"28": "SulawesiSelatan",
	"29": "SulawesiTengah",
	"30": "SulawesiTenggara",
	"31": "SulawesiUtara",
	"32": "SumateraBarat",
	"33": "SumateraSelatan",
	"34": "SumateraUtara",
	"35": "Indonesia",
}

const BMKG_URL = "https://data.bmkg.go.id/"
const path = BMKG_URL + "/DataMKG/MEWS/DigitalForecast/DigitalForecast-"

func GetPath(key string) string {
	return path + LocationKey[key] + ".xml"
}
