# Backend для сервиса soloanvill

api сервис для взаимодействия с [front'a](https://github.com/Anvill1/soloanvill.ru) с другими системами

## Необходимое для запуска

Для запуска приложения необходимо подключение к postgresql и jenkins

Обязательные переменные окружения для запуска приложения:

- SOLOANVILL_DATABASE_HOST
- SOLOANVILL_DATABASE_PORT
- SOLOANVILL_DATABASE_NAME (определна в Dockerifle, можно переопределить)
- SOLOANVILL_DATABASE_USER
- SOLOANVILL_DATABASE_PASSWORD
- SOLOANVILL_JENKINS_HOST
- SOLOANVILL_JENKINS_LOGIN
- SOLOANVILL_JENKINS_TOKEN
