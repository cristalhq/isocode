package currency

// ISO 4217 alphabetic codes.
const (
	AED Code = "AED"
	AFN Code = "AFN"
	ALL Code = "ALL"
	AMD Code = "AMD"
	AOA Code = "AOA"
	ARS Code = "ARS"
	AUD Code = "AUD"
	AWG Code = "AWG"
	AZN Code = "AZN"
	BAM Code = "BAM"
	BBD Code = "BBD"
	BDT Code = "BDT"
	BHD Code = "BHD"
	BIF Code = "BIF"
	BMD Code = "BMD"
	BND Code = "BND"
	BOB Code = "BOB"
	BOV Code = "BOV"
	BRL Code = "BRL"
	BSD Code = "BSD"
	BTN Code = "BTN"
	BWP Code = "BWP"
	BYN Code = "BYN"
	BZD Code = "BZD"
	CAD Code = "CAD"
	CDF Code = "CDF"
	CHE Code = "CHE"
	CHF Code = "CHF"
	CHW Code = "CHW"
	CLF Code = "CLF"
	CLP Code = "CLP"
	CNY Code = "CNY"
	COP Code = "COP"
	COU Code = "COU"
	CRC Code = "CRC"
	CUP Code = "CUP"
	CVE Code = "CVE"
	CZK Code = "CZK"
	DJF Code = "DJF"
	DKK Code = "DKK"
	DOP Code = "DOP"
	DZD Code = "DZD"
	EGP Code = "EGP"
	ERN Code = "ERN"
	ETB Code = "ETB"
	EUR Code = "EUR"
	FJD Code = "FJD"
	FKP Code = "FKP"
	GBP Code = "GBP"
	GEL Code = "GEL"
	GHS Code = "GHS"
	GIP Code = "GIP"
	GMD Code = "GMD"
	GNF Code = "GNF"
	GTQ Code = "GTQ"
	GYD Code = "GYD"
	HKD Code = "HKD"
	HNL Code = "HNL"
	HTG Code = "HTG"
	HUF Code = "HUF"
	IDR Code = "IDR"
	ILS Code = "ILS"
	INR Code = "INR"
	IQD Code = "IQD"
	IRR Code = "IRR"
	ISK Code = "ISK"
	JMD Code = "JMD"
	JOD Code = "JOD"
	JPY Code = "JPY"
	KES Code = "KES"
	KGS Code = "KGS"
	KHR Code = "KHR"
	KMF Code = "KMF"
	KPW Code = "KPW"
	KRW Code = "KRW"
	KWD Code = "KWD"
	KYD Code = "KYD"
	KZT Code = "KZT"
	LAK Code = "LAK"
	LBP Code = "LBP"
	LKR Code = "LKR"
	LRD Code = "LRD"
	LSL Code = "LSL"
	LYD Code = "LYD"
	MAD Code = "MAD"
	MDL Code = "MDL"
	MGA Code = "MGA"
	MKD Code = "MKD"
	MMK Code = "MMK"
	MNT Code = "MNT"
	MOP Code = "MOP"
	MRU Code = "MRU"
	MUR Code = "MUR"
	MVR Code = "MVR"
	MWK Code = "MWK"
	MXN Code = "MXN"
	MXV Code = "MXV"
	MYR Code = "MYR"
	MZN Code = "MZN"
	NAD Code = "NAD"
	NGN Code = "NGN"
	NIO Code = "NIO"
	NOK Code = "NOK"
	NPR Code = "NPR"
	NZD Code = "NZD"
	OMR Code = "OMR"
	PAB Code = "PAB"
	PEN Code = "PEN"
	PGK Code = "PGK"
	PHP Code = "PHP"
	PKR Code = "PKR"
	PLN Code = "PLN"
	PYG Code = "PYG"
	QAR Code = "QAR"
	RON Code = "RON"
	RSD Code = "RSD"
	RUB Code = "RUB"
	RWF Code = "RWF"
	SAR Code = "SAR"
	SBD Code = "SBD"
	SCR Code = "SCR"
	SDG Code = "SDG"
	SEK Code = "SEK"
	SGD Code = "SGD"
	SHP Code = "SHP"
	SLE Code = "SLE"
	SOS Code = "SOS"
	SRD Code = "SRD"
	SSP Code = "SSP"
	STN Code = "STN"
	SVC Code = "SVC"
	SYP Code = "SYP"
	SZL Code = "SZL"
	THB Code = "THB"
	TJS Code = "TJS"
	TMT Code = "TMT"
	TND Code = "TND"
	TOP Code = "TOP"
	TRY Code = "TRY"
	TTD Code = "TTD"
	TWD Code = "TWD"
	TZS Code = "TZS"
	UAH Code = "UAH"
	UGX Code = "UGX"
	USD Code = "USD"
	USN Code = "USN"
	UYI Code = "UYI"
	UYU Code = "UYU"
	UYW Code = "UYW"
	UZS Code = "UZS"
	VED Code = "VED"
	VES Code = "VES"
	VND Code = "VND"
	VUV Code = "VUV"
	WST Code = "WST"
	XAD Code = "XAD"
	XAF Code = "XAF"
	XAG Code = "XAG"
	XAU Code = "XAU"
	XBA Code = "XBA"
	XBB Code = "XBB"
	XBC Code = "XBC"
	XBD Code = "XBD"
	XCD Code = "XCD"
	XCG Code = "XCG"
	XDR Code = "XDR"
	XOF Code = "XOF"
	XPD Code = "XPD"
	XPF Code = "XPF"
	XPT Code = "XPT"
	XSU Code = "XSU"
	XTS Code = "XTS"
	XUA Code = "XUA"
	XXX Code = "XXX"
	YER Code = "YER"
	ZAR Code = "ZAR"
	ZMW Code = "ZMW"
	ZWG Code = "ZWG"
)

