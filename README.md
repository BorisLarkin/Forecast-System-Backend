# Лабораторная работа №4

- **Цель работы**: Завершение бэкенда для `SPA`
- **Порядок показа**: выполнить метод аутентификации через `swagger` в режиме инкогнито, получить список заявок в `swagger`; из браузера использовать содержимое `куки` из вкладки `Application` или `JWT` из ответа аутентификации для заголовков `cookie` или `authorization` соответственно в остальных запросах через `insomnia`/`postman`. Далее в `insomnia`/`postman` выполнить GET списка заявок: 401/403 для гостя, для создателя только его заявки. Выполнить PUT завершения заявки: для создателя 403 статус, для модератора успех и обновление полей. Выполнить GET списка заявок - для модератора все заявки. Показать содержимое `Redis` через `cmd`, для сессий показать пользователя
- **Контрольные вопросы**: куки, сессия, Redis, JWT, авторизация и аутентификация, SSO, двухфакторная аутентификаци, RSA
- **Sequence диаграмма**: весь набор `HTTP` запросов по бизнес-процессу без БД и нативного приложения: аутентификация, список услуг без черновика, добавление услуги в заявку, еще раз список услуг с черновиком, просмотр черновой заявки, редактирование заявки, формирование заявки, список заявок, завершение модератором из второго фронтенда, список заявок с расчетом. Добавить домены в качестве `Lifeline`, при добавлении сообщений выбирать методы доменов из диаграммы классов, передавать ключевые входные и выходные данные через `arguments` в скобках у `Message`
- **Задание**: Добавление авторизации и `swagger` в веб-сервис, подготовка ТЗ

Реализовать методы бэкенда для `аутентификации` и `регистрации`. Авторизация через хранение сессий и куки. Автозаполнение пользователя в таблице `заявок` при создании новой. Добавить описание методов для `swagger`.

Добавить проверку `Permissions` для методов `модератора`. Без авторизации в `Swagger` должно быть доступно только чтение-получение данных через API, с авторизацией - методы `пользователя`, а для `модератора` доступны все методы.
