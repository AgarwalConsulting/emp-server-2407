# ReSTful API (Representational State Transfer)

CRUD (Create, Read, Update, Destroy) => HTTP Methods (GET, POST, PATCH, OPTIONS, ...)

## Employee Management Server (JSON API)

```
CRUD      Action          HTTP Method         URL                    Req Body              Resp Body
-------------------------------------------------------------------------------------------------------------
Read      Index           GET               /employees                  -                   [{...}, ...]
Read      Show            GET               /employees/{id}             -                   {...}
Create    Create          POST              /employees                 {...}                {id: , ...}
Update    Update          PUT               /employees/{id}            {...}                {...} / -
Update    Update          PATCH             /employees/{id}           some attrs            {...} / -
Destroy   Delete          DELETE            /employees/{id}             -                   {...} / -
```
