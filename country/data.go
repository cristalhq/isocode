package country

// ISO 3166-1 codes named by their alpha-2 forms.
const (
	AD Code = 20
	AE Code = 784
	AF Code = 4
	AG Code = 28
	AI Code = 660
	AL Code = 8
	AM Code = 51
	AO Code = 24
	AQ Code = 10
	AR Code = 32
	AS Code = 16
	AT Code = 40
	AU Code = 36
	AW Code = 533
	AX Code = 248
	AZ Code = 31
	BA Code = 70
	BB Code = 52
	BD Code = 50
	BE Code = 56
	BF Code = 854
	BG Code = 100
	BH Code = 48
	BI Code = 108
	BJ Code = 204
	BL Code = 652
	BM Code = 60
	BN Code = 96
	BO Code = 68
	BQ Code = 535
	BR Code = 76
	BS Code = 44
	BT Code = 64
	BV Code = 74
	BW Code = 72
	BY Code = 112
	BZ Code = 84
	CA Code = 124
	CC Code = 166
	CD Code = 180
	CF Code = 140
	CG Code = 178
	CH Code = 756
	CI Code = 384
	CK Code = 184
	CL Code = 152
	CM Code = 120
	CN Code = 156
	CO Code = 170
	CR Code = 188
	CU Code = 192
	CV Code = 132
	CW Code = 531
	CX Code = 162
	CY Code = 196
	CZ Code = 203
	DE Code = 276
	DJ Code = 262
	DK Code = 208
	DM Code = 212
	DO Code = 214
	DZ Code = 12
	EC Code = 218
	EE Code = 233
	EG Code = 818
	EH Code = 732
	ER Code = 232
	ES Code = 724
	ET Code = 231
	FI Code = 246
	FJ Code = 242
	FK Code = 238
	FM Code = 583
	FO Code = 234
	FR Code = 250
	GA Code = 266
	GB Code = 826
	GD Code = 308
	GE Code = 268
	GF Code = 254
	GG Code = 831
	GH Code = 288
	GI Code = 292
	GL Code = 304
	GM Code = 270
	GN Code = 324
	GP Code = 312
	GQ Code = 226
	GR Code = 300
	GS Code = 239
	GT Code = 320
	GU Code = 316
	GW Code = 624
	GY Code = 328
	HK Code = 344
	HM Code = 334
	HN Code = 340
	HR Code = 191
	HT Code = 332
	HU Code = 348
	ID Code = 360
	IE Code = 372
	IL Code = 376
	IM Code = 833
	IN Code = 356
	IO Code = 86
	IQ Code = 368
	IR Code = 364
	IS Code = 352
	IT Code = 380
	JE Code = 832
	JM Code = 388
	JO Code = 400
	JP Code = 392
	KE Code = 404
	KG Code = 417
	KH Code = 116
	KI Code = 296
	KM Code = 174
	KN Code = 659
	KP Code = 408
	KR Code = 410
	KW Code = 414
	KY Code = 136
	KZ Code = 398
	LA Code = 418
	LB Code = 422
	LC Code = 662
	LI Code = 438
	LK Code = 144
	LR Code = 430
	LS Code = 426
	LT Code = 440
	LU Code = 442
	LV Code = 428
	LY Code = 434
	MA Code = 504
	MC Code = 492
	MD Code = 498
	ME Code = 499
	MF Code = 663
	MG Code = 450
	MH Code = 584
	MK Code = 807
	ML Code = 466
	MM Code = 104
	MN Code = 496
	MO Code = 446
	MP Code = 580
	MQ Code = 474
	MR Code = 478
	MS Code = 500
	MT Code = 470
	MU Code = 480
	MV Code = 462
	MW Code = 454
	MX Code = 484
	MY Code = 458
	MZ Code = 508
	NA Code = 516
	NC Code = 540
	NE Code = 562
	NF Code = 574
	NG Code = 566
	NI Code = 558
	NL Code = 528
	NO Code = 578
	NP Code = 524
	NR Code = 520
	NU Code = 570
	NZ Code = 554
	OM Code = 512
	PA Code = 591
	PE Code = 604
	PF Code = 258
	PG Code = 598
	PH Code = 608
	PK Code = 586
	PL Code = 616
	PM Code = 666
	PN Code = 612
	PR Code = 630
	PS Code = 275
	PT Code = 620
	PW Code = 585
	PY Code = 600
	QA Code = 634
	RE Code = 638
	RO Code = 642
	RS Code = 688
	RU Code = 643
	RW Code = 646
	SA Code = 682
	SB Code = 90
	SC Code = 690
	SD Code = 729
	SE Code = 752
	SG Code = 702
	SH Code = 654
	SI Code = 705
	SJ Code = 744
	SK Code = 703
	SL Code = 694
	SM Code = 674
	SN Code = 686
	SO Code = 706
	SR Code = 740
	SS Code = 728
	ST Code = 678
	SV Code = 222
	SX Code = 534
	SY Code = 760
	SZ Code = 748
	TC Code = 796
	TD Code = 148
	TF Code = 260
	TG Code = 768
	TH Code = 764
	TJ Code = 762
	TK Code = 772
	TL Code = 626
	TM Code = 795
	TN Code = 788
	TO Code = 776
	TR Code = 792
	TT Code = 780
	TV Code = 798
	TW Code = 158
	TZ Code = 834
	UA Code = 804
	UG Code = 800
	UM Code = 581
	US Code = 840
	UY Code = 858
	UZ Code = 860
	VA Code = 336
	VC Code = 670
	VE Code = 862
	VG Code = 92
	VI Code = 850
	VN Code = 704
	VU Code = 548
	WF Code = 876
	WS Code = 882
	YE Code = 887
	YT Code = 175
	ZA Code = 710
	ZM Code = 894
	ZW Code = 716
)

