# Goverwatch
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
DEscription: Displays in-detail metadata about a specific hero<br/>
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

