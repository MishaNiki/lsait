# API
## ARTICLE
### Section
---
##### __GET__ : /article/section
___response:___
__OK:__
200 - Все прошло гладко
```json
{
    "sections" : [
    	{
            "uuid": String,
    		"name": String,
    		"background": String,
    		"description": String
    	},
        ...
    ]
}
```
__Error:__

---
##### __GET__ : /article/section/{uuidSection}
___response:___
__OK:__
200 - Все прошло гладко
```json
{
	"themes": [
		{
			"uuid": String,
			"name": String,
			"atricles": [
				{
					"id": Integer,
					"name": String,
                    "auth": String,
                    "idAuth": String,
                },
                ...
            ]
      	},
        ...
	]
}
```
__Error:__

404 - нет такого раздела

---
### Theme
---
##### __GET__ : /article/theme/{uuidTheme}

##### ___response:___
__OK:__
200 - Все прошло гладко

```json
{
	"uuid": String,
    "name": String,
    "atricles":[
        {
            "id": Integer,
            "name": String,
            "auth": String,
            "idAuth": String,
        },
        ...
    ]
}
```
__Error:__
404 - нет такой темы

---
### Article
---
#####  __GET__ : /article/article/{uuidArticle}

#####  ___response:___
__OK:__
200 - Все прошло гладко

```json
{
    "id": Integer,
    "text": String,
    "auth": String,
    "idAuth": String,
}
```
__Error:__
404 - нет такой статьи

---
#####  __POST__ : /article/article
___request:___
```json
{
	"accessToken": String,
    "text": String,
    "id": Integer 	/* ID черновика */
}
```
___response:___
__OK:__
200 - Все прошло гладко
__Error:__
400 - плохой запрос
406 - access токен  протух

---
#####  __PUT__ : /article/article
___request:___
```json
{
	"accessToken": String,
	"uuid": String,			/* Идентификатор статьи */
    "text": String
}
```
___response:___
__OK:__
200 - Все прошло гладко
__Error:__
400 - плохой запрос
406 - access токен  протух

---
#####  __DELETE__ : /article/article
___request:___
```json
{
	"accessToken": String,
	"uuid": String
}
```
___response:___
__OK:__
200 - Все прошло гладко
__Error:__
400 - плохой запрос
406 - access токен  протух

---

### Draft

---
##### GET : /article/draft
___request:___
___response:___
__OK:__
200 - Все прошло гладко

```json
{
	"drafts": [
		{
			"name": String,
            "id": Integer
		},
	],
}
```
__Error:__
404 - нет такого черновика

---
##### GET : /article/draft/{id}
___request:___
___response:___
__OK:__
200 - Все прошло гладко

```json
{
    "accessToken": String,
	"text": String
}
```
__Error:__
404 - нет такого черновика

---
##### POST: /article/draft
___request:___
```json
{
	"accessToken": String,
    
}
```
___response:___
__OK:__
200 - Все прошло гладко
__Error:__
404 - нет такого черновика

---
##### PUT: /article/draft
___request:___
```json
{
    "accessToken": String,
    "id": Integer,
	"text": String
}
```
___response:___
__OK:__
201 - Обновлено
__Error:__
404 - нет такого черновика

---
##### DELETE: /article/draft
___request:___
```json
{
    "accessToken": String,
	"id": Integer
}
```
___response:___
__OK:__
200 - Все прошло гладко
__Error:__
400 - плохой запрос
404 - нет такого черновика

---