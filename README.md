# GOverwatch
Unofficial Overwatch API programmed from scratch using GoLang<br/>
<br/>
API Calls:<br/>
<br/>
URL: /heroes<br/>
Description: Displays a list of all heroes and some light metadata<br/>
Sample:<br/>
```json
{
	"Status": "200",
	"RequestType": "heroesIndex",
	"Data": [
		{
			"Name": "Genji",
			"Roles": "[\"OFFENSE\"]",
			"Portrait": "https://blzgdapipro-a.akamaihd.net/hero/genji/hero-select-portrait.png"
		},
		{
			"Name": "McCree",
			"Roles": "[\"OFFENSE\"]",
			"Portrait": "https://blzgdapipro-a.akamaihd.net/hero/mccree/hero-select-portrait.png"
		}, 
		...
	]
}
```
<br/>
URL: /heroes/{heroName}<br/>
Description: Displays in-detail metadata about a specific hero<br/>
Sample:<br/>

```json
{
	"Status": "200",
	"RequestType": "heroDetail",
	"Data": {
		"Name": "hanzo",
		"Role": "Defense",
		"Description": "Hanzo’s versatile arrows can reveal his enemies or fragment to strike multiple targets. He can scale walls to fire his bow from on high, or summon a titanic spirit dragon.",
		"Difficulty": 3,
		"Videos": [
			"https://blzgdapipro-a.akamaihd.net/hero/hanzo/idle-video.webm",
			"https://blzgdapipro-a.akamaihd.net/hero/hanzo/idle-video.mp4"
		],
		"Abilities": [
			{
				"Name": "Storm Bow",
				"Description": "Hanzo nocks and fires an arrow at his target.",
				"Icon": "https://blzgdapipro-a.akamaihd.net/hero/hanzo/ability-storm-bow/icon-ability.png"
			},
			{
				"Name": "Sonic Arrow",
				"Description": "Hanzo launches an arrow that contains a sonar tracking device. Any enemy within its detection radius is visibly marked, making them easier for Hanzo and his allies to hunt down.",
				"Icon": "https://blzgdapipro-a.akamaihd.net/hero/hanzo/ability-sonic-arrow/icon-ability.png"
			},
			{
				"Name": "Scatter Arrow",
				"Description": "Hanzo shoots a  fragmenting arrow that ricochets off walls and objects and can strike multiple targets at once.",
				"Icon": "https://blzgdapipro-a.akamaihd.net/hero/hanzo/ability-scatter-arrow/icon-ability.png"
			},
			{
				"Name": "Dragonstrike",
				"Description": "Hanzo summons a Spirit Dragon which travels through the air in a line. It passes through walls in its way,  devouring any enemies it encounters.",
				"Icon": "https://blzgdapipro-a.akamaihd.net/hero/hanzo/ability-dragon-strike/icon-ability.png"
			}
		]
	}
}
```

<br/>
URL: /players/{platform}/{region}/{gameMode}/{battleTag}<br/>
Description: Displays detailed profile information and statistics about a certain player (battle tag)<br/>
Sample:<br/>

