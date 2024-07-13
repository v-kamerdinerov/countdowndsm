package main

var defaultFont = Font{
	':': {
		"   ",
		"██╗",
		"╚═╝",
		"██╗",
		"╚═╝",
		"   ",
	},
	'0': {
		" ██████╗ ",
		"██╔═████╗",
		"██║██╔██║",
		"████╔╝██║",
		"╚██████╔╝",
		" ╚═════╝ ",
	},
	'1': {
		" ██╗",
		"███║",
		"╚██║",
		" ██║",
		" ██║",
		" ╚═╝",
	},
	'2': {
		"██████╗ ",
		"╚════██╗",
		" █████╔╝",
		"██╔═══╝ ",
		"███████╗",
		"╚══════╝",
	},
	'3': {
		"██████╗ ",
		"╚════██╗",
		" █████╔╝",
		" ╚═══██╗",
		"██████╔╝",
		"╚═════╝ ",
	},
	'4': {
		"██╗  ██╗",
		"██║  ██║",
		"███████║",
		"╚════██║",
		"     ██║",
		"     ╚═╝",
	},
	'5': {
		"███████╗",
		"██╔════╝",
		"███████╗",
		"╚════██║",
		"███████║",
		"╚══════╝",
	},
	'6': {
		" ██████╗ ",
		"██╔════╝ ",
		"███████╗ ",
		"██╔═══██╗",
		"╚██████╔╝",
		" ╚═════╝ ",
	},
	'7': {
		"███████╗",
		"╚════██║",
		"    ██╔╝",
		"   ██╔╝ ",
		"   ██║  ",
		"   ╚═╝  ",
	},
	'8': {
		" █████╗ ",
		"██╔══██╗",
		"╚█████╔╝",
		"██╔══██╗",
		"╚█████╔╝",
		" ╚════╝ ",
	},
	'9': {
		" █████╗ ",
		"██╔══██╗",
		"╚██████║",
		" ╚═══██║",
		" █████╔╝",
		" ╚════╝ ",
	},
}

