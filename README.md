# NewsGoClient

Клиент для получения новостей по RSS ссылке

Поддерживает запуск/остановку работы сервера по чтению новостей, получение заголовков и новостей по подстроке заголовка

1) Запустить сервер "go run main" (из директории сервера сопуствующего так и отдельно клонированного)
2) Запустить клиент "go run main" (из корня репозитория)

Примеры RSS ссылок:
 - https://lenta.ru/rss
 - https://meduza.io/rss/all
 
Примеры URL ссылок
 - https://yandex.ru/news (+ css правило article)
 - https://ria.ru (+ css правило .cell-list__item)
 - http://lenta.ru (+ css правило .item)
 
 P.S.
Наблюдаются шероховатости в работе и обработке ошибок, https://meduza.io как url не работает (get-запрос возвращает неполную? страницу)