```json
{
	"Status": "200",
	"RequestType": "PlayerDetail",
	"Data": {
		"Name": "Gethos",
		"Battletag": "Gethos-1743",
		"Portrait": "https://blzgdapipro-a.akamaihd.net/game/unlocks/0x02500000000008C1.png",
		"Rank": [
			"https://blzgdapipro-a.akamaihd.net/game/rank-icons/season-2/rank-4.png",
			"24832483"
		],
		"FeaturedStats": [
			{
				"Name": "Eliminations - Average",
				"Value": "13.84",
				"Icon": "<svg viewBox=\"0 0 32 32\" class=\"icon\"><use xlink:href=\"https://overwatch-a.akamaihd.net/img/icons/career-icons-9f59a643a3181e5bb684871c67ae3b62b6476ddeebd073310fe61baf6de3322ebeb80ec0e1f31a6817d164c6b2856c9f1830a3eeb11fb2c1d119a11dfba17437.svg#0x0860000000000383\"></use></svg>"
			},
			...
		],
		"CareerStats": [
			{
				"Name": "Combat",
				"Icon": "<use xlink:href=\"https://overwatch-a.akamaihd.net/img/icons/career-icons-9f59a643a3181e5bb684871c67ae3b62b6476ddeebd073310fe61baf6de3322ebeb80ec0e1f31a6817d164c6b2856c9f1830a3eeb11fb2c1d119a11dfba17437.svg#combat\"></use>",
				"Stats": [
					{
						"Name": "Melee Final Blows",
						"Value": "95",
						"Icon": ""
					},
					...
				]
			},
			{
				"Name": "Assists",
				"Icon": "<use xlink:href=\"https://overwatch-a.akamaihd.net/img/icons/career-icons-9f59a643a3181e5bb684871c67ae3b62b6476ddeebd073310fe61baf6de3322ebeb80ec0e1f31a6817d164c6b2856c9f1830a3eeb11fb2c1d119a11dfba17437.svg#assists\"></use>",
				"Stats": [
					{
						"Name": "Healing Done",
						"Value": "529,596",
						"Icon": ""
					},
					...
				]
			},
			{
				"Name": "Best",
				"Icon": "<use xlink:href=\"https://overwatch-a.akamaihd.net/img/icons/career-icons-9f59a643a3181e5bb684871c67ae3b62b6476ddeebd073310fe61baf6de3322ebeb80ec0e1f31a6817d164c6b2856c9f1830a3eeb11fb2c1d119a11dfba17437.svg#best\"></use>",
				"Stats": [
					{
						"Name": "Eliminations - Most in Game",
						"Value": "44",
						"Icon": ""
					},
					...
				]
			},
			{
				"Name": "Average",
				"Icon": "<use xlink:href=\"https://overwatch-a.akamaihd.net/img/icons/career-icons-9f59a643a3181e5bb684871c67ae3b62b6476ddeebd073310fe61baf6de3322ebeb80ec0e1f31a6817d164c6b2856c9f1830a3eeb11fb2c1d119a11dfba17437.svg#average\"></use>",
				"Stats": [
					{
						"Name": "Melee Final Blows - Average",
						"Value": "0.14",
						"Icon": ""
					},
					...
				]
			},
			{
				"Name": "Deaths",
				"Icon": "<use xlink:href=\"https://overwatch-a.akamaihd.net/img/icons/career-icons-9f59a643a3181e5bb684871c67ae3b62b6476ddeebd073310fe61baf6de3322ebeb80ec0e1f31a6817d164c6b2856c9f1830a3eeb11fb2c1d119a11dfba17437.svg#deaths\"></use>",
				"Stats": [
					{
						"Name": "Deaths",
						"Value": "4,459",
						"Icon": ""
					},
					...
				]
			},
			{
				"Name": "Match Awards",
				"Icon": "<use xlink:href=\"https://overwatch-a.akamaihd.net/img/icons/career-icons-9f59a643a3181e5bb684871c67ae3b62b6476ddeebd073310fe61baf6de3322ebeb80ec0e1f31a6817d164c6b2856c9f1830a3eeb11fb2c1d119a11dfba17437.svg#matchawards\"></use>",
				"Stats": [
					{
						"Name": "Cards",
						"Value": "195",
						"Icon": ""
					},
					...
				]
			},
			{
				"Name": "Game",
				"Icon": "<use xlink:href=\"https://overwatch-a.akamaihd.net/img/icons/career-icons-9f59a643a3181e5bb684871c67ae3b62b6476ddeebd073310fe61baf6de3322ebeb80ec0e1f31a6817d164c6b2856c9f1830a3eeb11fb2c1d119a11dfba17437.svg#game\"></use>",
				"Stats": [
					{
						"Name": "Games Won",
						"Value": "353",
						"Icon": ""
					},
					...
				]
			},
			{
				"Name": "Miscellaneous",
				"Icon": "<use xlink:href=\"https://overwatch-a.akamaihd.net/img/icons/career-icons-9f59a643a3181e5bb684871c67ae3b62b6476ddeebd073310fe61baf6de3322ebeb80ec0e1f31a6817d164c6b2856c9f1830a3eeb11fb2c1d119a11dfba17437.svg#misc\"></use>",
				"Stats": [
					{
						"Name": "Melee Final Blows - Most in Game",
						"Value": "3",
						"Icon": ""
					},
					...
				]
			},
			{
				"Name": "Hero Specific",
				"Icon": "<use xlink:href=\"https://overwatch-a.akamaihd.net/img/icons/career-icons-9f59a643a3181e5bb684871c67ae3b62b6476ddeebd073310fe61baf6de3322ebeb80ec0e1f31a6817d164c6b2856c9f1830a3eeb11fb2c1d119a11dfba17437.svg#0x02E0000000000002\"></use>",
				"Stats": [
					{
						"Name": "Souls Consumed",
						"Value": "376",
						"Icon": ""
					},
					...
				]
			},
			{
				"Name": "Combat",
				"Icon": "<use xlink:href=\"https://overwatch-a.akamaihd.net/img/icons/career-icons-9f59a643a3181e5bb684871c67ae3b62b6476ddeebd073310fe61baf6de3322ebeb80ec0e1f31a6817d164c6b2856c9f1830a3eeb11fb2c1d119a11dfba17437.svg#combat\"></use>",
				"Stats": [
					{
						"Name": "Eliminations",
						"Value": "539",
						"Icon": ""
					},
					...
				]
			},
			{
				"Name": "Assists",
				"Icon": "<use xlink:href=\"https://overwatch-a.akamaihd.net/img/icons/career-icons-9f59a643a3181e5bb684871c67ae3b62b6476ddeebd073310fe61baf6de3322ebeb80ec0e1f31a6817d164c6b2856c9f1830a3eeb11fb2c1d119a11dfba17437.svg#assists\"></use>",
				"Stats": [
					{
						"Name": "Healing Done",
						"Value": "13,582",
						"Icon": ""
					},
					...
				]
			},
			{
				"Name": "Best",
				"Icon": "<use xlink:href=\"https://overwatch-a.akamaihd.net/img/icons/career-icons-9f59a643a3181e5bb684871c67ae3b62b6476ddeebd073310fe61baf6de3322ebeb80ec0e1f31a6817d164c6b2856c9f1830a3eeb11fb2c1d119a11dfba17437.svg#best\"></use>",
				"Stats": [
					{
						"Name": "Eliminations - Most in Life",
						"Value": "13",
						"Icon": ""
					},
					...
				]
			},
			{
				"Name": "Average",
				"Icon": "<use xlink:href=\"https://overwatch-a.akamaihd.net/img/icons/career-icons-9f59a643a3181e5bb684871c67ae3b62b6476ddeebd073310fe61baf6de3322ebeb80ec0e1f31a6817d164c6b2856c9f1830a3eeb11fb2c1d119a11dfba17437.svg#average\"></use>",
				"Stats": [
					{
						"Name": "Melee Final Blows - Average",
						"Value": "0.13",
						"Icon": ""
					},
					...
				]
			},
	    ],
	    "HeroMetrics": [
			{
				"Name": "Time Played",
				"Heroes": [
					{
						"Name": "Roadhog",
						"Value": "2 hours",
						"Image": "https://blzgdapipro-a.akamaihd.net/game/heroes/small/0x02E0000000000040.png",
						"Percent": "1"
					},
					...
				]
			}
			...
		]
	}
}
```