var table = map[Code]Info{
	AED: {Code: AED, Numeric: "784", Name: "UAE Dirham", MinorUnit: 2},
	AFN: {Code: AFN, Numeric: "971", Name: "Afghani", MinorUnit: 2},
	ALL: {Code: ALL, Numeric: "008", Name: "Lek", MinorUnit: 2},
	AMD: {Code: AMD, Numeric: "051", Name: "Armenian Dram", MinorUnit: 2},
	AOA: {Code: AOA, Numeric: "973", Name: "Kwanza", MinorUnit: 2},
	ARS: {Code: ARS, Numeric: "032", Name: "Argentine Peso", MinorUnit: 2},
	AUD: {Code: AUD, Numeric: "036", Name: "Australian Dollar", MinorUnit: 2},
	AWG: {Code: AWG, Numeric: "533", Name: "Aruban Florin", MinorUnit: 2},
	AZN: {Code: AZN, Numeric: "944", Name: "Azerbaijan Manat", MinorUnit: 2},
	BAM: {Code: BAM, Numeric: "977", Name: "Convertible Mark", MinorUnit: 2},
	BBD: {Code: BBD, Numeric: "052", Name: "Barbados Dollar", MinorUnit: 2},
	BDT: {Code: BDT, Numeric: "050", Name: "Taka", MinorUnit: 2},
	BHD: {Code: BHD, Numeric: "048", Name: "Bahraini Dinar", MinorUnit: 3},
	BIF: {Code: BIF, Numeric: "108", Name: "Burundi Franc", MinorUnit: 0},
	BMD: {Code: BMD, Numeric: "060", Name: "Bermudian Dollar", MinorUnit: 2},
	BND: {Code: BND, Numeric: "096", Name: "Brunei Dollar", MinorUnit: 2},
	BOB: {Code: BOB, Numeric: "068", Name: "Boliviano", MinorUnit: 2},
	BOV: {Code: BOV, Numeric: "984", Name: "Mvdol", MinorUnit: 2},
	BRL: {Code: BRL, Numeric: "986", Name: "Brazilian Real", MinorUnit: 2},
	BSD: {Code: BSD, Numeric: "044", Name: "Bahamian Dollar", MinorUnit: 2},
	BTN: {Code: BTN, Numeric: "064", Name: "Ngultrum", MinorUnit: 2},
	BWP: {Code: BWP, Numeric: "072", Name: "Pula", MinorUnit: 2},
	BYN: {Code: BYN, Numeric: "933", Name: "Belarusian Ruble", MinorUnit: 2},
	BZD: {Code: BZD, Numeric: "084", Name: "Belize Dollar", MinorUnit: 2},
	CAD: {Code: CAD, Numeric: "124", Name: "Canadian Dollar", MinorUnit: 2},
	CDF: {Code: CDF, Numeric: "976", Name: "Congolese Franc", MinorUnit: 2},
	CHE: {Code: CHE, Numeric: "947", Name: "WIR Euro", MinorUnit: 2},
	CHF: {Code: CHF, Numeric: "756", Name: "Swiss Franc", MinorUnit: 2},
	CHW: {Code: CHW, Numeric: "948", Name: "WIR Franc", MinorUnit: 2},
	CLF: {Code: CLF, Numeric: "990", Name: "Unidad de Fomento", MinorUnit: 4},
	CLP: {Code: CLP, Numeric: "152", Name: "Chilean Peso", MinorUnit: 0},
	CNY: {Code: CNY, Numeric: "156", Name: "Yuan Renminbi", MinorUnit: 2},
	COP: {Code: COP, Numeric: "170", Name: "Colombian Peso", MinorUnit: 2},
	COU: {Code: COU, Numeric: "970", Name: "Unidad de Valor Real", MinorUnit: 2},
	CRC: {Code: CRC, Numeric: "188", Name: "Costa Rican Colon", MinorUnit: 2},
	CUP: {Code: CUP, Numeric: "192", Name: "Cuban Peso", MinorUnit: 2},
	CVE: {Code: CVE, Numeric: "132", Name: "Cabo Verde Escudo", MinorUnit: 2},
	CZK: {Code: CZK, Numeric: "203", Name: "Czech Koruna", MinorUnit: 2},
	DJF: {Code: DJF, Numeric: "262", Name: "Djibouti Franc", MinorUnit: 0},
	DKK: {Code: DKK, Numeric: "208", Name: "Danish Krone", MinorUnit: 2},
	DOP: {Code: DOP, Numeric: "214", Name: "Dominican Peso", MinorUnit: 2},
	DZD: {Code: DZD, Numeric: "012", Name: "Algerian Dinar", MinorUnit: 2},
	EGP: {Code: EGP, Numeric: "818", Name: "Egyptian Pound", MinorUnit: 2},
	ERN: {Code: ERN, Numeric: "232", Name: "Nakfa", MinorUnit: 2},
	ETB: {Code: ETB, Numeric: "230", Name: "Ethiopian Birr", MinorUnit: 2},
	EUR: {Code: EUR, Numeric: "978", Name: "Euro", MinorUnit: 2},
	FJD: {Code: FJD, Numeric: "242", Name: "Fiji Dollar", MinorUnit: 2},
	FKP: {Code: FKP, Numeric: "238", Name: "Falkland Islands Pound", MinorUnit: 2},
	GBP: {Code: GBP, Numeric: "826", Name: "Pound Sterling", MinorUnit: 2},
	GEL: {Code: GEL, Numeric: "981", Name: "Lari", MinorUnit: 2},
	GHS: {Code: GHS, Numeric: "936", Name: "Ghana Cedi", MinorUnit: 2},
	GIP: {Code: GIP, Numeric: "292", Name: "Gibraltar Pound", MinorUnit: 2},
	GMD: {Code: GMD, Numeric: "270", Name: "Dalasi", MinorUnit: 2},
	GNF: {Code: GNF, Numeric: "324", Name: "Guinean Franc", MinorUnit: 0},
	GTQ: {Code: GTQ, Numeric: "320", Name: "Quetzal", MinorUnit: 2},
	GYD: {Code: GYD, Numeric: "328", Name: "Guyana Dollar", MinorUnit: 2},
	HKD: {Code: HKD, Numeric: "344", Name: "Hong Kong Dollar", MinorUnit: 2},
	HNL: {Code: HNL, Numeric: "340", Name: "Lempira", MinorUnit: 2},
	HTG: {Code: HTG, Numeric: "332", Name: "Gourde", MinorUnit: 2},
	HUF: {Code: HUF, Numeric: "348", Name: "Forint", MinorUnit: 2},
	IDR: {Code: IDR, Numeric: "360", Name: "Rupiah", MinorUnit: 2},
	ILS: {Code: ILS, Numeric: "376", Name: "New Israeli Sheqel", MinorUnit: 2},
	INR: {Code: INR, Numeric: "356", Name: "Indian Rupee", MinorUnit: 2},
	IQD: {Code: IQD, Numeric: "368", Name: "Iraqi Dinar", MinorUnit: 3},
	IRR: {Code: IRR, Numeric: "364", Name: "Iranian Rial", MinorUnit: 2},
	ISK: {Code: ISK, Numeric: "352", Name: "Iceland Krona", MinorUnit: 0},
	JMD: {Code: JMD, Numeric: "388", Name: "Jamaican Dollar", MinorUnit: 2},
	JOD: {Code: JOD, Numeric: "400", Name: "Jordanian Dinar", MinorUnit: 3},
	JPY: {Code: JPY, Numeric: "392", Name: "Yen", MinorUnit: 0},
	KES: {Code: KES, Numeric: "404", Name: "Kenyan Shilling", MinorUnit: 2},
	KGS: {Code: KGS, Numeric: "417", Name: "Som", MinorUnit: 2},
	KHR: {Code: KHR, Numeric: "116", Name: "Riel", MinorUnit: 2},
	KMF: {Code: KMF, Numeric: "174", Name: "Comorian Franc ", MinorUnit: 0},
	KPW: {Code: KPW, Numeric: "408", Name: "North Korean Won", MinorUnit: 2},
	KRW: {Code: KRW, Numeric: "410", Name: "Won", MinorUnit: 0},
	KWD: {Code: KWD, Numeric: "414", Name: "Kuwaiti Dinar", MinorUnit: 3},
	KYD: {Code: KYD, Numeric: "136", Name: "Cayman Islands Dollar", MinorUnit: 2},
	KZT: {Code: KZT, Numeric: "398", Name: "Tenge", MinorUnit: 2},
	LAK: {Code: LAK, Numeric: "418", Name: "Lao Kip", MinorUnit: 2},
	LBP: {Code: LBP, Numeric: "422", Name: "Lebanese Pound", MinorUnit: 2},
	LKR: {Code: LKR, Numeric: "144", Name: "Sri Lanka Rupee", MinorUnit: 2},
	LRD: {Code: LRD, Numeric: "430", Name: "Liberian Dollar", MinorUnit: 2},
	LSL: {Code: LSL, Numeric: "426", Name: "Loti", MinorUnit: 2},
	LYD: {Code: LYD, Numeric: "434", Name: "Libyan Dinar", MinorUnit: 3},
	MAD: {Code: MAD, Numeric: "504", Name: "Moroccan Dirham", MinorUnit: 2},
	MDL: {Code: MDL, Numeric: "498", Name: "Moldovan Leu", MinorUnit: 2},
	MGA: {Code: MGA, Numeric: "969", Name: "Malagasy Ariary", MinorUnit: 2},
	MKD: {Code: MKD, Numeric: "807", Name: "Denar", MinorUnit: 2},
	MMK: {Code: MMK, Numeric: "104", Name: "Kyat", MinorUnit: 2},
	MNT: {Code: MNT, Numeric: "496", Name: "Tugrik", MinorUnit: 2},
	MOP: {Code: MOP, Numeric: "446", Name: "Pataca", MinorUnit: 2},
	MRU: {Code: MRU, Numeric: "929", Name: "Ouguiya", MinorUnit: 2},
	MUR: {Code: MUR, Numeric: "480", Name: "Mauritius Rupee", MinorUnit: 2},
	MVR: {Code: MVR, Numeric: "462", Name: "Rufiyaa", MinorUnit: 2},
	MWK: {Code: MWK, Numeric: "454", Name: "Malawi Kwacha", MinorUnit: 2},
	MXN: {Code: MXN, Numeric: "484", Name: "Mexican Peso", MinorUnit: 2},
	MXV: {Code: MXV, Numeric: "979", Name: "Mexican Unidad de Inversion (UDI)", MinorUnit: 2},
	MYR: {Code: MYR, Numeric: "458", Name: "Malaysian Ringgit", MinorUnit: 2},
	MZN: {Code: MZN, Numeric: "943", Name: "Mozambique Metical", MinorUnit: 2},
	NAD: {Code: NAD, Numeric: "516", Name: "Namibia Dollar", MinorUnit: 2},
	NGN: {Code: NGN, Numeric: "566", Name: "Naira", MinorUnit: 2},
	NIO: {Code: NIO, Numeric: "558", Name: "Cordoba Oro", MinorUnit: 2},
	NOK: {Code: NOK, Numeric: "578", Name: "Norwegian Krone", MinorUnit: 2},
	NPR: {Code: NPR, Numeric: "524", Name: "Nepalese Rupee", MinorUnit: 2},
	NZD: {Code: NZD, Numeric: "554", Name: "New Zealand Dollar", MinorUnit: 2},
	OMR: {Code: OMR, Numeric: "512", Name: "Rial Omani", MinorUnit: 3},
	PAB: {Code: PAB, Numeric: "590", Name: "Balboa", MinorUnit: 2},
	PEN: {Code: PEN, Numeric: "604", Name: "Sol", MinorUnit: 2},
	PGK: {Code: PGK, Numeric: "598", Name: "Kina", MinorUnit: 2},
	PHP: {Code: PHP, Numeric: "608", Name: "Philippine Peso", MinorUnit: 2},
	PKR: {Code: PKR, Numeric: "586", Name: "Pakistan Rupee", MinorUnit: 2},
	PLN: {Code: PLN, Numeric: "985", Name: "Zloty", MinorUnit: 2},
	PYG: {Code: PYG, Numeric: "600", Name: "Guarani", MinorUnit: 0},
	QAR: {Code: QAR, Numeric: "634", Name: "Qatari Rial", MinorUnit: 2},
	RON: {Code: RON, Numeric: "946", Name: "Romanian Leu", MinorUnit: 2},
	RSD: {Code: RSD, Numeric: "941", Name: "Serbian Dinar", MinorUnit: 2},
	RUB: {Code: RUB, Numeric: "643", Name: "Russian Ruble", MinorUnit: 2},
	RWF: {Code: RWF, Numeric: "646", Name: "Rwanda Franc", MinorUnit: 0},
	SAR: {Code: SAR, Numeric: "682", Name: "Saudi Riyal", MinorUnit: 2},
	SBD: {Code: SBD, Numeric: "090", Name: "Solomon Islands Dollar", MinorUnit: 2},
	SCR: {Code: SCR, Numeric: "690", Name: "Seychelles Rupee", MinorUnit: 2},
	SDG: {Code: SDG, Numeric: "938", Name: "Sudanese Pound", MinorUnit: 2},
	SEK: {Code: SEK, Numeric: "752", Name: "Swedish Krona", MinorUnit: 2},
	SGD: {Code: SGD, Numeric: "702", Name: "Singapore Dollar", MinorUnit: 2},
	SHP: {Code: SHP, Numeric: "654", Name: "Saint Helena Pound", MinorUnit: 2},
	SLE: {Code: SLE, Numeric: "925", Name: "Leone", MinorUnit: 2},
	SOS: {Code: SOS, Numeric: "706", Name: "Somali Shilling", MinorUnit: 2},
	SRD: {Code: SRD, Numeric: "968", Name: "Surinam Dollar", MinorUnit: 2},
	SSP: {Code: SSP, Numeric: "728", Name: "South Sudanese Pound", MinorUnit: 2},
	STN: {Code: STN, Numeric: "930", Name: "Dobra", MinorUnit: 2},
	SVC: {Code: SVC, Numeric: "222", Name: "El Salvador Colon", MinorUnit: 2},
	SYP: {Code: SYP, Numeric: "760", Name: "Syrian Pound", MinorUnit: 2},
	SZL: {Code: SZL, Numeric: "748", Name: "Lilangeni", MinorUnit: 2},
	THB: {Code: THB, Numeric: "764", Name: "Baht", MinorUnit: 2},
	TJS: {Code: TJS, Numeric: "972", Name: "Somoni", MinorUnit: 2},
	TMT: {Code: TMT, Numeric: "934", Name: "Turkmenistan New Manat", MinorUnit: 2},
	TND: {Code: TND, Numeric: "788", Name: "Tunisian Dinar", MinorUnit: 3},
	TOP: {Code: TOP, Numeric: "776", Name: "Pa’anga", MinorUnit: 2},
	TRY: {Code: TRY, Numeric: "949", Name: "Turkish Lira", MinorUnit: 2},
	TTD: {Code: TTD, Numeric: "780", Name: "Trinidad and Tobago Dollar", MinorUnit: 2},
	TWD: {Code: TWD, Numeric: "901", Name: "New Taiwan Dollar", MinorUnit: 2},
	TZS: {Code: TZS, Numeric: "834", Name: "Tanzanian Shilling", MinorUnit: 2},
	UAH: {Code: UAH, Numeric: "980", Name: "Hryvnia", MinorUnit: 2},
	UGX: {Code: UGX, Numeric: "800", Name: "Uganda Shilling", MinorUnit: 0},
	USD: {Code: USD, Numeric: "840", Name: "US Dollar", MinorUnit: 2},
	USN: {Code: USN, Numeric: "997", Name: "US Dollar (Next day)", MinorUnit: 2},
	UYI: {Code: UYI, Numeric: "940", Name: "Uruguay Peso en Unidades Indexadas (UI)", MinorUnit: 0},
	UYU: {Code: UYU, Numeric: "858", Name: "Peso Uruguayo", MinorUnit: 2},
	UYW: {Code: UYW, Numeric: "927", Name: "Unidad Previsional", MinorUnit: 4},
	UZS: {Code: UZS, Numeric: "860", Name: "Uzbekistan Sum", MinorUnit: 2},
	VED: {Code: VED, Numeric: "926", Name: "Bolívar Soberano", MinorUnit: 2},
	VES: {Code: VES, Numeric: "928", Name: "Bolívar Soberano", MinorUnit: 2},
	VND: {Code: VND, Numeric: "704", Name: "Dong", MinorUnit: 0},
	VUV: {Code: VUV, Numeric: "548", Name: "Vatu", MinorUnit: 0},
	WST: {Code: WST, Numeric: "882", Name: "Tala", MinorUnit: 2},
	XAD: {Code: XAD, Numeric: "396", Name: "Arab Accounting Dinar", MinorUnit: 2},
	XAF: {Code: XAF, Numeric: "950", Name: "CFA Franc BEAC", MinorUnit: 0},
	XAG: {Code: XAG, Numeric: "961", Name: "Silver", MinorUnit: -1},
	XAU: {Code: XAU, Numeric: "959", Name: "Gold", MinorUnit: -1},
	XBA: {Code: XBA, Numeric: "955", Name: "Bond Markets Unit European Composite Unit (EURCO)", MinorUnit: -1},
	XBB: {Code: XBB, Numeric: "956", Name: "Bond Markets Unit European Monetary Unit (E.M.U.-6)", MinorUnit: -1},
	XBC: {Code: XBC, Numeric: "957", Name: "Bond Markets Unit European Unit of Account 9 (E.U.A.-9)", MinorUnit: -1},
	XBD: {Code: XBD, Numeric: "958", Name: "Bond Markets Unit European Unit of Account 17 (E.U.A.-17)", MinorUnit: -1},
	XCD: {Code: XCD, Numeric: "951", Name: "East Caribbean Dollar", MinorUnit: 2},
	XCG: {Code: XCG, Numeric: "532", Name: "Caribbean Guilder", MinorUnit: 2},
	XDR: {Code: XDR, Numeric: "960", Name: "SDR (Special Drawing Right)", MinorUnit: -1},
	XOF: {Code: XOF, Numeric: "952", Name: "CFA Franc BCEAO", MinorUnit: 0},
	XPD: {Code: XPD, Numeric: "964", Name: "Palladium", MinorUnit: -1},
	XPF: {Code: XPF, Numeric: "953", Name: "CFP Franc", MinorUnit: 0},
	XPT: {Code: XPT, Numeric: "962", Name: "Platinum", MinorUnit: -1},
	XSU: {Code: XSU, Numeric: "994", Name: "Sucre", MinorUnit: -1},
	XTS: {Code: XTS, Numeric: "963", Name: "Codes specifically reserved for testing purposes", MinorUnit: -1},
	XUA: {Code: XUA, Numeric: "965", Name: "ADB Unit of Account", MinorUnit: -1},
	XXX: {Code: XXX, Numeric: "999", Name: "The codes assigned for transactions where no currency is involved", MinorUnit: -1},
	YER: {Code: YER, Numeric: "886", Name: "Yemeni Rial", MinorUnit: 2},
	ZAR: {Code: ZAR, Numeric: "710", Name: "Rand", MinorUnit: 2},
	ZMW: {Code: ZMW, Numeric: "967", Name: "Zambian Kwacha", MinorUnit: 2},
	ZWG: {Code: ZWG, Numeric: "924", Name: "Zimbabwe Gold", MinorUnit: 2},
}

