# API

## ARTICLE

### Section
---
##### __GET__ : /article/section

___response:___

__OK:__

- 200 - Все прошло гладко
```
{
  "sections" : [
    {
      "uuid": String,
      "title": String,
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

- 200 - Все прошло гладко

```
{
  "section" : {
    "uuid": String,
    "title": String,
    "themes": [
	  {
	    "title": String,
		"atricles": [
		  {
			"id": Integer,
			"title": String,
			"description": String,
			"date": String
		  },
		  ...
		]
	  },
	  ...
	]
  }
}
```

__Error:__

- 404 - нет такого раздела

---

##### __GET__: /section/theme

___response:___

__OK:__

- 200 - Все прошло гладко

```
{
  "section" : {
    "uuid": String,
    "title": String,
    "themes": [
	  {
	    "id": Integer,
        "title": String,
	  },
	  ...
	]
  }
}
```

---
### Article

> __Примечение:__ Для того чтобы выложить статью нужно выполнит следующию последовательность операций:
>
> 1. Запрос [POST]: /article/draft -  создаст черновик
> 2. Запрос [GET]: /article/edit - предоставит айди черновика для редактирования
> 3. Зарос [POST]: /article/article - переведёт черновик в состояние статьи

---
#####  __GET__ : /article/article/{id}

___response:___

__OK:__

- 200 - Все прошло гладко

```
{
  "id": Integer,
  "title": String,
  "text": String,
  "description": String,
}
```

__Error:__

- 404 - нет такой статьи

---
#####  __POST__ : /article/article

___request:___

___header:__ Token - access token 

```
{
  "id": Integer, 	/* ID черновика */
  "text": String,
  "title": String,
  "description": String,
  "idTheme": Integer
}
```

___response:___

__OK:__

- 200 - Все прошло гладко

__Error:__

- 400 - плохой запрос
- 401 - неавторизован
- 426 - access токен  протух

---

#####  __PUT__ : /article/article

___request:___

__header:__ Token - access token

```
{
	"id": Integer,
	"text": String,
	"title": String,
	"description": String,
	"idTheme": Integer
}
```
___response:___

__OK:__

- 200 - Все прошло гладко

__Error:__

- 400 - плохой запрос
- 401 - неавторизован
- 426 - access токен  протух

---

#####  __DELETE__ : /article/article

___request:___

___header:___  Token - access token

```
{
	"id": Integer
}
```
___response:___

__OK:__

- 200 - Все прошло гладко

__Error:__

- 400 - плохой запрос
- 401 - неавторизован
- 426 - access токен  протух

---
### Edit
---

#####  __GET__ : /article/edit/{id}

___request:___

___header:___  Token - access token

___response:___

__OK:__

- 200 - Все прошло гладко

```
{
  "id": Integer,
  "text": String,
  "title": String,
  "date": String,
  "description": String,
  "idTheme": Integer
}
```

__Error:__

- 400 - плохой запрос
- 401 - неавторизован
- 426 - access токен  протух

---

### Draft
---

#####  __POST__ : /article/draft

___request:___

___header:___  Token - access token

```
{
  "title": String,
  "description": String
}
```

___response:___

__OK:__

- 200 - Все прошло гладко

__Error:__

- 400 - плохой запрос
- 401 - неавторизован
- 426 - access токен  протух

---
### Profile
---

#####  __GET__ : /article/profile

___request:___

___header:___  Token - access token

___response:___

__OK:__

- 200 - Все прошло гладко

```
{
  "profile": {
    "name": String,
    "surname": String,
    "description": String,
    "position": String,
    "drafts": [
      {
        "title": String,
        "description": String,
        "date": String
      },
      ...
    ],
    "articles": [
      {
        "title": String,
        "description": String,
        "date": String
      },
      ...
    ]
  }
}
```

__Error:__

- 400 - плохой запрос
- 401 - неавторизован
- 426 - access токен  протух

---

#####  __PUT__ : /article/profile

___request:___

___header:___  Token - access token

```
{
  "name": String,
  "surname": String,
  "description": String,
  "position": String
}
```

___response:___

__OK:__

- 200 - Все прошло гладко

__Error:__

- 400 - плохой запрос
- 401 - неавторизован
- 426 - access токен  протух