var defaultFontSmall = Font{
	'a': {
		"▄▀█ ",
		"█▀█ ",
	},
	'á': {
		"▄▀█ ",
		"█▀█ ",
	},
	'b': {
		"█▄▄ ",
		"█▄█ ",
	},
	'c': {
		"█▀▀ ",
		"█▄▄ ",
	},
	'd': {
		"█▀▄ ",
		"█▄▀ ",
	},
	'é': {
		"█▀▀ ",
		"██▄ ",
	},
	'e': {
		"█▀▀ ",
		"██▄ ",
	},
	'f': {
		"█▀▀ ",
		"█▀░ ",
	},
	'g': {
		"█▀▀ ",
		"█▄█ ",
	},
	'h': {
		"█░█ ",
		"█▀█ ",
	},
	'i': {
		"█ ",
		"█ ",
	},
	'í': {
		"█ ",
		"█ ",
	},
	'j': {
		"░░█ ",
		"█▄█ ",
	},
	'k': {
		"█▄▀ ",
		"█░█ ",
	},
	'l': {
		"█░░ ",
		"█▄▄ ",
	},
	'm': {
		"█▀▄▀█ ",
		"█░▀░█ ",
	},
	'n': {
		"█▄░█ ",
		"█░▀█  ",
	},
	'o': {
		"█▀█ ",
		"█▄█ ",
	},
	'p': {
		"█▀█ ",
		"█▀▀ ",
	},
	'q': {
		"█▀█ ",
		"▀▀█ ",
	},
	'r': {
		"█▀█ ",
		"█▀▄ ",
	},
	's': {
		"█▀ ",
		"▄█ ",
	},
	't': {
		"▀█▀ ",
		"░█░ ",
	},
	'u': {
		"█░█ ",
		"█▄█ ",
	},
	'ü': {
		"█░█ ",
		"█▄█ ",
	},
	'v': {
		"█░█ ",
		"▀▄▀ ",
	},
	'w': {
		"█░█░█ ",
		"▀▄▀▄▀ ",
	},
	'x': {
		"▀▄▀ ",
		"█░█ ",
	},
	'y': {
		"█▄█ ",
		"░█░ ",
	},
	'z': {
		"▀█ ",
		"█▄ ",
	},
	'а': {
		"▄▀█ ",
		"█▀█ ",
	},
	'б': {
		"█▀▀░ ",
		"██▄█ ",
	},
	'в': {
		"█▀█░ ",
		"██▄█ ",
	},
	'г': {
		"█▀█ ",
		"█░░ ",
	},
	'д': {
		"░█▀█░ ",
		"▄█▄█▄ ",
	},
	'е': {
		"█▀▀ ",
		"██▄ ",
	},
	'ё': {
		"█▘▘ ",
		"██▄ ",
	},
	'ж': {
		"▀▄█▄▀ ",
		"▄▀█▀▄ ",
	},
	'з': {
		"▀█▙ ",
		"▄█▛ ",
	},
	'и': {
		"█░▄█ ",
		"█▀░█ ",
	},
	'й': {
		"█░▙█ ",
		"█▀░█ ",
	},
	'к': {
		"█▄▀ ",
		"█░█ ",
	},
	'л': {
		"▄▀▄ ",
		"█░█ ",
	},
	'м': {
		"█▀▄▀█ ",
		"█░▀░█ ",
	},
	'н': {
		"█░█ ",
		"█▀█ ",
	},
	'о': {
		"█▀█ ",
		"█▄█ ",
	},
	'п': {
		"█▀█ ",
		"█░█ ",
	},
	'р': {
		"█▀█ ",
		"█▀▀ ",
	},
	'с': {
		"█▀▀ ",
		"█▄▄ ",
	},
	'т': {
		"▀█▀ ",
		"░█░ ",
	},
	'у': {
		"█░█ ",
		"▄▀░ ",
	},
	'ф': {
		"█▜▛█ ",
		"▀██▀ ",
	},
	'х': {
		"▀▄▀ ",
		"█░█ ",
	},
	'ц': {
		"█░█░ ",
		"█▄█▄ ",
	},
	'ч': {
		"█░█ ",
		"▀▀█ ",
	},
	'ш': {
		"█░█░█ ",
		"█▄█▄█ ",
	},
	'щ': {
		"█░█░█░ ",
		"█▄█▄█▄ ",
	},
	'ъ': {
		"▀█▄▄ ",
		"░█▄█ ",
	},
	'ы': {
		"█▄▄░█ ",
		"█▄█░█ ",
	},
	'ь': {
		"█▄▄ ",
		"█▄█ ",
	},
	'э': {
		"▀▜▙ ",
		"▄▟▛ ",
	},
	'ю': {
		"█▄█▀█ ",
		"█░█▄█ ",
	},
	'я': {
		"█▀█ ",
		"▄▀█ ",
	},
	'1': {
		"▄█ ",
		"░█ ",
	},
	'2': {
		"▀█ ",
		"█▄ ",
	},
	'3': {
		"▜█ ",
		"▟█ ",
	},
	'4': {
		"█░█ ",
		"▀▀█ ",
	},
	'5': {
		"█▀ ",
		"▄█ ",
	},
	'6': {
		"█▄▄ ",
		"█▄█ ",
	},
	'7': {
		"▀▀█ ",
		"░░█ ",
	},
	'8': {
		"█▜█ ",
		"█▙█",
	},
	'9': {
		"█▀█ ",
		"▀▀█ ",
	},
	'0': {
		"█▀█ ",
		"█▄█ ",
	},
	'-': {
		"▄▄ ",
		"░░ ",
	},
	'=': {
		"▀▀ ",
		"▀▀ ",
	},
	',': {
		"░ ",
		"█ ",
	},
	'.': {
		"░ ",
		"▄ ",
	},
	'/': {
		"░░▄▀ ",
		"▄▀░░ ",
	},
	'?': {
		"▀█ ",
		"░▄ ",
	},
	'\\': {
		"▀▄░░ ",
		"░░▀▄ ",
	},
	'_': {
		"░░ ",
		"▄▄ ",
	},
	'!': {
		"█ ",
		"▄ ",
	},
	' ': {
		"░░ ",
		"░░ ",
	},
	':': {
		"▄ ",
		"▄ ",
	},
	'(': {
		"▄▀ ",
		"▀▄ ",
	},
	')': {
		"▀▄ ",
		"▄▀ ",
	},
}
