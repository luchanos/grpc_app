# grpc_app

- Создаем файл buf.gen.yanl в котором описываем правила по которым будет собираться наш протобаф файл: язык, 
версия синтаксиса, какие плагины ставим и так далее.
- Создаем файл .proto , в котором описываем все наши типы сообщение, которые будут летать 
на нашем сервере. Там надо указать поля сообщений и место (индекс) этих полей в самом сообщении. Ключ сообщения - message.
- В файле также нужно не забывать указать название пакета, который мы описываем в этом файле.
- Также в файле описываем наши сервисы через ключ service.
- rpc в составе сервиса - это таукой тип запроса. Мы указываем имя метода (например GetUser) и принимаемый аргумент
самого запроса (например GetUserRequest - имеет тип message).

Замечание - во время описаний сервисов надо избегать указаний типов общего вида, потому что
это может привести к проблемам - надо делать под них отдельные структуры (сообщения).

После того, как мы опишем все наши сообщения и сервисы и запустим buf gen, в результате мы получим длинную портянку,
в которой будет содержаться весь код компонентов и для сервиса и для клиента.

Замечание - обязательно нужно прикрутить к нашему buf.gen.yaml файлу использование плагина go-grpc, а то ни клиент, ни
сервер не сгенерятся.

Когда мы получим сгенерившиеся файлы, то тогда мы можем отыскать там интерфейс для сервера и для клиента. Наша задача
состоит в том, чтобы взять эти интерфейсы и реализовать их на своей стороне с конкретными реализациями методов.

После того, как я прописал свою реализацию сервиса я должен прилепить его к имплементации grpc сервера. То есть получается как
в случае с любыми сервисами в составе приложения. Наш сервис - это приложение и мы туда прописываем наши штуки - 
это могут быть клиенты к БД и так далее, но в данном конкретном случае нас интересует только наш grpc сервис.

Для создания grpc сервера мы пользуемся библиотекой grpc.

Получается у нас такая петрушка - grpc библиотека на Go нам нужна для того, чтобы получить некий скелет с командами,
а протобаф и плагины позволяют автоматически генерировать кодовую базу для наших сервисов.

## А теперь про стриминговый сервис
Муть стримингового сервиса, что он передает какие-то значения 1 раз в n единиц времени при сохранении стабильного
соединения с клиентом.

Когда определяем протофайл для стримингового сервиса, то мы должы сделать указание на то, что наш метод возвращает
объект стрима. Протобаф сам поймет, что там надо будет собрать.

Когда хотим, чтобы клиент тоже стримил, тогда нам надо указать, что в качестве аргумента он принимает в себя параметр
стрима.

# Вопросики
- Что означает require_unimplemented_servers=false в yaml файле для buf?


# Общие моменты про Go
Если на меня ругается GoLand при попытке импорта, но всё работает, то можно посмотреть вот 
сюда - https://stackoverflow.com/questions/61845013/package-xxx-is-not-in-goroot-when-building-a-go-project

Ещё если он будет ругаться на undefined, то можно запустить в GoLand как директорию. Тут есть немного подробностей:
https://stackoverflow.com/questions/54388407/command-line-arguments-handling-undefined-error-message