var table = map[Code]Info{
	AD: {Code: AD, Alpha2: "AD", Alpha3: "AND", Numeric: "020", Name: "Andorra", Flag: "🇦🇩"},
	AE: {Code: AE, Alpha2: "AE", Alpha3: "ARE", Numeric: "784", Name: "United Arab Emirates", Flag: "🇦🇪"},
	AF: {Code: AF, Alpha2: "AF", Alpha3: "AFG", Numeric: "004", Name: "Afghanistan", Flag: "🇦🇫"},
	AG: {Code: AG, Alpha2: "AG", Alpha3: "ATG", Numeric: "028", Name: "Antigua and Barbuda", Flag: "🇦🇬"},
	AI: {Code: AI, Alpha2: "AI", Alpha3: "AIA", Numeric: "660", Name: "Anguilla", Flag: "🇦🇮"},
	AL: {Code: AL, Alpha2: "AL", Alpha3: "ALB", Numeric: "008", Name: "Albania", Flag: "🇦🇱"},
	AM: {Code: AM, Alpha2: "AM", Alpha3: "ARM", Numeric: "051", Name: "Armenia", Flag: "🇦🇲"},
	AO: {Code: AO, Alpha2: "AO", Alpha3: "AGO", Numeric: "024", Name: "Angola", Flag: "🇦🇴"},
	AQ: {Code: AQ, Alpha2: "AQ", Alpha3: "ATA", Numeric: "010", Name: "Antarctica", Flag: "🇦🇶"},
	AR: {Code: AR, Alpha2: "AR", Alpha3: "ARG", Numeric: "032", Name: "Argentina", Flag: "🇦🇷"},
	AS: {Code: AS, Alpha2: "AS", Alpha3: "ASM", Numeric: "016", Name: "American Samoa", Flag: "🇦🇸"},
	AT: {Code: AT, Alpha2: "AT", Alpha3: "AUT", Numeric: "040", Name: "Austria", Flag: "🇦🇹"},
	AU: {Code: AU, Alpha2: "AU", Alpha3: "AUS", Numeric: "036", Name: "Australia", Flag: "🇦🇺"},
	AW: {Code: AW, Alpha2: "AW", Alpha3: "ABW", Numeric: "533", Name: "Aruba", Flag: "🇦🇼"},
	AX: {Code: AX, Alpha2: "AX", Alpha3: "ALA", Numeric: "248", Name: "Åland Islands", Flag: "🇦🇽"},
	AZ: {Code: AZ, Alpha2: "AZ", Alpha3: "AZE", Numeric: "031", Name: "Azerbaijan", Flag: "🇦🇿"},
	BA: {Code: BA, Alpha2: "BA", Alpha3: "BIH", Numeric: "070", Name: "Bosnia and Herzegovina", Flag: "🇧🇦"},
	BB: {Code: BB, Alpha2: "BB", Alpha3: "BRB", Numeric: "052", Name: "Barbados", Flag: "🇧🇧"},
	BD: {Code: BD, Alpha2: "BD", Alpha3: "BGD", Numeric: "050", Name: "Bangladesh", Flag: "🇧🇩"},
	BE: {Code: BE, Alpha2: "BE", Alpha3: "BEL", Numeric: "056", Name: "Belgium", Flag: "🇧🇪"},
	BF: {Code: BF, Alpha2: "BF", Alpha3: "BFA", Numeric: "854", Name: "Burkina Faso", Flag: "🇧🇫"},
	BG: {Code: BG, Alpha2: "BG", Alpha3: "BGR", Numeric: "100", Name: "Bulgaria", Flag: "🇧🇬"},
	BH: {Code: BH, Alpha2: "BH", Alpha3: "BHR", Numeric: "048", Name: "Bahrain", Flag: "🇧🇭"},
	BI: {Code: BI, Alpha2: "BI", Alpha3: "BDI", Numeric: "108", Name: "Burundi", Flag: "🇧🇮"},
	BJ: {Code: BJ, Alpha2: "BJ", Alpha3: "BEN", Numeric: "204", Name: "Benin", Flag: "🇧🇯"},
	BL: {Code: BL, Alpha2: "BL", Alpha3: "BLM", Numeric: "652", Name: "Saint Barthélemy", Flag: "🇧🇱"},
	BM: {Code: BM, Alpha2: "BM", Alpha3: "BMU", Numeric: "060", Name: "Bermuda", Flag: "🇧🇲"},
	BN: {Code: BN, Alpha2: "BN", Alpha3: "BRN", Numeric: "096", Name: "Brunei Darussalam", Flag: "🇧🇳"},
	BO: {Code: BO, Alpha2: "BO", Alpha3: "BOL", Numeric: "068", Name: "Bolivia, Plurinational State of", Flag: "🇧🇴"},
	BQ: {Code: BQ, Alpha2: "BQ", Alpha3: "BES", Numeric: "535", Name: "Bonaire, Sint Eustatius and Saba", Flag: "🇧🇶"},
	BR: {Code: BR, Alpha2: "BR", Alpha3: "BRA", Numeric: "076", Name: "Brazil", Flag: "🇧🇷"},
	BS: {Code: BS, Alpha2: "BS", Alpha3: "BHS", Numeric: "044", Name: "Bahamas", Flag: "🇧🇸"},
	BT: {Code: BT, Alpha2: "BT", Alpha3: "BTN", Numeric: "064", Name: "Bhutan", Flag: "🇧🇹"},
	BV: {Code: BV, Alpha2: "BV", Alpha3: "BVT", Numeric: "074", Name: "Bouvet Island", Flag: "🇧🇻"},
	BW: {Code: BW, Alpha2: "BW", Alpha3: "BWA", Numeric: "072", Name: "Botswana", Flag: "🇧🇼"},
	BY: {Code: BY, Alpha2: "BY", Alpha3: "BLR", Numeric: "112", Name: "Belarus", Flag: "🇧🇾"},
	BZ: {Code: BZ, Alpha2: "BZ", Alpha3: "BLZ", Numeric: "084", Name: "Belize", Flag: "🇧🇿"},
	CA: {Code: CA, Alpha2: "CA", Alpha3: "CAN", Numeric: "124", Name: "Canada", Flag: "🇨🇦"},
	CC: {Code: CC, Alpha2: "CC", Alpha3: "CCK", Numeric: "166", Name: "Cocos (Keeling) Islands", Flag: "🇨🇨"},
	CD: {Code: CD, Alpha2: "CD", Alpha3: "COD", Numeric: "180", Name: "Congo, The Democratic Republic of the", Flag: "🇨🇩"},
	CF: {Code: CF, Alpha2: "CF", Alpha3: "CAF", Numeric: "140", Name: "Central African Republic", Flag: "🇨🇫"},
	CG: {Code: CG, Alpha2: "CG", Alpha3: "COG", Numeric: "178", Name: "Congo", Flag: "🇨🇬"},
	CH: {Code: CH, Alpha2: "CH", Alpha3: "CHE", Numeric: "756", Name: "Switzerland", Flag: "🇨🇭"},
	CI: {Code: CI, Alpha2: "CI", Alpha3: "CIV", Numeric: "384", Name: "Côte d'Ivoire", Flag: "🇨🇮"},
	CK: {Code: CK, Alpha2: "CK", Alpha3: "COK", Numeric: "184", Name: "Cook Islands", Flag: "🇨🇰"},
	CL: {Code: CL, Alpha2: "CL", Alpha3: "CHL", Numeric: "152", Name: "Chile", Flag: "🇨🇱"},
	CM: {Code: CM, Alpha2: "CM", Alpha3: "CMR", Numeric: "120", Name: "Cameroon", Flag: "🇨🇲"},
	CN: {Code: CN, Alpha2: "CN", Alpha3: "CHN", Numeric: "156", Name: "China", Flag: "🇨🇳"},
	CO: {Code: CO, Alpha2: "CO", Alpha3: "COL", Numeric: "170", Name: "Colombia", Flag: "🇨🇴"},
	CR: {Code: CR, Alpha2: "CR", Alpha3: "CRI", Numeric: "188", Name: "Costa Rica", Flag: "🇨🇷"},
	CU: {Code: CU, Alpha2: "CU", Alpha3: "CUB", Numeric: "192", Name: "Cuba", Flag: "🇨🇺"},
	CV: {Code: CV, Alpha2: "CV", Alpha3: "CPV", Numeric: "132", Name: "Cabo Verde", Flag: "🇨🇻"},
	CW: {Code: CW, Alpha2: "CW", Alpha3: "CUW", Numeric: "531", Name: "Curaçao", Flag: "🇨🇼"},
	CX: {Code: CX, Alpha2: "CX", Alpha3: "CXR", Numeric: "162", Name: "Christmas Island", Flag: "🇨🇽"},
	CY: {Code: CY, Alpha2: "CY", Alpha3: "CYP", Numeric: "196", Name: "Cyprus", Flag: "🇨🇾"},
	CZ: {Code: CZ, Alpha2: "CZ", Alpha3: "CZE", Numeric: "203", Name: "Czechia", Flag: "🇨🇿"},
	DE: {Code: DE, Alpha2: "DE", Alpha3: "DEU", Numeric: "276", Name: "Germany", Flag: "🇩🇪"},
	DJ: {Code: DJ, Alpha2: "DJ", Alpha3: "DJI", Numeric: "262", Name: "Djibouti", Flag: "🇩🇯"},
	DK: {Code: DK, Alpha2: "DK", Alpha3: "DNK", Numeric: "208", Name: "Denmark", Flag: "🇩🇰"},
	DM: {Code: DM, Alpha2: "DM", Alpha3: "DMA", Numeric: "212", Name: "Dominica", Flag: "🇩🇲"},
	DO: {Code: DO, Alpha2: "DO", Alpha3: "DOM", Numeric: "214", Name: "Dominican Republic", Flag: "🇩🇴"},
	DZ: {Code: DZ, Alpha2: "DZ", Alpha3: "DZA", Numeric: "012", Name: "Algeria", Flag: "🇩🇿"},
	EC: {Code: EC, Alpha2: "EC", Alpha3: "ECU", Numeric: "218", Name: "Ecuador", Flag: "🇪🇨"},
	EE: {Code: EE, Alpha2: "EE", Alpha3: "EST", Numeric: "233", Name: "Estonia", Flag: "🇪🇪"},
	EG: {Code: EG, Alpha2: "EG", Alpha3: "EGY", Numeric: "818", Name: "Egypt", Flag: "🇪🇬"},
	EH: {Code: EH, Alpha2: "EH", Alpha3: "ESH", Numeric: "732", Name: "Western Sahara", Flag: "🇪🇭"},
	ER: {Code: ER, Alpha2: "ER", Alpha3: "ERI", Numeric: "232", Name: "Eritrea", Flag: "🇪🇷"},
	ES: {Code: ES, Alpha2: "ES", Alpha3: "ESP", Numeric: "724", Name: "Spain", Flag: "🇪🇸"},
	ET: {Code: ET, Alpha2: "ET", Alpha3: "ETH", Numeric: "231", Name: "Ethiopia", Flag: "🇪🇹"},
	FI: {Code: FI, Alpha2: "FI", Alpha3: "FIN", Numeric: "246", Name: "Finland", Flag: "🇫🇮"},
	FJ: {Code: FJ, Alpha2: "FJ", Alpha3: "FJI", Numeric: "242", Name: "Fiji", Flag: "🇫🇯"},
	FK: {Code: FK, Alpha2: "FK", Alpha3: "FLK", Numeric: "238", Name: "Falkland Islands (Malvinas)", Flag: "🇫🇰"},
	FM: {Code: FM, Alpha2: "FM", Alpha3: "FSM", Numeric: "583", Name: "Micronesia, Federated States of", Flag: "🇫🇲"},
	FO: {Code: FO, Alpha2: "FO", Alpha3: "FRO", Numeric: "234", Name: "Faroe Islands", Flag: "🇫🇴"},
	FR: {Code: FR, Alpha2: "FR", Alpha3: "FRA", Numeric: "250", Name: "France", Flag: "🇫🇷"},
	GA: {Code: GA, Alpha2: "GA", Alpha3: "GAB", Numeric: "266", Name: "Gabon", Flag: "🇬🇦"},
	GB: {Code: GB, Alpha2: "GB", Alpha3: "GBR", Numeric: "826", Name: "United Kingdom", Flag: "🇬🇧"},
	GD: {Code: GD, Alpha2: "GD", Alpha3: "GRD", Numeric: "308", Name: "Grenada", Flag: "🇬🇩"},
	GE: {Code: GE, Alpha2: "GE", Alpha3: "GEO", Numeric: "268", Name: "Georgia", Flag: "🇬🇪"},
	GF: {Code: GF, Alpha2: "GF", Alpha3: "GUF", Numeric: "254", Name: "French Guiana", Flag: "🇬🇫"},
	GG: {Code: GG, Alpha2: "GG", Alpha3: "GGY", Numeric: "831", Name: "Guernsey", Flag: "🇬🇬"},
	GH: {Code: GH, Alpha2: "GH", Alpha3: "GHA", Numeric: "288", Name: "Ghana", Flag: "🇬🇭"},
	GI: {Code: GI, Alpha2: "GI", Alpha3: "GIB", Numeric: "292", Name: "Gibraltar", Flag: "🇬🇮"},
	GL: {Code: GL, Alpha2: "GL", Alpha3: "GRL", Numeric: "304", Name: "Greenland", Flag: "🇬🇱"},
	GM: {Code: GM, Alpha2: "GM", Alpha3: "GMB", Numeric: "270", Name: "Gambia", Flag: "🇬🇲"},
	GN: {Code: GN, Alpha2: "GN", Alpha3: "GIN", Numeric: "324", Name: "Guinea", Flag: "🇬🇳"},
	GP: {Code: GP, Alpha2: "GP", Alpha3: "GLP", Numeric: "312", Name: "Guadeloupe", Flag: "🇬🇵"},
	GQ: {Code: GQ, Alpha2: "GQ", Alpha3: "GNQ", Numeric: "226", Name: "Equatorial Guinea", Flag: "🇬🇶"},
	GR: {Code: GR, Alpha2: "GR", Alpha3: "GRC", Numeric: "300", Name: "Greece", Flag: "🇬🇷"},
	GS: {Code: GS, Alpha2: "GS", Alpha3: "SGS", Numeric: "239", Name: "South Georgia and the South Sandwich Islands", Flag: "🇬🇸"},
	GT: {Code: GT, Alpha2: "GT", Alpha3: "GTM", Numeric: "320", Name: "Guatemala", Flag: "🇬🇹"},
	GU: {Code: GU, Alpha2: "GU", Alpha3: "GUM", Numeric: "316", Name: "Guam", Flag: "🇬🇺"},
	GW: {Code: GW, Alpha2: "GW", Alpha3: "GNB", Numeric: "624", Name: "Guinea-Bissau", Flag: "🇬🇼"},
	GY: {Code: GY, Alpha2: "GY", Alpha3: "GUY", Numeric: "328", Name: "Guyana", Flag: "🇬🇾"},
	HK: {Code: HK, Alpha2: "HK", Alpha3: "HKG", Numeric: "344", Name: "Hong Kong", Flag: "🇭🇰"},
	HM: {Code: HM, Alpha2: "HM", Alpha3: "HMD", Numeric: "334", Name: "Heard Island and McDonald Islands", Flag: "🇭🇲"},
	HN: {Code: HN, Alpha2: "HN", Alpha3: "HND", Numeric: "340", Name: "Honduras", Flag: "🇭🇳"},
	HR: {Code: HR, Alpha2: "HR", Alpha3: "HRV", Numeric: "191", Name: "Croatia", Flag: "🇭🇷"},
	HT: {Code: HT, Alpha2: "HT", Alpha3: "HTI", Numeric: "332", Name: "Haiti", Flag: "🇭🇹"},
	HU: {Code: HU, Alpha2: "HU", Alpha3: "HUN", Numeric: "348", Name: "Hungary", Flag: "🇭🇺"},
	ID: {Code: ID, Alpha2: "ID", Alpha3: "IDN", Numeric: "360", Name: "Indonesia", Flag: "🇮🇩"},
	IE: {Code: IE, Alpha2: "IE", Alpha3: "IRL", Numeric: "372", Name: "Ireland", Flag: "🇮🇪"},
	IL: {Code: IL, Alpha2: "IL", Alpha3: "ISR", Numeric: "376", Name: "Israel", Flag: "🇮🇱"},
	IM: {Code: IM, Alpha2: "IM", Alpha3: "IMN", Numeric: "833", Name: "Isle of Man", Flag: "🇮🇲"},
	IN: {Code: IN, Alpha2: "IN", Alpha3: "IND", Numeric: "356", Name: "India", Flag: "🇮🇳"},
	IO: {Code: IO, Alpha2: "IO", Alpha3: "IOT", Numeric: "086", Name: "British Indian Ocean Territory", Flag: "🇮🇴"},
	IQ: {Code: IQ, Alpha2: "IQ", Alpha3: "IRQ", Numeric: "368", Name: "Iraq", Flag: "🇮🇶"},
	IR: {Code: IR, Alpha2: "IR", Alpha3: "IRN", Numeric: "364", Name: "Iran, Islamic Republic of", Flag: "🇮🇷"},
	IS: {Code: IS, Alpha2: "IS", Alpha3: "ISL", Numeric: "352", Name: "Iceland", Flag: "🇮🇸"},
	IT: {Code: IT, Alpha2: "IT", Alpha3: "ITA", Numeric: "380", Name: "Italy", Flag: "🇮🇹"},
	JE: {Code: JE, Alpha2: "JE", Alpha3: "JEY", Numeric: "832", Name: "Jersey", Flag: "🇯🇪"},
	JM: {Code: JM, Alpha2: "JM", Alpha3: "JAM", Numeric: "388", Name: "Jamaica", Flag: "🇯🇲"},
	JO: {Code: JO, Alpha2: "JO", Alpha3: "JOR", Numeric: "400", Name: "Jordan", Flag: "🇯🇴"},
	JP: {Code: JP, Alpha2: "JP", Alpha3: "JPN", Numeric: "392", Name: "Japan", Flag: "🇯🇵"},
	KE: {Code: KE, Alpha2: "KE", Alpha3: "KEN", Numeric: "404", Name: "Kenya", Flag: "🇰🇪"},
	KG: {Code: KG, Alpha2: "KG", Alpha3: "KGZ", Numeric: "417", Name: "Kyrgyzstan", Flag: "🇰🇬"},
	KH: {Code: KH, Alpha2: "KH", Alpha3: "KHM", Numeric: "116", Name: "Cambodia", Flag: "🇰🇭"},
	KI: {Code: KI, Alpha2: "KI", Alpha3: "KIR", Numeric: "296", Name: "Kiribati", Flag: "🇰🇮"},
	KM: {Code: KM, Alpha2: "KM", Alpha3: "COM", Numeric: "174", Name: "Comoros", Flag: "🇰🇲"},
	KN: {Code: KN, Alpha2: "KN", Alpha3: "KNA", Numeric: "659", Name: "Saint Kitts and Nevis", Flag: "🇰🇳"},
	KP: {Code: KP, Alpha2: "KP", Alpha3: "PRK", Numeric: "408", Name: "Korea, Democratic People's Republic of", Flag: "🇰🇵"},
	KR: {Code: KR, Alpha2: "KR", Alpha3: "KOR", Numeric: "410", Name: "Korea, Republic of", Flag: "🇰🇷"},
	KW: {Code: KW, Alpha2: "KW", Alpha3: "KWT", Numeric: "414", Name: "Kuwait", Flag: "🇰🇼"},
	KY: {Code: KY, Alpha2: "KY", Alpha3: "CYM", Numeric: "136", Name: "Cayman Islands", Flag: "🇰🇾"},
	KZ: {Code: KZ, Alpha2: "KZ", Alpha3: "KAZ", Numeric: "398", Name: "Kazakhstan", Flag: "🇰🇿"},
	LA: {Code: LA, Alpha2: "LA", Alpha3: "LAO", Numeric: "418", Name: "Lao People's Democratic Republic", Flag: "🇱🇦"},
	LB: {Code: LB, Alpha2: "LB", Alpha3: "LBN", Numeric: "422", Name: "Lebanon", Flag: "🇱🇧"},
	LC: {Code: LC, Alpha2: "LC", Alpha3: "LCA", Numeric: "662", Name: "Saint Lucia", Flag: "🇱🇨"},
	LI: {Code: LI, Alpha2: "LI", Alpha3: "LIE", Numeric: "438", Name: "Liechtenstein", Flag: "🇱🇮"},
	LK: {Code: LK, Alpha2: "LK", Alpha3: "LKA", Numeric: "144", Name: "Sri Lanka", Flag: "🇱🇰"},
	LR: {Code: LR, Alpha2: "LR", Alpha3: "LBR", Numeric: "430", Name: "Liberia", Flag: "🇱🇷"},
	LS: {Code: LS, Alpha2: "LS", Alpha3: "LSO", Numeric: "426", Name: "Lesotho", Flag: "🇱🇸"},
	LT: {Code: LT, Alpha2: "LT", Alpha3: "LTU", Numeric: "440", Name: "Lithuania", Flag: "🇱🇹"},
	LU: {Code: LU, Alpha2: "LU", Alpha3: "LUX", Numeric: "442", Name: "Luxembourg", Flag: "🇱🇺"},
	LV: {Code: LV, Alpha2: "LV", Alpha3: "LVA", Numeric: "428", Name: "Latvia", Flag: "🇱🇻"},
	LY: {Code: LY, Alpha2: "LY", Alpha3: "LBY", Numeric: "434", Name: "Libya", Flag: "🇱🇾"},
	MA: {Code: MA, Alpha2: "MA", Alpha3: "MAR", Numeric: "504", Name: "Morocco", Flag: "🇲🇦"},
	MC: {Code: MC, Alpha2: "MC", Alpha3: "MCO", Numeric: "492", Name: "Monaco", Flag: "🇲🇨"},
	MD: {Code: MD, Alpha2: "MD", Alpha3: "MDA", Numeric: "498", Name: "Moldova, Republic of", Flag: "🇲🇩"},
	ME: {Code: ME, Alpha2: "ME", Alpha3: "MNE", Numeric: "499", Name: "Montenegro", Flag: "🇲🇪"},
	MF: {Code: MF, Alpha2: "MF", Alpha3: "MAF", Numeric: "663", Name: "Saint Martin (French part)", Flag: "🇲🇫"},
	MG: {Code: MG, Alpha2: "MG", Alpha3: "MDG", Numeric: "450", Name: "Madagascar", Flag: "🇲🇬"},
	MH: {Code: MH, Alpha2: "MH", Alpha3: "MHL", Numeric: "584", Name: "Marshall Islands", Flag: "🇲🇭"},
	MK: {Code: MK, Alpha2: "MK", Alpha3: "MKD", Numeric: "807", Name: "North Macedonia", Flag: "🇲🇰"},
	ML: {Code: ML, Alpha2: "ML", Alpha3: "MLI", Numeric: "466", Name: "Mali", Flag: "🇲🇱"},
	MM: {Code: MM, Alpha2: "MM", Alpha3: "MMR", Numeric: "104", Name: "Myanmar", Flag: "🇲🇲"},
	MN: {Code: MN, Alpha2: "MN", Alpha3: "MNG", Numeric: "496", Name: "Mongolia", Flag: "🇲🇳"},
	MO: {Code: MO, Alpha2: "MO", Alpha3: "MAC", Numeric: "446", Name: "Macao", Flag: "🇲🇴"},
	MP: {Code: MP, Alpha2: "MP", Alpha3: "MNP", Numeric: "580", Name: "Northern Mariana Islands", Flag: "🇲🇵"},
	MQ: {Code: MQ, Alpha2: "MQ", Alpha3: "MTQ", Numeric: "474", Name: "Martinique", Flag: "🇲🇶"},
	MR: {Code: MR, Alpha2: "MR", Alpha3: "MRT", Numeric: "478", Name: "Mauritania", Flag: "🇲🇷"},
	MS: {Code: MS, Alpha2: "MS", Alpha3: "MSR", Numeric: "500", Name: "Montserrat", Flag: "🇲🇸"},
	MT: {Code: MT, Alpha2: "MT", Alpha3: "MLT", Numeric: "470", Name: "Malta", Flag: "🇲🇹"},
	MU: {Code: MU, Alpha2: "MU", Alpha3: "MUS", Numeric: "480", Name: "Mauritius", Flag: "🇲🇺"},
	MV: {Code: MV, Alpha2: "MV", Alpha3: "MDV", Numeric: "462", Name: "Maldives", Flag: "🇲🇻"},
	MW: {Code: MW, Alpha2: "MW", Alpha3: "MWI", Numeric: "454", Name: "Malawi", Flag: "🇲🇼"},
	MX: {Code: MX, Alpha2: "MX", Alpha3: "MEX", Numeric: "484", Name: "Mexico", Flag: "🇲🇽"},
	MY: {Code: MY, Alpha2: "MY", Alpha3: "MYS", Numeric: "458", Name: "Malaysia", Flag: "🇲🇾"},
	MZ: {Code: MZ, Alpha2: "MZ", Alpha3: "MOZ", Numeric: "508", Name: "Mozambique", Flag: "🇲🇿"},
	NA: {Code: NA, Alpha2: "NA", Alpha3: "NAM", Numeric: "516", Name: "Namibia", Flag: "🇳🇦"},
	NC: {Code: NC, Alpha2: "NC", Alpha3: "NCL", Numeric: "540", Name: "New Caledonia", Flag: "🇳🇨"},
	NE: {Code: NE, Alpha2: "NE", Alpha3: "NER", Numeric: "562", Name: "Niger", Flag: "🇳🇪"},
	NF: {Code: NF, Alpha2: "NF", Alpha3: "NFK", Numeric: "574", Name: "Norfolk Island", Flag: "🇳🇫"},
	NG: {Code: NG, Alpha2: "NG", Alpha3: "NGA", Numeric: "566", Name: "Nigeria", Flag: "🇳🇬"},
	NI: {Code: NI, Alpha2: "NI", Alpha3: "NIC", Numeric: "558", Name: "Nicaragua", Flag: "🇳🇮"},
	NL: {Code: NL, Alpha2: "NL", Alpha3: "NLD", Numeric: "528", Name: "Netherlands", Flag: "🇳🇱"},
	NO: {Code: NO, Alpha2: "NO", Alpha3: "NOR", Numeric: "578", Name: "Norway", Flag: "🇳🇴"},
	NP: {Code: NP, Alpha2: "NP", Alpha3: "NPL", Numeric: "524", Name: "Nepal", Flag: "🇳🇵"},
	NR: {Code: NR, Alpha2: "NR", Alpha3: "NRU", Numeric: "520", Name: "Nauru", Flag: "🇳🇷"},
	NU: {Code: NU, Alpha2: "NU", Alpha3: "NIU", Numeric: "570", Name: "Niue", Flag: "🇳🇺"},
	NZ: {Code: NZ, Alpha2: "NZ", Alpha3: "NZL", Numeric: "554", Name: "New Zealand", Flag: "🇳🇿"},
	OM: {Code: OM, Alpha2: "OM", Alpha3: "OMN", Numeric: "512", Name: "Oman", Flag: "🇴🇲"},
	PA: {Code: PA, Alpha2: "PA", Alpha3: "PAN", Numeric: "591", Name: "Panama", Flag: "🇵🇦"},
	PE: {Code: PE, Alpha2: "PE", Alpha3: "PER", Numeric: "604", Name: "Peru", Flag: "🇵🇪"},
	PF: {Code: PF, Alpha2: "PF", Alpha3: "PYF", Numeric: "258", Name: "French Polynesia", Flag: "🇵🇫"},
	PG: {Code: PG, Alpha2: "PG", Alpha3: "PNG", Numeric: "598", Name: "Papua New Guinea", Flag: "🇵🇬"},
	PH: {Code: PH, Alpha2: "PH", Alpha3: "PHL", Numeric: "608", Name: "Philippines", Flag: "🇵🇭"},
	PK: {Code: PK, Alpha2: "PK", Alpha3: "PAK", Numeric: "586", Name: "Pakistan", Flag: "🇵🇰"},
	PL: {Code: PL, Alpha2: "PL", Alpha3: "POL", Numeric: "616", Name: "Poland", Flag: "🇵🇱"},
	PM: {Code: PM, Alpha2: "PM", Alpha3: "SPM", Numeric: "666", Name: "Saint Pierre and Miquelon", Flag: "🇵🇲"},
	PN: {Code: PN, Alpha2: "PN", Alpha3: "PCN", Numeric: "612", Name: "Pitcairn", Flag: "🇵🇳"},
	PR: {Code: PR, Alpha2: "PR", Alpha3: "PRI", Numeric: "630", Name: "Puerto Rico", Flag: "🇵🇷"},
	PS: {Code: PS, Alpha2: "PS", Alpha3: "PSE", Numeric: "275", Name: "Palestine, State of", Flag: "🇵🇸"},
	PT: {Code: PT, Alpha2: "PT", Alpha3: "PRT", Numeric: "620", Name: "Portugal", Flag: "🇵🇹"},
	PW: {Code: PW, Alpha2: "PW", Alpha3: "PLW", Numeric: "585", Name: "Palau", Flag: "🇵🇼"},
	PY: {Code: PY, Alpha2: "PY", Alpha3: "PRY", Numeric: "600", Name: "Paraguay", Flag: "🇵🇾"},
	QA: {Code: QA, Alpha2: "QA", Alpha3: "QAT", Numeric: "634", Name: "Qatar", Flag: "🇶🇦"},
	RE: {Code: RE, Alpha2: "RE", Alpha3: "REU", Numeric: "638", Name: "Réunion", Flag: "🇷🇪"},
	RO: {Code: RO, Alpha2: "RO", Alpha3: "ROU", Numeric: "642", Name: "Romania", Flag: "🇷🇴"},
	RS: {Code: RS, Alpha2: "RS", Alpha3: "SRB", Numeric: "688", Name: "Serbia", Flag: "🇷🇸"},
	RU: {Code: RU, Alpha2: "RU", Alpha3: "RUS", Numeric: "643", Name: "Russian Federation", Flag: "🇷🇺"},
	RW: {Code: RW, Alpha2: "RW", Alpha3: "RWA", Numeric: "646", Name: "Rwanda", Flag: "🇷🇼"},
	SA: {Code: SA, Alpha2: "SA", Alpha3: "SAU", Numeric: "682", Name: "Saudi Arabia", Flag: "🇸🇦"},
	SB: {Code: SB, Alpha2: "SB", Alpha3: "SLB", Numeric: "090", Name: "Solomon Islands", Flag: "🇸🇧"},
	SC: {Code: SC, Alpha2: "SC", Alpha3: "SYC", Numeric: "690", Name: "Seychelles", Flag: "🇸🇨"},
	SD: {Code: SD, Alpha2: "SD", Alpha3: "SDN", Numeric: "729", Name: "Sudan", Flag: "🇸🇩"},
	SE: {Code: SE, Alpha2: "SE", Alpha3: "SWE", Numeric: "752", Name: "Sweden", Flag: "🇸🇪"},
	SG: {Code: SG, Alpha2: "SG", Alpha3: "SGP", Numeric: "702", Name: "Singapore", Flag: "🇸🇬"},
	SH: {Code: SH, Alpha2: "SH", Alpha3: "SHN", Numeric: "654", Name: "Saint Helena, Ascension and Tristan da Cunha", Flag: "🇸🇭"},
	SI: {Code: SI, Alpha2: "SI", Alpha3: "SVN", Numeric: "705", Name: "Slovenia", Flag: "🇸🇮"},
	SJ: {Code: SJ, Alpha2: "SJ", Alpha3: "SJM", Numeric: "744", Name: "Svalbard and Jan Mayen", Flag: "🇸🇯"},
	SK: {Code: SK, Alpha2: "SK", Alpha3: "SVK", Numeric: "703", Name: "Slovakia", Flag: "🇸🇰"},
	SL: {Code: SL, Alpha2: "SL", Alpha3: "SLE", Numeric: "694", Name: "Sierra Leone", Flag: "🇸🇱"},
	SM: {Code: SM, Alpha2: "SM", Alpha3: "SMR", Numeric: "674", Name: "San Marino", Flag: "🇸🇲"},
	SN: {Code: SN, Alpha2: "SN", Alpha3: "SEN", Numeric: "686", Name: "Senegal", Flag: "🇸🇳"},
	SO: {Code: SO, Alpha2: "SO", Alpha3: "SOM", Numeric: "706", Name: "Somalia", Flag: "🇸🇴"},
	SR: {Code: SR, Alpha2: "SR", Alpha3: "SUR", Numeric: "740", Name: "Suriname", Flag: "🇸🇷"},
	SS: {Code: SS, Alpha2: "SS", Alpha3: "SSD", Numeric: "728", Name: "South Sudan", Flag: "🇸🇸"},
	ST: {Code: ST, Alpha2: "ST", Alpha3: "STP", Numeric: "678", Name: "Sao Tome and Principe", Flag: "🇸🇹"},
	SV: {Code: SV, Alpha2: "SV", Alpha3: "SLV", Numeric: "222", Name: "El Salvador", Flag: "🇸🇻"},
	SX: {Code: SX, Alpha2: "SX", Alpha3: "SXM", Numeric: "534", Name: "Sint Maarten (Dutch part)", Flag: "🇸🇽"},
	SY: {Code: SY, Alpha2: "SY", Alpha3: "SYR", Numeric: "760", Name: "Syrian Arab Republic", Flag: "🇸🇾"},
	SZ: {Code: SZ, Alpha2: "SZ", Alpha3: "SWZ", Numeric: "748", Name: "Eswatini", Flag: "🇸🇿"},
	TC: {Code: TC, Alpha2: "TC", Alpha3: "TCA", Numeric: "796", Name: "Turks and Caicos Islands", Flag: "🇹🇨"},
	TD: {Code: TD, Alpha2: "TD", Alpha3: "TCD", Numeric: "148", Name: "Chad", Flag: "🇹🇩"},
	TF: {Code: TF, Alpha2: "TF", Alpha3: "ATF", Numeric: "260", Name: "French Southern Territories", Flag: "🇹🇫"},
	TG: {Code: TG, Alpha2: "TG", Alpha3: "TGO", Numeric: "768", Name: "Togo", Flag: "🇹🇬"},
	TH: {Code: TH, Alpha2: "TH", Alpha3: "THA", Numeric: "764", Name: "Thailand", Flag: "🇹🇭"},
	TJ: {Code: TJ, Alpha2: "TJ", Alpha3: "TJK", Numeric: "762", Name: "Tajikistan", Flag: "🇹🇯"},
	TK: {Code: TK, Alpha2: "TK", Alpha3: "TKL", Numeric: "772", Name: "Tokelau", Flag: "🇹🇰"},
	TL: {Code: TL, Alpha2: "TL", Alpha3: "TLS", Numeric: "626", Name: "Timor-Leste", Flag: "🇹🇱"},
	TM: {Code: TM, Alpha2: "TM", Alpha3: "TKM", Numeric: "795", Name: "Turkmenistan", Flag: "🇹🇲"},
	TN: {Code: TN, Alpha2: "TN", Alpha3: "TUN", Numeric: "788", Name: "Tunisia", Flag: "🇹🇳"},
	TO: {Code: TO, Alpha2: "TO", Alpha3: "TON", Numeric: "776", Name: "Tonga", Flag: "🇹🇴"},
	TR: {Code: TR, Alpha2: "TR", Alpha3: "TUR", Numeric: "792", Name: "Türkiye", Flag: "🇹🇷"},
	TT: {Code: TT, Alpha2: "TT", Alpha3: "TTO", Numeric: "780", Name: "Trinidad and Tobago", Flag: "🇹🇹"},
	TV: {Code: TV, Alpha2: "TV", Alpha3: "TUV", Numeric: "798", Name: "Tuvalu", Flag: "🇹🇻"},
	TW: {Code: TW, Alpha2: "TW", Alpha3: "TWN", Numeric: "158", Name: "Taiwan, Province of China", Flag: "🇹🇼"},
	TZ: {Code: TZ, Alpha2: "TZ", Alpha3: "TZA", Numeric: "834", Name: "Tanzania, United Republic of", Flag: "🇹🇿"},
	UA: {Code: UA, Alpha2: "UA", Alpha3: "UKR", Numeric: "804", Name: "Ukraine", Flag: "🇺🇦"},
	UG: {Code: UG, Alpha2: "UG", Alpha3: "UGA", Numeric: "800", Name: "Uganda", Flag: "🇺🇬"},
	UM: {Code: UM, Alpha2: "UM", Alpha3: "UMI", Numeric: "581", Name: "United States Minor Outlying Islands", Flag: "🇺🇲"},
	US: {Code: US, Alpha2: "US", Alpha3: "USA", Numeric: "840", Name: "United States", Flag: "🇺🇸"},
	UY: {Code: UY, Alpha2: "UY", Alpha3: "URY", Numeric: "858", Name: "Uruguay", Flag: "🇺🇾"},
	UZ: {Code: UZ, Alpha2: "UZ", Alpha3: "UZB", Numeric: "860", Name: "Uzbekistan", Flag: "🇺🇿"},
	VA: {Code: VA, Alpha2: "VA", Alpha3: "VAT", Numeric: "336", Name: "Holy See (Vatican City State)", Flag: "🇻🇦"},
	VC: {Code: VC, Alpha2: "VC", Alpha3: "VCT", Numeric: "670", Name: "Saint Vincent and the Grenadines", Flag: "🇻🇨"},
	VE: {Code: VE, Alpha2: "VE", Alpha3: "VEN", Numeric: "862", Name: "Venezuela, Bolivarian Republic of", Flag: "🇻🇪"},
	VG: {Code: VG, Alpha2: "VG", Alpha3: "VGB", Numeric: "092", Name: "Virgin Islands, British", Flag: "🇻🇬"},
	VI: {Code: VI, Alpha2: "VI", Alpha3: "VIR", Numeric: "850", Name: "Virgin Islands, U.S.", Flag: "🇻🇮"},
	VN: {Code: VN, Alpha2: "VN", Alpha3: "VNM", Numeric: "704", Name: "Viet Nam", Flag: "🇻🇳"},
	VU: {Code: VU, Alpha2: "VU", Alpha3: "VUT", Numeric: "548", Name: "Vanuatu", Flag: "🇻🇺"},
	WF: {Code: WF, Alpha2: "WF", Alpha3: "WLF", Numeric: "876", Name: "Wallis and Futuna", Flag: "🇼🇫"},
	WS: {Code: WS, Alpha2: "WS", Alpha3: "WSM", Numeric: "882", Name: "Samoa", Flag: "🇼🇸"},
	YE: {Code: YE, Alpha2: "YE", Alpha3: "YEM", Numeric: "887", Name: "Yemen", Flag: "🇾🇪"},
	YT: {Code: YT, Alpha2: "YT", Alpha3: "MYT", Numeric: "175", Name: "Mayotte", Flag: "🇾🇹"},
	ZA: {Code: ZA, Alpha2: "ZA", Alpha3: "ZAF", Numeric: "710", Name: "South Africa", Flag: "🇿🇦"},
	ZM: {Code: ZM, Alpha2: "ZM", Alpha3: "ZMB", Numeric: "894", Name: "Zambia", Flag: "🇿🇲"},
	ZW: {Code: ZW, Alpha2: "ZW", Alpha3: "ZWE", Numeric: "716", Name: "Zimbabwe", Flag: "🇿🇼"},
}

