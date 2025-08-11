# 07.08.2025
примеры запросов
создание задачи
curl -X POST http://localhost:8080/api/v1/tasks

добавление URL в задачу
curl -X POST http://localhost:8080/api/v1/target-to-task/1 \
     -H "Content-Type: application/json" \
     -d '{
  "urls": [
    "https://zagorie.ru/upload/iblock/4ea/4eae10bf98dde4f7356ebef161d365d5.pdf",
    "https://cs13.pikabu.ru/post_img/big/2019/06/27/5/156161920716813977.jpg",
    "https://cs13.pikabu.ru/post_img/big/2019/06/27/5/156161920716813977.jpg"
  ]
}'

получение статуса задачи по id
curl -X GET http://localhost:8080/api/v1/task-status/1

так же можно кидать запросы через postman, коллекция в папке postman_collection