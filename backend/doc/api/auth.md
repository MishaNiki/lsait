# API
## AUTH

#### Signup
---
##### __POST__ : /auth/signup
___request:___
```
{
  "email": String
  "password": String
}
```
___response:___
__OK:__
201 - Все прошло гладко, на почту придёт сообщение со ссылкой на подтверждение
__Error:__
400 - Ошибка в json;
401 - Такой email уже есть в базе данных
406 - Данные не валидны, например  длина пароля

---
##### __POST__: /auth/signup/confirm
___request:___
```
{
  "key": String
}
```
___response:___
__OK:__
200 - Все прошло гладко

```
{
  "accessToken": String
  "refrashToken": String
}
```
__Error:__
400 - Ошибка в json
401 - Такого ключа нет либо срок его действия истёк

---

### Login
---
##### __POST__ : /auth/signup
___request:___
```
{
  "email": String
  "password": String
}
```
___response:___
__OK:__
200 - Все прошло гладко
```
{
  "accessToken": String
  "refrashToken": String
}
```
__Error:__
400 - Ошибка в json;
401 - Ошибка авторизации нет такого email или неверный пароль

---

### Forgot

##### __POST__ : /auth/forgot
___response:___
```
{
  "email": String
}
```
___responce:___
__OK:__
200 - Все прошло гладко, на почту придёт сообщение со ссылкой на подтверждение
__Error:__:
400 - Ошибка в json;
401 - Такой email нет в базе данных

---
##### __POST__: /auth/forgot/confirm
___request:___
```
{
  "key": String
  "password": String
}
```
___response:___
__OK:__
200 - Все прошло гладко
```
{
  "accessToken": String
  "refrashToken": String
}
```
__Error:__
400 - Ошибка в json
401 - Такого ключа нет либо срок его действия истёк

---

### Refrash
---
##### __POST__: /auth/refrash
___request:___
```
{
  "refrashToken": String
}
```
___response:___
__OK:__
200 - Все прошло гладко
```
{
  "accessToken": String
  "refrashToken": String
}
```
__Error:__
400 - Ошибка в json
401 - такого токена нет, либо токеном кто-то воспользовался 

---

### Logout
---
##### __POST__: /auth/logout
___request:___
```
{
  "refrashToken": String
}
```
___response:___
__OK:__ 
200 - Все прошло гладко

__Error:__
400 - Ошибка в json
401 - такого токена нет, либо токеном кто-то воспользовался

---
