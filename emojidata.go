package emoji

// Emojis - Map of Emoji Runes as Hex keys to their description
var Emojis = map[string]Emoji{
	"1F600": {
		Key:        "1F600",
		Value:      "😀",
		Descriptor: "grinning face",
	},
	"1F603": {
		Key:        "1F603",
		Value:      "😃",
		Descriptor: "grinning face with big eyes",
	},
	"1F604": {
		Key:        "1F604",
		Value:      "😄",
		Descriptor: "grinning face with smiling eyes",
	},
	"1F601": {
		Key:        "1F601",
		Value:      "😁",
		Descriptor: "beaming face with smiling eyes",
	},
	"1F606": {
		Key:        "1F606",
		Value:      "😆",
		Descriptor: "grinning squinting face",
	},
	"1F605": {
		Key:        "1F605",
		Value:      "😅",
		Descriptor: "grinning face with sweat",
	},
	"1F923": {
		Key:        "1F923",
		Value:      "🤣",
		Descriptor: "rolling on the floor laughing",
	},
	"1F602": {
		Key:        "1F602",
		Value:      "😂",
		Descriptor: "face with tears of joy",
	},
	"1F642": {
		Key:        "1F642",
		Value:      "🙂",
		Descriptor: "slightly smiling face",
	},
	"1F643": {
		Key:        "1F643",
		Value:      "🙃",
		Descriptor: "upside-down face",
	},
	"1FAE0": {
		Key:        "1FAE0",
		Value:      "🫠",
		Descriptor: "melting face",
	},
	"1F609": {
		Key:        "1F609",
		Value:      "😉",
		Descriptor: "winking face",
	},
	"1F60A": {
		Key:        "1F60A",
		Value:      "😊",
		Descriptor: "smiling face with smiling eyes",
	},
	"1F607": {
		Key:        "1F607",
		Value:      "😇",
		Descriptor: "smiling face with halo",
	},
	"1F970": {
		Key:        "1F970",
		Value:      "🥰",
		Descriptor: "smiling face with hearts",
	},
	"1F60D": {
		Key:        "1F60D",
		Value:      "😍",
		Descriptor: "smiling face with heart-eyes",
	},
	"1F929": {
		Key:        "1F929",
		Value:      "🤩",
		Descriptor: "star-struck",
	},
	"1F618": {
		Key:        "1F618",
		Value:      "😘",
		Descriptor: "face blowing a kiss",
	},
	"1F617": {
		Key:        "1F617",
		Value:      "😗",
		Descriptor: "kissing face",
	},
	"263A-FE0F": {
		Key:        "263A-FE0F",
		Value:      "☺️",
		Descriptor: "smiling face",
	},
	"1F61A": {
		Key:        "1F61A",
		Value:      "😚",
		Descriptor: "kissing face with closed eyes",
	},
	"1F619": {
		Key:        "1F619",
		Value:      "😙",
		Descriptor: "kissing face with smiling eyes",
	},
	"1F972": {
		Key:        "1F972",
		Value:      "🥲",
		Descriptor: "smiling face with tear",
	},
	"1F60B": {
		Key:        "1F60B",
		Value:      "😋",
		Descriptor: "face savoring food",
	},
	"1F61B": {
		Key:        "1F61B",
		Value:      "😛",
		Descriptor: "face with tongue",
	},
	"1F61C": {
		Key:        "1F61C",
		Value:      "😜",
		Descriptor: "winking face with tongue",
	},
	"1F92A": {
		Key:        "1F92A",
		Value:      "🤪",
		Descriptor: "zany face",
	},
	"1F61D": {
		Key:        "1F61D",
		Value:      "😝",
		Descriptor: "squinting face with tongue",
	},
	"1F911": {
		Key:        "1F911",
		Value:      "🤑",
		Descriptor: "money-mouth face",
	},
	"1F917": {
		Key:        "1F917",
		Value:      "🤗",
		Descriptor: "smiling face with open hands",
	},
	"1F92D": {
		Key:        "1F92D",
		Value:      "🤭",
		Descriptor: "face with hand over mouth",
	},
	"1FAE2": {
		Key:        "1FAE2",
		Value:      "🫢",
		Descriptor: "face with open eyes and hand over mouth",
	},
	"1FAE3": {
		Key:        "1FAE3",
		Value:      "🫣",
		Descriptor: "face with peeking eye",
	},
	"1F92B": {
		Key:        "1F92B",
		Value:      "🤫",
		Descriptor: "shushing face",
	},
	"1F914": {
		Key:        "1F914",
		Value:      "🤔",
		Descriptor: "thinking face",
	},
	"1FAE1": {
		Key:        "1FAE1",
		Value:      "🫡",
		Descriptor: "saluting face",
	},
	"1F910": {
		Key:        "1F910",
		Value:      "🤐",
		Descriptor: "zipper-mouth face",
	},
	"1F928": {
		Key:        "1F928",
		Value:      "🤨",
		Descriptor: "face with raised eyebrow",
	},
	"1F610": {
		Key:        "1F610",
		Value:      "😐",
		Descriptor: "neutral face",
	},
	"1F611": {
		Key:        "1F611",
		Value:      "😑",
		Descriptor: "expressionless face",
	},
	"1F636": {
		Key:        "1F636",
		Value:      "😶",
		Descriptor: "face without mouth",
	},
	"1FAE5": {
		Key:        "1FAE5",
		Value:      "🫥",
		Descriptor: "dotted line face",
	},
	"1F636-200D-1F32B-FE0F": {
		Key:        "1F636-200D-1F32B-FE0F",
		Value:      "😶‍🌫️",
		Descriptor: "face in clouds",
	},
	"1F60F": {
		Key:        "1F60F",
		Value:      "😏",
		Descriptor: "smirking face",
	},
	"1F612": {
		Key:        "1F612",
		Value:      "😒",
		Descriptor: "unamused face",
	},
	"1F644": {
		Key:        "1F644",
		Value:      "🙄",
		Descriptor: "face with rolling eyes",
	},
	"1F62C": {
		Key:        "1F62C",
		Value:      "😬",
		Descriptor: "grimacing face",
	},
	"1F62E-200D-1F4A8": {
		Key:        "1F62E-200D-1F4A8",
		Value:      "😮‍💨",
		Descriptor: "face exhaling",
	},
	"1F925": {
		Key:        "1F925",
		Value:      "🤥",
		Descriptor: "lying face",
	},
	"1FAE8": {
		Key:        "1FAE8",
		Value:      "🫨",
		Descriptor: "shaking face",
	},
	"1F642-200D-2194-FE0F": {
		Key:        "1F642-200D-2194-FE0F",
		Value:      "🙂‍↔️",
		Descriptor: "head shaking horizontally",
	},
	"1F642-200D-2195-FE0F": {
		Key:        "1F642-200D-2195-FE0F",
		Value:      "🙂‍↕️",
		Descriptor: "head shaking vertically",
	},
	"1F60C": {
		Key:        "1F60C",
		Value:      "😌",
		Descriptor: "relieved face",
	},
	"1F614": {
		Key:        "1F614",
		Value:      "😔",
		Descriptor: "pensive face",
	},
	"1F62A": {
		Key:        "1F62A",
		Value:      "😪",
		Descriptor: "sleepy face",
	},
	"1F924": {
		Key:        "1F924",
		Value:      "🤤",
		Descriptor: "drooling face",
	},
	"1F634": {
		Key:        "1F634",
		Value:      "😴",
		Descriptor: "sleeping face",
	},
	"1F637": {
		Key:        "1F637",
		Value:      "😷",
		Descriptor: "face with medical mask",
	},
	"1F912": {
		Key:        "1F912",
		Value:      "🤒",
		Descriptor: "face with thermometer",
	},
	"1F915": {
		Key:        "1F915",
		Value:      "🤕",
		Descriptor: "face with head-bandage",
	},
	"1F922": {
		Key:        "1F922",
		Value:      "🤢",
		Descriptor: "nauseated face",
	},
	"1F92E": {
		Key:        "1F92E",
		Value:      "🤮",
		Descriptor: "face vomiting",
	},
	"1F927": {
		Key:        "1F927",
		Value:      "🤧",
		Descriptor: "sneezing face",
	},
	"1F975": {
		Key:        "1F975",
		Value:      "🥵",
		Descriptor: "hot face",
	},
	"1F976": {
		Key:        "1F976",
		Value:      "🥶",
		Descriptor: "cold face",
	},
	"1F974": {
		Key:        "1F974",
		Value:      "🥴",
		Descriptor: "woozy face",
	},
	"1F635": {
		Key:        "1F635",
		Value:      "😵",
		Descriptor: "face with crossed-out eyes",
	},
	"1F635-200D-1F4AB": {
		Key:        "1F635-200D-1F4AB",
		Value:      "😵‍💫",
		Descriptor: "face with spiral eyes",
	},
	"1F92F": {
		Key:        "1F92F",
		Value:      "🤯",
		Descriptor: "exploding head",
	},
	"1F920": {
		Key:        "1F920",
		Value:      "🤠",
		Descriptor: "cowboy hat face",
	},
	"1F973": {
		Key:        "1F973",
		Value:      "🥳",
		Descriptor: "partying face",
	},
	"1F978": {
		Key:        "1F978",
		Value:      "🥸",
		Descriptor: "disguised face",
	},
	"1F60E": {
		Key:        "1F60E",
		Value:      "😎",
		Descriptor: "smiling face with sunglasses",
	},
	"1F913": {
		Key:        "1F913",
		Value:      "🤓",
		Descriptor: "nerd face",
	},
	"1F9D0": {
		Key:        "1F9D0",
		Value:      "🧐",
		Descriptor: "face with monocle",
	},
	"1F615": {
		Key:        "1F615",
		Value:      "😕",
		Descriptor: "confused face",
	},
	"1FAE4": {
		Key:        "1FAE4",
		Value:      "🫤",
		Descriptor: "face with diagonal mouth",
	},
	"1F61F": {
		Key:        "1F61F",
		Value:      "😟",
		Descriptor: "worried face",
	},
	"1F641": {
		Key:        "1F641",
		Value:      "🙁",
		Descriptor: "slightly frowning face",
	},
	"2639-FE0F": {
		Key:        "2639-FE0F",
		Value:      "☹️",
		Descriptor: "frowning face",
	},
	"1F62E": {
		Key:        "1F62E",
		Value:      "😮",
		Descriptor: "face with open mouth",
	},
	"1F62F": {
		Key:        "1F62F",
		Value:      "😯",
		Descriptor: "hushed face",
	},
	"1F632": {
		Key:        "1F632",
		Value:      "😲",
		Descriptor: "astonished face",
	},
	"1F633": {
		Key:        "1F633",
		Value:      "😳",
		Descriptor: "flushed face",
	},
	"1F97A": {
		Key:        "1F97A",
		Value:      "🥺",
		Descriptor: "pleading face",
	},
	"1F979": {
		Key:        "1F979",
		Value:      "🥹",
		Descriptor: "face holding back tears",
	},
	"1F626": {
		Key:        "1F626",
		Value:      "😦",
		Descriptor: "frowning face with open mouth",
	},
	"1F627": {
		Key:        "1F627",
		Value:      "😧",
		Descriptor: "anguished face",
	},
	"1F628": {
		Key:        "1F628",
		Value:      "😨",
		Descriptor: "fearful face",
	},
	"1F630": {
		Key:        "1F630",
		Value:      "😰",
		Descriptor: "anxious face with sweat",
	},
	"1F625": {
		Key:        "1F625",
		Value:      "😥",
		Descriptor: "sad but relieved face",
	},
	"1F622": {
		Key:        "1F622",
		Value:      "😢",
		Descriptor: "crying face",
	},
	"1F62D": {
		Key:        "1F62D",
		Value:      "😭",
		Descriptor: "loudly crying face",
	},
	"1F631": {
		Key:        "1F631",
		Value:      "😱",
		Descriptor: "face screaming in fear",
	},
	"1F616": {
		Key:        "1F616",
		Value:      "😖",
		Descriptor: "confounded face",
	},
	"1F623": {
		Key:        "1F623",
		Value:      "😣",
		Descriptor: "persevering face",
	},
	"1F61E": {
		Key:        "1F61E",
		Value:      "😞",
		Descriptor: "disappointed face",
	},
	"1F613": {
		Key:        "1F613",
		Value:      "😓",
		Descriptor: "downcast face with sweat",
	},
	"1F629": {
		Key:        "1F629",
		Value:      "😩",
		Descriptor: "weary face",
	},
	"1F62B": {
		Key:        "1F62B",
		Value:      "😫",
		Descriptor: "tired face",
	},
	"1F971": {
		Key:        "1F971",
		Value:      "🥱",
		Descriptor: "yawning face",
	},
	"1F624": {
		Key:        "1F624",
		Value:      "😤",
		Descriptor: "face with steam from nose",
	},
	"1F621": {
		Key:        "1F621",
		Value:      "😡",
		Descriptor: "enraged face",
	},
	"1F620": {
		Key:        "1F620",
		Value:      "😠",
		Descriptor: "angry face",
	},
	"1F92C": {
		Key:        "1F92C",
		Value:      "🤬",
		Descriptor: "face with symbols on mouth",
	},
	"1F608": {
		Key:        "1F608",
		Value:      "😈",
		Descriptor: "smiling face with horns",
	},
	"1F47F": {
		Key:        "1F47F",
		Value:      "👿",
		Descriptor: "angry face with horns",
	},
	"1F480": {
		Key:        "1F480",
		Value:      "💀",
		Descriptor: "skull",
	},
	"2620-FE0F": {
		Key:        "2620-FE0F",
		Value:      "☠️",
		Descriptor: "skull and crossbones",
	},
	"1F4A9": {
		Key:        "1F4A9",
		Value:      "💩",
		Descriptor: "pile of poo",
	},
	"1F921": {
		Key:        "1F921",
		Value:      "🤡",
		Descriptor: "clown face",
	},
	"1F479": {
		Key:        "1F479",
		Value:      "👹",
		Descriptor: "ogre",
	},
	"1F47A": {
		Key:        "1F47A",
		Value:      "👺",
		Descriptor: "goblin",
	},
	"1F47B": {
		Key:        "1F47B",
		Value:      "👻",
		Descriptor: "ghost",
	},
	"1F47D": {
		Key:        "1F47D",
		Value:      "👽",
		Descriptor: "alien",
	},
	"1F47E": {
		Key:        "1F47E",
		Value:      "👾",
		Descriptor: "alien monster",
	},
	"1F916": {
		Key:        "1F916",
		Value:      "🤖",
		Descriptor: "robot",
	},
	"1F63A": {
		Key:        "1F63A",
		Value:      "😺",
		Descriptor: "grinning cat",
	},
	"1F638": {
		Key:        "1F638",
		Value:      "😸",
		Descriptor: "grinning cat with smiling eyes",
	},
	"1F639": {
		Key:        "1F639",
		Value:      "😹",
		Descriptor: "cat with tears of joy",
	},
	"1F63B": {
		Key:        "1F63B",
		Value:      "😻",
		Descriptor: "smiling cat with heart-eyes",
	},
	"1F63C": {
		Key:        "1F63C",
		Value:      "😼",
		Descriptor: "cat with wry smile",
	},
	"1F63D": {
		Key:        "1F63D",
		Value:      "😽",
		Descriptor: "kissing cat",
	},
	"1F640": {
		Key:        "1F640",
		Value:      "🙀",
		Descriptor: "weary cat",
	},
	"1F63F": {
		Key:        "1F63F",
		Value:      "😿",
		Descriptor: "crying cat",
	},
	"1F63E": {
		Key:        "1F63E",
		Value:      "😾",
		Descriptor: "pouting cat",
	},
	"1F648": {
		Key:        "1F648",
		Value:      "🙈",
		Descriptor: "see-no-evil monkey",
	},
	"1F649": {
		Key:        "1F649",
		Value:      "🙉",
		Descriptor: "hear-no-evil monkey",
	},
	"1F64A": {
		Key:        "1F64A",
		Value:      "🙊",
		Descriptor: "speak-no-evil monkey",
	},
	"1F48C": {
		Key:        "1F48C",
		Value:      "💌",
		Descriptor: "love letter",
	},
	"1F498": {
		Key:        "1F498",
		Value:      "💘",
		Descriptor: "heart with arrow",
	},
	"1F49D": {
		Key:        "1F49D",
		Value:      "💝",
		Descriptor: "heart with ribbon",
	},
	"1F496": {
		Key:        "1F496",
		Value:      "💖",
		Descriptor: "sparkling heart",
	},
	"1F497": {
		Key:        "1F497",
		Value:      "💗",
		Descriptor: "growing heart",
	},
	"1F493": {
		Key:        "1F493",
		Value:      "💓",
		Descriptor: "beating heart",
	},
	"1F49E": {
		Key:        "1F49E",
		Value:      "💞",
		Descriptor: "revolving hearts",
	},
	"1F495": {
		Key:        "1F495",
		Value:      "💕",
		Descriptor: "two hearts",
	},
	"1F49F": {
		Key:        "1F49F",
		Value:      "💟",
		Descriptor: "heart decoration",
	},
	"2763-FE0F": {
		Key:        "2763-FE0F",
		Value:      "❣️",
		Descriptor: "heart exclamation",
	},
	"1F494": {
		Key:        "1F494",
		Value:      "💔",
		Descriptor: "broken heart",
	},
	"2764-FE0F-200D-1F525": {
		Key:        "2764-FE0F-200D-1F525",
		Value:      "❤️‍🔥",
		Descriptor: "heart on fire",
	},
	"2764-FE0F-200D-1FA79": {
		Key:        "2764-FE0F-200D-1FA79",
		Value:      "❤️‍🩹",
		Descriptor: "mending heart",
	},
	"2764-FE0F": {
		Key:        "2764-FE0F",
		Value:      "❤️",
		Descriptor: "red heart",
	},
	"1FA77": {
		Key:        "1FA77",
		Value:      "🩷",
		Descriptor: "pink heart",
	},
	"1F9E1": {
		Key:        "1F9E1",
		Value:      "🧡",
		Descriptor: "orange heart",
	},
	"1F49B": {
		Key:        "1F49B",
		Value:      "💛",
		Descriptor: "yellow heart",
	},
	"1F49A": {
		Key:        "1F49A",
		Value:      "💚",
		Descriptor: "green heart",
	},
	"1F499": {
		Key:        "1F499",
		Value:      "💙",
		Descriptor: "blue heart",
	},
	"1FA75": {
		Key:        "1FA75",
		Value:      "🩵",
		Descriptor: "light blue heart",
	},
	"1F49C": {
		Key:        "1F49C",
		Value:      "💜",
		Descriptor: "purple heart",
	},
	"1F90E": {
		Key:        "1F90E",
		Value:      "🤎",
		Descriptor: "brown heart",
	},
	"1F5A4": {
		Key:        "1F5A4",
		Value:      "🖤",
		Descriptor: "black heart",
	},
	"1FA76": {
		Key:        "1FA76",
		Value:      "🩶",
		Descriptor: "grey heart",
	},
	"1F90D": {
		Key:        "1F90D",
		Value:      "🤍",
		Descriptor: "white heart",
	},
	"1F48B": {
		Key:        "1F48B",
		Value:      "💋",
		Descriptor: "kiss mark",
	},
	"1F4AF": {
		Key:        "1F4AF",
		Value:      "💯",
		Descriptor: "hundred points",
	},
	"1F4A2": {
		Key:        "1F4A2",
		Value:      "💢",
		Descriptor: "anger symbol",
	},
	"1F4A5": {
		Key:        "1F4A5",
		Value:      "💥",
		Descriptor: "collision",
	},
	"1F4AB": {
		Key:        "1F4AB",
		Value:      "💫",
		Descriptor: "dizzy",
	},
	"1F4A6": {
		Key:        "1F4A6",
		Value:      "💦",
		Descriptor: "sweat droplets",
	},
	"1F4A8": {
		Key:        "1F4A8",
		Value:      "💨",
		Descriptor: "dashing away",
	},
	"1F573-FE0F": {
		Key:        "1F573-FE0F",
		Value:      "🕳️",
		Descriptor: "hole",
	},
	"1F4AC": {
		Key:        "1F4AC",
		Value:      "💬",
		Descriptor: "speech balloon",
	},
	"1F441-FE0F-200D-1F5E8-FE0F": {
		Key:        "1F441-FE0F-200D-1F5E8-FE0F",
		Value:      "👁️‍🗨️",
		Descriptor: "eye in speech bubble",
	},
	"1F5E8-FE0F": {
		Key:        "1F5E8-FE0F",
		Value:      "🗨️",
		Descriptor: "left speech bubble",
	},
	"1F5EF-FE0F": {
		Key:        "1F5EF-FE0F",
		Value:      "🗯️",
		Descriptor: "right anger bubble",
	},
	"1F4AD": {
		Key:        "1F4AD",
		Value:      "💭",
		Descriptor: "thought balloon",
	},
	"1F4A4": {
		Key:        "1F4A4",
		Value:      "💤",
		Descriptor: "ZZZ",
	},
	"1F44B": {
		Key:        "1F44B",
		Value:      "👋",
		Descriptor: "waving hand",
	},
	"1F44B-1F3FB": {
		Key:        "1F44B-1F3FB",
		Value:      "👋🏻",
		Descriptor: "waving hand: light skin tone",
	},
	"1F44B-1F3FC": {
		Key:        "1F44B-1F3FC",
		Value:      "👋🏼",
		Descriptor: "waving hand: medium-light skin tone",
	},
	"1F44B-1F3FD": {
		Key:        "1F44B-1F3FD",
		Value:      "👋🏽",
		Descriptor: "waving hand: medium skin tone",
	},
	"1F44B-1F3FE": {
		Key:        "1F44B-1F3FE",
		Value:      "👋🏾",
		Descriptor: "waving hand: medium-dark skin tone",
	},
	"1F44B-1F3FF": {
		Key:        "1F44B-1F3FF",
		Value:      "👋🏿",
		Descriptor: "waving hand: dark skin tone",
	},
	"1F91A": {
		Key:        "1F91A",
		Value:      "🤚",
		Descriptor: "raised back of hand",
	},
	"1F91A-1F3FB": {
		Key:        "1F91A-1F3FB",
		Value:      "🤚🏻",
		Descriptor: "raised back of hand: light skin tone",
	},
	"1F91A-1F3FC": {
		Key:        "1F91A-1F3FC",
		Value:      "🤚🏼",
		Descriptor: "raised back of hand: medium-light skin tone",
	},
	"1F91A-1F3FD": {
		Key:        "1F91A-1F3FD",
		Value:      "🤚🏽",
		Descriptor: "raised back of hand: medium skin tone",
	},
	"1F91A-1F3FE": {
		Key:        "1F91A-1F3FE",
		Value:      "🤚🏾",
		Descriptor: "raised back of hand: medium-dark skin tone",
	},
	"1F91A-1F3FF": {
		Key:        "1F91A-1F3FF",
		Value:      "🤚🏿",
		Descriptor: "raised back of hand: dark skin tone",
	},
	"1F590-FE0F": {
		Key:        "1F590-FE0F",
		Value:      "🖐️",
		Descriptor: "hand with fingers splayed",
	},
	"1F590-1F3FB": {
		Key:        "1F590-1F3FB",
		Value:      "🖐🏻",
		Descriptor: "hand with fingers splayed: light skin tone",
	},
	"1F590-1F3FC": {
		Key:        "1F590-1F3FC",
		Value:      "🖐🏼",
		Descriptor: "hand with fingers splayed: medium-light skin tone",
	},
	"1F590-1F3FD": {
		Key:        "1F590-1F3FD",
		Value:      "🖐🏽",
		Descriptor: "hand with fingers splayed: medium skin tone",
	},
	"1F590-1F3FE": {
		Key:        "1F590-1F3FE",
		Value:      "🖐🏾",
		Descriptor: "hand with fingers splayed: medium-dark skin tone",
	},
	"1F590-1F3FF": {
		Key:        "1F590-1F3FF",
		Value:      "🖐🏿",
		Descriptor: "hand with fingers splayed: dark skin tone",
	},
	"270B": {
		Key:        "270B",
		Value:      "✋",
		Descriptor: "raised hand",
	},
	"270B-1F3FB": {
		Key:        "270B-1F3FB",
		Value:      "✋🏻",
		Descriptor: "raised hand: light skin tone",
	},
	"270B-1F3FC": {
		Key:        "270B-1F3FC",
		Value:      "✋🏼",
		Descriptor: "raised hand: medium-light skin tone",
	},
	"270B-1F3FD": {
		Key:        "270B-1F3FD",
		Value:      "✋🏽",
		Descriptor: "raised hand: medium skin tone",
	},
	"270B-1F3FE": {
		Key:        "270B-1F3FE",
		Value:      "✋🏾",
		Descriptor: "raised hand: medium-dark skin tone",
	},
	"270B-1F3FF": {
		Key:        "270B-1F3FF",
		Value:      "✋🏿",
		Descriptor: "raised hand: dark skin tone",
	},
	"1F596": {
		Key:        "1F596",
		Value:      "🖖",
		Descriptor: "vulcan salute",
	},
	"1F596-1F3FB": {
		Key:        "1F596-1F3FB",
		Value:      "🖖🏻",
		Descriptor: "vulcan salute: light skin tone",
	},
	"1F596-1F3FC": {
		Key:        "1F596-1F3FC",
		Value:      "🖖🏼",
		Descriptor: "vulcan salute: medium-light skin tone",
	},
	"1F596-1F3FD": {
		Key:        "1F596-1F3FD",
		Value:      "🖖🏽",
		Descriptor: "vulcan salute: medium skin tone",
	},
	"1F596-1F3FE": {
		Key:        "1F596-1F3FE",
		Value:      "🖖🏾",
		Descriptor: "vulcan salute: medium-dark skin tone",
	},
	"1F596-1F3FF": {
		Key:        "1F596-1F3FF",
		Value:      "🖖🏿",
		Descriptor: "vulcan salute: dark skin tone",
	},
	"1FAF1": {
		Key:        "1FAF1",
		Value:      "🫱",
		Descriptor: "rightwards hand",
	},
	"1FAF1-1F3FB": {
		Key:        "1FAF1-1F3FB",
		Value:      "🫱🏻",
		Descriptor: "rightwards hand: light skin tone",
	},
	"1FAF1-1F3FC": {
		Key:        "1FAF1-1F3FC",
		Value:      "🫱🏼",
		Descriptor: "rightwards hand: medium-light skin tone",
	},
	"1FAF1-1F3FD": {
		Key:        "1FAF1-1F3FD",
		Value:      "🫱🏽",
		Descriptor: "rightwards hand: medium skin tone",
	},
	"1FAF1-1F3FE": {
		Key:        "1FAF1-1F3FE",
		Value:      "🫱🏾",
		Descriptor: "rightwards hand: medium-dark skin tone",
	},
	"1FAF1-1F3FF": {
		Key:        "1FAF1-1F3FF",
		Value:      "🫱🏿",
		Descriptor: "rightwards hand: dark skin tone",
	},
	"1FAF2": {
		Key:        "1FAF2",
		Value:      "🫲",
		Descriptor: "leftwards hand",
	},
	"1FAF2-1F3FB": {
		Key:        "1FAF2-1F3FB",
		Value:      "🫲🏻",
		Descriptor: "leftwards hand: light skin tone",
	},
	"1FAF2-1F3FC": {
		Key:        "1FAF2-1F3FC",
		Value:      "🫲🏼",
		Descriptor: "leftwards hand: medium-light skin tone",
	},
	"1FAF2-1F3FD": {
		Key:        "1FAF2-1F3FD",
		Value:      "🫲🏽",
		Descriptor: "leftwards hand: medium skin tone",
	},
	"1FAF2-1F3FE": {
		Key:        "1FAF2-1F3FE",
		Value:      "🫲🏾",
		Descriptor: "leftwards hand: medium-dark skin tone",
	},
	"1FAF2-1F3FF": {
		Key:        "1FAF2-1F3FF",
		Value:      "🫲🏿",
		Descriptor: "leftwards hand: dark skin tone",
	},
	"1FAF3": {
		Key:        "1FAF3",
		Value:      "🫳",
		Descriptor: "palm down hand",
	},
	"1FAF3-1F3FB": {
		Key:        "1FAF3-1F3FB",
		Value:      "🫳🏻",
		Descriptor: "palm down hand: light skin tone",
	},
	"1FAF3-1F3FC": {
		Key:        "1FAF3-1F3FC",
		Value:      "🫳🏼",
		Descriptor: "palm down hand: medium-light skin tone",
	},
	"1FAF3-1F3FD": {
		Key:        "1FAF3-1F3FD",
		Value:      "🫳🏽",
		Descriptor: "palm down hand: medium skin tone",
	},
	"1FAF3-1F3FE": {
		Key:        "1FAF3-1F3FE",
		Value:      "🫳🏾",
		Descriptor: "palm down hand: medium-dark skin tone",
	},
	"1FAF3-1F3FF": {
		Key:        "1FAF3-1F3FF",
		Value:      "🫳🏿",
		Descriptor: "palm down hand: dark skin tone",
	},
	"1FAF4": {
		Key:        "1FAF4",
		Value:      "🫴",
		Descriptor: "palm up hand",
	},
	"1FAF4-1F3FB": {
		Key:        "1FAF4-1F3FB",
		Value:      "🫴🏻",
		Descriptor: "palm up hand: light skin tone",
	},
	"1FAF4-1F3FC": {
		Key:        "1FAF4-1F3FC",
		Value:      "🫴🏼",
		Descriptor: "palm up hand: medium-light skin tone",
	},
	"1FAF4-1F3FD": {
		Key:        "1FAF4-1F3FD",
		Value:      "🫴🏽",
		Descriptor: "palm up hand: medium skin tone",
	},
	"1FAF4-1F3FE": {
		Key:        "1FAF4-1F3FE",
		Value:      "🫴🏾",
		Descriptor: "palm up hand: medium-dark skin tone",
	},
	"1FAF4-1F3FF": {
		Key:        "1FAF4-1F3FF",
		Value:      "🫴🏿",
		Descriptor: "palm up hand: dark skin tone",
	},
	"1FAF7": {
		Key:        "1FAF7",
		Value:      "🫷",
		Descriptor: "leftwards pushing hand",
	},
	"1FAF7-1F3FB": {
		Key:        "1FAF7-1F3FB",
		Value:      "🫷🏻",
		Descriptor: "leftwards pushing hand: light skin tone",
	},
	"1FAF7-1F3FC": {
		Key:        "1FAF7-1F3FC",
		Value:      "🫷🏼",
		Descriptor: "leftwards pushing hand: medium-light skin tone",
	},
	"1FAF7-1F3FD": {
		Key:        "1FAF7-1F3FD",
		Value:      "🫷🏽",
		Descriptor: "leftwards pushing hand: medium skin tone",
	},
	"1FAF7-1F3FE": {
		Key:        "1FAF7-1F3FE",
		Value:      "🫷🏾",
		Descriptor: "leftwards pushing hand: medium-dark skin tone",
	},
	"1FAF7-1F3FF": {
		Key:        "1FAF7-1F3FF",
		Value:      "🫷🏿",
		Descriptor: "leftwards pushing hand: dark skin tone",
	},
	"1FAF8": {
		Key:        "1FAF8",
		Value:      "🫸",
		Descriptor: "rightwards pushing hand",
	},
	"1FAF8-1F3FB": {
		Key:        "1FAF8-1F3FB",
		Value:      "🫸🏻",
		Descriptor: "rightwards pushing hand: light skin tone",
	},
	"1FAF8-1F3FC": {
		Key:        "1FAF8-1F3FC",
		Value:      "🫸🏼",
		Descriptor: "rightwards pushing hand: medium-light skin tone",
	},
	"1FAF8-1F3FD": {
		Key:        "1FAF8-1F3FD",
		Value:      "🫸🏽",
		Descriptor: "rightwards pushing hand: medium skin tone",
	},
	"1FAF8-1F3FE": {
		Key:        "1FAF8-1F3FE",
		Value:      "🫸🏾",
		Descriptor: "rightwards pushing hand: medium-dark skin tone",
	},
	"1FAF8-1F3FF": {
		Key:        "1FAF8-1F3FF",
		Value:      "🫸🏿",
		Descriptor: "rightwards pushing hand: dark skin tone",
	},
	"1F44C": {
		Key:        "1F44C",
		Value:      "👌",
		Descriptor: "OK hand",
	},
	"1F44C-1F3FB": {
		Key:        "1F44C-1F3FB",
		Value:      "👌🏻",
		Descriptor: "OK hand: light skin tone",
	},
	"1F44C-1F3FC": {
		Key:        "1F44C-1F3FC",
		Value:      "👌🏼",
		Descriptor: "OK hand: medium-light skin tone",
	},
	"1F44C-1F3FD": {
		Key:        "1F44C-1F3FD",
		Value:      "👌🏽",
		Descriptor: "OK hand: medium skin tone",
	},
	"1F44C-1F3FE": {
		Key:        "1F44C-1F3FE",
		Value:      "👌🏾",
		Descriptor: "OK hand: medium-dark skin tone",
	},
	"1F44C-1F3FF": {
		Key:        "1F44C-1F3FF",
		Value:      "👌🏿",
		Descriptor: "OK hand: dark skin tone",
	},
	"1F90C": {
		Key:        "1F90C",
		Value:      "🤌",
		Descriptor: "pinched fingers",
	},
	"1F90C-1F3FB": {
		Key:        "1F90C-1F3FB",
		Value:      "🤌🏻",
		Descriptor: "pinched fingers: light skin tone",
	},
	"1F90C-1F3FC": {
		Key:        "1F90C-1F3FC",
		Value:      "🤌🏼",
		Descriptor: "pinched fingers: medium-light skin tone",
	},
	"1F90C-1F3FD": {
		Key:        "1F90C-1F3FD",
		Value:      "🤌🏽",
		Descriptor: "pinched fingers: medium skin tone",
	},
	"1F90C-1F3FE": {
		Key:        "1F90C-1F3FE",
		Value:      "🤌🏾",
		Descriptor: "pinched fingers: medium-dark skin tone",
	},
	"1F90C-1F3FF": {
		Key:        "1F90C-1F3FF",
		Value:      "🤌🏿",
		Descriptor: "pinched fingers: dark skin tone",
	},
	"1F90F": {
		Key:        "1F90F",
		Value:      "🤏",
		Descriptor: "pinching hand",
	},
	"1F90F-1F3FB": {
		Key:        "1F90F-1F3FB",
		Value:      "🤏🏻",
		Descriptor: "pinching hand: light skin tone",
	},
	"1F90F-1F3FC": {
		Key:        "1F90F-1F3FC",
		Value:      "🤏🏼",
		Descriptor: "pinching hand: medium-light skin tone",
	},
	"1F90F-1F3FD": {
		Key:        "1F90F-1F3FD",
		Value:      "🤏🏽",
		Descriptor: "pinching hand: medium skin tone",
	},
	"1F90F-1F3FE": {
		Key:        "1F90F-1F3FE",
		Value:      "🤏🏾",
		Descriptor: "pinching hand: medium-dark skin tone",
	},
	"1F90F-1F3FF": {
		Key:        "1F90F-1F3FF",
		Value:      "🤏🏿",
		Descriptor: "pinching hand: dark skin tone",
	},
	"270C-FE0F": {
		Key:        "270C-FE0F",
		Value:      "✌️",
		Descriptor: "victory hand",
	},
	"270C-1F3FB": {
		Key:        "270C-1F3FB",
		Value:      "✌🏻",
		Descriptor: "victory hand: light skin tone",
	},
	"270C-1F3FC": {
		Key:        "270C-1F3FC",
		Value:      "✌🏼",
		Descriptor: "victory hand: medium-light skin tone",
	},
	"270C-1F3FD": {
		Key:        "270C-1F3FD",
		Value:      "✌🏽",
		Descriptor: "victory hand: medium skin tone",
	},
	"270C-1F3FE": {
		Key:        "270C-1F3FE",
		Value:      "✌🏾",
		Descriptor: "victory hand: medium-dark skin tone",
	},
	"270C-1F3FF": {
		Key:        "270C-1F3FF",
		Value:      "✌🏿",
		Descriptor: "victory hand: dark skin tone",
	},
	"1F91E": {
		Key:        "1F91E",
		Value:      "🤞",
		Descriptor: "crossed fingers",
	},
	"1F91E-1F3FB": {
		Key:        "1F91E-1F3FB",
		Value:      "🤞🏻",
		Descriptor: "crossed fingers: light skin tone",
	},
	"1F91E-1F3FC": {
		Key:        "1F91E-1F3FC",
		Value:      "🤞🏼",
		Descriptor: "crossed fingers: medium-light skin tone",
	},
	"1F91E-1F3FD": {
		Key:        "1F91E-1F3FD",
		Value:      "🤞🏽",
		Descriptor: "crossed fingers: medium skin tone",
	},
	"1F91E-1F3FE": {
		Key:        "1F91E-1F3FE",
		Value:      "🤞🏾",
		Descriptor: "crossed fingers: medium-dark skin tone",
	},
	"1F91E-1F3FF": {
		Key:        "1F91E-1F3FF",
		Value:      "🤞🏿",
		Descriptor: "crossed fingers: dark skin tone",
	},
	"1FAF0": {
		Key:        "1FAF0",
		Value:      "🫰",
		Descriptor: "hand with index finger and thumb crossed",
	},
	"1FAF0-1F3FB": {
		Key:        "1FAF0-1F3FB",
		Value:      "🫰🏻",
		Descriptor: "hand with index finger and thumb crossed: light skin tone",
	},
	"1FAF0-1F3FC": {
		Key:        "1FAF0-1F3FC",
		Value:      "🫰🏼",
		Descriptor: "hand with index finger and thumb crossed: medium-light skin tone",
	},
	"1FAF0-1F3FD": {
		Key:        "1FAF0-1F3FD",
		Value:      "🫰🏽",
		Descriptor: "hand with index finger and thumb crossed: medium skin tone",
	},
	"1FAF0-1F3FE": {
		Key:        "1FAF0-1F3FE",
		Value:      "🫰🏾",
		Descriptor: "hand with index finger and thumb crossed: medium-dark skin tone",
	},
	"1FAF0-1F3FF": {
		Key:        "1FAF0-1F3FF",
		Value:      "🫰🏿",
		Descriptor: "hand with index finger and thumb crossed: dark skin tone",
	},
	"1F91F": {
		Key:        "1F91F",
		Value:      "🤟",
		Descriptor: "love-you gesture",
	},
	"1F91F-1F3FB": {
		Key:        "1F91F-1F3FB",
		Value:      "🤟🏻",
		Descriptor: "love-you gesture: light skin tone",
	},
	"1F91F-1F3FC": {
		Key:        "1F91F-1F3FC",
		Value:      "🤟🏼",
		Descriptor: "love-you gesture: medium-light skin tone",
	},
	"1F91F-1F3FD": {
		Key:        "1F91F-1F3FD",
		Value:      "🤟🏽",
		Descriptor: "love-you gesture: medium skin tone",
	},
	"1F91F-1F3FE": {
		Key:        "1F91F-1F3FE",
		Value:      "🤟🏾",
		Descriptor: "love-you gesture: medium-dark skin tone",
	},
	"1F91F-1F3FF": {
		Key:        "1F91F-1F3FF",
		Value:      "🤟🏿",
		Descriptor: "love-you gesture: dark skin tone",
	},
	"1F918": {
		Key:        "1F918",
		Value:      "🤘",
		Descriptor: "sign of the horns",
	},
	"1F918-1F3FB": {
		Key:        "1F918-1F3FB",
		Value:      "🤘🏻",
		Descriptor: "sign of the horns: light skin tone",
	},
	"1F918-1F3FC": {
		Key:        "1F918-1F3FC",
		Value:      "🤘🏼",
		Descriptor: "sign of the horns: medium-light skin tone",
	},
	"1F918-1F3FD": {
		Key:        "1F918-1F3FD",
		Value:      "🤘🏽",
		Descriptor: "sign of the horns: medium skin tone",
	},
	"1F918-1F3FE": {
		Key:        "1F918-1F3FE",
		Value:      "🤘🏾",
		Descriptor: "sign of the horns: medium-dark skin tone",
	},
	"1F918-1F3FF": {
		Key:        "1F918-1F3FF",
		Value:      "🤘🏿",
		Descriptor: "sign of the horns: dark skin tone",
	},
	"1F919": {
		Key:        "1F919",
		Value:      "🤙",
		Descriptor: "call me hand",
	},
	"1F919-1F3FB": {
		Key:        "1F919-1F3FB",
		Value:      "🤙🏻",
		Descriptor: "call me hand: light skin tone",
	},
	"1F919-1F3FC": {
		Key:        "1F919-1F3FC",
		Value:      "🤙🏼",
		Descriptor: "call me hand: medium-light skin tone",
	},
	"1F919-1F3FD": {
		Key:        "1F919-1F3FD",
		Value:      "🤙🏽",
		Descriptor: "call me hand: medium skin tone",
	},
	"1F919-1F3FE": {
		Key:        "1F919-1F3FE",
		Value:      "🤙🏾",
		Descriptor: "call me hand: medium-dark skin tone",
	},
	"1F919-1F3FF": {
		Key:        "1F919-1F3FF",
		Value:      "🤙🏿",
		Descriptor: "call me hand: dark skin tone",
	},
	"1F448": {
		Key:        "1F448",
		Value:      "👈",
		Descriptor: "backhand index pointing left",
	},
	"1F448-1F3FB": {
		Key:        "1F448-1F3FB",
		Value:      "👈🏻",
		Descriptor: "backhand index pointing left: light skin tone",
	},
	"1F448-1F3FC": {
		Key:        "1F448-1F3FC",
		Value:      "👈🏼",
		Descriptor: "backhand index pointing left: medium-light skin tone",
	},
	"1F448-1F3FD": {
		Key:        "1F448-1F3FD",
		Value:      "👈🏽",
		Descriptor: "backhand index pointing left: medium skin tone",
	},
	"1F448-1F3FE": {
		Key:        "1F448-1F3FE",
		Value:      "👈🏾",
		Descriptor: "backhand index pointing left: medium-dark skin tone",
	},
	"1F448-1F3FF": {
		Key:        "1F448-1F3FF",
		Value:      "👈🏿",
		Descriptor: "backhand index pointing left: dark skin tone",
	},
	"1F449": {
		Key:        "1F449",
		Value:      "👉",
		Descriptor: "backhand index pointing right",
	},
	"1F449-1F3FB": {
		Key:        "1F449-1F3FB",
		Value:      "👉🏻",
		Descriptor: "backhand index pointing right: light skin tone",
	},
	"1F449-1F3FC": {
		Key:        "1F449-1F3FC",
		Value:      "👉🏼",
		Descriptor: "backhand index pointing right: medium-light skin tone",
	},
	"1F449-1F3FD": {
		Key:        "1F449-1F3FD",
		Value:      "👉🏽",
		Descriptor: "backhand index pointing right: medium skin tone",
	},
	"1F449-1F3FE": {
		Key:        "1F449-1F3FE",
		Value:      "👉🏾",
		Descriptor: "backhand index pointing right: medium-dark skin tone",
	},
	"1F449-1F3FF": {
		Key:        "1F449-1F3FF",
		Value:      "👉🏿",
		Descriptor: "backhand index pointing right: dark skin tone",
	},
	"1F446": {
		Key:        "1F446",
		Value:      "👆",
		Descriptor: "backhand index pointing up",
	},
	"1F446-1F3FB": {
		Key:        "1F446-1F3FB",
		Value:      "👆🏻",
		Descriptor: "backhand index pointing up: light skin tone",
	},
	"1F446-1F3FC": {
		Key:        "1F446-1F3FC",
		Value:      "👆🏼",
		Descriptor: "backhand index pointing up: medium-light skin tone",
	},
	"1F446-1F3FD": {
		Key:        "1F446-1F3FD",
		Value:      "👆🏽",
		Descriptor: "backhand index pointing up: medium skin tone",
	},
	"1F446-1F3FE": {
		Key:        "1F446-1F3FE",
		Value:      "👆🏾",
		Descriptor: "backhand index pointing up: medium-dark skin tone",
	},
	"1F446-1F3FF": {
		Key:        "1F446-1F3FF",
		Value:      "👆🏿",
		Descriptor: "backhand index pointing up: dark skin tone",
	},
	"1F595": {
		Key:        "1F595",
		Value:      "🖕",
		Descriptor: "middle finger",
	},
	"1F595-1F3FB": {
		Key:        "1F595-1F3FB",
		Value:      "🖕🏻",
		Descriptor: "middle finger: light skin tone",
	},
	"1F595-1F3FC": {
		Key:        "1F595-1F3FC",
		Value:      "🖕🏼",
		Descriptor: "middle finger: medium-light skin tone",
	},
	"1F595-1F3FD": {
		Key:        "1F595-1F3FD",
		Value:      "🖕🏽",
		Descriptor: "middle finger: medium skin tone",
	},
	"1F595-1F3FE": {
		Key:        "1F595-1F3FE",
		Value:      "🖕🏾",
		Descriptor: "middle finger: medium-dark skin tone",
	},
	"1F595-1F3FF": {
		Key:        "1F595-1F3FF",
		Value:      "🖕🏿",
		Descriptor: "middle finger: dark skin tone",
	},
	"1F447": {
		Key:        "1F447",
		Value:      "👇",
		Descriptor: "backhand index pointing down",
	},
	"1F447-1F3FB": {
		Key:        "1F447-1F3FB",
		Value:      "👇🏻",
		Descriptor: "backhand index pointing down: light skin tone",
	},
	"1F447-1F3FC": {
		Key:        "1F447-1F3FC",
		Value:      "👇🏼",
		Descriptor: "backhand index pointing down: medium-light skin tone",
	},
	"1F447-1F3FD": {
		Key:        "1F447-1F3FD",
		Value:      "👇🏽",
		Descriptor: "backhand index pointing down: medium skin tone",
	},
	"1F447-1F3FE": {
		Key:        "1F447-1F3FE",
		Value:      "👇🏾",
		Descriptor: "backhand index pointing down: medium-dark skin tone",
	},
	"1F447-1F3FF": {
		Key:        "1F447-1F3FF",
		Value:      "👇🏿",
		Descriptor: "backhand index pointing down: dark skin tone",
	},
	"261D-FE0F": {
		Key:        "261D-FE0F",
		Value:      "☝️",
		Descriptor: "index pointing up",
	},
	"261D-1F3FB": {
		Key:        "261D-1F3FB",
		Value:      "☝🏻",
		Descriptor: "index pointing up: light skin tone",
	},
	"261D-1F3FC": {
		Key:        "261D-1F3FC",
		Value:      "☝🏼",
		Descriptor: "index pointing up: medium-light skin tone",
	},
	"261D-1F3FD": {
		Key:        "261D-1F3FD",
		Value:      "☝🏽",
		Descriptor: "index pointing up: medium skin tone",
	},
	"261D-1F3FE": {
		Key:        "261D-1F3FE",
		Value:      "☝🏾",
		Descriptor: "index pointing up: medium-dark skin tone",
	},
	"261D-1F3FF": {
		Key:        "261D-1F3FF",
		Value:      "☝🏿",
		Descriptor: "index pointing up: dark skin tone",
	},
	"1FAF5": {
		Key:        "1FAF5",
		Value:      "🫵",
		Descriptor: "index pointing at the viewer",
	},
	"1FAF5-1F3FB": {
		Key:        "1FAF5-1F3FB",
		Value:      "🫵🏻",
		Descriptor: "index pointing at the viewer: light skin tone",
	},
	"1FAF5-1F3FC": {
		Key:        "1FAF5-1F3FC",
		Value:      "🫵🏼",
		Descriptor: "index pointing at the viewer: medium-light skin tone",
	},
	"1FAF5-1F3FD": {
		Key:        "1FAF5-1F3FD",
		Value:      "🫵🏽",
		Descriptor: "index pointing at the viewer: medium skin tone",
	},
	"1FAF5-1F3FE": {
		Key:        "1FAF5-1F3FE",
		Value:      "🫵🏾",
		Descriptor: "index pointing at the viewer: medium-dark skin tone",
	},
	"1FAF5-1F3FF": {
		Key:        "1FAF5-1F3FF",
		Value:      "🫵🏿",
		Descriptor: "index pointing at the viewer: dark skin tone",
	},
	"1F44D": {
		Key:        "1F44D",
		Value:      "👍",
		Descriptor: "thumbs up",
	},
	"1F44D-1F3FB": {
		Key:        "1F44D-1F3FB",
		Value:      "👍🏻",
		Descriptor: "thumbs up: light skin tone",
	},
	"1F44D-1F3FC": {
		Key:        "1F44D-1F3FC",
		Value:      "👍🏼",
		Descriptor: "thumbs up: medium-light skin tone",
	},
	"1F44D-1F3FD": {
		Key:        "1F44D-1F3FD",
		Value:      "👍🏽",
		Descriptor: "thumbs up: medium skin tone",
	},
	"1F44D-1F3FE": {
		Key:        "1F44D-1F3FE",
		Value:      "👍🏾",
		Descriptor: "thumbs up: medium-dark skin tone",
	},
	"1F44D-1F3FF": {
		Key:        "1F44D-1F3FF",
		Value:      "👍🏿",
		Descriptor: "thumbs up: dark skin tone",
	},
	"1F44E": {
		Key:        "1F44E",
		Value:      "👎",
		Descriptor: "thumbs down",
	},
	"1F44E-1F3FB": {
		Key:        "1F44E-1F3FB",
		Value:      "👎🏻",
		Descriptor: "thumbs down: light skin tone",
	},
	"1F44E-1F3FC": {
		Key:        "1F44E-1F3FC",
		Value:      "👎🏼",
		Descriptor: "thumbs down: medium-light skin tone",
	},
	"1F44E-1F3FD": {
		Key:        "1F44E-1F3FD",
		Value:      "👎🏽",
		Descriptor: "thumbs down: medium skin tone",
	},
	"1F44E-1F3FE": {
		Key:        "1F44E-1F3FE",
		Value:      "👎🏾",
		Descriptor: "thumbs down: medium-dark skin tone",
	},
	"1F44E-1F3FF": {
		Key:        "1F44E-1F3FF",
		Value:      "👎🏿",
		Descriptor: "thumbs down: dark skin tone",
	},
	"270A": {
		Key:        "270A",
		Value:      "✊",
		Descriptor: "raised fist",
	},
	"270A-1F3FB": {
		Key:        "270A-1F3FB",
		Value:      "✊🏻",
		Descriptor: "raised fist: light skin tone",
	},
	"270A-1F3FC": {
		Key:        "270A-1F3FC",
		Value:      "✊🏼",
		Descriptor: "raised fist: medium-light skin tone",
	},
	"270A-1F3FD": {
		Key:        "270A-1F3FD",
		Value:      "✊🏽",
		Descriptor: "raised fist: medium skin tone",
	},
	"270A-1F3FE": {
		Key:        "270A-1F3FE",
		Value:      "✊🏾",
		Descriptor: "raised fist: medium-dark skin tone",
	},
	"270A-1F3FF": {
		Key:        "270A-1F3FF",
		Value:      "✊🏿",
		Descriptor: "raised fist: dark skin tone",
	},
	"1F44A": {
		Key:        "1F44A",
		Value:      "👊",
		Descriptor: "oncoming fist",
	},
	"1F44A-1F3FB": {
		Key:        "1F44A-1F3FB",
		Value:      "👊🏻",
		Descriptor: "oncoming fist: light skin tone",
	},
	"1F44A-1F3FC": {
		Key:        "1F44A-1F3FC",
		Value:      "👊🏼",
		Descriptor: "oncoming fist: medium-light skin tone",
	},
	"1F44A-1F3FD": {
		Key:        "1F44A-1F3FD",
		Value:      "👊🏽",
		Descriptor: "oncoming fist: medium skin tone",
	},
	"1F44A-1F3FE": {
		Key:        "1F44A-1F3FE",
		Value:      "👊🏾",
		Descriptor: "oncoming fist: medium-dark skin tone",
	},
	"1F44A-1F3FF": {
		Key:        "1F44A-1F3FF",
		Value:      "👊🏿",
		Descriptor: "oncoming fist: dark skin tone",
	},
	"1F91B": {
		Key:        "1F91B",
		Value:      "🤛",
		Descriptor: "left-facing fist",
	},
	"1F91B-1F3FB": {
		Key:        "1F91B-1F3FB",
		Value:      "🤛🏻",
		Descriptor: "left-facing fist: light skin tone",
	},
	"1F91B-1F3FC": {
		Key:        "1F91B-1F3FC",
		Value:      "🤛🏼",
		Descriptor: "left-facing fist: medium-light skin tone",
	},
	"1F91B-1F3FD": {
		Key:        "1F91B-1F3FD",
		Value:      "🤛🏽",
		Descriptor: "left-facing fist: medium skin tone",
	},
	"1F91B-1F3FE": {
		Key:        "1F91B-1F3FE",
		Value:      "🤛🏾",
		Descriptor: "left-facing fist: medium-dark skin tone",
	},
	"1F91B-1F3FF": {
		Key:        "1F91B-1F3FF",
		Value:      "🤛🏿",
		Descriptor: "left-facing fist: dark skin tone",
	},
	"1F91C": {
		Key:        "1F91C",
		Value:      "🤜",
		Descriptor: "right-facing fist",
	},
	"1F91C-1F3FB": {
		Key:        "1F91C-1F3FB",
		Value:      "🤜🏻",
		Descriptor: "right-facing fist: light skin tone",
	},
	"1F91C-1F3FC": {
		Key:        "1F91C-1F3FC",
		Value:      "🤜🏼",
		Descriptor: "right-facing fist: medium-light skin tone",
	},
	"1F91C-1F3FD": {
		Key:        "1F91C-1F3FD",
		Value:      "🤜🏽",
		Descriptor: "right-facing fist: medium skin tone",
	},
	"1F91C-1F3FE": {
		Key:        "1F91C-1F3FE",
		Value:      "🤜🏾",
		Descriptor: "right-facing fist: medium-dark skin tone",
	},
	"1F91C-1F3FF": {
		Key:        "1F91C-1F3FF",
		Value:      "🤜🏿",
		Descriptor: "right-facing fist: dark skin tone",
	},
	"1F44F": {
		Key:        "1F44F",
		Value:      "👏",
		Descriptor: "clapping hands",
	},
	"1F44F-1F3FB": {
		Key:        "1F44F-1F3FB",
		Value:      "👏🏻",
		Descriptor: "clapping hands: light skin tone",
	},
	"1F44F-1F3FC": {
		Key:        "1F44F-1F3FC",
		Value:      "👏🏼",
		Descriptor: "clapping hands: medium-light skin tone",
	},
	"1F44F-1F3FD": {
		Key:        "1F44F-1F3FD",
		Value:      "👏🏽",
		Descriptor: "clapping hands: medium skin tone",
	},
	"1F44F-1F3FE": {
		Key:        "1F44F-1F3FE",
		Value:      "👏🏾",
		Descriptor: "clapping hands: medium-dark skin tone",
	},
	"1F44F-1F3FF": {
		Key:        "1F44F-1F3FF",
		Value:      "👏🏿",
		Descriptor: "clapping hands: dark skin tone",
	},
	"1F64C": {
		Key:        "1F64C",
		Value:      "🙌",
		Descriptor: "raising hands",
	},
	"1F64C-1F3FB": {
		Key:        "1F64C-1F3FB",
		Value:      "🙌🏻",
		Descriptor: "raising hands: light skin tone",
	},
	"1F64C-1F3FC": {
		Key:        "1F64C-1F3FC",
		Value:      "🙌🏼",
		Descriptor: "raising hands: medium-light skin tone",
	},
	"1F64C-1F3FD": {
		Key:        "1F64C-1F3FD",
		Value:      "🙌🏽",
		Descriptor: "raising hands: medium skin tone",
	},
	"1F64C-1F3FE": {
		Key:        "1F64C-1F3FE",
		Value:      "🙌🏾",
		Descriptor: "raising hands: medium-dark skin tone",
	},
	"1F64C-1F3FF": {
		Key:        "1F64C-1F3FF",
		Value:      "🙌🏿",
		Descriptor: "raising hands: dark skin tone",
	},
	"1FAF6": {
		Key:        "1FAF6",
		Value:      "🫶",
		Descriptor: "heart hands",
	},
	"1FAF6-1F3FB": {
		Key:        "1FAF6-1F3FB",
		Value:      "🫶🏻",
		Descriptor: "heart hands: light skin tone",
	},
	"1FAF6-1F3FC": {
		Key:        "1FAF6-1F3FC",
		Value:      "🫶🏼",
		Descriptor: "heart hands: medium-light skin tone",
	},
	"1FAF6-1F3FD": {
		Key:        "1FAF6-1F3FD",
		Value:      "🫶🏽",
		Descriptor: "heart hands: medium skin tone",
	},
	"1FAF6-1F3FE": {
		Key:        "1FAF6-1F3FE",
		Value:      "🫶🏾",
		Descriptor: "heart hands: medium-dark skin tone",
	},
	"1FAF6-1F3FF": {
		Key:        "1FAF6-1F3FF",
		Value:      "🫶🏿",
		Descriptor: "heart hands: dark skin tone",
	},
	"1F450": {
		Key:        "1F450",
		Value:      "👐",
		Descriptor: "open hands",
	},
	"1F450-1F3FB": {
		Key:        "1F450-1F3FB",
		Value:      "👐🏻",
		Descriptor: "open hands: light skin tone",
	},
	"1F450-1F3FC": {
		Key:        "1F450-1F3FC",
		Value:      "👐🏼",
		Descriptor: "open hands: medium-light skin tone",
	},
	"1F450-1F3FD": {
		Key:        "1F450-1F3FD",
		Value:      "👐🏽",
		Descriptor: "open hands: medium skin tone",
	},
	"1F450-1F3FE": {
		Key:        "1F450-1F3FE",
		Value:      "👐🏾",
		Descriptor: "open hands: medium-dark skin tone",
	},
	"1F450-1F3FF": {
		Key:        "1F450-1F3FF",
		Value:      "👐🏿",
		Descriptor: "open hands: dark skin tone",
	},
	"1F932": {
		Key:        "1F932",
		Value:      "🤲",
		Descriptor: "palms up together",
	},
	"1F932-1F3FB": {
		Key:        "1F932-1F3FB",
		Value:      "🤲🏻",
		Descriptor: "palms up together: light skin tone",
	},
	"1F932-1F3FC": {
		Key:        "1F932-1F3FC",
		Value:      "🤲🏼",
		Descriptor: "palms up together: medium-light skin tone",
	},
	"1F932-1F3FD": {
		Key:        "1F932-1F3FD",
		Value:      "🤲🏽",
		Descriptor: "palms up together: medium skin tone",
	},
	"1F932-1F3FE": {
		Key:        "1F932-1F3FE",
		Value:      "🤲🏾",
		Descriptor: "palms up together: medium-dark skin tone",
	},
	"1F932-1F3FF": {
		Key:        "1F932-1F3FF",
		Value:      "🤲🏿",
		Descriptor: "palms up together: dark skin tone",
	},
	"1F91D": {
		Key:        "1F91D",
		Value:      "🤝",
		Descriptor: "handshake",
	},
	"1F91D-1F3FB": {
		Key:        "1F91D-1F3FB",
		Value:      "🤝🏻",
		Descriptor: "handshake: light skin tone",
	},
	"1F91D-1F3FC": {
		Key:        "1F91D-1F3FC",
		Value:      "🤝🏼",
		Descriptor: "handshake: medium-light skin tone",
	},
	"1F91D-1F3FD": {
		Key:        "1F91D-1F3FD",
		Value:      "🤝🏽",
		Descriptor: "handshake: medium skin tone",
	},
	"1F91D-1F3FE": {
		Key:        "1F91D-1F3FE",
		Value:      "🤝🏾",
		Descriptor: "handshake: medium-dark skin tone",
	},
	"1F91D-1F3FF": {
		Key:        "1F91D-1F3FF",
		Value:      "🤝🏿",
		Descriptor: "handshake: dark skin tone",
	},
	"1FAF1-1F3FB-200D-1FAF2-1F3FC": {
		Key:        "1FAF1-1F3FB-200D-1FAF2-1F3FC",
		Value:      "🫱🏻‍🫲🏼",
		Descriptor: "handshake: light skin tone, medium-light skin tone",
	},
	"1FAF1-1F3FB-200D-1FAF2-1F3FD": {
		Key:        "1FAF1-1F3FB-200D-1FAF2-1F3FD",
		Value:      "🫱🏻‍🫲🏽",
		Descriptor: "handshake: light skin tone, medium skin tone",
	},
	"1FAF1-1F3FB-200D-1FAF2-1F3FE": {
		Key:        "1FAF1-1F3FB-200D-1FAF2-1F3FE",
		Value:      "🫱🏻‍🫲🏾",
		Descriptor: "handshake: light skin tone, medium-dark skin tone",
	},
	"1FAF1-1F3FB-200D-1FAF2-1F3FF": {
		Key:        "1FAF1-1F3FB-200D-1FAF2-1F3FF",
		Value:      "🫱🏻‍🫲🏿",
		Descriptor: "handshake: light skin tone, dark skin tone",
	},
	"1FAF1-1F3FC-200D-1FAF2-1F3FB": {
		Key:        "1FAF1-1F3FC-200D-1FAF2-1F3FB",
		Value:      "🫱🏼‍🫲🏻",
		Descriptor: "handshake: medium-light skin tone, light skin tone",
	},
	"1FAF1-1F3FC-200D-1FAF2-1F3FD": {
		Key:        "1FAF1-1F3FC-200D-1FAF2-1F3FD",
		Value:      "🫱🏼‍🫲🏽",
		Descriptor: "handshake: medium-light skin tone, medium skin tone",
	},
	"1FAF1-1F3FC-200D-1FAF2-1F3FE": {
		Key:        "1FAF1-1F3FC-200D-1FAF2-1F3FE",
		Value:      "🫱🏼‍🫲🏾",
		Descriptor: "handshake: medium-light skin tone, medium-dark skin tone",
	},
	"1FAF1-1F3FC-200D-1FAF2-1F3FF": {
		Key:        "1FAF1-1F3FC-200D-1FAF2-1F3FF",
		Value:      "🫱🏼‍🫲🏿",
		Descriptor: "handshake: medium-light skin tone, dark skin tone",
	},
	"1FAF1-1F3FD-200D-1FAF2-1F3FB": {
		Key:        "1FAF1-1F3FD-200D-1FAF2-1F3FB",
		Value:      "🫱🏽‍🫲🏻",
		Descriptor: "handshake: medium skin tone, light skin tone",
	},
	"1FAF1-1F3FD-200D-1FAF2-1F3FC": {
		Key:        "1FAF1-1F3FD-200D-1FAF2-1F3FC",
		Value:      "🫱🏽‍🫲🏼",
		Descriptor: "handshake: medium skin tone, medium-light skin tone",
	},
	"1FAF1-1F3FD-200D-1FAF2-1F3FE": {
		Key:        "1FAF1-1F3FD-200D-1FAF2-1F3FE",
		Value:      "🫱🏽‍🫲🏾",
		Descriptor: "handshake: medium skin tone, medium-dark skin tone",
	},
	"1FAF1-1F3FD-200D-1FAF2-1F3FF": {
		Key:        "1FAF1-1F3FD-200D-1FAF2-1F3FF",
		Value:      "🫱🏽‍🫲🏿",
		Descriptor: "handshake: medium skin tone, dark skin tone",
	},
	"1FAF1-1F3FE-200D-1FAF2-1F3FB": {
		Key:        "1FAF1-1F3FE-200D-1FAF2-1F3FB",
		Value:      "🫱🏾‍🫲🏻",
		Descriptor: "handshake: medium-dark skin tone, light skin tone",
	},
	"1FAF1-1F3FE-200D-1FAF2-1F3FC": {
		Key:        "1FAF1-1F3FE-200D-1FAF2-1F3FC",
		Value:      "🫱🏾‍🫲🏼",
		Descriptor: "handshake: medium-dark skin tone, medium-light skin tone",
	},
	"1FAF1-1F3FE-200D-1FAF2-1F3FD": {
		Key:        "1FAF1-1F3FE-200D-1FAF2-1F3FD",
		Value:      "🫱🏾‍🫲🏽",
		Descriptor: "handshake: medium-dark skin tone, medium skin tone",
	},
	"1FAF1-1F3FE-200D-1FAF2-1F3FF": {
		Key:        "1FAF1-1F3FE-200D-1FAF2-1F3FF",
		Value:      "🫱🏾‍🫲🏿",
		Descriptor: "handshake: medium-dark skin tone, dark skin tone",
	},
	"1FAF1-1F3FF-200D-1FAF2-1F3FB": {
		Key:        "1FAF1-1F3FF-200D-1FAF2-1F3FB",
		Value:      "🫱🏿‍🫲🏻",
		Descriptor: "handshake: dark skin tone, light skin tone",
	},
	"1FAF1-1F3FF-200D-1FAF2-1F3FC": {
		Key:        "1FAF1-1F3FF-200D-1FAF2-1F3FC",
		Value:      "🫱🏿‍🫲🏼",
		Descriptor: "handshake: dark skin tone, medium-light skin tone",
	},
	"1FAF1-1F3FF-200D-1FAF2-1F3FD": {
		Key:        "1FAF1-1F3FF-200D-1FAF2-1F3FD",
		Value:      "🫱🏿‍🫲🏽",
		Descriptor: "handshake: dark skin tone, medium skin tone",
	},
	"1FAF1-1F3FF-200D-1FAF2-1F3FE": {
		Key:        "1FAF1-1F3FF-200D-1FAF2-1F3FE",
		Value:      "🫱🏿‍🫲🏾",
		Descriptor: "handshake: dark skin tone, medium-dark skin tone",
	},
	"1F64F": {
		Key:        "1F64F",
		Value:      "🙏",
		Descriptor: "folded hands",
	},
	"1F64F-1F3FB": {
		Key:        "1F64F-1F3FB",
		Value:      "🙏🏻",
		Descriptor: "folded hands: light skin tone",
	},
	"1F64F-1F3FC": {
		Key:        "1F64F-1F3FC",
		Value:      "🙏🏼",
		Descriptor: "folded hands: medium-light skin tone",
	},
	"1F64F-1F3FD": {
		Key:        "1F64F-1F3FD",
		Value:      "🙏🏽",
		Descriptor: "folded hands: medium skin tone",
	},
	"1F64F-1F3FE": {
		Key:        "1F64F-1F3FE",
		Value:      "🙏🏾",
		Descriptor: "folded hands: medium-dark skin tone",
	},
	"1F64F-1F3FF": {
		Key:        "1F64F-1F3FF",
		Value:      "🙏🏿",
		Descriptor: "folded hands: dark skin tone",
	},
	"270D-FE0F": {
		Key:        "270D-FE0F",
		Value:      "✍️",
		Descriptor: "writing hand",
	},
	"270D-1F3FB": {
		Key:        "270D-1F3FB",
		Value:      "✍🏻",
		Descriptor: "writing hand: light skin tone",
	},
	"270D-1F3FC": {
		Key:        "270D-1F3FC",
		Value:      "✍🏼",
		Descriptor: "writing hand: medium-light skin tone",
	},
	"270D-1F3FD": {
		Key:        "270D-1F3FD",
		Value:      "✍🏽",
		Descriptor: "writing hand: medium skin tone",
	},
	"270D-1F3FE": {
		Key:        "270D-1F3FE",
		Value:      "✍🏾",
		Descriptor: "writing hand: medium-dark skin tone",
	},
	"270D-1F3FF": {
		Key:        "270D-1F3FF",
		Value:      "✍🏿",
		Descriptor: "writing hand: dark skin tone",
	},
	"1F485": {
		Key:        "1F485",
		Value:      "💅",
		Descriptor: "nail polish",
	},
	"1F485-1F3FB": {
		Key:        "1F485-1F3FB",
		Value:      "💅🏻",
		Descriptor: "nail polish: light skin tone",
	},
	"1F485-1F3FC": {
		Key:        "1F485-1F3FC",
		Value:      "💅🏼",
		Descriptor: "nail polish: medium-light skin tone",
	},
	"1F485-1F3FD": {
		Key:        "1F485-1F3FD",
		Value:      "💅🏽",
		Descriptor: "nail polish: medium skin tone",
	},
	"1F485-1F3FE": {
		Key:        "1F485-1F3FE",
		Value:      "💅🏾",
		Descriptor: "nail polish: medium-dark skin tone",
	},
	"1F485-1F3FF": {
		Key:        "1F485-1F3FF",
		Value:      "💅🏿",
		Descriptor: "nail polish: dark skin tone",
	},
	"1F933": {
		Key:        "1F933",
		Value:      "🤳",
		Descriptor: "selfie",
	},
	"1F933-1F3FB": {
		Key:        "1F933-1F3FB",
		Value:      "🤳🏻",
		Descriptor: "selfie: light skin tone",
	},
	"1F933-1F3FC": {
		Key:        "1F933-1F3FC",
		Value:      "🤳🏼",
		Descriptor: "selfie: medium-light skin tone",
	},
	"1F933-1F3FD": {
		Key:        "1F933-1F3FD",
		Value:      "🤳🏽",
		Descriptor: "selfie: medium skin tone",
	},
	"1F933-1F3FE": {
		Key:        "1F933-1F3FE",
		Value:      "🤳🏾",
		Descriptor: "selfie: medium-dark skin tone",
	},
	"1F933-1F3FF": {
		Key:        "1F933-1F3FF",
		Value:      "🤳🏿",
		Descriptor: "selfie: dark skin tone",
	},
	"1F4AA": {
		Key:        "1F4AA",
		Value:      "💪",
		Descriptor: "flexed biceps",
	},
	"1F4AA-1F3FB": {
		Key:        "1F4AA-1F3FB",
		Value:      "💪🏻",
		Descriptor: "flexed biceps: light skin tone",
	},
	"1F4AA-1F3FC": {
		Key:        "1F4AA-1F3FC",
		Value:      "💪🏼",
		Descriptor: "flexed biceps: medium-light skin tone",
	},
	"1F4AA-1F3FD": {
		Key:        "1F4AA-1F3FD",
		Value:      "💪🏽",
		Descriptor: "flexed biceps: medium skin tone",
	},
	"1F4AA-1F3FE": {
		Key:        "1F4AA-1F3FE",
		Value:      "💪🏾",
		Descriptor: "flexed biceps: medium-dark skin tone",
	},
	"1F4AA-1F3FF": {
		Key:        "1F4AA-1F3FF",
		Value:      "💪🏿",
		Descriptor: "flexed biceps: dark skin tone",
	},
	"1F9BE": {
		Key:        "1F9BE",
		Value:      "🦾",
		Descriptor: "mechanical arm",
	},
	"1F9BF": {
		Key:        "1F9BF",
		Value:      "🦿",
		Descriptor: "mechanical leg",
	},
	"1F9B5": {
		Key:        "1F9B5",
		Value:      "🦵",
		Descriptor: "leg",
	},
	"1F9B5-1F3FB": {
		Key:        "1F9B5-1F3FB",
		Value:      "🦵🏻",
		Descriptor: "leg: light skin tone",
	},
	"1F9B5-1F3FC": {
		Key:        "1F9B5-1F3FC",
		Value:      "🦵🏼",
		Descriptor: "leg: medium-light skin tone",
	},
	"1F9B5-1F3FD": {
		Key:        "1F9B5-1F3FD",
		Value:      "🦵🏽",
		Descriptor: "leg: medium skin tone",
	},
	"1F9B5-1F3FE": {
		Key:        "1F9B5-1F3FE",
		Value:      "🦵🏾",
		Descriptor: "leg: medium-dark skin tone",
	},
	"1F9B5-1F3FF": {
		Key:        "1F9B5-1F3FF",
		Value:      "🦵🏿",
		Descriptor: "leg: dark skin tone",
	},
	"1F9B6": {
		Key:        "1F9B6",
		Value:      "🦶",
		Descriptor: "foot",
	},
	"1F9B6-1F3FB": {
		Key:        "1F9B6-1F3FB",
		Value:      "🦶🏻",
		Descriptor: "foot: light skin tone",
	},
	"1F9B6-1F3FC": {
		Key:        "1F9B6-1F3FC",
		Value:      "🦶🏼",
		Descriptor: "foot: medium-light skin tone",
	},
	"1F9B6-1F3FD": {
		Key:        "1F9B6-1F3FD",
		Value:      "🦶🏽",
		Descriptor: "foot: medium skin tone",
	},
	"1F9B6-1F3FE": {
		Key:        "1F9B6-1F3FE",
		Value:      "🦶🏾",
		Descriptor: "foot: medium-dark skin tone",
	},
	"1F9B6-1F3FF": {
		Key:        "1F9B6-1F3FF",
		Value:      "🦶🏿",
		Descriptor: "foot: dark skin tone",
	},
	"1F442": {
		Key:        "1F442",
		Value:      "👂",
		Descriptor: "ear",
	},
	"1F442-1F3FB": {
		Key:        "1F442-1F3FB",
		Value:      "👂🏻",
		Descriptor: "ear: light skin tone",
	},
	"1F442-1F3FC": {
		Key:        "1F442-1F3FC",
		Value:      "👂🏼",
		Descriptor: "ear: medium-light skin tone",
	},
	"1F442-1F3FD": {
		Key:        "1F442-1F3FD",
		Value:      "👂🏽",
		Descriptor: "ear: medium skin tone",
	},
	"1F442-1F3FE": {
		Key:        "1F442-1F3FE",
		Value:      "👂🏾",
		Descriptor: "ear: medium-dark skin tone",
	},
	"1F442-1F3FF": {
		Key:        "1F442-1F3FF",
		Value:      "👂🏿",
		Descriptor: "ear: dark skin tone",
	},
	"1F9BB": {
		Key:        "1F9BB",
		Value:      "🦻",
		Descriptor: "ear with hearing aid",
	},
	"1F9BB-1F3FB": {
		Key:        "1F9BB-1F3FB",
		Value:      "🦻🏻",
		Descriptor: "ear with hearing aid: light skin tone",
	},
	"1F9BB-1F3FC": {
		Key:        "1F9BB-1F3FC",
		Value:      "🦻🏼",
		Descriptor: "ear with hearing aid: medium-light skin tone",
	},
	"1F9BB-1F3FD": {
		Key:        "1F9BB-1F3FD",
		Value:      "🦻🏽",
		Descriptor: "ear with hearing aid: medium skin tone",
	},
	"1F9BB-1F3FE": {
		Key:        "1F9BB-1F3FE",
		Value:      "🦻🏾",
		Descriptor: "ear with hearing aid: medium-dark skin tone",
	},
	"1F9BB-1F3FF": {
		Key:        "1F9BB-1F3FF",
		Value:      "🦻🏿",
		Descriptor: "ear with hearing aid: dark skin tone",
	},
	"1F443": {
		Key:        "1F443",
		Value:      "👃",
		Descriptor: "nose",
	},
	"1F443-1F3FB": {
		Key:        "1F443-1F3FB",
		Value:      "👃🏻",
		Descriptor: "nose: light skin tone",
	},
	"1F443-1F3FC": {
		Key:        "1F443-1F3FC",
		Value:      "👃🏼",
		Descriptor: "nose: medium-light skin tone",
	},
	"1F443-1F3FD": {
		Key:        "1F443-1F3FD",
		Value:      "👃🏽",
		Descriptor: "nose: medium skin tone",
	},
	"1F443-1F3FE": {
		Key:        "1F443-1F3FE",
		Value:      "👃🏾",
		Descriptor: "nose: medium-dark skin tone",
	},
	"1F443-1F3FF": {
		Key:        "1F443-1F3FF",
		Value:      "👃🏿",
		Descriptor: "nose: dark skin tone",
	},
	"1F9E0": {
		Key:        "1F9E0",
		Value:      "🧠",
		Descriptor: "brain",
	},
	"1FAC0": {
		Key:        "1FAC0",
		Value:      "🫀",
		Descriptor: "anatomical heart",
	},
	"1FAC1": {
		Key:        "1FAC1",
		Value:      "🫁",
		Descriptor: "lungs",
	},
	"1F9B7": {
		Key:        "1F9B7",
		Value:      "🦷",
		Descriptor: "tooth",
	},
	"1F9B4": {
		Key:        "1F9B4",
		Value:      "🦴",
		Descriptor: "bone",
	},
	"1F440": {
		Key:        "1F440",
		Value:      "👀",
		Descriptor: "eyes",
	},
	"1F441-FE0F": {
		Key:        "1F441-FE0F",
		Value:      "👁️",
		Descriptor: "eye",
	},
	"1F445": {
		Key:        "1F445",
		Value:      "👅",
		Descriptor: "tongue",
	},
	"1F444": {
		Key:        "1F444",
		Value:      "👄",
		Descriptor: "mouth",
	},
	"1FAE6": {
		Key:        "1FAE6",
		Value:      "🫦",
		Descriptor: "biting lip",
	},
	"1F476": {
		Key:        "1F476",
		Value:      "👶",
		Descriptor: "baby",
	},
	"1F476-1F3FB": {
		Key:        "1F476-1F3FB",
		Value:      "👶🏻",
		Descriptor: "baby: light skin tone",
	},
	"1F476-1F3FC": {
		Key:        "1F476-1F3FC",
		Value:      "👶🏼",
		Descriptor: "baby: medium-light skin tone",
	},
	"1F476-1F3FD": {
		Key:        "1F476-1F3FD",
		Value:      "👶🏽",
		Descriptor: "baby: medium skin tone",
	},
	"1F476-1F3FE": {
		Key:        "1F476-1F3FE",
		Value:      "👶🏾",
		Descriptor: "baby: medium-dark skin tone",
	},
	"1F476-1F3FF": {
		Key:        "1F476-1F3FF",
		Value:      "👶🏿",
		Descriptor: "baby: dark skin tone",
	},
	"1F9D2": {
		Key:        "1F9D2",
		Value:      "🧒",
		Descriptor: "child",
	},
	"1F9D2-1F3FB": {
		Key:        "1F9D2-1F3FB",
		Value:      "🧒🏻",
		Descriptor: "child: light skin tone",
	},
	"1F9D2-1F3FC": {
		Key:        "1F9D2-1F3FC",
		Value:      "🧒🏼",
		Descriptor: "child: medium-light skin tone",
	},
	"1F9D2-1F3FD": {
		Key:        "1F9D2-1F3FD",
		Value:      "🧒🏽",
		Descriptor: "child: medium skin tone",
	},
	"1F9D2-1F3FE": {
		Key:        "1F9D2-1F3FE",
		Value:      "🧒🏾",
		Descriptor: "child: medium-dark skin tone",
	},
	"1F9D2-1F3FF": {
		Key:        "1F9D2-1F3FF",
		Value:      "🧒🏿",
		Descriptor: "child: dark skin tone",
	},
	"1F466": {
		Key:        "1F466",
		Value:      "👦",
		Descriptor: "boy",
	},
	"1F466-1F3FB": {
		Key:        "1F466-1F3FB",
		Value:      "👦🏻",
		Descriptor: "boy: light skin tone",
	},
	"1F466-1F3FC": {
		Key:        "1F466-1F3FC",
		Value:      "👦🏼",
		Descriptor: "boy: medium-light skin tone",
	},
	"1F466-1F3FD": {
		Key:        "1F466-1F3FD",
		Value:      "👦🏽",
		Descriptor: "boy: medium skin tone",
	},
	"1F466-1F3FE": {
		Key:        "1F466-1F3FE",
		Value:      "👦🏾",
		Descriptor: "boy: medium-dark skin tone",
	},
	"1F466-1F3FF": {
		Key:        "1F466-1F3FF",
		Value:      "👦🏿",
		Descriptor: "boy: dark skin tone",
	},
	"1F467": {
		Key:        "1F467",
		Value:      "👧",
		Descriptor: "girl",
	},
	"1F467-1F3FB": {
		Key:        "1F467-1F3FB",
		Value:      "👧🏻",
		Descriptor: "girl: light skin tone",
	},
	"1F467-1F3FC": {
		Key:        "1F467-1F3FC",
		Value:      "👧🏼",
		Descriptor: "girl: medium-light skin tone",
	},
	"1F467-1F3FD": {
		Key:        "1F467-1F3FD",
		Value:      "👧🏽",
		Descriptor: "girl: medium skin tone",
	},
	"1F467-1F3FE": {
		Key:        "1F467-1F3FE",
		Value:      "👧🏾",
		Descriptor: "girl: medium-dark skin tone",
	},
	"1F467-1F3FF": {
		Key:        "1F467-1F3FF",
		Value:      "👧🏿",
		Descriptor: "girl: dark skin tone",
	},
	"1F9D1": {
		Key:        "1F9D1",
		Value:      "🧑",
		Descriptor: "person",
	},
	"1F9D1-1F3FB": {
		Key:        "1F9D1-1F3FB",
		Value:      "🧑🏻",
		Descriptor: "person: light skin tone",
	},
	"1F9D1-1F3FC": {
		Key:        "1F9D1-1F3FC",
		Value:      "🧑🏼",
		Descriptor: "person: medium-light skin tone",
	},
	"1F9D1-1F3FD": {
		Key:        "1F9D1-1F3FD",
		Value:      "🧑🏽",
		Descriptor: "person: medium skin tone",
	},
	"1F9D1-1F3FE": {
		Key:        "1F9D1-1F3FE",
		Value:      "🧑🏾",
		Descriptor: "person: medium-dark skin tone",
	},
	"1F9D1-1F3FF": {
		Key:        "1F9D1-1F3FF",
		Value:      "🧑🏿",
		Descriptor: "person: dark skin tone",
	},
	"1F471": {
		Key:        "1F471",
		Value:      "👱",
		Descriptor: "person: blond hair",
	},
	"1F471-1F3FB": {
		Key:        "1F471-1F3FB",
		Value:      "👱🏻",
		Descriptor: "person: light skin tone, blond hair",
	},
	"1F471-1F3FC": {
		Key:        "1F471-1F3FC",
		Value:      "👱🏼",
		Descriptor: "person: medium-light skin tone, blond hair",
	},
	"1F471-1F3FD": {
		Key:        "1F471-1F3FD",
		Value:      "👱🏽",
		Descriptor: "person: medium skin tone, blond hair",
	},
	"1F471-1F3FE": {
		Key:        "1F471-1F3FE",
		Value:      "👱🏾",
		Descriptor: "person: medium-dark skin tone, blond hair",
	},
	"1F471-1F3FF": {
		Key:        "1F471-1F3FF",
		Value:      "👱🏿",
		Descriptor: "person: dark skin tone, blond hair",
	},
	"1F468": {
		Key:        "1F468",
		Value:      "👨",
		Descriptor: "man",
	},
	"1F468-1F3FB": {
		Key:        "1F468-1F3FB",
		Value:      "👨🏻",
		Descriptor: "man: light skin tone",
	},
	"1F468-1F3FC": {
		Key:        "1F468-1F3FC",
		Value:      "👨🏼",
		Descriptor: "man: medium-light skin tone",
	},
	"1F468-1F3FD": {
		Key:        "1F468-1F3FD",
		Value:      "👨🏽",
		Descriptor: "man: medium skin tone",
	},
	"1F468-1F3FE": {
		Key:        "1F468-1F3FE",
		Value:      "👨🏾",
		Descriptor: "man: medium-dark skin tone",
	},
	"1F468-1F3FF": {
		Key:        "1F468-1F3FF",
		Value:      "👨🏿",
		Descriptor: "man: dark skin tone",
	},
	"1F9D4": {
		Key:        "1F9D4",
		Value:      "🧔",
		Descriptor: "person: beard",
	},
	"1F9D4-1F3FB": {
		Key:        "1F9D4-1F3FB",
		Value:      "🧔🏻",
		Descriptor: "person: light skin tone, beard",
	},
	"1F9D4-1F3FC": {
		Key:        "1F9D4-1F3FC",
		Value:      "🧔🏼",
		Descriptor: "person: medium-light skin tone, beard",
	},
	"1F9D4-1F3FD": {
		Key:        "1F9D4-1F3FD",
		Value:      "🧔🏽",
		Descriptor: "person: medium skin tone, beard",
	},
	"1F9D4-1F3FE": {
		Key:        "1F9D4-1F3FE",
		Value:      "🧔🏾",
		Descriptor: "person: medium-dark skin tone, beard",
	},
	"1F9D4-1F3FF": {
		Key:        "1F9D4-1F3FF",
		Value:      "🧔🏿",
		Descriptor: "person: dark skin tone, beard",
	},
	"1F9D4-200D-2642-FE0F": {
		Key:        "1F9D4-200D-2642-FE0F",
		Value:      "🧔‍♂️",
		Descriptor: "man: beard",
	},
	"1F9D4-1F3FB-200D-2642-FE0F": {
		Key:        "1F9D4-1F3FB-200D-2642-FE0F",
		Value:      "🧔🏻‍♂️",
		Descriptor: "man: light skin tone, beard",
	},
	"1F9D4-1F3FC-200D-2642-FE0F": {
		Key:        "1F9D4-1F3FC-200D-2642-FE0F",
		Value:      "🧔🏼‍♂️",
		Descriptor: "man: medium-light skin tone, beard",
	},
	"1F9D4-1F3FD-200D-2642-FE0F": {
		Key:        "1F9D4-1F3FD-200D-2642-FE0F",
		Value:      "🧔🏽‍♂️",
		Descriptor: "man: medium skin tone, beard",
	},
	"1F9D4-1F3FE-200D-2642-FE0F": {
		Key:        "1F9D4-1F3FE-200D-2642-FE0F",
		Value:      "🧔🏾‍♂️",
		Descriptor: "man: medium-dark skin tone, beard",
	},
	"1F9D4-1F3FF-200D-2642-FE0F": {
		Key:        "1F9D4-1F3FF-200D-2642-FE0F",
		Value:      "🧔🏿‍♂️",
		Descriptor: "man: dark skin tone, beard",
	},
	"1F9D4-200D-2640-FE0F": {
		Key:        "1F9D4-200D-2640-FE0F",
		Value:      "🧔‍♀️",
		Descriptor: "woman: beard",
	},
	"1F9D4-1F3FB-200D-2640-FE0F": {
		Key:        "1F9D4-1F3FB-200D-2640-FE0F",
		Value:      "🧔🏻‍♀️",
		Descriptor: "woman: light skin tone, beard",
	},
	"1F9D4-1F3FC-200D-2640-FE0F": {
		Key:        "1F9D4-1F3FC-200D-2640-FE0F",
		Value:      "🧔🏼‍♀️",
		Descriptor: "woman: medium-light skin tone, beard",
	},
	"1F9D4-1F3FD-200D-2640-FE0F": {
		Key:        "1F9D4-1F3FD-200D-2640-FE0F",
		Value:      "🧔🏽‍♀️",
		Descriptor: "woman: medium skin tone, beard",
	},
	"1F9D4-1F3FE-200D-2640-FE0F": {
		Key:        "1F9D4-1F3FE-200D-2640-FE0F",
		Value:      "🧔🏾‍♀️",
		Descriptor: "woman: medium-dark skin tone, beard",
	},
	"1F9D4-1F3FF-200D-2640-FE0F": {
		Key:        "1F9D4-1F3FF-200D-2640-FE0F",
		Value:      "🧔🏿‍♀️",
		Descriptor: "woman: dark skin tone, beard",
	},
	"1F468-200D-1F9B0": {
		Key:        "1F468-200D-1F9B0",
		Value:      "👨‍🦰",
		Descriptor: "man: red hair",
	},
	"1F468-1F3FB-200D-1F9B0": {
		Key:        "1F468-1F3FB-200D-1F9B0",
		Value:      "👨🏻‍🦰",
		Descriptor: "man: light skin tone, red hair",
	},
	"1F468-1F3FC-200D-1F9B0": {
		Key:        "1F468-1F3FC-200D-1F9B0",
		Value:      "👨🏼‍🦰",
		Descriptor: "man: medium-light skin tone, red hair",
	},
	"1F468-1F3FD-200D-1F9B0": {
		Key:        "1F468-1F3FD-200D-1F9B0",
		Value:      "👨🏽‍🦰",
		Descriptor: "man: medium skin tone, red hair",
	},
	"1F468-1F3FE-200D-1F9B0": {
		Key:        "1F468-1F3FE-200D-1F9B0",
		Value:      "👨🏾‍🦰",
		Descriptor: "man: medium-dark skin tone, red hair",
	},
	"1F468-1F3FF-200D-1F9B0": {
		Key:        "1F468-1F3FF-200D-1F9B0",
		Value:      "👨🏿‍🦰",
		Descriptor: "man: dark skin tone, red hair",
	},
	"1F468-200D-1F9B1": {
		Key:        "1F468-200D-1F9B1",
		Value:      "👨‍🦱",
		Descriptor: "man: curly hair",
	},
	"1F468-1F3FB-200D-1F9B1": {
		Key:        "1F468-1F3FB-200D-1F9B1",
		Value:      "👨🏻‍🦱",
		Descriptor: "man: light skin tone, curly hair",
	},
	"1F468-1F3FC-200D-1F9B1": {
		Key:        "1F468-1F3FC-200D-1F9B1",
		Value:      "👨🏼‍🦱",
		Descriptor: "man: medium-light skin tone, curly hair",
	},
	"1F468-1F3FD-200D-1F9B1": {
		Key:        "1F468-1F3FD-200D-1F9B1",
		Value:      "👨🏽‍🦱",
		Descriptor: "man: medium skin tone, curly hair",
	},
	"1F468-1F3FE-200D-1F9B1": {
		Key:        "1F468-1F3FE-200D-1F9B1",
		Value:      "👨🏾‍🦱",
		Descriptor: "man: medium-dark skin tone, curly hair",
	},
	"1F468-1F3FF-200D-1F9B1": {
		Key:        "1F468-1F3FF-200D-1F9B1",
		Value:      "👨🏿‍🦱",
		Descriptor: "man: dark skin tone, curly hair",
	},
	"1F468-200D-1F9B3": {
		Key:        "1F468-200D-1F9B3",
		Value:      "👨‍🦳",
		Descriptor: "man: white hair",
	},
	"1F468-1F3FB-200D-1F9B3": {
		Key:        "1F468-1F3FB-200D-1F9B3",
		Value:      "👨🏻‍🦳",
		Descriptor: "man: light skin tone, white hair",
	},
	"1F468-1F3FC-200D-1F9B3": {
		Key:        "1F468-1F3FC-200D-1F9B3",
		Value:      "👨🏼‍🦳",
		Descriptor: "man: medium-light skin tone, white hair",
	},
	"1F468-1F3FD-200D-1F9B3": {
		Key:        "1F468-1F3FD-200D-1F9B3",
		Value:      "👨🏽‍🦳",
		Descriptor: "man: medium skin tone, white hair",
	},
	"1F468-1F3FE-200D-1F9B3": {
		Key:        "1F468-1F3FE-200D-1F9B3",
		Value:      "👨🏾‍🦳",
		Descriptor: "man: medium-dark skin tone, white hair",
	},
	"1F468-1F3FF-200D-1F9B3": {
		Key:        "1F468-1F3FF-200D-1F9B3",
		Value:      "👨🏿‍🦳",
		Descriptor: "man: dark skin tone, white hair",
	},
	"1F468-200D-1F9B2": {
		Key:        "1F468-200D-1F9B2",
		Value:      "👨‍🦲",
		Descriptor: "man: bald",
	},
	"1F468-1F3FB-200D-1F9B2": {
		Key:        "1F468-1F3FB-200D-1F9B2",
		Value:      "👨🏻‍🦲",
		Descriptor: "man: light skin tone, bald",
	},
	"1F468-1F3FC-200D-1F9B2": {
		Key:        "1F468-1F3FC-200D-1F9B2",
		Value:      "👨🏼‍🦲",
		Descriptor: "man: medium-light skin tone, bald",
	},
	"1F468-1F3FD-200D-1F9B2": {
		Key:        "1F468-1F3FD-200D-1F9B2",
		Value:      "👨🏽‍🦲",
		Descriptor: "man: medium skin tone, bald",
	},
	"1F468-1F3FE-200D-1F9B2": {
		Key:        "1F468-1F3FE-200D-1F9B2",
		Value:      "👨🏾‍🦲",
		Descriptor: "man: medium-dark skin tone, bald",
	},
	"1F468-1F3FF-200D-1F9B2": {
		Key:        "1F468-1F3FF-200D-1F9B2",
		Value:      "👨🏿‍🦲",
		Descriptor: "man: dark skin tone, bald",
	},
	"1F469": {
		Key:        "1F469",
		Value:      "👩",
		Descriptor: "woman",
	},
	"1F469-1F3FB": {
		Key:        "1F469-1F3FB",
		Value:      "👩🏻",
		Descriptor: "woman: light skin tone",
	},
	"1F469-1F3FC": {
		Key:        "1F469-1F3FC",
		Value:      "👩🏼",
		Descriptor: "woman: medium-light skin tone",
	},
	"1F469-1F3FD": {
		Key:        "1F469-1F3FD",
		Value:      "👩🏽",
		Descriptor: "woman: medium skin tone",
	},
	"1F469-1F3FE": {
		Key:        "1F469-1F3FE",
		Value:      "👩🏾",
		Descriptor: "woman: medium-dark skin tone",
	},
	"1F469-1F3FF": {
		Key:        "1F469-1F3FF",
		Value:      "👩🏿",
		Descriptor: "woman: dark skin tone",
	},
	"1F469-200D-1F9B0": {
		Key:        "1F469-200D-1F9B0",
		Value:      "👩‍🦰",
		Descriptor: "woman: red hair",
	},
	"1F469-1F3FB-200D-1F9B0": {
		Key:        "1F469-1F3FB-200D-1F9B0",
		Value:      "👩🏻‍🦰",
		Descriptor: "woman: light skin tone, red hair",
	},
	"1F469-1F3FC-200D-1F9B0": {
		Key:        "1F469-1F3FC-200D-1F9B0",
		Value:      "👩🏼‍🦰",
		Descriptor: "woman: medium-light skin tone, red hair",
	},
	"1F469-1F3FD-200D-1F9B0": {
		Key:        "1F469-1F3FD-200D-1F9B0",
		Value:      "👩🏽‍🦰",
		Descriptor: "woman: medium skin tone, red hair",
	},
	"1F469-1F3FE-200D-1F9B0": {
		Key:        "1F469-1F3FE-200D-1F9B0",
		Value:      "👩🏾‍🦰",
		Descriptor: "woman: medium-dark skin tone, red hair",
	},
	"1F469-1F3FF-200D-1F9B0": {
		Key:        "1F469-1F3FF-200D-1F9B0",
		Value:      "👩🏿‍🦰",
		Descriptor: "woman: dark skin tone, red hair",
	},
	"1F9D1-200D-1F9B0": {
		Key:        "1F9D1-200D-1F9B0",
		Value:      "🧑‍🦰",
		Descriptor: "person: red hair",
	},
	"1F9D1-1F3FB-200D-1F9B0": {
		Key:        "1F9D1-1F3FB-200D-1F9B0",
		Value:      "🧑🏻‍🦰",
		Descriptor: "person: light skin tone, red hair",
	},
	"1F9D1-1F3FC-200D-1F9B0": {
		Key:        "1F9D1-1F3FC-200D-1F9B0",
		Value:      "🧑🏼‍🦰",
		Descriptor: "person: medium-light skin tone, red hair",
	},
	"1F9D1-1F3FD-200D-1F9B0": {
		Key:        "1F9D1-1F3FD-200D-1F9B0",
		Value:      "🧑🏽‍🦰",
		Descriptor: "person: medium skin tone, red hair",
	},
	"1F9D1-1F3FE-200D-1F9B0": {
		Key:        "1F9D1-1F3FE-200D-1F9B0",
		Value:      "🧑🏾‍🦰",
		Descriptor: "person: medium-dark skin tone, red hair",
	},
	"1F9D1-1F3FF-200D-1F9B0": {
		Key:        "1F9D1-1F3FF-200D-1F9B0",
		Value:      "🧑🏿‍🦰",
		Descriptor: "person: dark skin tone, red hair",
	},
	"1F469-200D-1F9B1": {
		Key:        "1F469-200D-1F9B1",
		Value:      "👩‍🦱",
		Descriptor: "woman: curly hair",
	},
	"1F469-1F3FB-200D-1F9B1": {
		Key:        "1F469-1F3FB-200D-1F9B1",
		Value:      "👩🏻‍🦱",
		Descriptor: "woman: light skin tone, curly hair",
	},
	"1F469-1F3FC-200D-1F9B1": {
		Key:        "1F469-1F3FC-200D-1F9B1",
		Value:      "👩🏼‍🦱",
		Descriptor: "woman: medium-light skin tone, curly hair",
	},
	"1F469-1F3FD-200D-1F9B1": {
		Key:        "1F469-1F3FD-200D-1F9B1",
		Value:      "👩🏽‍🦱",
		Descriptor: "woman: medium skin tone, curly hair",
	},
	"1F469-1F3FE-200D-1F9B1": {
		Key:        "1F469-1F3FE-200D-1F9B1",
		Value:      "👩🏾‍🦱",
		Descriptor: "woman: medium-dark skin tone, curly hair",
	},
	"1F469-1F3FF-200D-1F9B1": {
		Key:        "1F469-1F3FF-200D-1F9B1",
		Value:      "👩🏿‍🦱",
		Descriptor: "woman: dark skin tone, curly hair",
	},
	"1F9D1-200D-1F9B1": {
		Key:        "1F9D1-200D-1F9B1",
		Value:      "🧑‍🦱",
		Descriptor: "person: curly hair",
	},
	"1F9D1-1F3FB-200D-1F9B1": {
		Key:        "1F9D1-1F3FB-200D-1F9B1",
		Value:      "🧑🏻‍🦱",
		Descriptor: "person: light skin tone, curly hair",
	},
	"1F9D1-1F3FC-200D-1F9B1": {
		Key:        "1F9D1-1F3FC-200D-1F9B1",
		Value:      "🧑🏼‍🦱",
		Descriptor: "person: medium-light skin tone, curly hair",
	},
	"1F9D1-1F3FD-200D-1F9B1": {
		Key:        "1F9D1-1F3FD-200D-1F9B1",
		Value:      "🧑🏽‍🦱",
		Descriptor: "person: medium skin tone, curly hair",
	},
	"1F9D1-1F3FE-200D-1F9B1": {
		Key:        "1F9D1-1F3FE-200D-1F9B1",
		Value:      "🧑🏾‍🦱",
		Descriptor: "person: medium-dark skin tone, curly hair",
	},
	"1F9D1-1F3FF-200D-1F9B1": {
		Key:        "1F9D1-1F3FF-200D-1F9B1",
		Value:      "🧑🏿‍🦱",
		Descriptor: "person: dark skin tone, curly hair",
	},
	"1F469-200D-1F9B3": {
		Key:        "1F469-200D-1F9B3",
		Value:      "👩‍🦳",
		Descriptor: "woman: white hair",
	},
	"1F469-1F3FB-200D-1F9B3": {
		Key:        "1F469-1F3FB-200D-1F9B3",
		Value:      "👩🏻‍🦳",
		Descriptor: "woman: light skin tone, white hair",
	},
	"1F469-1F3FC-200D-1F9B3": {
		Key:        "1F469-1F3FC-200D-1F9B3",
		Value:      "👩🏼‍🦳",
		Descriptor: "woman: medium-light skin tone, white hair",
	},
	"1F469-1F3FD-200D-1F9B3": {
		Key:        "1F469-1F3FD-200D-1F9B3",
		Value:      "👩🏽‍🦳",
		Descriptor: "woman: medium skin tone, white hair",
	},
	"1F469-1F3FE-200D-1F9B3": {
		Key:        "1F469-1F3FE-200D-1F9B3",
		Value:      "👩🏾‍🦳",
		Descriptor: "woman: medium-dark skin tone, white hair",
	},
	"1F469-1F3FF-200D-1F9B3": {
		Key:        "1F469-1F3FF-200D-1F9B3",
		Value:      "👩🏿‍🦳",
		Descriptor: "woman: dark skin tone, white hair",
	},
	"1F9D1-200D-1F9B3": {
		Key:        "1F9D1-200D-1F9B3",
		Value:      "🧑‍🦳",
		Descriptor: "person: white hair",
	},
	"1F9D1-1F3FB-200D-1F9B3": {
		Key:        "1F9D1-1F3FB-200D-1F9B3",
		Value:      "🧑🏻‍🦳",
		Descriptor: "person: light skin tone, white hair",
	},
	"1F9D1-1F3FC-200D-1F9B3": {
		Key:        "1F9D1-1F3FC-200D-1F9B3",
		Value:      "🧑🏼‍🦳",
		Descriptor: "person: medium-light skin tone, white hair",
	},
	"1F9D1-1F3FD-200D-1F9B3": {
		Key:        "1F9D1-1F3FD-200D-1F9B3",
		Value:      "🧑🏽‍🦳",
		Descriptor: "person: medium skin tone, white hair",
	},
	"1F9D1-1F3FE-200D-1F9B3": {
		Key:        "1F9D1-1F3FE-200D-1F9B3",
		Value:      "🧑🏾‍🦳",
		Descriptor: "person: medium-dark skin tone, white hair",
	},
	"1F9D1-1F3FF-200D-1F9B3": {
		Key:        "1F9D1-1F3FF-200D-1F9B3",
		Value:      "🧑🏿‍🦳",
		Descriptor: "person: dark skin tone, white hair",
	},
	"1F469-200D-1F9B2": {
		Key:        "1F469-200D-1F9B2",
		Value:      "👩‍🦲",
		Descriptor: "woman: bald",
	},
	"1F469-1F3FB-200D-1F9B2": {
		Key:        "1F469-1F3FB-200D-1F9B2",
		Value:      "👩🏻‍🦲",
		Descriptor: "woman: light skin tone, bald",
	},
	"1F469-1F3FC-200D-1F9B2": {
		Key:        "1F469-1F3FC-200D-1F9B2",
		Value:      "👩🏼‍🦲",
		Descriptor: "woman: medium-light skin tone, bald",
	},
	"1F469-1F3FD-200D-1F9B2": {
		Key:        "1F469-1F3FD-200D-1F9B2",
		Value:      "👩🏽‍🦲",
		Descriptor: "woman: medium skin tone, bald",
	},
	"1F469-1F3FE-200D-1F9B2": {
		Key:        "1F469-1F3FE-200D-1F9B2",
		Value:      "👩🏾‍🦲",
		Descriptor: "woman: medium-dark skin tone, bald",
	},
	"1F469-1F3FF-200D-1F9B2": {
		Key:        "1F469-1F3FF-200D-1F9B2",
		Value:      "👩🏿‍🦲",
		Descriptor: "woman: dark skin tone, bald",
	},
	"1F9D1-200D-1F9B2": {
		Key:        "1F9D1-200D-1F9B2",
		Value:      "🧑‍🦲",
		Descriptor: "person: bald",
	},
	"1F9D1-1F3FB-200D-1F9B2": {
		Key:        "1F9D1-1F3FB-200D-1F9B2",
		Value:      "🧑🏻‍🦲",
		Descriptor: "person: light skin tone, bald",
	},
	"1F9D1-1F3FC-200D-1F9B2": {
		Key:        "1F9D1-1F3FC-200D-1F9B2",
		Value:      "🧑🏼‍🦲",
		Descriptor: "person: medium-light skin tone, bald",
	},
	"1F9D1-1F3FD-200D-1F9B2": {
		Key:        "1F9D1-1F3FD-200D-1F9B2",
		Value:      "🧑🏽‍🦲",
		Descriptor: "person: medium skin tone, bald",
	},
	"1F9D1-1F3FE-200D-1F9B2": {
		Key:        "1F9D1-1F3FE-200D-1F9B2",
		Value:      "🧑🏾‍🦲",
		Descriptor: "person: medium-dark skin tone, bald",
	},
	"1F9D1-1F3FF-200D-1F9B2": {
		Key:        "1F9D1-1F3FF-200D-1F9B2",
		Value:      "🧑🏿‍🦲",
		Descriptor: "person: dark skin tone, bald",
	},
	"1F471-200D-2640-FE0F": {
		Key:        "1F471-200D-2640-FE0F",
		Value:      "👱‍♀️",
		Descriptor: "woman: blond hair",
	},
	"1F471-1F3FB-200D-2640-FE0F": {
		Key:        "1F471-1F3FB-200D-2640-FE0F",
		Value:      "👱🏻‍♀️",
		Descriptor: "woman: light skin tone, blond hair",
	},
	"1F471-1F3FC-200D-2640-FE0F": {
		Key:        "1F471-1F3FC-200D-2640-FE0F",
		Value:      "👱🏼‍♀️",
		Descriptor: "woman: medium-light skin tone, blond hair",
	},
	"1F471-1F3FD-200D-2640-FE0F": {
		Key:        "1F471-1F3FD-200D-2640-FE0F",
		Value:      "👱🏽‍♀️",
		Descriptor: "woman: medium skin tone, blond hair",
	},
	"1F471-1F3FE-200D-2640-FE0F": {
		Key:        "1F471-1F3FE-200D-2640-FE0F",
		Value:      "👱🏾‍♀️",
		Descriptor: "woman: medium-dark skin tone, blond hair",
	},
	"1F471-1F3FF-200D-2640-FE0F": {
		Key:        "1F471-1F3FF-200D-2640-FE0F",
		Value:      "👱🏿‍♀️",
		Descriptor: "woman: dark skin tone, blond hair",
	},
	"1F471-200D-2642-FE0F": {
		Key:        "1F471-200D-2642-FE0F",
		Value:      "👱‍♂️",
		Descriptor: "man: blond hair",
	},
	"1F471-1F3FB-200D-2642-FE0F": {
		Key:        "1F471-1F3FB-200D-2642-FE0F",
		Value:      "👱🏻‍♂️",
		Descriptor: "man: light skin tone, blond hair",
	},
	"1F471-1F3FC-200D-2642-FE0F": {
		Key:        "1F471-1F3FC-200D-2642-FE0F",
		Value:      "👱🏼‍♂️",
		Descriptor: "man: medium-light skin tone, blond hair",
	},
	"1F471-1F3FD-200D-2642-FE0F": {
		Key:        "1F471-1F3FD-200D-2642-FE0F",
		Value:      "👱🏽‍♂️",
		Descriptor: "man: medium skin tone, blond hair",
	},
	"1F471-1F3FE-200D-2642-FE0F": {
		Key:        "1F471-1F3FE-200D-2642-FE0F",
		Value:      "👱🏾‍♂️",
		Descriptor: "man: medium-dark skin tone, blond hair",
	},
	"1F471-1F3FF-200D-2642-FE0F": {
		Key:        "1F471-1F3FF-200D-2642-FE0F",
		Value:      "👱🏿‍♂️",
		Descriptor: "man: dark skin tone, blond hair",
	},
	"1F9D3": {
		Key:        "1F9D3",
		Value:      "🧓",
		Descriptor: "older person",
	},
	"1F9D3-1F3FB": {
		Key:        "1F9D3-1F3FB",
		Value:      "🧓🏻",
		Descriptor: "older person: light skin tone",
	},
	"1F9D3-1F3FC": {
		Key:        "1F9D3-1F3FC",
		Value:      "🧓🏼",
		Descriptor: "older person: medium-light skin tone",
	},
	"1F9D3-1F3FD": {
		Key:        "1F9D3-1F3FD",
		Value:      "🧓🏽",
		Descriptor: "older person: medium skin tone",
	},
	"1F9D3-1F3FE": {
		Key:        "1F9D3-1F3FE",
		Value:      "🧓🏾",
		Descriptor: "older person: medium-dark skin tone",
	},
	"1F9D3-1F3FF": {
		Key:        "1F9D3-1F3FF",
		Value:      "🧓🏿",
		Descriptor: "older person: dark skin tone",
	},
	"1F474": {
		Key:        "1F474",
		Value:      "👴",
		Descriptor: "old man",
	},
	"1F474-1F3FB": {
		Key:        "1F474-1F3FB",
		Value:      "👴🏻",
		Descriptor: "old man: light skin tone",
	},
	"1F474-1F3FC": {
		Key:        "1F474-1F3FC",
		Value:      "👴🏼",
		Descriptor: "old man: medium-light skin tone",
	},
	"1F474-1F3FD": {
		Key:        "1F474-1F3FD",
		Value:      "👴🏽",
		Descriptor: "old man: medium skin tone",
	},
	"1F474-1F3FE": {
		Key:        "1F474-1F3FE",
		Value:      "👴🏾",
		Descriptor: "old man: medium-dark skin tone",
	},
	"1F474-1F3FF": {
		Key:        "1F474-1F3FF",
		Value:      "👴🏿",
		Descriptor: "old man: dark skin tone",
	},
	"1F475": {
		Key:        "1F475",
		Value:      "👵",
		Descriptor: "old woman",
	},
	"1F475-1F3FB": {
		Key:        "1F475-1F3FB",
		Value:      "👵🏻",
		Descriptor: "old woman: light skin tone",
	},
	"1F475-1F3FC": {
		Key:        "1F475-1F3FC",
		Value:      "👵🏼",
		Descriptor: "old woman: medium-light skin tone",
	},
	"1F475-1F3FD": {
		Key:        "1F475-1F3FD",
		Value:      "👵🏽",
		Descriptor: "old woman: medium skin tone",
	},
	"1F475-1F3FE": {
		Key:        "1F475-1F3FE",
		Value:      "👵🏾",
		Descriptor: "old woman: medium-dark skin tone",
	},
	"1F475-1F3FF": {
		Key:        "1F475-1F3FF",
		Value:      "👵🏿",
		Descriptor: "old woman: dark skin tone",
	},
	"1F64D": {
		Key:        "1F64D",
		Value:      "🙍",
		Descriptor: "person frowning",
	},
	"1F64D-1F3FB": {
		Key:        "1F64D-1F3FB",
		Value:      "🙍🏻",
		Descriptor: "person frowning: light skin tone",
	},
	"1F64D-1F3FC": {
		Key:        "1F64D-1F3FC",
		Value:      "🙍🏼",
		Descriptor: "person frowning: medium-light skin tone",
	},
	"1F64D-1F3FD": {
		Key:        "1F64D-1F3FD",
		Value:      "🙍🏽",
		Descriptor: "person frowning: medium skin tone",
	},
	"1F64D-1F3FE": {
		Key:        "1F64D-1F3FE",
		Value:      "🙍🏾",
		Descriptor: "person frowning: medium-dark skin tone",
	},
	"1F64D-1F3FF": {
		Key:        "1F64D-1F3FF",
		Value:      "🙍🏿",
		Descriptor: "person frowning: dark skin tone",
	},
	"1F64D-200D-2642-FE0F": {
		Key:        "1F64D-200D-2642-FE0F",
		Value:      "🙍‍♂️",
		Descriptor: "man frowning",
	},
	"1F64D-1F3FB-200D-2642-FE0F": {
		Key:        "1F64D-1F3FB-200D-2642-FE0F",
		Value:      "🙍🏻‍♂️",
		Descriptor: "man frowning: light skin tone",
	},
	"1F64D-1F3FC-200D-2642-FE0F": {
		Key:        "1F64D-1F3FC-200D-2642-FE0F",
		Value:      "🙍🏼‍♂️",
		Descriptor: "man frowning: medium-light skin tone",
	},
	"1F64D-1F3FD-200D-2642-FE0F": {
		Key:        "1F64D-1F3FD-200D-2642-FE0F",
		Value:      "🙍🏽‍♂️",
		Descriptor: "man frowning: medium skin tone",
	},
	"1F64D-1F3FE-200D-2642-FE0F": {
		Key:        "1F64D-1F3FE-200D-2642-FE0F",
		Value:      "🙍🏾‍♂️",
		Descriptor: "man frowning: medium-dark skin tone",
	},
	"1F64D-1F3FF-200D-2642-FE0F": {
		Key:        "1F64D-1F3FF-200D-2642-FE0F",
		Value:      "🙍🏿‍♂️",
		Descriptor: "man frowning: dark skin tone",
	},
	"1F64D-200D-2640-FE0F": {
		Key:        "1F64D-200D-2640-FE0F",
		Value:      "🙍‍♀️",
		Descriptor: "woman frowning",
	},
	"1F64D-1F3FB-200D-2640-FE0F": {
		Key:        "1F64D-1F3FB-200D-2640-FE0F",
		Value:      "🙍🏻‍♀️",
		Descriptor: "woman frowning: light skin tone",
	},
	"1F64D-1F3FC-200D-2640-FE0F": {
		Key:        "1F64D-1F3FC-200D-2640-FE0F",
		Value:      "🙍🏼‍♀️",
		Descriptor: "woman frowning: medium-light skin tone",
	},
	"1F64D-1F3FD-200D-2640-FE0F": {
		Key:        "1F64D-1F3FD-200D-2640-FE0F",
		Value:      "🙍🏽‍♀️",
		Descriptor: "woman frowning: medium skin tone",
	},
	"1F64D-1F3FE-200D-2640-FE0F": {
		Key:        "1F64D-1F3FE-200D-2640-FE0F",
		Value:      "🙍🏾‍♀️",
		Descriptor: "woman frowning: medium-dark skin tone",
	},
	"1F64D-1F3FF-200D-2640-FE0F": {
		Key:        "1F64D-1F3FF-200D-2640-FE0F",
		Value:      "🙍🏿‍♀️",
		Descriptor: "woman frowning: dark skin tone",
	},
	"1F64E": {
		Key:        "1F64E",
		Value:      "🙎",
		Descriptor: "person pouting",
	},
	"1F64E-1F3FB": {
		Key:        "1F64E-1F3FB",
		Value:      "🙎🏻",
		Descriptor: "person pouting: light skin tone",
	},
	"1F64E-1F3FC": {
		Key:        "1F64E-1F3FC",
		Value:      "🙎🏼",
		Descriptor: "person pouting: medium-light skin tone",
	},
	"1F64E-1F3FD": {
		Key:        "1F64E-1F3FD",
		Value:      "🙎🏽",
		Descriptor: "person pouting: medium skin tone",
	},
	"1F64E-1F3FE": {
		Key:        "1F64E-1F3FE",
		Value:      "🙎🏾",
		Descriptor: "person pouting: medium-dark skin tone",
	},
	"1F64E-1F3FF": {
		Key:        "1F64E-1F3FF",
		Value:      "🙎🏿",
		Descriptor: "person pouting: dark skin tone",
	},
	"1F64E-200D-2642-FE0F": {
		Key:        "1F64E-200D-2642-FE0F",
		Value:      "🙎‍♂️",
		Descriptor: "man pouting",
	},
	"1F64E-1F3FB-200D-2642-FE0F": {
		Key:        "1F64E-1F3FB-200D-2642-FE0F",
		Value:      "🙎🏻‍♂️",
		Descriptor: "man pouting: light skin tone",
	},
	"1F64E-1F3FC-200D-2642-FE0F": {
		Key:        "1F64E-1F3FC-200D-2642-FE0F",
		Value:      "🙎🏼‍♂️",
		Descriptor: "man pouting: medium-light skin tone",
	},
	"1F64E-1F3FD-200D-2642-FE0F": {
		Key:        "1F64E-1F3FD-200D-2642-FE0F",
		Value:      "🙎🏽‍♂️",
		Descriptor: "man pouting: medium skin tone",
	},
	"1F64E-1F3FE-200D-2642-FE0F": {
		Key:        "1F64E-1F3FE-200D-2642-FE0F",
		Value:      "🙎🏾‍♂️",
		Descriptor: "man pouting: medium-dark skin tone",
	},
	"1F64E-1F3FF-200D-2642-FE0F": {
		Key:        "1F64E-1F3FF-200D-2642-FE0F",
		Value:      "🙎🏿‍♂️",
		Descriptor: "man pouting: dark skin tone",
	},
	"1F64E-200D-2640-FE0F": {
		Key:        "1F64E-200D-2640-FE0F",
		Value:      "🙎‍♀️",
		Descriptor: "woman pouting",
	},
	"1F64E-1F3FB-200D-2640-FE0F": {
		Key:        "1F64E-1F3FB-200D-2640-FE0F",
		Value:      "🙎🏻‍♀️",
		Descriptor: "woman pouting: light skin tone",
	},
	"1F64E-1F3FC-200D-2640-FE0F": {
		Key:        "1F64E-1F3FC-200D-2640-FE0F",
		Value:      "🙎🏼‍♀️",
		Descriptor: "woman pouting: medium-light skin tone",
	},
	"1F64E-1F3FD-200D-2640-FE0F": {
		Key:        "1F64E-1F3FD-200D-2640-FE0F",
		Value:      "🙎🏽‍♀️",
		Descriptor: "woman pouting: medium skin tone",
	},
	"1F64E-1F3FE-200D-2640-FE0F": {
		Key:        "1F64E-1F3FE-200D-2640-FE0F",
		Value:      "🙎🏾‍♀️",
		Descriptor: "woman pouting: medium-dark skin tone",
	},
	"1F64E-1F3FF-200D-2640-FE0F": {
		Key:        "1F64E-1F3FF-200D-2640-FE0F",
		Value:      "🙎🏿‍♀️",
		Descriptor: "woman pouting: dark skin tone",
	},
	"1F645": {
		Key:        "1F645",
		Value:      "🙅",
		Descriptor: "person gesturing NO",
	},
	"1F645-1F3FB": {
		Key:        "1F645-1F3FB",
		Value:      "🙅🏻",
		Descriptor: "person gesturing NO: light skin tone",
	},
	"1F645-1F3FC": {
		Key:        "1F645-1F3FC",
		Value:      "🙅🏼",
		Descriptor: "person gesturing NO: medium-light skin tone",
	},
	"1F645-1F3FD": {
		Key:        "1F645-1F3FD",
		Value:      "🙅🏽",
		Descriptor: "person gesturing NO: medium skin tone",
	},
	"1F645-1F3FE": {
		Key:        "1F645-1F3FE",
		Value:      "🙅🏾",
		Descriptor: "person gesturing NO: medium-dark skin tone",
	},
	"1F645-1F3FF": {
		Key:        "1F645-1F3FF",
		Value:      "🙅🏿",
		Descriptor: "person gesturing NO: dark skin tone",
	},
	"1F645-200D-2642-FE0F": {
		Key:        "1F645-200D-2642-FE0F",
		Value:      "🙅‍♂️",
		Descriptor: "man gesturing NO",
	},
	"1F645-1F3FB-200D-2642-FE0F": {
		Key:        "1F645-1F3FB-200D-2642-FE0F",
		Value:      "🙅🏻‍♂️",
		Descriptor: "man gesturing NO: light skin tone",
	},
	"1F645-1F3FC-200D-2642-FE0F": {
		Key:        "1F645-1F3FC-200D-2642-FE0F",
		Value:      "🙅🏼‍♂️",
		Descriptor: "man gesturing NO: medium-light skin tone",
	},
	"1F645-1F3FD-200D-2642-FE0F": {
		Key:        "1F645-1F3FD-200D-2642-FE0F",
		Value:      "🙅🏽‍♂️",
		Descriptor: "man gesturing NO: medium skin tone",
	},
	"1F645-1F3FE-200D-2642-FE0F": {
		Key:        "1F645-1F3FE-200D-2642-FE0F",
		Value:      "🙅🏾‍♂️",
		Descriptor: "man gesturing NO: medium-dark skin tone",
	},
	"1F645-1F3FF-200D-2642-FE0F": {
		Key:        "1F645-1F3FF-200D-2642-FE0F",
		Value:      "🙅🏿‍♂️",
		Descriptor: "man gesturing NO: dark skin tone",
	},
	"1F645-200D-2640-FE0F": {
		Key:        "1F645-200D-2640-FE0F",
		Value:      "🙅‍♀️",
		Descriptor: "woman gesturing NO",
	},
	"1F645-1F3FB-200D-2640-FE0F": {
		Key:        "1F645-1F3FB-200D-2640-FE0F",
		Value:      "🙅🏻‍♀️",
		Descriptor: "woman gesturing NO: light skin tone",
	},
	"1F645-1F3FC-200D-2640-FE0F": {
		Key:        "1F645-1F3FC-200D-2640-FE0F",
		Value:      "🙅🏼‍♀️",
		Descriptor: "woman gesturing NO: medium-light skin tone",
	},
	"1F645-1F3FD-200D-2640-FE0F": {
		Key:        "1F645-1F3FD-200D-2640-FE0F",
		Value:      "🙅🏽‍♀️",
		Descriptor: "woman gesturing NO: medium skin tone",
	},
	"1F645-1F3FE-200D-2640-FE0F": {
		Key:        "1F645-1F3FE-200D-2640-FE0F",
		Value:      "🙅🏾‍♀️",
		Descriptor: "woman gesturing NO: medium-dark skin tone",
	},
	"1F645-1F3FF-200D-2640-FE0F": {
		Key:        "1F645-1F3FF-200D-2640-FE0F",
		Value:      "🙅🏿‍♀️",
		Descriptor: "woman gesturing NO: dark skin tone",
	},
	"1F646": {
		Key:        "1F646",
		Value:      "🙆",
		Descriptor: "person gesturing OK",
	},
	"1F646-1F3FB": {
		Key:        "1F646-1F3FB",
		Value:      "🙆🏻",
		Descriptor: "person gesturing OK: light skin tone",
	},
	"1F646-1F3FC": {
		Key:        "1F646-1F3FC",
		Value:      "🙆🏼",
		Descriptor: "person gesturing OK: medium-light skin tone",
	},
	"1F646-1F3FD": {
		Key:        "1F646-1F3FD",
		Value:      "🙆🏽",
		Descriptor: "person gesturing OK: medium skin tone",
	},
	"1F646-1F3FE": {
		Key:        "1F646-1F3FE",
		Value:      "🙆🏾",
		Descriptor: "person gesturing OK: medium-dark skin tone",
	},
	"1F646-1F3FF": {
		Key:        "1F646-1F3FF",
		Value:      "🙆🏿",
		Descriptor: "person gesturing OK: dark skin tone",
	},
	"1F646-200D-2642-FE0F": {
		Key:        "1F646-200D-2642-FE0F",
		Value:      "🙆‍♂️",
		Descriptor: "man gesturing OK",
	},
	"1F646-1F3FB-200D-2642-FE0F": {
		Key:        "1F646-1F3FB-200D-2642-FE0F",
		Value:      "🙆🏻‍♂️",
		Descriptor: "man gesturing OK: light skin tone",
	},
	"1F646-1F3FC-200D-2642-FE0F": {
		Key:        "1F646-1F3FC-200D-2642-FE0F",
		Value:      "🙆🏼‍♂️",
		Descriptor: "man gesturing OK: medium-light skin tone",
	},
	"1F646-1F3FD-200D-2642-FE0F": {
		Key:        "1F646-1F3FD-200D-2642-FE0F",
		Value:      "🙆🏽‍♂️",
		Descriptor: "man gesturing OK: medium skin tone",
	},
	"1F646-1F3FE-200D-2642-FE0F": {
		Key:        "1F646-1F3FE-200D-2642-FE0F",
		Value:      "🙆🏾‍♂️",
		Descriptor: "man gesturing OK: medium-dark skin tone",
	},
	"1F646-1F3FF-200D-2642-FE0F": {
		Key:        "1F646-1F3FF-200D-2642-FE0F",
		Value:      "🙆🏿‍♂️",
		Descriptor: "man gesturing OK: dark skin tone",
	},
	"1F646-200D-2640-FE0F": {
		Key:        "1F646-200D-2640-FE0F",
		Value:      "🙆‍♀️",
		Descriptor: "woman gesturing OK",
	},
	"1F646-1F3FB-200D-2640-FE0F": {
		Key:        "1F646-1F3FB-200D-2640-FE0F",
		Value:      "🙆🏻‍♀️",
		Descriptor: "woman gesturing OK: light skin tone",
	},
	"1F646-1F3FC-200D-2640-FE0F": {
		Key:        "1F646-1F3FC-200D-2640-FE0F",
		Value:      "🙆🏼‍♀️",
		Descriptor: "woman gesturing OK: medium-light skin tone",
	},
	"1F646-1F3FD-200D-2640-FE0F": {
		Key:        "1F646-1F3FD-200D-2640-FE0F",
		Value:      "🙆🏽‍♀️",
		Descriptor: "woman gesturing OK: medium skin tone",
	},
	"1F646-1F3FE-200D-2640-FE0F": {
		Key:        "1F646-1F3FE-200D-2640-FE0F",
		Value:      "🙆🏾‍♀️",
		Descriptor: "woman gesturing OK: medium-dark skin tone",
	},
	"1F646-1F3FF-200D-2640-FE0F": {
		Key:        "1F646-1F3FF-200D-2640-FE0F",
		Value:      "🙆🏿‍♀️",
		Descriptor: "woman gesturing OK: dark skin tone",
	},
	"1F481": {
		Key:        "1F481",
		Value:      "💁",
		Descriptor: "person tipping hand",
	},
	"1F481-1F3FB": {
		Key:        "1F481-1F3FB",
		Value:      "💁🏻",
		Descriptor: "person tipping hand: light skin tone",
	},
	"1F481-1F3FC": {
		Key:        "1F481-1F3FC",
		Value:      "💁🏼",
		Descriptor: "person tipping hand: medium-light skin tone",
	},
	"1F481-1F3FD": {
		Key:        "1F481-1F3FD",
		Value:      "💁🏽",
		Descriptor: "person tipping hand: medium skin tone",
	},
	"1F481-1F3FE": {
		Key:        "1F481-1F3FE",
		Value:      "💁🏾",
		Descriptor: "person tipping hand: medium-dark skin tone",
	},
	"1F481-1F3FF": {
		Key:        "1F481-1F3FF",
		Value:      "💁🏿",
		Descriptor: "person tipping hand: dark skin tone",
	},
	"1F481-200D-2642-FE0F": {
		Key:        "1F481-200D-2642-FE0F",
		Value:      "💁‍♂️",
		Descriptor: "man tipping hand",
	},
	"1F481-1F3FB-200D-2642-FE0F": {
		Key:        "1F481-1F3FB-200D-2642-FE0F",
		Value:      "💁🏻‍♂️",
		Descriptor: "man tipping hand: light skin tone",
	},
	"1F481-1F3FC-200D-2642-FE0F": {
		Key:        "1F481-1F3FC-200D-2642-FE0F",
		Value:      "💁🏼‍♂️",
		Descriptor: "man tipping hand: medium-light skin tone",
	},
	"1F481-1F3FD-200D-2642-FE0F": {
		Key:        "1F481-1F3FD-200D-2642-FE0F",
		Value:      "💁🏽‍♂️",
		Descriptor: "man tipping hand: medium skin tone",
	},
	"1F481-1F3FE-200D-2642-FE0F": {
		Key:        "1F481-1F3FE-200D-2642-FE0F",
		Value:      "💁🏾‍♂️",
		Descriptor: "man tipping hand: medium-dark skin tone",
	},
	"1F481-1F3FF-200D-2642-FE0F": {
		Key:        "1F481-1F3FF-200D-2642-FE0F",
		Value:      "💁🏿‍♂️",
		Descriptor: "man tipping hand: dark skin tone",
	},
	"1F481-200D-2640-FE0F": {
		Key:        "1F481-200D-2640-FE0F",
		Value:      "💁‍♀️",
		Descriptor: "woman tipping hand",
	},
	"1F481-1F3FB-200D-2640-FE0F": {
		Key:        "1F481-1F3FB-200D-2640-FE0F",
		Value:      "💁🏻‍♀️",
		Descriptor: "woman tipping hand: light skin tone",
	},
	"1F481-1F3FC-200D-2640-FE0F": {
		Key:        "1F481-1F3FC-200D-2640-FE0F",
		Value:      "💁🏼‍♀️",
		Descriptor: "woman tipping hand: medium-light skin tone",
	},
	"1F481-1F3FD-200D-2640-FE0F": {
		Key:        "1F481-1F3FD-200D-2640-FE0F",
		Value:      "💁🏽‍♀️",
		Descriptor: "woman tipping hand: medium skin tone",
	},
	"1F481-1F3FE-200D-2640-FE0F": {
		Key:        "1F481-1F3FE-200D-2640-FE0F",
		Value:      "💁🏾‍♀️",
		Descriptor: "woman tipping hand: medium-dark skin tone",
	},
	"1F481-1F3FF-200D-2640-FE0F": {
		Key:        "1F481-1F3FF-200D-2640-FE0F",
		Value:      "💁🏿‍♀️",
		Descriptor: "woman tipping hand: dark skin tone",
	},
	"1F64B": {
		Key:        "1F64B",
		Value:      "🙋",
		Descriptor: "person raising hand",
	},
	"1F64B-1F3FB": {
		Key:        "1F64B-1F3FB",
		Value:      "🙋🏻",
		Descriptor: "person raising hand: light skin tone",
	},
	"1F64B-1F3FC": {
		Key:        "1F64B-1F3FC",
		Value:      "🙋🏼",
		Descriptor: "person raising hand: medium-light skin tone",
	},
	"1F64B-1F3FD": {
		Key:        "1F64B-1F3FD",
		Value:      "🙋🏽",
		Descriptor: "person raising hand: medium skin tone",
	},
	"1F64B-1F3FE": {
		Key:        "1F64B-1F3FE",
		Value:      "🙋🏾",
		Descriptor: "person raising hand: medium-dark skin tone",
	},
	"1F64B-1F3FF": {
		Key:        "1F64B-1F3FF",
		Value:      "🙋🏿",
		Descriptor: "person raising hand: dark skin tone",
	},
	"1F64B-200D-2642-FE0F": {
		Key:        "1F64B-200D-2642-FE0F",
		Value:      "🙋‍♂️",
		Descriptor: "man raising hand",
	},
	"1F64B-1F3FB-200D-2642-FE0F": {
		Key:        "1F64B-1F3FB-200D-2642-FE0F",
		Value:      "🙋🏻‍♂️",
		Descriptor: "man raising hand: light skin tone",
	},
	"1F64B-1F3FC-200D-2642-FE0F": {
		Key:        "1F64B-1F3FC-200D-2642-FE0F",
		Value:      "🙋🏼‍♂️",
		Descriptor: "man raising hand: medium-light skin tone",
	},
	"1F64B-1F3FD-200D-2642-FE0F": {
		Key:        "1F64B-1F3FD-200D-2642-FE0F",
		Value:      "🙋🏽‍♂️",
		Descriptor: "man raising hand: medium skin tone",
	},
	"1F64B-1F3FE-200D-2642-FE0F": {
		Key:        "1F64B-1F3FE-200D-2642-FE0F",
		Value:      "🙋🏾‍♂️",
		Descriptor: "man raising hand: medium-dark skin tone",
	},
	"1F64B-1F3FF-200D-2642-FE0F": {
		Key:        "1F64B-1F3FF-200D-2642-FE0F",
		Value:      "🙋🏿‍♂️",
		Descriptor: "man raising hand: dark skin tone",
	},
	"1F64B-200D-2640-FE0F": {
		Key:        "1F64B-200D-2640-FE0F",
		Value:      "🙋‍♀️",
		Descriptor: "woman raising hand",
	},
	"1F64B-1F3FB-200D-2640-FE0F": {
		Key:        "1F64B-1F3FB-200D-2640-FE0F",
		Value:      "🙋🏻‍♀️",
		Descriptor: "woman raising hand: light skin tone",
	},
	"1F64B-1F3FC-200D-2640-FE0F": {
		Key:        "1F64B-1F3FC-200D-2640-FE0F",
		Value:      "🙋🏼‍♀️",
		Descriptor: "woman raising hand: medium-light skin tone",
	},
	"1F64B-1F3FD-200D-2640-FE0F": {
		Key:        "1F64B-1F3FD-200D-2640-FE0F",
		Value:      "🙋🏽‍♀️",
		Descriptor: "woman raising hand: medium skin tone",
	},
	"1F64B-1F3FE-200D-2640-FE0F": {
		Key:        "1F64B-1F3FE-200D-2640-FE0F",
		Value:      "🙋🏾‍♀️",
		Descriptor: "woman raising hand: medium-dark skin tone",
	},
	"1F64B-1F3FF-200D-2640-FE0F": {
		Key:        "1F64B-1F3FF-200D-2640-FE0F",
		Value:      "🙋🏿‍♀️",
		Descriptor: "woman raising hand: dark skin tone",
	},
	"1F9CF": {
		Key:        "1F9CF",
		Value:      "🧏",
		Descriptor: "deaf person",
	},
	"1F9CF-1F3FB": {
		Key:        "1F9CF-1F3FB",
		Value:      "🧏🏻",
		Descriptor: "deaf person: light skin tone",
	},
	"1F9CF-1F3FC": {
		Key:        "1F9CF-1F3FC",
		Value:      "🧏🏼",
		Descriptor: "deaf person: medium-light skin tone",
	},
	"1F9CF-1F3FD": {
		Key:        "1F9CF-1F3FD",
		Value:      "🧏🏽",
		Descriptor: "deaf person: medium skin tone",
	},
	"1F9CF-1F3FE": {
		Key:        "1F9CF-1F3FE",
		Value:      "🧏🏾",
		Descriptor: "deaf person: medium-dark skin tone",
	},
	"1F9CF-1F3FF": {
		Key:        "1F9CF-1F3FF",
		Value:      "🧏🏿",
		Descriptor: "deaf person: dark skin tone",
	},
	"1F9CF-200D-2642-FE0F": {
		Key:        "1F9CF-200D-2642-FE0F",
		Value:      "🧏‍♂️",
		Descriptor: "deaf man",
	},
	"1F9CF-1F3FB-200D-2642-FE0F": {
		Key:        "1F9CF-1F3FB-200D-2642-FE0F",
		Value:      "🧏🏻‍♂️",
		Descriptor: "deaf man: light skin tone",
	},
	"1F9CF-1F3FC-200D-2642-FE0F": {
		Key:        "1F9CF-1F3FC-200D-2642-FE0F",
		Value:      "🧏🏼‍♂️",
		Descriptor: "deaf man: medium-light skin tone",
	},
	"1F9CF-1F3FD-200D-2642-FE0F": {
		Key:        "1F9CF-1F3FD-200D-2642-FE0F",
		Value:      "🧏🏽‍♂️",
		Descriptor: "deaf man: medium skin tone",
	},
	"1F9CF-1F3FE-200D-2642-FE0F": {
		Key:        "1F9CF-1F3FE-200D-2642-FE0F",
		Value:      "🧏🏾‍♂️",
		Descriptor: "deaf man: medium-dark skin tone",
	},
	"1F9CF-1F3FF-200D-2642-FE0F": {
		Key:        "1F9CF-1F3FF-200D-2642-FE0F",
		Value:      "🧏🏿‍♂️",
		Descriptor: "deaf man: dark skin tone",
	},
	"1F9CF-200D-2640-FE0F": {
		Key:        "1F9CF-200D-2640-FE0F",
		Value:      "🧏‍♀️",
		Descriptor: "deaf woman",
	},
	"1F9CF-1F3FB-200D-2640-FE0F": {
		Key:        "1F9CF-1F3FB-200D-2640-FE0F",
		Value:      "🧏🏻‍♀️",
		Descriptor: "deaf woman: light skin tone",
	},
	"1F9CF-1F3FC-200D-2640-FE0F": {
		Key:        "1F9CF-1F3FC-200D-2640-FE0F",
		Value:      "🧏🏼‍♀️",
		Descriptor: "deaf woman: medium-light skin tone",
	},
	"1F9CF-1F3FD-200D-2640-FE0F": {
		Key:        "1F9CF-1F3FD-200D-2640-FE0F",
		Value:      "🧏🏽‍♀️",
		Descriptor: "deaf woman: medium skin tone",
	},
	"1F9CF-1F3FE-200D-2640-FE0F": {
		Key:        "1F9CF-1F3FE-200D-2640-FE0F",
		Value:      "🧏🏾‍♀️",
		Descriptor: "deaf woman: medium-dark skin tone",
	},
	"1F9CF-1F3FF-200D-2640-FE0F": {
		Key:        "1F9CF-1F3FF-200D-2640-FE0F",
		Value:      "🧏🏿‍♀️",
		Descriptor: "deaf woman: dark skin tone",
	},
	"1F647": {
		Key:        "1F647",
		Value:      "🙇",
		Descriptor: "person bowing",
	},
	"1F647-1F3FB": {
		Key:        "1F647-1F3FB",
		Value:      "🙇🏻",
		Descriptor: "person bowing: light skin tone",
	},
	"1F647-1F3FC": {
		Key:        "1F647-1F3FC",
		Value:      "🙇🏼",
		Descriptor: "person bowing: medium-light skin tone",
	},
	"1F647-1F3FD": {
		Key:        "1F647-1F3FD",
		Value:      "🙇🏽",
		Descriptor: "person bowing: medium skin tone",
	},
	"1F647-1F3FE": {
		Key:        "1F647-1F3FE",
		Value:      "🙇🏾",
		Descriptor: "person bowing: medium-dark skin tone",
	},
	"1F647-1F3FF": {
		Key:        "1F647-1F3FF",
		Value:      "🙇🏿",
		Descriptor: "person bowing: dark skin tone",
	},
	"1F647-200D-2642-FE0F": {
		Key:        "1F647-200D-2642-FE0F",
		Value:      "🙇‍♂️",
		Descriptor: "man bowing",
	},
	"1F647-1F3FB-200D-2642-FE0F": {
		Key:        "1F647-1F3FB-200D-2642-FE0F",
		Value:      "🙇🏻‍♂️",
		Descriptor: "man bowing: light skin tone",
	},
	"1F647-1F3FC-200D-2642-FE0F": {
		Key:        "1F647-1F3FC-200D-2642-FE0F",
		Value:      "🙇🏼‍♂️",
		Descriptor: "man bowing: medium-light skin tone",
	},
	"1F647-1F3FD-200D-2642-FE0F": {
		Key:        "1F647-1F3FD-200D-2642-FE0F",
		Value:      "🙇🏽‍♂️",
		Descriptor: "man bowing: medium skin tone",
	},
	"1F647-1F3FE-200D-2642-FE0F": {
		Key:        "1F647-1F3FE-200D-2642-FE0F",
		Value:      "🙇🏾‍♂️",
		Descriptor: "man bowing: medium-dark skin tone",
	},
	"1F647-1F3FF-200D-2642-FE0F": {
		Key:        "1F647-1F3FF-200D-2642-FE0F",
		Value:      "🙇🏿‍♂️",
		Descriptor: "man bowing: dark skin tone",
	},
	"1F647-200D-2640-FE0F": {
		Key:        "1F647-200D-2640-FE0F",
		Value:      "🙇‍♀️",
		Descriptor: "woman bowing",
	},
	"1F647-1F3FB-200D-2640-FE0F": {
		Key:        "1F647-1F3FB-200D-2640-FE0F",
		Value:      "🙇🏻‍♀️",
		Descriptor: "woman bowing: light skin tone",
	},
	"1F647-1F3FC-200D-2640-FE0F": {
		Key:        "1F647-1F3FC-200D-2640-FE0F",
		Value:      "🙇🏼‍♀️",
		Descriptor: "woman bowing: medium-light skin tone",
	},
	"1F647-1F3FD-200D-2640-FE0F": {
		Key:        "1F647-1F3FD-200D-2640-FE0F",
		Value:      "🙇🏽‍♀️",
		Descriptor: "woman bowing: medium skin tone",
	},
	"1F647-1F3FE-200D-2640-FE0F": {
		Key:        "1F647-1F3FE-200D-2640-FE0F",
		Value:      "🙇🏾‍♀️",
		Descriptor: "woman bowing: medium-dark skin tone",
	},
	"1F647-1F3FF-200D-2640-FE0F": {
		Key:        "1F647-1F3FF-200D-2640-FE0F",
		Value:      "🙇🏿‍♀️",
		Descriptor: "woman bowing: dark skin tone",
	},
	"1F926": {
		Key:        "1F926",
		Value:      "🤦",
		Descriptor: "person facepalming",
	},
	"1F926-1F3FB": {
		Key:        "1F926-1F3FB",
		Value:      "🤦🏻",
		Descriptor: "person facepalming: light skin tone",
	},
	"1F926-1F3FC": {
		Key:        "1F926-1F3FC",
		Value:      "🤦🏼",
		Descriptor: "person facepalming: medium-light skin tone",
	},
	"1F926-1F3FD": {
		Key:        "1F926-1F3FD",
		Value:      "🤦🏽",
		Descriptor: "person facepalming: medium skin tone",
	},
	"1F926-1F3FE": {
		Key:        "1F926-1F3FE",
		Value:      "🤦🏾",
		Descriptor: "person facepalming: medium-dark skin tone",
	},
	"1F926-1F3FF": {
		Key:        "1F926-1F3FF",
		Value:      "🤦🏿",
		Descriptor: "person facepalming: dark skin tone",
	},
	"1F926-200D-2642-FE0F": {
		Key:        "1F926-200D-2642-FE0F",
		Value:      "🤦‍♂️",
		Descriptor: "man facepalming",
	},
	"1F926-1F3FB-200D-2642-FE0F": {
		Key:        "1F926-1F3FB-200D-2642-FE0F",
		Value:      "🤦🏻‍♂️",
		Descriptor: "man facepalming: light skin tone",
	},
	"1F926-1F3FC-200D-2642-FE0F": {
		Key:        "1F926-1F3FC-200D-2642-FE0F",
		Value:      "🤦🏼‍♂️",
		Descriptor: "man facepalming: medium-light skin tone",
	},
	"1F926-1F3FD-200D-2642-FE0F": {
		Key:        "1F926-1F3FD-200D-2642-FE0F",
		Value:      "🤦🏽‍♂️",
		Descriptor: "man facepalming: medium skin tone",
	},
	"1F926-1F3FE-200D-2642-FE0F": {
		Key:        "1F926-1F3FE-200D-2642-FE0F",
		Value:      "🤦🏾‍♂️",
		Descriptor: "man facepalming: medium-dark skin tone",
	},
	"1F926-1F3FF-200D-2642-FE0F": {
		Key:        "1F926-1F3FF-200D-2642-FE0F",
		Value:      "🤦🏿‍♂️",
		Descriptor: "man facepalming: dark skin tone",
	},
	"1F926-200D-2640-FE0F": {
		Key:        "1F926-200D-2640-FE0F",
		Value:      "🤦‍♀️",
		Descriptor: "woman facepalming",
	},
	"1F926-1F3FB-200D-2640-FE0F": {
		Key:        "1F926-1F3FB-200D-2640-FE0F",
		Value:      "🤦🏻‍♀️",
		Descriptor: "woman facepalming: light skin tone",
	},
	"1F926-1F3FC-200D-2640-FE0F": {
		Key:        "1F926-1F3FC-200D-2640-FE0F",
		Value:      "🤦🏼‍♀️",
		Descriptor: "woman facepalming: medium-light skin tone",
	},
	"1F926-1F3FD-200D-2640-FE0F": {
		Key:        "1F926-1F3FD-200D-2640-FE0F",
		Value:      "🤦🏽‍♀️",
		Descriptor: "woman facepalming: medium skin tone",
	},
	"1F926-1F3FE-200D-2640-FE0F": {
		Key:        "1F926-1F3FE-200D-2640-FE0F",
		Value:      "🤦🏾‍♀️",
		Descriptor: "woman facepalming: medium-dark skin tone",
	},
	"1F926-1F3FF-200D-2640-FE0F": {
		Key:        "1F926-1F3FF-200D-2640-FE0F",
		Value:      "🤦🏿‍♀️",
		Descriptor: "woman facepalming: dark skin tone",
	},
	"1F937": {
		Key:        "1F937",
		Value:      "🤷",
		Descriptor: "person shrugging",
	},
	"1F937-1F3FB": {
		Key:        "1F937-1F3FB",
		Value:      "🤷🏻",
		Descriptor: "person shrugging: light skin tone",
	},
	"1F937-1F3FC": {
		Key:        "1F937-1F3FC",
		Value:      "🤷🏼",
		Descriptor: "person shrugging: medium-light skin tone",
	},
	"1F937-1F3FD": {
		Key:        "1F937-1F3FD",
		Value:      "🤷🏽",
		Descriptor: "person shrugging: medium skin tone",
	},
	"1F937-1F3FE": {
		Key:        "1F937-1F3FE",
		Value:      "🤷🏾",
		Descriptor: "person shrugging: medium-dark skin tone",
	},
	"1F937-1F3FF": {
		Key:        "1F937-1F3FF",
		Value:      "🤷🏿",
		Descriptor: "person shrugging: dark skin tone",
	},
	"1F937-200D-2642-FE0F": {
		Key:        "1F937-200D-2642-FE0F",
		Value:      "🤷‍♂️",
		Descriptor: "man shrugging",
	},
	"1F937-1F3FB-200D-2642-FE0F": {
		Key:        "1F937-1F3FB-200D-2642-FE0F",
		Value:      "🤷🏻‍♂️",
		Descriptor: "man shrugging: light skin tone",
	},
	"1F937-1F3FC-200D-2642-FE0F": {
		Key:        "1F937-1F3FC-200D-2642-FE0F",
		Value:      "🤷🏼‍♂️",
		Descriptor: "man shrugging: medium-light skin tone",
	},
	"1F937-1F3FD-200D-2642-FE0F": {
		Key:        "1F937-1F3FD-200D-2642-FE0F",
		Value:      "🤷🏽‍♂️",
		Descriptor: "man shrugging: medium skin tone",
	},
	"1F937-1F3FE-200D-2642-FE0F": {
		Key:        "1F937-1F3FE-200D-2642-FE0F",
		Value:      "🤷🏾‍♂️",
		Descriptor: "man shrugging: medium-dark skin tone",
	},
	"1F937-1F3FF-200D-2642-FE0F": {
		Key:        "1F937-1F3FF-200D-2642-FE0F",
		Value:      "🤷🏿‍♂️",
		Descriptor: "man shrugging: dark skin tone",
	},
	"1F937-200D-2640-FE0F": {
		Key:        "1F937-200D-2640-FE0F",
		Value:      "🤷‍♀️",
		Descriptor: "woman shrugging",
	},
	"1F937-1F3FB-200D-2640-FE0F": {
		Key:        "1F937-1F3FB-200D-2640-FE0F",
		Value:      "🤷🏻‍♀️",
		Descriptor: "woman shrugging: light skin tone",
	},
	"1F937-1F3FC-200D-2640-FE0F": {
		Key:        "1F937-1F3FC-200D-2640-FE0F",
		Value:      "🤷🏼‍♀️",
		Descriptor: "woman shrugging: medium-light skin tone",
	},
	"1F937-1F3FD-200D-2640-FE0F": {
		Key:        "1F937-1F3FD-200D-2640-FE0F",
		Value:      "🤷🏽‍♀️",
		Descriptor: "woman shrugging: medium skin tone",
	},
	"1F937-1F3FE-200D-2640-FE0F": {
		Key:        "1F937-1F3FE-200D-2640-FE0F",
		Value:      "🤷🏾‍♀️",
		Descriptor: "woman shrugging: medium-dark skin tone",
	},
	"1F937-1F3FF-200D-2640-FE0F": {
		Key:        "1F937-1F3FF-200D-2640-FE0F",
		Value:      "🤷🏿‍♀️",
		Descriptor: "woman shrugging: dark skin tone",
	},
	"1F9D1-200D-2695-FE0F": {
		Key:        "1F9D1-200D-2695-FE0F",
		Value:      "🧑‍⚕️",
		Descriptor: "health worker",
	},
	"1F9D1-1F3FB-200D-2695-FE0F": {
		Key:        "1F9D1-1F3FB-200D-2695-FE0F",
		Value:      "🧑🏻‍⚕️",
		Descriptor: "health worker: light skin tone",
	},
	"1F9D1-1F3FC-200D-2695-FE0F": {
		Key:        "1F9D1-1F3FC-200D-2695-FE0F",
		Value:      "🧑🏼‍⚕️",
		Descriptor: "health worker: medium-light skin tone",
	},
	"1F9D1-1F3FD-200D-2695-FE0F": {
		Key:        "1F9D1-1F3FD-200D-2695-FE0F",
		Value:      "🧑🏽‍⚕️",
		Descriptor: "health worker: medium skin tone",
	},
	"1F9D1-1F3FE-200D-2695-FE0F": {
		Key:        "1F9D1-1F3FE-200D-2695-FE0F",
		Value:      "🧑🏾‍⚕️",
		Descriptor: "health worker: medium-dark skin tone",
	},
	"1F9D1-1F3FF-200D-2695-FE0F": {
		Key:        "1F9D1-1F3FF-200D-2695-FE0F",
		Value:      "🧑🏿‍⚕️",
		Descriptor: "health worker: dark skin tone",
	},
	"1F468-200D-2695-FE0F": {
		Key:        "1F468-200D-2695-FE0F",
		Value:      "👨‍⚕️",
		Descriptor: "man health worker",
	},
	"1F468-1F3FB-200D-2695-FE0F": {
		Key:        "1F468-1F3FB-200D-2695-FE0F",
		Value:      "👨🏻‍⚕️",
		Descriptor: "man health worker: light skin tone",
	},
	"1F468-1F3FC-200D-2695-FE0F": {
		Key:        "1F468-1F3FC-200D-2695-FE0F",
		Value:      "👨🏼‍⚕️",
		Descriptor: "man health worker: medium-light skin tone",
	},
	"1F468-1F3FD-200D-2695-FE0F": {
		Key:        "1F468-1F3FD-200D-2695-FE0F",
		Value:      "👨🏽‍⚕️",
		Descriptor: "man health worker: medium skin tone",
	},
	"1F468-1F3FE-200D-2695-FE0F": {
		Key:        "1F468-1F3FE-200D-2695-FE0F",
		Value:      "👨🏾‍⚕️",
		Descriptor: "man health worker: medium-dark skin tone",
	},
	"1F468-1F3FF-200D-2695-FE0F": {
		Key:        "1F468-1F3FF-200D-2695-FE0F",
		Value:      "👨🏿‍⚕️",
		Descriptor: "man health worker: dark skin tone",
	},
	"1F469-200D-2695-FE0F": {
		Key:        "1F469-200D-2695-FE0F",
		Value:      "👩‍⚕️",
		Descriptor: "woman health worker",
	},
	"1F469-1F3FB-200D-2695-FE0F": {
		Key:        "1F469-1F3FB-200D-2695-FE0F",
		Value:      "👩🏻‍⚕️",
		Descriptor: "woman health worker: light skin tone",
	},
	"1F469-1F3FC-200D-2695-FE0F": {
		Key:        "1F469-1F3FC-200D-2695-FE0F",
		Value:      "👩🏼‍⚕️",
		Descriptor: "woman health worker: medium-light skin tone",
	},
	"1F469-1F3FD-200D-2695-FE0F": {
		Key:        "1F469-1F3FD-200D-2695-FE0F",
		Value:      "👩🏽‍⚕️",
		Descriptor: "woman health worker: medium skin tone",
	},
	"1F469-1F3FE-200D-2695-FE0F": {
		Key:        "1F469-1F3FE-200D-2695-FE0F",
		Value:      "👩🏾‍⚕️",
		Descriptor: "woman health worker: medium-dark skin tone",
	},
	"1F469-1F3FF-200D-2695-FE0F": {
		Key:        "1F469-1F3FF-200D-2695-FE0F",
		Value:      "👩🏿‍⚕️",
		Descriptor: "woman health worker: dark skin tone",
	},
	"1F9D1-200D-1F393": {
		Key:        "1F9D1-200D-1F393",
		Value:      "🧑‍🎓",
		Descriptor: "student",
	},
	"1F9D1-1F3FB-200D-1F393": {
		Key:        "1F9D1-1F3FB-200D-1F393",
		Value:      "🧑🏻‍🎓",
		Descriptor: "student: light skin tone",
	},
	"1F9D1-1F3FC-200D-1F393": {
		Key:        "1F9D1-1F3FC-200D-1F393",
		Value:      "🧑🏼‍🎓",
		Descriptor: "student: medium-light skin tone",
	},
	"1F9D1-1F3FD-200D-1F393": {
		Key:        "1F9D1-1F3FD-200D-1F393",
		Value:      "🧑🏽‍🎓",
		Descriptor: "student: medium skin tone",
	},
	"1F9D1-1F3FE-200D-1F393": {
		Key:        "1F9D1-1F3FE-200D-1F393",
		Value:      "🧑🏾‍🎓",
		Descriptor: "student: medium-dark skin tone",
	},
	"1F9D1-1F3FF-200D-1F393": {
		Key:        "1F9D1-1F3FF-200D-1F393",
		Value:      "🧑🏿‍🎓",
		Descriptor: "student: dark skin tone",
	},
	"1F468-200D-1F393": {
		Key:        "1F468-200D-1F393",
		Value:      "👨‍🎓",
		Descriptor: "man student",
	},
	"1F468-1F3FB-200D-1F393": {
		Key:        "1F468-1F3FB-200D-1F393",
		Value:      "👨🏻‍🎓",
		Descriptor: "man student: light skin tone",
	},
	"1F468-1F3FC-200D-1F393": {
		Key:        "1F468-1F3FC-200D-1F393",
		Value:      "👨🏼‍🎓",
		Descriptor: "man student: medium-light skin tone",
	},
	"1F468-1F3FD-200D-1F393": {
		Key:        "1F468-1F3FD-200D-1F393",
		Value:      "👨🏽‍🎓",
		Descriptor: "man student: medium skin tone",
	},
	"1F468-1F3FE-200D-1F393": {
		Key:        "1F468-1F3FE-200D-1F393",
		Value:      "👨🏾‍🎓",
		Descriptor: "man student: medium-dark skin tone",
	},
	"1F468-1F3FF-200D-1F393": {
		Key:        "1F468-1F3FF-200D-1F393",
		Value:      "👨🏿‍🎓",
		Descriptor: "man student: dark skin tone",
	},
	"1F469-200D-1F393": {
		Key:        "1F469-200D-1F393",
		Value:      "👩‍🎓",
		Descriptor: "woman student",
	},
	"1F469-1F3FB-200D-1F393": {
		Key:        "1F469-1F3FB-200D-1F393",
		Value:      "👩🏻‍🎓",
		Descriptor: "woman student: light skin tone",
	},
	"1F469-1F3FC-200D-1F393": {
		Key:        "1F469-1F3FC-200D-1F393",
		Value:      "👩🏼‍🎓",
		Descriptor: "woman student: medium-light skin tone",
	},
	"1F469-1F3FD-200D-1F393": {
		Key:        "1F469-1F3FD-200D-1F393",
		Value:      "👩🏽‍🎓",
		Descriptor: "woman student: medium skin tone",
	},
	"1F469-1F3FE-200D-1F393": {
		Key:        "1F469-1F3FE-200D-1F393",
		Value:      "👩🏾‍🎓",
		Descriptor: "woman student: medium-dark skin tone",
	},
	"1F469-1F3FF-200D-1F393": {
		Key:        "1F469-1F3FF-200D-1F393",
		Value:      "👩🏿‍🎓",
		Descriptor: "woman student: dark skin tone",
	},
	"1F9D1-200D-1F3EB": {
		Key:        "1F9D1-200D-1F3EB",
		Value:      "🧑‍🏫",
		Descriptor: "teacher",
	},
	"1F9D1-1F3FB-200D-1F3EB": {
		Key:        "1F9D1-1F3FB-200D-1F3EB",
		Value:      "🧑🏻‍🏫",
		Descriptor: "teacher: light skin tone",
	},
	"1F9D1-1F3FC-200D-1F3EB": {
		Key:        "1F9D1-1F3FC-200D-1F3EB",
		Value:      "🧑🏼‍🏫",
		Descriptor: "teacher: medium-light skin tone",
	},
	"1F9D1-1F3FD-200D-1F3EB": {
		Key:        "1F9D1-1F3FD-200D-1F3EB",
		Value:      "🧑🏽‍🏫",
		Descriptor: "teacher: medium skin tone",
	},
	"1F9D1-1F3FE-200D-1F3EB": {
		Key:        "1F9D1-1F3FE-200D-1F3EB",
		Value:      "🧑🏾‍🏫",
		Descriptor: "teacher: medium-dark skin tone",
	},
	"1F9D1-1F3FF-200D-1F3EB": {
		Key:        "1F9D1-1F3FF-200D-1F3EB",
		Value:      "🧑🏿‍🏫",
		Descriptor: "teacher: dark skin tone",
	},
	"1F468-200D-1F3EB": {
		Key:        "1F468-200D-1F3EB",
		Value:      "👨‍🏫",
		Descriptor: "man teacher",
	},
	"1F468-1F3FB-200D-1F3EB": {
		Key:        "1F468-1F3FB-200D-1F3EB",
		Value:      "👨🏻‍🏫",
		Descriptor: "man teacher: light skin tone",
	},
	"1F468-1F3FC-200D-1F3EB": {
		Key:        "1F468-1F3FC-200D-1F3EB",
		Value:      "👨🏼‍🏫",
		Descriptor: "man teacher: medium-light skin tone",
	},
	"1F468-1F3FD-200D-1F3EB": {
		Key:        "1F468-1F3FD-200D-1F3EB",
		Value:      "👨🏽‍🏫",
		Descriptor: "man teacher: medium skin tone",
	},
	"1F468-1F3FE-200D-1F3EB": {
		Key:        "1F468-1F3FE-200D-1F3EB",
		Value:      "👨🏾‍🏫",
		Descriptor: "man teacher: medium-dark skin tone",
	},
	"1F468-1F3FF-200D-1F3EB": {
		Key:        "1F468-1F3FF-200D-1F3EB",
		Value:      "👨🏿‍🏫",
		Descriptor: "man teacher: dark skin tone",
	},
	"1F469-200D-1F3EB": {
		Key:        "1F469-200D-1F3EB",
		Value:      "👩‍🏫",
		Descriptor: "woman teacher",
	},
	"1F469-1F3FB-200D-1F3EB": {
		Key:        "1F469-1F3FB-200D-1F3EB",
		Value:      "👩🏻‍🏫",
		Descriptor: "woman teacher: light skin tone",
	},
	"1F469-1F3FC-200D-1F3EB": {
		Key:        "1F469-1F3FC-200D-1F3EB",
		Value:      "👩🏼‍🏫",
		Descriptor: "woman teacher: medium-light skin tone",
	},
	"1F469-1F3FD-200D-1F3EB": {
		Key:        "1F469-1F3FD-200D-1F3EB",
		Value:      "👩🏽‍🏫",
		Descriptor: "woman teacher: medium skin tone",
	},
	"1F469-1F3FE-200D-1F3EB": {
		Key:        "1F469-1F3FE-200D-1F3EB",
		Value:      "👩🏾‍🏫",
		Descriptor: "woman teacher: medium-dark skin tone",
	},
	"1F469-1F3FF-200D-1F3EB": {
		Key:        "1F469-1F3FF-200D-1F3EB",
		Value:      "👩🏿‍🏫",
		Descriptor: "woman teacher: dark skin tone",
	},
	"1F9D1-200D-2696-FE0F": {
		Key:        "1F9D1-200D-2696-FE0F",
		Value:      "🧑‍⚖️",
		Descriptor: "judge",
	},
	"1F9D1-1F3FB-200D-2696-FE0F": {
		Key:        "1F9D1-1F3FB-200D-2696-FE0F",
		Value:      "🧑🏻‍⚖️",
		Descriptor: "judge: light skin tone",
	},
	"1F9D1-1F3FC-200D-2696-FE0F": {
		Key:        "1F9D1-1F3FC-200D-2696-FE0F",
		Value:      "🧑🏼‍⚖️",
		Descriptor: "judge: medium-light skin tone",
	},
	"1F9D1-1F3FD-200D-2696-FE0F": {
		Key:        "1F9D1-1F3FD-200D-2696-FE0F",
		Value:      "🧑🏽‍⚖️",
		Descriptor: "judge: medium skin tone",
	},
	"1F9D1-1F3FE-200D-2696-FE0F": {
		Key:        "1F9D1-1F3FE-200D-2696-FE0F",
		Value:      "🧑🏾‍⚖️",
		Descriptor: "judge: medium-dark skin tone",
	},
	"1F9D1-1F3FF-200D-2696-FE0F": {
		Key:        "1F9D1-1F3FF-200D-2696-FE0F",
		Value:      "🧑🏿‍⚖️",
		Descriptor: "judge: dark skin tone",
	},
	"1F468-200D-2696-FE0F": {
		Key:        "1F468-200D-2696-FE0F",
		Value:      "👨‍⚖️",
		Descriptor: "man judge",
	},
	"1F468-1F3FB-200D-2696-FE0F": {
		Key:        "1F468-1F3FB-200D-2696-FE0F",
		Value:      "👨🏻‍⚖️",
		Descriptor: "man judge: light skin tone",
	},
	"1F468-1F3FC-200D-2696-FE0F": {
		Key:        "1F468-1F3FC-200D-2696-FE0F",
		Value:      "👨🏼‍⚖️",
		Descriptor: "man judge: medium-light skin tone",
	},
	"1F468-1F3FD-200D-2696-FE0F": {
		Key:        "1F468-1F3FD-200D-2696-FE0F",
		Value:      "👨🏽‍⚖️",
		Descriptor: "man judge: medium skin tone",
	},
	"1F468-1F3FE-200D-2696-FE0F": {
		Key:        "1F468-1F3FE-200D-2696-FE0F",
		Value:      "👨🏾‍⚖️",
		Descriptor: "man judge: medium-dark skin tone",
	},
	"1F468-1F3FF-200D-2696-FE0F": {
		Key:        "1F468-1F3FF-200D-2696-FE0F",
		Value:      "👨🏿‍⚖️",
		Descriptor: "man judge: dark skin tone",
	},
	"1F469-200D-2696-FE0F": {
		Key:        "1F469-200D-2696-FE0F",
		Value:      "👩‍⚖️",
		Descriptor: "woman judge",
	},
	"1F469-1F3FB-200D-2696-FE0F": {
		Key:        "1F469-1F3FB-200D-2696-FE0F",
		Value:      "👩🏻‍⚖️",
		Descriptor: "woman judge: light skin tone",
	},
	"1F469-1F3FC-200D-2696-FE0F": {
		Key:        "1F469-1F3FC-200D-2696-FE0F",
		Value:      "👩🏼‍⚖️",
		Descriptor: "woman judge: medium-light skin tone",
	},
	"1F469-1F3FD-200D-2696-FE0F": {
		Key:        "1F469-1F3FD-200D-2696-FE0F",
		Value:      "👩🏽‍⚖️",
		Descriptor: "woman judge: medium skin tone",
	},
	"1F469-1F3FE-200D-2696-FE0F": {
		Key:        "1F469-1F3FE-200D-2696-FE0F",
		Value:      "👩🏾‍⚖️",
		Descriptor: "woman judge: medium-dark skin tone",
	},
	"1F469-1F3FF-200D-2696-FE0F": {
		Key:        "1F469-1F3FF-200D-2696-FE0F",
		Value:      "👩🏿‍⚖️",
		Descriptor: "woman judge: dark skin tone",
	},
	"1F9D1-200D-1F33E": {
		Key:        "1F9D1-200D-1F33E",
		Value:      "🧑‍🌾",
		Descriptor: "farmer",
	},
	"1F9D1-1F3FB-200D-1F33E": {
		Key:        "1F9D1-1F3FB-200D-1F33E",
		Value:      "🧑🏻‍🌾",
		Descriptor: "farmer: light skin tone",
	},
	"1F9D1-1F3FC-200D-1F33E": {
		Key:        "1F9D1-1F3FC-200D-1F33E",
		Value:      "🧑🏼‍🌾",
		Descriptor: "farmer: medium-light skin tone",
	},
	"1F9D1-1F3FD-200D-1F33E": {
		Key:        "1F9D1-1F3FD-200D-1F33E",
		Value:      "🧑🏽‍🌾",
		Descriptor: "farmer: medium skin tone",
	},
	"1F9D1-1F3FE-200D-1F33E": {
		Key:        "1F9D1-1F3FE-200D-1F33E",
		Value:      "🧑🏾‍🌾",
		Descriptor: "farmer: medium-dark skin tone",
	},
	"1F9D1-1F3FF-200D-1F33E": {
		Key:        "1F9D1-1F3FF-200D-1F33E",
		Value:      "🧑🏿‍🌾",
		Descriptor: "farmer: dark skin tone",
	},
	"1F468-200D-1F33E": {
		Key:        "1F468-200D-1F33E",
		Value:      "👨‍🌾",
		Descriptor: "man farmer",
	},
	"1F468-1F3FB-200D-1F33E": {
		Key:        "1F468-1F3FB-200D-1F33E",
		Value:      "👨🏻‍🌾",
		Descriptor: "man farmer: light skin tone",
	},
	"1F468-1F3FC-200D-1F33E": {
		Key:        "1F468-1F3FC-200D-1F33E",
		Value:      "👨🏼‍🌾",
		Descriptor: "man farmer: medium-light skin tone",
	},
	"1F468-1F3FD-200D-1F33E": {
		Key:        "1F468-1F3FD-200D-1F33E",
		Value:      "👨🏽‍🌾",
		Descriptor: "man farmer: medium skin tone",
	},
	"1F468-1F3FE-200D-1F33E": {
		Key:        "1F468-1F3FE-200D-1F33E",
		Value:      "👨🏾‍🌾",
		Descriptor: "man farmer: medium-dark skin tone",
	},
	"1F468-1F3FF-200D-1F33E": {
		Key:        "1F468-1F3FF-200D-1F33E",
		Value:      "👨🏿‍🌾",
		Descriptor: "man farmer: dark skin tone",
	},
	"1F469-200D-1F33E": {
		Key:        "1F469-200D-1F33E",
		Value:      "👩‍🌾",
		Descriptor: "woman farmer",
	},
	"1F469-1F3FB-200D-1F33E": {
		Key:        "1F469-1F3FB-200D-1F33E",
		Value:      "👩🏻‍🌾",
		Descriptor: "woman farmer: light skin tone",
	},
	"1F469-1F3FC-200D-1F33E": {
		Key:        "1F469-1F3FC-200D-1F33E",
		Value:      "👩🏼‍🌾",
		Descriptor: "woman farmer: medium-light skin tone",
	},
	"1F469-1F3FD-200D-1F33E": {
		Key:        "1F469-1F3FD-200D-1F33E",
		Value:      "👩🏽‍🌾",
		Descriptor: "woman farmer: medium skin tone",
	},
	"1F469-1F3FE-200D-1F33E": {
		Key:        "1F469-1F3FE-200D-1F33E",
		Value:      "👩🏾‍🌾",
		Descriptor: "woman farmer: medium-dark skin tone",
	},
	"1F469-1F3FF-200D-1F33E": {
		Key:        "1F469-1F3FF-200D-1F33E",
		Value:      "👩🏿‍🌾",
		Descriptor: "woman farmer: dark skin tone",
	},
	"1F9D1-200D-1F373": {
		Key:        "1F9D1-200D-1F373",
		Value:      "🧑‍🍳",
		Descriptor: "cook",
	},
	"1F9D1-1F3FB-200D-1F373": {
		Key:        "1F9D1-1F3FB-200D-1F373",
		Value:      "🧑🏻‍🍳",
		Descriptor: "cook: light skin tone",
	},
	"1F9D1-1F3FC-200D-1F373": {
		Key:        "1F9D1-1F3FC-200D-1F373",
		Value:      "🧑🏼‍🍳",
		Descriptor: "cook: medium-light skin tone",
	},
	"1F9D1-1F3FD-200D-1F373": {
		Key:        "1F9D1-1F3FD-200D-1F373",
		Value:      "🧑🏽‍🍳",
		Descriptor: "cook: medium skin tone",
	},
	"1F9D1-1F3FE-200D-1F373": {
		Key:        "1F9D1-1F3FE-200D-1F373",
		Value:      "🧑🏾‍🍳",
		Descriptor: "cook: medium-dark skin tone",
	},
	"1F9D1-1F3FF-200D-1F373": {
		Key:        "1F9D1-1F3FF-200D-1F373",
		Value:      "🧑🏿‍🍳",
		Descriptor: "cook: dark skin tone",
	},
	"1F468-200D-1F373": {
		Key:        "1F468-200D-1F373",
		Value:      "👨‍🍳",
		Descriptor: "man cook",
	},
	"1F468-1F3FB-200D-1F373": {
		Key:        "1F468-1F3FB-200D-1F373",
		Value:      "👨🏻‍🍳",
		Descriptor: "man cook: light skin tone",
	},
	"1F468-1F3FC-200D-1F373": {
		Key:        "1F468-1F3FC-200D-1F373",
		Value:      "👨🏼‍🍳",
		Descriptor: "man cook: medium-light skin tone",
	},
	"1F468-1F3FD-200D-1F373": {
		Key:        "1F468-1F3FD-200D-1F373",
		Value:      "👨🏽‍🍳",
		Descriptor: "man cook: medium skin tone",
	},
	"1F468-1F3FE-200D-1F373": {
		Key:        "1F468-1F3FE-200D-1F373",
		Value:      "👨🏾‍🍳",
		Descriptor: "man cook: medium-dark skin tone",
	},
	"1F468-1F3FF-200D-1F373": {
		Key:        "1F468-1F3FF-200D-1F373",
		Value:      "👨🏿‍🍳",
		Descriptor: "man cook: dark skin tone",
	},
	"1F469-200D-1F373": {
		Key:        "1F469-200D-1F373",
		Value:      "👩‍🍳",
		Descriptor: "woman cook",
	},
	"1F469-1F3FB-200D-1F373": {
		Key:        "1F469-1F3FB-200D-1F373",
		Value:      "👩🏻‍🍳",
		Descriptor: "woman cook: light skin tone",
	},
	"1F469-1F3FC-200D-1F373": {
		Key:        "1F469-1F3FC-200D-1F373",
		Value:      "👩🏼‍🍳",
		Descriptor: "woman cook: medium-light skin tone",
	},
	"1F469-1F3FD-200D-1F373": {
		Key:        "1F469-1F3FD-200D-1F373",
		Value:      "👩🏽‍🍳",
		Descriptor: "woman cook: medium skin tone",
	},
	"1F469-1F3FE-200D-1F373": {
		Key:        "1F469-1F3FE-200D-1F373",
		Value:      "👩🏾‍🍳",
		Descriptor: "woman cook: medium-dark skin tone",
	},
	"1F469-1F3FF-200D-1F373": {
		Key:        "1F469-1F3FF-200D-1F373",
		Value:      "👩🏿‍🍳",
		Descriptor: "woman cook: dark skin tone",
	},
	"1F9D1-200D-1F527": {
		Key:        "1F9D1-200D-1F527",
		Value:      "🧑‍🔧",
		Descriptor: "mechanic",
	},
	"1F9D1-1F3FB-200D-1F527": {
		Key:        "1F9D1-1F3FB-200D-1F527",
		Value:      "🧑🏻‍🔧",
		Descriptor: "mechanic: light skin tone",
	},
	"1F9D1-1F3FC-200D-1F527": {
		Key:        "1F9D1-1F3FC-200D-1F527",
		Value:      "🧑🏼‍🔧",
		Descriptor: "mechanic: medium-light skin tone",
	},
	"1F9D1-1F3FD-200D-1F527": {
		Key:        "1F9D1-1F3FD-200D-1F527",
		Value:      "🧑🏽‍🔧",
		Descriptor: "mechanic: medium skin tone",
	},
	"1F9D1-1F3FE-200D-1F527": {
		Key:        "1F9D1-1F3FE-200D-1F527",
		Value:      "🧑🏾‍🔧",
		Descriptor: "mechanic: medium-dark skin tone",
	},
	"1F9D1-1F3FF-200D-1F527": {
		Key:        "1F9D1-1F3FF-200D-1F527",
		Value:      "🧑🏿‍🔧",
		Descriptor: "mechanic: dark skin tone",
	},
	"1F468-200D-1F527": {
		Key:        "1F468-200D-1F527",
		Value:      "👨‍🔧",
		Descriptor: "man mechanic",
	},
	"1F468-1F3FB-200D-1F527": {
		Key:        "1F468-1F3FB-200D-1F527",
		Value:      "👨🏻‍🔧",
		Descriptor: "man mechanic: light skin tone",
	},
	"1F468-1F3FC-200D-1F527": {
		Key:        "1F468-1F3FC-200D-1F527",
		Value:      "👨🏼‍🔧",
		Descriptor: "man mechanic: medium-light skin tone",
	},
	"1F468-1F3FD-200D-1F527": {
		Key:        "1F468-1F3FD-200D-1F527",
		Value:      "👨🏽‍🔧",
		Descriptor: "man mechanic: medium skin tone",
	},
	"1F468-1F3FE-200D-1F527": {
		Key:        "1F468-1F3FE-200D-1F527",
		Value:      "👨🏾‍🔧",
		Descriptor: "man mechanic: medium-dark skin tone",
	},
	"1F468-1F3FF-200D-1F527": {
		Key:        "1F468-1F3FF-200D-1F527",
		Value:      "👨🏿‍🔧",
		Descriptor: "man mechanic: dark skin tone",
	},
	"1F469-200D-1F527": {
		Key:        "1F469-200D-1F527",
		Value:      "👩‍🔧",
		Descriptor: "woman mechanic",
	},
	"1F469-1F3FB-200D-1F527": {
		Key:        "1F469-1F3FB-200D-1F527",
		Value:      "👩🏻‍🔧",
		Descriptor: "woman mechanic: light skin tone",
	},
	"1F469-1F3FC-200D-1F527": {
		Key:        "1F469-1F3FC-200D-1F527",
		Value:      "👩🏼‍🔧",
		Descriptor: "woman mechanic: medium-light skin tone",
	},
	"1F469-1F3FD-200D-1F527": {
		Key:        "1F469-1F3FD-200D-1F527",
		Value:      "👩🏽‍🔧",
		Descriptor: "woman mechanic: medium skin tone",
	},
	"1F469-1F3FE-200D-1F527": {
		Key:        "1F469-1F3FE-200D-1F527",
		Value:      "👩🏾‍🔧",
		Descriptor: "woman mechanic: medium-dark skin tone",
	},
	"1F469-1F3FF-200D-1F527": {
		Key:        "1F469-1F3FF-200D-1F527",
		Value:      "👩🏿‍🔧",
		Descriptor: "woman mechanic: dark skin tone",
	},
	"1F9D1-200D-1F3ED": {
		Key:        "1F9D1-200D-1F3ED",
		Value:      "🧑‍🏭",
		Descriptor: "factory worker",
	},
	"1F9D1-1F3FB-200D-1F3ED": {
		Key:        "1F9D1-1F3FB-200D-1F3ED",
		Value:      "🧑🏻‍🏭",
		Descriptor: "factory worker: light skin tone",
	},
	"1F9D1-1F3FC-200D-1F3ED": {
		Key:        "1F9D1-1F3FC-200D-1F3ED",
		Value:      "🧑🏼‍🏭",
		Descriptor: "factory worker: medium-light skin tone",
	},
	"1F9D1-1F3FD-200D-1F3ED": {
		Key:        "1F9D1-1F3FD-200D-1F3ED",
		Value:      "🧑🏽‍🏭",
		Descriptor: "factory worker: medium skin tone",
	},
	"1F9D1-1F3FE-200D-1F3ED": {
		Key:        "1F9D1-1F3FE-200D-1F3ED",
		Value:      "🧑🏾‍🏭",
		Descriptor: "factory worker: medium-dark skin tone",
	},
	"1F9D1-1F3FF-200D-1F3ED": {
		Key:        "1F9D1-1F3FF-200D-1F3ED",
		Value:      "🧑🏿‍🏭",
		Descriptor: "factory worker: dark skin tone",
	},
	"1F468-200D-1F3ED": {
		Key:        "1F468-200D-1F3ED",
		Value:      "👨‍🏭",
		Descriptor: "man factory worker",
	},
	"1F468-1F3FB-200D-1F3ED": {
		Key:        "1F468-1F3FB-200D-1F3ED",
		Value:      "👨🏻‍🏭",
		Descriptor: "man factory worker: light skin tone",
	},
	"1F468-1F3FC-200D-1F3ED": {
		Key:        "1F468-1F3FC-200D-1F3ED",
		Value:      "👨🏼‍🏭",
		Descriptor: "man factory worker: medium-light skin tone",
	},
	"1F468-1F3FD-200D-1F3ED": {
		Key:        "1F468-1F3FD-200D-1F3ED",
		Value:      "👨🏽‍🏭",
		Descriptor: "man factory worker: medium skin tone",
	},
	"1F468-1F3FE-200D-1F3ED": {
		Key:        "1F468-1F3FE-200D-1F3ED",
		Value:      "👨🏾‍🏭",
		Descriptor: "man factory worker: medium-dark skin tone",
	},
	"1F468-1F3FF-200D-1F3ED": {
		Key:        "1F468-1F3FF-200D-1F3ED",
		Value:      "👨🏿‍🏭",
		Descriptor: "man factory worker: dark skin tone",
	},
	"1F469-200D-1F3ED": {
		Key:        "1F469-200D-1F3ED",
		Value:      "👩‍🏭",
		Descriptor: "woman factory worker",
	},
	"1F469-1F3FB-200D-1F3ED": {
		Key:        "1F469-1F3FB-200D-1F3ED",
		Value:      "👩🏻‍🏭",
		Descriptor: "woman factory worker: light skin tone",
	},
	"1F469-1F3FC-200D-1F3ED": {
		Key:        "1F469-1F3FC-200D-1F3ED",
		Value:      "👩🏼‍🏭",
		Descriptor: "woman factory worker: medium-light skin tone",
	},
	"1F469-1F3FD-200D-1F3ED": {
		Key:        "1F469-1F3FD-200D-1F3ED",
		Value:      "👩🏽‍🏭",
		Descriptor: "woman factory worker: medium skin tone",
	},
	"1F469-1F3FE-200D-1F3ED": {
		Key:        "1F469-1F3FE-200D-1F3ED",
		Value:      "👩🏾‍🏭",
		Descriptor: "woman factory worker: medium-dark skin tone",
	},
	"1F469-1F3FF-200D-1F3ED": {
		Key:        "1F469-1F3FF-200D-1F3ED",
		Value:      "👩🏿‍🏭",
		Descriptor: "woman factory worker: dark skin tone",
	},
	"1F9D1-200D-1F4BC": {
		Key:        "1F9D1-200D-1F4BC",
		Value:      "🧑‍💼",
		Descriptor: "office worker",
	},
	"1F9D1-1F3FB-200D-1F4BC": {
		Key:        "1F9D1-1F3FB-200D-1F4BC",
		Value:      "🧑🏻‍💼",
		Descriptor: "office worker: light skin tone",
	},
	"1F9D1-1F3FC-200D-1F4BC": {
		Key:        "1F9D1-1F3FC-200D-1F4BC",
		Value:      "🧑🏼‍💼",
		Descriptor: "office worker: medium-light skin tone",
	},
	"1F9D1-1F3FD-200D-1F4BC": {
		Key:        "1F9D1-1F3FD-200D-1F4BC",
		Value:      "🧑🏽‍💼",
		Descriptor: "office worker: medium skin tone",
	},
	"1F9D1-1F3FE-200D-1F4BC": {
		Key:        "1F9D1-1F3FE-200D-1F4BC",
		Value:      "🧑🏾‍💼",
		Descriptor: "office worker: medium-dark skin tone",
	},
	"1F9D1-1F3FF-200D-1F4BC": {
		Key:        "1F9D1-1F3FF-200D-1F4BC",
		Value:      "🧑🏿‍💼",
		Descriptor: "office worker: dark skin tone",
	},
	"1F468-200D-1F4BC": {
		Key:        "1F468-200D-1F4BC",
		Value:      "👨‍💼",
		Descriptor: "man office worker",
	},
	"1F468-1F3FB-200D-1F4BC": {
		Key:        "1F468-1F3FB-200D-1F4BC",
		Value:      "👨🏻‍💼",
		Descriptor: "man office worker: light skin tone",
	},
	"1F468-1F3FC-200D-1F4BC": {
		Key:        "1F468-1F3FC-200D-1F4BC",
		Value:      "👨🏼‍💼",
		Descriptor: "man office worker: medium-light skin tone",
	},
	"1F468-1F3FD-200D-1F4BC": {
		Key:        "1F468-1F3FD-200D-1F4BC",
		Value:      "👨🏽‍💼",
		Descriptor: "man office worker: medium skin tone",
	},
	"1F468-1F3FE-200D-1F4BC": {
		Key:        "1F468-1F3FE-200D-1F4BC",
		Value:      "👨🏾‍💼",
		Descriptor: "man office worker: medium-dark skin tone",
	},
	"1F468-1F3FF-200D-1F4BC": {
		Key:        "1F468-1F3FF-200D-1F4BC",
		Value:      "👨🏿‍💼",
		Descriptor: "man office worker: dark skin tone",
	},
	"1F469-200D-1F4BC": {
		Key:        "1F469-200D-1F4BC",
		Value:      "👩‍💼",
		Descriptor: "woman office worker",
	},
	"1F469-1F3FB-200D-1F4BC": {
		Key:        "1F469-1F3FB-200D-1F4BC",
		Value:      "👩🏻‍💼",
		Descriptor: "woman office worker: light skin tone",
	},
	"1F469-1F3FC-200D-1F4BC": {
		Key:        "1F469-1F3FC-200D-1F4BC",
		Value:      "👩🏼‍💼",
		Descriptor: "woman office worker: medium-light skin tone",
	},
	"1F469-1F3FD-200D-1F4BC": {
		Key:        "1F469-1F3FD-200D-1F4BC",
		Value:      "👩🏽‍💼",
		Descriptor: "woman office worker: medium skin tone",
	},
	"1F469-1F3FE-200D-1F4BC": {
		Key:        "1F469-1F3FE-200D-1F4BC",
		Value:      "👩🏾‍💼",
		Descriptor: "woman office worker: medium-dark skin tone",
	},
	"1F469-1F3FF-200D-1F4BC": {
		Key:        "1F469-1F3FF-200D-1F4BC",
		Value:      "👩🏿‍💼",
		Descriptor: "woman office worker: dark skin tone",
	},
	"1F9D1-200D-1F52C": {
		Key:        "1F9D1-200D-1F52C",
		Value:      "🧑‍🔬",
		Descriptor: "scientist",
	},
	"1F9D1-1F3FB-200D-1F52C": {
		Key:        "1F9D1-1F3FB-200D-1F52C",
		Value:      "🧑🏻‍🔬",
		Descriptor: "scientist: light skin tone",
	},
	"1F9D1-1F3FC-200D-1F52C": {
		Key:        "1F9D1-1F3FC-200D-1F52C",
		Value:      "🧑🏼‍🔬",
		Descriptor: "scientist: medium-light skin tone",
	},
	"1F9D1-1F3FD-200D-1F52C": {
		Key:        "1F9D1-1F3FD-200D-1F52C",
		Value:      "🧑🏽‍🔬",
		Descriptor: "scientist: medium skin tone",
	},
	"1F9D1-1F3FE-200D-1F52C": {
		Key:        "1F9D1-1F3FE-200D-1F52C",
		Value:      "🧑🏾‍🔬",
		Descriptor: "scientist: medium-dark skin tone",
	},
	"1F9D1-1F3FF-200D-1F52C": {
		Key:        "1F9D1-1F3FF-200D-1F52C",
		Value:      "🧑🏿‍🔬",
		Descriptor: "scientist: dark skin tone",
	},
	"1F468-200D-1F52C": {
		Key:        "1F468-200D-1F52C",
		Value:      "👨‍🔬",
		Descriptor: "man scientist",
	},
	"1F468-1F3FB-200D-1F52C": {
		Key:        "1F468-1F3FB-200D-1F52C",
		Value:      "👨🏻‍🔬",
		Descriptor: "man scientist: light skin tone",
	},
	"1F468-1F3FC-200D-1F52C": {
		Key:        "1F468-1F3FC-200D-1F52C",
		Value:      "👨🏼‍🔬",
		Descriptor: "man scientist: medium-light skin tone",
	},
	"1F468-1F3FD-200D-1F52C": {
		Key:        "1F468-1F3FD-200D-1F52C",
		Value:      "👨🏽‍🔬",
		Descriptor: "man scientist: medium skin tone",
	},
	"1F468-1F3FE-200D-1F52C": {
		Key:        "1F468-1F3FE-200D-1F52C",
		Value:      "👨🏾‍🔬",
		Descriptor: "man scientist: medium-dark skin tone",
	},
	"1F468-1F3FF-200D-1F52C": {
		Key:        "1F468-1F3FF-200D-1F52C",
		Value:      "👨🏿‍🔬",
		Descriptor: "man scientist: dark skin tone",
	},
	"1F469-200D-1F52C": {
		Key:        "1F469-200D-1F52C",
		Value:      "👩‍🔬",
		Descriptor: "woman scientist",
	},
	"1F469-1F3FB-200D-1F52C": {
		Key:        "1F469-1F3FB-200D-1F52C",
		Value:      "👩🏻‍🔬",
		Descriptor: "woman scientist: light skin tone",
	},
	"1F469-1F3FC-200D-1F52C": {
		Key:        "1F469-1F3FC-200D-1F52C",
		Value:      "👩🏼‍🔬",
		Descriptor: "woman scientist: medium-light skin tone",
	},
	"1F469-1F3FD-200D-1F52C": {
		Key:        "1F469-1F3FD-200D-1F52C",
		Value:      "👩🏽‍🔬",
		Descriptor: "woman scientist: medium skin tone",
	},
	"1F469-1F3FE-200D-1F52C": {
		Key:        "1F469-1F3FE-200D-1F52C",
		Value:      "👩🏾‍🔬",
		Descriptor: "woman scientist: medium-dark skin tone",
	},
	"1F469-1F3FF-200D-1F52C": {
		Key:        "1F469-1F3FF-200D-1F52C",
		Value:      "👩🏿‍🔬",
		Descriptor: "woman scientist: dark skin tone",
	},
	"1F9D1-200D-1F4BB": {
		Key:        "1F9D1-200D-1F4BB",
		Value:      "🧑‍💻",
		Descriptor: "technologist",
	},
	"1F9D1-1F3FB-200D-1F4BB": {
		Key:        "1F9D1-1F3FB-200D-1F4BB",
		Value:      "🧑🏻‍💻",
		Descriptor: "technologist: light skin tone",
	},
	"1F9D1-1F3FC-200D-1F4BB": {
		Key:        "1F9D1-1F3FC-200D-1F4BB",
		Value:      "🧑🏼‍💻",
		Descriptor: "technologist: medium-light skin tone",
	},
	"1F9D1-1F3FD-200D-1F4BB": {
		Key:        "1F9D1-1F3FD-200D-1F4BB",
		Value:      "🧑🏽‍💻",
		Descriptor: "technologist: medium skin tone",
	},
	"1F9D1-1F3FE-200D-1F4BB": {
		Key:        "1F9D1-1F3FE-200D-1F4BB",
		Value:      "🧑🏾‍💻",
		Descriptor: "technologist: medium-dark skin tone",
	},
	"1F9D1-1F3FF-200D-1F4BB": {
		Key:        "1F9D1-1F3FF-200D-1F4BB",
		Value:      "🧑🏿‍💻",
		Descriptor: "technologist: dark skin tone",
	},
	"1F468-200D-1F4BB": {
		Key:        "1F468-200D-1F4BB",
		Value:      "👨‍💻",
		Descriptor: "man technologist",
	},
	"1F468-1F3FB-200D-1F4BB": {
		Key:        "1F468-1F3FB-200D-1F4BB",
		Value:      "👨🏻‍💻",
		Descriptor: "man technologist: light skin tone",
	},
	"1F468-1F3FC-200D-1F4BB": {
		Key:        "1F468-1F3FC-200D-1F4BB",
		Value:      "👨🏼‍💻",
		Descriptor: "man technologist: medium-light skin tone",
	},
	"1F468-1F3FD-200D-1F4BB": {
		Key:        "1F468-1F3FD-200D-1F4BB",
		Value:      "👨🏽‍💻",
		Descriptor: "man technologist: medium skin tone",
	},
	"1F468-1F3FE-200D-1F4BB": {
		Key:        "1F468-1F3FE-200D-1F4BB",
		Value:      "👨🏾‍💻",
		Descriptor: "man technologist: medium-dark skin tone",
	},
	"1F468-1F3FF-200D-1F4BB": {
		Key:        "1F468-1F3FF-200D-1F4BB",
		Value:      "👨🏿‍💻",
		Descriptor: "man technologist: dark skin tone",
	},
	"1F469-200D-1F4BB": {
		Key:        "1F469-200D-1F4BB",
		Value:      "👩‍💻",
		Descriptor: "woman technologist",
	},
	"1F469-1F3FB-200D-1F4BB": {
		Key:        "1F469-1F3FB-200D-1F4BB",
		Value:      "👩🏻‍💻",
		Descriptor: "woman technologist: light skin tone",
	},
	"1F469-1F3FC-200D-1F4BB": {
		Key:        "1F469-1F3FC-200D-1F4BB",
		Value:      "👩🏼‍💻",
		Descriptor: "woman technologist: medium-light skin tone",
	},
	"1F469-1F3FD-200D-1F4BB": {
		Key:        "1F469-1F3FD-200D-1F4BB",
		Value:      "👩🏽‍💻",
		Descriptor: "woman technologist: medium skin tone",
	},
	"1F469-1F3FE-200D-1F4BB": {
		Key:        "1F469-1F3FE-200D-1F4BB",
		Value:      "👩🏾‍💻",
		Descriptor: "woman technologist: medium-dark skin tone",
	},
	"1F469-1F3FF-200D-1F4BB": {
		Key:        "1F469-1F3FF-200D-1F4BB",
		Value:      "👩🏿‍💻",
		Descriptor: "woman technologist: dark skin tone",
	},
	"1F9D1-200D-1F3A4": {
		Key:        "1F9D1-200D-1F3A4",
		Value:      "🧑‍🎤",
		Descriptor: "singer",
	},
	"1F9D1-1F3FB-200D-1F3A4": {
		Key:        "1F9D1-1F3FB-200D-1F3A4",
		Value:      "🧑🏻‍🎤",
		Descriptor: "singer: light skin tone",
	},
	"1F9D1-1F3FC-200D-1F3A4": {
		Key:        "1F9D1-1F3FC-200D-1F3A4",
		Value:      "🧑🏼‍🎤",
		Descriptor: "singer: medium-light skin tone",
	},
	"1F9D1-1F3FD-200D-1F3A4": {
		Key:        "1F9D1-1F3FD-200D-1F3A4",
		Value:      "🧑🏽‍🎤",
		Descriptor: "singer: medium skin tone",
	},
	"1F9D1-1F3FE-200D-1F3A4": {
		Key:        "1F9D1-1F3FE-200D-1F3A4",
		Value:      "🧑🏾‍🎤",
		Descriptor: "singer: medium-dark skin tone",
	},
	"1F9D1-1F3FF-200D-1F3A4": {
		Key:        "1F9D1-1F3FF-200D-1F3A4",
		Value:      "🧑🏿‍🎤",
		Descriptor: "singer: dark skin tone",
	},
	"1F468-200D-1F3A4": {
		Key:        "1F468-200D-1F3A4",
		Value:      "👨‍🎤",
		Descriptor: "man singer",
	},
	"1F468-1F3FB-200D-1F3A4": {
		Key:        "1F468-1F3FB-200D-1F3A4",
		Value:      "👨🏻‍🎤",
		Descriptor: "man singer: light skin tone",
	},
	"1F468-1F3FC-200D-1F3A4": {
		Key:        "1F468-1F3FC-200D-1F3A4",
		Value:      "👨🏼‍🎤",
		Descriptor: "man singer: medium-light skin tone",
	},
	"1F468-1F3FD-200D-1F3A4": {
		Key:        "1F468-1F3FD-200D-1F3A4",
		Value:      "👨🏽‍🎤",
		Descriptor: "man singer: medium skin tone",
	},
	"1F468-1F3FE-200D-1F3A4": {
		Key:        "1F468-1F3FE-200D-1F3A4",
		Value:      "👨🏾‍🎤",
		Descriptor: "man singer: medium-dark skin tone",
	},
	"1F468-1F3FF-200D-1F3A4": {
		Key:        "1F468-1F3FF-200D-1F3A4",
		Value:      "👨🏿‍🎤",
		Descriptor: "man singer: dark skin tone",
	},
	"1F469-200D-1F3A4": {
		Key:        "1F469-200D-1F3A4",
		Value:      "👩‍🎤",
		Descriptor: "woman singer",
	},
	"1F469-1F3FB-200D-1F3A4": {
		Key:        "1F469-1F3FB-200D-1F3A4",
		Value:      "👩🏻‍🎤",
		Descriptor: "woman singer: light skin tone",
	},
	"1F469-1F3FC-200D-1F3A4": {
		Key:        "1F469-1F3FC-200D-1F3A4",
		Value:      "👩🏼‍🎤",
		Descriptor: "woman singer: medium-light skin tone",
	},
	"1F469-1F3FD-200D-1F3A4": {
		Key:        "1F469-1F3FD-200D-1F3A4",
		Value:      "👩🏽‍🎤",
		Descriptor: "woman singer: medium skin tone",
	},
	"1F469-1F3FE-200D-1F3A4": {
		Key:        "1F469-1F3FE-200D-1F3A4",
		Value:      "👩🏾‍🎤",
		Descriptor: "woman singer: medium-dark skin tone",
	},
	"1F469-1F3FF-200D-1F3A4": {
		Key:        "1F469-1F3FF-200D-1F3A4",
		Value:      "👩🏿‍🎤",
		Descriptor: "woman singer: dark skin tone",
	},
	"1F9D1-200D-1F3A8": {
		Key:        "1F9D1-200D-1F3A8",
		Value:      "🧑‍🎨",
		Descriptor: "artist",
	},
	"1F9D1-1F3FB-200D-1F3A8": {
		Key:        "1F9D1-1F3FB-200D-1F3A8",
		Value:      "🧑🏻‍🎨",
		Descriptor: "artist: light skin tone",
	},
	"1F9D1-1F3FC-200D-1F3A8": {
		Key:        "1F9D1-1F3FC-200D-1F3A8",
		Value:      "🧑🏼‍🎨",
		Descriptor: "artist: medium-light skin tone",
	},
	"1F9D1-1F3FD-200D-1F3A8": {
		Key:        "1F9D1-1F3FD-200D-1F3A8",
		Value:      "🧑🏽‍🎨",
		Descriptor: "artist: medium skin tone",
	},
	"1F9D1-1F3FE-200D-1F3A8": {
		Key:        "1F9D1-1F3FE-200D-1F3A8",
		Value:      "🧑🏾‍🎨",
		Descriptor: "artist: medium-dark skin tone",
	},
	"1F9D1-1F3FF-200D-1F3A8": {
		Key:        "1F9D1-1F3FF-200D-1F3A8",
		Value:      "🧑🏿‍🎨",
		Descriptor: "artist: dark skin tone",
	},
	"1F468-200D-1F3A8": {
		Key:        "1F468-200D-1F3A8",
		Value:      "👨‍🎨",
		Descriptor: "man artist",
	},
	"1F468-1F3FB-200D-1F3A8": {
		Key:        "1F468-1F3FB-200D-1F3A8",
		Value:      "👨🏻‍🎨",
		Descriptor: "man artist: light skin tone",
	},
	"1F468-1F3FC-200D-1F3A8": {
		Key:        "1F468-1F3FC-200D-1F3A8",
		Value:      "👨🏼‍🎨",
		Descriptor: "man artist: medium-light skin tone",
	},
	"1F468-1F3FD-200D-1F3A8": {
		Key:        "1F468-1F3FD-200D-1F3A8",
		Value:      "👨🏽‍🎨",
		Descriptor: "man artist: medium skin tone",
	},
	"1F468-1F3FE-200D-1F3A8": {
		Key:        "1F468-1F3FE-200D-1F3A8",
		Value:      "👨🏾‍🎨",
		Descriptor: "man artist: medium-dark skin tone",
	},
	"1F468-1F3FF-200D-1F3A8": {
		Key:        "1F468-1F3FF-200D-1F3A8",
		Value:      "👨🏿‍🎨",
		Descriptor: "man artist: dark skin tone",
	},
	"1F469-200D-1F3A8": {
		Key:        "1F469-200D-1F3A8",
		Value:      "👩‍🎨",
		Descriptor: "woman artist",
	},
	"1F469-1F3FB-200D-1F3A8": {
		Key:        "1F469-1F3FB-200D-1F3A8",
		Value:      "👩🏻‍🎨",
		Descriptor: "woman artist: light skin tone",
	},
	"1F469-1F3FC-200D-1F3A8": {
		Key:        "1F469-1F3FC-200D-1F3A8",
		Value:      "👩🏼‍🎨",
		Descriptor: "woman artist: medium-light skin tone",
	},
	"1F469-1F3FD-200D-1F3A8": {
		Key:        "1F469-1F3FD-200D-1F3A8",
		Value:      "👩🏽‍🎨",
		Descriptor: "woman artist: medium skin tone",
	},
	"1F469-1F3FE-200D-1F3A8": {
		Key:        "1F469-1F3FE-200D-1F3A8",
		Value:      "👩🏾‍🎨",
		Descriptor: "woman artist: medium-dark skin tone",
	},
	"1F469-1F3FF-200D-1F3A8": {
		Key:        "1F469-1F3FF-200D-1F3A8",
		Value:      "👩🏿‍🎨",
		Descriptor: "woman artist: dark skin tone",
	},
	"1F9D1-200D-2708-FE0F": {
		Key:        "1F9D1-200D-2708-FE0F",
		Value:      "🧑‍✈️",
		Descriptor: "pilot",
	},
	"1F9D1-1F3FB-200D-2708-FE0F": {
		Key:        "1F9D1-1F3FB-200D-2708-FE0F",
		Value:      "🧑🏻‍✈️",
		Descriptor: "pilot: light skin tone",
	},
	"1F9D1-1F3FC-200D-2708-FE0F": {
		Key:        "1F9D1-1F3FC-200D-2708-FE0F",
		Value:      "🧑🏼‍✈️",
		Descriptor: "pilot: medium-light skin tone",
	},
	"1F9D1-1F3FD-200D-2708-FE0F": {
		Key:        "1F9D1-1F3FD-200D-2708-FE0F",
		Value:      "🧑🏽‍✈️",
		Descriptor: "pilot: medium skin tone",
	},
	"1F9D1-1F3FE-200D-2708-FE0F": {
		Key:        "1F9D1-1F3FE-200D-2708-FE0F",
		Value:      "🧑🏾‍✈️",
		Descriptor: "pilot: medium-dark skin tone",
	},
	"1F9D1-1F3FF-200D-2708-FE0F": {
		Key:        "1F9D1-1F3FF-200D-2708-FE0F",
		Value:      "🧑🏿‍✈️",
		Descriptor: "pilot: dark skin tone",
	},
	"1F468-200D-2708-FE0F": {
		Key:        "1F468-200D-2708-FE0F",
		Value:      "👨‍✈️",
		Descriptor: "man pilot",
	},
	"1F468-1F3FB-200D-2708-FE0F": {
		Key:        "1F468-1F3FB-200D-2708-FE0F",
		Value:      "👨🏻‍✈️",
		Descriptor: "man pilot: light skin tone",
	},
	"1F468-1F3FC-200D-2708-FE0F": {
		Key:        "1F468-1F3FC-200D-2708-FE0F",
		Value:      "👨🏼‍✈️",
		Descriptor: "man pilot: medium-light skin tone",
	},
	"1F468-1F3FD-200D-2708-FE0F": {
		Key:        "1F468-1F3FD-200D-2708-FE0F",
		Value:      "👨🏽‍✈️",
		Descriptor: "man pilot: medium skin tone",
	},
	"1F468-1F3FE-200D-2708-FE0F": {
		Key:        "1F468-1F3FE-200D-2708-FE0F",
		Value:      "👨🏾‍✈️",
		Descriptor: "man pilot: medium-dark skin tone",
	},
	"1F468-1F3FF-200D-2708-FE0F": {
		Key:        "1F468-1F3FF-200D-2708-FE0F",
		Value:      "👨🏿‍✈️",
		Descriptor: "man pilot: dark skin tone",
	},
	"1F469-200D-2708-FE0F": {
		Key:        "1F469-200D-2708-FE0F",
		Value:      "👩‍✈️",
		Descriptor: "woman pilot",
	},
	"1F469-1F3FB-200D-2708-FE0F": {
		Key:        "1F469-1F3FB-200D-2708-FE0F",
		Value:      "👩🏻‍✈️",
		Descriptor: "woman pilot: light skin tone",
	},
	"1F469-1F3FC-200D-2708-FE0F": {
		Key:        "1F469-1F3FC-200D-2708-FE0F",
		Value:      "👩🏼‍✈️",
		Descriptor: "woman pilot: medium-light skin tone",
	},
	"1F469-1F3FD-200D-2708-FE0F": {
		Key:        "1F469-1F3FD-200D-2708-FE0F",
		Value:      "👩🏽‍✈️",
		Descriptor: "woman pilot: medium skin tone",
	},
	"1F469-1F3FE-200D-2708-FE0F": {
		Key:        "1F469-1F3FE-200D-2708-FE0F",
		Value:      "👩🏾‍✈️",
		Descriptor: "woman pilot: medium-dark skin tone",
	},
	"1F469-1F3FF-200D-2708-FE0F": {
		Key:        "1F469-1F3FF-200D-2708-FE0F",
		Value:      "👩🏿‍✈️",
		Descriptor: "woman pilot: dark skin tone",
	},
	"1F9D1-200D-1F680": {
		Key:        "1F9D1-200D-1F680",
		Value:      "🧑‍🚀",
		Descriptor: "astronaut",
	},
	"1F9D1-1F3FB-200D-1F680": {
		Key:        "1F9D1-1F3FB-200D-1F680",
		Value:      "🧑🏻‍🚀",
		Descriptor: "astronaut: light skin tone",
	},
	"1F9D1-1F3FC-200D-1F680": {
		Key:        "1F9D1-1F3FC-200D-1F680",
		Value:      "🧑🏼‍🚀",
		Descriptor: "astronaut: medium-light skin tone",
	},
	"1F9D1-1F3FD-200D-1F680": {
		Key:        "1F9D1-1F3FD-200D-1F680",
		Value:      "🧑🏽‍🚀",
		Descriptor: "astronaut: medium skin tone",
	},
	"1F9D1-1F3FE-200D-1F680": {
		Key:        "1F9D1-1F3FE-200D-1F680",
		Value:      "🧑🏾‍🚀",
		Descriptor: "astronaut: medium-dark skin tone",
	},
	"1F9D1-1F3FF-200D-1F680": {
		Key:        "1F9D1-1F3FF-200D-1F680",
		Value:      "🧑🏿‍🚀",
		Descriptor: "astronaut: dark skin tone",
	},
	"1F468-200D-1F680": {
		Key:        "1F468-200D-1F680",
		Value:      "👨‍🚀",
		Descriptor: "man astronaut",
	},
	"1F468-1F3FB-200D-1F680": {
		Key:        "1F468-1F3FB-200D-1F680",
		Value:      "👨🏻‍🚀",
		Descriptor: "man astronaut: light skin tone",
	},
	"1F468-1F3FC-200D-1F680": {
		Key:        "1F468-1F3FC-200D-1F680",
		Value:      "👨🏼‍🚀",
		Descriptor: "man astronaut: medium-light skin tone",
	},
	"1F468-1F3FD-200D-1F680": {
		Key:        "1F468-1F3FD-200D-1F680",
		Value:      "👨🏽‍🚀",
		Descriptor: "man astronaut: medium skin tone",
	},
	"1F468-1F3FE-200D-1F680": {
		Key:        "1F468-1F3FE-200D-1F680",
		Value:      "👨🏾‍🚀",
		Descriptor: "man astronaut: medium-dark skin tone",
	},
	"1F468-1F3FF-200D-1F680": {
		Key:        "1F468-1F3FF-200D-1F680",
		Value:      "👨🏿‍🚀",
		Descriptor: "man astronaut: dark skin tone",
	},
	"1F469-200D-1F680": {
		Key:        "1F469-200D-1F680",
		Value:      "👩‍🚀",
		Descriptor: "woman astronaut",
	},
	"1F469-1F3FB-200D-1F680": {
		Key:        "1F469-1F3FB-200D-1F680",
		Value:      "👩🏻‍🚀",
		Descriptor: "woman astronaut: light skin tone",
	},
	"1F469-1F3FC-200D-1F680": {
		Key:        "1F469-1F3FC-200D-1F680",
		Value:      "👩🏼‍🚀",
		Descriptor: "woman astronaut: medium-light skin tone",
	},
	"1F469-1F3FD-200D-1F680": {
		Key:        "1F469-1F3FD-200D-1F680",
		Value:      "👩🏽‍🚀",
		Descriptor: "woman astronaut: medium skin tone",
	},
	"1F469-1F3FE-200D-1F680": {
		Key:        "1F469-1F3FE-200D-1F680",
		Value:      "👩🏾‍🚀",
		Descriptor: "woman astronaut: medium-dark skin tone",
	},
	"1F469-1F3FF-200D-1F680": {
		Key:        "1F469-1F3FF-200D-1F680",
		Value:      "👩🏿‍🚀",
		Descriptor: "woman astronaut: dark skin tone",
	},
	"1F9D1-200D-1F692": {
		Key:        "1F9D1-200D-1F692",
		Value:      "🧑‍🚒",
		Descriptor: "firefighter",
	},
	"1F9D1-1F3FB-200D-1F692": {
		Key:        "1F9D1-1F3FB-200D-1F692",
		Value:      "🧑🏻‍🚒",
		Descriptor: "firefighter: light skin tone",
	},
	"1F9D1-1F3FC-200D-1F692": {
		Key:        "1F9D1-1F3FC-200D-1F692",
		Value:      "🧑🏼‍🚒",
		Descriptor: "firefighter: medium-light skin tone",
	},
	"1F9D1-1F3FD-200D-1F692": {
		Key:        "1F9D1-1F3FD-200D-1F692",
		Value:      "🧑🏽‍🚒",
		Descriptor: "firefighter: medium skin tone",
	},
	"1F9D1-1F3FE-200D-1F692": {
		Key:        "1F9D1-1F3FE-200D-1F692",
		Value:      "🧑🏾‍🚒",
		Descriptor: "firefighter: medium-dark skin tone",
	},
	"1F9D1-1F3FF-200D-1F692": {
		Key:        "1F9D1-1F3FF-200D-1F692",
		Value:      "🧑🏿‍🚒",
		Descriptor: "firefighter: dark skin tone",
	},
	"1F468-200D-1F692": {
		Key:        "1F468-200D-1F692",
		Value:      "👨‍🚒",
		Descriptor: "man firefighter",
	},
	"1F468-1F3FB-200D-1F692": {
		Key:        "1F468-1F3FB-200D-1F692",
		Value:      "👨🏻‍🚒",
		Descriptor: "man firefighter: light skin tone",
	},
	"1F468-1F3FC-200D-1F692": {
		Key:        "1F468-1F3FC-200D-1F692",
		Value:      "👨🏼‍🚒",
		Descriptor: "man firefighter: medium-light skin tone",
	},
	"1F468-1F3FD-200D-1F692": {
		Key:        "1F468-1F3FD-200D-1F692",
		Value:      "👨🏽‍🚒",
		Descriptor: "man firefighter: medium skin tone",
	},
	"1F468-1F3FE-200D-1F692": {
		Key:        "1F468-1F3FE-200D-1F692",
		Value:      "👨🏾‍🚒",
		Descriptor: "man firefighter: medium-dark skin tone",
	},
	"1F468-1F3FF-200D-1F692": {
		Key:        "1F468-1F3FF-200D-1F692",
		Value:      "👨🏿‍🚒",
		Descriptor: "man firefighter: dark skin tone",
	},
	"1F469-200D-1F692": {
		Key:        "1F469-200D-1F692",
		Value:      "👩‍🚒",
		Descriptor: "woman firefighter",
	},
	"1F469-1F3FB-200D-1F692": {
		Key:        "1F469-1F3FB-200D-1F692",
		Value:      "👩🏻‍🚒",
		Descriptor: "woman firefighter: light skin tone",
	},
	"1F469-1F3FC-200D-1F692": {
		Key:        "1F469-1F3FC-200D-1F692",
		Value:      "👩🏼‍🚒",
		Descriptor: "woman firefighter: medium-light skin tone",
	},
	"1F469-1F3FD-200D-1F692": {
		Key:        "1F469-1F3FD-200D-1F692",
		Value:      "👩🏽‍🚒",
		Descriptor: "woman firefighter: medium skin tone",
	},
	"1F469-1F3FE-200D-1F692": {
		Key:        "1F469-1F3FE-200D-1F692",
		Value:      "👩🏾‍🚒",
		Descriptor: "woman firefighter: medium-dark skin tone",
	},
	"1F469-1F3FF-200D-1F692": {
		Key:        "1F469-1F3FF-200D-1F692",
		Value:      "👩🏿‍🚒",
		Descriptor: "woman firefighter: dark skin tone",
	},
	"1F46E": {
		Key:        "1F46E",
		Value:      "👮",
		Descriptor: "police officer",
	},
	"1F46E-1F3FB": {
		Key:        "1F46E-1F3FB",
		Value:      "👮🏻",
		Descriptor: "police officer: light skin tone",
	},
	"1F46E-1F3FC": {
		Key:        "1F46E-1F3FC",
		Value:      "👮🏼",
		Descriptor: "police officer: medium-light skin tone",
	},
	"1F46E-1F3FD": {
		Key:        "1F46E-1F3FD",
		Value:      "👮🏽",
		Descriptor: "police officer: medium skin tone",
	},
	"1F46E-1F3FE": {
		Key:        "1F46E-1F3FE",
		Value:      "👮🏾",
		Descriptor: "police officer: medium-dark skin tone",
	},
	"1F46E-1F3FF": {
		Key:        "1F46E-1F3FF",
		Value:      "👮🏿",
		Descriptor: "police officer: dark skin tone",
	},
	"1F46E-200D-2642-FE0F": {
		Key:        "1F46E-200D-2642-FE0F",
		Value:      "👮‍♂️",
		Descriptor: "man police officer",
	},
	"1F46E-1F3FB-200D-2642-FE0F": {
		Key:        "1F46E-1F3FB-200D-2642-FE0F",
		Value:      "👮🏻‍♂️",
		Descriptor: "man police officer: light skin tone",
	},
	"1F46E-1F3FC-200D-2642-FE0F": {
		Key:        "1F46E-1F3FC-200D-2642-FE0F",
		Value:      "👮🏼‍♂️",
		Descriptor: "man police officer: medium-light skin tone",
	},
	"1F46E-1F3FD-200D-2642-FE0F": {
		Key:        "1F46E-1F3FD-200D-2642-FE0F",
		Value:      "👮🏽‍♂️",
		Descriptor: "man police officer: medium skin tone",
	},
	"1F46E-1F3FE-200D-2642-FE0F": {
		Key:        "1F46E-1F3FE-200D-2642-FE0F",
		Value:      "👮🏾‍♂️",
		Descriptor: "man police officer: medium-dark skin tone",
	},
	"1F46E-1F3FF-200D-2642-FE0F": {
		Key:        "1F46E-1F3FF-200D-2642-FE0F",
		Value:      "👮🏿‍♂️",
		Descriptor: "man police officer: dark skin tone",
	},
	"1F46E-200D-2640-FE0F": {
		Key:        "1F46E-200D-2640-FE0F",
		Value:      "👮‍♀️",
		Descriptor: "woman police officer",
	},
	"1F46E-1F3FB-200D-2640-FE0F": {
		Key:        "1F46E-1F3FB-200D-2640-FE0F",
		Value:      "👮🏻‍♀️",
		Descriptor: "woman police officer: light skin tone",
	},
	"1F46E-1F3FC-200D-2640-FE0F": {
		Key:        "1F46E-1F3FC-200D-2640-FE0F",
		Value:      "👮🏼‍♀️",
		Descriptor: "woman police officer: medium-light skin tone",
	},
	"1F46E-1F3FD-200D-2640-FE0F": {
		Key:        "1F46E-1F3FD-200D-2640-FE0F",
		Value:      "👮🏽‍♀️",
		Descriptor: "woman police officer: medium skin tone",
	},
	"1F46E-1F3FE-200D-2640-FE0F": {
		Key:        "1F46E-1F3FE-200D-2640-FE0F",
		Value:      "👮🏾‍♀️",
		Descriptor: "woman police officer: medium-dark skin tone",
	},
	"1F46E-1F3FF-200D-2640-FE0F": {
		Key:        "1F46E-1F3FF-200D-2640-FE0F",
		Value:      "👮🏿‍♀️",
		Descriptor: "woman police officer: dark skin tone",
	},
	"1F575-FE0F": {
		Key:        "1F575-FE0F",
		Value:      "🕵️",
		Descriptor: "detective",
	},
	"1F575-1F3FB": {
		Key:        "1F575-1F3FB",
		Value:      "🕵🏻",
		Descriptor: "detective: light skin tone",
	},
	"1F575-1F3FC": {
		Key:        "1F575-1F3FC",
		Value:      "🕵🏼",
		Descriptor: "detective: medium-light skin tone",
	},
	"1F575-1F3FD": {
		Key:        "1F575-1F3FD",
		Value:      "🕵🏽",
		Descriptor: "detective: medium skin tone",
	},
	"1F575-1F3FE": {
		Key:        "1F575-1F3FE",
		Value:      "🕵🏾",
		Descriptor: "detective: medium-dark skin tone",
	},
	"1F575-1F3FF": {
		Key:        "1F575-1F3FF",
		Value:      "🕵🏿",
		Descriptor: "detective: dark skin tone",
	},
	"1F575-FE0F-200D-2642-FE0F": {
		Key:        "1F575-FE0F-200D-2642-FE0F",
		Value:      "🕵️‍♂️",
		Descriptor: "man detective",
	},
	"1F575-1F3FB-200D-2642-FE0F": {
		Key:        "1F575-1F3FB-200D-2642-FE0F",
		Value:      "🕵🏻‍♂️",
		Descriptor: "man detective: light skin tone",
	},
	"1F575-1F3FC-200D-2642-FE0F": {
		Key:        "1F575-1F3FC-200D-2642-FE0F",
		Value:      "🕵🏼‍♂️",
		Descriptor: "man detective: medium-light skin tone",
	},
	"1F575-1F3FD-200D-2642-FE0F": {
		Key:        "1F575-1F3FD-200D-2642-FE0F",
		Value:      "🕵🏽‍♂️",
		Descriptor: "man detective: medium skin tone",
	},
	"1F575-1F3FE-200D-2642-FE0F": {
		Key:        "1F575-1F3FE-200D-2642-FE0F",
		Value:      "🕵🏾‍♂️",
		Descriptor: "man detective: medium-dark skin tone",
	},
	"1F575-1F3FF-200D-2642-FE0F": {
		Key:        "1F575-1F3FF-200D-2642-FE0F",
		Value:      "🕵🏿‍♂️",
		Descriptor: "man detective: dark skin tone",
	},
	"1F575-FE0F-200D-2640-FE0F": {
		Key:        "1F575-FE0F-200D-2640-FE0F",
		Value:      "🕵️‍♀️",
		Descriptor: "woman detective",
	},
	"1F575-1F3FB-200D-2640-FE0F": {
		Key:        "1F575-1F3FB-200D-2640-FE0F",
		Value:      "🕵🏻‍♀️",
		Descriptor: "woman detective: light skin tone",
	},
	"1F575-1F3FC-200D-2640-FE0F": {
		Key:        "1F575-1F3FC-200D-2640-FE0F",
		Value:      "🕵🏼‍♀️",
		Descriptor: "woman detective: medium-light skin tone",
	},
	"1F575-1F3FD-200D-2640-FE0F": {
		Key:        "1F575-1F3FD-200D-2640-FE0F",
		Value:      "🕵🏽‍♀️",
		Descriptor: "woman detective: medium skin tone",
	},
	"1F575-1F3FE-200D-2640-FE0F": {
		Key:        "1F575-1F3FE-200D-2640-FE0F",
		Value:      "🕵🏾‍♀️",
		Descriptor: "woman detective: medium-dark skin tone",
	},
	"1F575-1F3FF-200D-2640-FE0F": {
		Key:        "1F575-1F3FF-200D-2640-FE0F",
		Value:      "🕵🏿‍♀️",
		Descriptor: "woman detective: dark skin tone",
	},
	"1F482": {
		Key:        "1F482",
		Value:      "💂",
		Descriptor: "guard",
	},
	"1F482-1F3FB": {
		Key:        "1F482-1F3FB",
		Value:      "💂🏻",
		Descriptor: "guard: light skin tone",
	},
	"1F482-1F3FC": {
		Key:        "1F482-1F3FC",
		Value:      "💂🏼",
		Descriptor: "guard: medium-light skin tone",
	},
	"1F482-1F3FD": {
		Key:        "1F482-1F3FD",
		Value:      "💂🏽",
		Descriptor: "guard: medium skin tone",
	},
	"1F482-1F3FE": {
		Key:        "1F482-1F3FE",
		Value:      "💂🏾",
		Descriptor: "guard: medium-dark skin tone",
	},
	"1F482-1F3FF": {
		Key:        "1F482-1F3FF",
		Value:      "💂🏿",
		Descriptor: "guard: dark skin tone",
	},
	"1F482-200D-2642-FE0F": {
		Key:        "1F482-200D-2642-FE0F",
		Value:      "💂‍♂️",
		Descriptor: "man guard",
	},
	"1F482-1F3FB-200D-2642-FE0F": {
		Key:        "1F482-1F3FB-200D-2642-FE0F",
		Value:      "💂🏻‍♂️",
		Descriptor: "man guard: light skin tone",
	},
	"1F482-1F3FC-200D-2642-FE0F": {
		Key:        "1F482-1F3FC-200D-2642-FE0F",
		Value:      "💂🏼‍♂️",
		Descriptor: "man guard: medium-light skin tone",
	},
	"1F482-1F3FD-200D-2642-FE0F": {
		Key:        "1F482-1F3FD-200D-2642-FE0F",
		Value:      "💂🏽‍♂️",
		Descriptor: "man guard: medium skin tone",
	},
	"1F482-1F3FE-200D-2642-FE0F": {
		Key:        "1F482-1F3FE-200D-2642-FE0F",
		Value:      "💂🏾‍♂️",
		Descriptor: "man guard: medium-dark skin tone",
	},
	"1F482-1F3FF-200D-2642-FE0F": {
		Key:        "1F482-1F3FF-200D-2642-FE0F",
		Value:      "💂🏿‍♂️",
		Descriptor: "man guard: dark skin tone",
	},
	"1F482-200D-2640-FE0F": {
		Key:        "1F482-200D-2640-FE0F",
		Value:      "💂‍♀️",
		Descriptor: "woman guard",
	},
	"1F482-1F3FB-200D-2640-FE0F": {
		Key:        "1F482-1F3FB-200D-2640-FE0F",
		Value:      "💂🏻‍♀️",
		Descriptor: "woman guard: light skin tone",
	},
	"1F482-1F3FC-200D-2640-FE0F": {
		Key:        "1F482-1F3FC-200D-2640-FE0F",
		Value:      "💂🏼‍♀️",
		Descriptor: "woman guard: medium-light skin tone",
	},
	"1F482-1F3FD-200D-2640-FE0F": {
		Key:        "1F482-1F3FD-200D-2640-FE0F",
		Value:      "💂🏽‍♀️",
		Descriptor: "woman guard: medium skin tone",
	},
	"1F482-1F3FE-200D-2640-FE0F": {
		Key:        "1F482-1F3FE-200D-2640-FE0F",
		Value:      "💂🏾‍♀️",
		Descriptor: "woman guard: medium-dark skin tone",
	},
	"1F482-1F3FF-200D-2640-FE0F": {
		Key:        "1F482-1F3FF-200D-2640-FE0F",
		Value:      "💂🏿‍♀️",
		Descriptor: "woman guard: dark skin tone",
	},
	"1F977": {
		Key:        "1F977",
		Value:      "🥷",
		Descriptor: "ninja",
	},
	"1F977-1F3FB": {
		Key:        "1F977-1F3FB",
		Value:      "🥷🏻",
		Descriptor: "ninja: light skin tone",
	},
	"1F977-1F3FC": {
		Key:        "1F977-1F3FC",
		Value:      "🥷🏼",
		Descriptor: "ninja: medium-light skin tone",
	},
	"1F977-1F3FD": {
		Key:        "1F977-1F3FD",
		Value:      "🥷🏽",
		Descriptor: "ninja: medium skin tone",
	},
	"1F977-1F3FE": {
		Key:        "1F977-1F3FE",
		Value:      "🥷🏾",
		Descriptor: "ninja: medium-dark skin tone",
	},
	"1F977-1F3FF": {
		Key:        "1F977-1F3FF",
		Value:      "🥷🏿",
		Descriptor: "ninja: dark skin tone",
	},
	"1F477": {
		Key:        "1F477",
		Value:      "👷",
		Descriptor: "construction worker",
	},
	"1F477-1F3FB": {
		Key:        "1F477-1F3FB",
		Value:      "👷🏻",
		Descriptor: "construction worker: light skin tone",
	},
	"1F477-1F3FC": {
		Key:        "1F477-1F3FC",
		Value:      "👷🏼",
		Descriptor: "construction worker: medium-light skin tone",
	},
	"1F477-1F3FD": {
		Key:        "1F477-1F3FD",
		Value:      "👷🏽",
		Descriptor: "construction worker: medium skin tone",
	},
	"1F477-1F3FE": {
		Key:        "1F477-1F3FE",
		Value:      "👷🏾",
		Descriptor: "construction worker: medium-dark skin tone",
	},
	"1F477-1F3FF": {
		Key:        "1F477-1F3FF",
		Value:      "👷🏿",
		Descriptor: "construction worker: dark skin tone",
	},
	"1F477-200D-2642-FE0F": {
		Key:        "1F477-200D-2642-FE0F",
		Value:      "👷‍♂️",
		Descriptor: "man construction worker",
	},
	"1F477-1F3FB-200D-2642-FE0F": {
		Key:        "1F477-1F3FB-200D-2642-FE0F",
		Value:      "👷🏻‍♂️",
		Descriptor: "man construction worker: light skin tone",
	},
	"1F477-1F3FC-200D-2642-FE0F": {
		Key:        "1F477-1F3FC-200D-2642-FE0F",
		Value:      "👷🏼‍♂️",
		Descriptor: "man construction worker: medium-light skin tone",
	},
	"1F477-1F3FD-200D-2642-FE0F": {
		Key:        "1F477-1F3FD-200D-2642-FE0F",
		Value:      "👷🏽‍♂️",
		Descriptor: "man construction worker: medium skin tone",
	},
	"1F477-1F3FE-200D-2642-FE0F": {
		Key:        "1F477-1F3FE-200D-2642-FE0F",
		Value:      "👷🏾‍♂️",
		Descriptor: "man construction worker: medium-dark skin tone",
	},
	"1F477-1F3FF-200D-2642-FE0F": {
		Key:        "1F477-1F3FF-200D-2642-FE0F",
		Value:      "👷🏿‍♂️",
		Descriptor: "man construction worker: dark skin tone",
	},
	"1F477-200D-2640-FE0F": {
		Key:        "1F477-200D-2640-FE0F",
		Value:      "👷‍♀️",
		Descriptor: "woman construction worker",
	},
	"1F477-1F3FB-200D-2640-FE0F": {
		Key:        "1F477-1F3FB-200D-2640-FE0F",
		Value:      "👷🏻‍♀️",
		Descriptor: "woman construction worker: light skin tone",
	},
	"1F477-1F3FC-200D-2640-FE0F": {
		Key:        "1F477-1F3FC-200D-2640-FE0F",
		Value:      "👷🏼‍♀️",
		Descriptor: "woman construction worker: medium-light skin tone",
	},
	"1F477-1F3FD-200D-2640-FE0F": {
		Key:        "1F477-1F3FD-200D-2640-FE0F",
		Value:      "👷🏽‍♀️",
		Descriptor: "woman construction worker: medium skin tone",
	},
	"1F477-1F3FE-200D-2640-FE0F": {
		Key:        "1F477-1F3FE-200D-2640-FE0F",
		Value:      "👷🏾‍♀️",
		Descriptor: "woman construction worker: medium-dark skin tone",
	},
	"1F477-1F3FF-200D-2640-FE0F": {
		Key:        "1F477-1F3FF-200D-2640-FE0F",
		Value:      "👷🏿‍♀️",
		Descriptor: "woman construction worker: dark skin tone",
	},
	"1FAC5": {
		Key:        "1FAC5",
		Value:      "🫅",
		Descriptor: "person with crown",
	},
	"1FAC5-1F3FB": {
		Key:        "1FAC5-1F3FB",
		Value:      "🫅🏻",
		Descriptor: "person with crown: light skin tone",
	},
	"1FAC5-1F3FC": {
		Key:        "1FAC5-1F3FC",
		Value:      "🫅🏼",
		Descriptor: "person with crown: medium-light skin tone",
	},
	"1FAC5-1F3FD": {
		Key:        "1FAC5-1F3FD",
		Value:      "🫅🏽",
		Descriptor: "person with crown: medium skin tone",
	},
	"1FAC5-1F3FE": {
		Key:        "1FAC5-1F3FE",
		Value:      "🫅🏾",
		Descriptor: "person with crown: medium-dark skin tone",
	},
	"1FAC5-1F3FF": {
		Key:        "1FAC5-1F3FF",
		Value:      "🫅🏿",
		Descriptor: "person with crown: dark skin tone",
	},
	"1F934": {
		Key:        "1F934",
		Value:      "🤴",
		Descriptor: "prince",
	},
	"1F934-1F3FB": {
		Key:        "1F934-1F3FB",
		Value:      "🤴🏻",
		Descriptor: "prince: light skin tone",
	},
	"1F934-1F3FC": {
		Key:        "1F934-1F3FC",
		Value:      "🤴🏼",
		Descriptor: "prince: medium-light skin tone",
	},
	"1F934-1F3FD": {
		Key:        "1F934-1F3FD",
		Value:      "🤴🏽",
		Descriptor: "prince: medium skin tone",
	},
	"1F934-1F3FE": {
		Key:        "1F934-1F3FE",
		Value:      "🤴🏾",
		Descriptor: "prince: medium-dark skin tone",
	},
	"1F934-1F3FF": {
		Key:        "1F934-1F3FF",
		Value:      "🤴🏿",
		Descriptor: "prince: dark skin tone",
	},
	"1F478": {
		Key:        "1F478",
		Value:      "👸",
		Descriptor: "princess",
	},
	"1F478-1F3FB": {
		Key:        "1F478-1F3FB",
		Value:      "👸🏻",
		Descriptor: "princess: light skin tone",
	},
	"1F478-1F3FC": {
		Key:        "1F478-1F3FC",
		Value:      "👸🏼",
		Descriptor: "princess: medium-light skin tone",
	},
	"1F478-1F3FD": {
		Key:        "1F478-1F3FD",
		Value:      "👸🏽",
		Descriptor: "princess: medium skin tone",
	},
	"1F478-1F3FE": {
		Key:        "1F478-1F3FE",
		Value:      "👸🏾",
		Descriptor: "princess: medium-dark skin tone",
	},
	"1F478-1F3FF": {
		Key:        "1F478-1F3FF",
		Value:      "👸🏿",
		Descriptor: "princess: dark skin tone",
	},
	"1F473": {
		Key:        "1F473",
		Value:      "👳",
		Descriptor: "person wearing turban",
	},
	"1F473-1F3FB": {
		Key:        "1F473-1F3FB",
		Value:      "👳🏻",
		Descriptor: "person wearing turban: light skin tone",
	},
	"1F473-1F3FC": {
		Key:        "1F473-1F3FC",
		Value:      "👳🏼",
		Descriptor: "person wearing turban: medium-light skin tone",
	},
	"1F473-1F3FD": {
		Key:        "1F473-1F3FD",
		Value:      "👳🏽",
		Descriptor: "person wearing turban: medium skin tone",
	},
	"1F473-1F3FE": {
		Key:        "1F473-1F3FE",
		Value:      "👳🏾",
		Descriptor: "person wearing turban: medium-dark skin tone",
	},
	"1F473-1F3FF": {
		Key:        "1F473-1F3FF",
		Value:      "👳🏿",
		Descriptor: "person wearing turban: dark skin tone",
	},
	"1F473-200D-2642-FE0F": {
		Key:        "1F473-200D-2642-FE0F",
		Value:      "👳‍♂️",
		Descriptor: "man wearing turban",
	},
	"1F473-1F3FB-200D-2642-FE0F": {
		Key:        "1F473-1F3FB-200D-2642-FE0F",
		Value:      "👳🏻‍♂️",
		Descriptor: "man wearing turban: light skin tone",
	},
	"1F473-1F3FC-200D-2642-FE0F": {
		Key:        "1F473-1F3FC-200D-2642-FE0F",
		Value:      "👳🏼‍♂️",
		Descriptor: "man wearing turban: medium-light skin tone",
	},
	"1F473-1F3FD-200D-2642-FE0F": {
		Key:        "1F473-1F3FD-200D-2642-FE0F",
		Value:      "👳🏽‍♂️",
		Descriptor: "man wearing turban: medium skin tone",
	},
	"1F473-1F3FE-200D-2642-FE0F": {
		Key:        "1F473-1F3FE-200D-2642-FE0F",
		Value:      "👳🏾‍♂️",
		Descriptor: "man wearing turban: medium-dark skin tone",
	},
	"1F473-1F3FF-200D-2642-FE0F": {
		Key:        "1F473-1F3FF-200D-2642-FE0F",
		Value:      "👳🏿‍♂️",
		Descriptor: "man wearing turban: dark skin tone",
	},
	"1F473-200D-2640-FE0F": {
		Key:        "1F473-200D-2640-FE0F",
		Value:      "👳‍♀️",
		Descriptor: "woman wearing turban",
	},
	"1F473-1F3FB-200D-2640-FE0F": {
		Key:        "1F473-1F3FB-200D-2640-FE0F",
		Value:      "👳🏻‍♀️",
		Descriptor: "woman wearing turban: light skin tone",
	},
	"1F473-1F3FC-200D-2640-FE0F": {
		Key:        "1F473-1F3FC-200D-2640-FE0F",
		Value:      "👳🏼‍♀️",
		Descriptor: "woman wearing turban: medium-light skin tone",
	},
	"1F473-1F3FD-200D-2640-FE0F": {
		Key:        "1F473-1F3FD-200D-2640-FE0F",
		Value:      "👳🏽‍♀️",
		Descriptor: "woman wearing turban: medium skin tone",
	},
	"1F473-1F3FE-200D-2640-FE0F": {
		Key:        "1F473-1F3FE-200D-2640-FE0F",
		Value:      "👳🏾‍♀️",
		Descriptor: "woman wearing turban: medium-dark skin tone",
	},
	"1F473-1F3FF-200D-2640-FE0F": {
		Key:        "1F473-1F3FF-200D-2640-FE0F",
		Value:      "👳🏿‍♀️",
		Descriptor: "woman wearing turban: dark skin tone",
	},
	"1F472": {
		Key:        "1F472",
		Value:      "👲",
		Descriptor: "person with skullcap",
	},
	"1F472-1F3FB": {
		Key:        "1F472-1F3FB",
		Value:      "👲🏻",
		Descriptor: "person with skullcap: light skin tone",
	},
	"1F472-1F3FC": {
		Key:        "1F472-1F3FC",
		Value:      "👲🏼",
		Descriptor: "person with skullcap: medium-light skin tone",
	},
	"1F472-1F3FD": {
		Key:        "1F472-1F3FD",
		Value:      "👲🏽",
		Descriptor: "person with skullcap: medium skin tone",
	},
	"1F472-1F3FE": {
		Key:        "1F472-1F3FE",
		Value:      "👲🏾",
		Descriptor: "person with skullcap: medium-dark skin tone",
	},
	"1F472-1F3FF": {
		Key:        "1F472-1F3FF",
		Value:      "👲🏿",
		Descriptor: "person with skullcap: dark skin tone",
	},
	"1F9D5": {
		Key:        "1F9D5",
		Value:      "🧕",
		Descriptor: "woman with headscarf",
	},
	"1F9D5-1F3FB": {
		Key:        "1F9D5-1F3FB",
		Value:      "🧕🏻",
		Descriptor: "woman with headscarf: light skin tone",
	},
	"1F9D5-1F3FC": {
		Key:        "1F9D5-1F3FC",
		Value:      "🧕🏼",
		Descriptor: "woman with headscarf: medium-light skin tone",
	},
	"1F9D5-1F3FD": {
		Key:        "1F9D5-1F3FD",
		Value:      "🧕🏽",
		Descriptor: "woman with headscarf: medium skin tone",
	},
	"1F9D5-1F3FE": {
		Key:        "1F9D5-1F3FE",
		Value:      "🧕🏾",
		Descriptor: "woman with headscarf: medium-dark skin tone",
	},
	"1F9D5-1F3FF": {
		Key:        "1F9D5-1F3FF",
		Value:      "🧕🏿",
		Descriptor: "woman with headscarf: dark skin tone",
	},
	"1F935": {
		Key:        "1F935",
		Value:      "🤵",
		Descriptor: "person in tuxedo",
	},
	"1F935-1F3FB": {
		Key:        "1F935-1F3FB",
		Value:      "🤵🏻",
		Descriptor: "person in tuxedo: light skin tone",
	},
	"1F935-1F3FC": {
		Key:        "1F935-1F3FC",
		Value:      "🤵🏼",
		Descriptor: "person in tuxedo: medium-light skin tone",
	},
	"1F935-1F3FD": {
		Key:        "1F935-1F3FD",
		Value:      "🤵🏽",
		Descriptor: "person in tuxedo: medium skin tone",
	},
	"1F935-1F3FE": {
		Key:        "1F935-1F3FE",
		Value:      "🤵🏾",
		Descriptor: "person in tuxedo: medium-dark skin tone",
	},
	"1F935-1F3FF": {
		Key:        "1F935-1F3FF",
		Value:      "🤵🏿",
		Descriptor: "person in tuxedo: dark skin tone",
	},
	"1F935-200D-2642-FE0F": {
		Key:        "1F935-200D-2642-FE0F",
		Value:      "🤵‍♂️",
		Descriptor: "man in tuxedo",
	},
	"1F935-1F3FB-200D-2642-FE0F": {
		Key:        "1F935-1F3FB-200D-2642-FE0F",
		Value:      "🤵🏻‍♂️",
		Descriptor: "man in tuxedo: light skin tone",
	},
	"1F935-1F3FC-200D-2642-FE0F": {
		Key:        "1F935-1F3FC-200D-2642-FE0F",
		Value:      "🤵🏼‍♂️",
		Descriptor: "man in tuxedo: medium-light skin tone",
	},
	"1F935-1F3FD-200D-2642-FE0F": {
		Key:        "1F935-1F3FD-200D-2642-FE0F",
		Value:      "🤵🏽‍♂️",
		Descriptor: "man in tuxedo: medium skin tone",
	},
	"1F935-1F3FE-200D-2642-FE0F": {
		Key:        "1F935-1F3FE-200D-2642-FE0F",
		Value:      "🤵🏾‍♂️",
		Descriptor: "man in tuxedo: medium-dark skin tone",
	},
	"1F935-1F3FF-200D-2642-FE0F": {
		Key:        "1F935-1F3FF-200D-2642-FE0F",
		Value:      "🤵🏿‍♂️",
		Descriptor: "man in tuxedo: dark skin tone",
	},
	"1F935-200D-2640-FE0F": {
		Key:        "1F935-200D-2640-FE0F",
		Value:      "🤵‍♀️",
		Descriptor: "woman in tuxedo",
	},
	"1F935-1F3FB-200D-2640-FE0F": {
		Key:        "1F935-1F3FB-200D-2640-FE0F",
		Value:      "🤵🏻‍♀️",
		Descriptor: "woman in tuxedo: light skin tone",
	},
	"1F935-1F3FC-200D-2640-FE0F": {
		Key:        "1F935-1F3FC-200D-2640-FE0F",
		Value:      "🤵🏼‍♀️",
		Descriptor: "woman in tuxedo: medium-light skin tone",
	},
	"1F935-1F3FD-200D-2640-FE0F": {
		Key:        "1F935-1F3FD-200D-2640-FE0F",
		Value:      "🤵🏽‍♀️",
		Descriptor: "woman in tuxedo: medium skin tone",
	},
	"1F935-1F3FE-200D-2640-FE0F": {
		Key:        "1F935-1F3FE-200D-2640-FE0F",
		Value:      "🤵🏾‍♀️",
		Descriptor: "woman in tuxedo: medium-dark skin tone",
	},
	"1F935-1F3FF-200D-2640-FE0F": {
		Key:        "1F935-1F3FF-200D-2640-FE0F",
		Value:      "🤵🏿‍♀️",
		Descriptor: "woman in tuxedo: dark skin tone",
	},
	"1F470": {
		Key:        "1F470",
		Value:      "👰",
		Descriptor: "person with veil",
	},
	"1F470-1F3FB": {
		Key:        "1F470-1F3FB",
		Value:      "👰🏻",
		Descriptor: "person with veil: light skin tone",
	},
	"1F470-1F3FC": {
		Key:        "1F470-1F3FC",
		Value:      "👰🏼",
		Descriptor: "person with veil: medium-light skin tone",
	},
	"1F470-1F3FD": {
		Key:        "1F470-1F3FD",
		Value:      "👰🏽",
		Descriptor: "person with veil: medium skin tone",
	},
	"1F470-1F3FE": {
		Key:        "1F470-1F3FE",
		Value:      "👰🏾",
		Descriptor: "person with veil: medium-dark skin tone",
	},
	"1F470-1F3FF": {
		Key:        "1F470-1F3FF",
		Value:      "👰🏿",
		Descriptor: "person with veil: dark skin tone",
	},
	"1F470-200D-2642-FE0F": {
		Key:        "1F470-200D-2642-FE0F",
		Value:      "👰‍♂️",
		Descriptor: "man with veil",
	},
	"1F470-1F3FB-200D-2642-FE0F": {
		Key:        "1F470-1F3FB-200D-2642-FE0F",
		Value:      "👰🏻‍♂️",
		Descriptor: "man with veil: light skin tone",
	},
	"1F470-1F3FC-200D-2642-FE0F": {
		Key:        "1F470-1F3FC-200D-2642-FE0F",
		Value:      "👰🏼‍♂️",
		Descriptor: "man with veil: medium-light skin tone",
	},
	"1F470-1F3FD-200D-2642-FE0F": {
		Key:        "1F470-1F3FD-200D-2642-FE0F",
		Value:      "👰🏽‍♂️",
		Descriptor: "man with veil: medium skin tone",
	},
	"1F470-1F3FE-200D-2642-FE0F": {
		Key:        "1F470-1F3FE-200D-2642-FE0F",
		Value:      "👰🏾‍♂️",
		Descriptor: "man with veil: medium-dark skin tone",
	},
	"1F470-1F3FF-200D-2642-FE0F": {
		Key:        "1F470-1F3FF-200D-2642-FE0F",
		Value:      "👰🏿‍♂️",
		Descriptor: "man with veil: dark skin tone",
	},
	"1F470-200D-2640-FE0F": {
		Key:        "1F470-200D-2640-FE0F",
		Value:      "👰‍♀️",
		Descriptor: "woman with veil",
	},
	"1F470-1F3FB-200D-2640-FE0F": {
		Key:        "1F470-1F3FB-200D-2640-FE0F",
		Value:      "👰🏻‍♀️",
		Descriptor: "woman with veil: light skin tone",
	},
	"1F470-1F3FC-200D-2640-FE0F": {
		Key:        "1F470-1F3FC-200D-2640-FE0F",
		Value:      "👰🏼‍♀️",
		Descriptor: "woman with veil: medium-light skin tone",
	},
	"1F470-1F3FD-200D-2640-FE0F": {
		Key:        "1F470-1F3FD-200D-2640-FE0F",
		Value:      "👰🏽‍♀️",
		Descriptor: "woman with veil: medium skin tone",
	},
	"1F470-1F3FE-200D-2640-FE0F": {
		Key:        "1F470-1F3FE-200D-2640-FE0F",
		Value:      "👰🏾‍♀️",
		Descriptor: "woman with veil: medium-dark skin tone",
	},
	"1F470-1F3FF-200D-2640-FE0F": {
		Key:        "1F470-1F3FF-200D-2640-FE0F",
		Value:      "👰🏿‍♀️",
		Descriptor: "woman with veil: dark skin tone",
	},
	"1F930": {
		Key:        "1F930",
		Value:      "🤰",
		Descriptor: "pregnant woman",
	},
	"1F930-1F3FB": {
		Key:        "1F930-1F3FB",
		Value:      "🤰🏻",
		Descriptor: "pregnant woman: light skin tone",
	},
	"1F930-1F3FC": {
		Key:        "1F930-1F3FC",
		Value:      "🤰🏼",
		Descriptor: "pregnant woman: medium-light skin tone",
	},
	"1F930-1F3FD": {
		Key:        "1F930-1F3FD",
		Value:      "🤰🏽",
		Descriptor: "pregnant woman: medium skin tone",
	},
	"1F930-1F3FE": {
		Key:        "1F930-1F3FE",
		Value:      "🤰🏾",
		Descriptor: "pregnant woman: medium-dark skin tone",
	},
	"1F930-1F3FF": {
		Key:        "1F930-1F3FF",
		Value:      "🤰🏿",
		Descriptor: "pregnant woman: dark skin tone",
	},
	"1FAC3": {
		Key:        "1FAC3",
		Value:      "🫃",
		Descriptor: "pregnant man",
	},
	"1FAC3-1F3FB": {
		Key:        "1FAC3-1F3FB",
		Value:      "🫃🏻",
		Descriptor: "pregnant man: light skin tone",
	},
	"1FAC3-1F3FC": {
		Key:        "1FAC3-1F3FC",
		Value:      "🫃🏼",
		Descriptor: "pregnant man: medium-light skin tone",
	},
	"1FAC3-1F3FD": {
		Key:        "1FAC3-1F3FD",
		Value:      "🫃🏽",
		Descriptor: "pregnant man: medium skin tone",
	},
	"1FAC3-1F3FE": {
		Key:        "1FAC3-1F3FE",
		Value:      "🫃🏾",
		Descriptor: "pregnant man: medium-dark skin tone",
	},
	"1FAC3-1F3FF": {
		Key:        "1FAC3-1F3FF",
		Value:      "🫃🏿",
		Descriptor: "pregnant man: dark skin tone",
	},
	"1FAC4": {
		Key:        "1FAC4",
		Value:      "🫄",
		Descriptor: "pregnant person",
	},
	"1FAC4-1F3FB": {
		Key:        "1FAC4-1F3FB",
		Value:      "🫄🏻",
		Descriptor: "pregnant person: light skin tone",
	},
	"1FAC4-1F3FC": {
		Key:        "1FAC4-1F3FC",
		Value:      "🫄🏼",
		Descriptor: "pregnant person: medium-light skin tone",
	},
	"1FAC4-1F3FD": {
		Key:        "1FAC4-1F3FD",
		Value:      "🫄🏽",
		Descriptor: "pregnant person: medium skin tone",
	},
	"1FAC4-1F3FE": {
		Key:        "1FAC4-1F3FE",
		Value:      "🫄🏾",
		Descriptor: "pregnant person: medium-dark skin tone",
	},
	"1FAC4-1F3FF": {
		Key:        "1FAC4-1F3FF",
		Value:      "🫄🏿",
		Descriptor: "pregnant person: dark skin tone",
	},
	"1F931": {
		Key:        "1F931",
		Value:      "🤱",
		Descriptor: "breast-feeding",
	},
	"1F931-1F3FB": {
		Key:        "1F931-1F3FB",
		Value:      "🤱🏻",
		Descriptor: "breast-feeding: light skin tone",
	},
	"1F931-1F3FC": {
		Key:        "1F931-1F3FC",
		Value:      "🤱🏼",
		Descriptor: "breast-feeding: medium-light skin tone",
	},
	"1F931-1F3FD": {
		Key:        "1F931-1F3FD",
		Value:      "🤱🏽",
		Descriptor: "breast-feeding: medium skin tone",
	},
	"1F931-1F3FE": {
		Key:        "1F931-1F3FE",
		Value:      "🤱🏾",
		Descriptor: "breast-feeding: medium-dark skin tone",
	},
	"1F931-1F3FF": {
		Key:        "1F931-1F3FF",
		Value:      "🤱🏿",
		Descriptor: "breast-feeding: dark skin tone",
	},
	"1F469-200D-1F37C": {
		Key:        "1F469-200D-1F37C",
		Value:      "👩‍🍼",
		Descriptor: "woman feeding baby",
	},
	"1F469-1F3FB-200D-1F37C": {
		Key:        "1F469-1F3FB-200D-1F37C",
		Value:      "👩🏻‍🍼",
		Descriptor: "woman feeding baby: light skin tone",
	},
	"1F469-1F3FC-200D-1F37C": {
		Key:        "1F469-1F3FC-200D-1F37C",
		Value:      "👩🏼‍🍼",
		Descriptor: "woman feeding baby: medium-light skin tone",
	},
	"1F469-1F3FD-200D-1F37C": {
		Key:        "1F469-1F3FD-200D-1F37C",
		Value:      "👩🏽‍🍼",
		Descriptor: "woman feeding baby: medium skin tone",
	},
	"1F469-1F3FE-200D-1F37C": {
		Key:        "1F469-1F3FE-200D-1F37C",
		Value:      "👩🏾‍🍼",
		Descriptor: "woman feeding baby: medium-dark skin tone",
	},
	"1F469-1F3FF-200D-1F37C": {
		Key:        "1F469-1F3FF-200D-1F37C",
		Value:      "👩🏿‍🍼",
		Descriptor: "woman feeding baby: dark skin tone",
	},
	"1F468-200D-1F37C": {
		Key:        "1F468-200D-1F37C",
		Value:      "👨‍🍼",
		Descriptor: "man feeding baby",
	},
	"1F468-1F3FB-200D-1F37C": {
		Key:        "1F468-1F3FB-200D-1F37C",
		Value:      "👨🏻‍🍼",
		Descriptor: "man feeding baby: light skin tone",
	},
	"1F468-1F3FC-200D-1F37C": {
		Key:        "1F468-1F3FC-200D-1F37C",
		Value:      "👨🏼‍🍼",
		Descriptor: "man feeding baby: medium-light skin tone",
	},
	"1F468-1F3FD-200D-1F37C": {
		Key:        "1F468-1F3FD-200D-1F37C",
		Value:      "👨🏽‍🍼",
		Descriptor: "man feeding baby: medium skin tone",
	},
	"1F468-1F3FE-200D-1F37C": {
		Key:        "1F468-1F3FE-200D-1F37C",
		Value:      "👨🏾‍🍼",
		Descriptor: "man feeding baby: medium-dark skin tone",
	},
	"1F468-1F3FF-200D-1F37C": {
		Key:        "1F468-1F3FF-200D-1F37C",
		Value:      "👨🏿‍🍼",
		Descriptor: "man feeding baby: dark skin tone",
	},
	"1F9D1-200D-1F37C": {
		Key:        "1F9D1-200D-1F37C",
		Value:      "🧑‍🍼",
		Descriptor: "person feeding baby",
	},
	"1F9D1-1F3FB-200D-1F37C": {
		Key:        "1F9D1-1F3FB-200D-1F37C",
		Value:      "🧑🏻‍🍼",
		Descriptor: "person feeding baby: light skin tone",
	},
	"1F9D1-1F3FC-200D-1F37C": {
		Key:        "1F9D1-1F3FC-200D-1F37C",
		Value:      "🧑🏼‍🍼",
		Descriptor: "person feeding baby: medium-light skin tone",
	},
	"1F9D1-1F3FD-200D-1F37C": {
		Key:        "1F9D1-1F3FD-200D-1F37C",
		Value:      "🧑🏽‍🍼",
		Descriptor: "person feeding baby: medium skin tone",
	},
	"1F9D1-1F3FE-200D-1F37C": {
		Key:        "1F9D1-1F3FE-200D-1F37C",
		Value:      "🧑🏾‍🍼",
		Descriptor: "person feeding baby: medium-dark skin tone",
	},
	"1F9D1-1F3FF-200D-1F37C": {
		Key:        "1F9D1-1F3FF-200D-1F37C",
		Value:      "🧑🏿‍🍼",
		Descriptor: "person feeding baby: dark skin tone",
	},
	"1F47C": {
		Key:        "1F47C",
		Value:      "👼",
		Descriptor: "baby angel",
	},
	"1F47C-1F3FB": {
		Key:        "1F47C-1F3FB",
		Value:      "👼🏻",
		Descriptor: "baby angel: light skin tone",
	},
	"1F47C-1F3FC": {
		Key:        "1F47C-1F3FC",
		Value:      "👼🏼",
		Descriptor: "baby angel: medium-light skin tone",
	},
	"1F47C-1F3FD": {
		Key:        "1F47C-1F3FD",
		Value:      "👼🏽",
		Descriptor: "baby angel: medium skin tone",
	},
	"1F47C-1F3FE": {
		Key:        "1F47C-1F3FE",
		Value:      "👼🏾",
		Descriptor: "baby angel: medium-dark skin tone",
	},
	"1F47C-1F3FF": {
		Key:        "1F47C-1F3FF",
		Value:      "👼🏿",
		Descriptor: "baby angel: dark skin tone",
	},
	"1F385": {
		Key:        "1F385",
		Value:      "🎅",
		Descriptor: "Santa Claus",
	},
	"1F385-1F3FB": {
		Key:        "1F385-1F3FB",
		Value:      "🎅🏻",
		Descriptor: "Santa Claus: light skin tone",
	},
	"1F385-1F3FC": {
		Key:        "1F385-1F3FC",
		Value:      "🎅🏼",
		Descriptor: "Santa Claus: medium-light skin tone",
	},
	"1F385-1F3FD": {
		Key:        "1F385-1F3FD",
		Value:      "🎅🏽",
		Descriptor: "Santa Claus: medium skin tone",
	},
	"1F385-1F3FE": {
		Key:        "1F385-1F3FE",
		Value:      "🎅🏾",
		Descriptor: "Santa Claus: medium-dark skin tone",
	},
	"1F385-1F3FF": {
		Key:        "1F385-1F3FF",
		Value:      "🎅🏿",
		Descriptor: "Santa Claus: dark skin tone",
	},
	"1F936": {
		Key:        "1F936",
		Value:      "🤶",
		Descriptor: "Mrs. Claus",
	},
	"1F936-1F3FB": {
		Key:        "1F936-1F3FB",
		Value:      "🤶🏻",
		Descriptor: "Mrs. Claus: light skin tone",
	},
	"1F936-1F3FC": {
		Key:        "1F936-1F3FC",
		Value:      "🤶🏼",
		Descriptor: "Mrs. Claus: medium-light skin tone",
	},
	"1F936-1F3FD": {
		Key:        "1F936-1F3FD",
		Value:      "🤶🏽",
		Descriptor: "Mrs. Claus: medium skin tone",
	},
	"1F936-1F3FE": {
		Key:        "1F936-1F3FE",
		Value:      "🤶🏾",
		Descriptor: "Mrs. Claus: medium-dark skin tone",
	},
	"1F936-1F3FF": {
		Key:        "1F936-1F3FF",
		Value:      "🤶🏿",
		Descriptor: "Mrs. Claus: dark skin tone",
	},
	"1F9D1-200D-1F384": {
		Key:        "1F9D1-200D-1F384",
		Value:      "🧑‍🎄",
		Descriptor: "mx claus",
	},
	"1F9D1-1F3FB-200D-1F384": {
		Key:        "1F9D1-1F3FB-200D-1F384",
		Value:      "🧑🏻‍🎄",
		Descriptor: "mx claus: light skin tone",
	},
	"1F9D1-1F3FC-200D-1F384": {
		Key:        "1F9D1-1F3FC-200D-1F384",
		Value:      "🧑🏼‍🎄",
		Descriptor: "mx claus: medium-light skin tone",
	},
	"1F9D1-1F3FD-200D-1F384": {
		Key:        "1F9D1-1F3FD-200D-1F384",
		Value:      "🧑🏽‍🎄",
		Descriptor: "mx claus: medium skin tone",
	},
	"1F9D1-1F3FE-200D-1F384": {
		Key:        "1F9D1-1F3FE-200D-1F384",
		Value:      "🧑🏾‍🎄",
		Descriptor: "mx claus: medium-dark skin tone",
	},
	"1F9D1-1F3FF-200D-1F384": {
		Key:        "1F9D1-1F3FF-200D-1F384",
		Value:      "🧑🏿‍🎄",
		Descriptor: "mx claus: dark skin tone",
	},
	"1F9B8": {
		Key:        "1F9B8",
		Value:      "🦸",
		Descriptor: "superhero",
	},
	"1F9B8-1F3FB": {
		Key:        "1F9B8-1F3FB",
		Value:      "🦸🏻",
		Descriptor: "superhero: light skin tone",
	},
	"1F9B8-1F3FC": {
		Key:        "1F9B8-1F3FC",
		Value:      "🦸🏼",
		Descriptor: "superhero: medium-light skin tone",
	},
	"1F9B8-1F3FD": {
		Key:        "1F9B8-1F3FD",
		Value:      "🦸🏽",
		Descriptor: "superhero: medium skin tone",
	},
	"1F9B8-1F3FE": {
		Key:        "1F9B8-1F3FE",
		Value:      "🦸🏾",
		Descriptor: "superhero: medium-dark skin tone",
	},
	"1F9B8-1F3FF": {
		Key:        "1F9B8-1F3FF",
		Value:      "🦸🏿",
		Descriptor: "superhero: dark skin tone",
	},
	"1F9B8-200D-2642-FE0F": {
		Key:        "1F9B8-200D-2642-FE0F",
		Value:      "🦸‍♂️",
		Descriptor: "man superhero",
	},
	"1F9B8-1F3FB-200D-2642-FE0F": {
		Key:        "1F9B8-1F3FB-200D-2642-FE0F",
		Value:      "🦸🏻‍♂️",
		Descriptor: "man superhero: light skin tone",
	},
	"1F9B8-1F3FC-200D-2642-FE0F": {
		Key:        "1F9B8-1F3FC-200D-2642-FE0F",
		Value:      "🦸🏼‍♂️",
		Descriptor: "man superhero: medium-light skin tone",
	},
	"1F9B8-1F3FD-200D-2642-FE0F": {
		Key:        "1F9B8-1F3FD-200D-2642-FE0F",
		Value:      "🦸🏽‍♂️",
		Descriptor: "man superhero: medium skin tone",
	},
	"1F9B8-1F3FE-200D-2642-FE0F": {
		Key:        "1F9B8-1F3FE-200D-2642-FE0F",
		Value:      "🦸🏾‍♂️",
		Descriptor: "man superhero: medium-dark skin tone",
	},
	"1F9B8-1F3FF-200D-2642-FE0F": {
		Key:        "1F9B8-1F3FF-200D-2642-FE0F",
		Value:      "🦸🏿‍♂️",
		Descriptor: "man superhero: dark skin tone",
	},
	"1F9B8-200D-2640-FE0F": {
		Key:        "1F9B8-200D-2640-FE0F",
		Value:      "🦸‍♀️",
		Descriptor: "woman superhero",
	},
	"1F9B8-1F3FB-200D-2640-FE0F": {
		Key:        "1F9B8-1F3FB-200D-2640-FE0F",
		Value:      "🦸🏻‍♀️",
		Descriptor: "woman superhero: light skin tone",
	},
	"1F9B8-1F3FC-200D-2640-FE0F": {
		Key:        "1F9B8-1F3FC-200D-2640-FE0F",
		Value:      "🦸🏼‍♀️",
		Descriptor: "woman superhero: medium-light skin tone",
	},
	"1F9B8-1F3FD-200D-2640-FE0F": {
		Key:        "1F9B8-1F3FD-200D-2640-FE0F",
		Value:      "🦸🏽‍♀️",
		Descriptor: "woman superhero: medium skin tone",
	},
	"1F9B8-1F3FE-200D-2640-FE0F": {
		Key:        "1F9B8-1F3FE-200D-2640-FE0F",
		Value:      "🦸🏾‍♀️",
		Descriptor: "woman superhero: medium-dark skin tone",
	},
	"1F9B8-1F3FF-200D-2640-FE0F": {
		Key:        "1F9B8-1F3FF-200D-2640-FE0F",
		Value:      "🦸🏿‍♀️",
		Descriptor: "woman superhero: dark skin tone",
	},
	"1F9B9": {
		Key:        "1F9B9",
		Value:      "🦹",
		Descriptor: "supervillain",
	},
	"1F9B9-1F3FB": {
		Key:        "1F9B9-1F3FB",
		Value:      "🦹🏻",
		Descriptor: "supervillain: light skin tone",
	},
	"1F9B9-1F3FC": {
		Key:        "1F9B9-1F3FC",
		Value:      "🦹🏼",
		Descriptor: "supervillain: medium-light skin tone",
	},
	"1F9B9-1F3FD": {
		Key:        "1F9B9-1F3FD",
		Value:      "🦹🏽",
		Descriptor: "supervillain: medium skin tone",
	},
	"1F9B9-1F3FE": {
		Key:        "1F9B9-1F3FE",
		Value:      "🦹🏾",
		Descriptor: "supervillain: medium-dark skin tone",
	},
	"1F9B9-1F3FF": {
		Key:        "1F9B9-1F3FF",
		Value:      "🦹🏿",
		Descriptor: "supervillain: dark skin tone",
	},
	"1F9B9-200D-2642-FE0F": {
		Key:        "1F9B9-200D-2642-FE0F",
		Value:      "🦹‍♂️",
		Descriptor: "man supervillain",
	},
	"1F9B9-1F3FB-200D-2642-FE0F": {
		Key:        "1F9B9-1F3FB-200D-2642-FE0F",
		Value:      "🦹🏻‍♂️",
		Descriptor: "man supervillain: light skin tone",
	},
	"1F9B9-1F3FC-200D-2642-FE0F": {
		Key:        "1F9B9-1F3FC-200D-2642-FE0F",
		Value:      "🦹🏼‍♂️",
		Descriptor: "man supervillain: medium-light skin tone",
	},
	"1F9B9-1F3FD-200D-2642-FE0F": {
		Key:        "1F9B9-1F3FD-200D-2642-FE0F",
		Value:      "🦹🏽‍♂️",
		Descriptor: "man supervillain: medium skin tone",
	},
	"1F9B9-1F3FE-200D-2642-FE0F": {
		Key:        "1F9B9-1F3FE-200D-2642-FE0F",
		Value:      "🦹🏾‍♂️",
		Descriptor: "man supervillain: medium-dark skin tone",
	},
	"1F9B9-1F3FF-200D-2642-FE0F": {
		Key:        "1F9B9-1F3FF-200D-2642-FE0F",
		Value:      "🦹🏿‍♂️",
		Descriptor: "man supervillain: dark skin tone",
	},
	"1F9B9-200D-2640-FE0F": {
		Key:        "1F9B9-200D-2640-FE0F",
		Value:      "🦹‍♀️",
		Descriptor: "woman supervillain",
	},
	"1F9B9-1F3FB-200D-2640-FE0F": {
		Key:        "1F9B9-1F3FB-200D-2640-FE0F",
		Value:      "🦹🏻‍♀️",
		Descriptor: "woman supervillain: light skin tone",
	},
	"1F9B9-1F3FC-200D-2640-FE0F": {
		Key:        "1F9B9-1F3FC-200D-2640-FE0F",
		Value:      "🦹🏼‍♀️",
		Descriptor: "woman supervillain: medium-light skin tone",
	},
	"1F9B9-1F3FD-200D-2640-FE0F": {
		Key:        "1F9B9-1F3FD-200D-2640-FE0F",
		Value:      "🦹🏽‍♀️",
		Descriptor: "woman supervillain: medium skin tone",
	},
	"1F9B9-1F3FE-200D-2640-FE0F": {
		Key:        "1F9B9-1F3FE-200D-2640-FE0F",
		Value:      "🦹🏾‍♀️",
		Descriptor: "woman supervillain: medium-dark skin tone",
	},
	"1F9B9-1F3FF-200D-2640-FE0F": {
		Key:        "1F9B9-1F3FF-200D-2640-FE0F",
		Value:      "🦹🏿‍♀️",
		Descriptor: "woman supervillain: dark skin tone",
	},
	"1F9D9": {
		Key:        "1F9D9",
		Value:      "🧙",
		Descriptor: "mage",
	},
	"1F9D9-1F3FB": {
		Key:        "1F9D9-1F3FB",
		Value:      "🧙🏻",
		Descriptor: "mage: light skin tone",
	},
	"1F9D9-1F3FC": {
		Key:        "1F9D9-1F3FC",
		Value:      "🧙🏼",
		Descriptor: "mage: medium-light skin tone",
	},
	"1F9D9-1F3FD": {
		Key:        "1F9D9-1F3FD",
		Value:      "🧙🏽",
		Descriptor: "mage: medium skin tone",
	},
	"1F9D9-1F3FE": {
		Key:        "1F9D9-1F3FE",
		Value:      "🧙🏾",
		Descriptor: "mage: medium-dark skin tone",
	},
	"1F9D9-1F3FF": {
		Key:        "1F9D9-1F3FF",
		Value:      "🧙🏿",
		Descriptor: "mage: dark skin tone",
	},
	"1F9D9-200D-2642-FE0F": {
		Key:        "1F9D9-200D-2642-FE0F",
		Value:      "🧙‍♂️",
		Descriptor: "man mage",
	},
	"1F9D9-1F3FB-200D-2642-FE0F": {
		Key:        "1F9D9-1F3FB-200D-2642-FE0F",
		Value:      "🧙🏻‍♂️",
		Descriptor: "man mage: light skin tone",
	},
	"1F9D9-1F3FC-200D-2642-FE0F": {
		Key:        "1F9D9-1F3FC-200D-2642-FE0F",
		Value:      "🧙🏼‍♂️",
		Descriptor: "man mage: medium-light skin tone",
	},
	"1F9D9-1F3FD-200D-2642-FE0F": {
		Key:        "1F9D9-1F3FD-200D-2642-FE0F",
		Value:      "🧙🏽‍♂️",
		Descriptor: "man mage: medium skin tone",
	},
	"1F9D9-1F3FE-200D-2642-FE0F": {
		Key:        "1F9D9-1F3FE-200D-2642-FE0F",
		Value:      "🧙🏾‍♂️",
		Descriptor: "man mage: medium-dark skin tone",
	},
	"1F9D9-1F3FF-200D-2642-FE0F": {
		Key:        "1F9D9-1F3FF-200D-2642-FE0F",
		Value:      "🧙🏿‍♂️",
		Descriptor: "man mage: dark skin tone",
	},
	"1F9D9-200D-2640-FE0F": {
		Key:        "1F9D9-200D-2640-FE0F",
		Value:      "🧙‍♀️",
		Descriptor: "woman mage",
	},
	"1F9D9-1F3FB-200D-2640-FE0F": {
		Key:        "1F9D9-1F3FB-200D-2640-FE0F",
		Value:      "🧙🏻‍♀️",
		Descriptor: "woman mage: light skin tone",
	},
	"1F9D9-1F3FC-200D-2640-FE0F": {
		Key:        "1F9D9-1F3FC-200D-2640-FE0F",
		Value:      "🧙🏼‍♀️",
		Descriptor: "woman mage: medium-light skin tone",
	},
	"1F9D9-1F3FD-200D-2640-FE0F": {
		Key:        "1F9D9-1F3FD-200D-2640-FE0F",
		Value:      "🧙🏽‍♀️",
		Descriptor: "woman mage: medium skin tone",
	},
	"1F9D9-1F3FE-200D-2640-FE0F": {
		Key:        "1F9D9-1F3FE-200D-2640-FE0F",
		Value:      "🧙🏾‍♀️",
		Descriptor: "woman mage: medium-dark skin tone",
	},
	"1F9D9-1F3FF-200D-2640-FE0F": {
		Key:        "1F9D9-1F3FF-200D-2640-FE0F",
		Value:      "🧙🏿‍♀️",
		Descriptor: "woman mage: dark skin tone",
	},
	"1F9DA": {
		Key:        "1F9DA",
		Value:      "🧚",
		Descriptor: "fairy",
	},
	"1F9DA-1F3FB": {
		Key:        "1F9DA-1F3FB",
		Value:      "🧚🏻",
		Descriptor: "fairy: light skin tone",
	},
	"1F9DA-1F3FC": {
		Key:        "1F9DA-1F3FC",
		Value:      "🧚🏼",
		Descriptor: "fairy: medium-light skin tone",
	},
	"1F9DA-1F3FD": {
		Key:        "1F9DA-1F3FD",
		Value:      "🧚🏽",
		Descriptor: "fairy: medium skin tone",
	},
	"1F9DA-1F3FE": {
		Key:        "1F9DA-1F3FE",
		Value:      "🧚🏾",
		Descriptor: "fairy: medium-dark skin tone",
	},
	"1F9DA-1F3FF": {
		Key:        "1F9DA-1F3FF",
		Value:      "🧚🏿",
		Descriptor: "fairy: dark skin tone",
	},
	"1F9DA-200D-2642-FE0F": {
		Key:        "1F9DA-200D-2642-FE0F",
		Value:      "🧚‍♂️",
		Descriptor: "man fairy",
	},
	"1F9DA-1F3FB-200D-2642-FE0F": {
		Key:        "1F9DA-1F3FB-200D-2642-FE0F",
		Value:      "🧚🏻‍♂️",
		Descriptor: "man fairy: light skin tone",
	},
	"1F9DA-1F3FC-200D-2642-FE0F": {
		Key:        "1F9DA-1F3FC-200D-2642-FE0F",
		Value:      "🧚🏼‍♂️",
		Descriptor: "man fairy: medium-light skin tone",
	},
	"1F9DA-1F3FD-200D-2642-FE0F": {
		Key:        "1F9DA-1F3FD-200D-2642-FE0F",
		Value:      "🧚🏽‍♂️",
		Descriptor: "man fairy: medium skin tone",
	},
	"1F9DA-1F3FE-200D-2642-FE0F": {
		Key:        "1F9DA-1F3FE-200D-2642-FE0F",
		Value:      "🧚🏾‍♂️",
		Descriptor: "man fairy: medium-dark skin tone",
	},
	"1F9DA-1F3FF-200D-2642-FE0F": {
		Key:        "1F9DA-1F3FF-200D-2642-FE0F",
		Value:      "🧚🏿‍♂️",
		Descriptor: "man fairy: dark skin tone",
	},
	"1F9DA-200D-2640-FE0F": {
		Key:        "1F9DA-200D-2640-FE0F",
		Value:      "🧚‍♀️",
		Descriptor: "woman fairy",
	},
	"1F9DA-1F3FB-200D-2640-FE0F": {
		Key:        "1F9DA-1F3FB-200D-2640-FE0F",
		Value:      "🧚🏻‍♀️",
		Descriptor: "woman fairy: light skin tone",
	},
	"1F9DA-1F3FC-200D-2640-FE0F": {
		Key:        "1F9DA-1F3FC-200D-2640-FE0F",
		Value:      "🧚🏼‍♀️",
		Descriptor: "woman fairy: medium-light skin tone",
	},
	"1F9DA-1F3FD-200D-2640-FE0F": {
		Key:        "1F9DA-1F3FD-200D-2640-FE0F",
		Value:      "🧚🏽‍♀️",
		Descriptor: "woman fairy: medium skin tone",
	},
	"1F9DA-1F3FE-200D-2640-FE0F": {
		Key:        "1F9DA-1F3FE-200D-2640-FE0F",
		Value:      "🧚🏾‍♀️",
		Descriptor: "woman fairy: medium-dark skin tone",
	},
	"1F9DA-1F3FF-200D-2640-FE0F": {
		Key:        "1F9DA-1F3FF-200D-2640-FE0F",
		Value:      "🧚🏿‍♀️",
		Descriptor: "woman fairy: dark skin tone",
	},
	"1F9DB": {
		Key:        "1F9DB",
		Value:      "🧛",
		Descriptor: "vampire",
	},
	"1F9DB-1F3FB": {
		Key:        "1F9DB-1F3FB",
		Value:      "🧛🏻",
		Descriptor: "vampire: light skin tone",
	},
	"1F9DB-1F3FC": {
		Key:        "1F9DB-1F3FC",
		Value:      "🧛🏼",
		Descriptor: "vampire: medium-light skin tone",
	},
	"1F9DB-1F3FD": {
		Key:        "1F9DB-1F3FD",
		Value:      "🧛🏽",
		Descriptor: "vampire: medium skin tone",
	},
	"1F9DB-1F3FE": {
		Key:        "1F9DB-1F3FE",
		Value:      "🧛🏾",
		Descriptor: "vampire: medium-dark skin tone",
	},
	"1F9DB-1F3FF": {
		Key:        "1F9DB-1F3FF",
		Value:      "🧛🏿",
		Descriptor: "vampire: dark skin tone",
	},
	"1F9DB-200D-2642-FE0F": {
		Key:        "1F9DB-200D-2642-FE0F",
		Value:      "🧛‍♂️",
		Descriptor: "man vampire",
	},
	"1F9DB-1F3FB-200D-2642-FE0F": {
		Key:        "1F9DB-1F3FB-200D-2642-FE0F",
		Value:      "🧛🏻‍♂️",
		Descriptor: "man vampire: light skin tone",
	},
	"1F9DB-1F3FC-200D-2642-FE0F": {
		Key:        "1F9DB-1F3FC-200D-2642-FE0F",
		Value:      "🧛🏼‍♂️",
		Descriptor: "man vampire: medium-light skin tone",
	},
	"1F9DB-1F3FD-200D-2642-FE0F": {
		Key:        "1F9DB-1F3FD-200D-2642-FE0F",
		Value:      "🧛🏽‍♂️",
		Descriptor: "man vampire: medium skin tone",
	},
	"1F9DB-1F3FE-200D-2642-FE0F": {
		Key:        "1F9DB-1F3FE-200D-2642-FE0F",
		Value:      "🧛🏾‍♂️",
		Descriptor: "man vampire: medium-dark skin tone",
	},
	"1F9DB-1F3FF-200D-2642-FE0F": {
		Key:        "1F9DB-1F3FF-200D-2642-FE0F",
		Value:      "🧛🏿‍♂️",
		Descriptor: "man vampire: dark skin tone",
	},
	"1F9DB-200D-2640-FE0F": {
		Key:        "1F9DB-200D-2640-FE0F",
		Value:      "🧛‍♀️",
		Descriptor: "woman vampire",
	},
	"1F9DB-1F3FB-200D-2640-FE0F": {
		Key:        "1F9DB-1F3FB-200D-2640-FE0F",
		Value:      "🧛🏻‍♀️",
		Descriptor: "woman vampire: light skin tone",
	},
	"1F9DB-1F3FC-200D-2640-FE0F": {
		Key:        "1F9DB-1F3FC-200D-2640-FE0F",
		Value:      "🧛🏼‍♀️",
		Descriptor: "woman vampire: medium-light skin tone",
	},
	"1F9DB-1F3FD-200D-2640-FE0F": {
		Key:        "1F9DB-1F3FD-200D-2640-FE0F",
		Value:      "🧛🏽‍♀️",
		Descriptor: "woman vampire: medium skin tone",
	},
	"1F9DB-1F3FE-200D-2640-FE0F": {
		Key:        "1F9DB-1F3FE-200D-2640-FE0F",
		Value:      "🧛🏾‍♀️",
		Descriptor: "woman vampire: medium-dark skin tone",
	},
	"1F9DB-1F3FF-200D-2640-FE0F": {
		Key:        "1F9DB-1F3FF-200D-2640-FE0F",
		Value:      "🧛🏿‍♀️",
		Descriptor: "woman vampire: dark skin tone",
	},
	"1F9DC": {
		Key:        "1F9DC",
		Value:      "🧜",
		Descriptor: "merperson",
	},
	"1F9DC-1F3FB": {
		Key:        "1F9DC-1F3FB",
		Value:      "🧜🏻",
		Descriptor: "merperson: light skin tone",
	},
	"1F9DC-1F3FC": {
		Key:        "1F9DC-1F3FC",
		Value:      "🧜🏼",
		Descriptor: "merperson: medium-light skin tone",
	},
	"1F9DC-1F3FD": {
		Key:        "1F9DC-1F3FD",
		Value:      "🧜🏽",
		Descriptor: "merperson: medium skin tone",
	},
	"1F9DC-1F3FE": {
		Key:        "1F9DC-1F3FE",
		Value:      "🧜🏾",
		Descriptor: "merperson: medium-dark skin tone",
	},
	"1F9DC-1F3FF": {
		Key:        "1F9DC-1F3FF",
		Value:      "🧜🏿",
		Descriptor: "merperson: dark skin tone",
	},
	"1F9DC-200D-2642-FE0F": {
		Key:        "1F9DC-200D-2642-FE0F",
		Value:      "🧜‍♂️",
		Descriptor: "merman",
	},
	"1F9DC-1F3FB-200D-2642-FE0F": {
		Key:        "1F9DC-1F3FB-200D-2642-FE0F",
		Value:      "🧜🏻‍♂️",
		Descriptor: "merman: light skin tone",
	},
	"1F9DC-1F3FC-200D-2642-FE0F": {
		Key:        "1F9DC-1F3FC-200D-2642-FE0F",
		Value:      "🧜🏼‍♂️",
		Descriptor: "merman: medium-light skin tone",
	},
	"1F9DC-1F3FD-200D-2642-FE0F": {
		Key:        "1F9DC-1F3FD-200D-2642-FE0F",
		Value:      "🧜🏽‍♂️",
		Descriptor: "merman: medium skin tone",
	},
	"1F9DC-1F3FE-200D-2642-FE0F": {
		Key:        "1F9DC-1F3FE-200D-2642-FE0F",
		Value:      "🧜🏾‍♂️",
		Descriptor: "merman: medium-dark skin tone",
	},
	"1F9DC-1F3FF-200D-2642-FE0F": {
		Key:        "1F9DC-1F3FF-200D-2642-FE0F",
		Value:      "🧜🏿‍♂️",
		Descriptor: "merman: dark skin tone",
	},
	"1F9DC-200D-2640-FE0F": {
		Key:        "1F9DC-200D-2640-FE0F",
		Value:      "🧜‍♀️",
		Descriptor: "mermaid",
	},
	"1F9DC-1F3FB-200D-2640-FE0F": {
		Key:        "1F9DC-1F3FB-200D-2640-FE0F",
		Value:      "🧜🏻‍♀️",
		Descriptor: "mermaid: light skin tone",
	},
	"1F9DC-1F3FC-200D-2640-FE0F": {
		Key:        "1F9DC-1F3FC-200D-2640-FE0F",
		Value:      "🧜🏼‍♀️",
		Descriptor: "mermaid: medium-light skin tone",
	},
	"1F9DC-1F3FD-200D-2640-FE0F": {
		Key:        "1F9DC-1F3FD-200D-2640-FE0F",
		Value:      "🧜🏽‍♀️",
		Descriptor: "mermaid: medium skin tone",
	},
	"1F9DC-1F3FE-200D-2640-FE0F": {
		Key:        "1F9DC-1F3FE-200D-2640-FE0F",
		Value:      "🧜🏾‍♀️",
		Descriptor: "mermaid: medium-dark skin tone",
	},
	"1F9DC-1F3FF-200D-2640-FE0F": {
		Key:        "1F9DC-1F3FF-200D-2640-FE0F",
		Value:      "🧜🏿‍♀️",
		Descriptor: "mermaid: dark skin tone",
	},
	"1F9DD": {
		Key:        "1F9DD",
		Value:      "🧝",
		Descriptor: "elf",
	},
	"1F9DD-1F3FB": {
		Key:        "1F9DD-1F3FB",
		Value:      "🧝🏻",
		Descriptor: "elf: light skin tone",
	},
	"1F9DD-1F3FC": {
		Key:        "1F9DD-1F3FC",
		Value:      "🧝🏼",
		Descriptor: "elf: medium-light skin tone",
	},
	"1F9DD-1F3FD": {
		Key:        "1F9DD-1F3FD",
		Value:      "🧝🏽",
		Descriptor: "elf: medium skin tone",
	},
	"1F9DD-1F3FE": {
		Key:        "1F9DD-1F3FE",
		Value:      "🧝🏾",
		Descriptor: "elf: medium-dark skin tone",
	},
	"1F9DD-1F3FF": {
		Key:        "1F9DD-1F3FF",
		Value:      "🧝🏿",
		Descriptor: "elf: dark skin tone",
	},
	"1F9DD-200D-2642-FE0F": {
		Key:        "1F9DD-200D-2642-FE0F",
		Value:      "🧝‍♂️",
		Descriptor: "man elf",
	},
	"1F9DD-1F3FB-200D-2642-FE0F": {
		Key:        "1F9DD-1F3FB-200D-2642-FE0F",
		Value:      "🧝🏻‍♂️",
		Descriptor: "man elf: light skin tone",
	},
	"1F9DD-1F3FC-200D-2642-FE0F": {
		Key:        "1F9DD-1F3FC-200D-2642-FE0F",
		Value:      "🧝🏼‍♂️",
		Descriptor: "man elf: medium-light skin tone",
	},
	"1F9DD-1F3FD-200D-2642-FE0F": {
		Key:        "1F9DD-1F3FD-200D-2642-FE0F",
		Value:      "🧝🏽‍♂️",
		Descriptor: "man elf: medium skin tone",
	},
	"1F9DD-1F3FE-200D-2642-FE0F": {
		Key:        "1F9DD-1F3FE-200D-2642-FE0F",
		Value:      "🧝🏾‍♂️",
		Descriptor: "man elf: medium-dark skin tone",
	},
	"1F9DD-1F3FF-200D-2642-FE0F": {
		Key:        "1F9DD-1F3FF-200D-2642-FE0F",
		Value:      "🧝🏿‍♂️",
		Descriptor: "man elf: dark skin tone",
	},
	"1F9DD-200D-2640-FE0F": {
		Key:        "1F9DD-200D-2640-FE0F",
		Value:      "🧝‍♀️",
		Descriptor: "woman elf",
	},
	"1F9DD-1F3FB-200D-2640-FE0F": {
		Key:        "1F9DD-1F3FB-200D-2640-FE0F",
		Value:      "🧝🏻‍♀️",
		Descriptor: "woman elf: light skin tone",
	},
	"1F9DD-1F3FC-200D-2640-FE0F": {
		Key:        "1F9DD-1F3FC-200D-2640-FE0F",
		Value:      "🧝🏼‍♀️",
		Descriptor: "woman elf: medium-light skin tone",
	},
	"1F9DD-1F3FD-200D-2640-FE0F": {
		Key:        "1F9DD-1F3FD-200D-2640-FE0F",
		Value:      "🧝🏽‍♀️",
		Descriptor: "woman elf: medium skin tone",
	},
	"1F9DD-1F3FE-200D-2640-FE0F": {
		Key:        "1F9DD-1F3FE-200D-2640-FE0F",
		Value:      "🧝🏾‍♀️",
		Descriptor: "woman elf: medium-dark skin tone",
	},
	"1F9DD-1F3FF-200D-2640-FE0F": {
		Key:        "1F9DD-1F3FF-200D-2640-FE0F",
		Value:      "🧝🏿‍♀️",
		Descriptor: "woman elf: dark skin tone",
	},
	"1F9DE": {
		Key:        "1F9DE",
		Value:      "🧞",
		Descriptor: "genie",
	},
	"1F9DE-200D-2642-FE0F": {
		Key:        "1F9DE-200D-2642-FE0F",
		Value:      "🧞‍♂️",
		Descriptor: "man genie",
	},
	"1F9DE-200D-2640-FE0F": {
		Key:        "1F9DE-200D-2640-FE0F",
		Value:      "🧞‍♀️",
		Descriptor: "woman genie",
	},
	"1F9DF": {
		Key:        "1F9DF",
		Value:      "🧟",
		Descriptor: "zombie",
	},
	"1F9DF-200D-2642-FE0F": {
		Key:        "1F9DF-200D-2642-FE0F",
		Value:      "🧟‍♂️",
		Descriptor: "man zombie",
	},
	"1F9DF-200D-2640-FE0F": {
		Key:        "1F9DF-200D-2640-FE0F",
		Value:      "🧟‍♀️",
		Descriptor: "woman zombie",
	},
	"1F9CC": {
		Key:        "1F9CC",
		Value:      "🧌",
		Descriptor: "troll",
	},
	"1F486": {
		Key:        "1F486",
		Value:      "💆",
		Descriptor: "person getting massage",
	},
	"1F486-1F3FB": {
		Key:        "1F486-1F3FB",
		Value:      "💆🏻",
		Descriptor: "person getting massage: light skin tone",
	},
	"1F486-1F3FC": {
		Key:        "1F486-1F3FC",
		Value:      "💆🏼",
		Descriptor: "person getting massage: medium-light skin tone",
	},
	"1F486-1F3FD": {
		Key:        "1F486-1F3FD",
		Value:      "💆🏽",
		Descriptor: "person getting massage: medium skin tone",
	},
	"1F486-1F3FE": {
		Key:        "1F486-1F3FE",
		Value:      "💆🏾",
		Descriptor: "person getting massage: medium-dark skin tone",
	},
	"1F486-1F3FF": {
		Key:        "1F486-1F3FF",
		Value:      "💆🏿",
		Descriptor: "person getting massage: dark skin tone",
	},
	"1F486-200D-2642-FE0F": {
		Key:        "1F486-200D-2642-FE0F",
		Value:      "💆‍♂️",
		Descriptor: "man getting massage",
	},
	"1F486-1F3FB-200D-2642-FE0F": {
		Key:        "1F486-1F3FB-200D-2642-FE0F",
		Value:      "💆🏻‍♂️",
		Descriptor: "man getting massage: light skin tone",
	},
	"1F486-1F3FC-200D-2642-FE0F": {
		Key:        "1F486-1F3FC-200D-2642-FE0F",
		Value:      "💆🏼‍♂️",
		Descriptor: "man getting massage: medium-light skin tone",
	},
	"1F486-1F3FD-200D-2642-FE0F": {
		Key:        "1F486-1F3FD-200D-2642-FE0F",
		Value:      "💆🏽‍♂️",
		Descriptor: "man getting massage: medium skin tone",
	},
	"1F486-1F3FE-200D-2642-FE0F": {
		Key:        "1F486-1F3FE-200D-2642-FE0F",
		Value:      "💆🏾‍♂️",
		Descriptor: "man getting massage: medium-dark skin tone",
	},
	"1F486-1F3FF-200D-2642-FE0F": {
		Key:        "1F486-1F3FF-200D-2642-FE0F",
		Value:      "💆🏿‍♂️",
		Descriptor: "man getting massage: dark skin tone",
	},
	"1F486-200D-2640-FE0F": {
		Key:        "1F486-200D-2640-FE0F",
		Value:      "💆‍♀️",
		Descriptor: "woman getting massage",
	},
	"1F486-1F3FB-200D-2640-FE0F": {
		Key:        "1F486-1F3FB-200D-2640-FE0F",
		Value:      "💆🏻‍♀️",
		Descriptor: "woman getting massage: light skin tone",
	},
	"1F486-1F3FC-200D-2640-FE0F": {
		Key:        "1F486-1F3FC-200D-2640-FE0F",
		Value:      "💆🏼‍♀️",
		Descriptor: "woman getting massage: medium-light skin tone",
	},
	"1F486-1F3FD-200D-2640-FE0F": {
		Key:        "1F486-1F3FD-200D-2640-FE0F",
		Value:      "💆🏽‍♀️",
		Descriptor: "woman getting massage: medium skin tone",
	},
	"1F486-1F3FE-200D-2640-FE0F": {
		Key:        "1F486-1F3FE-200D-2640-FE0F",
		Value:      "💆🏾‍♀️",
		Descriptor: "woman getting massage: medium-dark skin tone",
	},
	"1F486-1F3FF-200D-2640-FE0F": {
		Key:        "1F486-1F3FF-200D-2640-FE0F",
		Value:      "💆🏿‍♀️",
		Descriptor: "woman getting massage: dark skin tone",
	},
	"1F487": {
		Key:        "1F487",
		Value:      "💇",
		Descriptor: "person getting haircut",
	},
	"1F487-1F3FB": {
		Key:        "1F487-1F3FB",
		Value:      "💇🏻",
		Descriptor: "person getting haircut: light skin tone",
	},
	"1F487-1F3FC": {
		Key:        "1F487-1F3FC",
		Value:      "💇🏼",
		Descriptor: "person getting haircut: medium-light skin tone",
	},
	"1F487-1F3FD": {
		Key:        "1F487-1F3FD",
		Value:      "💇🏽",
		Descriptor: "person getting haircut: medium skin tone",
	},
	"1F487-1F3FE": {
		Key:        "1F487-1F3FE",
		Value:      "💇🏾",
		Descriptor: "person getting haircut: medium-dark skin tone",
	},
	"1F487-1F3FF": {
		Key:        "1F487-1F3FF",
		Value:      "💇🏿",
		Descriptor: "person getting haircut: dark skin tone",
	},
	"1F487-200D-2642-FE0F": {
		Key:        "1F487-200D-2642-FE0F",
		Value:      "💇‍♂️",
		Descriptor: "man getting haircut",
	},
	"1F487-1F3FB-200D-2642-FE0F": {
		Key:        "1F487-1F3FB-200D-2642-FE0F",
		Value:      "💇🏻‍♂️",
		Descriptor: "man getting haircut: light skin tone",
	},
	"1F487-1F3FC-200D-2642-FE0F": {
		Key:        "1F487-1F3FC-200D-2642-FE0F",
		Value:      "💇🏼‍♂️",
		Descriptor: "man getting haircut: medium-light skin tone",
	},
	"1F487-1F3FD-200D-2642-FE0F": {
		Key:        "1F487-1F3FD-200D-2642-FE0F",
		Value:      "💇🏽‍♂️",
		Descriptor: "man getting haircut: medium skin tone",
	},
	"1F487-1F3FE-200D-2642-FE0F": {
		Key:        "1F487-1F3FE-200D-2642-FE0F",
		Value:      "💇🏾‍♂️",
		Descriptor: "man getting haircut: medium-dark skin tone",
	},
	"1F487-1F3FF-200D-2642-FE0F": {
		Key:        "1F487-1F3FF-200D-2642-FE0F",
		Value:      "💇🏿‍♂️",
		Descriptor: "man getting haircut: dark skin tone",
	},
	"1F487-200D-2640-FE0F": {
		Key:        "1F487-200D-2640-FE0F",
		Value:      "💇‍♀️",
		Descriptor: "woman getting haircut",
	},
	"1F487-1F3FB-200D-2640-FE0F": {
		Key:        "1F487-1F3FB-200D-2640-FE0F",
		Value:      "💇🏻‍♀️",
		Descriptor: "woman getting haircut: light skin tone",
	},
	"1F487-1F3FC-200D-2640-FE0F": {
		Key:        "1F487-1F3FC-200D-2640-FE0F",
		Value:      "💇🏼‍♀️",
		Descriptor: "woman getting haircut: medium-light skin tone",
	},
	"1F487-1F3FD-200D-2640-FE0F": {
		Key:        "1F487-1F3FD-200D-2640-FE0F",
		Value:      "💇🏽‍♀️",
		Descriptor: "woman getting haircut: medium skin tone",
	},
	"1F487-1F3FE-200D-2640-FE0F": {
		Key:        "1F487-1F3FE-200D-2640-FE0F",
		Value:      "💇🏾‍♀️",
		Descriptor: "woman getting haircut: medium-dark skin tone",
	},
	"1F487-1F3FF-200D-2640-FE0F": {
		Key:        "1F487-1F3FF-200D-2640-FE0F",
		Value:      "💇🏿‍♀️",
		Descriptor: "woman getting haircut: dark skin tone",
	},
	"1F6B6": {
		Key:        "1F6B6",
		Value:      "🚶",
		Descriptor: "person walking",
	},
	"1F6B6-1F3FB": {
		Key:        "1F6B6-1F3FB",
		Value:      "🚶🏻",
		Descriptor: "person walking: light skin tone",
	},
	"1F6B6-1F3FC": {
		Key:        "1F6B6-1F3FC",
		Value:      "🚶🏼",
		Descriptor: "person walking: medium-light skin tone",
	},
	"1F6B6-1F3FD": {
		Key:        "1F6B6-1F3FD",
		Value:      "🚶🏽",
		Descriptor: "person walking: medium skin tone",
	},
	"1F6B6-1F3FE": {
		Key:        "1F6B6-1F3FE",
		Value:      "🚶🏾",
		Descriptor: "person walking: medium-dark skin tone",
	},
	"1F6B6-1F3FF": {
		Key:        "1F6B6-1F3FF",
		Value:      "🚶🏿",
		Descriptor: "person walking: dark skin tone",
	},
	"1F6B6-200D-2642-FE0F": {
		Key:        "1F6B6-200D-2642-FE0F",
		Value:      "🚶‍♂️",
		Descriptor: "man walking",
	},
	"1F6B6-1F3FB-200D-2642-FE0F": {
		Key:        "1F6B6-1F3FB-200D-2642-FE0F",
		Value:      "🚶🏻‍♂️",
		Descriptor: "man walking: light skin tone",
	},
	"1F6B6-1F3FC-200D-2642-FE0F": {
		Key:        "1F6B6-1F3FC-200D-2642-FE0F",
		Value:      "🚶🏼‍♂️",
		Descriptor: "man walking: medium-light skin tone",
	},
	"1F6B6-1F3FD-200D-2642-FE0F": {
		Key:        "1F6B6-1F3FD-200D-2642-FE0F",
		Value:      "🚶🏽‍♂️",
		Descriptor: "man walking: medium skin tone",
	},
	"1F6B6-1F3FE-200D-2642-FE0F": {
		Key:        "1F6B6-1F3FE-200D-2642-FE0F",
		Value:      "🚶🏾‍♂️",
		Descriptor: "man walking: medium-dark skin tone",
	},
	"1F6B6-1F3FF-200D-2642-FE0F": {
		Key:        "1F6B6-1F3FF-200D-2642-FE0F",
		Value:      "🚶🏿‍♂️",
		Descriptor: "man walking: dark skin tone",
	},
	"1F6B6-200D-2640-FE0F": {
		Key:        "1F6B6-200D-2640-FE0F",
		Value:      "🚶‍♀️",
		Descriptor: "woman walking",
	},
	"1F6B6-1F3FB-200D-2640-FE0F": {
		Key:        "1F6B6-1F3FB-200D-2640-FE0F",
		Value:      "🚶🏻‍♀️",
		Descriptor: "woman walking: light skin tone",
	},
	"1F6B6-1F3FC-200D-2640-FE0F": {
		Key:        "1F6B6-1F3FC-200D-2640-FE0F",
		Value:      "🚶🏼‍♀️",
		Descriptor: "woman walking: medium-light skin tone",
	},
	"1F6B6-1F3FD-200D-2640-FE0F": {
		Key:        "1F6B6-1F3FD-200D-2640-FE0F",
		Value:      "🚶🏽‍♀️",
		Descriptor: "woman walking: medium skin tone",
	},
	"1F6B6-1F3FE-200D-2640-FE0F": {
		Key:        "1F6B6-1F3FE-200D-2640-FE0F",
		Value:      "🚶🏾‍♀️",
		Descriptor: "woman walking: medium-dark skin tone",
	},
	"1F6B6-1F3FF-200D-2640-FE0F": {
		Key:        "1F6B6-1F3FF-200D-2640-FE0F",
		Value:      "🚶🏿‍♀️",
		Descriptor: "woman walking: dark skin tone",
	},
	"1F6B6-200D-27A1-FE0F": {
		Key:        "1F6B6-200D-27A1-FE0F",
		Value:      "🚶‍➡️",
		Descriptor: "person walking facing right",
	},
	"1F6B6-1F3FB-200D-27A1-FE0F": {
		Key:        "1F6B6-1F3FB-200D-27A1-FE0F",
		Value:      "🚶🏻‍➡️",
		Descriptor: "person walking facing right: light skin tone",
	},
	"1F6B6-1F3FC-200D-27A1-FE0F": {
		Key:        "1F6B6-1F3FC-200D-27A1-FE0F",
		Value:      "🚶🏼‍➡️",
		Descriptor: "person walking facing right: medium-light skin tone",
	},
	"1F6B6-1F3FD-200D-27A1-FE0F": {
		Key:        "1F6B6-1F3FD-200D-27A1-FE0F",
		Value:      "🚶🏽‍➡️",
		Descriptor: "person walking facing right: medium skin tone",
	},
	"1F6B6-1F3FE-200D-27A1-FE0F": {
		Key:        "1F6B6-1F3FE-200D-27A1-FE0F",
		Value:      "🚶🏾‍➡️",
		Descriptor: "person walking facing right: medium-dark skin tone",
	},
	"1F6B6-1F3FF-200D-27A1-FE0F": {
		Key:        "1F6B6-1F3FF-200D-27A1-FE0F",
		Value:      "🚶🏿‍➡️",
		Descriptor: "person walking facing right: dark skin tone",
	},
	"1F6B6-200D-2640-FE0F-200D-27A1-FE0F": {
		Key:        "1F6B6-200D-2640-FE0F-200D-27A1-FE0F",
		Value:      "🚶‍♀️‍➡️",
		Descriptor: "woman walking facing right",
	},
	"1F6B6-1F3FB-200D-2640-FE0F-200D-27A1-FE0F": {
		Key:        "1F6B6-1F3FB-200D-2640-FE0F-200D-27A1-FE0F",
		Value:      "🚶🏻‍♀️‍➡️",
		Descriptor: "woman walking facing right: light skin tone",
	},
	"1F6B6-1F3FC-200D-2640-FE0F-200D-27A1-FE0F": {
		Key:        "1F6B6-1F3FC-200D-2640-FE0F-200D-27A1-FE0F",
		Value:      "🚶🏼‍♀️‍➡️",
		Descriptor: "woman walking facing right: medium-light skin tone",
	},
	"1F6B6-1F3FD-200D-2640-FE0F-200D-27A1-FE0F": {
		Key:        "1F6B6-1F3FD-200D-2640-FE0F-200D-27A1-FE0F",
		Value:      "🚶🏽‍♀️‍➡️",
		Descriptor: "woman walking facing right: medium skin tone",
	},
	"1F6B6-1F3FE-200D-2640-FE0F-200D-27A1-FE0F": {
		Key:        "1F6B6-1F3FE-200D-2640-FE0F-200D-27A1-FE0F",
		Value:      "🚶🏾‍♀️‍➡️",
		Descriptor: "woman walking facing right: medium-dark skin tone",
	},
	"1F6B6-1F3FF-200D-2640-FE0F-200D-27A1-FE0F": {
		Key:        "1F6B6-1F3FF-200D-2640-FE0F-200D-27A1-FE0F",
		Value:      "🚶🏿‍♀️‍➡️",
		Descriptor: "woman walking facing right: dark skin tone",
	},
	"1F6B6-200D-2642-FE0F-200D-27A1-FE0F": {
		Key:        "1F6B6-200D-2642-FE0F-200D-27A1-FE0F",
		Value:      "🚶‍♂️‍➡️",
		Descriptor: "man walking facing right",
	},
	"1F6B6-1F3FB-200D-2642-FE0F-200D-27A1-FE0F": {
		Key:        "1F6B6-1F3FB-200D-2642-FE0F-200D-27A1-FE0F",
		Value:      "🚶🏻‍♂️‍➡️",
		Descriptor: "man walking facing right: light skin tone",
	},
	"1F6B6-1F3FC-200D-2642-FE0F-200D-27A1-FE0F": {
		Key:        "1F6B6-1F3FC-200D-2642-FE0F-200D-27A1-FE0F",
		Value:      "🚶🏼‍♂️‍➡️",
		Descriptor: "man walking facing right: medium-light skin tone",
	},
	"1F6B6-1F3FD-200D-2642-FE0F-200D-27A1-FE0F": {
		Key:        "1F6B6-1F3FD-200D-2642-FE0F-200D-27A1-FE0F",
		Value:      "🚶🏽‍♂️‍➡️",
		Descriptor: "man walking facing right: medium skin tone",
	},
	"1F6B6-1F3FE-200D-2642-FE0F-200D-27A1-FE0F": {
		Key:        "1F6B6-1F3FE-200D-2642-FE0F-200D-27A1-FE0F",
		Value:      "🚶🏾‍♂️‍➡️",
		Descriptor: "man walking facing right: medium-dark skin tone",
	},
	"1F6B6-1F3FF-200D-2642-FE0F-200D-27A1-FE0F": {
		Key:        "1F6B6-1F3FF-200D-2642-FE0F-200D-27A1-FE0F",
		Value:      "🚶🏿‍♂️‍➡️",
		Descriptor: "man walking facing right: dark skin tone",
	},
	"1F9CD": {
		Key:        "1F9CD",
		Value:      "🧍",
		Descriptor: "person standing",
	},
	"1F9CD-1F3FB": {
		Key:        "1F9CD-1F3FB",
		Value:      "🧍🏻",
		Descriptor: "person standing: light skin tone",
	},
	"1F9CD-1F3FC": {
		Key:        "1F9CD-1F3FC",
		Value:      "🧍🏼",
		Descriptor: "person standing: medium-light skin tone",
	},
	"1F9CD-1F3FD": {
		Key:        "1F9CD-1F3FD",
		Value:      "🧍🏽",
		Descriptor: "person standing: medium skin tone",
	},
	"1F9CD-1F3FE": {
		Key:        "1F9CD-1F3FE",
		Value:      "🧍🏾",
		Descriptor: "person standing: medium-dark skin tone",
	},
	"1F9CD-1F3FF": {
		Key:        "1F9CD-1F3FF",
		Value:      "🧍🏿",
		Descriptor: "person standing: dark skin tone",
	},
	"1F9CD-200D-2642-FE0F": {
		Key:        "1F9CD-200D-2642-FE0F",
		Value:      "🧍‍♂️",
		Descriptor: "man standing",
	},
	"1F9CD-1F3FB-200D-2642-FE0F": {
		Key:        "1F9CD-1F3FB-200D-2642-FE0F",
		Value:      "🧍🏻‍♂️",
		Descriptor: "man standing: light skin tone",
	},
	"1F9CD-1F3FC-200D-2642-FE0F": {
		Key:        "1F9CD-1F3FC-200D-2642-FE0F",
		Value:      "🧍🏼‍♂️",
		Descriptor: "man standing: medium-light skin tone",
	},
	"1F9CD-1F3FD-200D-2642-FE0F": {
		Key:        "1F9CD-1F3FD-200D-2642-FE0F",
		Value:      "🧍🏽‍♂️",
		Descriptor: "man standing: medium skin tone",
	},
	"1F9CD-1F3FE-200D-2642-FE0F": {
		Key:        "1F9CD-1F3FE-200D-2642-FE0F",
		Value:      "🧍🏾‍♂️",
		Descriptor: "man standing: medium-dark skin tone",
	},
	"1F9CD-1F3FF-200D-2642-FE0F": {
		Key:        "1F9CD-1F3FF-200D-2642-FE0F",
		Value:      "🧍🏿‍♂️",
		Descriptor: "man standing: dark skin tone",
	},
	"1F9CD-200D-2640-FE0F": {
		Key:        "1F9CD-200D-2640-FE0F",
		Value:      "🧍‍♀️",
		Descriptor: "woman standing",
	},
	"1F9CD-1F3FB-200D-2640-FE0F": {
		Key:        "1F9CD-1F3FB-200D-2640-FE0F",
		Value:      "🧍🏻‍♀️",
		Descriptor: "woman standing: light skin tone",
	},
	"1F9CD-1F3FC-200D-2640-FE0F": {
		Key:        "1F9CD-1F3FC-200D-2640-FE0F",
		Value:      "🧍🏼‍♀️",
		Descriptor: "woman standing: medium-light skin tone",
	},
	"1F9CD-1F3FD-200D-2640-FE0F": {
		Key:        "1F9CD-1F3FD-200D-2640-FE0F",
		Value:      "🧍🏽‍♀️",
		Descriptor: "woman standing: medium skin tone",
	},
	"1F9CD-1F3FE-200D-2640-FE0F": {
		Key:        "1F9CD-1F3FE-200D-2640-FE0F",
		Value:      "🧍🏾‍♀️",
		Descriptor: "woman standing: medium-dark skin tone",
	},
	"1F9CD-1F3FF-200D-2640-FE0F": {
		Key:        "1F9CD-1F3FF-200D-2640-FE0F",
		Value:      "🧍🏿‍♀️",
		Descriptor: "woman standing: dark skin tone",
	},
	"1F9CE": {
		Key:        "1F9CE",
		Value:      "🧎",
		Descriptor: "person kneeling",
	},
	"1F9CE-1F3FB": {
		Key:        "1F9CE-1F3FB",
		Value:      "🧎🏻",
		Descriptor: "person kneeling: light skin tone",
	},
	"1F9CE-1F3FC": {
		Key:        "1F9CE-1F3FC",
		Value:      "🧎🏼",
		Descriptor: "person kneeling: medium-light skin tone",
	},
	"1F9CE-1F3FD": {
		Key:        "1F9CE-1F3FD",
		Value:      "🧎🏽",
		Descriptor: "person kneeling: medium skin tone",
	},
	"1F9CE-1F3FE": {
		Key:        "1F9CE-1F3FE",
		Value:      "🧎🏾",
		Descriptor: "person kneeling: medium-dark skin tone",
	},
	"1F9CE-1F3FF": {
		Key:        "1F9CE-1F3FF",
		Value:      "🧎🏿",
		Descriptor: "person kneeling: dark skin tone",
	},
	"1F9CE-200D-2642-FE0F": {
		Key:        "1F9CE-200D-2642-FE0F",
		Value:      "🧎‍♂️",
		Descriptor: "man kneeling",
	},
	"1F9CE-1F3FB-200D-2642-FE0F": {
		Key:        "1F9CE-1F3FB-200D-2642-FE0F",
		Value:      "🧎🏻‍♂️",
		Descriptor: "man kneeling: light skin tone",
	},
	"1F9CE-1F3FC-200D-2642-FE0F": {
		Key:        "1F9CE-1F3FC-200D-2642-FE0F",
		Value:      "🧎🏼‍♂️",
		Descriptor: "man kneeling: medium-light skin tone",
	},
	"1F9CE-1F3FD-200D-2642-FE0F": {
		Key:        "1F9CE-1F3FD-200D-2642-FE0F",
		Value:      "🧎🏽‍♂️",
		Descriptor: "man kneeling: medium skin tone",
	},
	"1F9CE-1F3FE-200D-2642-FE0F": {
		Key:        "1F9CE-1F3FE-200D-2642-FE0F",
		Value:      "🧎🏾‍♂️",
		Descriptor: "man kneeling: medium-dark skin tone",
	},
	"1F9CE-1F3FF-200D-2642-FE0F": {
		Key:        "1F9CE-1F3FF-200D-2642-FE0F",
		Value:      "🧎🏿‍♂️",
		Descriptor: "man kneeling: dark skin tone",
	},
	"1F9CE-200D-2640-FE0F": {
		Key:        "1F9CE-200D-2640-FE0F",
		Value:      "🧎‍♀️",
		Descriptor: "woman kneeling",
	},
	"1F9CE-1F3FB-200D-2640-FE0F": {
		Key:        "1F9CE-1F3FB-200D-2640-FE0F",
		Value:      "🧎🏻‍♀️",
		Descriptor: "woman kneeling: light skin tone",
	},
	"1F9CE-1F3FC-200D-2640-FE0F": {
		Key:        "1F9CE-1F3FC-200D-2640-FE0F",
		Value:      "🧎🏼‍♀️",
		Descriptor: "woman kneeling: medium-light skin tone",
	},
	"1F9CE-1F3FD-200D-2640-FE0F": {
		Key:        "1F9CE-1F3FD-200D-2640-FE0F",
		Value:      "🧎🏽‍♀️",
		Descriptor: "woman kneeling: medium skin tone",
	},
	"1F9CE-1F3FE-200D-2640-FE0F": {
		Key:        "1F9CE-1F3FE-200D-2640-FE0F",
		Value:      "🧎🏾‍♀️",
		Descriptor: "woman kneeling: medium-dark skin tone",
	},
	"1F9CE-1F3FF-200D-2640-FE0F": {
		Key:        "1F9CE-1F3FF-200D-2640-FE0F",
		Value:      "🧎🏿‍♀️",
		Descriptor: "woman kneeling: dark skin tone",
	},
	"1F9CE-200D-27A1-FE0F": {
		Key:        "1F9CE-200D-27A1-FE0F",
		Value:      "🧎‍➡️",
		Descriptor: "person kneeling facing right",
	},
	"1F9CE-1F3FB-200D-27A1-FE0F": {
		Key:        "1F9CE-1F3FB-200D-27A1-FE0F",
		Value:      "🧎🏻‍➡️",
		Descriptor: "person kneeling facing right: light skin tone",
	},
	"1F9CE-1F3FC-200D-27A1-FE0F": {
		Key:        "1F9CE-1F3FC-200D-27A1-FE0F",
		Value:      "🧎🏼‍➡️",
		Descriptor: "person kneeling facing right: medium-light skin tone",
	},
	"1F9CE-1F3FD-200D-27A1-FE0F": {
		Key:        "1F9CE-1F3FD-200D-27A1-FE0F",
		Value:      "🧎🏽‍➡️",
		Descriptor: "person kneeling facing right: medium skin tone",
	},
	"1F9CE-1F3FE-200D-27A1-FE0F": {
		Key:        "1F9CE-1F3FE-200D-27A1-FE0F",
		Value:      "🧎🏾‍➡️",
		Descriptor: "person kneeling facing right: medium-dark skin tone",
	},
	"1F9CE-1F3FF-200D-27A1-FE0F": {
		Key:        "1F9CE-1F3FF-200D-27A1-FE0F",
		Value:      "🧎🏿‍➡️",
		Descriptor: "person kneeling facing right: dark skin tone",
	},
	"1F9CE-200D-2640-FE0F-200D-27A1-FE0F": {
		Key:        "1F9CE-200D-2640-FE0F-200D-27A1-FE0F",
		Value:      "🧎‍♀️‍➡️",
		Descriptor: "woman kneeling facing right",
	},
	"1F9CE-1F3FB-200D-2640-FE0F-200D-27A1-FE0F": {
		Key:        "1F9CE-1F3FB-200D-2640-FE0F-200D-27A1-FE0F",
		Value:      "🧎🏻‍♀️‍➡️",
		Descriptor: "woman kneeling facing right: light skin tone",
	},
	"1F9CE-1F3FC-200D-2640-FE0F-200D-27A1-FE0F": {
		Key:        "1F9CE-1F3FC-200D-2640-FE0F-200D-27A1-FE0F",
		Value:      "🧎🏼‍♀️‍➡️",
		Descriptor: "woman kneeling facing right: medium-light skin tone",
	},
	"1F9CE-1F3FD-200D-2640-FE0F-200D-27A1-FE0F": {
		Key:        "1F9CE-1F3FD-200D-2640-FE0F-200D-27A1-FE0F",
		Value:      "🧎🏽‍♀️‍➡️",
		Descriptor: "woman kneeling facing right: medium skin tone",
	},
	"1F9CE-1F3FE-200D-2640-FE0F-200D-27A1-FE0F": {
		Key:        "1F9CE-1F3FE-200D-2640-FE0F-200D-27A1-FE0F",
		Value:      "🧎🏾‍♀️‍➡️",
		Descriptor: "woman kneeling facing right: medium-dark skin tone",
	},
	"1F9CE-1F3FF-200D-2640-FE0F-200D-27A1-FE0F": {
		Key:        "1F9CE-1F3FF-200D-2640-FE0F-200D-27A1-FE0F",
		Value:      "🧎🏿‍♀️‍➡️",
		Descriptor: "woman kneeling facing right: dark skin tone",
	},
	"1F9CE-200D-2642-FE0F-200D-27A1-FE0F": {
		Key:        "1F9CE-200D-2642-FE0F-200D-27A1-FE0F",
		Value:      "🧎‍♂️‍➡️",
		Descriptor: "man kneeling facing right",
	},
	"1F9CE-1F3FB-200D-2642-FE0F-200D-27A1-FE0F": {
		Key:        "1F9CE-1F3FB-200D-2642-FE0F-200D-27A1-FE0F",
		Value:      "🧎🏻‍♂️‍➡️",
		Descriptor: "man kneeling facing right: light skin tone",
	},
	"1F9CE-1F3FC-200D-2642-FE0F-200D-27A1-FE0F": {
		Key:        "1F9CE-1F3FC-200D-2642-FE0F-200D-27A1-FE0F",
		Value:      "🧎🏼‍♂️‍➡️",
		Descriptor: "man kneeling facing right: medium-light skin tone",
	},
	"1F9CE-1F3FD-200D-2642-FE0F-200D-27A1-FE0F": {
		Key:        "1F9CE-1F3FD-200D-2642-FE0F-200D-27A1-FE0F",
		Value:      "🧎🏽‍♂️‍➡️",
		Descriptor: "man kneeling facing right: medium skin tone",
	},
	"1F9CE-1F3FE-200D-2642-FE0F-200D-27A1-FE0F": {
		Key:        "1F9CE-1F3FE-200D-2642-FE0F-200D-27A1-FE0F",
		Value:      "🧎🏾‍♂️‍➡️",
		Descriptor: "man kneeling facing right: medium-dark skin tone",
	},
	"1F9CE-1F3FF-200D-2642-FE0F-200D-27A1-FE0F": {
		Key:        "1F9CE-1F3FF-200D-2642-FE0F-200D-27A1-FE0F",
		Value:      "🧎🏿‍♂️‍➡️",
		Descriptor: "man kneeling facing right: dark skin tone",
	},
	"1F9D1-200D-1F9AF": {
		Key:        "1F9D1-200D-1F9AF",
		Value:      "🧑‍🦯",
		Descriptor: "person with white cane",
	},
	"1F9D1-1F3FB-200D-1F9AF": {
		Key:        "1F9D1-1F3FB-200D-1F9AF",
		Value:      "🧑🏻‍🦯",
		Descriptor: "person with white cane: light skin tone",
	},
	"1F9D1-1F3FC-200D-1F9AF": {
		Key:        "1F9D1-1F3FC-200D-1F9AF",
		Value:      "🧑🏼‍🦯",
		Descriptor: "person with white cane: medium-light skin tone",
	},
	"1F9D1-1F3FD-200D-1F9AF": {
		Key:        "1F9D1-1F3FD-200D-1F9AF",
		Value:      "🧑🏽‍🦯",
		Descriptor: "person with white cane: medium skin tone",
	},
	"1F9D1-1F3FE-200D-1F9AF": {
		Key:        "1F9D1-1F3FE-200D-1F9AF",
		Value:      "🧑🏾‍🦯",
		Descriptor: "person with white cane: medium-dark skin tone",
	},
	"1F9D1-1F3FF-200D-1F9AF": {
		Key:        "1F9D1-1F3FF-200D-1F9AF",
		Value:      "🧑🏿‍🦯",
		Descriptor: "person with white cane: dark skin tone",
	},
	"1F9D1-200D-1F9AF-200D-27A1-FE0F": {
		Key:        "1F9D1-200D-1F9AF-200D-27A1-FE0F",
		Value:      "🧑‍🦯‍➡️",
		Descriptor: "person with white cane facing right",
	},
	"1F9D1-1F3FB-200D-1F9AF-200D-27A1-FE0F": {
		Key:        "1F9D1-1F3FB-200D-1F9AF-200D-27A1-FE0F",
		Value:      "🧑🏻‍🦯‍➡️",
		Descriptor: "person with white cane facing right: light skin tone",
	},
	"1F9D1-1F3FC-200D-1F9AF-200D-27A1-FE0F": {
		Key:        "1F9D1-1F3FC-200D-1F9AF-200D-27A1-FE0F",
		Value:      "🧑🏼‍🦯‍➡️",
		Descriptor: "person with white cane facing right: medium-light skin tone",
	},
	"1F9D1-1F3FD-200D-1F9AF-200D-27A1-FE0F": {
		Key:        "1F9D1-1F3FD-200D-1F9AF-200D-27A1-FE0F",
		Value:      "🧑🏽‍🦯‍➡️",
		Descriptor: "person with white cane facing right: medium skin tone",
	},
	"1F9D1-1F3FE-200D-1F9AF-200D-27A1-FE0F": {
		Key:        "1F9D1-1F3FE-200D-1F9AF-200D-27A1-FE0F",
		Value:      "🧑🏾‍🦯‍➡️",
		Descriptor: "person with white cane facing right: medium-dark skin tone",
	},
	"1F9D1-1F3FF-200D-1F9AF-200D-27A1-FE0F": {
		Key:        "1F9D1-1F3FF-200D-1F9AF-200D-27A1-FE0F",
		Value:      "🧑🏿‍🦯‍➡️",
		Descriptor: "person with white cane facing right: dark skin tone",
	},
	"1F468-200D-1F9AF": {
		Key:        "1F468-200D-1F9AF",
		Value:      "👨‍🦯",
		Descriptor: "man with white cane",
	},
	"1F468-1F3FB-200D-1F9AF": {
		Key:        "1F468-1F3FB-200D-1F9AF",
		Value:      "👨🏻‍🦯",
		Descriptor: "man with white cane: light skin tone",
	},
	"1F468-1F3FC-200D-1F9AF": {
		Key:        "1F468-1F3FC-200D-1F9AF",
		Value:      "👨🏼‍🦯",
		Descriptor: "man with white cane: medium-light skin tone",
	},
	"1F468-1F3FD-200D-1F9AF": {
		Key:        "1F468-1F3FD-200D-1F9AF",
		Value:      "👨🏽‍🦯",
		Descriptor: "man with white cane: medium skin tone",
	},
	"1F468-1F3FE-200D-1F9AF": {
		Key:        "1F468-1F3FE-200D-1F9AF",
		Value:      "👨🏾‍🦯",
		Descriptor: "man with white cane: medium-dark skin tone",
	},
	"1F468-1F3FF-200D-1F9AF": {
		Key:        "1F468-1F3FF-200D-1F9AF",
		Value:      "👨🏿‍🦯",
		Descriptor: "man with white cane: dark skin tone",
	},
	"1F468-200D-1F9AF-200D-27A1-FE0F": {
		Key:        "1F468-200D-1F9AF-200D-27A1-FE0F",
		Value:      "👨‍🦯‍➡️",
		Descriptor: "man with white cane facing right",
	},
	"1F468-1F3FB-200D-1F9AF-200D-27A1-FE0F": {
		Key:        "1F468-1F3FB-200D-1F9AF-200D-27A1-FE0F",
		Value:      "👨🏻‍🦯‍➡️",
		Descriptor: "man with white cane facing right: light skin tone",
	},
	"1F468-1F3FC-200D-1F9AF-200D-27A1-FE0F": {
		Key:        "1F468-1F3FC-200D-1F9AF-200D-27A1-FE0F",
		Value:      "👨🏼‍🦯‍➡️",
		Descriptor: "man with white cane facing right: medium-light skin tone",
	},
	"1F468-1F3FD-200D-1F9AF-200D-27A1-FE0F": {
		Key:        "1F468-1F3FD-200D-1F9AF-200D-27A1-FE0F",
		Value:      "👨🏽‍🦯‍➡️",
		Descriptor: "man with white cane facing right: medium skin tone",
	},
	"1F468-1F3FE-200D-1F9AF-200D-27A1-FE0F": {
		Key:        "1F468-1F3FE-200D-1F9AF-200D-27A1-FE0F",
		Value:      "👨🏾‍🦯‍➡️",
		Descriptor: "man with white cane facing right: medium-dark skin tone",
	},
	"1F468-1F3FF-200D-1F9AF-200D-27A1-FE0F": {
		Key:        "1F468-1F3FF-200D-1F9AF-200D-27A1-FE0F",
		Value:      "👨🏿‍🦯‍➡️",
		Descriptor: "man with white cane facing right: dark skin tone",
	},
	"1F469-200D-1F9AF": {
		Key:        "1F469-200D-1F9AF",
		Value:      "👩‍🦯",
		Descriptor: "woman with white cane",
	},
	"1F469-1F3FB-200D-1F9AF": {
		Key:        "1F469-1F3FB-200D-1F9AF",
		Value:      "👩🏻‍🦯",
		Descriptor: "woman with white cane: light skin tone",
	},
	"1F469-1F3FC-200D-1F9AF": {
		Key:        "1F469-1F3FC-200D-1F9AF",
		Value:      "👩🏼‍🦯",
		Descriptor: "woman with white cane: medium-light skin tone",
	},
	"1F469-1F3FD-200D-1F9AF": {
		Key:        "1F469-1F3FD-200D-1F9AF",
		Value:      "👩🏽‍🦯",
		Descriptor: "woman with white cane: medium skin tone",
	},
	"1F469-1F3FE-200D-1F9AF": {
		Key:        "1F469-1F3FE-200D-1F9AF",
		Value:      "👩🏾‍🦯",
		Descriptor: "woman with white cane: medium-dark skin tone",
	},
	"1F469-1F3FF-200D-1F9AF": {
		Key:        "1F469-1F3FF-200D-1F9AF",
		Value:      "👩🏿‍🦯",
		Descriptor: "woman with white cane: dark skin tone",
	},
	"1F469-200D-1F9AF-200D-27A1-FE0F": {
		Key:        "1F469-200D-1F9AF-200D-27A1-FE0F",
		Value:      "👩‍🦯‍➡️",
		Descriptor: "woman with white cane facing right",
	},
	"1F469-1F3FB-200D-1F9AF-200D-27A1-FE0F": {
		Key:        "1F469-1F3FB-200D-1F9AF-200D-27A1-FE0F",
		Value:      "👩🏻‍🦯‍➡️",
		Descriptor: "woman with white cane facing right: light skin tone",
	},
	"1F469-1F3FC-200D-1F9AF-200D-27A1-FE0F": {
		Key:        "1F469-1F3FC-200D-1F9AF-200D-27A1-FE0F",
		Value:      "👩🏼‍🦯‍➡️",
		Descriptor: "woman with white cane facing right: medium-light skin tone",
	},
	"1F469-1F3FD-200D-1F9AF-200D-27A1-FE0F": {
		Key:        "1F469-1F3FD-200D-1F9AF-200D-27A1-FE0F",
		Value:      "👩🏽‍🦯‍➡️",
		Descriptor: "woman with white cane facing right: medium skin tone",
	},
	"1F469-1F3FE-200D-1F9AF-200D-27A1-FE0F": {
		Key:        "1F469-1F3FE-200D-1F9AF-200D-27A1-FE0F",
		Value:      "👩🏾‍🦯‍➡️",
		Descriptor: "woman with white cane facing right: medium-dark skin tone",
	},
	"1F469-1F3FF-200D-1F9AF-200D-27A1-FE0F": {
		Key:        "1F469-1F3FF-200D-1F9AF-200D-27A1-FE0F",
		Value:      "👩🏿‍🦯‍➡️",
		Descriptor: "woman with white cane facing right: dark skin tone",
	},
	"1F9D1-200D-1F9BC": {
		Key:        "1F9D1-200D-1F9BC",
		Value:      "🧑‍🦼",
		Descriptor: "person in motorized wheelchair",
	},
	"1F9D1-1F3FB-200D-1F9BC": {
		Key:        "1F9D1-1F3FB-200D-1F9BC",
		Value:      "🧑🏻‍🦼",
		Descriptor: "person in motorized wheelchair: light skin tone",
	},
	"1F9D1-1F3FC-200D-1F9BC": {
		Key:        "1F9D1-1F3FC-200D-1F9BC",
		Value:      "🧑🏼‍🦼",
		Descriptor: "person in motorized wheelchair: medium-light skin tone",
	},
	"1F9D1-1F3FD-200D-1F9BC": {
		Key:        "1F9D1-1F3FD-200D-1F9BC",
		Value:      "🧑🏽‍🦼",
		Descriptor: "person in motorized wheelchair: medium skin tone",
	},
	"1F9D1-1F3FE-200D-1F9BC": {
		Key:        "1F9D1-1F3FE-200D-1F9BC",
		Value:      "🧑🏾‍🦼",
		Descriptor: "person in motorized wheelchair: medium-dark skin tone",
	},
	"1F9D1-1F3FF-200D-1F9BC": {
		Key:        "1F9D1-1F3FF-200D-1F9BC",
		Value:      "🧑🏿‍🦼",
		Descriptor: "person in motorized wheelchair: dark skin tone",
	},
	"1F9D1-200D-1F9BC-200D-27A1-FE0F": {
		Key:        "1F9D1-200D-1F9BC-200D-27A1-FE0F",
		Value:      "🧑‍🦼‍➡️",
		Descriptor: "person in motorized wheelchair facing right",
	},
	"1F9D1-1F3FB-200D-1F9BC-200D-27A1-FE0F": {
		Key:        "1F9D1-1F3FB-200D-1F9BC-200D-27A1-FE0F",
		Value:      "🧑🏻‍🦼‍➡️",
		Descriptor: "person in motorized wheelchair facing right: light skin tone",
	},
	"1F9D1-1F3FC-200D-1F9BC-200D-27A1-FE0F": {
		Key:        "1F9D1-1F3FC-200D-1F9BC-200D-27A1-FE0F",
		Value:      "🧑🏼‍🦼‍➡️",
		Descriptor: "person in motorized wheelchair facing right: medium-light skin tone",
	},
	"1F9D1-1F3FD-200D-1F9BC-200D-27A1-FE0F": {
		Key:        "1F9D1-1F3FD-200D-1F9BC-200D-27A1-FE0F",
		Value:      "🧑🏽‍🦼‍➡️",
		Descriptor: "person in motorized wheelchair facing right: medium skin tone",
	},
	"1F9D1-1F3FE-200D-1F9BC-200D-27A1-FE0F": {
		Key:        "1F9D1-1F3FE-200D-1F9BC-200D-27A1-FE0F",
		Value:      "🧑🏾‍🦼‍➡️",
		Descriptor: "person in motorized wheelchair facing right: medium-dark skin tone",
	},
	"1F9D1-1F3FF-200D-1F9BC-200D-27A1-FE0F": {
		Key:        "1F9D1-1F3FF-200D-1F9BC-200D-27A1-FE0F",
		Value:      "🧑🏿‍🦼‍➡️",
		Descriptor: "person in motorized wheelchair facing right: dark skin tone",
	},
	"1F468-200D-1F9BC": {
		Key:        "1F468-200D-1F9BC",
		Value:      "👨‍🦼",
		Descriptor: "man in motorized wheelchair",
	},
	"1F468-1F3FB-200D-1F9BC": {
		Key:        "1F468-1F3FB-200D-1F9BC",
		Value:      "👨🏻‍🦼",
		Descriptor: "man in motorized wheelchair: light skin tone",
	},
	"1F468-1F3FC-200D-1F9BC": {
		Key:        "1F468-1F3FC-200D-1F9BC",
		Value:      "👨🏼‍🦼",
		Descriptor: "man in motorized wheelchair: medium-light skin tone",
	},
	"1F468-1F3FD-200D-1F9BC": {
		Key:        "1F468-1F3FD-200D-1F9BC",
		Value:      "👨🏽‍🦼",
		Descriptor: "man in motorized wheelchair: medium skin tone",
	},
	"1F468-1F3FE-200D-1F9BC": {
		Key:        "1F468-1F3FE-200D-1F9BC",
		Value:      "👨🏾‍🦼",
		Descriptor: "man in motorized wheelchair: medium-dark skin tone",
	},
	"1F468-1F3FF-200D-1F9BC": {
		Key:        "1F468-1F3FF-200D-1F9BC",
		Value:      "👨🏿‍🦼",
		Descriptor: "man in motorized wheelchair: dark skin tone",
	},
	"1F468-200D-1F9BC-200D-27A1-FE0F": {
		Key:        "1F468-200D-1F9BC-200D-27A1-FE0F",
		Value:      "👨‍🦼‍➡️",
		Descriptor: "man in motorized wheelchair facing right",
	},
	"1F468-1F3FB-200D-1F9BC-200D-27A1-FE0F": {
		Key:        "1F468-1F3FB-200D-1F9BC-200D-27A1-FE0F",
		Value:      "👨🏻‍🦼‍➡️",
		Descriptor: "man in motorized wheelchair facing right: light skin tone",
	},
	"1F468-1F3FC-200D-1F9BC-200D-27A1-FE0F": {
		Key:        "1F468-1F3FC-200D-1F9BC-200D-27A1-FE0F",
		Value:      "👨🏼‍🦼‍➡️",
		Descriptor: "man in motorized wheelchair facing right: medium-light skin tone",
	},
	"1F468-1F3FD-200D-1F9BC-200D-27A1-FE0F": {
		Key:        "1F468-1F3FD-200D-1F9BC-200D-27A1-FE0F",
		Value:      "👨🏽‍🦼‍➡️",
		Descriptor: "man in motorized wheelchair facing right: medium skin tone",
	},
	"1F468-1F3FE-200D-1F9BC-200D-27A1-FE0F": {
		Key:        "1F468-1F3FE-200D-1F9BC-200D-27A1-FE0F",
		Value:      "👨🏾‍🦼‍➡️",
		Descriptor: "man in motorized wheelchair facing right: medium-dark skin tone",
	},
	"1F468-1F3FF-200D-1F9BC-200D-27A1-FE0F": {
		Key:        "1F468-1F3FF-200D-1F9BC-200D-27A1-FE0F",
		Value:      "👨🏿‍🦼‍➡️",
		Descriptor: "man in motorized wheelchair facing right: dark skin tone",
	},
	"1F469-200D-1F9BC": {
		Key:        "1F469-200D-1F9BC",
		Value:      "👩‍🦼",
		Descriptor: "woman in motorized wheelchair",
	},
	"1F469-1F3FB-200D-1F9BC": {
		Key:        "1F469-1F3FB-200D-1F9BC",
		Value:      "👩🏻‍🦼",
		Descriptor: "woman in motorized wheelchair: light skin tone",
	},
	"1F469-1F3FC-200D-1F9BC": {
		Key:        "1F469-1F3FC-200D-1F9BC",
		Value:      "👩🏼‍🦼",
		Descriptor: "woman in motorized wheelchair: medium-light skin tone",
	},
	"1F469-1F3FD-200D-1F9BC": {
		Key:        "1F469-1F3FD-200D-1F9BC",
		Value:      "👩🏽‍🦼",
		Descriptor: "woman in motorized wheelchair: medium skin tone",
	},
	"1F469-1F3FE-200D-1F9BC": {
		Key:        "1F469-1F3FE-200D-1F9BC",
		Value:      "👩🏾‍🦼",
		Descriptor: "woman in motorized wheelchair: medium-dark skin tone",
	},
	"1F469-1F3FF-200D-1F9BC": {
		Key:        "1F469-1F3FF-200D-1F9BC",
		Value:      "👩🏿‍🦼",
		Descriptor: "woman in motorized wheelchair: dark skin tone",
	},
	"1F469-200D-1F9BC-200D-27A1-FE0F": {
		Key:        "1F469-200D-1F9BC-200D-27A1-FE0F",
		Value:      "👩‍🦼‍➡️",
		Descriptor: "woman in motorized wheelchair facing right",
	},
	"1F469-1F3FB-200D-1F9BC-200D-27A1-FE0F": {
		Key:        "1F469-1F3FB-200D-1F9BC-200D-27A1-FE0F",
		Value:      "👩🏻‍🦼‍➡️",
		Descriptor: "woman in motorized wheelchair facing right: light skin tone",
	},
	"1F469-1F3FC-200D-1F9BC-200D-27A1-FE0F": {
		Key:        "1F469-1F3FC-200D-1F9BC-200D-27A1-FE0F",
		Value:      "👩🏼‍🦼‍➡️",
		Descriptor: "woman in motorized wheelchair facing right: medium-light skin tone",
	},
	"1F469-1F3FD-200D-1F9BC-200D-27A1-FE0F": {
		Key:        "1F469-1F3FD-200D-1F9BC-200D-27A1-FE0F",
		Value:      "👩🏽‍🦼‍➡️",
		Descriptor: "woman in motorized wheelchair facing right: medium skin tone",
	},
	"1F469-1F3FE-200D-1F9BC-200D-27A1-FE0F": {
		Key:        "1F469-1F3FE-200D-1F9BC-200D-27A1-FE0F",
		Value:      "👩🏾‍🦼‍➡️",
		Descriptor: "woman in motorized wheelchair facing right: medium-dark skin tone",
	},
	"1F469-1F3FF-200D-1F9BC-200D-27A1-FE0F": {
		Key:        "1F469-1F3FF-200D-1F9BC-200D-27A1-FE0F",
		Value:      "👩🏿‍🦼‍➡️",
		Descriptor: "woman in motorized wheelchair facing right: dark skin tone",
	},
	"1F9D1-200D-1F9BD": {
		Key:        "1F9D1-200D-1F9BD",
		Value:      "🧑‍🦽",
		Descriptor: "person in manual wheelchair",
	},
	"1F9D1-1F3FB-200D-1F9BD": {
		Key:        "1F9D1-1F3FB-200D-1F9BD",
		Value:      "🧑🏻‍🦽",
		Descriptor: "person in manual wheelchair: light skin tone",
	},
	"1F9D1-1F3FC-200D-1F9BD": {
		Key:        "1F9D1-1F3FC-200D-1F9BD",
		Value:      "🧑🏼‍🦽",
		Descriptor: "person in manual wheelchair: medium-light skin tone",
	},
	"1F9D1-1F3FD-200D-1F9BD": {
		Key:        "1F9D1-1F3FD-200D-1F9BD",
		Value:      "🧑🏽‍🦽",
		Descriptor: "person in manual wheelchair: medium skin tone",
	},
	"1F9D1-1F3FE-200D-1F9BD": {
		Key:        "1F9D1-1F3FE-200D-1F9BD",
		Value:      "🧑🏾‍🦽",
		Descriptor: "person in manual wheelchair: medium-dark skin tone",
	},
	"1F9D1-1F3FF-200D-1F9BD": {
		Key:        "1F9D1-1F3FF-200D-1F9BD",
		Value:      "🧑🏿‍🦽",
		Descriptor: "person in manual wheelchair: dark skin tone",
	},
	"1F9D1-200D-1F9BD-200D-27A1-FE0F": {
		Key:        "1F9D1-200D-1F9BD-200D-27A1-FE0F",
		Value:      "🧑‍🦽‍➡️",
		Descriptor: "person in manual wheelchair facing right",
	},
	"1F9D1-1F3FB-200D-1F9BD-200D-27A1-FE0F": {
		Key:        "1F9D1-1F3FB-200D-1F9BD-200D-27A1-FE0F",
		Value:      "🧑🏻‍🦽‍➡️",
		Descriptor: "person in manual wheelchair facing right: light skin tone",
	},
	"1F9D1-1F3FC-200D-1F9BD-200D-27A1-FE0F": {
		Key:        "1F9D1-1F3FC-200D-1F9BD-200D-27A1-FE0F",
		Value:      "🧑🏼‍🦽‍➡️",
		Descriptor: "person in manual wheelchair facing right: medium-light skin tone",
	},
	"1F9D1-1F3FD-200D-1F9BD-200D-27A1-FE0F": {
		Key:        "1F9D1-1F3FD-200D-1F9BD-200D-27A1-FE0F",
		Value:      "🧑🏽‍🦽‍➡️",
		Descriptor: "person in manual wheelchair facing right: medium skin tone",
	},
	"1F9D1-1F3FE-200D-1F9BD-200D-27A1-FE0F": {
		Key:        "1F9D1-1F3FE-200D-1F9BD-200D-27A1-FE0F",
		Value:      "🧑🏾‍🦽‍➡️",
		Descriptor: "person in manual wheelchair facing right: medium-dark skin tone",
	},
	"1F9D1-1F3FF-200D-1F9BD-200D-27A1-FE0F": {
		Key:        "1F9D1-1F3FF-200D-1F9BD-200D-27A1-FE0F",
		Value:      "🧑🏿‍🦽‍➡️",
		Descriptor: "person in manual wheelchair facing right: dark skin tone",
	},
	"1F468-200D-1F9BD": {
		Key:        "1F468-200D-1F9BD",
		Value:      "👨‍🦽",
		Descriptor: "man in manual wheelchair",
	},
	"1F468-1F3FB-200D-1F9BD": {
		Key:        "1F468-1F3FB-200D-1F9BD",
		Value:      "👨🏻‍🦽",
		Descriptor: "man in manual wheelchair: light skin tone",
	},
	"1F468-1F3FC-200D-1F9BD": {
		Key:        "1F468-1F3FC-200D-1F9BD",
		Value:      "👨🏼‍🦽",
		Descriptor: "man in manual wheelchair: medium-light skin tone",
	},
	"1F468-1F3FD-200D-1F9BD": {
		Key:        "1F468-1F3FD-200D-1F9BD",
		Value:      "👨🏽‍🦽",
		Descriptor: "man in manual wheelchair: medium skin tone",
	},
	"1F468-1F3FE-200D-1F9BD": {
		Key:        "1F468-1F3FE-200D-1F9BD",
		Value:      "👨🏾‍🦽",
		Descriptor: "man in manual wheelchair: medium-dark skin tone",
	},
	"1F468-1F3FF-200D-1F9BD": {
		Key:        "1F468-1F3FF-200D-1F9BD",
		Value:      "👨🏿‍🦽",
		Descriptor: "man in manual wheelchair: dark skin tone",
	},
	"1F468-200D-1F9BD-200D-27A1-FE0F": {
		Key:        "1F468-200D-1F9BD-200D-27A1-FE0F",
		Value:      "👨‍🦽‍➡️",
		Descriptor: "man in manual wheelchair facing right",
	},
	"1F468-1F3FB-200D-1F9BD-200D-27A1-FE0F": {
		Key:        "1F468-1F3FB-200D-1F9BD-200D-27A1-FE0F",
		Value:      "👨🏻‍🦽‍➡️",
		Descriptor: "man in manual wheelchair facing right: light skin tone",
	},
	"1F468-1F3FC-200D-1F9BD-200D-27A1-FE0F": {
		Key:        "1F468-1F3FC-200D-1F9BD-200D-27A1-FE0F",
		Value:      "👨🏼‍🦽‍➡️",
		Descriptor: "man in manual wheelchair facing right: medium-light skin tone",
	},
	"1F468-1F3FD-200D-1F9BD-200D-27A1-FE0F": {
		Key:        "1F468-1F3FD-200D-1F9BD-200D-27A1-FE0F",
		Value:      "👨🏽‍🦽‍➡️",
		Descriptor: "man in manual wheelchair facing right: medium skin tone",
	},
	"1F468-1F3FE-200D-1F9BD-200D-27A1-FE0F": {
		Key:        "1F468-1F3FE-200D-1F9BD-200D-27A1-FE0F",
		Value:      "👨🏾‍🦽‍➡️",
		Descriptor: "man in manual wheelchair facing right: medium-dark skin tone",
	},
	"1F468-1F3FF-200D-1F9BD-200D-27A1-FE0F": {
		Key:        "1F468-1F3FF-200D-1F9BD-200D-27A1-FE0F",
		Value:      "👨🏿‍🦽‍➡️",
		Descriptor: "man in manual wheelchair facing right: dark skin tone",
	},
	"1F469-200D-1F9BD": {
		Key:        "1F469-200D-1F9BD",
		Value:      "👩‍🦽",
		Descriptor: "woman in manual wheelchair",
	},
	"1F469-1F3FB-200D-1F9BD": {
		Key:        "1F469-1F3FB-200D-1F9BD",
		Value:      "👩🏻‍🦽",
		Descriptor: "woman in manual wheelchair: light skin tone",
	},
	"1F469-1F3FC-200D-1F9BD": {
		Key:        "1F469-1F3FC-200D-1F9BD",
		Value:      "👩🏼‍🦽",
		Descriptor: "woman in manual wheelchair: medium-light skin tone",
	},
	"1F469-1F3FD-200D-1F9BD": {
		Key:        "1F469-1F3FD-200D-1F9BD",
		Value:      "👩🏽‍🦽",
		Descriptor: "woman in manual wheelchair: medium skin tone",
	},
	"1F469-1F3FE-200D-1F9BD": {
		Key:        "1F469-1F3FE-200D-1F9BD",
		Value:      "👩🏾‍🦽",
		Descriptor: "woman in manual wheelchair: medium-dark skin tone",
	},
	"1F469-1F3FF-200D-1F9BD": {
		Key:        "1F469-1F3FF-200D-1F9BD",
		Value:      "👩🏿‍🦽",
		Descriptor: "woman in manual wheelchair: dark skin tone",
	},
	"1F469-200D-1F9BD-200D-27A1-FE0F": {
		Key:        "1F469-200D-1F9BD-200D-27A1-FE0F",
		Value:      "👩‍🦽‍➡️",
		Descriptor: "woman in manual wheelchair facing right",
	},
	"1F469-1F3FB-200D-1F9BD-200D-27A1-FE0F": {
		Key:        "1F469-1F3FB-200D-1F9BD-200D-27A1-FE0F",
		Value:      "👩🏻‍🦽‍➡️",
		Descriptor: "woman in manual wheelchair facing right: light skin tone",
	},
	"1F469-1F3FC-200D-1F9BD-200D-27A1-FE0F": {
		Key:        "1F469-1F3FC-200D-1F9BD-200D-27A1-FE0F",
		Value:      "👩🏼‍🦽‍➡️",
		Descriptor: "woman in manual wheelchair facing right: medium-light skin tone",
	},
	"1F469-1F3FD-200D-1F9BD-200D-27A1-FE0F": {
		Key:        "1F469-1F3FD-200D-1F9BD-200D-27A1-FE0F",
		Value:      "👩🏽‍🦽‍➡️",
		Descriptor: "woman in manual wheelchair facing right: medium skin tone",
	},
	"1F469-1F3FE-200D-1F9BD-200D-27A1-FE0F": {
		Key:        "1F469-1F3FE-200D-1F9BD-200D-27A1-FE0F",
		Value:      "👩🏾‍🦽‍➡️",
		Descriptor: "woman in manual wheelchair facing right: medium-dark skin tone",
	},
	"1F469-1F3FF-200D-1F9BD-200D-27A1-FE0F": {
		Key:        "1F469-1F3FF-200D-1F9BD-200D-27A1-FE0F",
		Value:      "👩🏿‍🦽‍➡️",
		Descriptor: "woman in manual wheelchair facing right: dark skin tone",
	},
	"1F3C3": {
		Key:        "1F3C3",
		Value:      "🏃",
		Descriptor: "person running",
	},
	"1F3C3-1F3FB": {
		Key:        "1F3C3-1F3FB",
		Value:      "🏃🏻",
		Descriptor: "person running: light skin tone",
	},
	"1F3C3-1F3FC": {
		Key:        "1F3C3-1F3FC",
		Value:      "🏃🏼",
		Descriptor: "person running: medium-light skin tone",
	},
	"1F3C3-1F3FD": {
		Key:        "1F3C3-1F3FD",
		Value:      "🏃🏽",
		Descriptor: "person running: medium skin tone",
	},
	"1F3C3-1F3FE": {
		Key:        "1F3C3-1F3FE",
		Value:      "🏃🏾",
		Descriptor: "person running: medium-dark skin tone",
	},
	"1F3C3-1F3FF": {
		Key:        "1F3C3-1F3FF",
		Value:      "🏃🏿",
		Descriptor: "person running: dark skin tone",
	},
	"1F3C3-200D-2642-FE0F": {
		Key:        "1F3C3-200D-2642-FE0F",
		Value:      "🏃‍♂️",
		Descriptor: "man running",
	},
	"1F3C3-1F3FB-200D-2642-FE0F": {
		Key:        "1F3C3-1F3FB-200D-2642-FE0F",
		Value:      "🏃🏻‍♂️",
		Descriptor: "man running: light skin tone",
	},
	"1F3C3-1F3FC-200D-2642-FE0F": {
		Key:        "1F3C3-1F3FC-200D-2642-FE0F",
		Value:      "🏃🏼‍♂️",
		Descriptor: "man running: medium-light skin tone",
	},
	"1F3C3-1F3FD-200D-2642-FE0F": {
		Key:        "1F3C3-1F3FD-200D-2642-FE0F",
		Value:      "🏃🏽‍♂️",
		Descriptor: "man running: medium skin tone",
	},
	"1F3C3-1F3FE-200D-2642-FE0F": {
		Key:        "1F3C3-1F3FE-200D-2642-FE0F",
		Value:      "🏃🏾‍♂️",
		Descriptor: "man running: medium-dark skin tone",
	},
	"1F3C3-1F3FF-200D-2642-FE0F": {
		Key:        "1F3C3-1F3FF-200D-2642-FE0F",
		Value:      "🏃🏿‍♂️",
		Descriptor: "man running: dark skin tone",
	},
	"1F3C3-200D-2640-FE0F": {
		Key:        "1F3C3-200D-2640-FE0F",
		Value:      "🏃‍♀️",
		Descriptor: "woman running",
	},
	"1F3C3-1F3FB-200D-2640-FE0F": {
		Key:        "1F3C3-1F3FB-200D-2640-FE0F",
		Value:      "🏃🏻‍♀️",
		Descriptor: "woman running: light skin tone",
	},
	"1F3C3-1F3FC-200D-2640-FE0F": {
		Key:        "1F3C3-1F3FC-200D-2640-FE0F",
		Value:      "🏃🏼‍♀️",
		Descriptor: "woman running: medium-light skin tone",
	},
	"1F3C3-1F3FD-200D-2640-FE0F": {
		Key:        "1F3C3-1F3FD-200D-2640-FE0F",
		Value:      "🏃🏽‍♀️",
		Descriptor: "woman running: medium skin tone",
	},
	"1F3C3-1F3FE-200D-2640-FE0F": {
		Key:        "1F3C3-1F3FE-200D-2640-FE0F",
		Value:      "🏃🏾‍♀️",
		Descriptor: "woman running: medium-dark skin tone",
	},
	"1F3C3-1F3FF-200D-2640-FE0F": {
		Key:        "1F3C3-1F3FF-200D-2640-FE0F",
		Value:      "🏃🏿‍♀️",
		Descriptor: "woman running: dark skin tone",
	},
	"1F3C3-200D-27A1-FE0F": {
		Key:        "1F3C3-200D-27A1-FE0F",
		Value:      "🏃‍➡️",
		Descriptor: "person running facing right",
	},
	"1F3C3-1F3FB-200D-27A1-FE0F": {
		Key:        "1F3C3-1F3FB-200D-27A1-FE0F",
		Value:      "🏃🏻‍➡️",
		Descriptor: "person running facing right: light skin tone",
	},
	"1F3C3-1F3FC-200D-27A1-FE0F": {
		Key:        "1F3C3-1F3FC-200D-27A1-FE0F",
		Value:      "🏃🏼‍➡️",
		Descriptor: "person running facing right: medium-light skin tone",
	},
	"1F3C3-1F3FD-200D-27A1-FE0F": {
		Key:        "1F3C3-1F3FD-200D-27A1-FE0F",
		Value:      "🏃🏽‍➡️",
		Descriptor: "person running facing right: medium skin tone",
	},
	"1F3C3-1F3FE-200D-27A1-FE0F": {
		Key:        "1F3C3-1F3FE-200D-27A1-FE0F",
		Value:      "🏃🏾‍➡️",
		Descriptor: "person running facing right: medium-dark skin tone",
	},
	"1F3C3-1F3FF-200D-27A1-FE0F": {
		Key:        "1F3C3-1F3FF-200D-27A1-FE0F",
		Value:      "🏃🏿‍➡️",
		Descriptor: "person running facing right: dark skin tone",
	},
	"1F3C3-200D-2640-FE0F-200D-27A1-FE0F": {
		Key:        "1F3C3-200D-2640-FE0F-200D-27A1-FE0F",
		Value:      "🏃‍♀️‍➡️",
		Descriptor: "woman running facing right",
	},
	"1F3C3-1F3FB-200D-2640-FE0F-200D-27A1-FE0F": {
		Key:        "1F3C3-1F3FB-200D-2640-FE0F-200D-27A1-FE0F",
		Value:      "🏃🏻‍♀️‍➡️",
		Descriptor: "woman running facing right: light skin tone",
	},
	"1F3C3-1F3FC-200D-2640-FE0F-200D-27A1-FE0F": {
		Key:        "1F3C3-1F3FC-200D-2640-FE0F-200D-27A1-FE0F",
		Value:      "🏃🏼‍♀️‍➡️",
		Descriptor: "woman running facing right: medium-light skin tone",
	},
	"1F3C3-1F3FD-200D-2640-FE0F-200D-27A1-FE0F": {
		Key:        "1F3C3-1F3FD-200D-2640-FE0F-200D-27A1-FE0F",
		Value:      "🏃🏽‍♀️‍➡️",
		Descriptor: "woman running facing right: medium skin tone",
	},
	"1F3C3-1F3FE-200D-2640-FE0F-200D-27A1-FE0F": {
		Key:        "1F3C3-1F3FE-200D-2640-FE0F-200D-27A1-FE0F",
		Value:      "🏃🏾‍♀️‍➡️",
		Descriptor: "woman running facing right: medium-dark skin tone",
	},
	"1F3C3-1F3FF-200D-2640-FE0F-200D-27A1-FE0F": {
		Key:        "1F3C3-1F3FF-200D-2640-FE0F-200D-27A1-FE0F",
		Value:      "🏃🏿‍♀️‍➡️",
		Descriptor: "woman running facing right: dark skin tone",
	},
	"1F3C3-200D-2642-FE0F-200D-27A1-FE0F": {
		Key:        "1F3C3-200D-2642-FE0F-200D-27A1-FE0F",
		Value:      "🏃‍♂️‍➡️",
		Descriptor: "man running facing right",
	},
	"1F3C3-1F3FB-200D-2642-FE0F-200D-27A1-FE0F": {
		Key:        "1F3C3-1F3FB-200D-2642-FE0F-200D-27A1-FE0F",
		Value:      "🏃🏻‍♂️‍➡️",
		Descriptor: "man running facing right: light skin tone",
	},
	"1F3C3-1F3FC-200D-2642-FE0F-200D-27A1-FE0F": {
		Key:        "1F3C3-1F3FC-200D-2642-FE0F-200D-27A1-FE0F",
		Value:      "🏃🏼‍♂️‍➡️",
		Descriptor: "man running facing right: medium-light skin tone",
	},
	"1F3C3-1F3FD-200D-2642-FE0F-200D-27A1-FE0F": {
		Key:        "1F3C3-1F3FD-200D-2642-FE0F-200D-27A1-FE0F",
		Value:      "🏃🏽‍♂️‍➡️",
		Descriptor: "man running facing right: medium skin tone",
	},
	"1F3C3-1F3FE-200D-2642-FE0F-200D-27A1-FE0F": {
		Key:        "1F3C3-1F3FE-200D-2642-FE0F-200D-27A1-FE0F",
		Value:      "🏃🏾‍♂️‍➡️",
		Descriptor: "man running facing right: medium-dark skin tone",
	},
	"1F3C3-1F3FF-200D-2642-FE0F-200D-27A1-FE0F": {
		Key:        "1F3C3-1F3FF-200D-2642-FE0F-200D-27A1-FE0F",
		Value:      "🏃🏿‍♂️‍➡️",
		Descriptor: "man running facing right: dark skin tone",
	},
	"1F483": {
		Key:        "1F483",
		Value:      "💃",
		Descriptor: "woman dancing",
	},
	"1F483-1F3FB": {
		Key:        "1F483-1F3FB",
		Value:      "💃🏻",
		Descriptor: "woman dancing: light skin tone",
	},
	"1F483-1F3FC": {
		Key:        "1F483-1F3FC",
		Value:      "💃🏼",
		Descriptor: "woman dancing: medium-light skin tone",
	},
	"1F483-1F3FD": {
		Key:        "1F483-1F3FD",
		Value:      "💃🏽",
		Descriptor: "woman dancing: medium skin tone",
	},
	"1F483-1F3FE": {
		Key:        "1F483-1F3FE",
		Value:      "💃🏾",
		Descriptor: "woman dancing: medium-dark skin tone",
	},
	"1F483-1F3FF": {
		Key:        "1F483-1F3FF",
		Value:      "💃🏿",
		Descriptor: "woman dancing: dark skin tone",
	},
	"1F57A": {
		Key:        "1F57A",
		Value:      "🕺",
		Descriptor: "man dancing",
	},
	"1F57A-1F3FB": {
		Key:        "1F57A-1F3FB",
		Value:      "🕺🏻",
		Descriptor: "man dancing: light skin tone",
	},
	"1F57A-1F3FC": {
		Key:        "1F57A-1F3FC",
		Value:      "🕺🏼",
		Descriptor: "man dancing: medium-light skin tone",
	},
	"1F57A-1F3FD": {
		Key:        "1F57A-1F3FD",
		Value:      "🕺🏽",
		Descriptor: "man dancing: medium skin tone",
	},
	"1F57A-1F3FE": {
		Key:        "1F57A-1F3FE",
		Value:      "🕺🏾",
		Descriptor: "man dancing: medium-dark skin tone",
	},
	"1F57A-1F3FF": {
		Key:        "1F57A-1F3FF",
		Value:      "🕺🏿",
		Descriptor: "man dancing: dark skin tone",
	},
	"1F574-FE0F": {
		Key:        "1F574-FE0F",
		Value:      "🕴️",
		Descriptor: "person in suit levitating",
	},
	"1F574-1F3FB": {
		Key:        "1F574-1F3FB",
		Value:      "🕴🏻",
		Descriptor: "person in suit levitating: light skin tone",
	},
	"1F574-1F3FC": {
		Key:        "1F574-1F3FC",
		Value:      "🕴🏼",
		Descriptor: "person in suit levitating: medium-light skin tone",
	},
	"1F574-1F3FD": {
		Key:        "1F574-1F3FD",
		Value:      "🕴🏽",
		Descriptor: "person in suit levitating: medium skin tone",
	},
	"1F574-1F3FE": {
		Key:        "1F574-1F3FE",
		Value:      "🕴🏾",
		Descriptor: "person in suit levitating: medium-dark skin tone",
	},
	"1F574-1F3FF": {
		Key:        "1F574-1F3FF",
		Value:      "🕴🏿",
		Descriptor: "person in suit levitating: dark skin tone",
	},
	"1F46F": {
		Key:        "1F46F",
		Value:      "👯",
		Descriptor: "people with bunny ears",
	},
	"1F46F-200D-2642-FE0F": {
		Key:        "1F46F-200D-2642-FE0F",
		Value:      "👯‍♂️",
		Descriptor: "men with bunny ears",
	},
	"1F46F-200D-2640-FE0F": {
		Key:        "1F46F-200D-2640-FE0F",
		Value:      "👯‍♀️",
		Descriptor: "women with bunny ears",
	},
	"1F9D6": {
		Key:        "1F9D6",
		Value:      "🧖",
		Descriptor: "person in steamy room",
	},
	"1F9D6-1F3FB": {
		Key:        "1F9D6-1F3FB",
		Value:      "🧖🏻",
		Descriptor: "person in steamy room: light skin tone",
	},
	"1F9D6-1F3FC": {
		Key:        "1F9D6-1F3FC",
		Value:      "🧖🏼",
		Descriptor: "person in steamy room: medium-light skin tone",
	},
	"1F9D6-1F3FD": {
		Key:        "1F9D6-1F3FD",
		Value:      "🧖🏽",
		Descriptor: "person in steamy room: medium skin tone",
	},
	"1F9D6-1F3FE": {
		Key:        "1F9D6-1F3FE",
		Value:      "🧖🏾",
		Descriptor: "person in steamy room: medium-dark skin tone",
	},
	"1F9D6-1F3FF": {
		Key:        "1F9D6-1F3FF",
		Value:      "🧖🏿",
		Descriptor: "person in steamy room: dark skin tone",
	},
	"1F9D6-200D-2642-FE0F": {
		Key:        "1F9D6-200D-2642-FE0F",
		Value:      "🧖‍♂️",
		Descriptor: "man in steamy room",
	},
	"1F9D6-1F3FB-200D-2642-FE0F": {
		Key:        "1F9D6-1F3FB-200D-2642-FE0F",
		Value:      "🧖🏻‍♂️",
		Descriptor: "man in steamy room: light skin tone",
	},
	"1F9D6-1F3FC-200D-2642-FE0F": {
		Key:        "1F9D6-1F3FC-200D-2642-FE0F",
		Value:      "🧖🏼‍♂️",
		Descriptor: "man in steamy room: medium-light skin tone",
	},
	"1F9D6-1F3FD-200D-2642-FE0F": {
		Key:        "1F9D6-1F3FD-200D-2642-FE0F",
		Value:      "🧖🏽‍♂️",
		Descriptor: "man in steamy room: medium skin tone",
	},
	"1F9D6-1F3FE-200D-2642-FE0F": {
		Key:        "1F9D6-1F3FE-200D-2642-FE0F",
		Value:      "🧖🏾‍♂️",
		Descriptor: "man in steamy room: medium-dark skin tone",
	},
	"1F9D6-1F3FF-200D-2642-FE0F": {
		Key:        "1F9D6-1F3FF-200D-2642-FE0F",
		Value:      "🧖🏿‍♂️",
		Descriptor: "man in steamy room: dark skin tone",
	},
	"1F9D6-200D-2640-FE0F": {
		Key:        "1F9D6-200D-2640-FE0F",
		Value:      "🧖‍♀️",
		Descriptor: "woman in steamy room",
	},
	"1F9D6-1F3FB-200D-2640-FE0F": {
		Key:        "1F9D6-1F3FB-200D-2640-FE0F",
		Value:      "🧖🏻‍♀️",
		Descriptor: "woman in steamy room: light skin tone",
	},
	"1F9D6-1F3FC-200D-2640-FE0F": {
		Key:        "1F9D6-1F3FC-200D-2640-FE0F",
		Value:      "🧖🏼‍♀️",
		Descriptor: "woman in steamy room: medium-light skin tone",
	},
	"1F9D6-1F3FD-200D-2640-FE0F": {
		Key:        "1F9D6-1F3FD-200D-2640-FE0F",
		Value:      "🧖🏽‍♀️",
		Descriptor: "woman in steamy room: medium skin tone",
	},
	"1F9D6-1F3FE-200D-2640-FE0F": {
		Key:        "1F9D6-1F3FE-200D-2640-FE0F",
		Value:      "🧖🏾‍♀️",
		Descriptor: "woman in steamy room: medium-dark skin tone",
	},
	"1F9D6-1F3FF-200D-2640-FE0F": {
		Key:        "1F9D6-1F3FF-200D-2640-FE0F",
		Value:      "🧖🏿‍♀️",
		Descriptor: "woman in steamy room: dark skin tone",
	},
	"1F9D7": {
		Key:        "1F9D7",
		Value:      "🧗",
		Descriptor: "person climbing",
	},
	"1F9D7-1F3FB": {
		Key:        "1F9D7-1F3FB",
		Value:      "🧗🏻",
		Descriptor: "person climbing: light skin tone",
	},
	"1F9D7-1F3FC": {
		Key:        "1F9D7-1F3FC",
		Value:      "🧗🏼",
		Descriptor: "person climbing: medium-light skin tone",
	},
	"1F9D7-1F3FD": {
		Key:        "1F9D7-1F3FD",
		Value:      "🧗🏽",
		Descriptor: "person climbing: medium skin tone",
	},
	"1F9D7-1F3FE": {
		Key:        "1F9D7-1F3FE",
		Value:      "🧗🏾",
		Descriptor: "person climbing: medium-dark skin tone",
	},
	"1F9D7-1F3FF": {
		Key:        "1F9D7-1F3FF",
		Value:      "🧗🏿",
		Descriptor: "person climbing: dark skin tone",
	},
	"1F9D7-200D-2642-FE0F": {
		Key:        "1F9D7-200D-2642-FE0F",
		Value:      "🧗‍♂️",
		Descriptor: "man climbing",
	},
	"1F9D7-1F3FB-200D-2642-FE0F": {
		Key:        "1F9D7-1F3FB-200D-2642-FE0F",
		Value:      "🧗🏻‍♂️",
		Descriptor: "man climbing: light skin tone",
	},
	"1F9D7-1F3FC-200D-2642-FE0F": {
		Key:        "1F9D7-1F3FC-200D-2642-FE0F",
		Value:      "🧗🏼‍♂️",
		Descriptor: "man climbing: medium-light skin tone",
	},
	"1F9D7-1F3FD-200D-2642-FE0F": {
		Key:        "1F9D7-1F3FD-200D-2642-FE0F",
		Value:      "🧗🏽‍♂️",
		Descriptor: "man climbing: medium skin tone",
	},
	"1F9D7-1F3FE-200D-2642-FE0F": {
		Key:        "1F9D7-1F3FE-200D-2642-FE0F",
		Value:      "🧗🏾‍♂️",
		Descriptor: "man climbing: medium-dark skin tone",
	},
	"1F9D7-1F3FF-200D-2642-FE0F": {
		Key:        "1F9D7-1F3FF-200D-2642-FE0F",
		Value:      "🧗🏿‍♂️",
		Descriptor: "man climbing: dark skin tone",
	},
	"1F9D7-200D-2640-FE0F": {
		Key:        "1F9D7-200D-2640-FE0F",
		Value:      "🧗‍♀️",
		Descriptor: "woman climbing",
	},
	"1F9D7-1F3FB-200D-2640-FE0F": {
		Key:        "1F9D7-1F3FB-200D-2640-FE0F",
		Value:      "🧗🏻‍♀️",
		Descriptor: "woman climbing: light skin tone",
	},
	"1F9D7-1F3FC-200D-2640-FE0F": {
		Key:        "1F9D7-1F3FC-200D-2640-FE0F",
		Value:      "🧗🏼‍♀️",
		Descriptor: "woman climbing: medium-light skin tone",
	},
	"1F9D7-1F3FD-200D-2640-FE0F": {
		Key:        "1F9D7-1F3FD-200D-2640-FE0F",
		Value:      "🧗🏽‍♀️",
		Descriptor: "woman climbing: medium skin tone",
	},
	"1F9D7-1F3FE-200D-2640-FE0F": {
		Key:        "1F9D7-1F3FE-200D-2640-FE0F",
		Value:      "🧗🏾‍♀️",
		Descriptor: "woman climbing: medium-dark skin tone",
	},
	"1F9D7-1F3FF-200D-2640-FE0F": {
		Key:        "1F9D7-1F3FF-200D-2640-FE0F",
		Value:      "🧗🏿‍♀️",
		Descriptor: "woman climbing: dark skin tone",
	},
	"1F93A": {
		Key:        "1F93A",
		Value:      "🤺",
		Descriptor: "person fencing",
	},
	"1F3C7": {
		Key:        "1F3C7",
		Value:      "🏇",
		Descriptor: "horse racing",
	},
	"1F3C7-1F3FB": {
		Key:        "1F3C7-1F3FB",
		Value:      "🏇🏻",
		Descriptor: "horse racing: light skin tone",
	},
	"1F3C7-1F3FC": {
		Key:        "1F3C7-1F3FC",
		Value:      "🏇🏼",
		Descriptor: "horse racing: medium-light skin tone",
	},
	"1F3C7-1F3FD": {
		Key:        "1F3C7-1F3FD",
		Value:      "🏇🏽",
		Descriptor: "horse racing: medium skin tone",
	},
	"1F3C7-1F3FE": {
		Key:        "1F3C7-1F3FE",
		Value:      "🏇🏾",
		Descriptor: "horse racing: medium-dark skin tone",
	},
	"1F3C7-1F3FF": {
		Key:        "1F3C7-1F3FF",
		Value:      "🏇🏿",
		Descriptor: "horse racing: dark skin tone",
	},
	"26F7-FE0F": {
		Key:        "26F7-FE0F",
		Value:      "⛷️",
		Descriptor: "skier",
	},
	"1F3C2": {
		Key:        "1F3C2",
		Value:      "🏂",
		Descriptor: "snowboarder",
	},
	"1F3C2-1F3FB": {
		Key:        "1F3C2-1F3FB",
		Value:      "🏂🏻",
		Descriptor: "snowboarder: light skin tone",
	},
	"1F3C2-1F3FC": {
		Key:        "1F3C2-1F3FC",
		Value:      "🏂🏼",
		Descriptor: "snowboarder: medium-light skin tone",
	},
	"1F3C2-1F3FD": {
		Key:        "1F3C2-1F3FD",
		Value:      "🏂🏽",
		Descriptor: "snowboarder: medium skin tone",
	},
	"1F3C2-1F3FE": {
		Key:        "1F3C2-1F3FE",
		Value:      "🏂🏾",
		Descriptor: "snowboarder: medium-dark skin tone",
	},
	"1F3C2-1F3FF": {
		Key:        "1F3C2-1F3FF",
		Value:      "🏂🏿",
		Descriptor: "snowboarder: dark skin tone",
	},
	"1F3CC-FE0F": {
		Key:        "1F3CC-FE0F",
		Value:      "🏌️",
		Descriptor: "person golfing",
	},
	"1F3CC-1F3FB": {
		Key:        "1F3CC-1F3FB",
		Value:      "🏌🏻",
		Descriptor: "person golfing: light skin tone",
	},
	"1F3CC-1F3FC": {
		Key:        "1F3CC-1F3FC",
		Value:      "🏌🏼",
		Descriptor: "person golfing: medium-light skin tone",
	},
	"1F3CC-1F3FD": {
		Key:        "1F3CC-1F3FD",
		Value:      "🏌🏽",
		Descriptor: "person golfing: medium skin tone",
	},
	"1F3CC-1F3FE": {
		Key:        "1F3CC-1F3FE",
		Value:      "🏌🏾",
		Descriptor: "person golfing: medium-dark skin tone",
	},
	"1F3CC-1F3FF": {
		Key:        "1F3CC-1F3FF",
		Value:      "🏌🏿",
		Descriptor: "person golfing: dark skin tone",
	},
	"1F3CC-FE0F-200D-2642-FE0F": {
		Key:        "1F3CC-FE0F-200D-2642-FE0F",
		Value:      "🏌️‍♂️",
		Descriptor: "man golfing",
	},
	"1F3CC-1F3FB-200D-2642-FE0F": {
		Key:        "1F3CC-1F3FB-200D-2642-FE0F",
		Value:      "🏌🏻‍♂️",
		Descriptor: "man golfing: light skin tone",
	},
	"1F3CC-1F3FC-200D-2642-FE0F": {
		Key:        "1F3CC-1F3FC-200D-2642-FE0F",
		Value:      "🏌🏼‍♂️",
		Descriptor: "man golfing: medium-light skin tone",
	},
	"1F3CC-1F3FD-200D-2642-FE0F": {
		Key:        "1F3CC-1F3FD-200D-2642-FE0F",
		Value:      "🏌🏽‍♂️",
		Descriptor: "man golfing: medium skin tone",
	},
	"1F3CC-1F3FE-200D-2642-FE0F": {
		Key:        "1F3CC-1F3FE-200D-2642-FE0F",
		Value:      "🏌🏾‍♂️",
		Descriptor: "man golfing: medium-dark skin tone",
	},
	"1F3CC-1F3FF-200D-2642-FE0F": {
		Key:        "1F3CC-1F3FF-200D-2642-FE0F",
		Value:      "🏌🏿‍♂️",
		Descriptor: "man golfing: dark skin tone",
	},
	"1F3CC-FE0F-200D-2640-FE0F": {
		Key:        "1F3CC-FE0F-200D-2640-FE0F",
		Value:      "🏌️‍♀️",
		Descriptor: "woman golfing",
	},
	"1F3CC-1F3FB-200D-2640-FE0F": {
		Key:        "1F3CC-1F3FB-200D-2640-FE0F",
		Value:      "🏌🏻‍♀️",
		Descriptor: "woman golfing: light skin tone",
	},
	"1F3CC-1F3FC-200D-2640-FE0F": {
		Key:        "1F3CC-1F3FC-200D-2640-FE0F",
		Value:      "🏌🏼‍♀️",
		Descriptor: "woman golfing: medium-light skin tone",
	},
	"1F3CC-1F3FD-200D-2640-FE0F": {
		Key:        "1F3CC-1F3FD-200D-2640-FE0F",
		Value:      "🏌🏽‍♀️",
		Descriptor: "woman golfing: medium skin tone",
	},
	"1F3CC-1F3FE-200D-2640-FE0F": {
		Key:        "1F3CC-1F3FE-200D-2640-FE0F",
		Value:      "🏌🏾‍♀️",
		Descriptor: "woman golfing: medium-dark skin tone",
	},
	"1F3CC-1F3FF-200D-2640-FE0F": {
		Key:        "1F3CC-1F3FF-200D-2640-FE0F",
		Value:      "🏌🏿‍♀️",
		Descriptor: "woman golfing: dark skin tone",
	},
	"1F3C4": {
		Key:        "1F3C4",
		Value:      "🏄",
		Descriptor: "person surfing",
	},
	"1F3C4-1F3FB": {
		Key:        "1F3C4-1F3FB",
		Value:      "🏄🏻",
		Descriptor: "person surfing: light skin tone",
	},
	"1F3C4-1F3FC": {
		Key:        "1F3C4-1F3FC",
		Value:      "🏄🏼",
		Descriptor: "person surfing: medium-light skin tone",
	},
	"1F3C4-1F3FD": {
		Key:        "1F3C4-1F3FD",
		Value:      "🏄🏽",
		Descriptor: "person surfing: medium skin tone",
	},
	"1F3C4-1F3FE": {
		Key:        "1F3C4-1F3FE",
		Value:      "🏄🏾",
		Descriptor: "person surfing: medium-dark skin tone",
	},
	"1F3C4-1F3FF": {
		Key:        "1F3C4-1F3FF",
		Value:      "🏄🏿",
		Descriptor: "person surfing: dark skin tone",
	},
	"1F3C4-200D-2642-FE0F": {
		Key:        "1F3C4-200D-2642-FE0F",
		Value:      "🏄‍♂️",
		Descriptor: "man surfing",
	},
	"1F3C4-1F3FB-200D-2642-FE0F": {
		Key:        "1F3C4-1F3FB-200D-2642-FE0F",
		Value:      "🏄🏻‍♂️",
		Descriptor: "man surfing: light skin tone",
	},
	"1F3C4-1F3FC-200D-2642-FE0F": {
		Key:        "1F3C4-1F3FC-200D-2642-FE0F",
		Value:      "🏄🏼‍♂️",
		Descriptor: "man surfing: medium-light skin tone",
	},
	"1F3C4-1F3FD-200D-2642-FE0F": {
		Key:        "1F3C4-1F3FD-200D-2642-FE0F",
		Value:      "🏄🏽‍♂️",
		Descriptor: "man surfing: medium skin tone",
	},
	"1F3C4-1F3FE-200D-2642-FE0F": {
		Key:        "1F3C4-1F3FE-200D-2642-FE0F",
		Value:      "🏄🏾‍♂️",
		Descriptor: "man surfing: medium-dark skin tone",
	},
	"1F3C4-1F3FF-200D-2642-FE0F": {
		Key:        "1F3C4-1F3FF-200D-2642-FE0F",
		Value:      "🏄🏿‍♂️",
		Descriptor: "man surfing: dark skin tone",
	},
	"1F3C4-200D-2640-FE0F": {
		Key:        "1F3C4-200D-2640-FE0F",
		Value:      "🏄‍♀️",
		Descriptor: "woman surfing",
	},
	"1F3C4-1F3FB-200D-2640-FE0F": {
		Key:        "1F3C4-1F3FB-200D-2640-FE0F",
		Value:      "🏄🏻‍♀️",
		Descriptor: "woman surfing: light skin tone",
	},
	"1F3C4-1F3FC-200D-2640-FE0F": {
		Key:        "1F3C4-1F3FC-200D-2640-FE0F",
		Value:      "🏄🏼‍♀️",
		Descriptor: "woman surfing: medium-light skin tone",
	},
	"1F3C4-1F3FD-200D-2640-FE0F": {
		Key:        "1F3C4-1F3FD-200D-2640-FE0F",
		Value:      "🏄🏽‍♀️",
		Descriptor: "woman surfing: medium skin tone",
	},
	"1F3C4-1F3FE-200D-2640-FE0F": {
		Key:        "1F3C4-1F3FE-200D-2640-FE0F",
		Value:      "🏄🏾‍♀️",
		Descriptor: "woman surfing: medium-dark skin tone",
	},
	"1F3C4-1F3FF-200D-2640-FE0F": {
		Key:        "1F3C4-1F3FF-200D-2640-FE0F",
		Value:      "🏄🏿‍♀️",
		Descriptor: "woman surfing: dark skin tone",
	},
	"1F6A3": {
		Key:        "1F6A3",
		Value:      "🚣",
		Descriptor: "person rowing boat",
	},
	"1F6A3-1F3FB": {
		Key:        "1F6A3-1F3FB",
		Value:      "🚣🏻",
		Descriptor: "person rowing boat: light skin tone",
	},
	"1F6A3-1F3FC": {
		Key:        "1F6A3-1F3FC",
		Value:      "🚣🏼",
		Descriptor: "person rowing boat: medium-light skin tone",
	},
	"1F6A3-1F3FD": {
		Key:        "1F6A3-1F3FD",
		Value:      "🚣🏽",
		Descriptor: "person rowing boat: medium skin tone",
	},
	"1F6A3-1F3FE": {
		Key:        "1F6A3-1F3FE",
		Value:      "🚣🏾",
		Descriptor: "person rowing boat: medium-dark skin tone",
	},
	"1F6A3-1F3FF": {
		Key:        "1F6A3-1F3FF",
		Value:      "🚣🏿",
		Descriptor: "person rowing boat: dark skin tone",
	},
	"1F6A3-200D-2642-FE0F": {
		Key:        "1F6A3-200D-2642-FE0F",
		Value:      "🚣‍♂️",
		Descriptor: "man rowing boat",
	},
	"1F6A3-1F3FB-200D-2642-FE0F": {
		Key:        "1F6A3-1F3FB-200D-2642-FE0F",
		Value:      "🚣🏻‍♂️",
		Descriptor: "man rowing boat: light skin tone",
	},
	"1F6A3-1F3FC-200D-2642-FE0F": {
		Key:        "1F6A3-1F3FC-200D-2642-FE0F",
		Value:      "🚣🏼‍♂️",
		Descriptor: "man rowing boat: medium-light skin tone",
	},
	"1F6A3-1F3FD-200D-2642-FE0F": {
		Key:        "1F6A3-1F3FD-200D-2642-FE0F",
		Value:      "🚣🏽‍♂️",
		Descriptor: "man rowing boat: medium skin tone",
	},
	"1F6A3-1F3FE-200D-2642-FE0F": {
		Key:        "1F6A3-1F3FE-200D-2642-FE0F",
		Value:      "🚣🏾‍♂️",
		Descriptor: "man rowing boat: medium-dark skin tone",
	},
	"1F6A3-1F3FF-200D-2642-FE0F": {
		Key:        "1F6A3-1F3FF-200D-2642-FE0F",
		Value:      "🚣🏿‍♂️",
		Descriptor: "man rowing boat: dark skin tone",
	},
	"1F6A3-200D-2640-FE0F": {
		Key:        "1F6A3-200D-2640-FE0F",
		Value:      "🚣‍♀️",
		Descriptor: "woman rowing boat",
	},
	"1F6A3-1F3FB-200D-2640-FE0F": {
		Key:        "1F6A3-1F3FB-200D-2640-FE0F",
		Value:      "🚣🏻‍♀️",
		Descriptor: "woman rowing boat: light skin tone",
	},
	"1F6A3-1F3FC-200D-2640-FE0F": {
		Key:        "1F6A3-1F3FC-200D-2640-FE0F",
		Value:      "🚣🏼‍♀️",
		Descriptor: "woman rowing boat: medium-light skin tone",
	},
	"1F6A3-1F3FD-200D-2640-FE0F": {
		Key:        "1F6A3-1F3FD-200D-2640-FE0F",
		Value:      "🚣🏽‍♀️",
		Descriptor: "woman rowing boat: medium skin tone",
	},
	"1F6A3-1F3FE-200D-2640-FE0F": {
		Key:        "1F6A3-1F3FE-200D-2640-FE0F",
		Value:      "🚣🏾‍♀️",
		Descriptor: "woman rowing boat: medium-dark skin tone",
	},
	"1F6A3-1F3FF-200D-2640-FE0F": {
		Key:        "1F6A3-1F3FF-200D-2640-FE0F",
		Value:      "🚣🏿‍♀️",
		Descriptor: "woman rowing boat: dark skin tone",
	},
	"1F3CA": {
		Key:        "1F3CA",
		Value:      "🏊",
		Descriptor: "person swimming",
	},
	"1F3CA-1F3FB": {
		Key:        "1F3CA-1F3FB",
		Value:      "🏊🏻",
		Descriptor: "person swimming: light skin tone",
	},
	"1F3CA-1F3FC": {
		Key:        "1F3CA-1F3FC",
		Value:      "🏊🏼",
		Descriptor: "person swimming: medium-light skin tone",
	},
	"1F3CA-1F3FD": {
		Key:        "1F3CA-1F3FD",
		Value:      "🏊🏽",
		Descriptor: "person swimming: medium skin tone",
	},
	"1F3CA-1F3FE": {
		Key:        "1F3CA-1F3FE",
		Value:      "🏊🏾",
		Descriptor: "person swimming: medium-dark skin tone",
	},
	"1F3CA-1F3FF": {
		Key:        "1F3CA-1F3FF",
		Value:      "🏊🏿",
		Descriptor: "person swimming: dark skin tone",
	},
	"1F3CA-200D-2642-FE0F": {
		Key:        "1F3CA-200D-2642-FE0F",
		Value:      "🏊‍♂️",
		Descriptor: "man swimming",
	},
	"1F3CA-1F3FB-200D-2642-FE0F": {
		Key:        "1F3CA-1F3FB-200D-2642-FE0F",
		Value:      "🏊🏻‍♂️",
		Descriptor: "man swimming: light skin tone",
	},
	"1F3CA-1F3FC-200D-2642-FE0F": {
		Key:        "1F3CA-1F3FC-200D-2642-FE0F",
		Value:      "🏊🏼‍♂️",
		Descriptor: "man swimming: medium-light skin tone",
	},
	"1F3CA-1F3FD-200D-2642-FE0F": {
		Key:        "1F3CA-1F3FD-200D-2642-FE0F",
		Value:      "🏊🏽‍♂️",
		Descriptor: "man swimming: medium skin tone",
	},
	"1F3CA-1F3FE-200D-2642-FE0F": {
		Key:        "1F3CA-1F3FE-200D-2642-FE0F",
		Value:      "🏊🏾‍♂️",
		Descriptor: "man swimming: medium-dark skin tone",
	},
	"1F3CA-1F3FF-200D-2642-FE0F": {
		Key:        "1F3CA-1F3FF-200D-2642-FE0F",
		Value:      "🏊🏿‍♂️",
		Descriptor: "man swimming: dark skin tone",
	},
	"1F3CA-200D-2640-FE0F": {
		Key:        "1F3CA-200D-2640-FE0F",
		Value:      "🏊‍♀️",
		Descriptor: "woman swimming",
	},
	"1F3CA-1F3FB-200D-2640-FE0F": {
		Key:        "1F3CA-1F3FB-200D-2640-FE0F",
		Value:      "🏊🏻‍♀️",
		Descriptor: "woman swimming: light skin tone",
	},
	"1F3CA-1F3FC-200D-2640-FE0F": {
		Key:        "1F3CA-1F3FC-200D-2640-FE0F",
		Value:      "🏊🏼‍♀️",
		Descriptor: "woman swimming: medium-light skin tone",
	},
	"1F3CA-1F3FD-200D-2640-FE0F": {
		Key:        "1F3CA-1F3FD-200D-2640-FE0F",
		Value:      "🏊🏽‍♀️",
		Descriptor: "woman swimming: medium skin tone",
	},
	"1F3CA-1F3FE-200D-2640-FE0F": {
		Key:        "1F3CA-1F3FE-200D-2640-FE0F",
		Value:      "🏊🏾‍♀️",
		Descriptor: "woman swimming: medium-dark skin tone",
	},
	"1F3CA-1F3FF-200D-2640-FE0F": {
		Key:        "1F3CA-1F3FF-200D-2640-FE0F",
		Value:      "🏊🏿‍♀️",
		Descriptor: "woman swimming: dark skin tone",
	},
	"26F9-FE0F": {
		Key:        "26F9-FE0F",
		Value:      "⛹️",
		Descriptor: "person bouncing ball",
	},
	"26F9-1F3FB": {
		Key:        "26F9-1F3FB",
		Value:      "⛹🏻",
		Descriptor: "person bouncing ball: light skin tone",
	},
	"26F9-1F3FC": {
		Key:        "26F9-1F3FC",
		Value:      "⛹🏼",
		Descriptor: "person bouncing ball: medium-light skin tone",
	},
	"26F9-1F3FD": {
		Key:        "26F9-1F3FD",
		Value:      "⛹🏽",
		Descriptor: "person bouncing ball: medium skin tone",
	},
	"26F9-1F3FE": {
		Key:        "26F9-1F3FE",
		Value:      "⛹🏾",
		Descriptor: "person bouncing ball: medium-dark skin tone",
	},
	"26F9-1F3FF": {
		Key:        "26F9-1F3FF",
		Value:      "⛹🏿",
		Descriptor: "person bouncing ball: dark skin tone",
	},
	"26F9-FE0F-200D-2642-FE0F": {
		Key:        "26F9-FE0F-200D-2642-FE0F",
		Value:      "⛹️‍♂️",
		Descriptor: "man bouncing ball",
	},
	"26F9-1F3FB-200D-2642-FE0F": {
		Key:        "26F9-1F3FB-200D-2642-FE0F",
		Value:      "⛹🏻‍♂️",
		Descriptor: "man bouncing ball: light skin tone",
	},
	"26F9-1F3FC-200D-2642-FE0F": {
		Key:        "26F9-1F3FC-200D-2642-FE0F",
		Value:      "⛹🏼‍♂️",
		Descriptor: "man bouncing ball: medium-light skin tone",
	},
	"26F9-1F3FD-200D-2642-FE0F": {
		Key:        "26F9-1F3FD-200D-2642-FE0F",
		Value:      "⛹🏽‍♂️",
		Descriptor: "man bouncing ball: medium skin tone",
	},
	"26F9-1F3FE-200D-2642-FE0F": {
		Key:        "26F9-1F3FE-200D-2642-FE0F",
		Value:      "⛹🏾‍♂️",
		Descriptor: "man bouncing ball: medium-dark skin tone",
	},
	"26F9-1F3FF-200D-2642-FE0F": {
		Key:        "26F9-1F3FF-200D-2642-FE0F",
		Value:      "⛹🏿‍♂️",
		Descriptor: "man bouncing ball: dark skin tone",
	},
	"26F9-FE0F-200D-2640-FE0F": {
		Key:        "26F9-FE0F-200D-2640-FE0F",
		Value:      "⛹️‍♀️",
		Descriptor: "woman bouncing ball",
	},
	"26F9-1F3FB-200D-2640-FE0F": {
		Key:        "26F9-1F3FB-200D-2640-FE0F",
		Value:      "⛹🏻‍♀️",
		Descriptor: "woman bouncing ball: light skin tone",
	},
	"26F9-1F3FC-200D-2640-FE0F": {
		Key:        "26F9-1F3FC-200D-2640-FE0F",
		Value:      "⛹🏼‍♀️",
		Descriptor: "woman bouncing ball: medium-light skin tone",
	},
	"26F9-1F3FD-200D-2640-FE0F": {
		Key:        "26F9-1F3FD-200D-2640-FE0F",
		Value:      "⛹🏽‍♀️",
		Descriptor: "woman bouncing ball: medium skin tone",
	},
	"26F9-1F3FE-200D-2640-FE0F": {
		Key:        "26F9-1F3FE-200D-2640-FE0F",
		Value:      "⛹🏾‍♀️",
		Descriptor: "woman bouncing ball: medium-dark skin tone",
	},
	"26F9-1F3FF-200D-2640-FE0F": {
		Key:        "26F9-1F3FF-200D-2640-FE0F",
		Value:      "⛹🏿‍♀️",
		Descriptor: "woman bouncing ball: dark skin tone",
	},
	"1F3CB-FE0F": {
		Key:        "1F3CB-FE0F",
		Value:      "🏋️",
		Descriptor: "person lifting weights",
	},
	"1F3CB-1F3FB": {
		Key:        "1F3CB-1F3FB",
		Value:      "🏋🏻",
		Descriptor: "person lifting weights: light skin tone",
	},
	"1F3CB-1F3FC": {
		Key:        "1F3CB-1F3FC",
		Value:      "🏋🏼",
		Descriptor: "person lifting weights: medium-light skin tone",
	},
	"1F3CB-1F3FD": {
		Key:        "1F3CB-1F3FD",
		Value:      "🏋🏽",
		Descriptor: "person lifting weights: medium skin tone",
	},
	"1F3CB-1F3FE": {
		Key:        "1F3CB-1F3FE",
		Value:      "🏋🏾",
		Descriptor: "person lifting weights: medium-dark skin tone",
	},
	"1F3CB-1F3FF": {
		Key:        "1F3CB-1F3FF",
		Value:      "🏋🏿",
		Descriptor: "person lifting weights: dark skin tone",
	},
	"1F3CB-FE0F-200D-2642-FE0F": {
		Key:        "1F3CB-FE0F-200D-2642-FE0F",
		Value:      "🏋️‍♂️",
		Descriptor: "man lifting weights",
	},
	"1F3CB-1F3FB-200D-2642-FE0F": {
		Key:        "1F3CB-1F3FB-200D-2642-FE0F",
		Value:      "🏋🏻‍♂️",
		Descriptor: "man lifting weights: light skin tone",
	},
	"1F3CB-1F3FC-200D-2642-FE0F": {
		Key:        "1F3CB-1F3FC-200D-2642-FE0F",
		Value:      "🏋🏼‍♂️",
		Descriptor: "man lifting weights: medium-light skin tone",
	},
	"1F3CB-1F3FD-200D-2642-FE0F": {
		Key:        "1F3CB-1F3FD-200D-2642-FE0F",
		Value:      "🏋🏽‍♂️",
		Descriptor: "man lifting weights: medium skin tone",
	},
	"1F3CB-1F3FE-200D-2642-FE0F": {
		Key:        "1F3CB-1F3FE-200D-2642-FE0F",
		Value:      "🏋🏾‍♂️",
		Descriptor: "man lifting weights: medium-dark skin tone",
	},
	"1F3CB-1F3FF-200D-2642-FE0F": {
		Key:        "1F3CB-1F3FF-200D-2642-FE0F",
		Value:      "🏋🏿‍♂️",
		Descriptor: "man lifting weights: dark skin tone",
	},
	"1F3CB-FE0F-200D-2640-FE0F": {
		Key:        "1F3CB-FE0F-200D-2640-FE0F",
		Value:      "🏋️‍♀️",
		Descriptor: "woman lifting weights",
	},
	"1F3CB-1F3FB-200D-2640-FE0F": {
		Key:        "1F3CB-1F3FB-200D-2640-FE0F",
		Value:      "🏋🏻‍♀️",
		Descriptor: "woman lifting weights: light skin tone",
	},
	"1F3CB-1F3FC-200D-2640-FE0F": {
		Key:        "1F3CB-1F3FC-200D-2640-FE0F",
		Value:      "🏋🏼‍♀️",
		Descriptor: "woman lifting weights: medium-light skin tone",
	},
	"1F3CB-1F3FD-200D-2640-FE0F": {
		Key:        "1F3CB-1F3FD-200D-2640-FE0F",
		Value:      "🏋🏽‍♀️",
		Descriptor: "woman lifting weights: medium skin tone",
	},
	"1F3CB-1F3FE-200D-2640-FE0F": {
		Key:        "1F3CB-1F3FE-200D-2640-FE0F",
		Value:      "🏋🏾‍♀️",
		Descriptor: "woman lifting weights: medium-dark skin tone",
	},
	"1F3CB-1F3FF-200D-2640-FE0F": {
		Key:        "1F3CB-1F3FF-200D-2640-FE0F",
		Value:      "🏋🏿‍♀️",
		Descriptor: "woman lifting weights: dark skin tone",
	},
	"1F6B4": {
		Key:        "1F6B4",
		Value:      "🚴",
		Descriptor: "person biking",
	},
	"1F6B4-1F3FB": {
		Key:        "1F6B4-1F3FB",
		Value:      "🚴🏻",
		Descriptor: "person biking: light skin tone",
	},
	"1F6B4-1F3FC": {
		Key:        "1F6B4-1F3FC",
		Value:      "🚴🏼",
		Descriptor: "person biking: medium-light skin tone",
	},
	"1F6B4-1F3FD": {
		Key:        "1F6B4-1F3FD",
		Value:      "🚴🏽",
		Descriptor: "person biking: medium skin tone",
	},
	"1F6B4-1F3FE": {
		Key:        "1F6B4-1F3FE",
		Value:      "🚴🏾",
		Descriptor: "person biking: medium-dark skin tone",
	},
	"1F6B4-1F3FF": {
		Key:        "1F6B4-1F3FF",
		Value:      "🚴🏿",
		Descriptor: "person biking: dark skin tone",
	},
	"1F6B4-200D-2642-FE0F": {
		Key:        "1F6B4-200D-2642-FE0F",
		Value:      "🚴‍♂️",
		Descriptor: "man biking",
	},
	"1F6B4-1F3FB-200D-2642-FE0F": {
		Key:        "1F6B4-1F3FB-200D-2642-FE0F",
		Value:      "🚴🏻‍♂️",
		Descriptor: "man biking: light skin tone",
	},
	"1F6B4-1F3FC-200D-2642-FE0F": {
		Key:        "1F6B4-1F3FC-200D-2642-FE0F",
		Value:      "🚴🏼‍♂️",
		Descriptor: "man biking: medium-light skin tone",
	},
	"1F6B4-1F3FD-200D-2642-FE0F": {
		Key:        "1F6B4-1F3FD-200D-2642-FE0F",
		Value:      "🚴🏽‍♂️",
		Descriptor: "man biking: medium skin tone",
	},
	"1F6B4-1F3FE-200D-2642-FE0F": {
		Key:        "1F6B4-1F3FE-200D-2642-FE0F",
		Value:      "🚴🏾‍♂️",
		Descriptor: "man biking: medium-dark skin tone",
	},
	"1F6B4-1F3FF-200D-2642-FE0F": {
		Key:        "1F6B4-1F3FF-200D-2642-FE0F",
		Value:      "🚴🏿‍♂️",
		Descriptor: "man biking: dark skin tone",
	},
	"1F6B4-200D-2640-FE0F": {
		Key:        "1F6B4-200D-2640-FE0F",
		Value:      "🚴‍♀️",
		Descriptor: "woman biking",
	},
	"1F6B4-1F3FB-200D-2640-FE0F": {
		Key:        "1F6B4-1F3FB-200D-2640-FE0F",
		Value:      "🚴🏻‍♀️",
		Descriptor: "woman biking: light skin tone",
	},
	"1F6B4-1F3FC-200D-2640-FE0F": {
		Key:        "1F6B4-1F3FC-200D-2640-FE0F",
		Value:      "🚴🏼‍♀️",
		Descriptor: "woman biking: medium-light skin tone",
	},
	"1F6B4-1F3FD-200D-2640-FE0F": {
		Key:        "1F6B4-1F3FD-200D-2640-FE0F",
		Value:      "🚴🏽‍♀️",
		Descriptor: "woman biking: medium skin tone",
	},
	"1F6B4-1F3FE-200D-2640-FE0F": {
		Key:        "1F6B4-1F3FE-200D-2640-FE0F",
		Value:      "🚴🏾‍♀️",
		Descriptor: "woman biking: medium-dark skin tone",
	},
	"1F6B4-1F3FF-200D-2640-FE0F": {
		Key:        "1F6B4-1F3FF-200D-2640-FE0F",
		Value:      "🚴🏿‍♀️",
		Descriptor: "woman biking: dark skin tone",
	},
	"1F6B5": {
		Key:        "1F6B5",
		Value:      "🚵",
		Descriptor: "person mountain biking",
	},
	"1F6B5-1F3FB": {
		Key:        "1F6B5-1F3FB",
		Value:      "🚵🏻",
		Descriptor: "person mountain biking: light skin tone",
	},
	"1F6B5-1F3FC": {
		Key:        "1F6B5-1F3FC",
		Value:      "🚵🏼",
		Descriptor: "person mountain biking: medium-light skin tone",
	},
	"1F6B5-1F3FD": {
		Key:        "1F6B5-1F3FD",
		Value:      "🚵🏽",
		Descriptor: "person mountain biking: medium skin tone",
	},
	"1F6B5-1F3FE": {
		Key:        "1F6B5-1F3FE",
		Value:      "🚵🏾",
		Descriptor: "person mountain biking: medium-dark skin tone",
	},
	"1F6B5-1F3FF": {
		Key:        "1F6B5-1F3FF",
		Value:      "🚵🏿",
		Descriptor: "person mountain biking: dark skin tone",
	},
	"1F6B5-200D-2642-FE0F": {
		Key:        "1F6B5-200D-2642-FE0F",
		Value:      "🚵‍♂️",
		Descriptor: "man mountain biking",
	},
	"1F6B5-1F3FB-200D-2642-FE0F": {
		Key:        "1F6B5-1F3FB-200D-2642-FE0F",
		Value:      "🚵🏻‍♂️",
		Descriptor: "man mountain biking: light skin tone",
	},
	"1F6B5-1F3FC-200D-2642-FE0F": {
		Key:        "1F6B5-1F3FC-200D-2642-FE0F",
		Value:      "🚵🏼‍♂️",
		Descriptor: "man mountain biking: medium-light skin tone",
	},
	"1F6B5-1F3FD-200D-2642-FE0F": {
		Key:        "1F6B5-1F3FD-200D-2642-FE0F",
		Value:      "🚵🏽‍♂️",
		Descriptor: "man mountain biking: medium skin tone",
	},
	"1F6B5-1F3FE-200D-2642-FE0F": {
		Key:        "1F6B5-1F3FE-200D-2642-FE0F",
		Value:      "🚵🏾‍♂️",
		Descriptor: "man mountain biking: medium-dark skin tone",
	},
	"1F6B5-1F3FF-200D-2642-FE0F": {
		Key:        "1F6B5-1F3FF-200D-2642-FE0F",
		Value:      "🚵🏿‍♂️",
		Descriptor: "man mountain biking: dark skin tone",
	},
	"1F6B5-200D-2640-FE0F": {
		Key:        "1F6B5-200D-2640-FE0F",
		Value:      "🚵‍♀️",
		Descriptor: "woman mountain biking",
	},
	"1F6B5-1F3FB-200D-2640-FE0F": {
		Key:        "1F6B5-1F3FB-200D-2640-FE0F",
		Value:      "🚵🏻‍♀️",
		Descriptor: "woman mountain biking: light skin tone",
	},
	"1F6B5-1F3FC-200D-2640-FE0F": {
		Key:        "1F6B5-1F3FC-200D-2640-FE0F",
		Value:      "🚵🏼‍♀️",
		Descriptor: "woman mountain biking: medium-light skin tone",
	},
	"1F6B5-1F3FD-200D-2640-FE0F": {
		Key:        "1F6B5-1F3FD-200D-2640-FE0F",
		Value:      "🚵🏽‍♀️",
		Descriptor: "woman mountain biking: medium skin tone",
	},
	"1F6B5-1F3FE-200D-2640-FE0F": {
		Key:        "1F6B5-1F3FE-200D-2640-FE0F",
		Value:      "🚵🏾‍♀️",
		Descriptor: "woman mountain biking: medium-dark skin tone",
	},
	"1F6B5-1F3FF-200D-2640-FE0F": {
		Key:        "1F6B5-1F3FF-200D-2640-FE0F",
		Value:      "🚵🏿‍♀️",
		Descriptor: "woman mountain biking: dark skin tone",
	},
	"1F938": {
		Key:        "1F938",
		Value:      "🤸",
		Descriptor: "person cartwheeling",
	},
	"1F938-1F3FB": {
		Key:        "1F938-1F3FB",
		Value:      "🤸🏻",
		Descriptor: "person cartwheeling: light skin tone",
	},
	"1F938-1F3FC": {
		Key:        "1F938-1F3FC",
		Value:      "🤸🏼",
		Descriptor: "person cartwheeling: medium-light skin tone",
	},
	"1F938-1F3FD": {
		Key:        "1F938-1F3FD",
		Value:      "🤸🏽",
		Descriptor: "person cartwheeling: medium skin tone",
	},
	"1F938-1F3FE": {
		Key:        "1F938-1F3FE",
		Value:      "🤸🏾",
		Descriptor: "person cartwheeling: medium-dark skin tone",
	},
	"1F938-1F3FF": {
		Key:        "1F938-1F3FF",
		Value:      "🤸🏿",
		Descriptor: "person cartwheeling: dark skin tone",
	},
	"1F938-200D-2642-FE0F": {
		Key:        "1F938-200D-2642-FE0F",
		Value:      "🤸‍♂️",
		Descriptor: "man cartwheeling",
	},
	"1F938-1F3FB-200D-2642-FE0F": {
		Key:        "1F938-1F3FB-200D-2642-FE0F",
		Value:      "🤸🏻‍♂️",
		Descriptor: "man cartwheeling: light skin tone",
	},
	"1F938-1F3FC-200D-2642-FE0F": {
		Key:        "1F938-1F3FC-200D-2642-FE0F",
		Value:      "🤸🏼‍♂️",
		Descriptor: "man cartwheeling: medium-light skin tone",
	},
	"1F938-1F3FD-200D-2642-FE0F": {
		Key:        "1F938-1F3FD-200D-2642-FE0F",
		Value:      "🤸🏽‍♂️",
		Descriptor: "man cartwheeling: medium skin tone",
	},
	"1F938-1F3FE-200D-2642-FE0F": {
		Key:        "1F938-1F3FE-200D-2642-FE0F",
		Value:      "🤸🏾‍♂️",
		Descriptor: "man cartwheeling: medium-dark skin tone",
	},
	"1F938-1F3FF-200D-2642-FE0F": {
		Key:        "1F938-1F3FF-200D-2642-FE0F",
		Value:      "🤸🏿‍♂️",
		Descriptor: "man cartwheeling: dark skin tone",
	},
	"1F938-200D-2640-FE0F": {
		Key:        "1F938-200D-2640-FE0F",
		Value:      "🤸‍♀️",
		Descriptor: "woman cartwheeling",
	},
	"1F938-1F3FB-200D-2640-FE0F": {
		Key:        "1F938-1F3FB-200D-2640-FE0F",
		Value:      "🤸🏻‍♀️",
		Descriptor: "woman cartwheeling: light skin tone",
	},
	"1F938-1F3FC-200D-2640-FE0F": {
		Key:        "1F938-1F3FC-200D-2640-FE0F",
		Value:      "🤸🏼‍♀️",
		Descriptor: "woman cartwheeling: medium-light skin tone",
	},
	"1F938-1F3FD-200D-2640-FE0F": {
		Key:        "1F938-1F3FD-200D-2640-FE0F",
		Value:      "🤸🏽‍♀️",
		Descriptor: "woman cartwheeling: medium skin tone",
	},
	"1F938-1F3FE-200D-2640-FE0F": {
		Key:        "1F938-1F3FE-200D-2640-FE0F",
		Value:      "🤸🏾‍♀️",
		Descriptor: "woman cartwheeling: medium-dark skin tone",
	},
	"1F938-1F3FF-200D-2640-FE0F": {
		Key:        "1F938-1F3FF-200D-2640-FE0F",
		Value:      "🤸🏿‍♀️",
		Descriptor: "woman cartwheeling: dark skin tone",
	},
	"1F93C": {
		Key:        "1F93C",
		Value:      "🤼",
		Descriptor: "people wrestling",
	},
	"1F93C-200D-2642-FE0F": {
		Key:        "1F93C-200D-2642-FE0F",
		Value:      "🤼‍♂️",
		Descriptor: "men wrestling",
	},
	"1F93C-200D-2640-FE0F": {
		Key:        "1F93C-200D-2640-FE0F",
		Value:      "🤼‍♀️",
		Descriptor: "women wrestling",
	},
	"1F93D": {
		Key:        "1F93D",
		Value:      "🤽",
		Descriptor: "person playing water polo",
	},
	"1F93D-1F3FB": {
		Key:        "1F93D-1F3FB",
		Value:      "🤽🏻",
		Descriptor: "person playing water polo: light skin tone",
	},
	"1F93D-1F3FC": {
		Key:        "1F93D-1F3FC",
		Value:      "🤽🏼",
		Descriptor: "person playing water polo: medium-light skin tone",
	},
	"1F93D-1F3FD": {
		Key:        "1F93D-1F3FD",
		Value:      "🤽🏽",
		Descriptor: "person playing water polo: medium skin tone",
	},
	"1F93D-1F3FE": {
		Key:        "1F93D-1F3FE",
		Value:      "🤽🏾",
		Descriptor: "person playing water polo: medium-dark skin tone",
	},
	"1F93D-1F3FF": {
		Key:        "1F93D-1F3FF",
		Value:      "🤽🏿",
		Descriptor: "person playing water polo: dark skin tone",
	},
	"1F93D-200D-2642-FE0F": {
		Key:        "1F93D-200D-2642-FE0F",
		Value:      "🤽‍♂️",
		Descriptor: "man playing water polo",
	},
	"1F93D-1F3FB-200D-2642-FE0F": {
		Key:        "1F93D-1F3FB-200D-2642-FE0F",
		Value:      "🤽🏻‍♂️",
		Descriptor: "man playing water polo: light skin tone",
	},
	"1F93D-1F3FC-200D-2642-FE0F": {
		Key:        "1F93D-1F3FC-200D-2642-FE0F",
		Value:      "🤽🏼‍♂️",
		Descriptor: "man playing water polo: medium-light skin tone",
	},
	"1F93D-1F3FD-200D-2642-FE0F": {
		Key:        "1F93D-1F3FD-200D-2642-FE0F",
		Value:      "🤽🏽‍♂️",
		Descriptor: "man playing water polo: medium skin tone",
	},
	"1F93D-1F3FE-200D-2642-FE0F": {
		Key:        "1F93D-1F3FE-200D-2642-FE0F",
		Value:      "🤽🏾‍♂️",
		Descriptor: "man playing water polo: medium-dark skin tone",
	},
	"1F93D-1F3FF-200D-2642-FE0F": {
		Key:        "1F93D-1F3FF-200D-2642-FE0F",
		Value:      "🤽🏿‍♂️",
		Descriptor: "man playing water polo: dark skin tone",
	},
	"1F93D-200D-2640-FE0F": {
		Key:        "1F93D-200D-2640-FE0F",
		Value:      "🤽‍♀️",
		Descriptor: "woman playing water polo",
	},
	"1F93D-1F3FB-200D-2640-FE0F": {
		Key:        "1F93D-1F3FB-200D-2640-FE0F",
		Value:      "🤽🏻‍♀️",
		Descriptor: "woman playing water polo: light skin tone",
	},
	"1F93D-1F3FC-200D-2640-FE0F": {
		Key:        "1F93D-1F3FC-200D-2640-FE0F",
		Value:      "🤽🏼‍♀️",
		Descriptor: "woman playing water polo: medium-light skin tone",
	},
	"1F93D-1F3FD-200D-2640-FE0F": {
		Key:        "1F93D-1F3FD-200D-2640-FE0F",
		Value:      "🤽🏽‍♀️",
		Descriptor: "woman playing water polo: medium skin tone",
	},
	"1F93D-1F3FE-200D-2640-FE0F": {
		Key:        "1F93D-1F3FE-200D-2640-FE0F",
		Value:      "🤽🏾‍♀️",
		Descriptor: "woman playing water polo: medium-dark skin tone",
	},
	"1F93D-1F3FF-200D-2640-FE0F": {
		Key:        "1F93D-1F3FF-200D-2640-FE0F",
		Value:      "🤽🏿‍♀️",
		Descriptor: "woman playing water polo: dark skin tone",
	},
	"1F93E": {
		Key:        "1F93E",
		Value:      "🤾",
		Descriptor: "person playing handball",
	},
	"1F93E-1F3FB": {
		Key:        "1F93E-1F3FB",
		Value:      "🤾🏻",
		Descriptor: "person playing handball: light skin tone",
	},
	"1F93E-1F3FC": {
		Key:        "1F93E-1F3FC",
		Value:      "🤾🏼",
		Descriptor: "person playing handball: medium-light skin tone",
	},
	"1F93E-1F3FD": {
		Key:        "1F93E-1F3FD",
		Value:      "🤾🏽",
		Descriptor: "person playing handball: medium skin tone",
	},
	"1F93E-1F3FE": {
		Key:        "1F93E-1F3FE",
		Value:      "🤾🏾",
		Descriptor: "person playing handball: medium-dark skin tone",
	},
	"1F93E-1F3FF": {
		Key:        "1F93E-1F3FF",
		Value:      "🤾🏿",
		Descriptor: "person playing handball: dark skin tone",
	},
	"1F93E-200D-2642-FE0F": {
		Key:        "1F93E-200D-2642-FE0F",
		Value:      "🤾‍♂️",
		Descriptor: "man playing handball",
	},
	"1F93E-1F3FB-200D-2642-FE0F": {
		Key:        "1F93E-1F3FB-200D-2642-FE0F",
		Value:      "🤾🏻‍♂️",
		Descriptor: "man playing handball: light skin tone",
	},
	"1F93E-1F3FC-200D-2642-FE0F": {
		Key:        "1F93E-1F3FC-200D-2642-FE0F",
		Value:      "🤾🏼‍♂️",
		Descriptor: "man playing handball: medium-light skin tone",
	},
	"1F93E-1F3FD-200D-2642-FE0F": {
		Key:        "1F93E-1F3FD-200D-2642-FE0F",
		Value:      "🤾🏽‍♂️",
		Descriptor: "man playing handball: medium skin tone",
	},
	"1F93E-1F3FE-200D-2642-FE0F": {
		Key:        "1F93E-1F3FE-200D-2642-FE0F",
		Value:      "🤾🏾‍♂️",
		Descriptor: "man playing handball: medium-dark skin tone",
	},
	"1F93E-1F3FF-200D-2642-FE0F": {
		Key:        "1F93E-1F3FF-200D-2642-FE0F",
		Value:      "🤾🏿‍♂️",
		Descriptor: "man playing handball: dark skin tone",
	},
	"1F93E-200D-2640-FE0F": {
		Key:        "1F93E-200D-2640-FE0F",
		Value:      "🤾‍♀️",
		Descriptor: "woman playing handball",
	},
	"1F93E-1F3FB-200D-2640-FE0F": {
		Key:        "1F93E-1F3FB-200D-2640-FE0F",
		Value:      "🤾🏻‍♀️",
		Descriptor: "woman playing handball: light skin tone",
	},
	"1F93E-1F3FC-200D-2640-FE0F": {
		Key:        "1F93E-1F3FC-200D-2640-FE0F",
		Value:      "🤾🏼‍♀️",
		Descriptor: "woman playing handball: medium-light skin tone",
	},
	"1F93E-1F3FD-200D-2640-FE0F": {
		Key:        "1F93E-1F3FD-200D-2640-FE0F",
		Value:      "🤾🏽‍♀️",
		Descriptor: "woman playing handball: medium skin tone",
	},
	"1F93E-1F3FE-200D-2640-FE0F": {
		Key:        "1F93E-1F3FE-200D-2640-FE0F",
		Value:      "🤾🏾‍♀️",
		Descriptor: "woman playing handball: medium-dark skin tone",
	},
	"1F93E-1F3FF-200D-2640-FE0F": {
		Key:        "1F93E-1F3FF-200D-2640-FE0F",
		Value:      "🤾🏿‍♀️",
		Descriptor: "woman playing handball: dark skin tone",
	},
	"1F939": {
		Key:        "1F939",
		Value:      "🤹",
		Descriptor: "person juggling",
	},
	"1F939-1F3FB": {
		Key:        "1F939-1F3FB",
		Value:      "🤹🏻",
		Descriptor: "person juggling: light skin tone",
	},
	"1F939-1F3FC": {
		Key:        "1F939-1F3FC",
		Value:      "🤹🏼",
		Descriptor: "person juggling: medium-light skin tone",
	},
	"1F939-1F3FD": {
		Key:        "1F939-1F3FD",
		Value:      "🤹🏽",
		Descriptor: "person juggling: medium skin tone",
	},
	"1F939-1F3FE": {
		Key:        "1F939-1F3FE",
		Value:      "🤹🏾",
		Descriptor: "person juggling: medium-dark skin tone",
	},
	"1F939-1F3FF": {
		Key:        "1F939-1F3FF",
		Value:      "🤹🏿",
		Descriptor: "person juggling: dark skin tone",
	},
	"1F939-200D-2642-FE0F": {
		Key:        "1F939-200D-2642-FE0F",
		Value:      "🤹‍♂️",
		Descriptor: "man juggling",
	},
	"1F939-1F3FB-200D-2642-FE0F": {
		Key:        "1F939-1F3FB-200D-2642-FE0F",
		Value:      "🤹🏻‍♂️",
		Descriptor: "man juggling: light skin tone",
	},
	"1F939-1F3FC-200D-2642-FE0F": {
		Key:        "1F939-1F3FC-200D-2642-FE0F",
		Value:      "🤹🏼‍♂️",
		Descriptor: "man juggling: medium-light skin tone",
	},
	"1F939-1F3FD-200D-2642-FE0F": {
		Key:        "1F939-1F3FD-200D-2642-FE0F",
		Value:      "🤹🏽‍♂️",
		Descriptor: "man juggling: medium skin tone",
	},
	"1F939-1F3FE-200D-2642-FE0F": {
		Key:        "1F939-1F3FE-200D-2642-FE0F",
		Value:      "🤹🏾‍♂️",
		Descriptor: "man juggling: medium-dark skin tone",
	},
	"1F939-1F3FF-200D-2642-FE0F": {
		Key:        "1F939-1F3FF-200D-2642-FE0F",
		Value:      "🤹🏿‍♂️",
		Descriptor: "man juggling: dark skin tone",
	},
	"1F939-200D-2640-FE0F": {
		Key:        "1F939-200D-2640-FE0F",
		Value:      "🤹‍♀️",
		Descriptor: "woman juggling",
	},
	"1F939-1F3FB-200D-2640-FE0F": {
		Key:        "1F939-1F3FB-200D-2640-FE0F",
		Value:      "🤹🏻‍♀️",
		Descriptor: "woman juggling: light skin tone",
	},
	"1F939-1F3FC-200D-2640-FE0F": {
		Key:        "1F939-1F3FC-200D-2640-FE0F",
		Value:      "🤹🏼‍♀️",
		Descriptor: "woman juggling: medium-light skin tone",
	},
	"1F939-1F3FD-200D-2640-FE0F": {
		Key:        "1F939-1F3FD-200D-2640-FE0F",
		Value:      "🤹🏽‍♀️",
		Descriptor: "woman juggling: medium skin tone",
	},
	"1F939-1F3FE-200D-2640-FE0F": {
		Key:        "1F939-1F3FE-200D-2640-FE0F",
		Value:      "🤹🏾‍♀️",
		Descriptor: "woman juggling: medium-dark skin tone",
	},
	"1F939-1F3FF-200D-2640-FE0F": {
		Key:        "1F939-1F3FF-200D-2640-FE0F",
		Value:      "🤹🏿‍♀️",
		Descriptor: "woman juggling: dark skin tone",
	},
	"1F9D8": {
		Key:        "1F9D8",
		Value:      "🧘",
		Descriptor: "person in lotus position",
	},
	"1F9D8-1F3FB": {
		Key:        "1F9D8-1F3FB",
		Value:      "🧘🏻",
		Descriptor: "person in lotus position: light skin tone",
	},
	"1F9D8-1F3FC": {
		Key:        "1F9D8-1F3FC",
		Value:      "🧘🏼",
		Descriptor: "person in lotus position: medium-light skin tone",
	},
	"1F9D8-1F3FD": {
		Key:        "1F9D8-1F3FD",
		Value:      "🧘🏽",
		Descriptor: "person in lotus position: medium skin tone",
	},
	"1F9D8-1F3FE": {
		Key:        "1F9D8-1F3FE",
		Value:      "🧘🏾",
		Descriptor: "person in lotus position: medium-dark skin tone",
	},
	"1F9D8-1F3FF": {
		Key:        "1F9D8-1F3FF",
		Value:      "🧘🏿",
		Descriptor: "person in lotus position: dark skin tone",
	},
	"1F9D8-200D-2642-FE0F": {
		Key:        "1F9D8-200D-2642-FE0F",
		Value:      "🧘‍♂️",
		Descriptor: "man in lotus position",
	},
	"1F9D8-1F3FB-200D-2642-FE0F": {
		Key:        "1F9D8-1F3FB-200D-2642-FE0F",
		Value:      "🧘🏻‍♂️",
		Descriptor: "man in lotus position: light skin tone",
	},
	"1F9D8-1F3FC-200D-2642-FE0F": {
		Key:        "1F9D8-1F3FC-200D-2642-FE0F",
		Value:      "🧘🏼‍♂️",
		Descriptor: "man in lotus position: medium-light skin tone",
	},
	"1F9D8-1F3FD-200D-2642-FE0F": {
		Key:        "1F9D8-1F3FD-200D-2642-FE0F",
		Value:      "🧘🏽‍♂️",
		Descriptor: "man in lotus position: medium skin tone",
	},
	"1F9D8-1F3FE-200D-2642-FE0F": {
		Key:        "1F9D8-1F3FE-200D-2642-FE0F",
		Value:      "🧘🏾‍♂️",
		Descriptor: "man in lotus position: medium-dark skin tone",
	},
	"1F9D8-1F3FF-200D-2642-FE0F": {
		Key:        "1F9D8-1F3FF-200D-2642-FE0F",
		Value:      "🧘🏿‍♂️",
		Descriptor: "man in lotus position: dark skin tone",
	},
	"1F9D8-200D-2640-FE0F": {
		Key:        "1F9D8-200D-2640-FE0F",
		Value:      "🧘‍♀️",
		Descriptor: "woman in lotus position",
	},
	"1F9D8-1F3FB-200D-2640-FE0F": {
		Key:        "1F9D8-1F3FB-200D-2640-FE0F",
		Value:      "🧘🏻‍♀️",
		Descriptor: "woman in lotus position: light skin tone",
	},
	"1F9D8-1F3FC-200D-2640-FE0F": {
		Key:        "1F9D8-1F3FC-200D-2640-FE0F",
		Value:      "🧘🏼‍♀️",
		Descriptor: "woman in lotus position: medium-light skin tone",
	},
	"1F9D8-1F3FD-200D-2640-FE0F": {
		Key:        "1F9D8-1F3FD-200D-2640-FE0F",
		Value:      "🧘🏽‍♀️",
		Descriptor: "woman in lotus position: medium skin tone",
	},
	"1F9D8-1F3FE-200D-2640-FE0F": {
		Key:        "1F9D8-1F3FE-200D-2640-FE0F",
		Value:      "🧘🏾‍♀️",
		Descriptor: "woman in lotus position: medium-dark skin tone",
	},
	"1F9D8-1F3FF-200D-2640-FE0F": {
		Key:        "1F9D8-1F3FF-200D-2640-FE0F",
		Value:      "🧘🏿‍♀️",
		Descriptor: "woman in lotus position: dark skin tone",
	},
	"1F6C0": {
		Key:        "1F6C0",
		Value:      "🛀",
		Descriptor: "person taking bath",
	},
	"1F6C0-1F3FB": {
		Key:        "1F6C0-1F3FB",
		Value:      "🛀🏻",
		Descriptor: "person taking bath: light skin tone",
	},
	"1F6C0-1F3FC": {
		Key:        "1F6C0-1F3FC",
		Value:      "🛀🏼",
		Descriptor: "person taking bath: medium-light skin tone",
	},
	"1F6C0-1F3FD": {
		Key:        "1F6C0-1F3FD",
		Value:      "🛀🏽",
		Descriptor: "person taking bath: medium skin tone",
	},
	"1F6C0-1F3FE": {
		Key:        "1F6C0-1F3FE",
		Value:      "🛀🏾",
		Descriptor: "person taking bath: medium-dark skin tone",
	},
	"1F6C0-1F3FF": {
		Key:        "1F6C0-1F3FF",
		Value:      "🛀🏿",
		Descriptor: "person taking bath: dark skin tone",
	},
	"1F6CC": {
		Key:        "1F6CC",
		Value:      "🛌",
		Descriptor: "person in bed",
	},
	"1F6CC-1F3FB": {
		Key:        "1F6CC-1F3FB",
		Value:      "🛌🏻",
		Descriptor: "person in bed: light skin tone",
	},
	"1F6CC-1F3FC": {
		Key:        "1F6CC-1F3FC",
		Value:      "🛌🏼",
		Descriptor: "person in bed: medium-light skin tone",
	},
	"1F6CC-1F3FD": {
		Key:        "1F6CC-1F3FD",
		Value:      "🛌🏽",
		Descriptor: "person in bed: medium skin tone",
	},
	"1F6CC-1F3FE": {
		Key:        "1F6CC-1F3FE",
		Value:      "🛌🏾",
		Descriptor: "person in bed: medium-dark skin tone",
	},
	"1F6CC-1F3FF": {
		Key:        "1F6CC-1F3FF",
		Value:      "🛌🏿",
		Descriptor: "person in bed: dark skin tone",
	},
	"1F9D1-200D-1F91D-200D-1F9D1": {
		Key:        "1F9D1-200D-1F91D-200D-1F9D1",
		Value:      "🧑‍🤝‍🧑",
		Descriptor: "people holding hands",
	},
	"1F9D1-1F3FB-200D-1F91D-200D-1F9D1-1F3FB": {
		Key:        "1F9D1-1F3FB-200D-1F91D-200D-1F9D1-1F3FB",
		Value:      "🧑🏻‍🤝‍🧑🏻",
		Descriptor: "people holding hands: light skin tone",
	},
	"1F9D1-1F3FB-200D-1F91D-200D-1F9D1-1F3FC": {
		Key:        "1F9D1-1F3FB-200D-1F91D-200D-1F9D1-1F3FC",
		Value:      "🧑🏻‍🤝‍🧑🏼",
		Descriptor: "people holding hands: light skin tone, medium-light skin tone",
	},
	"1F9D1-1F3FB-200D-1F91D-200D-1F9D1-1F3FD": {
		Key:        "1F9D1-1F3FB-200D-1F91D-200D-1F9D1-1F3FD",
		Value:      "🧑🏻‍🤝‍🧑🏽",
		Descriptor: "people holding hands: light skin tone, medium skin tone",
	},
	"1F9D1-1F3FB-200D-1F91D-200D-1F9D1-1F3FE": {
		Key:        "1F9D1-1F3FB-200D-1F91D-200D-1F9D1-1F3FE",
		Value:      "🧑🏻‍🤝‍🧑🏾",
		Descriptor: "people holding hands: light skin tone, medium-dark skin tone",
	},
	"1F9D1-1F3FB-200D-1F91D-200D-1F9D1-1F3FF": {
		Key:        "1F9D1-1F3FB-200D-1F91D-200D-1F9D1-1F3FF",
		Value:      "🧑🏻‍🤝‍🧑🏿",
		Descriptor: "people holding hands: light skin tone, dark skin tone",
	},
	"1F9D1-1F3FC-200D-1F91D-200D-1F9D1-1F3FB": {
		Key:        "1F9D1-1F3FC-200D-1F91D-200D-1F9D1-1F3FB",
		Value:      "🧑🏼‍🤝‍🧑🏻",
		Descriptor: "people holding hands: medium-light skin tone, light skin tone",
	},
	"1F9D1-1F3FC-200D-1F91D-200D-1F9D1-1F3FC": {
		Key:        "1F9D1-1F3FC-200D-1F91D-200D-1F9D1-1F3FC",
		Value:      "🧑🏼‍🤝‍🧑🏼",
		Descriptor: "people holding hands: medium-light skin tone",
	},
	"1F9D1-1F3FC-200D-1F91D-200D-1F9D1-1F3FD": {
		Key:        "1F9D1-1F3FC-200D-1F91D-200D-1F9D1-1F3FD",
		Value:      "🧑🏼‍🤝‍🧑🏽",
		Descriptor: "people holding hands: medium-light skin tone, medium skin tone",
	},
	"1F9D1-1F3FC-200D-1F91D-200D-1F9D1-1F3FE": {
		Key:        "1F9D1-1F3FC-200D-1F91D-200D-1F9D1-1F3FE",
		Value:      "🧑🏼‍🤝‍🧑🏾",
		Descriptor: "people holding hands: medium-light skin tone, medium-dark skin tone",
	},
	"1F9D1-1F3FC-200D-1F91D-200D-1F9D1-1F3FF": {
		Key:        "1F9D1-1F3FC-200D-1F91D-200D-1F9D1-1F3FF",
		Value:      "🧑🏼‍🤝‍🧑🏿",
		Descriptor: "people holding hands: medium-light skin tone, dark skin tone",
	},
	"1F9D1-1F3FD-200D-1F91D-200D-1F9D1-1F3FB": {
		Key:        "1F9D1-1F3FD-200D-1F91D-200D-1F9D1-1F3FB",
		Value:      "🧑🏽‍🤝‍🧑🏻",
		Descriptor: "people holding hands: medium skin tone, light skin tone",
	},
	"1F9D1-1F3FD-200D-1F91D-200D-1F9D1-1F3FC": {
		Key:        "1F9D1-1F3FD-200D-1F91D-200D-1F9D1-1F3FC",
		Value:      "🧑🏽‍🤝‍🧑🏼",
		Descriptor: "people holding hands: medium skin tone, medium-light skin tone",
	},
	"1F9D1-1F3FD-200D-1F91D-200D-1F9D1-1F3FD": {
		Key:        "1F9D1-1F3FD-200D-1F91D-200D-1F9D1-1F3FD",
		Value:      "🧑🏽‍🤝‍🧑🏽",
		Descriptor: "people holding hands: medium skin tone",
	},
	"1F9D1-1F3FD-200D-1F91D-200D-1F9D1-1F3FE": {
		Key:        "1F9D1-1F3FD-200D-1F91D-200D-1F9D1-1F3FE",
		Value:      "🧑🏽‍🤝‍🧑🏾",
		Descriptor: "people holding hands: medium skin tone, medium-dark skin tone",
	},
	"1F9D1-1F3FD-200D-1F91D-200D-1F9D1-1F3FF": {
		Key:        "1F9D1-1F3FD-200D-1F91D-200D-1F9D1-1F3FF",
		Value:      "🧑🏽‍🤝‍🧑🏿",
		Descriptor: "people holding hands: medium skin tone, dark skin tone",
	},
	"1F9D1-1F3FE-200D-1F91D-200D-1F9D1-1F3FB": {
		Key:        "1F9D1-1F3FE-200D-1F91D-200D-1F9D1-1F3FB",
		Value:      "🧑🏾‍🤝‍🧑🏻",
		Descriptor: "people holding hands: medium-dark skin tone, light skin tone",
	},
	"1F9D1-1F3FE-200D-1F91D-200D-1F9D1-1F3FC": {
		Key:        "1F9D1-1F3FE-200D-1F91D-200D-1F9D1-1F3FC",
		Value:      "🧑🏾‍🤝‍🧑🏼",
		Descriptor: "people holding hands: medium-dark skin tone, medium-light skin tone",
	},
	"1F9D1-1F3FE-200D-1F91D-200D-1F9D1-1F3FD": {
		Key:        "1F9D1-1F3FE-200D-1F91D-200D-1F9D1-1F3FD",
		Value:      "🧑🏾‍🤝‍🧑🏽",
		Descriptor: "people holding hands: medium-dark skin tone, medium skin tone",
	},
	"1F9D1-1F3FE-200D-1F91D-200D-1F9D1-1F3FE": {
		Key:        "1F9D1-1F3FE-200D-1F91D-200D-1F9D1-1F3FE",
		Value:      "🧑🏾‍🤝‍🧑🏾",
		Descriptor: "people holding hands: medium-dark skin tone",
	},
	"1F9D1-1F3FE-200D-1F91D-200D-1F9D1-1F3FF": {
		Key:        "1F9D1-1F3FE-200D-1F91D-200D-1F9D1-1F3FF",
		Value:      "🧑🏾‍🤝‍🧑🏿",
		Descriptor: "people holding hands: medium-dark skin tone, dark skin tone",
	},
	"1F9D1-1F3FF-200D-1F91D-200D-1F9D1-1F3FB": {
		Key:        "1F9D1-1F3FF-200D-1F91D-200D-1F9D1-1F3FB",
		Value:      "🧑🏿‍🤝‍🧑🏻",
		Descriptor: "people holding hands: dark skin tone, light skin tone",
	},
	"1F9D1-1F3FF-200D-1F91D-200D-1F9D1-1F3FC": {
		Key:        "1F9D1-1F3FF-200D-1F91D-200D-1F9D1-1F3FC",
		Value:      "🧑🏿‍🤝‍🧑🏼",
		Descriptor: "people holding hands: dark skin tone, medium-light skin tone",
	},
	"1F9D1-1F3FF-200D-1F91D-200D-1F9D1-1F3FD": {
		Key:        "1F9D1-1F3FF-200D-1F91D-200D-1F9D1-1F3FD",
		Value:      "🧑🏿‍🤝‍🧑🏽",
		Descriptor: "people holding hands: dark skin tone, medium skin tone",
	},
	"1F9D1-1F3FF-200D-1F91D-200D-1F9D1-1F3FE": {
		Key:        "1F9D1-1F3FF-200D-1F91D-200D-1F9D1-1F3FE",
		Value:      "🧑🏿‍🤝‍🧑🏾",
		Descriptor: "people holding hands: dark skin tone, medium-dark skin tone",
	},
	"1F9D1-1F3FF-200D-1F91D-200D-1F9D1-1F3FF": {
		Key:        "1F9D1-1F3FF-200D-1F91D-200D-1F9D1-1F3FF",
		Value:      "🧑🏿‍🤝‍🧑🏿",
		Descriptor: "people holding hands: dark skin tone",
	},
	"1F46D": {
		Key:        "1F46D",
		Value:      "👭",
		Descriptor: "women holding hands",
	},
	"1F46D-1F3FB": {
		Key:        "1F46D-1F3FB",
		Value:      "👭🏻",
		Descriptor: "women holding hands: light skin tone",
	},
	"1F469-1F3FB-200D-1F91D-200D-1F469-1F3FC": {
		Key:        "1F469-1F3FB-200D-1F91D-200D-1F469-1F3FC",
		Value:      "👩🏻‍🤝‍👩🏼",
		Descriptor: "women holding hands: light skin tone, medium-light skin tone",
	},
	"1F469-1F3FB-200D-1F91D-200D-1F469-1F3FD": {
		Key:        "1F469-1F3FB-200D-1F91D-200D-1F469-1F3FD",
		Value:      "👩🏻‍🤝‍👩🏽",
		Descriptor: "women holding hands: light skin tone, medium skin tone",
	},
	"1F469-1F3FB-200D-1F91D-200D-1F469-1F3FE": {
		Key:        "1F469-1F3FB-200D-1F91D-200D-1F469-1F3FE",
		Value:      "👩🏻‍🤝‍👩🏾",
		Descriptor: "women holding hands: light skin tone, medium-dark skin tone",
	},
	"1F469-1F3FB-200D-1F91D-200D-1F469-1F3FF": {
		Key:        "1F469-1F3FB-200D-1F91D-200D-1F469-1F3FF",
		Value:      "👩🏻‍🤝‍👩🏿",
		Descriptor: "women holding hands: light skin tone, dark skin tone",
	},
	"1F469-1F3FC-200D-1F91D-200D-1F469-1F3FB": {
		Key:        "1F469-1F3FC-200D-1F91D-200D-1F469-1F3FB",
		Value:      "👩🏼‍🤝‍👩🏻",
		Descriptor: "women holding hands: medium-light skin tone, light skin tone",
	},
	"1F46D-1F3FC": {
		Key:        "1F46D-1F3FC",
		Value:      "👭🏼",
		Descriptor: "women holding hands: medium-light skin tone",
	},
	"1F469-1F3FC-200D-1F91D-200D-1F469-1F3FD": {
		Key:        "1F469-1F3FC-200D-1F91D-200D-1F469-1F3FD",
		Value:      "👩🏼‍🤝‍👩🏽",
		Descriptor: "women holding hands: medium-light skin tone, medium skin tone",
	},
	"1F469-1F3FC-200D-1F91D-200D-1F469-1F3FE": {
		Key:        "1F469-1F3FC-200D-1F91D-200D-1F469-1F3FE",
		Value:      "👩🏼‍🤝‍👩🏾",
		Descriptor: "women holding hands: medium-light skin tone, medium-dark skin tone",
	},
	"1F469-1F3FC-200D-1F91D-200D-1F469-1F3FF": {
		Key:        "1F469-1F3FC-200D-1F91D-200D-1F469-1F3FF",
		Value:      "👩🏼‍🤝‍👩🏿",
		Descriptor: "women holding hands: medium-light skin tone, dark skin tone",
	},
	"1F469-1F3FD-200D-1F91D-200D-1F469-1F3FB": {
		Key:        "1F469-1F3FD-200D-1F91D-200D-1F469-1F3FB",
		Value:      "👩🏽‍🤝‍👩🏻",
		Descriptor: "women holding hands: medium skin tone, light skin tone",
	},
	"1F469-1F3FD-200D-1F91D-200D-1F469-1F3FC": {
		Key:        "1F469-1F3FD-200D-1F91D-200D-1F469-1F3FC",
		Value:      "👩🏽‍🤝‍👩🏼",
		Descriptor: "women holding hands: medium skin tone, medium-light skin tone",
	},
	"1F46D-1F3FD": {
		Key:        "1F46D-1F3FD",
		Value:      "👭🏽",
		Descriptor: "women holding hands: medium skin tone",
	},
	"1F469-1F3FD-200D-1F91D-200D-1F469-1F3FE": {
		Key:        "1F469-1F3FD-200D-1F91D-200D-1F469-1F3FE",
		Value:      "👩🏽‍🤝‍👩🏾",
		Descriptor: "women holding hands: medium skin tone, medium-dark skin tone",
	},
	"1F469-1F3FD-200D-1F91D-200D-1F469-1F3FF": {
		Key:        "1F469-1F3FD-200D-1F91D-200D-1F469-1F3FF",
		Value:      "👩🏽‍🤝‍👩🏿",
		Descriptor: "women holding hands: medium skin tone, dark skin tone",
	},
	"1F469-1F3FE-200D-1F91D-200D-1F469-1F3FB": {
		Key:        "1F469-1F3FE-200D-1F91D-200D-1F469-1F3FB",
		Value:      "👩🏾‍🤝‍👩🏻",
		Descriptor: "women holding hands: medium-dark skin tone, light skin tone",
	},
	"1F469-1F3FE-200D-1F91D-200D-1F469-1F3FC": {
		Key:        "1F469-1F3FE-200D-1F91D-200D-1F469-1F3FC",
		Value:      "👩🏾‍🤝‍👩🏼",
		Descriptor: "women holding hands: medium-dark skin tone, medium-light skin tone",
	},
	"1F469-1F3FE-200D-1F91D-200D-1F469-1F3FD": {
		Key:        "1F469-1F3FE-200D-1F91D-200D-1F469-1F3FD",
		Value:      "👩🏾‍🤝‍👩🏽",
		Descriptor: "women holding hands: medium-dark skin tone, medium skin tone",
	},
	"1F46D-1F3FE": {
		Key:        "1F46D-1F3FE",
		Value:      "👭🏾",
		Descriptor: "women holding hands: medium-dark skin tone",
	},
	"1F469-1F3FE-200D-1F91D-200D-1F469-1F3FF": {
		Key:        "1F469-1F3FE-200D-1F91D-200D-1F469-1F3FF",
		Value:      "👩🏾‍🤝‍👩🏿",
		Descriptor: "women holding hands: medium-dark skin tone, dark skin tone",
	},
	"1F469-1F3FF-200D-1F91D-200D-1F469-1F3FB": {
		Key:        "1F469-1F3FF-200D-1F91D-200D-1F469-1F3FB",
		Value:      "👩🏿‍🤝‍👩🏻",
		Descriptor: "women holding hands: dark skin tone, light skin tone",
	},
	"1F469-1F3FF-200D-1F91D-200D-1F469-1F3FC": {
		Key:        "1F469-1F3FF-200D-1F91D-200D-1F469-1F3FC",
		Value:      "👩🏿‍🤝‍👩🏼",
		Descriptor: "women holding hands: dark skin tone, medium-light skin tone",
	},
	"1F469-1F3FF-200D-1F91D-200D-1F469-1F3FD": {
		Key:        "1F469-1F3FF-200D-1F91D-200D-1F469-1F3FD",
		Value:      "👩🏿‍🤝‍👩🏽",
		Descriptor: "women holding hands: dark skin tone, medium skin tone",
	},
	"1F469-1F3FF-200D-1F91D-200D-1F469-1F3FE": {
		Key:        "1F469-1F3FF-200D-1F91D-200D-1F469-1F3FE",
		Value:      "👩🏿‍🤝‍👩🏾",
		Descriptor: "women holding hands: dark skin tone, medium-dark skin tone",
	},
	"1F46D-1F3FF": {
		Key:        "1F46D-1F3FF",
		Value:      "👭🏿",
		Descriptor: "women holding hands: dark skin tone",
	},
	"1F46B": {
		Key:        "1F46B",
		Value:      "👫",
		Descriptor: "woman and man holding hands",
	},
	"1F46B-1F3FB": {
		Key:        "1F46B-1F3FB",
		Value:      "👫🏻",
		Descriptor: "woman and man holding hands: light skin tone",
	},
	"1F469-1F3FB-200D-1F91D-200D-1F468-1F3FC": {
		Key:        "1F469-1F3FB-200D-1F91D-200D-1F468-1F3FC",
		Value:      "👩🏻‍🤝‍👨🏼",
		Descriptor: "woman and man holding hands: light skin tone, medium-light skin tone",
	},
	"1F469-1F3FB-200D-1F91D-200D-1F468-1F3FD": {
		Key:        "1F469-1F3FB-200D-1F91D-200D-1F468-1F3FD",
		Value:      "👩🏻‍🤝‍👨🏽",
		Descriptor: "woman and man holding hands: light skin tone, medium skin tone",
	},
	"1F469-1F3FB-200D-1F91D-200D-1F468-1F3FE": {
		Key:        "1F469-1F3FB-200D-1F91D-200D-1F468-1F3FE",
		Value:      "👩🏻‍🤝‍👨🏾",
		Descriptor: "woman and man holding hands: light skin tone, medium-dark skin tone",
	},
	"1F469-1F3FB-200D-1F91D-200D-1F468-1F3FF": {
		Key:        "1F469-1F3FB-200D-1F91D-200D-1F468-1F3FF",
		Value:      "👩🏻‍🤝‍👨🏿",
		Descriptor: "woman and man holding hands: light skin tone, dark skin tone",
	},
	"1F469-1F3FC-200D-1F91D-200D-1F468-1F3FB": {
		Key:        "1F469-1F3FC-200D-1F91D-200D-1F468-1F3FB",
		Value:      "👩🏼‍🤝‍👨🏻",
		Descriptor: "woman and man holding hands: medium-light skin tone, light skin tone",
	},
	"1F46B-1F3FC": {
		Key:        "1F46B-1F3FC",
		Value:      "👫🏼",
		Descriptor: "woman and man holding hands: medium-light skin tone",
	},
	"1F469-1F3FC-200D-1F91D-200D-1F468-1F3FD": {
		Key:        "1F469-1F3FC-200D-1F91D-200D-1F468-1F3FD",
		Value:      "👩🏼‍🤝‍👨🏽",
		Descriptor: "woman and man holding hands: medium-light skin tone, medium skin tone",
	},
	"1F469-1F3FC-200D-1F91D-200D-1F468-1F3FE": {
		Key:        "1F469-1F3FC-200D-1F91D-200D-1F468-1F3FE",
		Value:      "👩🏼‍🤝‍👨🏾",
		Descriptor: "woman and man holding hands: medium-light skin tone, medium-dark skin tone",
	},
	"1F469-1F3FC-200D-1F91D-200D-1F468-1F3FF": {
		Key:        "1F469-1F3FC-200D-1F91D-200D-1F468-1F3FF",
		Value:      "👩🏼‍🤝‍👨🏿",
		Descriptor: "woman and man holding hands: medium-light skin tone, dark skin tone",
	},
	"1F469-1F3FD-200D-1F91D-200D-1F468-1F3FB": {
		Key:        "1F469-1F3FD-200D-1F91D-200D-1F468-1F3FB",
		Value:      "👩🏽‍🤝‍👨🏻",
		Descriptor: "woman and man holding hands: medium skin tone, light skin tone",
	},
	"1F469-1F3FD-200D-1F91D-200D-1F468-1F3FC": {
		Key:        "1F469-1F3FD-200D-1F91D-200D-1F468-1F3FC",
		Value:      "👩🏽‍🤝‍👨🏼",
		Descriptor: "woman and man holding hands: medium skin tone, medium-light skin tone",
	},
	"1F46B-1F3FD": {
		Key:        "1F46B-1F3FD",
		Value:      "👫🏽",
		Descriptor: "woman and man holding hands: medium skin tone",
	},
	"1F469-1F3FD-200D-1F91D-200D-1F468-1F3FE": {
		Key:        "1F469-1F3FD-200D-1F91D-200D-1F468-1F3FE",
		Value:      "👩🏽‍🤝‍👨🏾",
		Descriptor: "woman and man holding hands: medium skin tone, medium-dark skin tone",
	},
	"1F469-1F3FD-200D-1F91D-200D-1F468-1F3FF": {
		Key:        "1F469-1F3FD-200D-1F91D-200D-1F468-1F3FF",
		Value:      "👩🏽‍🤝‍👨🏿",
		Descriptor: "woman and man holding hands: medium skin tone, dark skin tone",
	},
	"1F469-1F3FE-200D-1F91D-200D-1F468-1F3FB": {
		Key:        "1F469-1F3FE-200D-1F91D-200D-1F468-1F3FB",
		Value:      "👩🏾‍🤝‍👨🏻",
		Descriptor: "woman and man holding hands: medium-dark skin tone, light skin tone",
	},
	"1F469-1F3FE-200D-1F91D-200D-1F468-1F3FC": {
		Key:        "1F469-1F3FE-200D-1F91D-200D-1F468-1F3FC",
		Value:      "👩🏾‍🤝‍👨🏼",
		Descriptor: "woman and man holding hands: medium-dark skin tone, medium-light skin tone",
	},
	"1F469-1F3FE-200D-1F91D-200D-1F468-1F3FD": {
		Key:        "1F469-1F3FE-200D-1F91D-200D-1F468-1F3FD",
		Value:      "👩🏾‍🤝‍👨🏽",
		Descriptor: "woman and man holding hands: medium-dark skin tone, medium skin tone",
	},
	"1F46B-1F3FE": {
		Key:        "1F46B-1F3FE",
		Value:      "👫🏾",
		Descriptor: "woman and man holding hands: medium-dark skin tone",
	},
	"1F469-1F3FE-200D-1F91D-200D-1F468-1F3FF": {
		Key:        "1F469-1F3FE-200D-1F91D-200D-1F468-1F3FF",
		Value:      "👩🏾‍🤝‍👨🏿",
		Descriptor: "woman and man holding hands: medium-dark skin tone, dark skin tone",
	},
	"1F469-1F3FF-200D-1F91D-200D-1F468-1F3FB": {
		Key:        "1F469-1F3FF-200D-1F91D-200D-1F468-1F3FB",
		Value:      "👩🏿‍🤝‍👨🏻",
		Descriptor: "woman and man holding hands: dark skin tone, light skin tone",
	},
	"1F469-1F3FF-200D-1F91D-200D-1F468-1F3FC": {
		Key:        "1F469-1F3FF-200D-1F91D-200D-1F468-1F3FC",
		Value:      "👩🏿‍🤝‍👨🏼",
		Descriptor: "woman and man holding hands: dark skin tone, medium-light skin tone",
	},
	"1F469-1F3FF-200D-1F91D-200D-1F468-1F3FD": {
		Key:        "1F469-1F3FF-200D-1F91D-200D-1F468-1F3FD",
		Value:      "👩🏿‍🤝‍👨🏽",
		Descriptor: "woman and man holding hands: dark skin tone, medium skin tone",
	},
	"1F469-1F3FF-200D-1F91D-200D-1F468-1F3FE": {
		Key:        "1F469-1F3FF-200D-1F91D-200D-1F468-1F3FE",
		Value:      "👩🏿‍🤝‍👨🏾",
		Descriptor: "woman and man holding hands: dark skin tone, medium-dark skin tone",
	},
	"1F46B-1F3FF": {
		Key:        "1F46B-1F3FF",
		Value:      "👫🏿",
		Descriptor: "woman and man holding hands: dark skin tone",
	},
	"1F46C": {
		Key:        "1F46C",
		Value:      "👬",
		Descriptor: "men holding hands",
	},
	"1F46C-1F3FB": {
		Key:        "1F46C-1F3FB",
		Value:      "👬🏻",
		Descriptor: "men holding hands: light skin tone",
	},
	"1F468-1F3FB-200D-1F91D-200D-1F468-1F3FC": {
		Key:        "1F468-1F3FB-200D-1F91D-200D-1F468-1F3FC",
		Value:      "👨🏻‍🤝‍👨🏼",
		Descriptor: "men holding hands: light skin tone, medium-light skin tone",
	},
	"1F468-1F3FB-200D-1F91D-200D-1F468-1F3FD": {
		Key:        "1F468-1F3FB-200D-1F91D-200D-1F468-1F3FD",
		Value:      "👨🏻‍🤝‍👨🏽",
		Descriptor: "men holding hands: light skin tone, medium skin tone",
	},
	"1F468-1F3FB-200D-1F91D-200D-1F468-1F3FE": {
		Key:        "1F468-1F3FB-200D-1F91D-200D-1F468-1F3FE",
		Value:      "👨🏻‍🤝‍👨🏾",
		Descriptor: "men holding hands: light skin tone, medium-dark skin tone",
	},
	"1F468-1F3FB-200D-1F91D-200D-1F468-1F3FF": {
		Key:        "1F468-1F3FB-200D-1F91D-200D-1F468-1F3FF",
		Value:      "👨🏻‍🤝‍👨🏿",
		Descriptor: "men holding hands: light skin tone, dark skin tone",
	},
	"1F468-1F3FC-200D-1F91D-200D-1F468-1F3FB": {
		Key:        "1F468-1F3FC-200D-1F91D-200D-1F468-1F3FB",
		Value:      "👨🏼‍🤝‍👨🏻",
		Descriptor: "men holding hands: medium-light skin tone, light skin tone",
	},
	"1F46C-1F3FC": {
		Key:        "1F46C-1F3FC",
		Value:      "👬🏼",
		Descriptor: "men holding hands: medium-light skin tone",
	},
	"1F468-1F3FC-200D-1F91D-200D-1F468-1F3FD": {
		Key:        "1F468-1F3FC-200D-1F91D-200D-1F468-1F3FD",
		Value:      "👨🏼‍🤝‍👨🏽",
		Descriptor: "men holding hands: medium-light skin tone, medium skin tone",
	},
	"1F468-1F3FC-200D-1F91D-200D-1F468-1F3FE": {
		Key:        "1F468-1F3FC-200D-1F91D-200D-1F468-1F3FE",
		Value:      "👨🏼‍🤝‍👨🏾",
		Descriptor: "men holding hands: medium-light skin tone, medium-dark skin tone",
	},
	"1F468-1F3FC-200D-1F91D-200D-1F468-1F3FF": {
		Key:        "1F468-1F3FC-200D-1F91D-200D-1F468-1F3FF",
		Value:      "👨🏼‍🤝‍👨🏿",
		Descriptor: "men holding hands: medium-light skin tone, dark skin tone",
	},
	"1F468-1F3FD-200D-1F91D-200D-1F468-1F3FB": {
		Key:        "1F468-1F3FD-200D-1F91D-200D-1F468-1F3FB",
		Value:      "👨🏽‍🤝‍👨🏻",
		Descriptor: "men holding hands: medium skin tone, light skin tone",
	},
	"1F468-1F3FD-200D-1F91D-200D-1F468-1F3FC": {
		Key:        "1F468-1F3FD-200D-1F91D-200D-1F468-1F3FC",
		Value:      "👨🏽‍🤝‍👨🏼",
		Descriptor: "men holding hands: medium skin tone, medium-light skin tone",
	},
	"1F46C-1F3FD": {
		Key:        "1F46C-1F3FD",
		Value:      "👬🏽",
		Descriptor: "men holding hands: medium skin tone",
	},
	"1F468-1F3FD-200D-1F91D-200D-1F468-1F3FE": {
		Key:        "1F468-1F3FD-200D-1F91D-200D-1F468-1F3FE",
		Value:      "👨🏽‍🤝‍👨🏾",
		Descriptor: "men holding hands: medium skin tone, medium-dark skin tone",
	},
	"1F468-1F3FD-200D-1F91D-200D-1F468-1F3FF": {
		Key:        "1F468-1F3FD-200D-1F91D-200D-1F468-1F3FF",
		Value:      "👨🏽‍🤝‍👨🏿",
		Descriptor: "men holding hands: medium skin tone, dark skin tone",
	},
	"1F468-1F3FE-200D-1F91D-200D-1F468-1F3FB": {
		Key:        "1F468-1F3FE-200D-1F91D-200D-1F468-1F3FB",
		Value:      "👨🏾‍🤝‍👨🏻",
		Descriptor: "men holding hands: medium-dark skin tone, light skin tone",
	},
	"1F468-1F3FE-200D-1F91D-200D-1F468-1F3FC": {
		Key:        "1F468-1F3FE-200D-1F91D-200D-1F468-1F3FC",
		Value:      "👨🏾‍🤝‍👨🏼",
		Descriptor: "men holding hands: medium-dark skin tone, medium-light skin tone",
	},
	"1F468-1F3FE-200D-1F91D-200D-1F468-1F3FD": {
		Key:        "1F468-1F3FE-200D-1F91D-200D-1F468-1F3FD",
		Value:      "👨🏾‍🤝‍👨🏽",
		Descriptor: "men holding hands: medium-dark skin tone, medium skin tone",
	},
	"1F46C-1F3FE": {
		Key:        "1F46C-1F3FE",
		Value:      "👬🏾",
		Descriptor: "men holding hands: medium-dark skin tone",
	},
	"1F468-1F3FE-200D-1F91D-200D-1F468-1F3FF": {
		Key:        "1F468-1F3FE-200D-1F91D-200D-1F468-1F3FF",
		Value:      "👨🏾‍🤝‍👨🏿",
		Descriptor: "men holding hands: medium-dark skin tone, dark skin tone",
	},
	"1F468-1F3FF-200D-1F91D-200D-1F468-1F3FB": {
		Key:        "1F468-1F3FF-200D-1F91D-200D-1F468-1F3FB",
		Value:      "👨🏿‍🤝‍👨🏻",
		Descriptor: "men holding hands: dark skin tone, light skin tone",
	},
	"1F468-1F3FF-200D-1F91D-200D-1F468-1F3FC": {
		Key:        "1F468-1F3FF-200D-1F91D-200D-1F468-1F3FC",
		Value:      "👨🏿‍🤝‍👨🏼",
		Descriptor: "men holding hands: dark skin tone, medium-light skin tone",
	},
	"1F468-1F3FF-200D-1F91D-200D-1F468-1F3FD": {
		Key:        "1F468-1F3FF-200D-1F91D-200D-1F468-1F3FD",
		Value:      "👨🏿‍🤝‍👨🏽",
		Descriptor: "men holding hands: dark skin tone, medium skin tone",
	},
	"1F468-1F3FF-200D-1F91D-200D-1F468-1F3FE": {
		Key:        "1F468-1F3FF-200D-1F91D-200D-1F468-1F3FE",
		Value:      "👨🏿‍🤝‍👨🏾",
		Descriptor: "men holding hands: dark skin tone, medium-dark skin tone",
	},
	"1F46C-1F3FF": {
		Key:        "1F46C-1F3FF",
		Value:      "👬🏿",
		Descriptor: "men holding hands: dark skin tone",
	},
	"1F48F": {
		Key:        "1F48F",
		Value:      "💏",
		Descriptor: "kiss",
	},
	"1F48F-1F3FB": {
		Key:        "1F48F-1F3FB",
		Value:      "💏🏻",
		Descriptor: "kiss: light skin tone",
	},
	"1F48F-1F3FC": {
		Key:        "1F48F-1F3FC",
		Value:      "💏🏼",
		Descriptor: "kiss: medium-light skin tone",
	},
	"1F48F-1F3FD": {
		Key:        "1F48F-1F3FD",
		Value:      "💏🏽",
		Descriptor: "kiss: medium skin tone",
	},
	"1F48F-1F3FE": {
		Key:        "1F48F-1F3FE",
		Value:      "💏🏾",
		Descriptor: "kiss: medium-dark skin tone",
	},
	"1F48F-1F3FF": {
		Key:        "1F48F-1F3FF",
		Value:      "💏🏿",
		Descriptor: "kiss: dark skin tone",
	},
	"1F9D1-1F3FB-200D-2764-FE0F-200D-1F48B-200D-1F9D1-1F3FC": {
		Key:        "1F9D1-1F3FB-200D-2764-FE0F-200D-1F48B-200D-1F9D1-1F3FC",
		Value:      "🧑🏻‍❤️‍💋‍🧑🏼",
		Descriptor: "kiss: person, person, light skin tone, medium-light skin tone",
	},
	"1F9D1-1F3FB-200D-2764-FE0F-200D-1F48B-200D-1F9D1-1F3FD": {
		Key:        "1F9D1-1F3FB-200D-2764-FE0F-200D-1F48B-200D-1F9D1-1F3FD",
		Value:      "🧑🏻‍❤️‍💋‍🧑🏽",
		Descriptor: "kiss: person, person, light skin tone, medium skin tone",
	},
	"1F9D1-1F3FB-200D-2764-FE0F-200D-1F48B-200D-1F9D1-1F3FE": {
		Key:        "1F9D1-1F3FB-200D-2764-FE0F-200D-1F48B-200D-1F9D1-1F3FE",
		Value:      "🧑🏻‍❤️‍💋‍🧑🏾",
		Descriptor: "kiss: person, person, light skin tone, medium-dark skin tone",
	},
	"1F9D1-1F3FB-200D-2764-FE0F-200D-1F48B-200D-1F9D1-1F3FF": {
		Key:        "1F9D1-1F3FB-200D-2764-FE0F-200D-1F48B-200D-1F9D1-1F3FF",
		Value:      "🧑🏻‍❤️‍💋‍🧑🏿",
		Descriptor: "kiss: person, person, light skin tone, dark skin tone",
	},
	"1F9D1-1F3FC-200D-2764-FE0F-200D-1F48B-200D-1F9D1-1F3FB": {
		Key:        "1F9D1-1F3FC-200D-2764-FE0F-200D-1F48B-200D-1F9D1-1F3FB",
		Value:      "🧑🏼‍❤️‍💋‍🧑🏻",
		Descriptor: "kiss: person, person, medium-light skin tone, light skin tone",
	},
	"1F9D1-1F3FC-200D-2764-FE0F-200D-1F48B-200D-1F9D1-1F3FD": {
		Key:        "1F9D1-1F3FC-200D-2764-FE0F-200D-1F48B-200D-1F9D1-1F3FD",
		Value:      "🧑🏼‍❤️‍💋‍🧑🏽",
		Descriptor: "kiss: person, person, medium-light skin tone, medium skin tone",
	},
	"1F9D1-1F3FC-200D-2764-FE0F-200D-1F48B-200D-1F9D1-1F3FE": {
		Key:        "1F9D1-1F3FC-200D-2764-FE0F-200D-1F48B-200D-1F9D1-1F3FE",
		Value:      "🧑🏼‍❤️‍💋‍🧑🏾",
		Descriptor: "kiss: person, person, medium-light skin tone, medium-dark skin tone",
	},
	"1F9D1-1F3FC-200D-2764-FE0F-200D-1F48B-200D-1F9D1-1F3FF": {
		Key:        "1F9D1-1F3FC-200D-2764-FE0F-200D-1F48B-200D-1F9D1-1F3FF",
		Value:      "🧑🏼‍❤️‍💋‍🧑🏿",
		Descriptor: "kiss: person, person, medium-light skin tone, dark skin tone",
	},
	"1F9D1-1F3FD-200D-2764-FE0F-200D-1F48B-200D-1F9D1-1F3FB": {
		Key:        "1F9D1-1F3FD-200D-2764-FE0F-200D-1F48B-200D-1F9D1-1F3FB",
		Value:      "🧑🏽‍❤️‍💋‍🧑🏻",
		Descriptor: "kiss: person, person, medium skin tone, light skin tone",
	},
	"1F9D1-1F3FD-200D-2764-FE0F-200D-1F48B-200D-1F9D1-1F3FC": {
		Key:        "1F9D1-1F3FD-200D-2764-FE0F-200D-1F48B-200D-1F9D1-1F3FC",
		Value:      "🧑🏽‍❤️‍💋‍🧑🏼",
		Descriptor: "kiss: person, person, medium skin tone, medium-light skin tone",
	},
	"1F9D1-1F3FD-200D-2764-FE0F-200D-1F48B-200D-1F9D1-1F3FE": {
		Key:        "1F9D1-1F3FD-200D-2764-FE0F-200D-1F48B-200D-1F9D1-1F3FE",
		Value:      "🧑🏽‍❤️‍💋‍🧑🏾",
		Descriptor: "kiss: person, person, medium skin tone, medium-dark skin tone",
	},
	"1F9D1-1F3FD-200D-2764-FE0F-200D-1F48B-200D-1F9D1-1F3FF": {
		Key:        "1F9D1-1F3FD-200D-2764-FE0F-200D-1F48B-200D-1F9D1-1F3FF",
		Value:      "🧑🏽‍❤️‍💋‍🧑🏿",
		Descriptor: "kiss: person, person, medium skin tone, dark skin tone",
	},
	"1F9D1-1F3FE-200D-2764-FE0F-200D-1F48B-200D-1F9D1-1F3FB": {
		Key:        "1F9D1-1F3FE-200D-2764-FE0F-200D-1F48B-200D-1F9D1-1F3FB",
		Value:      "🧑🏾‍❤️‍💋‍🧑🏻",
		Descriptor: "kiss: person, person, medium-dark skin tone, light skin tone",
	},
	"1F9D1-1F3FE-200D-2764-FE0F-200D-1F48B-200D-1F9D1-1F3FC": {
		Key:        "1F9D1-1F3FE-200D-2764-FE0F-200D-1F48B-200D-1F9D1-1F3FC",
		Value:      "🧑🏾‍❤️‍💋‍🧑🏼",
		Descriptor: "kiss: person, person, medium-dark skin tone, medium-light skin tone",
	},
	"1F9D1-1F3FE-200D-2764-FE0F-200D-1F48B-200D-1F9D1-1F3FD": {
		Key:        "1F9D1-1F3FE-200D-2764-FE0F-200D-1F48B-200D-1F9D1-1F3FD",
		Value:      "🧑🏾‍❤️‍💋‍🧑🏽",
		Descriptor: "kiss: person, person, medium-dark skin tone, medium skin tone",
	},
	"1F9D1-1F3FE-200D-2764-FE0F-200D-1F48B-200D-1F9D1-1F3FF": {
		Key:        "1F9D1-1F3FE-200D-2764-FE0F-200D-1F48B-200D-1F9D1-1F3FF",
		Value:      "🧑🏾‍❤️‍💋‍🧑🏿",
		Descriptor: "kiss: person, person, medium-dark skin tone, dark skin tone",
	},
	"1F9D1-1F3FF-200D-2764-FE0F-200D-1F48B-200D-1F9D1-1F3FB": {
		Key:        "1F9D1-1F3FF-200D-2764-FE0F-200D-1F48B-200D-1F9D1-1F3FB",
		Value:      "🧑🏿‍❤️‍💋‍🧑🏻",
		Descriptor: "kiss: person, person, dark skin tone, light skin tone",
	},
	"1F9D1-1F3FF-200D-2764-FE0F-200D-1F48B-200D-1F9D1-1F3FC": {
		Key:        "1F9D1-1F3FF-200D-2764-FE0F-200D-1F48B-200D-1F9D1-1F3FC",
		Value:      "🧑🏿‍❤️‍💋‍🧑🏼",
		Descriptor: "kiss: person, person, dark skin tone, medium-light skin tone",
	},
	"1F9D1-1F3FF-200D-2764-FE0F-200D-1F48B-200D-1F9D1-1F3FD": {
		Key:        "1F9D1-1F3FF-200D-2764-FE0F-200D-1F48B-200D-1F9D1-1F3FD",
		Value:      "🧑🏿‍❤️‍💋‍🧑🏽",
		Descriptor: "kiss: person, person, dark skin tone, medium skin tone",
	},
	"1F9D1-1F3FF-200D-2764-FE0F-200D-1F48B-200D-1F9D1-1F3FE": {
		Key:        "1F9D1-1F3FF-200D-2764-FE0F-200D-1F48B-200D-1F9D1-1F3FE",
		Value:      "🧑🏿‍❤️‍💋‍🧑🏾",
		Descriptor: "kiss: person, person, dark skin tone, medium-dark skin tone",
	},
	"1F469-200D-2764-FE0F-200D-1F48B-200D-1F468": {
		Key:        "1F469-200D-2764-FE0F-200D-1F48B-200D-1F468",
		Value:      "👩‍❤️‍💋‍👨",
		Descriptor: "kiss: woman, man",
	},
	"1F469-1F3FB-200D-2764-FE0F-200D-1F48B-200D-1F468-1F3FB": {
		Key:        "1F469-1F3FB-200D-2764-FE0F-200D-1F48B-200D-1F468-1F3FB",
		Value:      "👩🏻‍❤️‍💋‍👨🏻",
		Descriptor: "kiss: woman, man, light skin tone",
	},
	"1F469-1F3FB-200D-2764-FE0F-200D-1F48B-200D-1F468-1F3FC": {
		Key:        "1F469-1F3FB-200D-2764-FE0F-200D-1F48B-200D-1F468-1F3FC",
		Value:      "👩🏻‍❤️‍💋‍👨🏼",
		Descriptor: "kiss: woman, man, light skin tone, medium-light skin tone",
	},
	"1F469-1F3FB-200D-2764-FE0F-200D-1F48B-200D-1F468-1F3FD": {
		Key:        "1F469-1F3FB-200D-2764-FE0F-200D-1F48B-200D-1F468-1F3FD",
		Value:      "👩🏻‍❤️‍💋‍👨🏽",
		Descriptor: "kiss: woman, man, light skin tone, medium skin tone",
	},
	"1F469-1F3FB-200D-2764-FE0F-200D-1F48B-200D-1F468-1F3FE": {
		Key:        "1F469-1F3FB-200D-2764-FE0F-200D-1F48B-200D-1F468-1F3FE",
		Value:      "👩🏻‍❤️‍💋‍👨🏾",
		Descriptor: "kiss: woman, man, light skin tone, medium-dark skin tone",
	},
	"1F469-1F3FB-200D-2764-FE0F-200D-1F48B-200D-1F468-1F3FF": {
		Key:        "1F469-1F3FB-200D-2764-FE0F-200D-1F48B-200D-1F468-1F3FF",
		Value:      "👩🏻‍❤️‍💋‍👨🏿",
		Descriptor: "kiss: woman, man, light skin tone, dark skin tone",
	},
	"1F469-1F3FC-200D-2764-FE0F-200D-1F48B-200D-1F468-1F3FB": {
		Key:        "1F469-1F3FC-200D-2764-FE0F-200D-1F48B-200D-1F468-1F3FB",
		Value:      "👩🏼‍❤️‍💋‍👨🏻",
		Descriptor: "kiss: woman, man, medium-light skin tone, light skin tone",
	},
	"1F469-1F3FC-200D-2764-FE0F-200D-1F48B-200D-1F468-1F3FC": {
		Key:        "1F469-1F3FC-200D-2764-FE0F-200D-1F48B-200D-1F468-1F3FC",
		Value:      "👩🏼‍❤️‍💋‍👨🏼",
		Descriptor: "kiss: woman, man, medium-light skin tone",
	},
	"1F469-1F3FC-200D-2764-FE0F-200D-1F48B-200D-1F468-1F3FD": {
		Key:        "1F469-1F3FC-200D-2764-FE0F-200D-1F48B-200D-1F468-1F3FD",
		Value:      "👩🏼‍❤️‍💋‍👨🏽",
		Descriptor: "kiss: woman, man, medium-light skin tone, medium skin tone",
	},
	"1F469-1F3FC-200D-2764-FE0F-200D-1F48B-200D-1F468-1F3FE": {
		Key:        "1F469-1F3FC-200D-2764-FE0F-200D-1F48B-200D-1F468-1F3FE",
		Value:      "👩🏼‍❤️‍💋‍👨🏾",
		Descriptor: "kiss: woman, man, medium-light skin tone, medium-dark skin tone",
	},
	"1F469-1F3FC-200D-2764-FE0F-200D-1F48B-200D-1F468-1F3FF": {
		Key:        "1F469-1F3FC-200D-2764-FE0F-200D-1F48B-200D-1F468-1F3FF",
		Value:      "👩🏼‍❤️‍💋‍👨🏿",
		Descriptor: "kiss: woman, man, medium-light skin tone, dark skin tone",
	},
	"1F469-1F3FD-200D-2764-FE0F-200D-1F48B-200D-1F468-1F3FB": {
		Key:        "1F469-1F3FD-200D-2764-FE0F-200D-1F48B-200D-1F468-1F3FB",
		Value:      "👩🏽‍❤️‍💋‍👨🏻",
		Descriptor: "kiss: woman, man, medium skin tone, light skin tone",
	},
	"1F469-1F3FD-200D-2764-FE0F-200D-1F48B-200D-1F468-1F3FC": {
		Key:        "1F469-1F3FD-200D-2764-FE0F-200D-1F48B-200D-1F468-1F3FC",
		Value:      "👩🏽‍❤️‍💋‍👨🏼",
		Descriptor: "kiss: woman, man, medium skin tone, medium-light skin tone",
	},
	"1F469-1F3FD-200D-2764-FE0F-200D-1F48B-200D-1F468-1F3FD": {
		Key:        "1F469-1F3FD-200D-2764-FE0F-200D-1F48B-200D-1F468-1F3FD",
		Value:      "👩🏽‍❤️‍💋‍👨🏽",
		Descriptor: "kiss: woman, man, medium skin tone",
	},
	"1F469-1F3FD-200D-2764-FE0F-200D-1F48B-200D-1F468-1F3FE": {
		Key:        "1F469-1F3FD-200D-2764-FE0F-200D-1F48B-200D-1F468-1F3FE",
		Value:      "👩🏽‍❤️‍💋‍👨🏾",
		Descriptor: "kiss: woman, man, medium skin tone, medium-dark skin tone",
	},
	"1F469-1F3FD-200D-2764-FE0F-200D-1F48B-200D-1F468-1F3FF": {
		Key:        "1F469-1F3FD-200D-2764-FE0F-200D-1F48B-200D-1F468-1F3FF",
		Value:      "👩🏽‍❤️‍💋‍👨🏿",
		Descriptor: "kiss: woman, man, medium skin tone, dark skin tone",
	},
	"1F469-1F3FE-200D-2764-FE0F-200D-1F48B-200D-1F468-1F3FB": {
		Key:        "1F469-1F3FE-200D-2764-FE0F-200D-1F48B-200D-1F468-1F3FB",
		Value:      "👩🏾‍❤️‍💋‍👨🏻",
		Descriptor: "kiss: woman, man, medium-dark skin tone, light skin tone",
	},
	"1F469-1F3FE-200D-2764-FE0F-200D-1F48B-200D-1F468-1F3FC": {
		Key:        "1F469-1F3FE-200D-2764-FE0F-200D-1F48B-200D-1F468-1F3FC",
		Value:      "👩🏾‍❤️‍💋‍👨🏼",
		Descriptor: "kiss: woman, man, medium-dark skin tone, medium-light skin tone",
	},
	"1F469-1F3FE-200D-2764-FE0F-200D-1F48B-200D-1F468-1F3FD": {
		Key:        "1F469-1F3FE-200D-2764-FE0F-200D-1F48B-200D-1F468-1F3FD",
		Value:      "👩🏾‍❤️‍💋‍👨🏽",
		Descriptor: "kiss: woman, man, medium-dark skin tone, medium skin tone",
	},
	"1F469-1F3FE-200D-2764-FE0F-200D-1F48B-200D-1F468-1F3FE": {
		Key:        "1F469-1F3FE-200D-2764-FE0F-200D-1F48B-200D-1F468-1F3FE",
		Value:      "👩🏾‍❤️‍💋‍👨🏾",
		Descriptor: "kiss: woman, man, medium-dark skin tone",
	},
	"1F469-1F3FE-200D-2764-FE0F-200D-1F48B-200D-1F468-1F3FF": {
		Key:        "1F469-1F3FE-200D-2764-FE0F-200D-1F48B-200D-1F468-1F3FF",
		Value:      "👩🏾‍❤️‍💋‍👨🏿",
		Descriptor: "kiss: woman, man, medium-dark skin tone, dark skin tone",
	},
	"1F469-1F3FF-200D-2764-FE0F-200D-1F48B-200D-1F468-1F3FB": {
		Key:        "1F469-1F3FF-200D-2764-FE0F-200D-1F48B-200D-1F468-1F3FB",
		Value:      "👩🏿‍❤️‍💋‍👨🏻",
		Descriptor: "kiss: woman, man, dark skin tone, light skin tone",
	},
	"1F469-1F3FF-200D-2764-FE0F-200D-1F48B-200D-1F468-1F3FC": {
		Key:        "1F469-1F3FF-200D-2764-FE0F-200D-1F48B-200D-1F468-1F3FC",
		Value:      "👩🏿‍❤️‍💋‍👨🏼",
		Descriptor: "kiss: woman, man, dark skin tone, medium-light skin tone",
	},
	"1F469-1F3FF-200D-2764-FE0F-200D-1F48B-200D-1F468-1F3FD": {
		Key:        "1F469-1F3FF-200D-2764-FE0F-200D-1F48B-200D-1F468-1F3FD",
		Value:      "👩🏿‍❤️‍💋‍👨🏽",
		Descriptor: "kiss: woman, man, dark skin tone, medium skin tone",
	},
	"1F469-1F3FF-200D-2764-FE0F-200D-1F48B-200D-1F468-1F3FE": {
		Key:        "1F469-1F3FF-200D-2764-FE0F-200D-1F48B-200D-1F468-1F3FE",
		Value:      "👩🏿‍❤️‍💋‍👨🏾",
		Descriptor: "kiss: woman, man, dark skin tone, medium-dark skin tone",
	},
	"1F469-1F3FF-200D-2764-FE0F-200D-1F48B-200D-1F468-1F3FF": {
		Key:        "1F469-1F3FF-200D-2764-FE0F-200D-1F48B-200D-1F468-1F3FF",
		Value:      "👩🏿‍❤️‍💋‍👨🏿",
		Descriptor: "kiss: woman, man, dark skin tone",
	},
	"1F468-200D-2764-FE0F-200D-1F48B-200D-1F468": {
		Key:        "1F468-200D-2764-FE0F-200D-1F48B-200D-1F468",
		Value:      "👨‍❤️‍💋‍👨",
		Descriptor: "kiss: man, man",
	},
	"1F468-1F3FB-200D-2764-FE0F-200D-1F48B-200D-1F468-1F3FB": {
		Key:        "1F468-1F3FB-200D-2764-FE0F-200D-1F48B-200D-1F468-1F3FB",
		Value:      "👨🏻‍❤️‍💋‍👨🏻",
		Descriptor: "kiss: man, man, light skin tone",
	},
	"1F468-1F3FB-200D-2764-FE0F-200D-1F48B-200D-1F468-1F3FC": {
		Key:        "1F468-1F3FB-200D-2764-FE0F-200D-1F48B-200D-1F468-1F3FC",
		Value:      "👨🏻‍❤️‍💋‍👨🏼",
		Descriptor: "kiss: man, man, light skin tone, medium-light skin tone",
	},
	"1F468-1F3FB-200D-2764-FE0F-200D-1F48B-200D-1F468-1F3FD": {
		Key:        "1F468-1F3FB-200D-2764-FE0F-200D-1F48B-200D-1F468-1F3FD",
		Value:      "👨🏻‍❤️‍💋‍👨🏽",
		Descriptor: "kiss: man, man, light skin tone, medium skin tone",
	},
	"1F468-1F3FB-200D-2764-FE0F-200D-1F48B-200D-1F468-1F3FE": {
		Key:        "1F468-1F3FB-200D-2764-FE0F-200D-1F48B-200D-1F468-1F3FE",
		Value:      "👨🏻‍❤️‍💋‍👨🏾",
		Descriptor: "kiss: man, man, light skin tone, medium-dark skin tone",
	},
	"1F468-1F3FB-200D-2764-FE0F-200D-1F48B-200D-1F468-1F3FF": {
		Key:        "1F468-1F3FB-200D-2764-FE0F-200D-1F48B-200D-1F468-1F3FF",
		Value:      "👨🏻‍❤️‍💋‍👨🏿",
		Descriptor: "kiss: man, man, light skin tone, dark skin tone",
	},
	"1F468-1F3FC-200D-2764-FE0F-200D-1F48B-200D-1F468-1F3FB": {
		Key:        "1F468-1F3FC-200D-2764-FE0F-200D-1F48B-200D-1F468-1F3FB",
		Value:      "👨🏼‍❤️‍💋‍👨🏻",
		Descriptor: "kiss: man, man, medium-light skin tone, light skin tone",
	},
	"1F468-1F3FC-200D-2764-FE0F-200D-1F48B-200D-1F468-1F3FC": {
		Key:        "1F468-1F3FC-200D-2764-FE0F-200D-1F48B-200D-1F468-1F3FC",
		Value:      "👨🏼‍❤️‍💋‍👨🏼",
		Descriptor: "kiss: man, man, medium-light skin tone",
	},
	"1F468-1F3FC-200D-2764-FE0F-200D-1F48B-200D-1F468-1F3FD": {
		Key:        "1F468-1F3FC-200D-2764-FE0F-200D-1F48B-200D-1F468-1F3FD",
		Value:      "👨🏼‍❤️‍💋‍👨🏽",
		Descriptor: "kiss: man, man, medium-light skin tone, medium skin tone",
	},
	"1F468-1F3FC-200D-2764-FE0F-200D-1F48B-200D-1F468-1F3FE": {
		Key:        "1F468-1F3FC-200D-2764-FE0F-200D-1F48B-200D-1F468-1F3FE",
		Value:      "👨🏼‍❤️‍💋‍👨🏾",
		Descriptor: "kiss: man, man, medium-light skin tone, medium-dark skin tone",
	},
	"1F468-1F3FC-200D-2764-FE0F-200D-1F48B-200D-1F468-1F3FF": {
		Key:        "1F468-1F3FC-200D-2764-FE0F-200D-1F48B-200D-1F468-1F3FF",
		Value:      "👨🏼‍❤️‍💋‍👨🏿",
		Descriptor: "kiss: man, man, medium-light skin tone, dark skin tone",
	},
	"1F468-1F3FD-200D-2764-FE0F-200D-1F48B-200D-1F468-1F3FB": {
		Key:        "1F468-1F3FD-200D-2764-FE0F-200D-1F48B-200D-1F468-1F3FB",
		Value:      "👨🏽‍❤️‍💋‍👨🏻",
		Descriptor: "kiss: man, man, medium skin tone, light skin tone",
	},
	"1F468-1F3FD-200D-2764-FE0F-200D-1F48B-200D-1F468-1F3FC": {
		Key:        "1F468-1F3FD-200D-2764-FE0F-200D-1F48B-200D-1F468-1F3FC",
		Value:      "👨🏽‍❤️‍💋‍👨🏼",
		Descriptor: "kiss: man, man, medium skin tone, medium-light skin tone",
	},
	"1F468-1F3FD-200D-2764-FE0F-200D-1F48B-200D-1F468-1F3FD": {
		Key:        "1F468-1F3FD-200D-2764-FE0F-200D-1F48B-200D-1F468-1F3FD",
		Value:      "👨🏽‍❤️‍💋‍👨🏽",
		Descriptor: "kiss: man, man, medium skin tone",
	},
	"1F468-1F3FD-200D-2764-FE0F-200D-1F48B-200D-1F468-1F3FE": {
		Key:        "1F468-1F3FD-200D-2764-FE0F-200D-1F48B-200D-1F468-1F3FE",
		Value:      "👨🏽‍❤️‍💋‍👨🏾",
		Descriptor: "kiss: man, man, medium skin tone, medium-dark skin tone",
	},
	"1F468-1F3FD-200D-2764-FE0F-200D-1F48B-200D-1F468-1F3FF": {
		Key:        "1F468-1F3FD-200D-2764-FE0F-200D-1F48B-200D-1F468-1F3FF",
		Value:      "👨🏽‍❤️‍💋‍👨🏿",
		Descriptor: "kiss: man, man, medium skin tone, dark skin tone",
	},
	"1F468-1F3FE-200D-2764-FE0F-200D-1F48B-200D-1F468-1F3FB": {
		Key:        "1F468-1F3FE-200D-2764-FE0F-200D-1F48B-200D-1F468-1F3FB",
		Value:      "👨🏾‍❤️‍💋‍👨🏻",
		Descriptor: "kiss: man, man, medium-dark skin tone, light skin tone",
	},
	"1F468-1F3FE-200D-2764-FE0F-200D-1F48B-200D-1F468-1F3FC": {
		Key:        "1F468-1F3FE-200D-2764-FE0F-200D-1F48B-200D-1F468-1F3FC",
		Value:      "👨🏾‍❤️‍💋‍👨🏼",
		Descriptor: "kiss: man, man, medium-dark skin tone, medium-light skin tone",
	},
	"1F468-1F3FE-200D-2764-FE0F-200D-1F48B-200D-1F468-1F3FD": {
		Key:        "1F468-1F3FE-200D-2764-FE0F-200D-1F48B-200D-1F468-1F3FD",
		Value:      "👨🏾‍❤️‍💋‍👨🏽",
		Descriptor: "kiss: man, man, medium-dark skin tone, medium skin tone",
	},
	"1F468-1F3FE-200D-2764-FE0F-200D-1F48B-200D-1F468-1F3FE": {
		Key:        "1F468-1F3FE-200D-2764-FE0F-200D-1F48B-200D-1F468-1F3FE",
		Value:      "👨🏾‍❤️‍💋‍👨🏾",
		Descriptor: "kiss: man, man, medium-dark skin tone",
	},
	"1F468-1F3FE-200D-2764-FE0F-200D-1F48B-200D-1F468-1F3FF": {
		Key:        "1F468-1F3FE-200D-2764-FE0F-200D-1F48B-200D-1F468-1F3FF",
		Value:      "👨🏾‍❤️‍💋‍👨🏿",
		Descriptor: "kiss: man, man, medium-dark skin tone, dark skin tone",
	},
	"1F468-1F3FF-200D-2764-FE0F-200D-1F48B-200D-1F468-1F3FB": {
		Key:        "1F468-1F3FF-200D-2764-FE0F-200D-1F48B-200D-1F468-1F3FB",
		Value:      "👨🏿‍❤️‍💋‍👨🏻",
		Descriptor: "kiss: man, man, dark skin tone, light skin tone",
	},
	"1F468-1F3FF-200D-2764-FE0F-200D-1F48B-200D-1F468-1F3FC": {
		Key:        "1F468-1F3FF-200D-2764-FE0F-200D-1F48B-200D-1F468-1F3FC",
		Value:      "👨🏿‍❤️‍💋‍👨🏼",
		Descriptor: "kiss: man, man, dark skin tone, medium-light skin tone",
	},
	"1F468-1F3FF-200D-2764-FE0F-200D-1F48B-200D-1F468-1F3FD": {
		Key:        "1F468-1F3FF-200D-2764-FE0F-200D-1F48B-200D-1F468-1F3FD",
		Value:      "👨🏿‍❤️‍💋‍👨🏽",
		Descriptor: "kiss: man, man, dark skin tone, medium skin tone",
	},
	"1F468-1F3FF-200D-2764-FE0F-200D-1F48B-200D-1F468-1F3FE": {
		Key:        "1F468-1F3FF-200D-2764-FE0F-200D-1F48B-200D-1F468-1F3FE",
		Value:      "👨🏿‍❤️‍💋‍👨🏾",
		Descriptor: "kiss: man, man, dark skin tone, medium-dark skin tone",
	},
	"1F468-1F3FF-200D-2764-FE0F-200D-1F48B-200D-1F468-1F3FF": {
		Key:        "1F468-1F3FF-200D-2764-FE0F-200D-1F48B-200D-1F468-1F3FF",
		Value:      "👨🏿‍❤️‍💋‍👨🏿",
		Descriptor: "kiss: man, man, dark skin tone",
	},
	"1F469-200D-2764-FE0F-200D-1F48B-200D-1F469": {
		Key:        "1F469-200D-2764-FE0F-200D-1F48B-200D-1F469",
		Value:      "👩‍❤️‍💋‍👩",
		Descriptor: "kiss: woman, woman",
	},
	"1F469-1F3FB-200D-2764-FE0F-200D-1F48B-200D-1F469-1F3FB": {
		Key:        "1F469-1F3FB-200D-2764-FE0F-200D-1F48B-200D-1F469-1F3FB",
		Value:      "👩🏻‍❤️‍💋‍👩🏻",
		Descriptor: "kiss: woman, woman, light skin tone",
	},
	"1F469-1F3FB-200D-2764-FE0F-200D-1F48B-200D-1F469-1F3FC": {
		Key:        "1F469-1F3FB-200D-2764-FE0F-200D-1F48B-200D-1F469-1F3FC",
		Value:      "👩🏻‍❤️‍💋‍👩🏼",
		Descriptor: "kiss: woman, woman, light skin tone, medium-light skin tone",
	},
	"1F469-1F3FB-200D-2764-FE0F-200D-1F48B-200D-1F469-1F3FD": {
		Key:        "1F469-1F3FB-200D-2764-FE0F-200D-1F48B-200D-1F469-1F3FD",
		Value:      "👩🏻‍❤️‍💋‍👩🏽",
		Descriptor: "kiss: woman, woman, light skin tone, medium skin tone",
	},
	"1F469-1F3FB-200D-2764-FE0F-200D-1F48B-200D-1F469-1F3FE": {
		Key:        "1F469-1F3FB-200D-2764-FE0F-200D-1F48B-200D-1F469-1F3FE",
		Value:      "👩🏻‍❤️‍💋‍👩🏾",
		Descriptor: "kiss: woman, woman, light skin tone, medium-dark skin tone",
	},
	"1F469-1F3FB-200D-2764-FE0F-200D-1F48B-200D-1F469-1F3FF": {
		Key:        "1F469-1F3FB-200D-2764-FE0F-200D-1F48B-200D-1F469-1F3FF",
		Value:      "👩🏻‍❤️‍💋‍👩🏿",
		Descriptor: "kiss: woman, woman, light skin tone, dark skin tone",
	},
	"1F469-1F3FC-200D-2764-FE0F-200D-1F48B-200D-1F469-1F3FB": {
		Key:        "1F469-1F3FC-200D-2764-FE0F-200D-1F48B-200D-1F469-1F3FB",
		Value:      "👩🏼‍❤️‍💋‍👩🏻",
		Descriptor: "kiss: woman, woman, medium-light skin tone, light skin tone",
	},
	"1F469-1F3FC-200D-2764-FE0F-200D-1F48B-200D-1F469-1F3FC": {
		Key:        "1F469-1F3FC-200D-2764-FE0F-200D-1F48B-200D-1F469-1F3FC",
		Value:      "👩🏼‍❤️‍💋‍👩🏼",
		Descriptor: "kiss: woman, woman, medium-light skin tone",
	},
	"1F469-1F3FC-200D-2764-FE0F-200D-1F48B-200D-1F469-1F3FD": {
		Key:        "1F469-1F3FC-200D-2764-FE0F-200D-1F48B-200D-1F469-1F3FD",
		Value:      "👩🏼‍❤️‍💋‍👩🏽",
		Descriptor: "kiss: woman, woman, medium-light skin tone, medium skin tone",
	},
	"1F469-1F3FC-200D-2764-FE0F-200D-1F48B-200D-1F469-1F3FE": {
		Key:        "1F469-1F3FC-200D-2764-FE0F-200D-1F48B-200D-1F469-1F3FE",
		Value:      "👩🏼‍❤️‍💋‍👩🏾",
		Descriptor: "kiss: woman, woman, medium-light skin tone, medium-dark skin tone",
	},
	"1F469-1F3FC-200D-2764-FE0F-200D-1F48B-200D-1F469-1F3FF": {
		Key:        "1F469-1F3FC-200D-2764-FE0F-200D-1F48B-200D-1F469-1F3FF",
		Value:      "👩🏼‍❤️‍💋‍👩🏿",
		Descriptor: "kiss: woman, woman, medium-light skin tone, dark skin tone",
	},
	"1F469-1F3FD-200D-2764-FE0F-200D-1F48B-200D-1F469-1F3FB": {
		Key:        "1F469-1F3FD-200D-2764-FE0F-200D-1F48B-200D-1F469-1F3FB",
		Value:      "👩🏽‍❤️‍💋‍👩🏻",
		Descriptor: "kiss: woman, woman, medium skin tone, light skin tone",
	},
	"1F469-1F3FD-200D-2764-FE0F-200D-1F48B-200D-1F469-1F3FC": {
		Key:        "1F469-1F3FD-200D-2764-FE0F-200D-1F48B-200D-1F469-1F3FC",
		Value:      "👩🏽‍❤️‍💋‍👩🏼",
		Descriptor: "kiss: woman, woman, medium skin tone, medium-light skin tone",
	},
	"1F469-1F3FD-200D-2764-FE0F-200D-1F48B-200D-1F469-1F3FD": {
		Key:        "1F469-1F3FD-200D-2764-FE0F-200D-1F48B-200D-1F469-1F3FD",
		Value:      "👩🏽‍❤️‍💋‍👩🏽",
		Descriptor: "kiss: woman, woman, medium skin tone",
	},
	"1F469-1F3FD-200D-2764-FE0F-200D-1F48B-200D-1F469-1F3FE": {
		Key:        "1F469-1F3FD-200D-2764-FE0F-200D-1F48B-200D-1F469-1F3FE",
		Value:      "👩🏽‍❤️‍💋‍👩🏾",
		Descriptor: "kiss: woman, woman, medium skin tone, medium-dark skin tone",
	},
	"1F469-1F3FD-200D-2764-FE0F-200D-1F48B-200D-1F469-1F3FF": {
		Key:        "1F469-1F3FD-200D-2764-FE0F-200D-1F48B-200D-1F469-1F3FF",
		Value:      "👩🏽‍❤️‍💋‍👩🏿",
		Descriptor: "kiss: woman, woman, medium skin tone, dark skin tone",
	},
	"1F469-1F3FE-200D-2764-FE0F-200D-1F48B-200D-1F469-1F3FB": {
		Key:        "1F469-1F3FE-200D-2764-FE0F-200D-1F48B-200D-1F469-1F3FB",
		Value:      "👩🏾‍❤️‍💋‍👩🏻",
		Descriptor: "kiss: woman, woman, medium-dark skin tone, light skin tone",
	},
	"1F469-1F3FE-200D-2764-FE0F-200D-1F48B-200D-1F469-1F3FC": {
		Key:        "1F469-1F3FE-200D-2764-FE0F-200D-1F48B-200D-1F469-1F3FC",
		Value:      "👩🏾‍❤️‍💋‍👩🏼",
		Descriptor: "kiss: woman, woman, medium-dark skin tone, medium-light skin tone",
	},
	"1F469-1F3FE-200D-2764-FE0F-200D-1F48B-200D-1F469-1F3FD": {
		Key:        "1F469-1F3FE-200D-2764-FE0F-200D-1F48B-200D-1F469-1F3FD",
		Value:      "👩🏾‍❤️‍💋‍👩🏽",
		Descriptor: "kiss: woman, woman, medium-dark skin tone, medium skin tone",
	},
	"1F469-1F3FE-200D-2764-FE0F-200D-1F48B-200D-1F469-1F3FE": {
		Key:        "1F469-1F3FE-200D-2764-FE0F-200D-1F48B-200D-1F469-1F3FE",
		Value:      "👩🏾‍❤️‍💋‍👩🏾",
		Descriptor: "kiss: woman, woman, medium-dark skin tone",
	},
	"1F469-1F3FE-200D-2764-FE0F-200D-1F48B-200D-1F469-1F3FF": {
		Key:        "1F469-1F3FE-200D-2764-FE0F-200D-1F48B-200D-1F469-1F3FF",
		Value:      "👩🏾‍❤️‍💋‍👩🏿",
		Descriptor: "kiss: woman, woman, medium-dark skin tone, dark skin tone",
	},
	"1F469-1F3FF-200D-2764-FE0F-200D-1F48B-200D-1F469-1F3FB": {
		Key:        "1F469-1F3FF-200D-2764-FE0F-200D-1F48B-200D-1F469-1F3FB",
		Value:      "👩🏿‍❤️‍💋‍👩🏻",
		Descriptor: "kiss: woman, woman, dark skin tone, light skin tone",
	},
	"1F469-1F3FF-200D-2764-FE0F-200D-1F48B-200D-1F469-1F3FC": {
		Key:        "1F469-1F3FF-200D-2764-FE0F-200D-1F48B-200D-1F469-1F3FC",
		Value:      "👩🏿‍❤️‍💋‍👩🏼",
		Descriptor: "kiss: woman, woman, dark skin tone, medium-light skin tone",
	},
	"1F469-1F3FF-200D-2764-FE0F-200D-1F48B-200D-1F469-1F3FD": {
		Key:        "1F469-1F3FF-200D-2764-FE0F-200D-1F48B-200D-1F469-1F3FD",
		Value:      "👩🏿‍❤️‍💋‍👩🏽",
		Descriptor: "kiss: woman, woman, dark skin tone, medium skin tone",
	},
	"1F469-1F3FF-200D-2764-FE0F-200D-1F48B-200D-1F469-1F3FE": {
		Key:        "1F469-1F3FF-200D-2764-FE0F-200D-1F48B-200D-1F469-1F3FE",
		Value:      "👩🏿‍❤️‍💋‍👩🏾",
		Descriptor: "kiss: woman, woman, dark skin tone, medium-dark skin tone",
	},
	"1F469-1F3FF-200D-2764-FE0F-200D-1F48B-200D-1F469-1F3FF": {
		Key:        "1F469-1F3FF-200D-2764-FE0F-200D-1F48B-200D-1F469-1F3FF",
		Value:      "👩🏿‍❤️‍💋‍👩🏿",
		Descriptor: "kiss: woman, woman, dark skin tone",
	},
	"1F491": {
		Key:        "1F491",
		Value:      "💑",
		Descriptor: "couple with heart",
	},
	"1F491-1F3FB": {
		Key:        "1F491-1F3FB",
		Value:      "💑🏻",
		Descriptor: "couple with heart: light skin tone",
	},
	"1F491-1F3FC": {
		Key:        "1F491-1F3FC",
		Value:      "💑🏼",
		Descriptor: "couple with heart: medium-light skin tone",
	},
	"1F491-1F3FD": {
		Key:        "1F491-1F3FD",
		Value:      "💑🏽",
		Descriptor: "couple with heart: medium skin tone",
	},
	"1F491-1F3FE": {
		Key:        "1F491-1F3FE",
		Value:      "💑🏾",
		Descriptor: "couple with heart: medium-dark skin tone",
	},
	"1F491-1F3FF": {
		Key:        "1F491-1F3FF",
		Value:      "💑🏿",
		Descriptor: "couple with heart: dark skin tone",
	},
	"1F9D1-1F3FB-200D-2764-FE0F-200D-1F9D1-1F3FC": {
		Key:        "1F9D1-1F3FB-200D-2764-FE0F-200D-1F9D1-1F3FC",
		Value:      "🧑🏻‍❤️‍🧑🏼",
		Descriptor: "couple with heart: person, person, light skin tone, medium-light skin tone",
	},
	"1F9D1-1F3FB-200D-2764-FE0F-200D-1F9D1-1F3FD": {
		Key:        "1F9D1-1F3FB-200D-2764-FE0F-200D-1F9D1-1F3FD",
		Value:      "🧑🏻‍❤️‍🧑🏽",
		Descriptor: "couple with heart: person, person, light skin tone, medium skin tone",
	},
	"1F9D1-1F3FB-200D-2764-FE0F-200D-1F9D1-1F3FE": {
		Key:        "1F9D1-1F3FB-200D-2764-FE0F-200D-1F9D1-1F3FE",
		Value:      "🧑🏻‍❤️‍🧑🏾",
		Descriptor: "couple with heart: person, person, light skin tone, medium-dark skin tone",
	},
	"1F9D1-1F3FB-200D-2764-FE0F-200D-1F9D1-1F3FF": {
		Key:        "1F9D1-1F3FB-200D-2764-FE0F-200D-1F9D1-1F3FF",
		Value:      "🧑🏻‍❤️‍🧑🏿",
		Descriptor: "couple with heart: person, person, light skin tone, dark skin tone",
	},
	"1F9D1-1F3FC-200D-2764-FE0F-200D-1F9D1-1F3FB": {
		Key:        "1F9D1-1F3FC-200D-2764-FE0F-200D-1F9D1-1F3FB",
		Value:      "🧑🏼‍❤️‍🧑🏻",
		Descriptor: "couple with heart: person, person, medium-light skin tone, light skin tone",
	},
	"1F9D1-1F3FC-200D-2764-FE0F-200D-1F9D1-1F3FD": {
		Key:        "1F9D1-1F3FC-200D-2764-FE0F-200D-1F9D1-1F3FD",
		Value:      "🧑🏼‍❤️‍🧑🏽",
		Descriptor: "couple with heart: person, person, medium-light skin tone, medium skin tone",
	},
	"1F9D1-1F3FC-200D-2764-FE0F-200D-1F9D1-1F3FE": {
		Key:        "1F9D1-1F3FC-200D-2764-FE0F-200D-1F9D1-1F3FE",
		Value:      "🧑🏼‍❤️‍🧑🏾",
		Descriptor: "couple with heart: person, person, medium-light skin tone, medium-dark skin tone",
	},
	"1F9D1-1F3FC-200D-2764-FE0F-200D-1F9D1-1F3FF": {
		Key:        "1F9D1-1F3FC-200D-2764-FE0F-200D-1F9D1-1F3FF",
		Value:      "🧑🏼‍❤️‍🧑🏿",
		Descriptor: "couple with heart: person, person, medium-light skin tone, dark skin tone",
	},
	"1F9D1-1F3FD-200D-2764-FE0F-200D-1F9D1-1F3FB": {
		Key:        "1F9D1-1F3FD-200D-2764-FE0F-200D-1F9D1-1F3FB",
		Value:      "🧑🏽‍❤️‍🧑🏻",
		Descriptor: "couple with heart: person, person, medium skin tone, light skin tone",
	},
	"1F9D1-1F3FD-200D-2764-FE0F-200D-1F9D1-1F3FC": {
		Key:        "1F9D1-1F3FD-200D-2764-FE0F-200D-1F9D1-1F3FC",
		Value:      "🧑🏽‍❤️‍🧑🏼",
		Descriptor: "couple with heart: person, person, medium skin tone, medium-light skin tone",
	},
	"1F9D1-1F3FD-200D-2764-FE0F-200D-1F9D1-1F3FE": {
		Key:        "1F9D1-1F3FD-200D-2764-FE0F-200D-1F9D1-1F3FE",
		Value:      "🧑🏽‍❤️‍🧑🏾",
		Descriptor: "couple with heart: person, person, medium skin tone, medium-dark skin tone",
	},
	"1F9D1-1F3FD-200D-2764-FE0F-200D-1F9D1-1F3FF": {
		Key:        "1F9D1-1F3FD-200D-2764-FE0F-200D-1F9D1-1F3FF",
		Value:      "🧑🏽‍❤️‍🧑🏿",
		Descriptor: "couple with heart: person, person, medium skin tone, dark skin tone",
	},
	"1F9D1-1F3FE-200D-2764-FE0F-200D-1F9D1-1F3FB": {
		Key:        "1F9D1-1F3FE-200D-2764-FE0F-200D-1F9D1-1F3FB",
		Value:      "🧑🏾‍❤️‍🧑🏻",
		Descriptor: "couple with heart: person, person, medium-dark skin tone, light skin tone",
	},
	"1F9D1-1F3FE-200D-2764-FE0F-200D-1F9D1-1F3FC": {
		Key:        "1F9D1-1F3FE-200D-2764-FE0F-200D-1F9D1-1F3FC",
		Value:      "🧑🏾‍❤️‍🧑🏼",
		Descriptor: "couple with heart: person, person, medium-dark skin tone, medium-light skin tone",
	},
	"1F9D1-1F3FE-200D-2764-FE0F-200D-1F9D1-1F3FD": {
		Key:        "1F9D1-1F3FE-200D-2764-FE0F-200D-1F9D1-1F3FD",
		Value:      "🧑🏾‍❤️‍🧑🏽",
		Descriptor: "couple with heart: person, person, medium-dark skin tone, medium skin tone",
	},
	"1F9D1-1F3FE-200D-2764-FE0F-200D-1F9D1-1F3FF": {
		Key:        "1F9D1-1F3FE-200D-2764-FE0F-200D-1F9D1-1F3FF",
		Value:      "🧑🏾‍❤️‍🧑🏿",
		Descriptor: "couple with heart: person, person, medium-dark skin tone, dark skin tone",
	},
	"1F9D1-1F3FF-200D-2764-FE0F-200D-1F9D1-1F3FB": {
		Key:        "1F9D1-1F3FF-200D-2764-FE0F-200D-1F9D1-1F3FB",
		Value:      "🧑🏿‍❤️‍🧑🏻",
		Descriptor: "couple with heart: person, person, dark skin tone, light skin tone",
	},
	"1F9D1-1F3FF-200D-2764-FE0F-200D-1F9D1-1F3FC": {
		Key:        "1F9D1-1F3FF-200D-2764-FE0F-200D-1F9D1-1F3FC",
		Value:      "🧑🏿‍❤️‍🧑🏼",
		Descriptor: "couple with heart: person, person, dark skin tone, medium-light skin tone",
	},
	"1F9D1-1F3FF-200D-2764-FE0F-200D-1F9D1-1F3FD": {
		Key:        "1F9D1-1F3FF-200D-2764-FE0F-200D-1F9D1-1F3FD",
		Value:      "🧑🏿‍❤️‍🧑🏽",
		Descriptor: "couple with heart: person, person, dark skin tone, medium skin tone",
	},
	"1F9D1-1F3FF-200D-2764-FE0F-200D-1F9D1-1F3FE": {
		Key:        "1F9D1-1F3FF-200D-2764-FE0F-200D-1F9D1-1F3FE",
		Value:      "🧑🏿‍❤️‍🧑🏾",
		Descriptor: "couple with heart: person, person, dark skin tone, medium-dark skin tone",
	},
	"1F469-200D-2764-FE0F-200D-1F468": {
		Key:        "1F469-200D-2764-FE0F-200D-1F468",
		Value:      "👩‍❤️‍👨",
		Descriptor: "couple with heart: woman, man",
	},
	"1F469-1F3FB-200D-2764-FE0F-200D-1F468-1F3FB": {
		Key:        "1F469-1F3FB-200D-2764-FE0F-200D-1F468-1F3FB",
		Value:      "👩🏻‍❤️‍👨🏻",
		Descriptor: "couple with heart: woman, man, light skin tone",
	},
	"1F469-1F3FB-200D-2764-FE0F-200D-1F468-1F3FC": {
		Key:        "1F469-1F3FB-200D-2764-FE0F-200D-1F468-1F3FC",
		Value:      "👩🏻‍❤️‍👨🏼",
		Descriptor: "couple with heart: woman, man, light skin tone, medium-light skin tone",
	},
	"1F469-1F3FB-200D-2764-FE0F-200D-1F468-1F3FD": {
		Key:        "1F469-1F3FB-200D-2764-FE0F-200D-1F468-1F3FD",
		Value:      "👩🏻‍❤️‍👨🏽",
		Descriptor: "couple with heart: woman, man, light skin tone, medium skin tone",
	},
	"1F469-1F3FB-200D-2764-FE0F-200D-1F468-1F3FE": {
		Key:        "1F469-1F3FB-200D-2764-FE0F-200D-1F468-1F3FE",
		Value:      "👩🏻‍❤️‍👨🏾",
		Descriptor: "couple with heart: woman, man, light skin tone, medium-dark skin tone",
	},
	"1F469-1F3FB-200D-2764-FE0F-200D-1F468-1F3FF": {
		Key:        "1F469-1F3FB-200D-2764-FE0F-200D-1F468-1F3FF",
		Value:      "👩🏻‍❤️‍👨🏿",
		Descriptor: "couple with heart: woman, man, light skin tone, dark skin tone",
	},
	"1F469-1F3FC-200D-2764-FE0F-200D-1F468-1F3FB": {
		Key:        "1F469-1F3FC-200D-2764-FE0F-200D-1F468-1F3FB",
		Value:      "👩🏼‍❤️‍👨🏻",
		Descriptor: "couple with heart: woman, man, medium-light skin tone, light skin tone",
	},
	"1F469-1F3FC-200D-2764-FE0F-200D-1F468-1F3FC": {
		Key:        "1F469-1F3FC-200D-2764-FE0F-200D-1F468-1F3FC",
		Value:      "👩🏼‍❤️‍👨🏼",
		Descriptor: "couple with heart: woman, man, medium-light skin tone",
	},
	"1F469-1F3FC-200D-2764-FE0F-200D-1F468-1F3FD": {
		Key:        "1F469-1F3FC-200D-2764-FE0F-200D-1F468-1F3FD",
		Value:      "👩🏼‍❤️‍👨🏽",
		Descriptor: "couple with heart: woman, man, medium-light skin tone, medium skin tone",
	},
	"1F469-1F3FC-200D-2764-FE0F-200D-1F468-1F3FE": {
		Key:        "1F469-1F3FC-200D-2764-FE0F-200D-1F468-1F3FE",
		Value:      "👩🏼‍❤️‍👨🏾",
		Descriptor: "couple with heart: woman, man, medium-light skin tone, medium-dark skin tone",
	},
	"1F469-1F3FC-200D-2764-FE0F-200D-1F468-1F3FF": {
		Key:        "1F469-1F3FC-200D-2764-FE0F-200D-1F468-1F3FF",
		Value:      "👩🏼‍❤️‍👨🏿",
		Descriptor: "couple with heart: woman, man, medium-light skin tone, dark skin tone",
	},
	"1F469-1F3FD-200D-2764-FE0F-200D-1F468-1F3FB": {
		Key:        "1F469-1F3FD-200D-2764-FE0F-200D-1F468-1F3FB",
		Value:      "👩🏽‍❤️‍👨🏻",
		Descriptor: "couple with heart: woman, man, medium skin tone, light skin tone",
	},
	"1F469-1F3FD-200D-2764-FE0F-200D-1F468-1F3FC": {
		Key:        "1F469-1F3FD-200D-2764-FE0F-200D-1F468-1F3FC",
		Value:      "👩🏽‍❤️‍👨🏼",
		Descriptor: "couple with heart: woman, man, medium skin tone, medium-light skin tone",
	},
	"1F469-1F3FD-200D-2764-FE0F-200D-1F468-1F3FD": {
		Key:        "1F469-1F3FD-200D-2764-FE0F-200D-1F468-1F3FD",
		Value:      "👩🏽‍❤️‍👨🏽",
		Descriptor: "couple with heart: woman, man, medium skin tone",
	},
	"1F469-1F3FD-200D-2764-FE0F-200D-1F468-1F3FE": {
		Key:        "1F469-1F3FD-200D-2764-FE0F-200D-1F468-1F3FE",
		Value:      "👩🏽‍❤️‍👨🏾",
		Descriptor: "couple with heart: woman, man, medium skin tone, medium-dark skin tone",
	},
	"1F469-1F3FD-200D-2764-FE0F-200D-1F468-1F3FF": {
		Key:        "1F469-1F3FD-200D-2764-FE0F-200D-1F468-1F3FF",
		Value:      "👩🏽‍❤️‍👨🏿",
		Descriptor: "couple with heart: woman, man, medium skin tone, dark skin tone",
	},
	"1F469-1F3FE-200D-2764-FE0F-200D-1F468-1F3FB": {
		Key:        "1F469-1F3FE-200D-2764-FE0F-200D-1F468-1F3FB",
		Value:      "👩🏾‍❤️‍👨🏻",
		Descriptor: "couple with heart: woman, man, medium-dark skin tone, light skin tone",
	},
	"1F469-1F3FE-200D-2764-FE0F-200D-1F468-1F3FC": {
		Key:        "1F469-1F3FE-200D-2764-FE0F-200D-1F468-1F3FC",
		Value:      "👩🏾‍❤️‍👨🏼",
		Descriptor: "couple with heart: woman, man, medium-dark skin tone, medium-light skin tone",
	},
	"1F469-1F3FE-200D-2764-FE0F-200D-1F468-1F3FD": {
		Key:        "1F469-1F3FE-200D-2764-FE0F-200D-1F468-1F3FD",
		Value:      "👩🏾‍❤️‍👨🏽",
		Descriptor: "couple with heart: woman, man, medium-dark skin tone, medium skin tone",
	},
	"1F469-1F3FE-200D-2764-FE0F-200D-1F468-1F3FE": {
		Key:        "1F469-1F3FE-200D-2764-FE0F-200D-1F468-1F3FE",
		Value:      "👩🏾‍❤️‍👨🏾",
		Descriptor: "couple with heart: woman, man, medium-dark skin tone",
	},
	"1F469-1F3FE-200D-2764-FE0F-200D-1F468-1F3FF": {
		Key:        "1F469-1F3FE-200D-2764-FE0F-200D-1F468-1F3FF",
		Value:      "👩🏾‍❤️‍👨🏿",
		Descriptor: "couple with heart: woman, man, medium-dark skin tone, dark skin tone",
	},
	"1F469-1F3FF-200D-2764-FE0F-200D-1F468-1F3FB": {
		Key:        "1F469-1F3FF-200D-2764-FE0F-200D-1F468-1F3FB",
		Value:      "👩🏿‍❤️‍👨🏻",
		Descriptor: "couple with heart: woman, man, dark skin tone, light skin tone",
	},
	"1F469-1F3FF-200D-2764-FE0F-200D-1F468-1F3FC": {
		Key:        "1F469-1F3FF-200D-2764-FE0F-200D-1F468-1F3FC",
		Value:      "👩🏿‍❤️‍👨🏼",
		Descriptor: "couple with heart: woman, man, dark skin tone, medium-light skin tone",
	},
	"1F469-1F3FF-200D-2764-FE0F-200D-1F468-1F3FD": {
		Key:        "1F469-1F3FF-200D-2764-FE0F-200D-1F468-1F3FD",
		Value:      "👩🏿‍❤️‍👨🏽",
		Descriptor: "couple with heart: woman, man, dark skin tone, medium skin tone",
	},
	"1F469-1F3FF-200D-2764-FE0F-200D-1F468-1F3FE": {
		Key:        "1F469-1F3FF-200D-2764-FE0F-200D-1F468-1F3FE",
		Value:      "👩🏿‍❤️‍👨🏾",
		Descriptor: "couple with heart: woman, man, dark skin tone, medium-dark skin tone",
	},
	"1F469-1F3FF-200D-2764-FE0F-200D-1F468-1F3FF": {
		Key:        "1F469-1F3FF-200D-2764-FE0F-200D-1F468-1F3FF",
		Value:      "👩🏿‍❤️‍👨🏿",
		Descriptor: "couple with heart: woman, man, dark skin tone",
	},
	"1F468-200D-2764-FE0F-200D-1F468": {
		Key:        "1F468-200D-2764-FE0F-200D-1F468",
		Value:      "👨‍❤️‍👨",
		Descriptor: "couple with heart: man, man",
	},
	"1F468-1F3FB-200D-2764-FE0F-200D-1F468-1F3FB": {
		Key:        "1F468-1F3FB-200D-2764-FE0F-200D-1F468-1F3FB",
		Value:      "👨🏻‍❤️‍👨🏻",
		Descriptor: "couple with heart: man, man, light skin tone",
	},
	"1F468-1F3FB-200D-2764-FE0F-200D-1F468-1F3FC": {
		Key:        "1F468-1F3FB-200D-2764-FE0F-200D-1F468-1F3FC",
		Value:      "👨🏻‍❤️‍👨🏼",
		Descriptor: "couple with heart: man, man, light skin tone, medium-light skin tone",
	},
	"1F468-1F3FB-200D-2764-FE0F-200D-1F468-1F3FD": {
		Key:        "1F468-1F3FB-200D-2764-FE0F-200D-1F468-1F3FD",
		Value:      "👨🏻‍❤️‍👨🏽",
		Descriptor: "couple with heart: man, man, light skin tone, medium skin tone",
	},
	"1F468-1F3FB-200D-2764-FE0F-200D-1F468-1F3FE": {
		Key:        "1F468-1F3FB-200D-2764-FE0F-200D-1F468-1F3FE",
		Value:      "👨🏻‍❤️‍👨🏾",
		Descriptor: "couple with heart: man, man, light skin tone, medium-dark skin tone",
	},
	"1F468-1F3FB-200D-2764-FE0F-200D-1F468-1F3FF": {
		Key:        "1F468-1F3FB-200D-2764-FE0F-200D-1F468-1F3FF",
		Value:      "👨🏻‍❤️‍👨🏿",
		Descriptor: "couple with heart: man, man, light skin tone, dark skin tone",
	},
	"1F468-1F3FC-200D-2764-FE0F-200D-1F468-1F3FB": {
		Key:        "1F468-1F3FC-200D-2764-FE0F-200D-1F468-1F3FB",
		Value:      "👨🏼‍❤️‍👨🏻",
		Descriptor: "couple with heart: man, man, medium-light skin tone, light skin tone",
	},
	"1F468-1F3FC-200D-2764-FE0F-200D-1F468-1F3FC": {
		Key:        "1F468-1F3FC-200D-2764-FE0F-200D-1F468-1F3FC",
		Value:      "👨🏼‍❤️‍👨🏼",
		Descriptor: "couple with heart: man, man, medium-light skin tone",
	},
	"1F468-1F3FC-200D-2764-FE0F-200D-1F468-1F3FD": {
		Key:        "1F468-1F3FC-200D-2764-FE0F-200D-1F468-1F3FD",
		Value:      "👨🏼‍❤️‍👨🏽",
		Descriptor: "couple with heart: man, man, medium-light skin tone, medium skin tone",
	},
	"1F468-1F3FC-200D-2764-FE0F-200D-1F468-1F3FE": {
		Key:        "1F468-1F3FC-200D-2764-FE0F-200D-1F468-1F3FE",
		Value:      "👨🏼‍❤️‍👨🏾",
		Descriptor: "couple with heart: man, man, medium-light skin tone, medium-dark skin tone",
	},
	"1F468-1F3FC-200D-2764-FE0F-200D-1F468-1F3FF": {
		Key:        "1F468-1F3FC-200D-2764-FE0F-200D-1F468-1F3FF",
		Value:      "👨🏼‍❤️‍👨🏿",
		Descriptor: "couple with heart: man, man, medium-light skin tone, dark skin tone",
	},
	"1F468-1F3FD-200D-2764-FE0F-200D-1F468-1F3FB": {
		Key:        "1F468-1F3FD-200D-2764-FE0F-200D-1F468-1F3FB",
		Value:      "👨🏽‍❤️‍👨🏻",
		Descriptor: "couple with heart: man, man, medium skin tone, light skin tone",
	},
	"1F468-1F3FD-200D-2764-FE0F-200D-1F468-1F3FC": {
		Key:        "1F468-1F3FD-200D-2764-FE0F-200D-1F468-1F3FC",
		Value:      "👨🏽‍❤️‍👨🏼",
		Descriptor: "couple with heart: man, man, medium skin tone, medium-light skin tone",
	},
	"1F468-1F3FD-200D-2764-FE0F-200D-1F468-1F3FD": {
		Key:        "1F468-1F3FD-200D-2764-FE0F-200D-1F468-1F3FD",
		Value:      "👨🏽‍❤️‍👨🏽",
		Descriptor: "couple with heart: man, man, medium skin tone",
	},
	"1F468-1F3FD-200D-2764-FE0F-200D-1F468-1F3FE": {
		Key:        "1F468-1F3FD-200D-2764-FE0F-200D-1F468-1F3FE",
		Value:      "👨🏽‍❤️‍👨🏾",
		Descriptor: "couple with heart: man, man, medium skin tone, medium-dark skin tone",
	},
	"1F468-1F3FD-200D-2764-FE0F-200D-1F468-1F3FF": {
		Key:        "1F468-1F3FD-200D-2764-FE0F-200D-1F468-1F3FF",
		Value:      "👨🏽‍❤️‍👨🏿",
		Descriptor: "couple with heart: man, man, medium skin tone, dark skin tone",
	},
	"1F468-1F3FE-200D-2764-FE0F-200D-1F468-1F3FB": {
		Key:        "1F468-1F3FE-200D-2764-FE0F-200D-1F468-1F3FB",
		Value:      "👨🏾‍❤️‍👨🏻",
		Descriptor: "couple with heart: man, man, medium-dark skin tone, light skin tone",
	},
	"1F468-1F3FE-200D-2764-FE0F-200D-1F468-1F3FC": {
		Key:        "1F468-1F3FE-200D-2764-FE0F-200D-1F468-1F3FC",
		Value:      "👨🏾‍❤️‍👨🏼",
		Descriptor: "couple with heart: man, man, medium-dark skin tone, medium-light skin tone",
	},
	"1F468-1F3FE-200D-2764-FE0F-200D-1F468-1F3FD": {
		Key:        "1F468-1F3FE-200D-2764-FE0F-200D-1F468-1F3FD",
		Value:      "👨🏾‍❤️‍👨🏽",
		Descriptor: "couple with heart: man, man, medium-dark skin tone, medium skin tone",
	},
	"1F468-1F3FE-200D-2764-FE0F-200D-1F468-1F3FE": {
		Key:        "1F468-1F3FE-200D-2764-FE0F-200D-1F468-1F3FE",
		Value:      "👨🏾‍❤️‍👨🏾",
		Descriptor: "couple with heart: man, man, medium-dark skin tone",
	},
	"1F468-1F3FE-200D-2764-FE0F-200D-1F468-1F3FF": {
		Key:        "1F468-1F3FE-200D-2764-FE0F-200D-1F468-1F3FF",
		Value:      "👨🏾‍❤️‍👨🏿",
		Descriptor: "couple with heart: man, man, medium-dark skin tone, dark skin tone",
	},
	"1F468-1F3FF-200D-2764-FE0F-200D-1F468-1F3FB": {
		Key:        "1F468-1F3FF-200D-2764-FE0F-200D-1F468-1F3FB",
		Value:      "👨🏿‍❤️‍👨🏻",
		Descriptor: "couple with heart: man, man, dark skin tone, light skin tone",
	},
	"1F468-1F3FF-200D-2764-FE0F-200D-1F468-1F3FC": {
		Key:        "1F468-1F3FF-200D-2764-FE0F-200D-1F468-1F3FC",
		Value:      "👨🏿‍❤️‍👨🏼",
		Descriptor: "couple with heart: man, man, dark skin tone, medium-light skin tone",
	},
	"1F468-1F3FF-200D-2764-FE0F-200D-1F468-1F3FD": {
		Key:        "1F468-1F3FF-200D-2764-FE0F-200D-1F468-1F3FD",
		Value:      "👨🏿‍❤️‍👨🏽",
		Descriptor: "couple with heart: man, man, dark skin tone, medium skin tone",
	},
	"1F468-1F3FF-200D-2764-FE0F-200D-1F468-1F3FE": {
		Key:        "1F468-1F3FF-200D-2764-FE0F-200D-1F468-1F3FE",
		Value:      "👨🏿‍❤️‍👨🏾",
		Descriptor: "couple with heart: man, man, dark skin tone, medium-dark skin tone",
	},
	"1F468-1F3FF-200D-2764-FE0F-200D-1F468-1F3FF": {
		Key:        "1F468-1F3FF-200D-2764-FE0F-200D-1F468-1F3FF",
		Value:      "👨🏿‍❤️‍👨🏿",
		Descriptor: "couple with heart: man, man, dark skin tone",
	},
	"1F469-200D-2764-FE0F-200D-1F469": {
		Key:        "1F469-200D-2764-FE0F-200D-1F469",
		Value:      "👩‍❤️‍👩",
		Descriptor: "couple with heart: woman, woman",
	},
	"1F469-1F3FB-200D-2764-FE0F-200D-1F469-1F3FB": {
		Key:        "1F469-1F3FB-200D-2764-FE0F-200D-1F469-1F3FB",
		Value:      "👩🏻‍❤️‍👩🏻",
		Descriptor: "couple with heart: woman, woman, light skin tone",
	},
	"1F469-1F3FB-200D-2764-FE0F-200D-1F469-1F3FC": {
		Key:        "1F469-1F3FB-200D-2764-FE0F-200D-1F469-1F3FC",
		Value:      "👩🏻‍❤️‍👩🏼",
		Descriptor: "couple with heart: woman, woman, light skin tone, medium-light skin tone",
	},
	"1F469-1F3FB-200D-2764-FE0F-200D-1F469-1F3FD": {
		Key:        "1F469-1F3FB-200D-2764-FE0F-200D-1F469-1F3FD",
		Value:      "👩🏻‍❤️‍👩🏽",
		Descriptor: "couple with heart: woman, woman, light skin tone, medium skin tone",
	},
	"1F469-1F3FB-200D-2764-FE0F-200D-1F469-1F3FE": {
		Key:        "1F469-1F3FB-200D-2764-FE0F-200D-1F469-1F3FE",
		Value:      "👩🏻‍❤️‍👩🏾",
		Descriptor: "couple with heart: woman, woman, light skin tone, medium-dark skin tone",
	},
	"1F469-1F3FB-200D-2764-FE0F-200D-1F469-1F3FF": {
		Key:        "1F469-1F3FB-200D-2764-FE0F-200D-1F469-1F3FF",
		Value:      "👩🏻‍❤️‍👩🏿",
		Descriptor: "couple with heart: woman, woman, light skin tone, dark skin tone",
	},
	"1F469-1F3FC-200D-2764-FE0F-200D-1F469-1F3FB": {
		Key:        "1F469-1F3FC-200D-2764-FE0F-200D-1F469-1F3FB",
		Value:      "👩🏼‍❤️‍👩🏻",
		Descriptor: "couple with heart: woman, woman, medium-light skin tone, light skin tone",
	},
	"1F469-1F3FC-200D-2764-FE0F-200D-1F469-1F3FC": {
		Key:        "1F469-1F3FC-200D-2764-FE0F-200D-1F469-1F3FC",
		Value:      "👩🏼‍❤️‍👩🏼",
		Descriptor: "couple with heart: woman, woman, medium-light skin tone",
	},
	"1F469-1F3FC-200D-2764-FE0F-200D-1F469-1F3FD": {
		Key:        "1F469-1F3FC-200D-2764-FE0F-200D-1F469-1F3FD",
		Value:      "👩🏼‍❤️‍👩🏽",
		Descriptor: "couple with heart: woman, woman, medium-light skin tone, medium skin tone",
	},
	"1F469-1F3FC-200D-2764-FE0F-200D-1F469-1F3FE": {
		Key:        "1F469-1F3FC-200D-2764-FE0F-200D-1F469-1F3FE",
		Value:      "👩🏼‍❤️‍👩🏾",
		Descriptor: "couple with heart: woman, woman, medium-light skin tone, medium-dark skin tone",
	},
	"1F469-1F3FC-200D-2764-FE0F-200D-1F469-1F3FF": {
		Key:        "1F469-1F3FC-200D-2764-FE0F-200D-1F469-1F3FF",
		Value:      "👩🏼‍❤️‍👩🏿",
		Descriptor: "couple with heart: woman, woman, medium-light skin tone, dark skin tone",
	},
	"1F469-1F3FD-200D-2764-FE0F-200D-1F469-1F3FB": {
		Key:        "1F469-1F3FD-200D-2764-FE0F-200D-1F469-1F3FB",
		Value:      "👩🏽‍❤️‍👩🏻",
		Descriptor: "couple with heart: woman, woman, medium skin tone, light skin tone",
	},
	"1F469-1F3FD-200D-2764-FE0F-200D-1F469-1F3FC": {
		Key:        "1F469-1F3FD-200D-2764-FE0F-200D-1F469-1F3FC",
		Value:      "👩🏽‍❤️‍👩🏼",
		Descriptor: "couple with heart: woman, woman, medium skin tone, medium-light skin tone",
	},
	"1F469-1F3FD-200D-2764-FE0F-200D-1F469-1F3FD": {
		Key:        "1F469-1F3FD-200D-2764-FE0F-200D-1F469-1F3FD",
		Value:      "👩🏽‍❤️‍👩🏽",
		Descriptor: "couple with heart: woman, woman, medium skin tone",
	},
	"1F469-1F3FD-200D-2764-FE0F-200D-1F469-1F3FE": {
		Key:        "1F469-1F3FD-200D-2764-FE0F-200D-1F469-1F3FE",
		Value:      "👩🏽‍❤️‍👩🏾",
		Descriptor: "couple with heart: woman, woman, medium skin tone, medium-dark skin tone",
	},
	"1F469-1F3FD-200D-2764-FE0F-200D-1F469-1F3FF": {
		Key:        "1F469-1F3FD-200D-2764-FE0F-200D-1F469-1F3FF",
		Value:      "👩🏽‍❤️‍👩🏿",
		Descriptor: "couple with heart: woman, woman, medium skin tone, dark skin tone",
	},
	"1F469-1F3FE-200D-2764-FE0F-200D-1F469-1F3FB": {
		Key:        "1F469-1F3FE-200D-2764-FE0F-200D-1F469-1F3FB",
		Value:      "👩🏾‍❤️‍👩🏻",
		Descriptor: "couple with heart: woman, woman, medium-dark skin tone, light skin tone",
	},
	"1F469-1F3FE-200D-2764-FE0F-200D-1F469-1F3FC": {
		Key:        "1F469-1F3FE-200D-2764-FE0F-200D-1F469-1F3FC",
		Value:      "👩🏾‍❤️‍👩🏼",
		Descriptor: "couple with heart: woman, woman, medium-dark skin tone, medium-light skin tone",
	},
	"1F469-1F3FE-200D-2764-FE0F-200D-1F469-1F3FD": {
		Key:        "1F469-1F3FE-200D-2764-FE0F-200D-1F469-1F3FD",
		Value:      "👩🏾‍❤️‍👩🏽",
		Descriptor: "couple with heart: woman, woman, medium-dark skin tone, medium skin tone",
	},
	"1F469-1F3FE-200D-2764-FE0F-200D-1F469-1F3FE": {
		Key:        "1F469-1F3FE-200D-2764-FE0F-200D-1F469-1F3FE",
		Value:      "👩🏾‍❤️‍👩🏾",
		Descriptor: "couple with heart: woman, woman, medium-dark skin tone",
	},
	"1F469-1F3FE-200D-2764-FE0F-200D-1F469-1F3FF": {
		Key:        "1F469-1F3FE-200D-2764-FE0F-200D-1F469-1F3FF",
		Value:      "👩🏾‍❤️‍👩🏿",
		Descriptor: "couple with heart: woman, woman, medium-dark skin tone, dark skin tone",
	},
	"1F469-1F3FF-200D-2764-FE0F-200D-1F469-1F3FB": {
		Key:        "1F469-1F3FF-200D-2764-FE0F-200D-1F469-1F3FB",
		Value:      "👩🏿‍❤️‍👩🏻",
		Descriptor: "couple with heart: woman, woman, dark skin tone, light skin tone",
	},
	"1F469-1F3FF-200D-2764-FE0F-200D-1F469-1F3FC": {
		Key:        "1F469-1F3FF-200D-2764-FE0F-200D-1F469-1F3FC",
		Value:      "👩🏿‍❤️‍👩🏼",
		Descriptor: "couple with heart: woman, woman, dark skin tone, medium-light skin tone",
	},
	"1F469-1F3FF-200D-2764-FE0F-200D-1F469-1F3FD": {
		Key:        "1F469-1F3FF-200D-2764-FE0F-200D-1F469-1F3FD",
		Value:      "👩🏿‍❤️‍👩🏽",
		Descriptor: "couple with heart: woman, woman, dark skin tone, medium skin tone",
	},
	"1F469-1F3FF-200D-2764-FE0F-200D-1F469-1F3FE": {
		Key:        "1F469-1F3FF-200D-2764-FE0F-200D-1F469-1F3FE",
		Value:      "👩🏿‍❤️‍👩🏾",
		Descriptor: "couple with heart: woman, woman, dark skin tone, medium-dark skin tone",
	},
	"1F469-1F3FF-200D-2764-FE0F-200D-1F469-1F3FF": {
		Key:        "1F469-1F3FF-200D-2764-FE0F-200D-1F469-1F3FF",
		Value:      "👩🏿‍❤️‍👩🏿",
		Descriptor: "couple with heart: woman, woman, dark skin tone",
	},
	"1F468-200D-1F469-200D-1F466": {
		Key:        "1F468-200D-1F469-200D-1F466",
		Value:      "👨‍👩‍👦",
		Descriptor: "family: man, woman, boy",
	},
	"1F468-200D-1F469-200D-1F467": {
		Key:        "1F468-200D-1F469-200D-1F467",
		Value:      "👨‍👩‍👧",
		Descriptor: "family: man, woman, girl",
	},
	"1F468-200D-1F469-200D-1F467-200D-1F466": {
		Key:        "1F468-200D-1F469-200D-1F467-200D-1F466",
		Value:      "👨‍👩‍👧‍👦",
		Descriptor: "family: man, woman, girl, boy",
	},
	"1F468-200D-1F469-200D-1F466-200D-1F466": {
		Key:        "1F468-200D-1F469-200D-1F466-200D-1F466",
		Value:      "👨‍👩‍👦‍👦",
		Descriptor: "family: man, woman, boy, boy",
	},
	"1F468-200D-1F469-200D-1F467-200D-1F467": {
		Key:        "1F468-200D-1F469-200D-1F467-200D-1F467",
		Value:      "👨‍👩‍👧‍👧",
		Descriptor: "family: man, woman, girl, girl",
	},
	"1F468-200D-1F468-200D-1F466": {
		Key:        "1F468-200D-1F468-200D-1F466",
		Value:      "👨‍👨‍👦",
		Descriptor: "family: man, man, boy",
	},
	"1F468-200D-1F468-200D-1F467": {
		Key:        "1F468-200D-1F468-200D-1F467",
		Value:      "👨‍👨‍👧",
		Descriptor: "family: man, man, girl",
	},
	"1F468-200D-1F468-200D-1F467-200D-1F466": {
		Key:        "1F468-200D-1F468-200D-1F467-200D-1F466",
		Value:      "👨‍👨‍👧‍👦",
		Descriptor: "family: man, man, girl, boy",
	},
	"1F468-200D-1F468-200D-1F466-200D-1F466": {
		Key:        "1F468-200D-1F468-200D-1F466-200D-1F466",
		Value:      "👨‍👨‍👦‍👦",
		Descriptor: "family: man, man, boy, boy",
	},
	"1F468-200D-1F468-200D-1F467-200D-1F467": {
		Key:        "1F468-200D-1F468-200D-1F467-200D-1F467",
		Value:      "👨‍👨‍👧‍👧",
		Descriptor: "family: man, man, girl, girl",
	},
	"1F469-200D-1F469-200D-1F466": {
		Key:        "1F469-200D-1F469-200D-1F466",
		Value:      "👩‍👩‍👦",
		Descriptor: "family: woman, woman, boy",
	},
	"1F469-200D-1F469-200D-1F467": {
		Key:        "1F469-200D-1F469-200D-1F467",
		Value:      "👩‍👩‍👧",
		Descriptor: "family: woman, woman, girl",
	},
	"1F469-200D-1F469-200D-1F467-200D-1F466": {
		Key:        "1F469-200D-1F469-200D-1F467-200D-1F466",
		Value:      "👩‍👩‍👧‍👦",
		Descriptor: "family: woman, woman, girl, boy",
	},
	"1F469-200D-1F469-200D-1F466-200D-1F466": {
		Key:        "1F469-200D-1F469-200D-1F466-200D-1F466",
		Value:      "👩‍👩‍👦‍👦",
		Descriptor: "family: woman, woman, boy, boy",
	},
	"1F469-200D-1F469-200D-1F467-200D-1F467": {
		Key:        "1F469-200D-1F469-200D-1F467-200D-1F467",
		Value:      "👩‍👩‍👧‍👧",
		Descriptor: "family: woman, woman, girl, girl",
	},
	"1F468-200D-1F466": {
		Key:        "1F468-200D-1F466",
		Value:      "👨‍👦",
		Descriptor: "family: man, boy",
	},
	"1F468-200D-1F466-200D-1F466": {
		Key:        "1F468-200D-1F466-200D-1F466",
		Value:      "👨‍👦‍👦",
		Descriptor: "family: man, boy, boy",
	},
	"1F468-200D-1F467": {
		Key:        "1F468-200D-1F467",
		Value:      "👨‍👧",
		Descriptor: "family: man, girl",
	},
	"1F468-200D-1F467-200D-1F466": {
		Key:        "1F468-200D-1F467-200D-1F466",
		Value:      "👨‍👧‍👦",
		Descriptor: "family: man, girl, boy",
	},
	"1F468-200D-1F467-200D-1F467": {
		Key:        "1F468-200D-1F467-200D-1F467",
		Value:      "👨‍👧‍👧",
		Descriptor: "family: man, girl, girl",
	},
	"1F469-200D-1F466": {
		Key:        "1F469-200D-1F466",
		Value:      "👩‍👦",
		Descriptor: "family: woman, boy",
	},
	"1F469-200D-1F466-200D-1F466": {
		Key:        "1F469-200D-1F466-200D-1F466",
		Value:      "👩‍👦‍👦",
		Descriptor: "family: woman, boy, boy",
	},
	"1F469-200D-1F467": {
		Key:        "1F469-200D-1F467",
		Value:      "👩‍👧",
		Descriptor: "family: woman, girl",
	},
	"1F469-200D-1F467-200D-1F466": {
		Key:        "1F469-200D-1F467-200D-1F466",
		Value:      "👩‍👧‍👦",
		Descriptor: "family: woman, girl, boy",
	},
	"1F469-200D-1F467-200D-1F467": {
		Key:        "1F469-200D-1F467-200D-1F467",
		Value:      "👩‍👧‍👧",
		Descriptor: "family: woman, girl, girl",
	},
	"1F5E3-FE0F": {
		Key:        "1F5E3-FE0F",
		Value:      "🗣️",
		Descriptor: "speaking head",
	},
	"1F464": {
		Key:        "1F464",
		Value:      "👤",
		Descriptor: "bust in silhouette",
	},
	"1F465": {
		Key:        "1F465",
		Value:      "👥",
		Descriptor: "busts in silhouette",
	},
	"1FAC2": {
		Key:        "1FAC2",
		Value:      "🫂",
		Descriptor: "people hugging",
	},
	"1F46A": {
		Key:        "1F46A",
		Value:      "👪",
		Descriptor: "family",
	},
	"1F9D1-200D-1F9D1-200D-1F9D2": {
		Key:        "1F9D1-200D-1F9D1-200D-1F9D2",
		Value:      "🧑‍🧑‍🧒",
		Descriptor: "family: adult, adult, child",
	},
	"1F9D1-200D-1F9D1-200D-1F9D2-200D-1F9D2": {
		Key:        "1F9D1-200D-1F9D1-200D-1F9D2-200D-1F9D2",
		Value:      "🧑‍🧑‍🧒‍🧒",
		Descriptor: "family: adult, adult, child, child",
	},
	"1F9D1-200D-1F9D2": {
		Key:        "1F9D1-200D-1F9D2",
		Value:      "🧑‍🧒",
		Descriptor: "family: adult, child",
	},
	"1F9D1-200D-1F9D2-200D-1F9D2": {
		Key:        "1F9D1-200D-1F9D2-200D-1F9D2",
		Value:      "🧑‍🧒‍🧒",
		Descriptor: "family: adult, child, child",
	},
	"1F463": {
		Key:        "1F463",
		Value:      "👣",
		Descriptor: "footprints",
	},
	"1F435": {
		Key:        "1F435",
		Value:      "🐵",
		Descriptor: "monkey face",
	},
	"1F412": {
		Key:        "1F412",
		Value:      "🐒",
		Descriptor: "monkey",
	},
	"1F98D": {
		Key:        "1F98D",
		Value:      "🦍",
		Descriptor: "gorilla",
	},
	"1F9A7": {
		Key:        "1F9A7",
		Value:      "🦧",
		Descriptor: "orangutan",
	},
	"1F436": {
		Key:        "1F436",
		Value:      "🐶",
		Descriptor: "dog face",
	},
	"1F415": {
		Key:        "1F415",
		Value:      "🐕",
		Descriptor: "dog",
	},
	"1F9AE": {
		Key:        "1F9AE",
		Value:      "🦮",
		Descriptor: "guide dog",
	},
	"1F415-200D-1F9BA": {
		Key:        "1F415-200D-1F9BA",
		Value:      "🐕‍🦺",
		Descriptor: "service dog",
	},
	"1F429": {
		Key:        "1F429",
		Value:      "🐩",
		Descriptor: "poodle",
	},
	"1F43A": {
		Key:        "1F43A",
		Value:      "🐺",
		Descriptor: "wolf",
	},
	"1F98A": {
		Key:        "1F98A",
		Value:      "🦊",
		Descriptor: "fox",
	},
	"1F99D": {
		Key:        "1F99D",
		Value:      "🦝",
		Descriptor: "raccoon",
	},
	"1F431": {
		Key:        "1F431",
		Value:      "🐱",
		Descriptor: "cat face",
	},
	"1F408": {
		Key:        "1F408",
		Value:      "🐈",
		Descriptor: "cat",
	},
	"1F408-200D-2B1B": {
		Key:        "1F408-200D-2B1B",
		Value:      "🐈‍⬛",
		Descriptor: "black cat",
	},
	"1F981": {
		Key:        "1F981",
		Value:      "🦁",
		Descriptor: "lion",
	},
	"1F42F": {
		Key:        "1F42F",
		Value:      "🐯",
		Descriptor: "tiger face",
	},
	"1F405": {
		Key:        "1F405",
		Value:      "🐅",
		Descriptor: "tiger",
	},
	"1F406": {
		Key:        "1F406",
		Value:      "🐆",
		Descriptor: "leopard",
	},
	"1F434": {
		Key:        "1F434",
		Value:      "🐴",
		Descriptor: "horse face",
	},
	"1FACE": {
		Key:        "1FACE",
		Value:      "🫎",
		Descriptor: "moose",
	},
	"1FACF": {
		Key:        "1FACF",
		Value:      "🫏",
		Descriptor: "donkey",
	},
	"1F40E": {
		Key:        "1F40E",
		Value:      "🐎",
		Descriptor: "horse",
	},
	"1F984": {
		Key:        "1F984",
		Value:      "🦄",
		Descriptor: "unicorn",
	},
	"1F993": {
		Key:        "1F993",
		Value:      "🦓",
		Descriptor: "zebra",
	},
	"1F98C": {
		Key:        "1F98C",
		Value:      "🦌",
		Descriptor: "deer",
	},
	"1F9AC": {
		Key:        "1F9AC",
		Value:      "🦬",
		Descriptor: "bison",
	},
	"1F42E": {
		Key:        "1F42E",
		Value:      "🐮",
		Descriptor: "cow face",
	},
	"1F402": {
		Key:        "1F402",
		Value:      "🐂",
		Descriptor: "ox",
	},
	"1F403": {
		Key:        "1F403",
		Value:      "🐃",
		Descriptor: "water buffalo",
	},
	"1F404": {
		Key:        "1F404",
		Value:      "🐄",
		Descriptor: "cow",
	},
	"1F437": {
		Key:        "1F437",
		Value:      "🐷",
		Descriptor: "pig face",
	},
	"1F416": {
		Key:        "1F416",
		Value:      "🐖",
		Descriptor: "pig",
	},
	"1F417": {
		Key:        "1F417",
		Value:      "🐗",
		Descriptor: "boar",
	},
	"1F43D": {
		Key:        "1F43D",
		Value:      "🐽",
		Descriptor: "pig nose",
	},
	"1F40F": {
		Key:        "1F40F",
		Value:      "🐏",
		Descriptor: "ram",
	},
	"1F411": {
		Key:        "1F411",
		Value:      "🐑",
		Descriptor: "ewe",
	},
	"1F410": {
		Key:        "1F410",
		Value:      "🐐",
		Descriptor: "goat",
	},
	"1F42A": {
		Key:        "1F42A",
		Value:      "🐪",
		Descriptor: "camel",
	},
	"1F42B": {
		Key:        "1F42B",
		Value:      "🐫",
		Descriptor: "two-hump camel",
	},
	"1F999": {
		Key:        "1F999",
		Value:      "🦙",
		Descriptor: "llama",
	},
	"1F992": {
		Key:        "1F992",
		Value:      "🦒",
		Descriptor: "giraffe",
	},
	"1F418": {
		Key:        "1F418",
		Value:      "🐘",
		Descriptor: "elephant",
	},
	"1F9A3": {
		Key:        "1F9A3",
		Value:      "🦣",
		Descriptor: "mammoth",
	},
	"1F98F": {
		Key:        "1F98F",
		Value:      "🦏",
		Descriptor: "rhinoceros",
	},
	"1F99B": {
		Key:        "1F99B",
		Value:      "🦛",
		Descriptor: "hippopotamus",
	},
	"1F42D": {
		Key:        "1F42D",
		Value:      "🐭",
		Descriptor: "mouse face",
	},
	"1F401": {
		Key:        "1F401",
		Value:      "🐁",
		Descriptor: "mouse",
	},
	"1F400": {
		Key:        "1F400",
		Value:      "🐀",
		Descriptor: "rat",
	},
	"1F439": {
		Key:        "1F439",
		Value:      "🐹",
		Descriptor: "hamster",
	},
	"1F430": {
		Key:        "1F430",
		Value:      "🐰",
		Descriptor: "rabbit face",
	},
	"1F407": {
		Key:        "1F407",
		Value:      "🐇",
		Descriptor: "rabbit",
	},
	"1F43F-FE0F": {
		Key:        "1F43F-FE0F",
		Value:      "🐿️",
		Descriptor: "chipmunk",
	},
	"1F9AB": {
		Key:        "1F9AB",
		Value:      "🦫",
		Descriptor: "beaver",
	},
	"1F994": {
		Key:        "1F994",
		Value:      "🦔",
		Descriptor: "hedgehog",
	},
	"1F987": {
		Key:        "1F987",
		Value:      "🦇",
		Descriptor: "bat",
	},
	"1F43B": {
		Key:        "1F43B",
		Value:      "🐻",
		Descriptor: "bear",
	},
	"1F43B-200D-2744-FE0F": {
		Key:        "1F43B-200D-2744-FE0F",
		Value:      "🐻‍❄️",
		Descriptor: "polar bear",
	},
	"1F428": {
		Key:        "1F428",
		Value:      "🐨",
		Descriptor: "koala",
	},
	"1F43C": {
		Key:        "1F43C",
		Value:      "🐼",
		Descriptor: "panda",
	},
	"1F9A5": {
		Key:        "1F9A5",
		Value:      "🦥",
		Descriptor: "sloth",
	},
	"1F9A6": {
		Key:        "1F9A6",
		Value:      "🦦",
		Descriptor: "otter",
	},
	"1F9A8": {
		Key:        "1F9A8",
		Value:      "🦨",
		Descriptor: "skunk",
	},
	"1F998": {
		Key:        "1F998",
		Value:      "🦘",
		Descriptor: "kangaroo",
	},
	"1F9A1": {
		Key:        "1F9A1",
		Value:      "🦡",
		Descriptor: "badger",
	},
	"1F43E": {
		Key:        "1F43E",
		Value:      "🐾",
		Descriptor: "paw prints",
	},
	"1F983": {
		Key:        "1F983",
		Value:      "🦃",
		Descriptor: "turkey",
	},
	"1F414": {
		Key:        "1F414",
		Value:      "🐔",
		Descriptor: "chicken",
	},
	"1F413": {
		Key:        "1F413",
		Value:      "🐓",
		Descriptor: "rooster",
	},
	"1F423": {
		Key:        "1F423",
		Value:      "🐣",
		Descriptor: "hatching chick",
	},
	"1F424": {
		Key:        "1F424",
		Value:      "🐤",
		Descriptor: "baby chick",
	},
	"1F425": {
		Key:        "1F425",
		Value:      "🐥",
		Descriptor: "front-facing baby chick",
	},
	"1F426": {
		Key:        "1F426",
		Value:      "🐦",
		Descriptor: "bird",
	},
	"1F427": {
		Key:        "1F427",
		Value:      "🐧",
		Descriptor: "penguin",
	},
	"1F54A-FE0F": {
		Key:        "1F54A-FE0F",
		Value:      "🕊️",
		Descriptor: "dove",
	},
	"1F985": {
		Key:        "1F985",
		Value:      "🦅",
		Descriptor: "eagle",
	},
	"1F986": {
		Key:        "1F986",
		Value:      "🦆",
		Descriptor: "duck",
	},
	"1F9A2": {
		Key:        "1F9A2",
		Value:      "🦢",
		Descriptor: "swan",
	},
	"1F989": {
		Key:        "1F989",
		Value:      "🦉",
		Descriptor: "owl",
	},
	"1F9A4": {
		Key:        "1F9A4",
		Value:      "🦤",
		Descriptor: "dodo",
	},
	"1FAB6": {
		Key:        "1FAB6",
		Value:      "🪶",
		Descriptor: "feather",
	},
	"1F9A9": {
		Key:        "1F9A9",
		Value:      "🦩",
		Descriptor: "flamingo",
	},
	"1F99A": {
		Key:        "1F99A",
		Value:      "🦚",
		Descriptor: "peacock",
	},
	"1F99C": {
		Key:        "1F99C",
		Value:      "🦜",
		Descriptor: "parrot",
	},
	"1FABD": {
		Key:        "1FABD",
		Value:      "🪽",
		Descriptor: "wing",
	},
	"1F426-200D-2B1B": {
		Key:        "1F426-200D-2B1B",
		Value:      "🐦‍⬛",
		Descriptor: "black bird",
	},
	"1FABF": {
		Key:        "1FABF",
		Value:      "🪿",
		Descriptor: "goose",
	},
	"1F426-200D-1F525": {
		Key:        "1F426-200D-1F525",
		Value:      "🐦‍🔥",
		Descriptor: "phoenix",
	},
	"1F438": {
		Key:        "1F438",
		Value:      "🐸",
		Descriptor: "frog",
	},
	"1F40A": {
		Key:        "1F40A",
		Value:      "🐊",
		Descriptor: "crocodile",
	},
	"1F422": {
		Key:        "1F422",
		Value:      "🐢",
		Descriptor: "turtle",
	},
	"1F98E": {
		Key:        "1F98E",
		Value:      "🦎",
		Descriptor: "lizard",
	},
	"1F40D": {
		Key:        "1F40D",
		Value:      "🐍",
		Descriptor: "snake",
	},
	"1F432": {
		Key:        "1F432",
		Value:      "🐲",
		Descriptor: "dragon face",
	},
	"1F409": {
		Key:        "1F409",
		Value:      "🐉",
		Descriptor: "dragon",
	},
	"1F995": {
		Key:        "1F995",
		Value:      "🦕",
		Descriptor: "sauropod",
	},
	"1F996": {
		Key:        "1F996",
		Value:      "🦖",
		Descriptor: "T-Rex",
	},
	"1F433": {
		Key:        "1F433",
		Value:      "🐳",
		Descriptor: "spouting whale",
	},
	"1F40B": {
		Key:        "1F40B",
		Value:      "🐋",
		Descriptor: "whale",
	},
	"1F42C": {
		Key:        "1F42C",
		Value:      "🐬",
		Descriptor: "dolphin",
	},
	"1F9AD": {
		Key:        "1F9AD",
		Value:      "🦭",
		Descriptor: "seal",
	},
	"1F41F": {
		Key:        "1F41F",
		Value:      "🐟",
		Descriptor: "fish",
	},
	"1F420": {
		Key:        "1F420",
		Value:      "🐠",
		Descriptor: "tropical fish",
	},
	"1F421": {
		Key:        "1F421",
		Value:      "🐡",
		Descriptor: "blowfish",
	},
	"1F988": {
		Key:        "1F988",
		Value:      "🦈",
		Descriptor: "shark",
	},
	"1F419": {
		Key:        "1F419",
		Value:      "🐙",
		Descriptor: "octopus",
	},
	"1F41A": {
		Key:        "1F41A",
		Value:      "🐚",
		Descriptor: "spiral shell",
	},
	"1FAB8": {
		Key:        "1FAB8",
		Value:      "🪸",
		Descriptor: "coral",
	},
	"1FABC": {
		Key:        "1FABC",
		Value:      "🪼",
		Descriptor: "jellyfish",
	},
	"1F40C": {
		Key:        "1F40C",
		Value:      "🐌",
		Descriptor: "snail",
	},
	"1F98B": {
		Key:        "1F98B",
		Value:      "🦋",
		Descriptor: "butterfly",
	},
	"1F41B": {
		Key:        "1F41B",
		Value:      "🐛",
		Descriptor: "bug",
	},
	"1F41C": {
		Key:        "1F41C",
		Value:      "🐜",
		Descriptor: "ant",
	},
	"1F41D": {
		Key:        "1F41D",
		Value:      "🐝",
		Descriptor: "honeybee",
	},
	"1FAB2": {
		Key:        "1FAB2",
		Value:      "🪲",
		Descriptor: "beetle",
	},
	"1F41E": {
		Key:        "1F41E",
		Value:      "🐞",
		Descriptor: "lady beetle",
	},
	"1F997": {
		Key:        "1F997",
		Value:      "🦗",
		Descriptor: "cricket",
	},
	"1FAB3": {
		Key:        "1FAB3",
		Value:      "🪳",
		Descriptor: "cockroach",
	},
	"1F577-FE0F": {
		Key:        "1F577-FE0F",
		Value:      "🕷️",
		Descriptor: "spider",
	},
	"1F578-FE0F": {
		Key:        "1F578-FE0F",
		Value:      "🕸️",
		Descriptor: "spider web",
	},
	"1F982": {
		Key:        "1F982",
		Value:      "🦂",
		Descriptor: "scorpion",
	},
	"1F99F": {
		Key:        "1F99F",
		Value:      "🦟",
		Descriptor: "mosquito",
	},
	"1FAB0": {
		Key:        "1FAB0",
		Value:      "🪰",
		Descriptor: "fly",
	},
	"1FAB1": {
		Key:        "1FAB1",
		Value:      "🪱",
		Descriptor: "worm",
	},
	"1F9A0": {
		Key:        "1F9A0",
		Value:      "🦠",
		Descriptor: "microbe",
	},
	"1F490": {
		Key:        "1F490",
		Value:      "💐",
		Descriptor: "bouquet",
	},
	"1F338": {
		Key:        "1F338",
		Value:      "🌸",
		Descriptor: "cherry blossom",
	},
	"1F4AE": {
		Key:        "1F4AE",
		Value:      "💮",
		Descriptor: "white flower",
	},
	"1FAB7": {
		Key:        "1FAB7",
		Value:      "🪷",
		Descriptor: "lotus",
	},
	"1F3F5-FE0F": {
		Key:        "1F3F5-FE0F",
		Value:      "🏵️",
		Descriptor: "rosette",
	},
	"1F339": {
		Key:        "1F339",
		Value:      "🌹",
		Descriptor: "rose",
	},
	"1F940": {
		Key:        "1F940",
		Value:      "🥀",
		Descriptor: "wilted flower",
	},
	"1F33A": {
		Key:        "1F33A",
		Value:      "🌺",
		Descriptor: "hibiscus",
	},
	"1F33B": {
		Key:        "1F33B",
		Value:      "🌻",
		Descriptor: "sunflower",
	},
	"1F33C": {
		Key:        "1F33C",
		Value:      "🌼",
		Descriptor: "blossom",
	},
	"1F337": {
		Key:        "1F337",
		Value:      "🌷",
		Descriptor: "tulip",
	},
	"1FABB": {
		Key:        "1FABB",
		Value:      "🪻",
		Descriptor: "hyacinth",
	},
	"1F331": {
		Key:        "1F331",
		Value:      "🌱",
		Descriptor: "seedling",
	},
	"1FAB4": {
		Key:        "1FAB4",
		Value:      "🪴",
		Descriptor: "potted plant",
	},
	"1F332": {
		Key:        "1F332",
		Value:      "🌲",
		Descriptor: "evergreen tree",
	},
	"1F333": {
		Key:        "1F333",
		Value:      "🌳",
		Descriptor: "deciduous tree",
	},
	"1F334": {
		Key:        "1F334",
		Value:      "🌴",
		Descriptor: "palm tree",
	},
	"1F335": {
		Key:        "1F335",
		Value:      "🌵",
		Descriptor: "cactus",
	},
	"1F33E": {
		Key:        "1F33E",
		Value:      "🌾",
		Descriptor: "sheaf of rice",
	},
	"1F33F": {
		Key:        "1F33F",
		Value:      "🌿",
		Descriptor: "herb",
	},
	"2618-FE0F": {
		Key:        "2618-FE0F",
		Value:      "☘️",
		Descriptor: "shamrock",
	},
	"1F340": {
		Key:        "1F340",
		Value:      "🍀",
		Descriptor: "four leaf clover",
	},
	"1F341": {
		Key:        "1F341",
		Value:      "🍁",
		Descriptor: "maple leaf",
	},
	"1F342": {
		Key:        "1F342",
		Value:      "🍂",
		Descriptor: "fallen leaf",
	},
	"1F343": {
		Key:        "1F343",
		Value:      "🍃",
		Descriptor: "leaf fluttering in wind",
	},
	"1FAB9": {
		Key:        "1FAB9",
		Value:      "🪹",
		Descriptor: "empty nest",
	},
	"1FABA": {
		Key:        "1FABA",
		Value:      "🪺",
		Descriptor: "nest with eggs",
	},
	"1F344": {
		Key:        "1F344",
		Value:      "🍄",
		Descriptor: "mushroom",
	},
	"1F347": {
		Key:        "1F347",
		Value:      "🍇",
		Descriptor: "grapes",
	},
	"1F348": {
		Key:        "1F348",
		Value:      "🍈",
		Descriptor: "melon",
	},
	"1F349": {
		Key:        "1F349",
		Value:      "🍉",
		Descriptor: "watermelon",
	},
	"1F34A": {
		Key:        "1F34A",
		Value:      "🍊",
		Descriptor: "tangerine",
	},
	"1F34B": {
		Key:        "1F34B",
		Value:      "🍋",
		Descriptor: "lemon",
	},
	"1F34B-200D-1F7E9": {
		Key:        "1F34B-200D-1F7E9",
		Value:      "🍋‍🟩",
		Descriptor: "lime",
	},
	"1F34C": {
		Key:        "1F34C",
		Value:      "🍌",
		Descriptor: "banana",
	},
	"1F34D": {
		Key:        "1F34D",
		Value:      "🍍",
		Descriptor: "pineapple",
	},
	"1F96D": {
		Key:        "1F96D",
		Value:      "🥭",
		Descriptor: "mango",
	},
	"1F34E": {
		Key:        "1F34E",
		Value:      "🍎",
		Descriptor: "red apple",
	},
	"1F34F": {
		Key:        "1F34F",
		Value:      "🍏",
		Descriptor: "green apple",
	},
	"1F350": {
		Key:        "1F350",
		Value:      "🍐",
		Descriptor: "pear",
	},
	"1F351": {
		Key:        "1F351",
		Value:      "🍑",
		Descriptor: "peach",
	},
	"1F352": {
		Key:        "1F352",
		Value:      "🍒",
		Descriptor: "cherries",
	},
	"1F353": {
		Key:        "1F353",
		Value:      "🍓",
		Descriptor: "strawberry",
	},
	"1FAD0": {
		Key:        "1FAD0",
		Value:      "🫐",
		Descriptor: "blueberries",
	},
	"1F95D": {
		Key:        "1F95D",
		Value:      "🥝",
		Descriptor: "kiwi fruit",
	},
	"1F345": {
		Key:        "1F345",
		Value:      "🍅",
		Descriptor: "tomato",
	},
	"1FAD2": {
		Key:        "1FAD2",
		Value:      "🫒",
		Descriptor: "olive",
	},
	"1F965": {
		Key:        "1F965",
		Value:      "🥥",
		Descriptor: "coconut",
	},
	"1F951": {
		Key:        "1F951",
		Value:      "🥑",
		Descriptor: "avocado",
	},
	"1F346": {
		Key:        "1F346",
		Value:      "🍆",
		Descriptor: "eggplant",
	},
	"1F954": {
		Key:        "1F954",
		Value:      "🥔",
		Descriptor: "potato",
	},
	"1F955": {
		Key:        "1F955",
		Value:      "🥕",
		Descriptor: "carrot",
	},
	"1F33D": {
		Key:        "1F33D",
		Value:      "🌽",
		Descriptor: "ear of corn",
	},
	"1F336-FE0F": {
		Key:        "1F336-FE0F",
		Value:      "🌶️",
		Descriptor: "hot pepper",
	},
	"1FAD1": {
		Key:        "1FAD1",
		Value:      "🫑",
		Descriptor: "bell pepper",
	},
	"1F952": {
		Key:        "1F952",
		Value:      "🥒",
		Descriptor: "cucumber",
	},
	"1F96C": {
		Key:        "1F96C",
		Value:      "🥬",
		Descriptor: "leafy green",
	},
	"1F966": {
		Key:        "1F966",
		Value:      "🥦",
		Descriptor: "broccoli",
	},
	"1F9C4": {
		Key:        "1F9C4",
		Value:      "🧄",
		Descriptor: "garlic",
	},
	"1F9C5": {
		Key:        "1F9C5",
		Value:      "🧅",
		Descriptor: "onion",
	},
	"1F95C": {
		Key:        "1F95C",
		Value:      "🥜",
		Descriptor: "peanuts",
	},
	"1FAD8": {
		Key:        "1FAD8",
		Value:      "🫘",
		Descriptor: "beans",
	},
	"1F330": {
		Key:        "1F330",
		Value:      "🌰",
		Descriptor: "chestnut",
	},
	"1FADA": {
		Key:        "1FADA",
		Value:      "🫚",
		Descriptor: "ginger root",
	},
	"1FADB": {
		Key:        "1FADB",
		Value:      "🫛",
		Descriptor: "pea pod",
	},
	"1F344-200D-1F7EB": {
		Key:        "1F344-200D-1F7EB",
		Value:      "🍄‍🟫",
		Descriptor: "brown mushroom",
	},
	"1F35E": {
		Key:        "1F35E",
		Value:      "🍞",
		Descriptor: "bread",
	},
	"1F950": {
		Key:        "1F950",
		Value:      "🥐",
		Descriptor: "croissant",
	},
	"1F956": {
		Key:        "1F956",
		Value:      "🥖",
		Descriptor: "baguette bread",
	},
	"1FAD3": {
		Key:        "1FAD3",
		Value:      "🫓",
		Descriptor: "flatbread",
	},
	"1F968": {
		Key:        "1F968",
		Value:      "🥨",
		Descriptor: "pretzel",
	},
	"1F96F": {
		Key:        "1F96F",
		Value:      "🥯",
		Descriptor: "bagel",
	},
	"1F95E": {
		Key:        "1F95E",
		Value:      "🥞",
		Descriptor: "pancakes",
	},
	"1F9C7": {
		Key:        "1F9C7",
		Value:      "🧇",
		Descriptor: "waffle",
	},
	"1F9C0": {
		Key:        "1F9C0",
		Value:      "🧀",
		Descriptor: "cheese wedge",
	},
	"1F356": {
		Key:        "1F356",
		Value:      "🍖",
		Descriptor: "meat on bone",
	},
	"1F357": {
		Key:        "1F357",
		Value:      "🍗",
		Descriptor: "poultry leg",
	},
	"1F969": {
		Key:        "1F969",
		Value:      "🥩",
		Descriptor: "cut of meat",
	},
	"1F953": {
		Key:        "1F953",
		Value:      "🥓",
		Descriptor: "bacon",
	},
	"1F354": {
		Key:        "1F354",
		Value:      "🍔",
		Descriptor: "hamburger",
	},
	"1F35F": {
		Key:        "1F35F",
		Value:      "🍟",
		Descriptor: "french fries",
	},
	"1F355": {
		Key:        "1F355",
		Value:      "🍕",
		Descriptor: "pizza",
	},
	"1F32D": {
		Key:        "1F32D",
		Value:      "🌭",
		Descriptor: "hot dog",
	},
	"1F96A": {
		Key:        "1F96A",
		Value:      "🥪",
		Descriptor: "sandwich",
	},
	"1F32E": {
		Key:        "1F32E",
		Value:      "🌮",
		Descriptor: "taco",
	},
	"1F32F": {
		Key:        "1F32F",
		Value:      "🌯",
		Descriptor: "burrito",
	},
	"1FAD4": {
		Key:        "1FAD4",
		Value:      "🫔",
		Descriptor: "tamale",
	},
	"1F959": {
		Key:        "1F959",
		Value:      "🥙",
		Descriptor: "stuffed flatbread",
	},
	"1F9C6": {
		Key:        "1F9C6",
		Value:      "🧆",
		Descriptor: "falafel",
	},
	"1F95A": {
		Key:        "1F95A",
		Value:      "🥚",
		Descriptor: "egg",
	},
	"1F373": {
		Key:        "1F373",
		Value:      "🍳",
		Descriptor: "cooking",
	},
	"1F958": {
		Key:        "1F958",
		Value:      "🥘",
		Descriptor: "shallow pan of food",
	},
	"1F372": {
		Key:        "1F372",
		Value:      "🍲",
		Descriptor: "pot of food",
	},
	"1FAD5": {
		Key:        "1FAD5",
		Value:      "🫕",
		Descriptor: "fondue",
	},
	"1F963": {
		Key:        "1F963",
		Value:      "🥣",
		Descriptor: "bowl with spoon",
	},
	"1F957": {
		Key:        "1F957",
		Value:      "🥗",
		Descriptor: "green salad",
	},
	"1F37F": {
		Key:        "1F37F",
		Value:      "🍿",
		Descriptor: "popcorn",
	},
	"1F9C8": {
		Key:        "1F9C8",
		Value:      "🧈",
		Descriptor: "butter",
	},
	"1F9C2": {
		Key:        "1F9C2",
		Value:      "🧂",
		Descriptor: "salt",
	},
	"1F96B": {
		Key:        "1F96B",
		Value:      "🥫",
		Descriptor: "canned food",
	},
	"1F371": {
		Key:        "1F371",
		Value:      "🍱",
		Descriptor: "bento box",
	},
	"1F358": {
		Key:        "1F358",
		Value:      "🍘",
		Descriptor: "rice cracker",
	},
	"1F359": {
		Key:        "1F359",
		Value:      "🍙",
		Descriptor: "rice ball",
	},
	"1F35A": {
		Key:        "1F35A",
		Value:      "🍚",
		Descriptor: "cooked rice",
	},
	"1F35B": {
		Key:        "1F35B",
		Value:      "🍛",
		Descriptor: "curry rice",
	},
	"1F35C": {
		Key:        "1F35C",
		Value:      "🍜",
		Descriptor: "steaming bowl",
	},
	"1F35D": {
		Key:        "1F35D",
		Value:      "🍝",
		Descriptor: "spaghetti",
	},
	"1F360": {
		Key:        "1F360",
		Value:      "🍠",
		Descriptor: "roasted sweet potato",
	},
	"1F362": {
		Key:        "1F362",
		Value:      "🍢",
		Descriptor: "oden",
	},
	"1F363": {
		Key:        "1F363",
		Value:      "🍣",
		Descriptor: "sushi",
	},
	"1F364": {
		Key:        "1F364",
		Value:      "🍤",
		Descriptor: "fried shrimp",
	},
	"1F365": {
		Key:        "1F365",
		Value:      "🍥",
		Descriptor: "fish cake with swirl",
	},
	"1F96E": {
		Key:        "1F96E",
		Value:      "🥮",
		Descriptor: "moon cake",
	},
	"1F361": {
		Key:        "1F361",
		Value:      "🍡",
		Descriptor: "dango",
	},
	"1F95F": {
		Key:        "1F95F",
		Value:      "🥟",
		Descriptor: "dumpling",
	},
	"1F960": {
		Key:        "1F960",
		Value:      "🥠",
		Descriptor: "fortune cookie",
	},
	"1F961": {
		Key:        "1F961",
		Value:      "🥡",
		Descriptor: "takeout box",
	},
	"1F980": {
		Key:        "1F980",
		Value:      "🦀",
		Descriptor: "crab",
	},
	"1F99E": {
		Key:        "1F99E",
		Value:      "🦞",
		Descriptor: "lobster",
	},
	"1F990": {
		Key:        "1F990",
		Value:      "🦐",
		Descriptor: "shrimp",
	},
	"1F991": {
		Key:        "1F991",
		Value:      "🦑",
		Descriptor: "squid",
	},
	"1F9AA": {
		Key:        "1F9AA",
		Value:      "🦪",
		Descriptor: "oyster",
	},
	"1F366": {
		Key:        "1F366",
		Value:      "🍦",
		Descriptor: "soft ice cream",
	},
	"1F367": {
		Key:        "1F367",
		Value:      "🍧",
		Descriptor: "shaved ice",
	},
	"1F368": {
		Key:        "1F368",
		Value:      "🍨",
		Descriptor: "ice cream",
	},
	"1F369": {
		Key:        "1F369",
		Value:      "🍩",
		Descriptor: "doughnut",
	},
	"1F36A": {
		Key:        "1F36A",
		Value:      "🍪",
		Descriptor: "cookie",
	},
	"1F382": {
		Key:        "1F382",
		Value:      "🎂",
		Descriptor: "birthday cake",
	},
	"1F370": {
		Key:        "1F370",
		Value:      "🍰",
		Descriptor: "shortcake",
	},
	"1F9C1": {
		Key:        "1F9C1",
		Value:      "🧁",
		Descriptor: "cupcake",
	},
	"1F967": {
		Key:        "1F967",
		Value:      "🥧",
		Descriptor: "pie",
	},
	"1F36B": {
		Key:        "1F36B",
		Value:      "🍫",
		Descriptor: "chocolate bar",
	},
	"1F36C": {
		Key:        "1F36C",
		Value:      "🍬",
		Descriptor: "candy",
	},
	"1F36D": {
		Key:        "1F36D",
		Value:      "🍭",
		Descriptor: "lollipop",
	},
	"1F36E": {
		Key:        "1F36E",
		Value:      "🍮",
		Descriptor: "custard",
	},
	"1F36F": {
		Key:        "1F36F",
		Value:      "🍯",
		Descriptor: "honey pot",
	},
	"1F37C": {
		Key:        "1F37C",
		Value:      "🍼",
		Descriptor: "baby bottle",
	},
	"1F95B": {
		Key:        "1F95B",
		Value:      "🥛",
		Descriptor: "glass of milk",
	},
	"2615": {
		Key:        "2615",
		Value:      "☕",
		Descriptor: "hot beverage",
	},
	"1FAD6": {
		Key:        "1FAD6",
		Value:      "🫖",
		Descriptor: "teapot",
	},
	"1F375": {
		Key:        "1F375",
		Value:      "🍵",
		Descriptor: "teacup without handle",
	},
	"1F376": {
		Key:        "1F376",
		Value:      "🍶",
		Descriptor: "sake",
	},
	"1F37E": {
		Key:        "1F37E",
		Value:      "🍾",
		Descriptor: "bottle with popping cork",
	},
	"1F377": {
		Key:        "1F377",
		Value:      "🍷",
		Descriptor: "wine glass",
	},
	"1F378": {
		Key:        "1F378",
		Value:      "🍸",
		Descriptor: "cocktail glass",
	},
	"1F379": {
		Key:        "1F379",
		Value:      "🍹",
		Descriptor: "tropical drink",
	},
	"1F37A": {
		Key:        "1F37A",
		Value:      "🍺",
		Descriptor: "beer mug",
	},
	"1F37B": {
		Key:        "1F37B",
		Value:      "🍻",
		Descriptor: "clinking beer mugs",
	},
	"1F942": {
		Key:        "1F942",
		Value:      "🥂",
		Descriptor: "clinking glasses",
	},
	"1F943": {
		Key:        "1F943",
		Value:      "🥃",
		Descriptor: "tumbler glass",
	},
	"1FAD7": {
		Key:        "1FAD7",
		Value:      "🫗",
		Descriptor: "pouring liquid",
	},
	"1F964": {
		Key:        "1F964",
		Value:      "🥤",
		Descriptor: "cup with straw",
	},
	"1F9CB": {
		Key:        "1F9CB",
		Value:      "🧋",
		Descriptor: "bubble tea",
	},
	"1F9C3": {
		Key:        "1F9C3",
		Value:      "🧃",
		Descriptor: "beverage box",
	},
	"1F9C9": {
		Key:        "1F9C9",
		Value:      "🧉",
		Descriptor: "mate",
	},
	"1F9CA": {
		Key:        "1F9CA",
		Value:      "🧊",
		Descriptor: "ice",
	},
	"1F962": {
		Key:        "1F962",
		Value:      "🥢",
		Descriptor: "chopsticks",
	},
	"1F37D-FE0F": {
		Key:        "1F37D-FE0F",
		Value:      "🍽️",
		Descriptor: "fork and knife with plate",
	},
	"1F374": {
		Key:        "1F374",
		Value:      "🍴",
		Descriptor: "fork and knife",
	},
	"1F944": {
		Key:        "1F944",
		Value:      "🥄",
		Descriptor: "spoon",
	},
	"1F52A": {
		Key:        "1F52A",
		Value:      "🔪",
		Descriptor: "kitchen knife",
	},
	"1FAD9": {
		Key:        "1FAD9",
		Value:      "🫙",
		Descriptor: "jar",
	},
	"1F3FA": {
		Key:        "1F3FA",
		Value:      "🏺",
		Descriptor: "amphora",
	},
	"1F30D": {
		Key:        "1F30D",
		Value:      "🌍",
		Descriptor: "globe showing Europe-Africa",
	},
	"1F30E": {
		Key:        "1F30E",
		Value:      "🌎",
		Descriptor: "globe showing Americas",
	},
	"1F30F": {
		Key:        "1F30F",
		Value:      "🌏",
		Descriptor: "globe showing Asia-Australia",
	},
	"1F310": {
		Key:        "1F310",
		Value:      "🌐",
		Descriptor: "globe with meridians",
	},
	"1F5FA-FE0F": {
		Key:        "1F5FA-FE0F",
		Value:      "🗺️",
		Descriptor: "world map",
	},
	"1F5FE": {
		Key:        "1F5FE",
		Value:      "🗾",
		Descriptor: "map of Japan",
	},
	"1F9ED": {
		Key:        "1F9ED",
		Value:      "🧭",
		Descriptor: "compass",
	},
	"1F3D4-FE0F": {
		Key:        "1F3D4-FE0F",
		Value:      "🏔️",
		Descriptor: "snow-capped mountain",
	},
	"26F0-FE0F": {
		Key:        "26F0-FE0F",
		Value:      "⛰️",
		Descriptor: "mountain",
	},
	"1F30B": {
		Key:        "1F30B",
		Value:      "🌋",
		Descriptor: "volcano",
	},
	"1F5FB": {
		Key:        "1F5FB",
		Value:      "🗻",
		Descriptor: "mount fuji",
	},
	"1F3D5-FE0F": {
		Key:        "1F3D5-FE0F",
		Value:      "🏕️",
		Descriptor: "camping",
	},
	"1F3D6-FE0F": {
		Key:        "1F3D6-FE0F",
		Value:      "🏖️",
		Descriptor: "beach with umbrella",
	},
	"1F3DC-FE0F": {
		Key:        "1F3DC-FE0F",
		Value:      "🏜️",
		Descriptor: "desert",
	},
	"1F3DD-FE0F": {
		Key:        "1F3DD-FE0F",
		Value:      "🏝️",
		Descriptor: "desert island",
	},
	"1F3DE-FE0F": {
		Key:        "1F3DE-FE0F",
		Value:      "🏞️",
		Descriptor: "national park",
	},
	"1F3DF-FE0F": {
		Key:        "1F3DF-FE0F",
		Value:      "🏟️",
		Descriptor: "stadium",
	},
	"1F3DB-FE0F": {
		Key:        "1F3DB-FE0F",
		Value:      "🏛️",
		Descriptor: "classical building",
	},
	"1F3D7-FE0F": {
		Key:        "1F3D7-FE0F",
		Value:      "🏗️",
		Descriptor: "building construction",
	},
	"1F9F1": {
		Key:        "1F9F1",
		Value:      "🧱",
		Descriptor: "brick",
	},
	"1FAA8": {
		Key:        "1FAA8",
		Value:      "🪨",
		Descriptor: "rock",
	},
	"1FAB5": {
		Key:        "1FAB5",
		Value:      "🪵",
		Descriptor: "wood",
	},
	"1F6D6": {
		Key:        "1F6D6",
		Value:      "🛖",
		Descriptor: "hut",
	},
	"1F3D8-FE0F": {
		Key:        "1F3D8-FE0F",
		Value:      "🏘️",
		Descriptor: "houses",
	},
	"1F3DA-FE0F": {
		Key:        "1F3DA-FE0F",
		Value:      "🏚️",
		Descriptor: "derelict house",
	},
	"1F3E0": {
		Key:        "1F3E0",
		Value:      "🏠",
		Descriptor: "house",
	},
	"1F3E1": {
		Key:        "1F3E1",
		Value:      "🏡",
		Descriptor: "house with garden",
	},
	"1F3E2": {
		Key:        "1F3E2",
		Value:      "🏢",
		Descriptor: "office building",
	},
	"1F3E3": {
		Key:        "1F3E3",
		Value:      "🏣",
		Descriptor: "Japanese post office",
	},
	"1F3E4": {
		Key:        "1F3E4",
		Value:      "🏤",
		Descriptor: "post office",
	},
	"1F3E5": {
		Key:        "1F3E5",
		Value:      "🏥",
		Descriptor: "hospital",
	},
	"1F3E6": {
		Key:        "1F3E6",
		Value:      "🏦",
		Descriptor: "bank",
	},
	"1F3E8": {
		Key:        "1F3E8",
		Value:      "🏨",
		Descriptor: "hotel",
	},
	"1F3E9": {
		Key:        "1F3E9",
		Value:      "🏩",
		Descriptor: "love hotel",
	},
	"1F3EA": {
		Key:        "1F3EA",
		Value:      "🏪",
		Descriptor: "convenience store",
	},
	"1F3EB": {
		Key:        "1F3EB",
		Value:      "🏫",
		Descriptor: "school",
	},
	"1F3EC": {
		Key:        "1F3EC",
		Value:      "🏬",
		Descriptor: "department store",
	},
	"1F3ED": {
		Key:        "1F3ED",
		Value:      "🏭",
		Descriptor: "factory",
	},
	"1F3EF": {
		Key:        "1F3EF",
		Value:      "🏯",
		Descriptor: "Japanese castle",
	},
	"1F3F0": {
		Key:        "1F3F0",
		Value:      "🏰",
		Descriptor: "castle",
	},
	"1F492": {
		Key:        "1F492",
		Value:      "💒",
		Descriptor: "wedding",
	},
	"1F5FC": {
		Key:        "1F5FC",
		Value:      "🗼",
		Descriptor: "Tokyo tower",
	},
	"1F5FD": {
		Key:        "1F5FD",
		Value:      "🗽",
		Descriptor: "Statue of Liberty",
	},
	"26EA": {
		Key:        "26EA",
		Value:      "⛪",
		Descriptor: "church",
	},
	"1F54C": {
		Key:        "1F54C",
		Value:      "🕌",
		Descriptor: "mosque",
	},
	"1F6D5": {
		Key:        "1F6D5",
		Value:      "🛕",
		Descriptor: "hindu temple",
	},
	"1F54D": {
		Key:        "1F54D",
		Value:      "🕍",
		Descriptor: "synagogue",
	},
	"26E9-FE0F": {
		Key:        "26E9-FE0F",
		Value:      "⛩️",
		Descriptor: "shinto shrine",
	},
	"1F54B": {
		Key:        "1F54B",
		Value:      "🕋",
		Descriptor: "kaaba",
	},
	"26F2": {
		Key:        "26F2",
		Value:      "⛲",
		Descriptor: "fountain",
	},
	"26FA": {
		Key:        "26FA",
		Value:      "⛺",
		Descriptor: "tent",
	},
	"1F301": {
		Key:        "1F301",
		Value:      "🌁",
		Descriptor: "foggy",
	},
	"1F303": {
		Key:        "1F303",
		Value:      "🌃",
		Descriptor: "night with stars",
	},
	"1F3D9-FE0F": {
		Key:        "1F3D9-FE0F",
		Value:      "🏙️",
		Descriptor: "cityscape",
	},
	"1F304": {
		Key:        "1F304",
		Value:      "🌄",
		Descriptor: "sunrise over mountains",
	},
	"1F305": {
		Key:        "1F305",
		Value:      "🌅",
		Descriptor: "sunrise",
	},
	"1F306": {
		Key:        "1F306",
		Value:      "🌆",
		Descriptor: "cityscape at dusk",
	},
	"1F307": {
		Key:        "1F307",
		Value:      "🌇",
		Descriptor: "sunset",
	},
	"1F309": {
		Key:        "1F309",
		Value:      "🌉",
		Descriptor: "bridge at night",
	},
	"2668-FE0F": {
		Key:        "2668-FE0F",
		Value:      "♨️",
		Descriptor: "hot springs",
	},
	"1F3A0": {
		Key:        "1F3A0",
		Value:      "🎠",
		Descriptor: "carousel horse",
	},
	"1F6DD": {
		Key:        "1F6DD",
		Value:      "🛝",
		Descriptor: "playground slide",
	},
	"1F3A1": {
		Key:        "1F3A1",
		Value:      "🎡",
		Descriptor: "ferris wheel",
	},
	"1F3A2": {
		Key:        "1F3A2",
		Value:      "🎢",
		Descriptor: "roller coaster",
	},
	"1F488": {
		Key:        "1F488",
		Value:      "💈",
		Descriptor: "barber pole",
	},
	"1F3AA": {
		Key:        "1F3AA",
		Value:      "🎪",
		Descriptor: "circus tent",
	},
	"1F682": {
		Key:        "1F682",
		Value:      "🚂",
		Descriptor: "locomotive",
	},
	"1F683": {
		Key:        "1F683",
		Value:      "🚃",
		Descriptor: "railway car",
	},
	"1F684": {
		Key:        "1F684",
		Value:      "🚄",
		Descriptor: "high-speed train",
	},
	"1F685": {
		Key:        "1F685",
		Value:      "🚅",
		Descriptor: "bullet train",
	},
	"1F686": {
		Key:        "1F686",
		Value:      "🚆",
		Descriptor: "train",
	},
	"1F687": {
		Key:        "1F687",
		Value:      "🚇",
		Descriptor: "metro",
	},
	"1F688": {
		Key:        "1F688",
		Value:      "🚈",
		Descriptor: "light rail",
	},
	"1F689": {
		Key:        "1F689",
		Value:      "🚉",
		Descriptor: "station",
	},
	"1F68A": {
		Key:        "1F68A",
		Value:      "🚊",
		Descriptor: "tram",
	},
	"1F69D": {
		Key:        "1F69D",
		Value:      "🚝",
		Descriptor: "monorail",
	},
	"1F69E": {
		Key:        "1F69E",
		Value:      "🚞",
		Descriptor: "mountain railway",
	},
	"1F68B": {
		Key:        "1F68B",
		Value:      "🚋",
		Descriptor: "tram car",
	},
	"1F68C": {
		Key:        "1F68C",
		Value:      "🚌",
		Descriptor: "bus",
	},
	"1F68D": {
		Key:        "1F68D",
		Value:      "🚍",
		Descriptor: "oncoming bus",
	},
	"1F68E": {
		Key:        "1F68E",
		Value:      "🚎",
		Descriptor: "trolleybus",
	},
	"1F690": {
		Key:        "1F690",
		Value:      "🚐",
		Descriptor: "minibus",
	},
	"1F691": {
		Key:        "1F691",
		Value:      "🚑",
		Descriptor: "ambulance",
	},
	"1F692": {
		Key:        "1F692",
		Value:      "🚒",
		Descriptor: "fire engine",
	},
	"1F693": {
		Key:        "1F693",
		Value:      "🚓",
		Descriptor: "police car",
	},
	"1F694": {
		Key:        "1F694",
		Value:      "🚔",
		Descriptor: "oncoming police car",
	},
	"1F695": {
		Key:        "1F695",
		Value:      "🚕",
		Descriptor: "taxi",
	},
	"1F696": {
		Key:        "1F696",
		Value:      "🚖",
		Descriptor: "oncoming taxi",
	},
	"1F697": {
		Key:        "1F697",
		Value:      "🚗",
		Descriptor: "automobile",
	},
	"1F698": {
		Key:        "1F698",
		Value:      "🚘",
		Descriptor: "oncoming automobile",
	},
	"1F699": {
		Key:        "1F699",
		Value:      "🚙",
		Descriptor: "sport utility vehicle",
	},
	"1F6FB": {
		Key:        "1F6FB",
		Value:      "🛻",
		Descriptor: "pickup truck",
	},
	"1F69A": {
		Key:        "1F69A",
		Value:      "🚚",
		Descriptor: "delivery truck",
	},
	"1F69B": {
		Key:        "1F69B",
		Value:      "🚛",
		Descriptor: "articulated lorry",
	},
	"1F69C": {
		Key:        "1F69C",
		Value:      "🚜",
		Descriptor: "tractor",
	},
	"1F3CE-FE0F": {
		Key:        "1F3CE-FE0F",
		Value:      "🏎️",
		Descriptor: "racing car",
	},
	"1F3CD-FE0F": {
		Key:        "1F3CD-FE0F",
		Value:      "🏍️",
		Descriptor: "motorcycle",
	},
	"1F6F5": {
		Key:        "1F6F5",
		Value:      "🛵",
		Descriptor: "motor scooter",
	},
	"1F9BD": {
		Key:        "1F9BD",
		Value:      "🦽",
		Descriptor: "manual wheelchair",
	},
	"1F9BC": {
		Key:        "1F9BC",
		Value:      "🦼",
		Descriptor: "motorized wheelchair",
	},
	"1F6FA": {
		Key:        "1F6FA",
		Value:      "🛺",
		Descriptor: "auto rickshaw",
	},
	"1F6B2": {
		Key:        "1F6B2",
		Value:      "🚲",
		Descriptor: "bicycle",
	},
	"1F6F4": {
		Key:        "1F6F4",
		Value:      "🛴",
		Descriptor: "kick scooter",
	},
	"1F6F9": {
		Key:        "1F6F9",
		Value:      "🛹",
		Descriptor: "skateboard",
	},
	"1F6FC": {
		Key:        "1F6FC",
		Value:      "🛼",
		Descriptor: "roller skate",
	},
	"1F68F": {
		Key:        "1F68F",
		Value:      "🚏",
		Descriptor: "bus stop",
	},
	"1F6E3-FE0F": {
		Key:        "1F6E3-FE0F",
		Value:      "🛣️",
		Descriptor: "motorway",
	},
	"1F6E4-FE0F": {
		Key:        "1F6E4-FE0F",
		Value:      "🛤️",
		Descriptor: "railway track",
	},
	"1F6E2-FE0F": {
		Key:        "1F6E2-FE0F",
		Value:      "🛢️",
		Descriptor: "oil drum",
	},
	"26FD": {
		Key:        "26FD",
		Value:      "⛽",
		Descriptor: "fuel pump",
	},
	"1F6DE": {
		Key:        "1F6DE",
		Value:      "🛞",
		Descriptor: "wheel",
	},
	"1F6A8": {
		Key:        "1F6A8",
		Value:      "🚨",
		Descriptor: "police car light",
	},
	"1F6A5": {
		Key:        "1F6A5",
		Value:      "🚥",
		Descriptor: "horizontal traffic light",
	},
	"1F6A6": {
		Key:        "1F6A6",
		Value:      "🚦",
		Descriptor: "vertical traffic light",
	},
	"1F6D1": {
		Key:        "1F6D1",
		Value:      "🛑",
		Descriptor: "stop sign",
	},
	"1F6A7": {
		Key:        "1F6A7",
		Value:      "🚧",
		Descriptor: "construction",
	},
	"2693": {
		Key:        "2693",
		Value:      "⚓",
		Descriptor: "anchor",
	},
	"1F6DF": {
		Key:        "1F6DF",
		Value:      "🛟",
		Descriptor: "ring buoy",
	},
	"26F5": {
		Key:        "26F5",
		Value:      "⛵",
		Descriptor: "sailboat",
	},
	"1F6F6": {
		Key:        "1F6F6",
		Value:      "🛶",
		Descriptor: "canoe",
	},
	"1F6A4": {
		Key:        "1F6A4",
		Value:      "🚤",
		Descriptor: "speedboat",
	},
	"1F6F3-FE0F": {
		Key:        "1F6F3-FE0F",
		Value:      "🛳️",
		Descriptor: "passenger ship",
	},
	"26F4-FE0F": {
		Key:        "26F4-FE0F",
		Value:      "⛴️",
		Descriptor: "ferry",
	},
	"1F6E5-FE0F": {
		Key:        "1F6E5-FE0F",
		Value:      "🛥️",
		Descriptor: "motor boat",
	},
	"1F6A2": {
		Key:        "1F6A2",
		Value:      "🚢",
		Descriptor: "ship",
	},
	"2708-FE0F": {
		Key:        "2708-FE0F",
		Value:      "✈️",
		Descriptor: "airplane",
	},
	"1F6E9-FE0F": {
		Key:        "1F6E9-FE0F",
		Value:      "🛩️",
		Descriptor: "small airplane",
	},
	"1F6EB": {
		Key:        "1F6EB",
		Value:      "🛫",
		Descriptor: "airplane departure",
	},
	"1F6EC": {
		Key:        "1F6EC",
		Value:      "🛬",
		Descriptor: "airplane arrival",
	},
	"1FA82": {
		Key:        "1FA82",
		Value:      "🪂",
		Descriptor: "parachute",
	},
	"1F4BA": {
		Key:        "1F4BA",
		Value:      "💺",
		Descriptor: "seat",
	},
	"1F681": {
		Key:        "1F681",
		Value:      "🚁",
		Descriptor: "helicopter",
	},
	"1F69F": {
		Key:        "1F69F",
		Value:      "🚟",
		Descriptor: "suspension railway",
	},
	"1F6A0": {
		Key:        "1F6A0",
		Value:      "🚠",
		Descriptor: "mountain cableway",
	},
	"1F6A1": {
		Key:        "1F6A1",
		Value:      "🚡",
		Descriptor: "aerial tramway",
	},
	"1F6F0-FE0F": {
		Key:        "1F6F0-FE0F",
		Value:      "🛰️",
		Descriptor: "satellite",
	},
	"1F680": {
		Key:        "1F680",
		Value:      "🚀",
		Descriptor: "rocket",
	},
	"1F6F8": {
		Key:        "1F6F8",
		Value:      "🛸",
		Descriptor: "flying saucer",
	},
	"1F6CE-FE0F": {
		Key:        "1F6CE-FE0F",
		Value:      "🛎️",
		Descriptor: "bellhop bell",
	},
	"1F9F3": {
		Key:        "1F9F3",
		Value:      "🧳",
		Descriptor: "luggage",
	},
	"231B": {
		Key:        "231B",
		Value:      "⌛",
		Descriptor: "hourglass done",
	},
	"23F3": {
		Key:        "23F3",
		Value:      "⏳",
		Descriptor: "hourglass not done",
	},
	"231A": {
		Key:        "231A",
		Value:      "⌚",
		Descriptor: "watch",
	},
	"23F0": {
		Key:        "23F0",
		Value:      "⏰",
		Descriptor: "alarm clock",
	},
	"23F1-FE0F": {
		Key:        "23F1-FE0F",
		Value:      "⏱️",
		Descriptor: "stopwatch",
	},
	"23F2-FE0F": {
		Key:        "23F2-FE0F",
		Value:      "⏲️",
		Descriptor: "timer clock",
	},
	"1F570-FE0F": {
		Key:        "1F570-FE0F",
		Value:      "🕰️",
		Descriptor: "mantelpiece clock",
	},
	"1F55B": {
		Key:        "1F55B",
		Value:      "🕛",
		Descriptor: "twelve o’clock",
	},
	"1F567": {
		Key:        "1F567",
		Value:      "🕧",
		Descriptor: "twelve-thirty",
	},
	"1F550": {
		Key:        "1F550",
		Value:      "🕐",
		Descriptor: "one o’clock",
	},
	"1F55C": {
		Key:        "1F55C",
		Value:      "🕜",
		Descriptor: "one-thirty",
	},
	"1F551": {
		Key:        "1F551",
		Value:      "🕑",
		Descriptor: "two o’clock",
	},
	"1F55D": {
		Key:        "1F55D",
		Value:      "🕝",
		Descriptor: "two-thirty",
	},
	"1F552": {
		Key:        "1F552",
		Value:      "🕒",
		Descriptor: "three o’clock",
	},
	"1F55E": {
		Key:        "1F55E",
		Value:      "🕞",
		Descriptor: "three-thirty",
	},
	"1F553": {
		Key:        "1F553",
		Value:      "🕓",
		Descriptor: "four o’clock",
	},
	"1F55F": {
		Key:        "1F55F",
		Value:      "🕟",
		Descriptor: "four-thirty",
	},
	"1F554": {
		Key:        "1F554",
		Value:      "🕔",
		Descriptor: "five o’clock",
	},
	"1F560": {
		Key:        "1F560",
		Value:      "🕠",
		Descriptor: "five-thirty",
	},
	"1F555": {
		Key:        "1F555",
		Value:      "🕕",
		Descriptor: "six o’clock",
	},
	"1F561": {
		Key:        "1F561",
		Value:      "🕡",
		Descriptor: "six-thirty",
	},
	"1F556": {
		Key:        "1F556",
		Value:      "🕖",
		Descriptor: "seven o’clock",
	},
	"1F562": {
		Key:        "1F562",
		Value:      "🕢",
		Descriptor: "seven-thirty",
	},
	"1F557": {
		Key:        "1F557",
		Value:      "🕗",
		Descriptor: "eight o’clock",
	},
	"1F563": {
		Key:        "1F563",
		Value:      "🕣",
		Descriptor: "eight-thirty",
	},
	"1F558": {
		Key:        "1F558",
		Value:      "🕘",
		Descriptor: "nine o’clock",
	},
	"1F564": {
		Key:        "1F564",
		Value:      "🕤",
		Descriptor: "nine-thirty",
	},
	"1F559": {
		Key:        "1F559",
		Value:      "🕙",
		Descriptor: "ten o’clock",
	},
	"1F565": {
		Key:        "1F565",
		Value:      "🕥",
		Descriptor: "ten-thirty",
	},
	"1F55A": {
		Key:        "1F55A",
		Value:      "🕚",
		Descriptor: "eleven o’clock",
	},
	"1F566": {
		Key:        "1F566",
		Value:      "🕦",
		Descriptor: "eleven-thirty",
	},
	"1F311": {
		Key:        "1F311",
		Value:      "🌑",
		Descriptor: "new moon",
	},
	"1F312": {
		Key:        "1F312",
		Value:      "🌒",
		Descriptor: "waxing crescent moon",
	},
	"1F313": {
		Key:        "1F313",
		Value:      "🌓",
		Descriptor: "first quarter moon",
	},
	"1F314": {
		Key:        "1F314",
		Value:      "🌔",
		Descriptor: "waxing gibbous moon",
	},
	"1F315": {
		Key:        "1F315",
		Value:      "🌕",
		Descriptor: "full moon",
	},
	"1F316": {
		Key:        "1F316",
		Value:      "🌖",
		Descriptor: "waning gibbous moon",
	},
	"1F317": {
		Key:        "1F317",
		Value:      "🌗",
		Descriptor: "last quarter moon",
	},
	"1F318": {
		Key:        "1F318",
		Value:      "🌘",
		Descriptor: "waning crescent moon",
	},
	"1F319": {
		Key:        "1F319",
		Value:      "🌙",
		Descriptor: "crescent moon",
	},
	"1F31A": {
		Key:        "1F31A",
		Value:      "🌚",
		Descriptor: "new moon face",
	},
	"1F31B": {
		Key:        "1F31B",
		Value:      "🌛",
		Descriptor: "first quarter moon face",
	},
	"1F31C": {
		Key:        "1F31C",
		Value:      "🌜",
		Descriptor: "last quarter moon face",
	},
	"1F321-FE0F": {
		Key:        "1F321-FE0F",
		Value:      "🌡️",
		Descriptor: "thermometer",
	},
	"2600-FE0F": {
		Key:        "2600-FE0F",
		Value:      "☀️",
		Descriptor: "sun",
	},
	"1F31D": {
		Key:        "1F31D",
		Value:      "🌝",
		Descriptor: "full moon face",
	},
	"1F31E": {
		Key:        "1F31E",
		Value:      "🌞",
		Descriptor: "sun with face",
	},
	"1FA90": {
		Key:        "1FA90",
		Value:      "🪐",
		Descriptor: "ringed planet",
	},
	"2B50": {
		Key:        "2B50",
		Value:      "⭐",
		Descriptor: "star",
	},
	"1F31F": {
		Key:        "1F31F",
		Value:      "🌟",
		Descriptor: "glowing star",
	},
	"1F320": {
		Key:        "1F320",
		Value:      "🌠",
		Descriptor: "shooting star",
	},
	"1F30C": {
		Key:        "1F30C",
		Value:      "🌌",
		Descriptor: "milky way",
	},
	"2601-FE0F": {
		Key:        "2601-FE0F",
		Value:      "☁️",
		Descriptor: "cloud",
	},
	"26C5": {
		Key:        "26C5",
		Value:      "⛅",
		Descriptor: "sun behind cloud",
	},
	"26C8-FE0F": {
		Key:        "26C8-FE0F",
		Value:      "⛈️",
		Descriptor: "cloud with lightning and rain",
	},
	"1F324-FE0F": {
		Key:        "1F324-FE0F",
		Value:      "🌤️",
		Descriptor: "sun behind small cloud",
	},
	"1F325-FE0F": {
		Key:        "1F325-FE0F",
		Value:      "🌥️",
		Descriptor: "sun behind large cloud",
	},
	"1F326-FE0F": {
		Key:        "1F326-FE0F",
		Value:      "🌦️",
		Descriptor: "sun behind rain cloud",
	},
	"1F327-FE0F": {
		Key:        "1F327-FE0F",
		Value:      "🌧️",
		Descriptor: "cloud with rain",
	},
	"1F328-FE0F": {
		Key:        "1F328-FE0F",
		Value:      "🌨️",
		Descriptor: "cloud with snow",
	},
	"1F329-FE0F": {
		Key:        "1F329-FE0F",
		Value:      "🌩️",
		Descriptor: "cloud with lightning",
	},
	"1F32A-FE0F": {
		Key:        "1F32A-FE0F",
		Value:      "🌪️",
		Descriptor: "tornado",
	},
	"1F32B-FE0F": {
		Key:        "1F32B-FE0F",
		Value:      "🌫️",
		Descriptor: "fog",
	},
	"1F32C-FE0F": {
		Key:        "1F32C-FE0F",
		Value:      "🌬️",
		Descriptor: "wind face",
	},
	"1F300": {
		Key:        "1F300",
		Value:      "🌀",
		Descriptor: "cyclone",
	},
	"1F308": {
		Key:        "1F308",
		Value:      "🌈",
		Descriptor: "rainbow",
	},
	"1F302": {
		Key:        "1F302",
		Value:      "🌂",
		Descriptor: "closed umbrella",
	},
	"2602-FE0F": {
		Key:        "2602-FE0F",
		Value:      "☂️",
		Descriptor: "umbrella",
	},
	"2614": {
		Key:        "2614",
		Value:      "☔",
		Descriptor: "umbrella with rain drops",
	},
	"26F1-FE0F": {
		Key:        "26F1-FE0F",
		Value:      "⛱️",
		Descriptor: "umbrella on ground",
	},
	"26A1": {
		Key:        "26A1",
		Value:      "⚡",
		Descriptor: "high voltage",
	},
	"2744-FE0F": {
		Key:        "2744-FE0F",
		Value:      "❄️",
		Descriptor: "snowflake",
	},
	"2603-FE0F": {
		Key:        "2603-FE0F",
		Value:      "☃️",
		Descriptor: "snowman",
	},
	"26C4": {
		Key:        "26C4",
		Value:      "⛄",
		Descriptor: "snowman without snow",
	},
	"2604-FE0F": {
		Key:        "2604-FE0F",
		Value:      "☄️",
		Descriptor: "comet",
	},
	"1F525": {
		Key:        "1F525",
		Value:      "🔥",
		Descriptor: "fire",
	},
	"1F4A7": {
		Key:        "1F4A7",
		Value:      "💧",
		Descriptor: "droplet",
	},
	"1F30A": {
		Key:        "1F30A",
		Value:      "🌊",
		Descriptor: "water wave",
	},
	"1F383": {
		Key:        "1F383",
		Value:      "🎃",
		Descriptor: "jack-o-lantern",
	},
	"1F384": {
		Key:        "1F384",
		Value:      "🎄",
		Descriptor: "Christmas tree",
	},
	"1F386": {
		Key:        "1F386",
		Value:      "🎆",
		Descriptor: "fireworks",
	},
	"1F387": {
		Key:        "1F387",
		Value:      "🎇",
		Descriptor: "sparkler",
	},
	"1F9E8": {
		Key:        "1F9E8",
		Value:      "🧨",
		Descriptor: "firecracker",
	},
	"2728": {
		Key:        "2728",
		Value:      "✨",
		Descriptor: "sparkles",
	},
	"1F388": {
		Key:        "1F388",
		Value:      "🎈",
		Descriptor: "balloon",
	},
	"1F389": {
		Key:        "1F389",
		Value:      "🎉",
		Descriptor: "party popper",
	},
	"1F38A": {
		Key:        "1F38A",
		Value:      "🎊",
		Descriptor: "confetti ball",
	},
	"1F38B": {
		Key:        "1F38B",
		Value:      "🎋",
		Descriptor: "tanabata tree",
	},
	"1F38D": {
		Key:        "1F38D",
		Value:      "🎍",
		Descriptor: "pine decoration",
	},
	"1F38E": {
		Key:        "1F38E",
		Value:      "🎎",
		Descriptor: "Japanese dolls",
	},
	"1F38F": {
		Key:        "1F38F",
		Value:      "🎏",
		Descriptor: "carp streamer",
	},
	"1F390": {
		Key:        "1F390",
		Value:      "🎐",
		Descriptor: "wind chime",
	},
	"1F391": {
		Key:        "1F391",
		Value:      "🎑",
		Descriptor: "moon viewing ceremony",
	},
	"1F9E7": {
		Key:        "1F9E7",
		Value:      "🧧",
		Descriptor: "red envelope",
	},
	"1F380": {
		Key:        "1F380",
		Value:      "🎀",
		Descriptor: "ribbon",
	},
	"1F381": {
		Key:        "1F381",
		Value:      "🎁",
		Descriptor: "wrapped gift",
	},
	"1F397-FE0F": {
		Key:        "1F397-FE0F",
		Value:      "🎗️",
		Descriptor: "reminder ribbon",
	},
	"1F39F-FE0F": {
		Key:        "1F39F-FE0F",
		Value:      "🎟️",
		Descriptor: "admission tickets",
	},
	"1F3AB": {
		Key:        "1F3AB",
		Value:      "🎫",
		Descriptor: "ticket",
	},
	"1F396-FE0F": {
		Key:        "1F396-FE0F",
		Value:      "🎖️",
		Descriptor: "military medal",
	},
	"1F3C6": {
		Key:        "1F3C6",
		Value:      "🏆",
		Descriptor: "trophy",
	},
	"1F3C5": {
		Key:        "1F3C5",
		Value:      "🏅",
		Descriptor: "sports medal",
	},
	"1F947": {
		Key:        "1F947",
		Value:      "🥇",
		Descriptor: "1st place medal",
	},
	"1F948": {
		Key:        "1F948",
		Value:      "🥈",
		Descriptor: "2nd place medal",
	},
	"1F949": {
		Key:        "1F949",
		Value:      "🥉",
		Descriptor: "3rd place medal",
	},
	"26BD": {
		Key:        "26BD",
		Value:      "⚽",
		Descriptor: "soccer ball",
	},
	"26BE": {
		Key:        "26BE",
		Value:      "⚾",
		Descriptor: "baseball",
	},
	"1F94E": {
		Key:        "1F94E",
		Value:      "🥎",
		Descriptor: "softball",
	},
	"1F3C0": {
		Key:        "1F3C0",
		Value:      "🏀",
		Descriptor: "basketball",
	},
	"1F3D0": {
		Key:        "1F3D0",
		Value:      "🏐",
		Descriptor: "volleyball",
	},
	"1F3C8": {
		Key:        "1F3C8",
		Value:      "🏈",
		Descriptor: "american football",
	},
	"1F3C9": {
		Key:        "1F3C9",
		Value:      "🏉",
		Descriptor: "rugby football",
	},
	"1F3BE": {
		Key:        "1F3BE",
		Value:      "🎾",
		Descriptor: "tennis",
	},
	"1F94F": {
		Key:        "1F94F",
		Value:      "🥏",
		Descriptor: "flying disc",
	},
	"1F3B3": {
		Key:        "1F3B3",
		Value:      "🎳",
		Descriptor: "bowling",
	},
	"1F3CF": {
		Key:        "1F3CF",
		Value:      "🏏",
		Descriptor: "cricket game",
	},
	"1F3D1": {
		Key:        "1F3D1",
		Value:      "🏑",
		Descriptor: "field hockey",
	},
	"1F3D2": {
		Key:        "1F3D2",
		Value:      "🏒",
		Descriptor: "ice hockey",
	},
	"1F94D": {
		Key:        "1F94D",
		Value:      "🥍",
		Descriptor: "lacrosse",
	},
	"1F3D3": {
		Key:        "1F3D3",
		Value:      "🏓",
		Descriptor: "ping pong",
	},
	"1F3F8": {
		Key:        "1F3F8",
		Value:      "🏸",
		Descriptor: "badminton",
	},
	"1F94A": {
		Key:        "1F94A",
		Value:      "🥊",
		Descriptor: "boxing glove",
	},
	"1F94B": {
		Key:        "1F94B",
		Value:      "🥋",
		Descriptor: "martial arts uniform",
	},
	"1F945": {
		Key:        "1F945",
		Value:      "🥅",
		Descriptor: "goal net",
	},
	"26F3": {
		Key:        "26F3",
		Value:      "⛳",
		Descriptor: "flag in hole",
	},
	"26F8-FE0F": {
		Key:        "26F8-FE0F",
		Value:      "⛸️",
		Descriptor: "ice skate",
	},
	"1F3A3": {
		Key:        "1F3A3",
		Value:      "🎣",
		Descriptor: "fishing pole",
	},
	"1F93F": {
		Key:        "1F93F",
		Value:      "🤿",
		Descriptor: "diving mask",
	},
	"1F3BD": {
		Key:        "1F3BD",
		Value:      "🎽",
		Descriptor: "running shirt",
	},
	"1F3BF": {
		Key:        "1F3BF",
		Value:      "🎿",
		Descriptor: "skis",
	},
	"1F6F7": {
		Key:        "1F6F7",
		Value:      "🛷",
		Descriptor: "sled",
	},
	"1F94C": {
		Key:        "1F94C",
		Value:      "🥌",
		Descriptor: "curling stone",
	},
	"1F3AF": {
		Key:        "1F3AF",
		Value:      "🎯",
		Descriptor: "bullseye",
	},
	"1FA80": {
		Key:        "1FA80",
		Value:      "🪀",
		Descriptor: "yo-yo",
	},
	"1FA81": {
		Key:        "1FA81",
		Value:      "🪁",
		Descriptor: "kite",
	},
	"1F52B": {
		Key:        "1F52B",
		Value:      "🔫",
		Descriptor: "water pistol",
	},
	"1F3B1": {
		Key:        "1F3B1",
		Value:      "🎱",
		Descriptor: "pool 8 ball",
	},
	"1F52E": {
		Key:        "1F52E",
		Value:      "🔮",
		Descriptor: "crystal ball",
	},
	"1FA84": {
		Key:        "1FA84",
		Value:      "🪄",
		Descriptor: "magic wand",
	},
	"1F3AE": {
		Key:        "1F3AE",
		Value:      "🎮",
		Descriptor: "video game",
	},
	"1F579-FE0F": {
		Key:        "1F579-FE0F",
		Value:      "🕹️",
		Descriptor: "joystick",
	},
	"1F3B0": {
		Key:        "1F3B0",
		Value:      "🎰",
		Descriptor: "slot machine",
	},
	"1F3B2": {
		Key:        "1F3B2",
		Value:      "🎲",
		Descriptor: "game die",
	},
	"1F9E9": {
		Key:        "1F9E9",
		Value:      "🧩",
		Descriptor: "puzzle piece",
	},
	"1F9F8": {
		Key:        "1F9F8",
		Value:      "🧸",
		Descriptor: "teddy bear",
	},
	"1FA85": {
		Key:        "1FA85",
		Value:      "🪅",
		Descriptor: "piñata",
	},
	"1FAA9": {
		Key:        "1FAA9",
		Value:      "🪩",
		Descriptor: "mirror ball",
	},
	"1FA86": {
		Key:        "1FA86",
		Value:      "🪆",
		Descriptor: "nesting dolls",
	},
	"2660-FE0F": {
		Key:        "2660-FE0F",
		Value:      "♠️",
		Descriptor: "spade suit",
	},
	"2665-FE0F": {
		Key:        "2665-FE0F",
		Value:      "♥️",
		Descriptor: "heart suit",
	},
	"2666-FE0F": {
		Key:        "2666-FE0F",
		Value:      "♦️",
		Descriptor: "diamond suit",
	},
	"2663-FE0F": {
		Key:        "2663-FE0F",
		Value:      "♣️",
		Descriptor: "club suit",
	},
	"265F-FE0F": {
		Key:        "265F-FE0F",
		Value:      "♟️",
		Descriptor: "chess pawn",
	},
	"1F0CF": {
		Key:        "1F0CF",
		Value:      "🃏",
		Descriptor: "joker",
	},
	"1F004": {
		Key:        "1F004",
		Value:      "🀄",
		Descriptor: "mahjong red dragon",
	},
	"1F3B4": {
		Key:        "1F3B4",
		Value:      "🎴",
		Descriptor: "flower playing cards",
	},
	"1F3AD": {
		Key:        "1F3AD",
		Value:      "🎭",
		Descriptor: "performing arts",
	},
	"1F5BC-FE0F": {
		Key:        "1F5BC-FE0F",
		Value:      "🖼️",
		Descriptor: "framed picture",
	},
	"1F3A8": {
		Key:        "1F3A8",
		Value:      "🎨",
		Descriptor: "artist palette",
	},
	"1F9F5": {
		Key:        "1F9F5",
		Value:      "🧵",
		Descriptor: "thread",
	},
	"1FAA1": {
		Key:        "1FAA1",
		Value:      "🪡",
		Descriptor: "sewing needle",
	},
	"1F9F6": {
		Key:        "1F9F6",
		Value:      "🧶",
		Descriptor: "yarn",
	},
	"1FAA2": {
		Key:        "1FAA2",
		Value:      "🪢",
		Descriptor: "knot",
	},
	"1F453": {
		Key:        "1F453",
		Value:      "👓",
		Descriptor: "glasses",
	},
	"1F576-FE0F": {
		Key:        "1F576-FE0F",
		Value:      "🕶️",
		Descriptor: "sunglasses",
	},
	"1F97D": {
		Key:        "1F97D",
		Value:      "🥽",
		Descriptor: "goggles",
	},
	"1F97C": {
		Key:        "1F97C",
		Value:      "🥼",
		Descriptor: "lab coat",
	},
	"1F9BA": {
		Key:        "1F9BA",
		Value:      "🦺",
		Descriptor: "safety vest",
	},
	"1F454": {
		Key:        "1F454",
		Value:      "👔",
		Descriptor: "necktie",
	},
	"1F455": {
		Key:        "1F455",
		Value:      "👕",
		Descriptor: "t-shirt",
	},
	"1F456": {
		Key:        "1F456",
		Value:      "👖",
		Descriptor: "jeans",
	},
	"1F9E3": {
		Key:        "1F9E3",
		Value:      "🧣",
		Descriptor: "scarf",
	},
	"1F9E4": {
		Key:        "1F9E4",
		Value:      "🧤",
		Descriptor: "gloves",
	},
	"1F9E5": {
		Key:        "1F9E5",
		Value:      "🧥",
		Descriptor: "coat",
	},
	"1F9E6": {
		Key:        "1F9E6",
		Value:      "🧦",
		Descriptor: "socks",
	},
	"1F457": {
		Key:        "1F457",
		Value:      "👗",
		Descriptor: "dress",
	},
	"1F458": {
		Key:        "1F458",
		Value:      "👘",
		Descriptor: "kimono",
	},
	"1F97B": {
		Key:        "1F97B",
		Value:      "🥻",
		Descriptor: "sari",
	},
	"1FA71": {
		Key:        "1FA71",
		Value:      "🩱",
		Descriptor: "one-piece swimsuit",
	},
	"1FA72": {
		Key:        "1FA72",
		Value:      "🩲",
		Descriptor: "briefs",
	},
	"1FA73": {
		Key:        "1FA73",
		Value:      "🩳",
		Descriptor: "shorts",
	},
	"1F459": {
		Key:        "1F459",
		Value:      "👙",
		Descriptor: "bikini",
	},
	"1F45A": {
		Key:        "1F45A",
		Value:      "👚",
		Descriptor: "woman’s clothes",
	},
	"1FAAD": {
		Key:        "1FAAD",
		Value:      "🪭",
		Descriptor: "folding hand fan",
	},
	"1F45B": {
		Key:        "1F45B",
		Value:      "👛",
		Descriptor: "purse",
	},
	"1F45C": {
		Key:        "1F45C",
		Value:      "👜",
		Descriptor: "handbag",
	},
	"1F45D": {
		Key:        "1F45D",
		Value:      "👝",
		Descriptor: "clutch bag",
	},
	"1F6CD-FE0F": {
		Key:        "1F6CD-FE0F",
		Value:      "🛍️",
		Descriptor: "shopping bags",
	},
	"1F392": {
		Key:        "1F392",
		Value:      "🎒",
		Descriptor: "backpack",
	},
	"1FA74": {
		Key:        "1FA74",
		Value:      "🩴",
		Descriptor: "thong sandal",
	},
	"1F45E": {
		Key:        "1F45E",
		Value:      "👞",
		Descriptor: "man’s shoe",
	},
	"1F45F": {
		Key:        "1F45F",
		Value:      "👟",
		Descriptor: "running shoe",
	},
	"1F97E": {
		Key:        "1F97E",
		Value:      "🥾",
		Descriptor: "hiking boot",
	},
	"1F97F": {
		Key:        "1F97F",
		Value:      "🥿",
		Descriptor: "flat shoe",
	},
	"1F460": {
		Key:        "1F460",
		Value:      "👠",
		Descriptor: "high-heeled shoe",
	},
	"1F461": {
		Key:        "1F461",
		Value:      "👡",
		Descriptor: "woman’s sandal",
	},
	"1FA70": {
		Key:        "1FA70",
		Value:      "🩰",
		Descriptor: "ballet shoes",
	},
	"1F462": {
		Key:        "1F462",
		Value:      "👢",
		Descriptor: "woman’s boot",
	},
	"1FAAE": {
		Key:        "1FAAE",
		Value:      "🪮",
		Descriptor: "hair pick",
	},
	"1F451": {
		Key:        "1F451",
		Value:      "👑",
		Descriptor: "crown",
	},
	"1F452": {
		Key:        "1F452",
		Value:      "👒",
		Descriptor: "woman’s hat",
	},
	"1F3A9": {
		Key:        "1F3A9",
		Value:      "🎩",
		Descriptor: "top hat",
	},
	"1F393": {
		Key:        "1F393",
		Value:      "🎓",
		Descriptor: "graduation cap",
	},
	"1F9E2": {
		Key:        "1F9E2",
		Value:      "🧢",
		Descriptor: "billed cap",
	},
	"1FA96": {
		Key:        "1FA96",
		Value:      "🪖",
		Descriptor: "military helmet",
	},
	"26D1-FE0F": {
		Key:        "26D1-FE0F",
		Value:      "⛑️",
		Descriptor: "rescue worker’s helmet",
	},
	"1F4FF": {
		Key:        "1F4FF",
		Value:      "📿",
		Descriptor: "prayer beads",
	},
	"1F484": {
		Key:        "1F484",
		Value:      "💄",
		Descriptor: "lipstick",
	},
	"1F48D": {
		Key:        "1F48D",
		Value:      "💍",
		Descriptor: "ring",
	},
	"1F48E": {
		Key:        "1F48E",
		Value:      "💎",
		Descriptor: "gem stone",
	},
	"1F507": {
		Key:        "1F507",
		Value:      "🔇",
		Descriptor: "muted speaker",
	},
	"1F508": {
		Key:        "1F508",
		Value:      "🔈",
		Descriptor: "speaker low volume",
	},
	"1F509": {
		Key:        "1F509",
		Value:      "🔉",
		Descriptor: "speaker medium volume",
	},
	"1F50A": {
		Key:        "1F50A",
		Value:      "🔊",
		Descriptor: "speaker high volume",
	},
	"1F4E2": {
		Key:        "1F4E2",
		Value:      "📢",
		Descriptor: "loudspeaker",
	},
	"1F4E3": {
		Key:        "1F4E3",
		Value:      "📣",
		Descriptor: "megaphone",
	},
	"1F4EF": {
		Key:        "1F4EF",
		Value:      "📯",
		Descriptor: "postal horn",
	},
	"1F514": {
		Key:        "1F514",
		Value:      "🔔",
		Descriptor: "bell",
	},
	"1F515": {
		Key:        "1F515",
		Value:      "🔕",
		Descriptor: "bell with slash",
	},
	"1F3BC": {
		Key:        "1F3BC",
		Value:      "🎼",
		Descriptor: "musical score",
	},
	"1F3B5": {
		Key:        "1F3B5",
		Value:      "🎵",
		Descriptor: "musical note",
	},
	"1F3B6": {
		Key:        "1F3B6",
		Value:      "🎶",
		Descriptor: "musical notes",
	},
	"1F399-FE0F": {
		Key:        "1F399-FE0F",
		Value:      "🎙️",
		Descriptor: "studio microphone",
	},
	"1F39A-FE0F": {
		Key:        "1F39A-FE0F",
		Value:      "🎚️",
		Descriptor: "level slider",
	},
	"1F39B-FE0F": {
		Key:        "1F39B-FE0F",
		Value:      "🎛️",
		Descriptor: "control knobs",
	},
	"1F3A4": {
		Key:        "1F3A4",
		Value:      "🎤",
		Descriptor: "microphone",
	},
	"1F3A7": {
		Key:        "1F3A7",
		Value:      "🎧",
		Descriptor: "headphone",
	},
	"1F4FB": {
		Key:        "1F4FB",
		Value:      "📻",
		Descriptor: "radio",
	},
	"1F3B7": {
		Key:        "1F3B7",
		Value:      "🎷",
		Descriptor: "saxophone",
	},
	"1FA97": {
		Key:        "1FA97",
		Value:      "🪗",
		Descriptor: "accordion",
	},
	"1F3B8": {
		Key:        "1F3B8",
		Value:      "🎸",
		Descriptor: "guitar",
	},
	"1F3B9": {
		Key:        "1F3B9",
		Value:      "🎹",
		Descriptor: "musical keyboard",
	},
	"1F3BA": {
		Key:        "1F3BA",
		Value:      "🎺",
		Descriptor: "trumpet",
	},
	"1F3BB": {
		Key:        "1F3BB",
		Value:      "🎻",
		Descriptor: "violin",
	},
	"1FA95": {
		Key:        "1FA95",
		Value:      "🪕",
		Descriptor: "banjo",
	},
	"1F941": {
		Key:        "1F941",
		Value:      "🥁",
		Descriptor: "drum",
	},
	"1FA98": {
		Key:        "1FA98",
		Value:      "🪘",
		Descriptor: "long drum",
	},
	"1FA87": {
		Key:        "1FA87",
		Value:      "🪇",
		Descriptor: "maracas",
	},
	"1FA88": {
		Key:        "1FA88",
		Value:      "🪈",
		Descriptor: "flute",
	},
	"1F4F1": {
		Key:        "1F4F1",
		Value:      "📱",
		Descriptor: "mobile phone",
	},
	"1F4F2": {
		Key:        "1F4F2",
		Value:      "📲",
		Descriptor: "mobile phone with arrow",
	},
	"260E-FE0F": {
		Key:        "260E-FE0F",
		Value:      "☎️",
		Descriptor: "telephone",
	},
	"1F4DE": {
		Key:        "1F4DE",
		Value:      "📞",
		Descriptor: "telephone receiver",
	},
	"1F4DF": {
		Key:        "1F4DF",
		Value:      "📟",
		Descriptor: "pager",
	},
	"1F4E0": {
		Key:        "1F4E0",
		Value:      "📠",
		Descriptor: "fax machine",
	},
	"1F50B": {
		Key:        "1F50B",
		Value:      "🔋",
		Descriptor: "battery",
	},
	"1FAAB": {
		Key:        "1FAAB",
		Value:      "🪫",
		Descriptor: "low battery",
	},
	"1F50C": {
		Key:        "1F50C",
		Value:      "🔌",
		Descriptor: "electric plug",
	},
	"1F4BB": {
		Key:        "1F4BB",
		Value:      "💻",
		Descriptor: "laptop",
	},
	"1F5A5-FE0F": {
		Key:        "1F5A5-FE0F",
		Value:      "🖥️",
		Descriptor: "desktop computer",
	},
	"1F5A8-FE0F": {
		Key:        "1F5A8-FE0F",
		Value:      "🖨️",
		Descriptor: "printer",
	},
	"2328-FE0F": {
		Key:        "2328-FE0F",
		Value:      "⌨️",
		Descriptor: "keyboard",
	},
	"1F5B1-FE0F": {
		Key:        "1F5B1-FE0F",
		Value:      "🖱️",
		Descriptor: "computer mouse",
	},
	"1F5B2-FE0F": {
		Key:        "1F5B2-FE0F",
		Value:      "🖲️",
		Descriptor: "trackball",
	},
	"1F4BD": {
		Key:        "1F4BD",
		Value:      "💽",
		Descriptor: "computer disk",
	},
	"1F4BE": {
		Key:        "1F4BE",
		Value:      "💾",
		Descriptor: "floppy disk",
	},
	"1F4BF": {
		Key:        "1F4BF",
		Value:      "💿",
		Descriptor: "optical disk",
	},
	"1F4C0": {
		Key:        "1F4C0",
		Value:      "📀",
		Descriptor: "dvd",
	},
	"1F9EE": {
		Key:        "1F9EE",
		Value:      "🧮",
		Descriptor: "abacus",
	},
	"1F3A5": {
		Key:        "1F3A5",
		Value:      "🎥",
		Descriptor: "movie camera",
	},
	"1F39E-FE0F": {
		Key:        "1F39E-FE0F",
		Value:      "🎞️",
		Descriptor: "film frames",
	},
	"1F4FD-FE0F": {
		Key:        "1F4FD-FE0F",
		Value:      "📽️",
		Descriptor: "film projector",
	},
	"1F3AC": {
		Key:        "1F3AC",
		Value:      "🎬",
		Descriptor: "clapper board",
	},
	"1F4FA": {
		Key:        "1F4FA",
		Value:      "📺",
		Descriptor: "television",
	},
	"1F4F7": {
		Key:        "1F4F7",
		Value:      "📷",
		Descriptor: "camera",
	},
	"1F4F8": {
		Key:        "1F4F8",
		Value:      "📸",
		Descriptor: "camera with flash",
	},
	"1F4F9": {
		Key:        "1F4F9",
		Value:      "📹",
		Descriptor: "video camera",
	},
	"1F4FC": {
		Key:        "1F4FC",
		Value:      "📼",
		Descriptor: "videocassette",
	},
	"1F50D": {
		Key:        "1F50D",
		Value:      "🔍",
		Descriptor: "magnifying glass tilted left",
	},
	"1F50E": {
		Key:        "1F50E",
		Value:      "🔎",
		Descriptor: "magnifying glass tilted right",
	},
	"1F56F-FE0F": {
		Key:        "1F56F-FE0F",
		Value:      "🕯️",
		Descriptor: "candle",
	},
	"1F4A1": {
		Key:        "1F4A1",
		Value:      "💡",
		Descriptor: "light bulb",
	},
	"1F526": {
		Key:        "1F526",
		Value:      "🔦",
		Descriptor: "flashlight",
	},
	"1F3EE": {
		Key:        "1F3EE",
		Value:      "🏮",
		Descriptor: "red paper lantern",
	},
	"1FA94": {
		Key:        "1FA94",
		Value:      "🪔",
		Descriptor: "diya lamp",
	},
	"1F4D4": {
		Key:        "1F4D4",
		Value:      "📔",
		Descriptor: "notebook with decorative cover",
	},
	"1F4D5": {
		Key:        "1F4D5",
		Value:      "📕",
		Descriptor: "closed book",
	},
	"1F4D6": {
		Key:        "1F4D6",
		Value:      "📖",
		Descriptor: "open book",
	},
	"1F4D7": {
		Key:        "1F4D7",
		Value:      "📗",
		Descriptor: "green book",
	},
	"1F4D8": {
		Key:        "1F4D8",
		Value:      "📘",
		Descriptor: "blue book",
	},
	"1F4D9": {
		Key:        "1F4D9",
		Value:      "📙",
		Descriptor: "orange book",
	},
	"1F4DA": {
		Key:        "1F4DA",
		Value:      "📚",
		Descriptor: "books",
	},
	"1F4D3": {
		Key:        "1F4D3",
		Value:      "📓",
		Descriptor: "notebook",
	},
	"1F4D2": {
		Key:        "1F4D2",
		Value:      "📒",
		Descriptor: "ledger",
	},
	"1F4C3": {
		Key:        "1F4C3",
		Value:      "📃",
		Descriptor: "page with curl",
	},
	"1F4DC": {
		Key:        "1F4DC",
		Value:      "📜",
		Descriptor: "scroll",
	},
	"1F4C4": {
		Key:        "1F4C4",
		Value:      "📄",
		Descriptor: "page facing up",
	},
	"1F4F0": {
		Key:        "1F4F0",
		Value:      "📰",
		Descriptor: "newspaper",
	},
	"1F5DE-FE0F": {
		Key:        "1F5DE-FE0F",
		Value:      "🗞️",
		Descriptor: "rolled-up newspaper",
	},
	"1F4D1": {
		Key:        "1F4D1",
		Value:      "📑",
		Descriptor: "bookmark tabs",
	},
	"1F516": {
		Key:        "1F516",
		Value:      "🔖",
		Descriptor: "bookmark",
	},
	"1F3F7-FE0F": {
		Key:        "1F3F7-FE0F",
		Value:      "🏷️",
		Descriptor: "label",
	},
	"1F4B0": {
		Key:        "1F4B0",
		Value:      "💰",
		Descriptor: "money bag",
	},
	"1FA99": {
		Key:        "1FA99",
		Value:      "🪙",
		Descriptor: "coin",
	},
	"1F4B4": {
		Key:        "1F4B4",
		Value:      "💴",
		Descriptor: "yen banknote",
	},
	"1F4B5": {
		Key:        "1F4B5",
		Value:      "💵",
		Descriptor: "dollar banknote",
	},
	"1F4B6": {
		Key:        "1F4B6",
		Value:      "💶",
		Descriptor: "euro banknote",
	},
	"1F4B7": {
		Key:        "1F4B7",
		Value:      "💷",
		Descriptor: "pound banknote",
	},
	"1F4B8": {
		Key:        "1F4B8",
		Value:      "💸",
		Descriptor: "money with wings",
	},
	"1F4B3": {
		Key:        "1F4B3",
		Value:      "💳",
		Descriptor: "credit card",
	},
	"1F9FE": {
		Key:        "1F9FE",
		Value:      "🧾",
		Descriptor: "receipt",
	},
	"1F4B9": {
		Key:        "1F4B9",
		Value:      "💹",
		Descriptor: "chart increasing with yen",
	},
	"2709-FE0F": {
		Key:        "2709-FE0F",
		Value:      "✉️",
		Descriptor: "envelope",
	},
	"1F4E7": {
		Key:        "1F4E7",
		Value:      "📧",
		Descriptor: "e-mail",
	},
	"1F4E8": {
		Key:        "1F4E8",
		Value:      "📨",
		Descriptor: "incoming envelope",
	},
	"1F4E9": {
		Key:        "1F4E9",
		Value:      "📩",
		Descriptor: "envelope with arrow",
	},
	"1F4E4": {
		Key:        "1F4E4",
		Value:      "📤",
		Descriptor: "outbox tray",
	},
	"1F4E5": {
		Key:        "1F4E5",
		Value:      "📥",
		Descriptor: "inbox tray",
	},
	"1F4E6": {
		Key:        "1F4E6",
		Value:      "📦",
		Descriptor: "package",
	},
	"1F4EB": {
		Key:        "1F4EB",
		Value:      "📫",
		Descriptor: "closed mailbox with raised flag",
	},
	"1F4EA": {
		Key:        "1F4EA",
		Value:      "📪",
		Descriptor: "closed mailbox with lowered flag",
	},
	"1F4EC": {
		Key:        "1F4EC",
		Value:      "📬",
		Descriptor: "open mailbox with raised flag",
	},
	"1F4ED": {
		Key:        "1F4ED",
		Value:      "📭",
		Descriptor: "open mailbox with lowered flag",
	},
	"1F4EE": {
		Key:        "1F4EE",
		Value:      "📮",
		Descriptor: "postbox",
	},
	"1F5F3-FE0F": {
		Key:        "1F5F3-FE0F",
		Value:      "🗳️",
		Descriptor: "ballot box with ballot",
	},
	"270F-FE0F": {
		Key:        "270F-FE0F",
		Value:      "✏️",
		Descriptor: "pencil",
	},
	"2712-FE0F": {
		Key:        "2712-FE0F",
		Value:      "✒️",
		Descriptor: "black nib",
	},
	"1F58B-FE0F": {
		Key:        "1F58B-FE0F",
		Value:      "🖋️",
		Descriptor: "fountain pen",
	},
	"1F58A-FE0F": {
		Key:        "1F58A-FE0F",
		Value:      "🖊️",
		Descriptor: "pen",
	},
	"1F58C-FE0F": {
		Key:        "1F58C-FE0F",
		Value:      "🖌️",
		Descriptor: "paintbrush",
	},
	"1F58D-FE0F": {
		Key:        "1F58D-FE0F",
		Value:      "🖍️",
		Descriptor: "crayon",
	},
	"1F4DD": {
		Key:        "1F4DD",
		Value:      "📝",
		Descriptor: "memo",
	},
	"1F4BC": {
		Key:        "1F4BC",
		Value:      "💼",
		Descriptor: "briefcase",
	},
	"1F4C1": {
		Key:        "1F4C1",
		Value:      "📁",
		Descriptor: "file folder",
	},
	"1F4C2": {
		Key:        "1F4C2",
		Value:      "📂",
		Descriptor: "open file folder",
	},
	"1F5C2-FE0F": {
		Key:        "1F5C2-FE0F",
		Value:      "🗂️",
		Descriptor: "card index dividers",
	},
	"1F4C5": {
		Key:        "1F4C5",
		Value:      "📅",
		Descriptor: "calendar",
	},
	"1F4C6": {
		Key:        "1F4C6",
		Value:      "📆",
		Descriptor: "tear-off calendar",
	},
	"1F5D2-FE0F": {
		Key:        "1F5D2-FE0F",
		Value:      "🗒️",
		Descriptor: "spiral notepad",
	},
	"1F5D3-FE0F": {
		Key:        "1F5D3-FE0F",
		Value:      "🗓️",
		Descriptor: "spiral calendar",
	},
	"1F4C7": {
		Key:        "1F4C7",
		Value:      "📇",
		Descriptor: "card index",
	},
	"1F4C8": {
		Key:        "1F4C8",
		Value:      "📈",
		Descriptor: "chart increasing",
	},
	"1F4C9": {
		Key:        "1F4C9",
		Value:      "📉",
		Descriptor: "chart decreasing",
	},
	"1F4CA": {
		Key:        "1F4CA",
		Value:      "📊",
		Descriptor: "bar chart",
	},
	"1F4CB": {
		Key:        "1F4CB",
		Value:      "📋",
		Descriptor: "clipboard",
	},
	"1F4CC": {
		Key:        "1F4CC",
		Value:      "📌",
		Descriptor: "pushpin",
	},
	"1F4CD": {
		Key:        "1F4CD",
		Value:      "📍",
		Descriptor: "round pushpin",
	},
	"1F4CE": {
		Key:        "1F4CE",
		Value:      "📎",
		Descriptor: "paperclip",
	},
	"1F587-FE0F": {
		Key:        "1F587-FE0F",
		Value:      "🖇️",
		Descriptor: "linked paperclips",
	},
	"1F4CF": {
		Key:        "1F4CF",
		Value:      "📏",
		Descriptor: "straight ruler",
	},
	"1F4D0": {
		Key:        "1F4D0",
		Value:      "📐",
		Descriptor: "triangular ruler",
	},
	"2702-FE0F": {
		Key:        "2702-FE0F",
		Value:      "✂️",
		Descriptor: "scissors",
	},
	"1F5C3-FE0F": {
		Key:        "1F5C3-FE0F",
		Value:      "🗃️",
		Descriptor: "card file box",
	},
	"1F5C4-FE0F": {
		Key:        "1F5C4-FE0F",
		Value:      "🗄️",
		Descriptor: "file cabinet",
	},
	"1F5D1-FE0F": {
		Key:        "1F5D1-FE0F",
		Value:      "🗑️",
		Descriptor: "wastebasket",
	},
	"1F512": {
		Key:        "1F512",
		Value:      "🔒",
		Descriptor: "locked",
	},
	"1F513": {
		Key:        "1F513",
		Value:      "🔓",
		Descriptor: "unlocked",
	},
	"1F50F": {
		Key:        "1F50F",
		Value:      "🔏",
		Descriptor: "locked with pen",
	},
	"1F510": {
		Key:        "1F510",
		Value:      "🔐",
		Descriptor: "locked with key",
	},
	"1F511": {
		Key:        "1F511",
		Value:      "🔑",
		Descriptor: "key",
	},
	"1F5DD-FE0F": {
		Key:        "1F5DD-FE0F",
		Value:      "🗝️",
		Descriptor: "old key",
	},
	"1F528": {
		Key:        "1F528",
		Value:      "🔨",
		Descriptor: "hammer",
	},
	"1FA93": {
		Key:        "1FA93",
		Value:      "🪓",
		Descriptor: "axe",
	},
	"26CF-FE0F": {
		Key:        "26CF-FE0F",
		Value:      "⛏️",
		Descriptor: "pick",
	},
	"2692-FE0F": {
		Key:        "2692-FE0F",
		Value:      "⚒️",
		Descriptor: "hammer and pick",
	},
	"1F6E0-FE0F": {
		Key:        "1F6E0-FE0F",
		Value:      "🛠️",
		Descriptor: "hammer and wrench",
	},
	"1F5E1-FE0F": {
		Key:        "1F5E1-FE0F",
		Value:      "🗡️",
		Descriptor: "dagger",
	},
	"2694-FE0F": {
		Key:        "2694-FE0F",
		Value:      "⚔️",
		Descriptor: "crossed swords",
	},
	"1F4A3": {
		Key:        "1F4A3",
		Value:      "💣",
		Descriptor: "bomb",
	},
	"1FA83": {
		Key:        "1FA83",
		Value:      "🪃",
		Descriptor: "boomerang",
	},
	"1F3F9": {
		Key:        "1F3F9",
		Value:      "🏹",
		Descriptor: "bow and arrow",
	},
	"1F6E1-FE0F": {
		Key:        "1F6E1-FE0F",
		Value:      "🛡️",
		Descriptor: "shield",
	},
	"1FA9A": {
		Key:        "1FA9A",
		Value:      "🪚",
		Descriptor: "carpentry saw",
	},
	"1F527": {
		Key:        "1F527",
		Value:      "🔧",
		Descriptor: "wrench",
	},
	"1FA9B": {
		Key:        "1FA9B",
		Value:      "🪛",
		Descriptor: "screwdriver",
	},
	"1F529": {
		Key:        "1F529",
		Value:      "🔩",
		Descriptor: "nut and bolt",
	},
	"2699-FE0F": {
		Key:        "2699-FE0F",
		Value:      "⚙️",
		Descriptor: "gear",
	},
	"1F5DC-FE0F": {
		Key:        "1F5DC-FE0F",
		Value:      "🗜️",
		Descriptor: "clamp",
	},
	"2696-FE0F": {
		Key:        "2696-FE0F",
		Value:      "⚖️",
		Descriptor: "balance scale",
	},
	"1F9AF": {
		Key:        "1F9AF",
		Value:      "🦯",
		Descriptor: "white cane",
	},
	"1F517": {
		Key:        "1F517",
		Value:      "🔗",
		Descriptor: "link",
	},
	"26D3-FE0F-200D-1F4A5": {
		Key:        "26D3-FE0F-200D-1F4A5",
		Value:      "⛓️‍💥",
		Descriptor: "broken chain",
	},
	"26D3-FE0F": {
		Key:        "26D3-FE0F",
		Value:      "⛓️",
		Descriptor: "chains",
	},
	"1FA9D": {
		Key:        "1FA9D",
		Value:      "🪝",
		Descriptor: "hook",
	},
	"1F9F0": {
		Key:        "1F9F0",
		Value:      "🧰",
		Descriptor: "toolbox",
	},
	"1F9F2": {
		Key:        "1F9F2",
		Value:      "🧲",
		Descriptor: "magnet",
	},
	"1FA9C": {
		Key:        "1FA9C",
		Value:      "🪜",
		Descriptor: "ladder",
	},
	"2697-FE0F": {
		Key:        "2697-FE0F",
		Value:      "⚗️",
		Descriptor: "alembic",
	},
	"1F9EA": {
		Key:        "1F9EA",
		Value:      "🧪",
		Descriptor: "test tube",
	},
	"1F9EB": {
		Key:        "1F9EB",
		Value:      "🧫",
		Descriptor: "petri dish",
	},
	"1F9EC": {
		Key:        "1F9EC",
		Value:      "🧬",
		Descriptor: "dna",
	},
	"1F52C": {
		Key:        "1F52C",
		Value:      "🔬",
		Descriptor: "microscope",
	},
	"1F52D": {
		Key:        "1F52D",
		Value:      "🔭",
		Descriptor: "telescope",
	},
	"1F4E1": {
		Key:        "1F4E1",
		Value:      "📡",
		Descriptor: "satellite antenna",
	},
	"1F489": {
		Key:        "1F489",
		Value:      "💉",
		Descriptor: "syringe",
	},
	"1FA78": {
		Key:        "1FA78",
		Value:      "🩸",
		Descriptor: "drop of blood",
	},
	"1F48A": {
		Key:        "1F48A",
		Value:      "💊",
		Descriptor: "pill",
	},
	"1FA79": {
		Key:        "1FA79",
		Value:      "🩹",
		Descriptor: "adhesive bandage",
	},
	"1FA7C": {
		Key:        "1FA7C",
		Value:      "🩼",
		Descriptor: "crutch",
	},
	"1FA7A": {
		Key:        "1FA7A",
		Value:      "🩺",
		Descriptor: "stethoscope",
	},
	"1FA7B": {
		Key:        "1FA7B",
		Value:      "🩻",
		Descriptor: "x-ray",
	},
	"1F6AA": {
		Key:        "1F6AA",
		Value:      "🚪",
		Descriptor: "door",
	},
	"1F6D7": {
		Key:        "1F6D7",
		Value:      "🛗",
		Descriptor: "elevator",
	},
	"1FA9E": {
		Key:        "1FA9E",
		Value:      "🪞",
		Descriptor: "mirror",
	},
	"1FA9F": {
		Key:        "1FA9F",
		Value:      "🪟",
		Descriptor: "window",
	},
	"1F6CF-FE0F": {
		Key:        "1F6CF-FE0F",
		Value:      "🛏️",
		Descriptor: "bed",
	},
	"1F6CB-FE0F": {
		Key:        "1F6CB-FE0F",
		Value:      "🛋️",
		Descriptor: "couch and lamp",
	},
	"1FA91": {
		Key:        "1FA91",
		Value:      "🪑",
		Descriptor: "chair",
	},
	"1F6BD": {
		Key:        "1F6BD",
		Value:      "🚽",
		Descriptor: "toilet",
	},
	"1FAA0": {
		Key:        "1FAA0",
		Value:      "🪠",
		Descriptor: "plunger",
	},
	"1F6BF": {
		Key:        "1F6BF",
		Value:      "🚿",
		Descriptor: "shower",
	},
	"1F6C1": {
		Key:        "1F6C1",
		Value:      "🛁",
		Descriptor: "bathtub",
	},
	"1FAA4": {
		Key:        "1FAA4",
		Value:      "🪤",
		Descriptor: "mouse trap",
	},
	"1FA92": {
		Key:        "1FA92",
		Value:      "🪒",
		Descriptor: "razor",
	},
	"1F9F4": {
		Key:        "1F9F4",
		Value:      "🧴",
		Descriptor: "lotion bottle",
	},
	"1F9F7": {
		Key:        "1F9F7",
		Value:      "🧷",
		Descriptor: "safety pin",
	},
	"1F9F9": {
		Key:        "1F9F9",
		Value:      "🧹",
		Descriptor: "broom",
	},
	"1F9FA": {
		Key:        "1F9FA",
		Value:      "🧺",
		Descriptor: "basket",
	},
	"1F9FB": {
		Key:        "1F9FB",
		Value:      "🧻",
		Descriptor: "roll of paper",
	},
	"1FAA3": {
		Key:        "1FAA3",
		Value:      "🪣",
		Descriptor: "bucket",
	},
	"1F9FC": {
		Key:        "1F9FC",
		Value:      "🧼",
		Descriptor: "soap",
	},
	"1FAE7": {
		Key:        "1FAE7",
		Value:      "🫧",
		Descriptor: "bubbles",
	},
	"1FAA5": {
		Key:        "1FAA5",
		Value:      "🪥",
		Descriptor: "toothbrush",
	},
	"1F9FD": {
		Key:        "1F9FD",
		Value:      "🧽",
		Descriptor: "sponge",
	},
	"1F9EF": {
		Key:        "1F9EF",
		Value:      "🧯",
		Descriptor: "fire extinguisher",
	},
	"1F6D2": {
		Key:        "1F6D2",
		Value:      "🛒",
		Descriptor: "shopping cart",
	},
	"1F6AC": {
		Key:        "1F6AC",
		Value:      "🚬",
		Descriptor: "cigarette",
	},
	"26B0-FE0F": {
		Key:        "26B0-FE0F",
		Value:      "⚰️",
		Descriptor: "coffin",
	},
	"1FAA6": {
		Key:        "1FAA6",
		Value:      "🪦",
		Descriptor: "headstone",
	},
	"26B1-FE0F": {
		Key:        "26B1-FE0F",
		Value:      "⚱️",
		Descriptor: "funeral urn",
	},
	"1F9FF": {
		Key:        "1F9FF",
		Value:      "🧿",
		Descriptor: "nazar amulet",
	},
	"1FAAC": {
		Key:        "1FAAC",
		Value:      "🪬",
		Descriptor: "hamsa",
	},
	"1F5FF": {
		Key:        "1F5FF",
		Value:      "🗿",
		Descriptor: "moai",
	},
	"1FAA7": {
		Key:        "1FAA7",
		Value:      "🪧",
		Descriptor: "placard",
	},
	"1FAAA": {
		Key:        "1FAAA",
		Value:      "🪪",
		Descriptor: "identification card",
	},
	"1F3E7": {
		Key:        "1F3E7",
		Value:      "🏧",
		Descriptor: "ATM sign",
	},
	"1F6AE": {
		Key:        "1F6AE",
		Value:      "🚮",
		Descriptor: "litter in bin sign",
	},
	"1F6B0": {
		Key:        "1F6B0",
		Value:      "🚰",
		Descriptor: "potable water",
	},
	"267F": {
		Key:        "267F",
		Value:      "♿",
		Descriptor: "wheelchair symbol",
	},
	"1F6B9": {
		Key:        "1F6B9",
		Value:      "🚹",
		Descriptor: "men’s room",
	},
	"1F6BA": {
		Key:        "1F6BA",
		Value:      "🚺",
		Descriptor: "women’s room",
	},
	"1F6BB": {
		Key:        "1F6BB",
		Value:      "🚻",
		Descriptor: "restroom",
	},
	"1F6BC": {
		Key:        "1F6BC",
		Value:      "🚼",
		Descriptor: "baby symbol",
	},
	"1F6BE": {
		Key:        "1F6BE",
		Value:      "🚾",
		Descriptor: "water closet",
	},
	"1F6C2": {
		Key:        "1F6C2",
		Value:      "🛂",
		Descriptor: "passport control",
	},
	"1F6C3": {
		Key:        "1F6C3",
		Value:      "🛃",
		Descriptor: "customs",
	},
	"1F6C4": {
		Key:        "1F6C4",
		Value:      "🛄",
		Descriptor: "baggage claim",
	},
	"1F6C5": {
		Key:        "1F6C5",
		Value:      "🛅",
		Descriptor: "left luggage",
	},
	"26A0-FE0F": {
		Key:        "26A0-FE0F",
		Value:      "⚠️",
		Descriptor: "warning",
	},
	"1F6B8": {
		Key:        "1F6B8",
		Value:      "🚸",
		Descriptor: "children crossing",
	},
	"26D4": {
		Key:        "26D4",
		Value:      "⛔",
		Descriptor: "no entry",
	},
	"1F6AB": {
		Key:        "1F6AB",
		Value:      "🚫",
		Descriptor: "prohibited",
	},
	"1F6B3": {
		Key:        "1F6B3",
		Value:      "🚳",
		Descriptor: "no bicycles",
	},
	"1F6AD": {
		Key:        "1F6AD",
		Value:      "🚭",
		Descriptor: "no smoking",
	},
	"1F6AF": {
		Key:        "1F6AF",
		Value:      "🚯",
		Descriptor: "no littering",
	},
	"1F6B1": {
		Key:        "1F6B1",
		Value:      "🚱",
		Descriptor: "non-potable water",
	},
	"1F6B7": {
		Key:        "1F6B7",
		Value:      "🚷",
		Descriptor: "no pedestrians",
	},
	"1F4F5": {
		Key:        "1F4F5",
		Value:      "📵",
		Descriptor: "no mobile phones",
	},
	"1F51E": {
		Key:        "1F51E",
		Value:      "🔞",
		Descriptor: "no one under eighteen",
	},
	"2622-FE0F": {
		Key:        "2622-FE0F",
		Value:      "☢️",
		Descriptor: "radioactive",
	},
	"2623-FE0F": {
		Key:        "2623-FE0F",
		Value:      "☣️",
		Descriptor: "biohazard",
	},
	"2B06-FE0F": {
		Key:        "2B06-FE0F",
		Value:      "⬆️",
		Descriptor: "up arrow",
	},
	"2197-FE0F": {
		Key:        "2197-FE0F",
		Value:      "↗️",
		Descriptor: "up-right arrow",
	},
	"27A1-FE0F": {
		Key:        "27A1-FE0F",
		Value:      "➡️",
		Descriptor: "right arrow",
	},
	"2198-FE0F": {
		Key:        "2198-FE0F",
		Value:      "↘️",
		Descriptor: "down-right arrow",
	},
	"2B07-FE0F": {
		Key:        "2B07-FE0F",
		Value:      "⬇️",
		Descriptor: "down arrow",
	},
	"2199-FE0F": {
		Key:        "2199-FE0F",
		Value:      "↙️",
		Descriptor: "down-left arrow",
	},
	"2B05-FE0F": {
		Key:        "2B05-FE0F",
		Value:      "⬅️",
		Descriptor: "left arrow",
	},
	"2196-FE0F": {
		Key:        "2196-FE0F",
		Value:      "↖️",
		Descriptor: "up-left arrow",
	},
	"2195-FE0F": {
		Key:        "2195-FE0F",
		Value:      "↕️",
		Descriptor: "up-down arrow",
	},
	"2194-FE0F": {
		Key:        "2194-FE0F",
		Value:      "↔️",
		Descriptor: "left-right arrow",
	},
	"21A9-FE0F": {
		Key:        "21A9-FE0F",
		Value:      "↩️",
		Descriptor: "right arrow curving left",
	},
	"21AA-FE0F": {
		Key:        "21AA-FE0F",
		Value:      "↪️",
		Descriptor: "left arrow curving right",
	},
	"2934-FE0F": {
		Key:        "2934-FE0F",
		Value:      "⤴️",
		Descriptor: "right arrow curving up",
	},
	"2935-FE0F": {
		Key:        "2935-FE0F",
		Value:      "⤵️",
		Descriptor: "right arrow curving down",
	},
	"1F503": {
		Key:        "1F503",
		Value:      "🔃",
		Descriptor: "clockwise vertical arrows",
	},
	"1F504": {
		Key:        "1F504",
		Value:      "🔄",
		Descriptor: "counterclockwise arrows button",
	},
	"1F519": {
		Key:        "1F519",
		Value:      "🔙",
		Descriptor: "BACK arrow",
	},
	"1F51A": {
		Key:        "1F51A",
		Value:      "🔚",
		Descriptor: "END arrow",
	},
	"1F51B": {
		Key:        "1F51B",
		Value:      "🔛",
		Descriptor: "ON! arrow",
	},
	"1F51C": {
		Key:        "1F51C",
		Value:      "🔜",
		Descriptor: "SOON arrow",
	},
	"1F51D": {
		Key:        "1F51D",
		Value:      "🔝",
		Descriptor: "TOP arrow",
	},
	"1F6D0": {
		Key:        "1F6D0",
		Value:      "🛐",
		Descriptor: "place of worship",
	},
	"269B-FE0F": {
		Key:        "269B-FE0F",
		Value:      "⚛️",
		Descriptor: "atom symbol",
	},
	"1F549-FE0F": {
		Key:        "1F549-FE0F",
		Value:      "🕉️",
		Descriptor: "om",
	},
	"2721-FE0F": {
		Key:        "2721-FE0F",
		Value:      "✡️",
		Descriptor: "star of David",
	},
	"2638-FE0F": {
		Key:        "2638-FE0F",
		Value:      "☸️",
		Descriptor: "wheel of dharma",
	},
	"262F-FE0F": {
		Key:        "262F-FE0F",
		Value:      "☯️",
		Descriptor: "yin yang",
	},
	"271D-FE0F": {
		Key:        "271D-FE0F",
		Value:      "✝️",
		Descriptor: "latin cross",
	},
	"2626-FE0F": {
		Key:        "2626-FE0F",
		Value:      "☦️",
		Descriptor: "orthodox cross",
	},
	"262A-FE0F": {
		Key:        "262A-FE0F",
		Value:      "☪️",
		Descriptor: "star and crescent",
	},
	"262E-FE0F": {
		Key:        "262E-FE0F",
		Value:      "☮️",
		Descriptor: "peace symbol",
	},
	"1F54E": {
		Key:        "1F54E",
		Value:      "🕎",
		Descriptor: "menorah",
	},
	"1F52F": {
		Key:        "1F52F",
		Value:      "🔯",
		Descriptor: "dotted six-pointed star",
	},
	"1FAAF": {
		Key:        "1FAAF",
		Value:      "🪯",
		Descriptor: "khanda",
	},
	"2648": {
		Key:        "2648",
		Value:      "♈",
		Descriptor: "Aries",
	},
	"2649": {
		Key:        "2649",
		Value:      "♉",
		Descriptor: "Taurus",
	},
	"264A": {
		Key:        "264A",
		Value:      "♊",
		Descriptor: "Gemini",
	},
	"264B": {
		Key:        "264B",
		Value:      "♋",
		Descriptor: "Cancer",
	},
	"264C": {
		Key:        "264C",
		Value:      "♌",
		Descriptor: "Leo",
	},
	"264D": {
		Key:        "264D",
		Value:      "♍",
		Descriptor: "Virgo",
	},
	"264E": {
		Key:        "264E",
		Value:      "♎",
		Descriptor: "Libra",
	},
	"264F": {
		Key:        "264F",
		Value:      "♏",
		Descriptor: "Scorpio",
	},
	"2650": {
		Key:        "2650",
		Value:      "♐",
		Descriptor: "Sagittarius",
	},
	"2651": {
		Key:        "2651",
		Value:      "♑",
		Descriptor: "Capricorn",
	},
	"2652": {
		Key:        "2652",
		Value:      "♒",
		Descriptor: "Aquarius",
	},
	"2653": {
		Key:        "2653",
		Value:      "♓",
		Descriptor: "Pisces",
	},
	"26CE": {
		Key:        "26CE",
		Value:      "⛎",
		Descriptor: "Ophiuchus",
	},
	"1F500": {
		Key:        "1F500",
		Value:      "🔀",
		Descriptor: "shuffle tracks button",
	},
	"1F501": {
		Key:        "1F501",
		Value:      "🔁",
		Descriptor: "repeat button",
	},
	"1F502": {
		Key:        "1F502",
		Value:      "🔂",
		Descriptor: "repeat single button",
	},
	"25B6-FE0F": {
		Key:        "25B6-FE0F",
		Value:      "▶️",
		Descriptor: "play button",
	},
	"23E9": {
		Key:        "23E9",
		Value:      "⏩",
		Descriptor: "fast-forward button",
	},
	"23ED-FE0F": {
		Key:        "23ED-FE0F",
		Value:      "⏭️",
		Descriptor: "next track button",
	},
	"23EF-FE0F": {
		Key:        "23EF-FE0F",
		Value:      "⏯️",
		Descriptor: "play or pause button",
	},
	"25C0-FE0F": {
		Key:        "25C0-FE0F",
		Value:      "◀️",
		Descriptor: "reverse button",
	},
	"23EA": {
		Key:        "23EA",
		Value:      "⏪",
		Descriptor: "fast reverse button",
	},
	"23EE-FE0F": {
		Key:        "23EE-FE0F",
		Value:      "⏮️",
		Descriptor: "last track button",
	},
	"1F53C": {
		Key:        "1F53C",
		Value:      "🔼",
		Descriptor: "upwards button",
	},
	"23EB": {
		Key:        "23EB",
		Value:      "⏫",
		Descriptor: "fast up button",
	},
	"1F53D": {
		Key:        "1F53D",
		Value:      "🔽",
		Descriptor: "downwards button",
	},
	"23EC": {
		Key:        "23EC",
		Value:      "⏬",
		Descriptor: "fast down button",
	},
	"23F8-FE0F": {
		Key:        "23F8-FE0F",
		Value:      "⏸️",
		Descriptor: "pause button",
	},
	"23F9-FE0F": {
		Key:        "23F9-FE0F",
		Value:      "⏹️",
		Descriptor: "stop button",
	},
	"23FA-FE0F": {
		Key:        "23FA-FE0F",
		Value:      "⏺️",
		Descriptor: "record button",
	},
	"23CF-FE0F": {
		Key:        "23CF-FE0F",
		Value:      "⏏️",
		Descriptor: "eject button",
	},
	"1F3A6": {
		Key:        "1F3A6",
		Value:      "🎦",
		Descriptor: "cinema",
	},
	"1F505": {
		Key:        "1F505",
		Value:      "🔅",
		Descriptor: "dim button",
	},
	"1F506": {
		Key:        "1F506",
		Value:      "🔆",
		Descriptor: "bright button",
	},
	"1F4F6": {
		Key:        "1F4F6",
		Value:      "📶",
		Descriptor: "antenna bars",
	},
	"1F6DC": {
		Key:        "1F6DC",
		Value:      "🛜",
		Descriptor: "wireless",
	},
	"1F4F3": {
		Key:        "1F4F3",
		Value:      "📳",
		Descriptor: "vibration mode",
	},
	"1F4F4": {
		Key:        "1F4F4",
		Value:      "📴",
		Descriptor: "mobile phone off",
	},
	"2640-FE0F": {
		Key:        "2640-FE0F",
		Value:      "♀️",
		Descriptor: "female sign",
	},
	"2642-FE0F": {
		Key:        "2642-FE0F",
		Value:      "♂️",
		Descriptor: "male sign",
	},
	"26A7-FE0F": {
		Key:        "26A7-FE0F",
		Value:      "⚧️",
		Descriptor: "transgender symbol",
	},
	"2716-FE0F": {
		Key:        "2716-FE0F",
		Value:      "✖️",
		Descriptor: "multiply",
	},
	"2795": {
		Key:        "2795",
		Value:      "➕",
		Descriptor: "plus",
	},
	"2796": {
		Key:        "2796",
		Value:      "➖",
		Descriptor: "minus",
	},
	"2797": {
		Key:        "2797",
		Value:      "➗",
		Descriptor: "divide",
	},
	"1F7F0": {
		Key:        "1F7F0",
		Value:      "🟰",
		Descriptor: "heavy equals sign",
	},
	"267E-FE0F": {
		Key:        "267E-FE0F",
		Value:      "♾️",
		Descriptor: "infinity",
	},
	"203C-FE0F": {
		Key:        "203C-FE0F",
		Value:      "‼️",
		Descriptor: "double exclamation mark",
	},
	"2049-FE0F": {
		Key:        "2049-FE0F",
		Value:      "⁉️",
		Descriptor: "exclamation question mark",
	},
	"2753": {
		Key:        "2753",
		Value:      "❓",
		Descriptor: "red question mark",
	},
	"2754": {
		Key:        "2754",
		Value:      "❔",
		Descriptor: "white question mark",
	},
	"2755": {
		Key:        "2755",
		Value:      "❕",
		Descriptor: "white exclamation mark",
	},
	"2757": {
		Key:        "2757",
		Value:      "❗",
		Descriptor: "red exclamation mark",
	},
	"3030-FE0F": {
		Key:        "3030-FE0F",
		Value:      "〰️",
		Descriptor: "wavy dash",
	},
	"1F4B1": {
		Key:        "1F4B1",
		Value:      "💱",
		Descriptor: "currency exchange",
	},
	"1F4B2": {
		Key:        "1F4B2",
		Value:      "💲",
		Descriptor: "heavy dollar sign",
	},
	"2695-FE0F": {
		Key:        "2695-FE0F",
		Value:      "⚕️",
		Descriptor: "medical symbol",
	},
	"267B-FE0F": {
		Key:        "267B-FE0F",
		Value:      "♻️",
		Descriptor: "recycling symbol",
	},
	"269C-FE0F": {
		Key:        "269C-FE0F",
		Value:      "⚜️",
		Descriptor: "fleur-de-lis",
	},
	"1F531": {
		Key:        "1F531",
		Value:      "🔱",
		Descriptor: "trident emblem",
	},
	"1F4DB": {
		Key:        "1F4DB",
		Value:      "📛",
		Descriptor: "name badge",
	},
	"1F530": {
		Key:        "1F530",
		Value:      "🔰",
		Descriptor: "Japanese symbol for beginner",
	},
	"2B55": {
		Key:        "2B55",
		Value:      "⭕",
		Descriptor: "hollow red circle",
	},
	"2705": {
		Key:        "2705",
		Value:      "✅",
		Descriptor: "check mark button",
	},
	"2611-FE0F": {
		Key:        "2611-FE0F",
		Value:      "☑️",
		Descriptor: "check box with check",
	},
	"2714-FE0F": {
		Key:        "2714-FE0F",
		Value:      "✔️",
		Descriptor: "check mark",
	},
	"274C": {
		Key:        "274C",
		Value:      "❌",
		Descriptor: "cross mark",
	},
	"274E": {
		Key:        "274E",
		Value:      "❎",
		Descriptor: "cross mark button",
	},
	"27B0": {
		Key:        "27B0",
		Value:      "➰",
		Descriptor: "curly loop",
	},
	"27BF": {
		Key:        "27BF",
		Value:      "➿",
		Descriptor: "double curly loop",
	},
	"303D-FE0F": {
		Key:        "303D-FE0F",
		Value:      "〽️",
		Descriptor: "part alternation mark",
	},
	"2733-FE0F": {
		Key:        "2733-FE0F",
		Value:      "✳️",
		Descriptor: "eight-spoked asterisk",
	},
	"2734-FE0F": {
		Key:        "2734-FE0F",
		Value:      "✴️",
		Descriptor: "eight-pointed star",
	},
	"2747-FE0F": {
		Key:        "2747-FE0F",
		Value:      "❇️",
		Descriptor: "sparkle",
	},
	"00A9-FE0F": {
		Key:        "00A9-FE0F",
		Value:      "©️",
		Descriptor: "copyright",
	},
	"00AE-FE0F": {
		Key:        "00AE-FE0F",
		Value:      "®️",
		Descriptor: "registered",
	},
	"2122-FE0F": {
		Key:        "2122-FE0F",
		Value:      "™️",
		Descriptor: "trade mark",
	},
	"0023-FE0F-20E3": {
		Key:        "0023-FE0F-20E3",
		Value:      "#️⃣",
		Descriptor: "keycap: #",
	},
	"002A-FE0F-20E3": {
		Key:        "002A-FE0F-20E3",
		Value:      "*️⃣",
		Descriptor: "keycap: *",
	},
	"0030-FE0F-20E3": {
		Key:        "0030-FE0F-20E3",
		Value:      "0️⃣",
		Descriptor: "keycap: 0",
	},
	"0031-FE0F-20E3": {
		Key:        "0031-FE0F-20E3",
		Value:      "1️⃣",
		Descriptor: "keycap: 1",
	},
	"0032-FE0F-20E3": {
		Key:        "0032-FE0F-20E3",
		Value:      "2️⃣",
		Descriptor: "keycap: 2",
	},
	"0033-FE0F-20E3": {
		Key:        "0033-FE0F-20E3",
		Value:      "3️⃣",
		Descriptor: "keycap: 3",
	},
	"0034-FE0F-20E3": {
		Key:        "0034-FE0F-20E3",
		Value:      "4️⃣",
		Descriptor: "keycap: 4",
	},
	"0035-FE0F-20E3": {
		Key:        "0035-FE0F-20E3",
		Value:      "5️⃣",
		Descriptor: "keycap: 5",
	},
	"0036-FE0F-20E3": {
		Key:        "0036-FE0F-20E3",
		Value:      "6️⃣",
		Descriptor: "keycap: 6",
	},
	"0037-FE0F-20E3": {
		Key:        "0037-FE0F-20E3",
		Value:      "7️⃣",
		Descriptor: "keycap: 7",
	},
	"0038-FE0F-20E3": {
		Key:        "0038-FE0F-20E3",
		Value:      "8️⃣",
		Descriptor: "keycap: 8",
	},
	"0039-FE0F-20E3": {
		Key:        "0039-FE0F-20E3",
		Value:      "9️⃣",
		Descriptor: "keycap: 9",
	},
	"1F51F": {
		Key:        "1F51F",
		Value:      "🔟",
		Descriptor: "keycap: 10",
	},
	"1F520": {
		Key:        "1F520",
		Value:      "🔠",
		Descriptor: "input latin uppercase",
	},
	"1F521": {
		Key:        "1F521",
		Value:      "🔡",
		Descriptor: "input latin lowercase",
	},
	"1F522": {
		Key:        "1F522",
		Value:      "🔢",
		Descriptor: "input numbers",
	},
	"1F523": {
		Key:        "1F523",
		Value:      "🔣",
		Descriptor: "input symbols",
	},
	"1F524": {
		Key:        "1F524",
		Value:      "🔤",
		Descriptor: "input latin letters",
	},
	"1F170-FE0F": {
		Key:        "1F170-FE0F",
		Value:      "🅰️",
		Descriptor: "A button (blood type)",
	},
	"1F18E": {
		Key:        "1F18E",
		Value:      "🆎",
		Descriptor: "AB button (blood type)",
	},
	"1F171-FE0F": {
		Key:        "1F171-FE0F",
		Value:      "🅱️",
		Descriptor: "B button (blood type)",
	},
	"1F191": {
		Key:        "1F191",
		Value:      "🆑",
		Descriptor: "CL button",
	},
	"1F192": {
		Key:        "1F192",
		Value:      "🆒",
		Descriptor: "COOL button",
	},
	"1F193": {
		Key:        "1F193",
		Value:      "🆓",
		Descriptor: "FREE button",
	},
	"2139-FE0F": {
		Key:        "2139-FE0F",
		Value:      "ℹ️",
		Descriptor: "information",
	},
	"1F194": {
		Key:        "1F194",
		Value:      "🆔",
		Descriptor: "ID button",
	},
	"24C2-FE0F": {
		Key:        "24C2-FE0F",
		Value:      "Ⓜ️",
		Descriptor: "circled M",
	},
	"1F195": {
		Key:        "1F195",
		Value:      "🆕",
		Descriptor: "NEW button",
	},
	"1F196": {
		Key:        "1F196",
		Value:      "🆖",
		Descriptor: "NG button",
	},
	"1F17E-FE0F": {
		Key:        "1F17E-FE0F",
		Value:      "🅾️",
		Descriptor: "O button (blood type)",
	},
	"1F197": {
		Key:        "1F197",
		Value:      "🆗",
		Descriptor: "OK button",
	},
	"1F17F-FE0F": {
		Key:        "1F17F-FE0F",
		Value:      "🅿️",
		Descriptor: "P button",
	},
	"1F198": {
		Key:        "1F198",
		Value:      "🆘",
		Descriptor: "SOS button",
	},
	"1F199": {
		Key:        "1F199",
		Value:      "🆙",
		Descriptor: "UP! button",
	},
	"1F19A": {
		Key:        "1F19A",
		Value:      "🆚",
		Descriptor: "VS button",
	},
	"1F201": {
		Key:        "1F201",
		Value:      "🈁",
		Descriptor: "Japanese “here” button",
	},
	"1F202-FE0F": {
		Key:        "1F202-FE0F",
		Value:      "🈂️",
		Descriptor: "Japanese “service charge” button",
	},
	"1F237-FE0F": {
		Key:        "1F237-FE0F",
		Value:      "🈷️",
		Descriptor: "Japanese “monthly amount” button",
	},
	"1F236": {
		Key:        "1F236",
		Value:      "🈶",
		Descriptor: "Japanese “not free of charge” button",
	},
	"1F22F": {
		Key:        "1F22F",
		Value:      "🈯",
		Descriptor: "Japanese “reserved” button",
	},
	"1F250": {
		Key:        "1F250",
		Value:      "🉐",
		Descriptor: "Japanese “bargain” button",
	},
	"1F239": {
		Key:        "1F239",
		Value:      "🈹",
		Descriptor: "Japanese “discount” button",
	},
	"1F21A": {
		Key:        "1F21A",
		Value:      "🈚",
		Descriptor: "Japanese “free of charge” button",
	},
	"1F232": {
		Key:        "1F232",
		Value:      "🈲",
		Descriptor: "Japanese “prohibited” button",
	},
	"1F251": {
		Key:        "1F251",
		Value:      "🉑",
		Descriptor: "Japanese “acceptable” button",
	},
	"1F238": {
		Key:        "1F238",
		Value:      "🈸",
		Descriptor: "Japanese “application” button",
	},
	"1F234": {
		Key:        "1F234",
		Value:      "🈴",
		Descriptor: "Japanese “passing grade” button",
	},
	"1F233": {
		Key:        "1F233",
		Value:      "🈳",
		Descriptor: "Japanese “vacancy” button",
	},
	"3297-FE0F": {
		Key:        "3297-FE0F",
		Value:      "㊗️",
		Descriptor: "Japanese “congratulations” button",
	},
	"3299-FE0F": {
		Key:        "3299-FE0F",
		Value:      "㊙️",
		Descriptor: "Japanese “secret” button",
	},
	"1F23A": {
		Key:        "1F23A",
		Value:      "🈺",
		Descriptor: "Japanese “open for business” button",
	},
	"1F235": {
		Key:        "1F235",
		Value:      "🈵",
		Descriptor: "Japanese “no vacancy” button",
	},
	"1F534": {
		Key:        "1F534",
		Value:      "🔴",
		Descriptor: "red circle",
	},
	"1F7E0": {
		Key:        "1F7E0",
		Value:      "🟠",
		Descriptor: "orange circle",
	},
	"1F7E1": {
		Key:        "1F7E1",
		Value:      "🟡",
		Descriptor: "yellow circle",
	},
	"1F7E2": {
		Key:        "1F7E2",
		Value:      "🟢",
		Descriptor: "green circle",
	},
	"1F535": {
		Key:        "1F535",
		Value:      "🔵",
		Descriptor: "blue circle",
	},
	"1F7E3": {
		Key:        "1F7E3",
		Value:      "🟣",
		Descriptor: "purple circle",
	},
	"1F7E4": {
		Key:        "1F7E4",
		Value:      "🟤",
		Descriptor: "brown circle",
	},
	"26AB": {
		Key:        "26AB",
		Value:      "⚫",
		Descriptor: "black circle",
	},
	"26AA": {
		Key:        "26AA",
		Value:      "⚪",
		Descriptor: "white circle",
	},
	"1F7E5": {
		Key:        "1F7E5",
		Value:      "🟥",
		Descriptor: "red square",
	},
	"1F7E7": {
		Key:        "1F7E7",
		Value:      "🟧",
		Descriptor: "orange square",
	},
	"1F7E8": {
		Key:        "1F7E8",
		Value:      "🟨",
		Descriptor: "yellow square",
	},
	"1F7E9": {
		Key:        "1F7E9",
		Value:      "🟩",
		Descriptor: "green square",
	},
	"1F7E6": {
		Key:        "1F7E6",
		Value:      "🟦",
		Descriptor: "blue square",
	},
	"1F7EA": {
		Key:        "1F7EA",
		Value:      "🟪",
		Descriptor: "purple square",
	},
	"1F7EB": {
		Key:        "1F7EB",
		Value:      "🟫",
		Descriptor: "brown square",
	},
	"2B1B": {
		Key:        "2B1B",
		Value:      "⬛",
		Descriptor: "black large square",
	},
	"2B1C": {
		Key:        "2B1C",
		Value:      "⬜",
		Descriptor: "white large square",
	},
	"25FC-FE0F": {
		Key:        "25FC-FE0F",
		Value:      "◼️",
		Descriptor: "black medium square",
	},
	"25FB-FE0F": {
		Key:        "25FB-FE0F",
		Value:      "◻️",
		Descriptor: "white medium square",
	},
	"25FE": {
		Key:        "25FE",
		Value:      "◾",
		Descriptor: "black medium-small square",
	},
	"25FD": {
		Key:        "25FD",
		Value:      "◽",
		Descriptor: "white medium-small square",
	},
	"25AA-FE0F": {
		Key:        "25AA-FE0F",
		Value:      "▪️",
		Descriptor: "black small square",
	},
	"25AB-FE0F": {
		Key:        "25AB-FE0F",
		Value:      "▫️",
		Descriptor: "white small square",
	},
	"1F536": {
		Key:        "1F536",
		Value:      "🔶",
		Descriptor: "large orange diamond",
	},
	"1F537": {
		Key:        "1F537",
		Value:      "🔷",
		Descriptor: "large blue diamond",
	},
	"1F538": {
		Key:        "1F538",
		Value:      "🔸",
		Descriptor: "small orange diamond",
	},
	"1F539": {
		Key:        "1F539",
		Value:      "🔹",
		Descriptor: "small blue diamond",
	},
	"1F53A": {
		Key:        "1F53A",
		Value:      "🔺",
		Descriptor: "red triangle pointed up",
	},
	"1F53B": {
		Key:        "1F53B",
		Value:      "🔻",
		Descriptor: "red triangle pointed down",
	},
	"1F4A0": {
		Key:        "1F4A0",
		Value:      "💠",
		Descriptor: "diamond with a dot",
	},
	"1F518": {
		Key:        "1F518",
		Value:      "🔘",
		Descriptor: "radio button",
	},
	"1F533": {
		Key:        "1F533",
		Value:      "🔳",
		Descriptor: "white square button",
	},
	"1F532": {
		Key:        "1F532",
		Value:      "🔲",
		Descriptor: "black square button",
	},
	"1F3C1": {
		Key:        "1F3C1",
		Value:      "🏁",
		Descriptor: "chequered flag",
	},
	"1F6A9": {
		Key:        "1F6A9",
		Value:      "🚩",
		Descriptor: "triangular flag",
	},
	"1F38C": {
		Key:        "1F38C",
		Value:      "🎌",
		Descriptor: "crossed flags",
	},
	"1F3F4": {
		Key:        "1F3F4",
		Value:      "🏴",
		Descriptor: "black flag",
	},
	"1F3F3-FE0F": {
		Key:        "1F3F3-FE0F",
		Value:      "🏳️",
		Descriptor: "white flag",
	},
	"1F3F3-FE0F-200D-1F308": {
		Key:        "1F3F3-FE0F-200D-1F308",
		Value:      "🏳️‍🌈",
		Descriptor: "rainbow flag",
	},
	"1F3F3-FE0F-200D-26A7-FE0F": {
		Key:        "1F3F3-FE0F-200D-26A7-FE0F",
		Value:      "🏳️‍⚧️",
		Descriptor: "transgender flag",
	},
	"1F3F4-200D-2620-FE0F": {
		Key:        "1F3F4-200D-2620-FE0F",
		Value:      "🏴‍☠️",
		Descriptor: "pirate flag",
	},
	"1F1E6-1F1E8": {
		Key:        "1F1E6-1F1E8",
		Value:      "🇦🇨",
		Descriptor: "flag: Ascension Island",
	},
	"1F1E6-1F1E9": {
		Key:        "1F1E6-1F1E9",
		Value:      "🇦🇩",
		Descriptor: "flag: Andorra",
	},
	"1F1E6-1F1EA": {
		Key:        "1F1E6-1F1EA",
		Value:      "🇦🇪",
		Descriptor: "flag: United Arab Emirates",
	},
	"1F1E6-1F1EB": {
		Key:        "1F1E6-1F1EB",
		Value:      "🇦🇫",
		Descriptor: "flag: Afghanistan",
	},
	"1F1E6-1F1EC": {
		Key:        "1F1E6-1F1EC",
		Value:      "🇦🇬",
		Descriptor: "flag: Antigua & Barbuda",
	},
	"1F1E6-1F1EE": {
		Key:        "1F1E6-1F1EE",
		Value:      "🇦🇮",
		Descriptor: "flag: Anguilla",
	},
	"1F1E6-1F1F1": {
		Key:        "1F1E6-1F1F1",
		Value:      "🇦🇱",
		Descriptor: "flag: Albania",
	},
	"1F1E6-1F1F2": {
		Key:        "1F1E6-1F1F2",
		Value:      "🇦🇲",
		Descriptor: "flag: Armenia",
	},
	"1F1E6-1F1F4": {
		Key:        "1F1E6-1F1F4",
		Value:      "🇦🇴",
		Descriptor: "flag: Angola",
	},
	"1F1E6-1F1F6": {
		Key:        "1F1E6-1F1F6",
		Value:      "🇦🇶",
		Descriptor: "flag: Antarctica",
	},
	"1F1E6-1F1F7": {
		Key:        "1F1E6-1F1F7",
		Value:      "🇦🇷",
		Descriptor: "flag: Argentina",
	},
	"1F1E6-1F1F8": {
		Key:        "1F1E6-1F1F8",
		Value:      "🇦🇸",
		Descriptor: "flag: American Samoa",
	},
	"1F1E6-1F1F9": {
		Key:        "1F1E6-1F1F9",
		Value:      "🇦🇹",
		Descriptor: "flag: Austria",
	},
	"1F1E6-1F1FA": {
		Key:        "1F1E6-1F1FA",
		Value:      "🇦🇺",
		Descriptor: "flag: Australia",
	},
	"1F1E6-1F1FC": {
		Key:        "1F1E6-1F1FC",
		Value:      "🇦🇼",
		Descriptor: "flag: Aruba",
	},
	"1F1E6-1F1FD": {
		Key:        "1F1E6-1F1FD",
		Value:      "🇦🇽",
		Descriptor: "flag: Åland Islands",
	},
	"1F1E6-1F1FF": {
		Key:        "1F1E6-1F1FF",
		Value:      "🇦🇿",
		Descriptor: "flag: Azerbaijan",
	},
	"1F1E7-1F1E6": {
		Key:        "1F1E7-1F1E6",
		Value:      "🇧🇦",
		Descriptor: "flag: Bosnia & Herzegovina",
	},
	"1F1E7-1F1E7": {
		Key:        "1F1E7-1F1E7",
		Value:      "🇧🇧",
		Descriptor: "flag: Barbados",
	},
	"1F1E7-1F1E9": {
		Key:        "1F1E7-1F1E9",
		Value:      "🇧🇩",
		Descriptor: "flag: Bangladesh",
	},
	"1F1E7-1F1EA": {
		Key:        "1F1E7-1F1EA",
		Value:      "🇧🇪",
		Descriptor: "flag: Belgium",
	},
	"1F1E7-1F1EB": {
		Key:        "1F1E7-1F1EB",
		Value:      "🇧🇫",
		Descriptor: "flag: Burkina Faso",
	},
	"1F1E7-1F1EC": {
		Key:        "1F1E7-1F1EC",
		Value:      "🇧🇬",
		Descriptor: "flag: Bulgaria",
	},
	"1F1E7-1F1ED": {
		Key:        "1F1E7-1F1ED",
		Value:      "🇧🇭",
		Descriptor: "flag: Bahrain",
	},
	"1F1E7-1F1EE": {
		Key:        "1F1E7-1F1EE",
		Value:      "🇧🇮",
		Descriptor: "flag: Burundi",
	},
	"1F1E7-1F1EF": {
		Key:        "1F1E7-1F1EF",
		Value:      "🇧🇯",
		Descriptor: "flag: Benin",
	},
	"1F1E7-1F1F1": {
		Key:        "1F1E7-1F1F1",
		Value:      "🇧🇱",
		Descriptor: "flag: St. Barthélemy",
	},
	"1F1E7-1F1F2": {
		Key:        "1F1E7-1F1F2",
		Value:      "🇧🇲",
		Descriptor: "flag: Bermuda",
	},
	"1F1E7-1F1F3": {
		Key:        "1F1E7-1F1F3",
		Value:      "🇧🇳",
		Descriptor: "flag: Brunei",
	},
	"1F1E7-1F1F4": {
		Key:        "1F1E7-1F1F4",
		Value:      "🇧🇴",
		Descriptor: "flag: Bolivia",
	},
	"1F1E7-1F1F6": {
		Key:        "1F1E7-1F1F6",
		Value:      "🇧🇶",
		Descriptor: "flag: Caribbean Netherlands",
	},
	"1F1E7-1F1F7": {
		Key:        "1F1E7-1F1F7",
		Value:      "🇧🇷",
		Descriptor: "flag: Brazil",
	},
	"1F1E7-1F1F8": {
		Key:        "1F1E7-1F1F8",
		Value:      "🇧🇸",
		Descriptor: "flag: Bahamas",
	},
	"1F1E7-1F1F9": {
		Key:        "1F1E7-1F1F9",
		Value:      "🇧🇹",
		Descriptor: "flag: Bhutan",
	},
	"1F1E7-1F1FB": {
		Key:        "1F1E7-1F1FB",
		Value:      "🇧🇻",
		Descriptor: "flag: Bouvet Island",
	},
	"1F1E7-1F1FC": {
		Key:        "1F1E7-1F1FC",
		Value:      "🇧🇼",
		Descriptor: "flag: Botswana",
	},
	"1F1E7-1F1FE": {
		Key:        "1F1E7-1F1FE",
		Value:      "🇧🇾",
		Descriptor: "flag: Belarus",
	},
	"1F1E7-1F1FF": {
		Key:        "1F1E7-1F1FF",
		Value:      "🇧🇿",
		Descriptor: "flag: Belize",
	},
	"1F1E8-1F1E6": {
		Key:        "1F1E8-1F1E6",
		Value:      "🇨🇦",
		Descriptor: "flag: Canada",
	},
	"1F1E8-1F1E8": {
		Key:        "1F1E8-1F1E8",
		Value:      "🇨🇨",
		Descriptor: "flag: Cocos (Keeling) Islands",
	},
	"1F1E8-1F1E9": {
		Key:        "1F1E8-1F1E9",
		Value:      "🇨🇩",
		Descriptor: "flag: Congo - Kinshasa",
	},
	"1F1E8-1F1EB": {
		Key:        "1F1E8-1F1EB",
		Value:      "🇨🇫",
		Descriptor: "flag: Central African Republic",
	},
	"1F1E8-1F1EC": {
		Key:        "1F1E8-1F1EC",
		Value:      "🇨🇬",
		Descriptor: "flag: Congo - Brazzaville",
	},
	"1F1E8-1F1ED": {
		Key:        "1F1E8-1F1ED",
		Value:      "🇨🇭",
		Descriptor: "flag: Switzerland",
	},
	"1F1E8-1F1EE": {
		Key:        "1F1E8-1F1EE",
		Value:      "🇨🇮",
		Descriptor: "flag: Côte d’Ivoire",
	},
	"1F1E8-1F1F0": {
		Key:        "1F1E8-1F1F0",
		Value:      "🇨🇰",
		Descriptor: "flag: Cook Islands",
	},
	"1F1E8-1F1F1": {
		Key:        "1F1E8-1F1F1",
		Value:      "🇨🇱",
		Descriptor: "flag: Chile",
	},
	"1F1E8-1F1F2": {
		Key:        "1F1E8-1F1F2",
		Value:      "🇨🇲",
		Descriptor: "flag: Cameroon",
	},
	"1F1E8-1F1F3": {
		Key:        "1F1E8-1F1F3",
		Value:      "🇨🇳",
		Descriptor: "flag: China",
	},
	"1F1E8-1F1F4": {
		Key:        "1F1E8-1F1F4",
		Value:      "🇨🇴",
		Descriptor: "flag: Colombia",
	},
	"1F1E8-1F1F5": {
		Key:        "1F1E8-1F1F5",
		Value:      "🇨🇵",
		Descriptor: "flag: Clipperton Island",
	},
	"1F1E8-1F1F7": {
		Key:        "1F1E8-1F1F7",
		Value:      "🇨🇷",
		Descriptor: "flag: Costa Rica",
	},
	"1F1E8-1F1FA": {
		Key:        "1F1E8-1F1FA",
		Value:      "🇨🇺",
		Descriptor: "flag: Cuba",
	},
	"1F1E8-1F1FB": {
		Key:        "1F1E8-1F1FB",
		Value:      "🇨🇻",
		Descriptor: "flag: Cape Verde",
	},
	"1F1E8-1F1FC": {
		Key:        "1F1E8-1F1FC",
		Value:      "🇨🇼",
		Descriptor: "flag: Curaçao",
	},
	"1F1E8-1F1FD": {
		Key:        "1F1E8-1F1FD",
		Value:      "🇨🇽",
		Descriptor: "flag: Christmas Island",
	},
	"1F1E8-1F1FE": {
		Key:        "1F1E8-1F1FE",
		Value:      "🇨🇾",
		Descriptor: "flag: Cyprus",
	},
	"1F1E8-1F1FF": {
		Key:        "1F1E8-1F1FF",
		Value:      "🇨🇿",
		Descriptor: "flag: Czechia",
	},
	"1F1E9-1F1EA": {
		Key:        "1F1E9-1F1EA",
		Value:      "🇩🇪",
		Descriptor: "flag: Germany",
	},
	"1F1E9-1F1EC": {
		Key:        "1F1E9-1F1EC",
		Value:      "🇩🇬",
		Descriptor: "flag: Diego Garcia",
	},
	"1F1E9-1F1EF": {
		Key:        "1F1E9-1F1EF",
		Value:      "🇩🇯",
		Descriptor: "flag: Djibouti",
	},
	"1F1E9-1F1F0": {
		Key:        "1F1E9-1F1F0",
		Value:      "🇩🇰",
		Descriptor: "flag: Denmark",
	},
	"1F1E9-1F1F2": {
		Key:        "1F1E9-1F1F2",
		Value:      "🇩🇲",
		Descriptor: "flag: Dominica",
	},
	"1F1E9-1F1F4": {
		Key:        "1F1E9-1F1F4",
		Value:      "🇩🇴",
		Descriptor: "flag: Dominican Republic",
	},
	"1F1E9-1F1FF": {
		Key:        "1F1E9-1F1FF",
		Value:      "🇩🇿",
		Descriptor: "flag: Algeria",
	},
	"1F1EA-1F1E6": {
		Key:        "1F1EA-1F1E6",
		Value:      "🇪🇦",
		Descriptor: "flag: Ceuta & Melilla",
	},
	"1F1EA-1F1E8": {
		Key:        "1F1EA-1F1E8",
		Value:      "🇪🇨",
		Descriptor: "flag: Ecuador",
	},
	"1F1EA-1F1EA": {
		Key:        "1F1EA-1F1EA",
		Value:      "🇪🇪",
		Descriptor: "flag: Estonia",
	},
	"1F1EA-1F1EC": {
		Key:        "1F1EA-1F1EC",
		Value:      "🇪🇬",
		Descriptor: "flag: Egypt",
	},
	"1F1EA-1F1ED": {
		Key:        "1F1EA-1F1ED",
		Value:      "🇪🇭",
		Descriptor: "flag: Western Sahara",
	},
	"1F1EA-1F1F7": {
		Key:        "1F1EA-1F1F7",
		Value:      "🇪🇷",
		Descriptor: "flag: Eritrea",
	},
	"1F1EA-1F1F8": {
		Key:        "1F1EA-1F1F8",
		Value:      "🇪🇸",
		Descriptor: "flag: Spain",
	},
	"1F1EA-1F1F9": {
		Key:        "1F1EA-1F1F9",
		Value:      "🇪🇹",
		Descriptor: "flag: Ethiopia",
	},
	"1F1EA-1F1FA": {
		Key:        "1F1EA-1F1FA",
		Value:      "🇪🇺",
		Descriptor: "flag: European Union",
	},
	"1F1EB-1F1EE": {
		Key:        "1F1EB-1F1EE",
		Value:      "🇫🇮",
		Descriptor: "flag: Finland",
	},
	"1F1EB-1F1EF": {
		Key:        "1F1EB-1F1EF",
		Value:      "🇫🇯",
		Descriptor: "flag: Fiji",
	},
	"1F1EB-1F1F0": {
		Key:        "1F1EB-1F1F0",
		Value:      "🇫🇰",
		Descriptor: "flag: Falkland Islands",
	},
	"1F1EB-1F1F2": {
		Key:        "1F1EB-1F1F2",
		Value:      "🇫🇲",
		Descriptor: "flag: Micronesia",
	},
	"1F1EB-1F1F4": {
		Key:        "1F1EB-1F1F4",
		Value:      "🇫🇴",
		Descriptor: "flag: Faroe Islands",
	},
	"1F1EB-1F1F7": {
		Key:        "1F1EB-1F1F7",
		Value:      "🇫🇷",
		Descriptor: "flag: France",
	},
	"1F1EC-1F1E6": {
		Key:        "1F1EC-1F1E6",
		Value:      "🇬🇦",
		Descriptor: "flag: Gabon",
	},
	"1F1EC-1F1E7": {
		Key:        "1F1EC-1F1E7",
		Value:      "🇬🇧",
		Descriptor: "flag: United Kingdom",
	},
	"1F1EC-1F1E9": {
		Key:        "1F1EC-1F1E9",
		Value:      "🇬🇩",
		Descriptor: "flag: Grenada",
	},
	"1F1EC-1F1EA": {
		Key:        "1F1EC-1F1EA",
		Value:      "🇬🇪",
		Descriptor: "flag: Georgia",
	},
	"1F1EC-1F1EB": {
		Key:        "1F1EC-1F1EB",
		Value:      "🇬🇫",
		Descriptor: "flag: French Guiana",
	},
	"1F1EC-1F1EC": {
		Key:        "1F1EC-1F1EC",
		Value:      "🇬🇬",
		Descriptor: "flag: Guernsey",
	},
	"1F1EC-1F1ED": {
		Key:        "1F1EC-1F1ED",
		Value:      "🇬🇭",
		Descriptor: "flag: Ghana",
	},
	"1F1EC-1F1EE": {
		Key:        "1F1EC-1F1EE",
		Value:      "🇬🇮",
		Descriptor: "flag: Gibraltar",
	},
	"1F1EC-1F1F1": {
		Key:        "1F1EC-1F1F1",
		Value:      "🇬🇱",
		Descriptor: "flag: Greenland",
	},
	"1F1EC-1F1F2": {
		Key:        "1F1EC-1F1F2",
		Value:      "🇬🇲",
		Descriptor: "flag: Gambia",
	},
	"1F1EC-1F1F3": {
		Key:        "1F1EC-1F1F3",
		Value:      "🇬🇳",
		Descriptor: "flag: Guinea",
	},
	"1F1EC-1F1F5": {
		Key:        "1F1EC-1F1F5",
		Value:      "🇬🇵",
		Descriptor: "flag: Guadeloupe",
	},
	"1F1EC-1F1F6": {
		Key:        "1F1EC-1F1F6",
		Value:      "🇬🇶",
		Descriptor: "flag: Equatorial Guinea",
	},
	"1F1EC-1F1F7": {
		Key:        "1F1EC-1F1F7",
		Value:      "🇬🇷",
		Descriptor: "flag: Greece",
	},
	"1F1EC-1F1F8": {
		Key:        "1F1EC-1F1F8",
		Value:      "🇬🇸",
		Descriptor: "flag: South Georgia & South Sandwich Islands",
	},
	"1F1EC-1F1F9": {
		Key:        "1F1EC-1F1F9",
		Value:      "🇬🇹",
		Descriptor: "flag: Guatemala",
	},
	"1F1EC-1F1FA": {
		Key:        "1F1EC-1F1FA",
		Value:      "🇬🇺",
		Descriptor: "flag: Guam",
	},
	"1F1EC-1F1FC": {
		Key:        "1F1EC-1F1FC",
		Value:      "🇬🇼",
		Descriptor: "flag: Guinea-Bissau",
	},
	"1F1EC-1F1FE": {
		Key:        "1F1EC-1F1FE",
		Value:      "🇬🇾",
		Descriptor: "flag: Guyana",
	},
	"1F1ED-1F1F0": {
		Key:        "1F1ED-1F1F0",
		Value:      "🇭🇰",
		Descriptor: "flag: Hong Kong SAR China",
	},
	"1F1ED-1F1F2": {
		Key:        "1F1ED-1F1F2",
		Value:      "🇭🇲",
		Descriptor: "flag: Heard & McDonald Islands",
	},
	"1F1ED-1F1F3": {
		Key:        "1F1ED-1F1F3",
		Value:      "🇭🇳",
		Descriptor: "flag: Honduras",
	},
	"1F1ED-1F1F7": {
		Key:        "1F1ED-1F1F7",
		Value:      "🇭🇷",
		Descriptor: "flag: Croatia",
	},
	"1F1ED-1F1F9": {
		Key:        "1F1ED-1F1F9",
		Value:      "🇭🇹",
		Descriptor: "flag: Haiti",
	},
	"1F1ED-1F1FA": {
		Key:        "1F1ED-1F1FA",
		Value:      "🇭🇺",
		Descriptor: "flag: Hungary",
	},
	"1F1EE-1F1E8": {
		Key:        "1F1EE-1F1E8",
		Value:      "🇮🇨",
		Descriptor: "flag: Canary Islands",
	},
	"1F1EE-1F1E9": {
		Key:        "1F1EE-1F1E9",
		Value:      "🇮🇩",
		Descriptor: "flag: Indonesia",
	},
	"1F1EE-1F1EA": {
		Key:        "1F1EE-1F1EA",
		Value:      "🇮🇪",
		Descriptor: "flag: Ireland",
	},
	"1F1EE-1F1F1": {
		Key:        "1F1EE-1F1F1",
		Value:      "🇮🇱",
		Descriptor: "flag: Israel",
	},
	"1F1EE-1F1F2": {
		Key:        "1F1EE-1F1F2",
		Value:      "🇮🇲",
		Descriptor: "flag: Isle of Man",
	},
	"1F1EE-1F1F3": {
		Key:        "1F1EE-1F1F3",
		Value:      "🇮🇳",
		Descriptor: "flag: India",
	},
	"1F1EE-1F1F4": {
		Key:        "1F1EE-1F1F4",
		Value:      "🇮🇴",
		Descriptor: "flag: British Indian Ocean Territory",
	},
	"1F1EE-1F1F6": {
		Key:        "1F1EE-1F1F6",
		Value:      "🇮🇶",
		Descriptor: "flag: Iraq",
	},
	"1F1EE-1F1F7": {
		Key:        "1F1EE-1F1F7",
		Value:      "🇮🇷",
		Descriptor: "flag: Iran",
	},
	"1F1EE-1F1F8": {
		Key:        "1F1EE-1F1F8",
		Value:      "🇮🇸",
		Descriptor: "flag: Iceland",
	},
	"1F1EE-1F1F9": {
		Key:        "1F1EE-1F1F9",
		Value:      "🇮🇹",
		Descriptor: "flag: Italy",
	},
	"1F1EF-1F1EA": {
		Key:        "1F1EF-1F1EA",
		Value:      "🇯🇪",
		Descriptor: "flag: Jersey",
	},
	"1F1EF-1F1F2": {
		Key:        "1F1EF-1F1F2",
		Value:      "🇯🇲",
		Descriptor: "flag: Jamaica",
	},
	"1F1EF-1F1F4": {
		Key:        "1F1EF-1F1F4",
		Value:      "🇯🇴",
		Descriptor: "flag: Jordan",
	},
	"1F1EF-1F1F5": {
		Key:        "1F1EF-1F1F5",
		Value:      "🇯🇵",
		Descriptor: "flag: Japan",
	},
	"1F1F0-1F1EA": {
		Key:        "1F1F0-1F1EA",
		Value:      "🇰🇪",
		Descriptor: "flag: Kenya",
	},
	"1F1F0-1F1EC": {
		Key:        "1F1F0-1F1EC",
		Value:      "🇰🇬",
		Descriptor: "flag: Kyrgyzstan",
	},
	"1F1F0-1F1ED": {
		Key:        "1F1F0-1F1ED",
		Value:      "🇰🇭",
		Descriptor: "flag: Cambodia",
	},
	"1F1F0-1F1EE": {
		Key:        "1F1F0-1F1EE",
		Value:      "🇰🇮",
		Descriptor: "flag: Kiribati",
	},
	"1F1F0-1F1F2": {
		Key:        "1F1F0-1F1F2",
		Value:      "🇰🇲",
		Descriptor: "flag: Comoros",
	},
	"1F1F0-1F1F3": {
		Key:        "1F1F0-1F1F3",
		Value:      "🇰🇳",
		Descriptor: "flag: St. Kitts & Nevis",
	},
	"1F1F0-1F1F5": {
		Key:        "1F1F0-1F1F5",
		Value:      "🇰🇵",
		Descriptor: "flag: North Korea",
	},
	"1F1F0-1F1F7": {
		Key:        "1F1F0-1F1F7",
		Value:      "🇰🇷",
		Descriptor: "flag: South Korea",
	},
	"1F1F0-1F1FC": {
		Key:        "1F1F0-1F1FC",
		Value:      "🇰🇼",
		Descriptor: "flag: Kuwait",
	},
	"1F1F0-1F1FE": {
		Key:        "1F1F0-1F1FE",
		Value:      "🇰🇾",
		Descriptor: "flag: Cayman Islands",
	},
	"1F1F0-1F1FF": {
		Key:        "1F1F0-1F1FF",
		Value:      "🇰🇿",
		Descriptor: "flag: Kazakhstan",
	},
	"1F1F1-1F1E6": {
		Key:        "1F1F1-1F1E6",
		Value:      "🇱🇦",
		Descriptor: "flag: Laos",
	},
	"1F1F1-1F1E7": {
		Key:        "1F1F1-1F1E7",
		Value:      "🇱🇧",
		Descriptor: "flag: Lebanon",
	},
	"1F1F1-1F1E8": {
		Key:        "1F1F1-1F1E8",
		Value:      "🇱🇨",
		Descriptor: "flag: St. Lucia",
	},
	"1F1F1-1F1EE": {
		Key:        "1F1F1-1F1EE",
		Value:      "🇱🇮",
		Descriptor: "flag: Liechtenstein",
	},
	"1F1F1-1F1F0": {
		Key:        "1F1F1-1F1F0",
		Value:      "🇱🇰",
		Descriptor: "flag: Sri Lanka",
	},
	"1F1F1-1F1F7": {
		Key:        "1F1F1-1F1F7",
		Value:      "🇱🇷",
		Descriptor: "flag: Liberia",
	},
	"1F1F1-1F1F8": {
		Key:        "1F1F1-1F1F8",
		Value:      "🇱🇸",
		Descriptor: "flag: Lesotho",
	},
	"1F1F1-1F1F9": {
		Key:        "1F1F1-1F1F9",
		Value:      "🇱🇹",
		Descriptor: "flag: Lithuania",
	},
	"1F1F1-1F1FA": {
		Key:        "1F1F1-1F1FA",
		Value:      "🇱🇺",
		Descriptor: "flag: Luxembourg",
	},
	"1F1F1-1F1FB": {
		Key:        "1F1F1-1F1FB",
		Value:      "🇱🇻",
		Descriptor: "flag: Latvia",
	},
	"1F1F1-1F1FE": {
		Key:        "1F1F1-1F1FE",
		Value:      "🇱🇾",
		Descriptor: "flag: Libya",
	},
	"1F1F2-1F1E6": {
		Key:        "1F1F2-1F1E6",
		Value:      "🇲🇦",
		Descriptor: "flag: Morocco",
	},
	"1F1F2-1F1E8": {
		Key:        "1F1F2-1F1E8",
		Value:      "🇲🇨",
		Descriptor: "flag: Monaco",
	},
	"1F1F2-1F1E9": {
		Key:        "1F1F2-1F1E9",
		Value:      "🇲🇩",
		Descriptor: "flag: Moldova",
	},
	"1F1F2-1F1EA": {
		Key:        "1F1F2-1F1EA",
		Value:      "🇲🇪",
		Descriptor: "flag: Montenegro",
	},
	"1F1F2-1F1EB": {
		Key:        "1F1F2-1F1EB",
		Value:      "🇲🇫",
		Descriptor: "flag: St. Martin",
	},
	"1F1F2-1F1EC": {
		Key:        "1F1F2-1F1EC",
		Value:      "🇲🇬",
		Descriptor: "flag: Madagascar",
	},
	"1F1F2-1F1ED": {
		Key:        "1F1F2-1F1ED",
		Value:      "🇲🇭",
		Descriptor: "flag: Marshall Islands",
	},
	"1F1F2-1F1F0": {
		Key:        "1F1F2-1F1F0",
		Value:      "🇲🇰",
		Descriptor: "flag: North Macedonia",
	},
	"1F1F2-1F1F1": {
		Key:        "1F1F2-1F1F1",
		Value:      "🇲🇱",
		Descriptor: "flag: Mali",
	},
	"1F1F2-1F1F2": {
		Key:        "1F1F2-1F1F2",
		Value:      "🇲🇲",
		Descriptor: "flag: Myanmar (Burma)",
	},
	"1F1F2-1F1F3": {
		Key:        "1F1F2-1F1F3",
		Value:      "🇲🇳",
		Descriptor: "flag: Mongolia",
	},
	"1F1F2-1F1F4": {
		Key:        "1F1F2-1F1F4",
		Value:      "🇲🇴",
		Descriptor: "flag: Macao SAR China",
	},
	"1F1F2-1F1F5": {
		Key:        "1F1F2-1F1F5",
		Value:      "🇲🇵",
		Descriptor: "flag: Northern Mariana Islands",
	},
	"1F1F2-1F1F6": {
		Key:        "1F1F2-1F1F6",
		Value:      "🇲🇶",
		Descriptor: "flag: Martinique",
	},
	"1F1F2-1F1F7": {
		Key:        "1F1F2-1F1F7",
		Value:      "🇲🇷",
		Descriptor: "flag: Mauritania",
	},
	"1F1F2-1F1F8": {
		Key:        "1F1F2-1F1F8",
		Value:      "🇲🇸",
		Descriptor: "flag: Montserrat",
	},
	"1F1F2-1F1F9": {
		Key:        "1F1F2-1F1F9",
		Value:      "🇲🇹",
		Descriptor: "flag: Malta",
	},
	"1F1F2-1F1FA": {
		Key:        "1F1F2-1F1FA",
		Value:      "🇲🇺",
		Descriptor: "flag: Mauritius",
	},
	"1F1F2-1F1FB": {
		Key:        "1F1F2-1F1FB",
		Value:      "🇲🇻",
		Descriptor: "flag: Maldives",
	},
	"1F1F2-1F1FC": {
		Key:        "1F1F2-1F1FC",
		Value:      "🇲🇼",
		Descriptor: "flag: Malawi",
	},
	"1F1F2-1F1FD": {
		Key:        "1F1F2-1F1FD",
		Value:      "🇲🇽",
		Descriptor: "flag: Mexico",
	},
	"1F1F2-1F1FE": {
		Key:        "1F1F2-1F1FE",
		Value:      "🇲🇾",
		Descriptor: "flag: Malaysia",
	},
	"1F1F2-1F1FF": {
		Key:        "1F1F2-1F1FF",
		Value:      "🇲🇿",
		Descriptor: "flag: Mozambique",
	},
	"1F1F3-1F1E6": {
		Key:        "1F1F3-1F1E6",
		Value:      "🇳🇦",
		Descriptor: "flag: Namibia",
	},
	"1F1F3-1F1E8": {
		Key:        "1F1F3-1F1E8",
		Value:      "🇳🇨",
		Descriptor: "flag: New Caledonia",
	},
	"1F1F3-1F1EA": {
		Key:        "1F1F3-1F1EA",
		Value:      "🇳🇪",
		Descriptor: "flag: Niger",
	},
	"1F1F3-1F1EB": {
		Key:        "1F1F3-1F1EB",
		Value:      "🇳🇫",
		Descriptor: "flag: Norfolk Island",
	},
	"1F1F3-1F1EC": {
		Key:        "1F1F3-1F1EC",
		Value:      "🇳🇬",
		Descriptor: "flag: Nigeria",
	},
	"1F1F3-1F1EE": {
		Key:        "1F1F3-1F1EE",
		Value:      "🇳🇮",
		Descriptor: "flag: Nicaragua",
	},
	"1F1F3-1F1F1": {
		Key:        "1F1F3-1F1F1",
		Value:      "🇳🇱",
		Descriptor: "flag: Netherlands",
	},
	"1F1F3-1F1F4": {
		Key:        "1F1F3-1F1F4",
		Value:      "🇳🇴",
		Descriptor: "flag: Norway",
	},
	"1F1F3-1F1F5": {
		Key:        "1F1F3-1F1F5",
		Value:      "🇳🇵",
		Descriptor: "flag: Nepal",
	},
	"1F1F3-1F1F7": {
		Key:        "1F1F3-1F1F7",
		Value:      "🇳🇷",
		Descriptor: "flag: Nauru",
	},
	"1F1F3-1F1FA": {
		Key:        "1F1F3-1F1FA",
		Value:      "🇳🇺",
		Descriptor: "flag: Niue",
	},
	"1F1F3-1F1FF": {
		Key:        "1F1F3-1F1FF",
		Value:      "🇳🇿",
		Descriptor: "flag: New Zealand",
	},
	"1F1F4-1F1F2": {
		Key:        "1F1F4-1F1F2",
		Value:      "🇴🇲",
		Descriptor: "flag: Oman",
	},
	"1F1F5-1F1E6": {
		Key:        "1F1F5-1F1E6",
		Value:      "🇵🇦",
		Descriptor: "flag: Panama",
	},
	"1F1F5-1F1EA": {
		Key:        "1F1F5-1F1EA",
		Value:      "🇵🇪",
		Descriptor: "flag: Peru",
	},
	"1F1F5-1F1EB": {
		Key:        "1F1F5-1F1EB",
		Value:      "🇵🇫",
		Descriptor: "flag: French Polynesia",
	},
	"1F1F5-1F1EC": {
		Key:        "1F1F5-1F1EC",
		Value:      "🇵🇬",
		Descriptor: "flag: Papua New Guinea",
	},
	"1F1F5-1F1ED": {
		Key:        "1F1F5-1F1ED",
		Value:      "🇵🇭",
		Descriptor: "flag: Philippines",
	},
	"1F1F5-1F1F0": {
		Key:        "1F1F5-1F1F0",
		Value:      "🇵🇰",
		Descriptor: "flag: Pakistan",
	},
	"1F1F5-1F1F1": {
		Key:        "1F1F5-1F1F1",
		Value:      "🇵🇱",
		Descriptor: "flag: Poland",
	},
	"1F1F5-1F1F2": {
		Key:        "1F1F5-1F1F2",
		Value:      "🇵🇲",
		Descriptor: "flag: St. Pierre & Miquelon",
	},
	"1F1F5-1F1F3": {
		Key:        "1F1F5-1F1F3",
		Value:      "🇵🇳",
		Descriptor: "flag: Pitcairn Islands",
	},
	"1F1F5-1F1F7": {
		Key:        "1F1F5-1F1F7",
		Value:      "🇵🇷",
		Descriptor: "flag: Puerto Rico",
	},
	"1F1F5-1F1F8": {
		Key:        "1F1F5-1F1F8",
		Value:      "🇵🇸",
		Descriptor: "flag: Palestinian Territories",
	},
	"1F1F5-1F1F9": {
		Key:        "1F1F5-1F1F9",
		Value:      "🇵🇹",
		Descriptor: "flag: Portugal",
	},
	"1F1F5-1F1FC": {
		Key:        "1F1F5-1F1FC",
		Value:      "🇵🇼",
		Descriptor: "flag: Palau",
	},
	"1F1F5-1F1FE": {
		Key:        "1F1F5-1F1FE",
		Value:      "🇵🇾",
		Descriptor: "flag: Paraguay",
	},
	"1F1F6-1F1E6": {
		Key:        "1F1F6-1F1E6",
		Value:      "🇶🇦",
		Descriptor: "flag: Qatar",
	},
	"1F1F7-1F1EA": {
		Key:        "1F1F7-1F1EA",
		Value:      "🇷🇪",
		Descriptor: "flag: Réunion",
	},
	"1F1F7-1F1F4": {
		Key:        "1F1F7-1F1F4",
		Value:      "🇷🇴",
		Descriptor: "flag: Romania",
	},
	"1F1F7-1F1F8": {
		Key:        "1F1F7-1F1F8",
		Value:      "🇷🇸",
		Descriptor: "flag: Serbia",
	},
	"1F1F7-1F1FA": {
		Key:        "1F1F7-1F1FA",
		Value:      "🇷🇺",
		Descriptor: "flag: Russia",
	},
	"1F1F7-1F1FC": {
		Key:        "1F1F7-1F1FC",
		Value:      "🇷🇼",
		Descriptor: "flag: Rwanda",
	},
	"1F1F8-1F1E6": {
		Key:        "1F1F8-1F1E6",
		Value:      "🇸🇦",
		Descriptor: "flag: Saudi Arabia",
	},
	"1F1F8-1F1E7": {
		Key:        "1F1F8-1F1E7",
		Value:      "🇸🇧",
		Descriptor: "flag: Solomon Islands",
	},
	"1F1F8-1F1E8": {
		Key:        "1F1F8-1F1E8",
		Value:      "🇸🇨",
		Descriptor: "flag: Seychelles",
	},
	"1F1F8-1F1E9": {
		Key:        "1F1F8-1F1E9",
		Value:      "🇸🇩",
		Descriptor: "flag: Sudan",
	},
	"1F1F8-1F1EA": {
		Key:        "1F1F8-1F1EA",
		Value:      "🇸🇪",
		Descriptor: "flag: Sweden",
	},
	"1F1F8-1F1EC": {
		Key:        "1F1F8-1F1EC",
		Value:      "🇸🇬",
		Descriptor: "flag: Singapore",
	},
	"1F1F8-1F1ED": {
		Key:        "1F1F8-1F1ED",
		Value:      "🇸🇭",
		Descriptor: "flag: St. Helena",
	},
	"1F1F8-1F1EE": {
		Key:        "1F1F8-1F1EE",
		Value:      "🇸🇮",
		Descriptor: "flag: Slovenia",
	},
	"1F1F8-1F1EF": {
		Key:        "1F1F8-1F1EF",
		Value:      "🇸🇯",
		Descriptor: "flag: Svalbard & Jan Mayen",
	},
	"1F1F8-1F1F0": {
		Key:        "1F1F8-1F1F0",
		Value:      "🇸🇰",
		Descriptor: "flag: Slovakia",
	},
	"1F1F8-1F1F1": {
		Key:        "1F1F8-1F1F1",
		Value:      "🇸🇱",
		Descriptor: "flag: Sierra Leone",
	},
	"1F1F8-1F1F2": {
		Key:        "1F1F8-1F1F2",
		Value:      "🇸🇲",
		Descriptor: "flag: San Marino",
	},
	"1F1F8-1F1F3": {
		Key:        "1F1F8-1F1F3",
		Value:      "🇸🇳",
		Descriptor: "flag: Senegal",
	},
	"1F1F8-1F1F4": {
		Key:        "1F1F8-1F1F4",
		Value:      "🇸🇴",
		Descriptor: "flag: Somalia",
	},
	"1F1F8-1F1F7": {
		Key:        "1F1F8-1F1F7",
		Value:      "🇸🇷",
		Descriptor: "flag: Suriname",
	},
	"1F1F8-1F1F8": {
		Key:        "1F1F8-1F1F8",
		Value:      "🇸🇸",
		Descriptor: "flag: South Sudan",
	},
	"1F1F8-1F1F9": {
		Key:        "1F1F8-1F1F9",
		Value:      "🇸🇹",
		Descriptor: "flag: São Tomé & Príncipe",
	},
	"1F1F8-1F1FB": {
		Key:        "1F1F8-1F1FB",
		Value:      "🇸🇻",
		Descriptor: "flag: El Salvador",
	},
	"1F1F8-1F1FD": {
		Key:        "1F1F8-1F1FD",
		Value:      "🇸🇽",
		Descriptor: "flag: Sint Maarten",
	},
	"1F1F8-1F1FE": {
		Key:        "1F1F8-1F1FE",
		Value:      "🇸🇾",
		Descriptor: "flag: Syria",
	},
	"1F1F8-1F1FF": {
		Key:        "1F1F8-1F1FF",
		Value:      "🇸🇿",
		Descriptor: "flag: Eswatini",
	},
	"1F1F9-1F1E6": {
		Key:        "1F1F9-1F1E6",
		Value:      "🇹🇦",
		Descriptor: "flag: Tristan da Cunha",
	},
	"1F1F9-1F1E8": {
		Key:        "1F1F9-1F1E8",
		Value:      "🇹🇨",
		Descriptor: "flag: Turks & Caicos Islands",
	},
	"1F1F9-1F1E9": {
		Key:        "1F1F9-1F1E9",
		Value:      "🇹🇩",
		Descriptor: "flag: Chad",
	},
	"1F1F9-1F1EB": {
		Key:        "1F1F9-1F1EB",
		Value:      "🇹🇫",
		Descriptor: "flag: French Southern Territories",
	},
	"1F1F9-1F1EC": {
		Key:        "1F1F9-1F1EC",
		Value:      "🇹🇬",
		Descriptor: "flag: Togo",
	},
	"1F1F9-1F1ED": {
		Key:        "1F1F9-1F1ED",
		Value:      "🇹🇭",
		Descriptor: "flag: Thailand",
	},
	"1F1F9-1F1EF": {
		Key:        "1F1F9-1F1EF",
		Value:      "🇹🇯",
		Descriptor: "flag: Tajikistan",
	},
	"1F1F9-1F1F0": {
		Key:        "1F1F9-1F1F0",
		Value:      "🇹🇰",
		Descriptor: "flag: Tokelau",
	},
	"1F1F9-1F1F1": {
		Key:        "1F1F9-1F1F1",
		Value:      "🇹🇱",
		Descriptor: "flag: Timor-Leste",
	},
	"1F1F9-1F1F2": {
		Key:        "1F1F9-1F1F2",
		Value:      "🇹🇲",
		Descriptor: "flag: Turkmenistan",
	},
	"1F1F9-1F1F3": {
		Key:        "1F1F9-1F1F3",
		Value:      "🇹🇳",
		Descriptor: "flag: Tunisia",
	},
	"1F1F9-1F1F4": {
		Key:        "1F1F9-1F1F4",
		Value:      "🇹🇴",
		Descriptor: "flag: Tonga",
	},
	"1F1F9-1F1F7": {
		Key:        "1F1F9-1F1F7",
		Value:      "🇹🇷",
		Descriptor: "flag: Türkiye",
	},
	"1F1F9-1F1F9": {
		Key:        "1F1F9-1F1F9",
		Value:      "🇹🇹",
		Descriptor: "flag: Trinidad & Tobago",
	},
	"1F1F9-1F1FB": {
		Key:        "1F1F9-1F1FB",
		Value:      "🇹🇻",
		Descriptor: "flag: Tuvalu",
	},
	"1F1F9-1F1FC": {
		Key:        "1F1F9-1F1FC",
		Value:      "🇹🇼",
		Descriptor: "flag: Taiwan",
	},
	"1F1F9-1F1FF": {
		Key:        "1F1F9-1F1FF",
		Value:      "🇹🇿",
		Descriptor: "flag: Tanzania",
	},
	"1F1FA-1F1E6": {
		Key:        "1F1FA-1F1E6",
		Value:      "🇺🇦",
		Descriptor: "flag: Ukraine",
	},
	"1F1FA-1F1EC": {
		Key:        "1F1FA-1F1EC",
		Value:      "🇺🇬",
		Descriptor: "flag: Uganda",
	},
	"1F1FA-1F1F2": {
		Key:        "1F1FA-1F1F2",
		Value:      "🇺🇲",
		Descriptor: "flag: U.S. Outlying Islands",
	},
	"1F1FA-1F1F3": {
		Key:        "1F1FA-1F1F3",
		Value:      "🇺🇳",
		Descriptor: "flag: United Nations",
	},
	"1F1FA-1F1F8": {
		Key:        "1F1FA-1F1F8",
		Value:      "🇺🇸",
		Descriptor: "flag: United States",
	},
	"1F1FA-1F1FE": {
		Key:        "1F1FA-1F1FE",
		Value:      "🇺🇾",
		Descriptor: "flag: Uruguay",
	},
	"1F1FA-1F1FF": {
		Key:        "1F1FA-1F1FF",
		Value:      "🇺🇿",
		Descriptor: "flag: Uzbekistan",
	},
	"1F1FB-1F1E6": {
		Key:        "1F1FB-1F1E6",
		Value:      "🇻🇦",
		Descriptor: "flag: Vatican City",
	},
	"1F1FB-1F1E8": {
		Key:        "1F1FB-1F1E8",
		Value:      "🇻🇨",
		Descriptor: "flag: St. Vincent & Grenadines",
	},
	"1F1FB-1F1EA": {
		Key:        "1F1FB-1F1EA",
		Value:      "🇻🇪",
		Descriptor: "flag: Venezuela",
	},
	"1F1FB-1F1EC": {
		Key:        "1F1FB-1F1EC",
		Value:      "🇻🇬",
		Descriptor: "flag: British Virgin Islands",
	},
	"1F1FB-1F1EE": {
		Key:        "1F1FB-1F1EE",
		Value:      "🇻🇮",
		Descriptor: "flag: U.S. Virgin Islands",
	},
	"1F1FB-1F1F3": {
		Key:        "1F1FB-1F1F3",
		Value:      "🇻🇳",
		Descriptor: "flag: Vietnam",
	},
	"1F1FB-1F1FA": {
		Key:        "1F1FB-1F1FA",
		Value:      "🇻🇺",
		Descriptor: "flag: Vanuatu",
	},
	"1F1FC-1F1EB": {
		Key:        "1F1FC-1F1EB",
		Value:      "🇼🇫",
		Descriptor: "flag: Wallis & Futuna",
	},
	"1F1FC-1F1F8": {
		Key:        "1F1FC-1F1F8",
		Value:      "🇼🇸",
		Descriptor: "flag: Samoa",
	},
	"1F1FD-1F1F0": {
		Key:        "1F1FD-1F1F0",
		Value:      "🇽🇰",
		Descriptor: "flag: Kosovo",
	},
	"1F1FE-1F1EA": {
		Key:        "1F1FE-1F1EA",
		Value:      "🇾🇪",
		Descriptor: "flag: Yemen",
	},
	"1F1FE-1F1F9": {
		Key:        "1F1FE-1F1F9",
		Value:      "🇾🇹",
		Descriptor: "flag: Mayotte",
	},
	"1F1FF-1F1E6": {
		Key:        "1F1FF-1F1E6",
		Value:      "🇿🇦",
		Descriptor: "flag: South Africa",
	},
	"1F1FF-1F1F2": {
		Key:        "1F1FF-1F1F2",
		Value:      "🇿🇲",
		Descriptor: "flag: Zambia",
	},
	"1F1FF-1F1FC": {
		Key:        "1F1FF-1F1FC",
		Value:      "🇿🇼",
		Descriptor: "flag: Zimbabwe",
	},
	"1F3F4-E0067-E0062-E0065-E006E-E0067-E007F": {
		Key:        "1F3F4-E0067-E0062-E0065-E006E-E0067-E007F",
		Value:      "🏴󠁧󠁢󠁥󠁮󠁧󠁿",
		Descriptor: "flag: England",
	},
	"1F3F4-E0067-E0062-E0073-E0063-E0074-E007F": {
		Key:        "1F3F4-E0067-E0062-E0073-E0063-E0074-E007F",
		Value:      "🏴󠁧󠁢󠁳󠁣󠁴󠁿",
		Descriptor: "flag: Scotland",
	},
	"1F3F4-E0067-E0062-E0077-E006C-E0073-E007F": {
		Key:        "1F3F4-E0067-E0062-E0077-E006C-E0073-E007F",
		Value:      "🏴󠁧󠁢󠁷󠁬󠁳󠁿",
		Descriptor: "flag: Wales",
	},
}
