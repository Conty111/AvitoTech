# Структура приложения
    app/
        source/
            logging/
                logging.go - пакет логирования
            
            storage/
                redisGO/
                    redisGo.go - пакет с реализацией Storage 
                    с помощью redis
                storage.go - файл с интерфейсом взаимодействия с БД
            
            web/
                gorilla_mux/
                    gorilla_mux.go - файл с http-роутером gorilla/mux, 
                    methods.go - файл с обработчиками запросов
                web.go - файл с интерфейсом для API
            
            config.go - статичные данные для приложения
            go.mod - зависимости
            main.go - точка входа приложения
        
        Dockerfile - файл для сборки docker образа
        go1.21.3.linux-amd64.tar.gz - архив с Golang-ом

# Как запускать
У вас должен быть установлен docker-compose.

    sudo docker-compose up -d

# Как использовать
Nginx будет принимать запросы на 8089 порту

Set value

    curl -X POST -H "Content-Type: application/json" -d '{ "name": "Dmitriy"}' http://localhost:8089/set_key

Get value

    curl -X GET http://localhost:8089/get_key?key=name

Delete value

    curl -X DELETE -H "Content-Type: application/json" -d '{"name": "Dmitriy"}' http://localhost:8089/del_key