var codes = [...]Code{
	AD,
	AE,
	AF,
	AG,
	AI,
	AL,
	AM,
	AO,
	AQ,
	AR,
	AS,
	AT,
	AU,
	AW,
	AX,
	AZ,
	BA,
	BB,
	BD,
	BE,
	BF,
	BG,
	BH,
	BI,
	BJ,
	BL,
	BM,
	BN,
	BO,
	BQ,
	BR,
	BS,
	BT,
	BV,
	BW,
	BY,
	BZ,
	CA,
	CC,
	CD,
	CF,
	CG,
	CH,
	CI,
	CK,
	CL,
	CM,
	CN,
	CO,
	CR,
	CU,
	CV,
	CW,
	CX,
	CY,
	CZ,
	DE,
	DJ,
	DK,
	DM,
	DO,
	DZ,
	EC,
	EE,
	EG,
	EH,
	ER,
	ES,
	ET,
	FI,
	FJ,
	FK,
	FM,
	FO,
	FR,
	GA,
	GB,
	GD,
	GE,
	GF,
	GG,
	GH,
	GI,
	GL,
	GM,
	GN,
	GP,
	GQ,
	GR,
	GS,
	GT,
	GU,
	GW,
	GY,
	HK,
	HM,
	HN,
	HR,
	HT,
	HU,
	ID,
	IE,
	IL,
	IM,
	IN,
	IO,
	IQ,
	IR,
	IS,
	IT,
	JE,
	JM,
	JO,
	JP,
	KE,
	KG,
	KH,
	KI,
	KM,
	KN,
	KP,
	KR,
	KW,
	KY,
	KZ,
	LA,
	LB,
	LC,
	LI,
	LK,
	LR,
	LS,
	LT,
	LU,
	LV,
	LY,
	MA,
	MC,
	MD,
	ME,
	MF,
	MG,
	MH,
	MK,
	ML,
	MM,
	MN,
	MO,
	MP,
	MQ,
	MR,
	MS,
	MT,
	MU,
	MV,
	MW,
	MX,
	MY,
	MZ,
	NA,
	NC,
	NE,
	NF,
	NG,
	NI,
	NL,
	NO,
	NP,
	NR,
	NU,
	NZ,
	OM,
	PA,
	PE,
	PF,
	PG,
	PH,
	PK,
	PL,
	PM,
	PN,
	PR,
	PS,
	PT,
	PW,
	PY,
	QA,
	RE,
	RO,
	RS,
	RU,
	RW,
	SA,
	SB,
	SC,
	SD,
	SE,
	SG,
	SH,
	SI,
	SJ,
	SK,
	SL,
	SM,
	SN,
	SO,
	SR,
	SS,
	ST,
	SV,
	SX,
	SY,
	SZ,
	TC,
	TD,
	TF,
	TG,
	TH,
	TJ,
	TK,
	TL,
	TM,
	TN,
	TO,
	TR,
	TT,
	TV,
	TW,
	TZ,
	UA,
	UG,
	UM,
	US,
	UY,
	UZ,
	VA,
	VC,
	VE,
	VG,
	VI,
	VN,
	VU,
	WF,
	WS,
	YE,
	YT,
	ZA,
	ZM,
	ZW,
}
