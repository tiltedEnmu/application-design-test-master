Задание выполнил.

Как потрогать:
- Клонируем репо
- Затем в папке открываем терминал и пишем: `go run ./cmd/main.go`
- Потом пишем запросы в том же терминале через curl или как вам удобно
- Эндпоинтов две штуки: `POST /order` - резервирует румы согласно заказу. 
Заказ оформляем в json по схеме из задания и кладем в тело запроса. 
И `GET /rooms` - в задании не видел, чтобы нужно было что-то подобное делать, 
но он так или иначе выводит список всех забронированных комнат. 
Тело запроса оставляем пустым. Удобно, чтобы смотреть изменения.

Примечания:
- Тестов нет, конфигов тоже нет (все захардкодил в мейне). Ну в задании не требовалось, я и не сделал
- Если есть какие то вопросы по поводу задания, писать в тг (https://t.me/MesanagiEnmu)