var codes = [...]Code{
	AED,
	AFN,
	ALL,
	AMD,
	AOA,
	ARS,
	AUD,
	AWG,
	AZN,
	BAM,
	BBD,
	BDT,
	BHD,
	BIF,
	BMD,
	BND,
	BOB,
	BOV,
	BRL,
	BSD,
	BTN,
	BWP,
	BYN,
	BZD,
	CAD,
	CDF,
	CHE,
	CHF,
	CHW,
	CLF,
	CLP,
	CNY,
	COP,
	COU,
	CRC,
	CUP,
	CVE,
	CZK,
	DJF,
	DKK,
	DOP,
	DZD,
	EGP,
	ERN,
	ETB,
	EUR,
	FJD,
	FKP,
	GBP,
	GEL,
	GHS,
	GIP,
	GMD,
	GNF,
	GTQ,
	GYD,
	HKD,
	HNL,
	HTG,
	HUF,
	IDR,
	ILS,
	INR,
	IQD,
	IRR,
	ISK,
	JMD,
	JOD,
	JPY,
	KES,
	KGS,
	KHR,
	KMF,
	KPW,
	KRW,
	KWD,
	KYD,
	KZT,
	LAK,
	LBP,
	LKR,
	LRD,
	LSL,
	LYD,
	MAD,
	MDL,
	MGA,
	MKD,
	MMK,
	MNT,
	MOP,
	MRU,
	MUR,
	MVR,
	MWK,
	MXN,
	MXV,
	MYR,
	MZN,
	NAD,
	NGN,
	NIO,
	NOK,
	NPR,
	NZD,
	OMR,
	PAB,
	PEN,
	PGK,
	PHP,
	PKR,
	PLN,
	PYG,
	QAR,
	RON,
	RSD,
	RUB,
	RWF,
	SAR,
	SBD,
	SCR,
	SDG,
	SEK,
	SGD,
	SHP,
	SLE,
	SOS,
	SRD,
	SSP,
	STN,
	SVC,
	SYP,
	SZL,
	THB,
	TJS,
	TMT,
	TND,
	TOP,
	TRY,
	TTD,
	TWD,
	TZS,
	UAH,
	UGX,
	USD,
	USN,
	UYI,
	UYU,
	UYW,
	UZS,
	VED,
	VES,
	VND,
	VUV,
	WST,
	XAD,
	XAF,
	XAG,
	XAU,
	XBA,
	XBB,
	XBC,
	XBD,
	XCD,
	XCG,
	XDR,
	XOF,
	XPD,
	XPF,
	XPT,
	XSU,
	XTS,
	XUA,
	XXX,
	YER,
	ZAR,
	ZMW,
	ZWG,
}
