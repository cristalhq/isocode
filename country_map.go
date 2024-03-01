package isocode

var countries = map[Country]countryDetails{
	AD: {Code: 20, Alpha2: "AD", Alpha3: "AND", Flag: "🇦🇩", Name: "Andorra"},
	AE: {Code: 784, Alpha2: "AE", Alpha3: "ARE", Flag: "🇦🇪", Name: "United Arab Emirates"},
	AF: {Code: 4, Alpha2: "AF", Alpha3: "AFG", Flag: "🇦🇫", Name: "Afghanistan"},
	AG: {Code: 28, Alpha2: "AG", Alpha3: "ATG", Flag: "🇦🇬", Name: "Antigua and Barbuda"},
	AI: {Code: 660, Alpha2: "AI", Alpha3: "AIA", Flag: "🇦🇮", Name: "Anguilla"},
	AL: {Code: 8, Alpha2: "AL", Alpha3: "ALB", Flag: "🇦🇱", Name: "Albania"},
	AM: {Code: 51, Alpha2: "AM", Alpha3: "ARM", Flag: "🇦🇲", Name: "Armenia"},
	AO: {Code: 24, Alpha2: "AO", Alpha3: "AGO", Flag: "🇦🇴", Name: "Angola"},
	AQ: {Code: 10, Alpha2: "AQ", Alpha3: "ATA", Flag: "🇦🇶", Name: "Antarctica"},
	AR: {Code: 32, Alpha2: "AR", Alpha3: "ARG", Flag: "🇦🇷", Name: "Argentina"},
	AS: {Code: 16, Alpha2: "AS", Alpha3: "ASM", Flag: "🇦🇸", Name: "American Samoa"},
	AT: {Code: 40, Alpha2: "AT", Alpha3: "AUT", Flag: "🇦🇹", Name: "Austria"},
	AU: {Code: 36, Alpha2: "AU", Alpha3: "AUS", Flag: "🇦🇺", Name: "Australia"},
	AW: {Code: 533, Alpha2: "AW", Alpha3: "ABW", Flag: "🇦🇼", Name: "Aruba"},
	AX: {Code: 248, Alpha2: "AX", Alpha3: "ALA", Flag: "🇦🇽", Name: "Åland Islands"},
	AZ: {Code: 31, Alpha2: "AZ", Alpha3: "AZE", Flag: "🇦🇿", Name: "Azerbaijan"},
	BA: {Code: 70, Alpha2: "BA", Alpha3: "BIH", Flag: "🇧🇦", Name: "Bosnia and Herzegovina"},
	BB: {Code: 52, Alpha2: "BB", Alpha3: "BRB", Flag: "🇧🇧", Name: "Barbados"},
	BD: {Code: 50, Alpha2: "BD", Alpha3: "BGD", Flag: "🇧🇩", Name: "Bangladesh"},
	BE: {Code: 56, Alpha2: "BE", Alpha3: "BEL", Flag: "🇧🇪", Name: "Belgium"},
	BF: {Code: 854, Alpha2: "BF", Alpha3: "BFA", Flag: "🇧🇫", Name: "Burkina Faso"},
	BG: {Code: 100, Alpha2: "BG", Alpha3: "BGR", Flag: "🇧🇬", Name: "Bulgaria"},
	BH: {Code: 48, Alpha2: "BH", Alpha3: "BHR", Flag: "🇧🇭", Name: "Bahrain"},
	BI: {Code: 108, Alpha2: "BI", Alpha3: "BDI", Flag: "🇧🇮", Name: "Burundi"},
	BJ: {Code: 204, Alpha2: "BJ", Alpha3: "BEN", Flag: "🇧🇯", Name: "Benin"},
	BL: {Code: 652, Alpha2: "BL", Alpha3: "BLM", Flag: "🇧🇱", Name: "Saint Barthélemy"},
	BM: {Code: 60, Alpha2: "BM", Alpha3: "BMU", Flag: "🇧🇲", Name: "Bermuda"},
	BN: {Code: 96, Alpha2: "BN", Alpha3: "BRN", Flag: "🇧🇳", Name: "Brunei Darussalam"},
	BO: {Code: 68, Alpha2: "BO", Alpha3: "BOL", Flag: "🇧🇴", Name: "Bolivia (Plurinational State of)"},
	BQ: {Code: 535, Alpha2: "BQ", Alpha3: "BES", Flag: "🇧🇶", Name: "Bonaire, Sint Eustatius and Saba"},
	BR: {Code: 76, Alpha2: "BR", Alpha3: "BRA", Flag: "🇧🇷", Name: "Brazil"},
	BS: {Code: 44, Alpha2: "BS", Alpha3: "BHS", Flag: "🇧🇸", Name: "Bahamas"},
	BT: {Code: 64, Alpha2: "BT", Alpha3: "BTN", Flag: "🇧🇹", Name: "Bhutan"},
	BV: {Code: 74, Alpha2: "BV", Alpha3: "BVT", Flag: "🇧🇻", Name: "Bouvet Island"},
	BW: {Code: 72, Alpha2: "BW", Alpha3: "BWA", Flag: "🇧🇼", Name: "Botswana"},
	BY: {Code: 112, Alpha2: "BY", Alpha3: "BLR", Flag: "🇧🇾", Name: "Belarus"},
	BZ: {Code: 84, Alpha2: "BZ", Alpha3: "BLZ", Flag: "🇧🇿", Name: "Belize"},
	CA: {Code: 124, Alpha2: "CA", Alpha3: "CAN", Flag: "🇨🇦", Name: "Canada"},
	CC: {Code: 166, Alpha2: "CC", Alpha3: "CCK", Flag: "🇨🇨", Name: "Cocos (Keeling) Islands"},
	CD: {Code: 180, Alpha2: "CD", Alpha3: "COD", Flag: "🇨🇩", Name: "Congo, Democratic Republic of the"},
	CF: {Code: 140, Alpha2: "CF", Alpha3: "CAF", Flag: "🇨🇫", Name: "Central African Republic"},
	CG: {Code: 178, Alpha2: "CG", Alpha3: "COG", Flag: "🇨🇬", Name: "Congo"},
	CH: {Code: 756, Alpha2: "CH", Alpha3: "CHE", Flag: "🇨🇭", Name: "Switzerland"},
	CI: {Code: 384, Alpha2: "CI", Alpha3: "CIV", Flag: "🇨🇮", Name: "Côte d'Ivoire"},
	CK: {Code: 184, Alpha2: "CK", Alpha3: "COK", Flag: "🇨🇰", Name: "Cook Islands"},
	CL: {Code: 152, Alpha2: "CL", Alpha3: "CHL", Flag: "🇨🇱", Name: "Chile"},
	CM: {Code: 120, Alpha2: "CM", Alpha3: "CMR", Flag: "🇨🇲", Name: "Cameroon"},
	CN: {Code: 156, Alpha2: "CN", Alpha3: "CHN", Flag: "🇨🇳", Name: "China"},
	CO: {Code: 170, Alpha2: "CO", Alpha3: "COL", Flag: "🇨🇴", Name: "Colombia"},
	CR: {Code: 188, Alpha2: "CR", Alpha3: "CRI", Flag: "🇨🇷", Name: "Costa Rica"},
	CU: {Code: 192, Alpha2: "CU", Alpha3: "CUB", Flag: "🇨🇺", Name: "Cuba"},
	CV: {Code: 132, Alpha2: "CV", Alpha3: "CPV", Flag: "🇨🇻", Name: "Cabo Verde"},
	CW: {Code: 531, Alpha2: "CW", Alpha3: "CUW", Flag: "🇨🇼", Name: "Curaçao"},
	CX: {Code: 162, Alpha2: "CX", Alpha3: "CXR", Flag: "🇨🇽", Name: "Christmas Island"},
	CY: {Code: 196, Alpha2: "CY", Alpha3: "CYP", Flag: "🇨🇾", Name: "Cyprus"},
	CZ: {Code: 203, Alpha2: "CZ", Alpha3: "CZE", Flag: "🇨🇿", Name: "Czechia"},
	DE: {Code: 276, Alpha2: "DE", Alpha3: "DEU", Flag: "🇩🇪", Name: "Germany"},
	DJ: {Code: 262, Alpha2: "DJ", Alpha3: "DJI", Flag: "🇩🇯", Name: "Djibouti"},
	DK: {Code: 208, Alpha2: "DK", Alpha3: "DNK", Flag: "🇩🇰", Name: "Denmark"},
	DM: {Code: 212, Alpha2: "DM", Alpha3: "DMA", Flag: "🇩🇲", Name: "Dominica"},
	DO: {Code: 214, Alpha2: "DO", Alpha3: "DOM", Flag: "🇩🇴", Name: "Dominican Republic"},
	DZ: {Code: 12, Alpha2: "DZ", Alpha3: "DZA", Flag: "🇩🇿", Name: "Algeria"},
	EC: {Code: 218, Alpha2: "EC", Alpha3: "ECU", Flag: "🇪🇨", Name: "Ecuador"},
	EE: {Code: 233, Alpha2: "EE", Alpha3: "EST", Flag: "🇪🇪", Name: "Estonia"},
	EG: {Code: 818, Alpha2: "EG", Alpha3: "EGY", Flag: "🇪🇬", Name: "Egypt"},
	EH: {Code: 732, Alpha2: "EH", Alpha3: "ESH", Flag: "🇪🇭", Name: "Western Sahara"},
	ER: {Code: 232, Alpha2: "ER", Alpha3: "ERI", Flag: "🇪🇷", Name: "Eritrea"},
	ES: {Code: 724, Alpha2: "ES", Alpha3: "ESP", Flag: "🇪🇸", Name: "Spain"},
	ET: {Code: 231, Alpha2: "ET", Alpha3: "ETH", Flag: "🇪🇹", Name: "Ethiopia"},
	FI: {Code: 246, Alpha2: "FI", Alpha3: "FIN", Flag: "🇫🇮", Name: "Finland"},
	FJ: {Code: 242, Alpha2: "FJ", Alpha3: "FJI", Flag: "🇫🇯", Name: "Fiji"},
	FK: {Code: 238, Alpha2: "FK", Alpha3: "FLK", Flag: "🇫🇰", Name: "Falkland Islands (Malvinas)"},
	FM: {Code: 583, Alpha2: "FM", Alpha3: "FSM", Flag: "🇫🇲", Name: "Micronesia (Federated States of)"},
	FO: {Code: 234, Alpha2: "FO", Alpha3: "FRO", Flag: "🇫🇴", Name: "Faroe Islands"},
	FR: {Code: 250, Alpha2: "FR", Alpha3: "FRA", Flag: "🇫🇷", Name: "France"},
	GA: {Code: 266, Alpha2: "GA", Alpha3: "GAB", Flag: "🇬🇦", Name: "Gabon"},
	GB: {Code: 826, Alpha2: "GB", Alpha3: "GBR", Flag: "🇬🇧", Name: "United Kingdom of Great Britain and Northern Ireland"},
	GD: {Code: 308, Alpha2: "GD", Alpha3: "GRD", Flag: "🇬🇩", Name: "Grenada"},
	GE: {Code: 268, Alpha2: "GE", Alpha3: "GEO", Flag: "🇬🇪", Name: "Georgia"},
	GF: {Code: 254, Alpha2: "GF", Alpha3: "GUF", Flag: "🇬🇫", Name: "French Guiana"},
	GG: {Code: 831, Alpha2: "GG", Alpha3: "GGY", Flag: "🇬🇬", Name: "Guernsey"},
	GH: {Code: 288, Alpha2: "GH", Alpha3: "GHA", Flag: "🇬🇭", Name: "Ghana"},
	GI: {Code: 292, Alpha2: "GI", Alpha3: "GIB", Flag: "🇬🇮", Name: "Gibraltar"},
	GL: {Code: 304, Alpha2: "GL", Alpha3: "GRL", Flag: "🇬🇱", Name: "Greenland"},
	GM: {Code: 270, Alpha2: "GM", Alpha3: "GMB", Flag: "🇬🇲", Name: "Gambia"},
	GN: {Code: 324, Alpha2: "GN", Alpha3: "GIN", Flag: "🇬🇳", Name: "Guinea"},
	GP: {Code: 312, Alpha2: "GP", Alpha3: "GLP", Flag: "🇬🇵", Name: "Guadeloupe"},
	GQ: {Code: 226, Alpha2: "GQ", Alpha3: "GNQ", Flag: "🇬🇶", Name: "Equatorial Guinea"},
	GR: {Code: 300, Alpha2: "GR", Alpha3: "GRC", Flag: "🇬🇷", Name: "Greece"},
	GS: {Code: 239, Alpha2: "GS", Alpha3: "SGS", Flag: "🇬🇸", Name: "South Georgia and the South Sandwich Islands"},
	GT: {Code: 320, Alpha2: "GT", Alpha3: "GTM", Flag: "🇬🇹", Name: "Guatemala"},
	GU: {Code: 316, Alpha2: "GU", Alpha3: "GUM", Flag: "🇬🇺", Name: "Guam"},
	GW: {Code: 624, Alpha2: "GW", Alpha3: "GNB", Flag: "🇬🇼", Name: "Guinea-Bissau"},
	GY: {Code: 328, Alpha2: "GY", Alpha3: "GUY", Flag: "🇬🇾", Name: "Guyana"},
	HK: {Code: 344, Alpha2: "HK", Alpha3: "HKG", Flag: "🇭🇰", Name: "Hong Kong"},
	HM: {Code: 334, Alpha2: "HM", Alpha3: "HMD", Flag: "🇭🇲", Name: "Heard Island and McDonald Islands"},
	HN: {Code: 340, Alpha2: "HN", Alpha3: "HND", Flag: "🇭🇳", Name: "Honduras"},
	HR: {Code: 191, Alpha2: "HR", Alpha3: "HRV", Flag: "🇭🇷", Name: "Croatia"},
	HT: {Code: 332, Alpha2: "HT", Alpha3: "HTI", Flag: "🇭🇹", Name: "Haiti"},
	HU: {Code: 348, Alpha2: "HU", Alpha3: "HUN", Flag: "🇭🇺", Name: "Hungary"},
	ID: {Code: 360, Alpha2: "ID", Alpha3: "IDN", Flag: "🇮🇩", Name: "Indonesia"},
	IE: {Code: 372, Alpha2: "IE", Alpha3: "IRL", Flag: "🇮🇪", Name: "Ireland"},
	IL: {Code: 376, Alpha2: "IL", Alpha3: "ISR", Flag: "🇮🇱", Name: "Israel"},
	IM: {Code: 833, Alpha2: "IM", Alpha3: "IMN", Flag: "🇮🇲", Name: "Isle of Man"},
	IN: {Code: 356, Alpha2: "IN", Alpha3: "IND", Flag: "🇮🇳", Name: "India"},
	IO: {Code: 86, Alpha2: "IO", Alpha3: "IOT", Flag: "🇮🇴", Name: "British Indian Ocean Territory"},
	IQ: {Code: 368, Alpha2: "IQ", Alpha3: "IRQ", Flag: "🇮🇶", Name: "Iraq"},
	IR: {Code: 364, Alpha2: "IR", Alpha3: "IRN", Flag: "🇮🇷", Name: "Iran (Islamic Republic of)"},
	IS: {Code: 352, Alpha2: "IS", Alpha3: "ISL", Flag: "🇮🇸", Name: "Iceland"},
	IT: {Code: 380, Alpha2: "IT", Alpha3: "ITA", Flag: "🇮🇹", Name: "Italy"},
	JE: {Code: 832, Alpha2: "JE", Alpha3: "JEY", Flag: "🇯🇪", Name: "Jersey"},
	JM: {Code: 388, Alpha2: "JM", Alpha3: "JAM", Flag: "🇯🇲", Name: "Jamaica"},
	JO: {Code: 400, Alpha2: "JO", Alpha3: "JOR", Flag: "🇯🇴", Name: "Jordan"},
	JP: {Code: 392, Alpha2: "JP", Alpha3: "JPN", Flag: "🇯🇵", Name: "Japan"},
	KE: {Code: 404, Alpha2: "KE", Alpha3: "KEN", Flag: "🇰🇪", Name: "Kenya"},
	KG: {Code: 417, Alpha2: "KG", Alpha3: "KGZ", Flag: "🇰🇬", Name: "Kyrgyzstan"},
	KH: {Code: 116, Alpha2: "KH", Alpha3: "KHM", Flag: "🇰🇭", Name: "Cambodia"},
	KI: {Code: 296, Alpha2: "KI", Alpha3: "KIR", Flag: "🇰🇮", Name: "Kiribati"},
	KM: {Code: 174, Alpha2: "KM", Alpha3: "COM", Flag: "🇰🇲", Name: "Comoros"},
	KN: {Code: 659, Alpha2: "KN", Alpha3: "KNA", Flag: "🇰🇳", Name: "Saint Kitts and Nevis"},
	KP: {Code: 408, Alpha2: "KP", Alpha3: "PRK", Flag: "🇰🇵", Name: "Korea (Democratic People's Republic of)"},
	KR: {Code: 410, Alpha2: "KR", Alpha3: "KOR", Flag: "🇰🇷", Name: "Korea, Republic of"},
	KW: {Code: 414, Alpha2: "KW", Alpha3: "KWT", Flag: "🇰🇼", Name: "Kuwait"},
	KY: {Code: 136, Alpha2: "KY", Alpha3: "CYM", Flag: "🇰🇾", Name: "Cayman Islands"},
	KZ: {Code: 398, Alpha2: "KZ", Alpha3: "KAZ", Flag: "🇰🇿", Name: "Kazakhstan"},
	LA: {Code: 418, Alpha2: "LA", Alpha3: "LAO", Flag: "🇱🇦", Name: "Lao People's Democratic Republic"},
	LB: {Code: 422, Alpha2: "LB", Alpha3: "LBN", Flag: "🇱🇧", Name: "Lebanon"},
	LC: {Code: 662, Alpha2: "LC", Alpha3: "LCA", Flag: "🇱🇨", Name: "Saint Lucia"},
	LI: {Code: 438, Alpha2: "LI", Alpha3: "LIE", Flag: "🇱🇮", Name: "Liechtenstein"},
	LK: {Code: 144, Alpha2: "LK", Alpha3: "LKA", Flag: "🇱🇰", Name: "Sri Lanka"},
	LR: {Code: 430, Alpha2: "LR", Alpha3: "LBR", Flag: "🇱🇷", Name: "Liberia"},
	LS: {Code: 426, Alpha2: "LS", Alpha3: "LSO", Flag: "🇱🇸", Name: "Lesotho"},
	LT: {Code: 440, Alpha2: "LT", Alpha3: "LTU", Flag: "🇱🇹", Name: "Lithuania"},
	LU: {Code: 442, Alpha2: "LU", Alpha3: "LUX", Flag: "🇱🇺", Name: "Luxembourg"},
	LV: {Code: 428, Alpha2: "LV", Alpha3: "LVA", Flag: "🇱🇻", Name: "Latvia"},
	LY: {Code: 434, Alpha2: "LY", Alpha3: "LBY", Flag: "🇱🇾", Name: "Libya"},
	MA: {Code: 504, Alpha2: "MA", Alpha3: "MAR", Flag: "🇲🇦", Name: "Morocco"},
	MC: {Code: 492, Alpha2: "MC", Alpha3: "MCO", Flag: "🇲🇨", Name: "Monaco"},
	MD: {Code: 498, Alpha2: "MD", Alpha3: "MDA", Flag: "🇲🇩", Name: "Moldova, Republic of"},
	ME: {Code: 499, Alpha2: "ME", Alpha3: "MNE", Flag: "🇲🇪", Name: "Montenegro"},
	MF: {Code: 663, Alpha2: "MF", Alpha3: "MAF", Flag: "🇲🇫", Name: "Saint Martin (French part)"},
	MG: {Code: 450, Alpha2: "MG", Alpha3: "MDG", Flag: "🇲🇬", Name: "Madagascar"},
	MH: {Code: 584, Alpha2: "MH", Alpha3: "MHL", Flag: "🇲🇭", Name: "Marshall Islands"},
	MK: {Code: 807, Alpha2: "MK", Alpha3: "MKD", Flag: "🇲🇰", Name: "North Macedonia"},
	ML: {Code: 466, Alpha2: "ML", Alpha3: "MLI", Flag: "🇲🇱", Name: "Mali"},
	MM: {Code: 104, Alpha2: "MM", Alpha3: "MMR", Flag: "🇲🇲", Name: "Myanmar"},
	MN: {Code: 496, Alpha2: "MN", Alpha3: "MNG", Flag: "🇲🇳", Name: "Mongolia"},
	MO: {Code: 446, Alpha2: "MO", Alpha3: "MAC", Flag: "🇲🇴", Name: "Macao"},
	MP: {Code: 580, Alpha2: "MP", Alpha3: "MNP", Flag: "🇲🇵", Name: "Northern Mariana Islands"},
	MQ: {Code: 474, Alpha2: "MQ", Alpha3: "MTQ", Flag: "🇲🇶", Name: "Martinique"},
	MR: {Code: 478, Alpha2: "MR", Alpha3: "MRT", Flag: "🇲🇷", Name: "Mauritania"},
	MS: {Code: 500, Alpha2: "MS", Alpha3: "MSR", Flag: "🇲🇸", Name: "Montserrat"},
	MT: {Code: 470, Alpha2: "MT", Alpha3: "MLT", Flag: "🇲🇹", Name: "Malta"},
	MU: {Code: 480, Alpha2: "MU", Alpha3: "MUS", Flag: "🇲🇺", Name: "Mauritius"},
	MV: {Code: 462, Alpha2: "MV", Alpha3: "MDV", Flag: "🇲🇻", Name: "Maldives"},
	MW: {Code: 454, Alpha2: "MW", Alpha3: "MWI", Flag: "🇲🇼", Name: "Malawi"},
	MX: {Code: 484, Alpha2: "MX", Alpha3: "MEX", Flag: "🇲🇽", Name: "Mexico"},
	MY: {Code: 458, Alpha2: "MY", Alpha3: "MYS", Flag: "🇲🇾", Name: "Malaysia"},
	MZ: {Code: 508, Alpha2: "MZ", Alpha3: "MOZ", Flag: "🇲🇿", Name: "Mozambique"},
	NA: {Code: 516, Alpha2: "NA", Alpha3: "NAM", Flag: "🇳🇦", Name: "Namibia"},
	NC: {Code: 540, Alpha2: "NC", Alpha3: "NCL", Flag: "🇳🇨", Name: "New Caledonia"},
	NE: {Code: 562, Alpha2: "NE", Alpha3: "NER", Flag: "🇳🇪", Name: "Niger"},
	NF: {Code: 574, Alpha2: "NF", Alpha3: "NFK", Flag: "🇳🇫", Name: "Norfolk Island"},
	NG: {Code: 566, Alpha2: "NG", Alpha3: "NGA", Flag: "🇳🇬", Name: "Nigeria"},
	NI: {Code: 558, Alpha2: "NI", Alpha3: "NIC", Flag: "🇳🇮", Name: "Nicaragua"},
	NL: {Code: 528, Alpha2: "NL", Alpha3: "NLD", Flag: "🇳🇱", Name: "Netherlands"},
	NO: {Code: 578, Alpha2: "NO", Alpha3: "NOR", Flag: "🇳🇴", Name: "Norway"},
	NP: {Code: 524, Alpha2: "NP", Alpha3: "NPL", Flag: "🇳🇵", Name: "Nepal"},
	NR: {Code: 520, Alpha2: "NR", Alpha3: "NRU", Flag: "🇳🇷", Name: "Nauru"},
	NU: {Code: 570, Alpha2: "NU", Alpha3: "NIU", Flag: "🇳🇺", Name: "Niue"},
	NZ: {Code: 554, Alpha2: "NZ", Alpha3: "NZL", Flag: "🇳🇿", Name: "New Zealand"},
	OM: {Code: 512, Alpha2: "OM", Alpha3: "OMN", Flag: "🇴🇲", Name: "Oman"},
	PA: {Code: 591, Alpha2: "PA", Alpha3: "PAN", Flag: "🇵🇦", Name: "Panama"},
	PE: {Code: 604, Alpha2: "PE", Alpha3: "PER", Flag: "🇵🇪", Name: "Peru"},
	PF: {Code: 258, Alpha2: "PF", Alpha3: "PYF", Flag: "🇵🇫", Name: "French Polynesia"},
	PG: {Code: 598, Alpha2: "PG", Alpha3: "PNG", Flag: "🇵🇬", Name: "Papua New Guinea"},
	PH: {Code: 608, Alpha2: "PH", Alpha3: "PHL", Flag: "🇵🇭", Name: "Philippines"},
	PK: {Code: 586, Alpha2: "PK", Alpha3: "PAK", Flag: "🇵🇰", Name: "Pakistan"},
	PL: {Code: 616, Alpha2: "PL", Alpha3: "POL", Flag: "🇵🇱", Name: "Poland"},
	PM: {Code: 666, Alpha2: "PM", Alpha3: "SPM", Flag: "🇵🇲", Name: "Saint Pierre and Miquelon"},
	PN: {Code: 612, Alpha2: "PN", Alpha3: "PCN", Flag: "🇵🇳", Name: "Pitcairn"},
	PR: {Code: 630, Alpha2: "PR", Alpha3: "PRI", Flag: "🇵🇷", Name: "Puerto Rico"},
	PS: {Code: 275, Alpha2: "PS", Alpha3: "PSE", Flag: "🇵🇸", Name: "Palestine, State of"},
	PT: {Code: 620, Alpha2: "PT", Alpha3: "PRT", Flag: "🇵🇹", Name: "Portugal"},
	PW: {Code: 585, Alpha2: "PW", Alpha3: "PLW", Flag: "🇵🇼", Name: "Palau"},
	PY: {Code: 600, Alpha2: "PY", Alpha3: "PRY", Flag: "🇵🇾", Name: "Paraguay"},
	QA: {Code: 634, Alpha2: "QA", Alpha3: "QAT", Flag: "🇶🇦", Name: "Qatar"},
	RE: {Code: 638, Alpha2: "RE", Alpha3: "REU", Flag: "🇷🇪", Name: "Réunion"},
	RO: {Code: 642, Alpha2: "RO", Alpha3: "ROU", Flag: "🇷🇴", Name: "Romania"},
	RS: {Code: 688, Alpha2: "RS", Alpha3: "SRB", Flag: "🇷🇸", Name: "Serbia"},
	RU: {Code: 643, Alpha2: "RU", Alpha3: "RUS", Flag: "🇷🇺", Name: "Russian Federation"},
	RW: {Code: 646, Alpha2: "RW", Alpha3: "RWA", Flag: "🇷🇼", Name: "Rwanda"},
	SA: {Code: 682, Alpha2: "SA", Alpha3: "SAU", Flag: "🇸🇦", Name: "Saudi Arabia"},
	SB: {Code: 90, Alpha2: "SB", Alpha3: "SLB", Flag: "🇸🇧", Name: "Solomon Islands"},
	SC: {Code: 690, Alpha2: "SC", Alpha3: "SYC", Flag: "🇸🇨", Name: "Seychelles"},
	SD: {Code: 729, Alpha2: "SD", Alpha3: "SDN", Flag: "🇸🇩", Name: "Sudan"},
	SE: {Code: 752, Alpha2: "SE", Alpha3: "SWE", Flag: "🇸🇪", Name: "Sweden"},
	SG: {Code: 702, Alpha2: "SG", Alpha3: "SGP", Flag: "🇸🇬", Name: "Singapore"},
	SH: {Code: 654, Alpha2: "SH", Alpha3: "SHN", Flag: "🇸🇭", Name: "Saint Helena, Ascension and Tristan da Cunha"},
	SI: {Code: 705, Alpha2: "SI", Alpha3: "SVN", Flag: "🇸🇮", Name: "Slovenia"},
	SJ: {Code: 744, Alpha2: "SJ", Alpha3: "SJM", Flag: "🇸🇯", Name: "Svalbard and Jan Mayen"},
	SK: {Code: 703, Alpha2: "SK", Alpha3: "SVK", Flag: "🇸🇰", Name: "Slovakia"},
	SL: {Code: 694, Alpha2: "SL", Alpha3: "SLE", Flag: "🇸🇱", Name: "Sierra Leone"},
	SM: {Code: 674, Alpha2: "SM", Alpha3: "SMR", Flag: "🇸🇲", Name: "San Marino"},
	SN: {Code: 686, Alpha2: "SN", Alpha3: "SEN", Flag: "🇸🇳", Name: "Senegal"},
	SO: {Code: 706, Alpha2: "SO", Alpha3: "SOM", Flag: "🇸🇴", Name: "Somalia"},
	SR: {Code: 740, Alpha2: "SR", Alpha3: "SUR", Flag: "🇸🇷", Name: "Suriname"},
	SS: {Code: 728, Alpha2: "SS", Alpha3: "SSD", Flag: "🇸🇸", Name: "South Sudan"},
	ST: {Code: 678, Alpha2: "ST", Alpha3: "STP", Flag: "🇸🇹", Name: "Sao Tome and Principe"},
	SV: {Code: 222, Alpha2: "SV", Alpha3: "SLV", Flag: "🇸🇻", Name: "El Salvador"},
	SX: {Code: 534, Alpha2: "SX", Alpha3: "SXM", Flag: "🇸🇽", Name: "Sint Maarten (Dutch part)"},
	SY: {Code: 760, Alpha2: "SY", Alpha3: "SYR", Flag: "🇸🇾", Name: "Syrian Arab Republic"},
	SZ: {Code: 748, Alpha2: "SZ", Alpha3: "SWZ", Flag: "🇸🇿", Name: "Eswatini"},
	TC: {Code: 796, Alpha2: "TC", Alpha3: "TCA", Flag: "🇹🇨", Name: "Turks and Caicos Islands"},
	TD: {Code: 148, Alpha2: "TD", Alpha3: "TCD", Flag: "🇹🇩", Name: "Chad"},
	TF: {Code: 260, Alpha2: "TF", Alpha3: "ATF", Flag: "🇹🇫", Name: "French Southern Territories"},
	TG: {Code: 768, Alpha2: "TG", Alpha3: "TGO", Flag: "🇹🇬", Name: "Togo"},
	TH: {Code: 764, Alpha2: "TH", Alpha3: "THA", Flag: "🇹🇭", Name: "Thailand"},
	TJ: {Code: 762, Alpha2: "TJ", Alpha3: "TJK", Flag: "🇹🇯", Name: "Tajikistan"},
	TK: {Code: 772, Alpha2: "TK", Alpha3: "TKL", Flag: "🇹🇰", Name: "Tokelau"},
	TL: {Code: 626, Alpha2: "TL", Alpha3: "TLS", Flag: "🇹🇱", Name: "Timor-Leste"},
	TM: {Code: 795, Alpha2: "TM", Alpha3: "TKM", Flag: "🇹🇲", Name: "Turkmenistan"},
	TN: {Code: 788, Alpha2: "TN", Alpha3: "TUN", Flag: "🇹🇳", Name: "Tunisia"},
	TO: {Code: 776, Alpha2: "TO", Alpha3: "TON", Flag: "🇹🇴", Name: "Tonga"},
	TR: {Code: 792, Alpha2: "TR", Alpha3: "TUR", Flag: "🇹🇷", Name: "Turkey"},
	TT: {Code: 780, Alpha2: "TT", Alpha3: "TTO", Flag: "🇹🇹", Name: "Trinidad and Tobago"},
	TV: {Code: 798, Alpha2: "TV", Alpha3: "TUV", Flag: "🇹🇻", Name: "Tuvalu"},
	TW: {Code: 158, Alpha2: "TW", Alpha3: "TWN", Flag: "🇹🇼", Name: "Taiwan, Province of China"},
	TZ: {Code: 834, Alpha2: "TZ", Alpha3: "TZA", Flag: "🇹🇿", Name: "Tanzania, United Republic of"},
	UA: {Code: 804, Alpha2: "UA", Alpha3: "UKR", Flag: "🇺🇦", Name: "Ukraine"},
	UG: {Code: 800, Alpha2: "UG", Alpha3: "UGA", Flag: "🇺🇬", Name: "Uganda"},
	UM: {Code: 581, Alpha2: "UM", Alpha3: "UMI", Flag: "🇺🇲", Name: "United States Minor Outlying Islands"},
	US: {Code: 840, Alpha2: "US", Alpha3: "USA", Flag: "🇺🇸", Name: "United States of America"},
	UY: {Code: 858, Alpha2: "UY", Alpha3: "URY", Flag: "🇺🇾", Name: "Uruguay"},
	UZ: {Code: 860, Alpha2: "UZ", Alpha3: "UZB", Flag: "🇺🇿", Name: "Uzbekistan"},
	VA: {Code: 336, Alpha2: "VA", Alpha3: "VAT", Flag: "🇻🇦", Name: "Holy See"},
	VC: {Code: 670, Alpha2: "VC", Alpha3: "VCT", Flag: "🇻🇨", Name: "Saint Vincent and the Grenadines"},
	VE: {Code: 862, Alpha2: "VE", Alpha3: "VEN", Flag: "🇻🇪", Name: "Venezuela (Bolivarian Republic of)"},
	VG: {Code: 92, Alpha2: "VG", Alpha3: "VGB", Flag: "🇻🇬", Name: "Virgin Islands (British)"},
	VI: {Code: 850, Alpha2: "VI", Alpha3: "VIR", Flag: "🇻🇮", Name: "Virgin Islands (U.S.)"},
	VN: {Code: 704, Alpha2: "VN", Alpha3: "VNM", Flag: "🇻🇳", Name: "Viet Nam"},
	VU: {Code: 548, Alpha2: "VU", Alpha3: "VUT", Flag: "🇻🇺", Name: "Vanuatu"},
	WF: {Code: 876, Alpha2: "WF", Alpha3: "WLF", Flag: "🇼🇫", Name: "Wallis and Futuna"},
	WS: {Code: 882, Alpha2: "WS", Alpha3: "WSM", Flag: "🇼🇸", Name: "Samoa"},
	YE: {Code: 887, Alpha2: "YE", Alpha3: "YEM", Flag: "🇾🇪", Name: "Yemen"},
	YT: {Code: 175, Alpha2: "YT", Alpha3: "MYT", Flag: "🇾🇹", Name: "Mayotte"},
	ZA: {Code: 710, Alpha2: "ZA", Alpha3: "ZAF", Flag: "🇿🇦", Name: "South Africa"},
	ZM: {Code: 894, Alpha2: "ZM", Alpha3: "ZMB", Flag: "🇿🇲", Name: "Zambia"},
	ZW: {Code: 716, Alpha2: "ZW", Alpha3: "ZWE", Flag: "🇿🇼", Name: "Zimbabwe"},
}
