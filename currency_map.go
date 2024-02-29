package isocode

var currencies = map[Currency]currencyDetails{
	AED: {Numeric: 784, Digits: 2, Name: "United Arab Emirates"},
	AFN: {Numeric: 971, Digits: 2, Name: "Afghan afghani"},
	ALL: {Numeric: 8, Digits: 2, Name: "Albanian lek"},
	AMD: {Numeric: 51, Digits: 2, Name: "Armenian dram"},
	ANG: {Numeric: 532, Digits: 2, Name: "Netherlands Antillean guilder"},
	AOA: {Numeric: 973, Digits: 2, Name: "Angolan kwanza"},
	ARS: {Numeric: 32, Digits: 2, Name: "Argentine peso"},
	AUD: {Numeric: 36, Digits: 2, Name: "Australian dollar"},
	AWG: {Numeric: 533, Digits: 2, Name: "Aruban florin"},
	AZN: {Numeric: 944, Digits: 2, Name: "Azerbaijani manat"},
	BAM: {Numeric: 977, Digits: 2, Name: "Bosnia and Herzegovina convertible mark"},
	BBD: {Numeric: 52, Digits: 2, Name: "Barbados dollar"},
	BDT: {Numeric: 50, Digits: 2, Name: "Bangladeshi taka"},
	BGN: {Numeric: 975, Digits: 2, Name: "Bulgarian lev"},
	BHD: {Numeric: 48, Digits: 3, Name: "Bahraini dinar"},
	BIF: {Numeric: 108, Digits: 0, Name: "Burundian franc"},
	BMD: {Numeric: 60, Digits: 2, Name: "Bermudian dollar"},
	BND: {Numeric: 96, Digits: 2, Name: "Brunei dollar"},
	BOB: {Numeric: 68, Digits: 2, Name: "Boliviano"},
	BOV: {Numeric: 984, Digits: 2, Name: "Bolivian Mvdol (funds code)"},
	BRL: {Numeric: 986, Digits: 2, Name: "Brazilian real"},
	BSD: {Numeric: 44, Digits: 2, Name: "Bahamian dollar"},
	BTN: {Numeric: 64, Digits: 2, Name: "Bhutanese ngultrum"},
	BWP: {Numeric: 72, Digits: 2, Name: "Botswana pula"},
	BYN: {Numeric: 933, Digits: 2, Name: "Belarusian ruble"},
	BZD: {Numeric: 84, Digits: 2, Name: "Belize dollar"},
	CAD: {Numeric: 124, Digits: 2, Name: "Canadian dollar"},
	CDF: {Numeric: 976, Digits: 2, Name: "Congolese franc"},
	CHE: {Numeric: 947, Digits: 2, Name: "WIR euro (complementary currency)"},
	CHF: {Numeric: 756, Digits: 2, Name: "Swiss franc"},
	CHW: {Numeric: 948, Digits: 2, Name: "WIR franc (complementary currency)"},
	CLF: {Numeric: 990, Digits: 4, Name: "Unidad de Fomento (funds code)"},
	CLP: {Numeric: 152, Digits: 0, Name: "Chilean peso"},
	COP: {Numeric: 170, Digits: 2, Name: "Colombian peso"},
	COU: {Numeric: 970, Digits: 2, Name: "Unidad de Valor Real (UVR)"},
	CRC: {Numeric: 188, Digits: 2, Name: "Costa Rican"},
	CUC: {Numeric: 931, Digits: 2, Name: "Cuban convertible"},
	CUP: {Numeric: 192, Digits: 2, Name: "Cuban peso"},
	CVE: {Numeric: 132, Digits: 2, Name: "Cape Verdean escudo"},
	CZK: {Numeric: 203, Digits: 2, Name: "Czech koruna"},
	DJF: {Numeric: 262, Digits: 0, Name: "Djiboutian franc"},
	DKK: {Numeric: 208, Digits: 2, Name: "Danish krone"},
	DOP: {Numeric: 214, Digits: 2, Name: "Dominican peso"},
	DZD: {Numeric: 12, Digits: 2, Name: "Algerian dinar"},
	EGP: {Numeric: 818, Digits: 2, Name: "Egyptian pound"},
	ERN: {Numeric: 232, Digits: 2, Name: "Eritrean nakfa"},
	ETB: {Numeric: 230, Digits: 2, Name: "Ethiopian birr"},
	EUR: {Numeric: 978, Digits: 2, Name: "Euro"},
	FJD: {Numeric: 242, Digits: 2, Name: "Fiji dollar"},
	FKP: {Numeric: 238, Digits: 2, Name: "Falkland Islands pound"},
	GBP: {Numeric: 826, Digits: 2, Name: "Pound sterling"},
	GEL: {Numeric: 981, Digits: 2, Name: "Georgian lari"},
	GHS: {Numeric: 936, Digits: 2, Name: "Ghanaian cedi"},
	GIP: {Numeric: 292, Digits: 2, Name: "Gibraltar pound"},
	GMD: {Numeric: 270, Digits: 2, Name: "Gambian dalasi"},
	GNF: {Numeric: 324, Digits: 0, Name: "Guinean franc"},
	GTQ: {Numeric: 320, Digits: 2, Name: "Guatemalan quetzal"},
	GYD: {Numeric: 328, Digits: 2, Name: "Guyanese dollar"},
	HKD: {Numeric: 344, Digits: 2, Name: "Hong Kong dollar"},
	HNL: {Numeric: 340, Digits: 2, Name: "Honduran lempira"},
	HTG: {Numeric: 332, Digits: 2, Name: "Haitian gourde"},
	HUF: {Numeric: 348, Digits: 2, Name: "Hungarian forint"},
	IDR: {Numeric: 360, Digits: 2, Name: "Indonesian rupiah"},
	ILS: {Numeric: 376, Digits: 2, Name: "Israeli new shekel"},
	INR: {Numeric: 356, Digits: 2, Name: "Indian rupee"},
	IQD: {Numeric: 368, Digits: 3, Name: "Iraqi dinar"},
	IRR: {Numeric: 364, Digits: 2, Name: "Iranian rial"},
	ISK: {Numeric: 352, Digits: 0, Name: "Icelandic króna (plural: krónur)"},
	JMD: {Numeric: 388, Digits: 2, Name: "Jamaican dollar"},
	JOD: {Numeric: 400, Digits: 3, Name: "Jordanian dinar"},
	JPY: {Numeric: 392, Digits: 0, Name: "Japanese yen"},
	KES: {Numeric: 404, Digits: 2, Name: "Kenyan shilling"},
	KGS: {Numeric: 417, Digits: 2, Name: "Kyrgyzstani som"},
	KHR: {Numeric: 116, Digits: 2, Name: "Cambodian riel"},
	KMF: {Numeric: 174, Digits: 0, Name: "Comoro franc"},
	KPW: {Numeric: 408, Digits: 2, Name: "North Korean won"},
	KRW: {Numeric: 410, Digits: 0, Name: "South Korean won"},
	KWD: {Numeric: 414, Digits: 3, Name: "Kuwaiti dinar"},
	KYD: {Numeric: 136, Digits: 2, Name: "Cayman Islands dollar"},
	KZT: {Numeric: 398, Digits: 2, Name: "Kazakhstani tenge"},
	LAK: {Numeric: 418, Digits: 2, Name: "Lao kip"},
	LBP: {Numeric: 422, Digits: 2, Name: "Lebanese pound"},
	LKR: {Numeric: 144, Digits: 2, Name: "Sri Lankan rupee"},
	LRD: {Numeric: 430, Digits: 2, Name: "Liberian dollar"},
	LSL: {Numeric: 426, Digits: 2, Name: "Lesotho loti"},
	LYD: {Numeric: 434, Digits: 3, Name: "Libyan dinar"},
	MAD: {Numeric: 504, Digits: 2, Name: "Moroccan dirham"},
	MDL: {Numeric: 498, Digits: 2, Name: "Moldovan leu"},
	MGA: {Numeric: 969, Digits: 2, Name: "Malagasy ariary"},
	MKD: {Numeric: 807, Digits: 2, Name: "Macedonian denar"},
	MMK: {Numeric: 104, Digits: 2, Name: "Myanmar kyat"},
	MNT: {Numeric: 496, Digits: 2, Name: "Mongolian tögrög"},
	MOP: {Numeric: 446, Digits: 2, Name: "Macanese pataca"},
	MRU: {Numeric: 929, Digits: 2, Name: "Mauritanian ouguiya"},
	MUR: {Numeric: 480, Digits: 2, Name: "Mauritian rupee"},
	MVR: {Numeric: 462, Digits: 2, Name: "Maldivian rufiyaa"},
	MWK: {Numeric: 454, Digits: 2, Name: "Malawian kwacha"},
	MXN: {Numeric: 484, Digits: 2, Name: "Mexican peso"},
	MXV: {Numeric: 979, Digits: 2, Name: "Mexican Unidad de Inversion (UDI) (funds code)"},
	MYR: {Numeric: 458, Digits: 2, Name: "Malaysian ringgit"},
	MZN: {Numeric: 943, Digits: 2, Name: "Mozambican metical"},
	NAD: {Numeric: 516, Digits: 2, Name: "Namibian dollar"},
	NGN: {Numeric: 566, Digits: 2, Name: "Nigerian naira"},
	NIO: {Numeric: 558, Digits: 2, Name: "Nicaraguan córdoba"},
	NOK: {Numeric: 578, Digits: 2, Name: "Norwegian krone"},
	NPR: {Numeric: 524, Digits: 2, Name: "Nepalese rupee"},
	NZD: {Numeric: 554, Digits: 2, Name: "New Zealand dollar"},
	OMR: {Numeric: 512, Digits: 3, Name: "Omani rial"},
	PAB: {Numeric: 590, Digits: 2, Name: "Panamanian balboa"},
	PEN: {Numeric: 604, Digits: 2, Name: "Peruvian sol"},
	PGK: {Numeric: 598, Digits: 2, Name: "Papua New"},
	PHP: {Numeric: 608, Digits: 2, Name: "Philippine peso"},
	PKR: {Numeric: 586, Digits: 2, Name: "Pakistani rupee"},
	PLN: {Numeric: 985, Digits: 2, Name: "Polish złoty"},
	PYG: {Numeric: 600, Digits: 0, Name: "Paraguayan guaraní"},
	QAR: {Numeric: 634, Digits: 2, Name: "Qatari riyal"},
	RON: {Numeric: 946, Digits: 2, Name: "Romanian leu"},
	RSD: {Numeric: 941, Digits: 2, Name: "Serbian dinar"},
	CNY: {Numeric: 156, Digits: 2, Name: "Renminbi"},
	RUB: {Numeric: 643, Digits: 2, Name: "Russian ruble"},
	RWF: {Numeric: 646, Digits: 0, Name: "Rwandan franc"},
	SAR: {Numeric: 682, Digits: 2, Name: "Saudi riyal"},
	SBD: {Numeric: 90, Digits: 2, Name: "Solomon Islands dollar"},
	SCR: {Numeric: 690, Digits: 2, Name: "Seychelles rupee"},
	SDG: {Numeric: 938, Digits: 2, Name: "Sudanese pound"},
	SEK: {Numeric: 752, Digits: 2, Name: "Swedish krona (plural: kronor)"},
	SGD: {Numeric: 702, Digits: 2, Name: "Singapore dollar"},
	SHP: {Numeric: 654, Digits: 2, Name: "Saint Helena pound"},
	SLE: {Numeric: 925, Digits: 2, Name: "Sierra Leonean leone (new leone)"},
	SLL: {Numeric: 694, Digits: 2, Name: "Sierra Leonean leone (old leone)"},
	SOS: {Numeric: 706, Digits: 2, Name: "Somali shilling"},
	SRD: {Numeric: 968, Digits: 2, Name: "Surinamese dollar"},
	SSP: {Numeric: 728, Digits: 2, Name: "South Sudanese pound"},
	STN: {Numeric: 930, Digits: 2, Name: "São Tomé and Príncipe dobra"},
	SVC: {Numeric: 222, Digits: 2, Name: "Salvadoran colón"},
	SYP: {Numeric: 760, Digits: 2, Name: "Syrian pound"},
	SZL: {Numeric: 748, Digits: 2, Name: "Swazi lilangeni"},
	THB: {Numeric: 764, Digits: 2, Name: "Thai baht"},
	TJS: {Numeric: 972, Digits: 2, Name: "Tajikistani somoni"},
	TMT: {Numeric: 934, Digits: 2, Name: "Turkmenistan manat"},
	TND: {Numeric: 788, Digits: 3, Name: "Tunisian dinar"},
	TOP: {Numeric: 776, Digits: 2, Name: "Tongan pa'anga"},
	TRY: {Numeric: 949, Digits: 2, Name: "Turkish lira"},
	TTD: {Numeric: 780, Digits: 2, Name: "Trinidad and"},
	TWD: {Numeric: 901, Digits: 2, Name: "New Taiwan dollar"},
	TZS: {Numeric: 834, Digits: 2, Name: "Tanzanian shilling"},
	UAH: {Numeric: 980, Digits: 2, Name: "Ukrainian hryvnia"},
	UGX: {Numeric: 800, Digits: 0, Name: "Ugandan shilling"},
	USD: {Numeric: 840, Digits: 2, Name: "United States dollar"},
	USN: {Numeric: 997, Digits: 2, Name: "United States dollar"},
	UYI: {Numeric: 940, Digits: 0, Name: "Uruguay Peso"},
	UYU: {Numeric: 858, Digits: 2, Name: "Uruguayan peso"},
	UYW: {Numeric: 927, Digits: 4, Name: "Unidad previsional"},
	UZS: {Numeric: 860, Digits: 2, Name: "Uzbekistan sum"},
	VED: {Numeric: 926, Digits: 2, Name: "Venezuelan digital bolívar"},
	VES: {Numeric: 928, Digits: 2, Name: "Venezuelan sovereign bolívar"},
	VND: {Numeric: 704, Digits: 0, Name: "Vietnamese đồng"},
	VUV: {Numeric: 548, Digits: 0, Name: "Vanuatu vatu"},
	WST: {Numeric: 882, Digits: 2, Name: "Samoan tala"},
	XAF: {Numeric: 950, Digits: 0, Name: "CFA franc"},
	XAG: {Numeric: 961, Digits: -1, Name: "Silver (one troy ounce)"},
	XAU: {Numeric: 959, Digits: -1, Name: "Gold (one troy ounce)"},
	XBA: {Numeric: 955, Digits: -1, Name: "European Composite"},
	XBB: {Numeric: 956, Digits: -1, Name: "European Monetary"},
	XBC: {Numeric: 957, Digits: -1, Name: "European Unit"},
	XBD: {Numeric: 958, Digits: -1, Name: "European Unit"},
	XCD: {Numeric: 951, Digits: 2, Name: "East Caribbean dollar"},
	XDR: {Numeric: 960, Digits: -1, Name: "Special drawing rights"},
	XOF: {Numeric: 952, Digits: 0, Name: "CFA franc"},
	XPD: {Numeric: 964, Digits: -1, Name: "Palladium (one troy ounce)"},
	XPF: {Numeric: 953, Digits: 0, Name: "CFP franc"},
	XPT: {Numeric: 962, Digits: -1, Name: "Platinum (one troy ounce)"},
	XSU: {Numeric: 994, Digits: -1, Name: "SUCRE"},
	XTS: {Numeric: 963, Digits: -1, Name: "Code reserved for testing"},
	XUA: {Numeric: 965, Digits: -1, Name: "ADB Unit of Account"},
	XXX: {Numeric: 999, Digits: -1, Name: "No currency"},
	YER: {Numeric: 886, Digits: 2, Name: "Yemeni rial"},
	ZAR: {Numeric: 710, Digits: 2, Name: "South African rand"},
	ZMW: {Numeric: 967, Digits: 2, Name: "Zambian kwacha"},
	ZWL: {Numeric: 932, Digits: 2, Name: "Zimbabwean dollar (fifth)"},
}
