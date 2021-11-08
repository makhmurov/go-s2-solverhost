# Напиши свой микросвервис на GO за 3 дня! WIP

Домашнее задание "Напиши свой микросвервис на GO за 3 дня!" @ IT Univer, 2021.10


Необходимо предоставить решение четырёх задач:

- "Циклическая ротация"
- "Чудные вхождения в массив"
- "Проверка последовательности"
- "Поиск отсутствующего элемента"

---

## Часть 1. Реализация задач

Реализации представлены в пакете [task](https://github.com/makhmurov/go-s2-solverhost/tree/dev/task):

- ["Циклическая ротация"](https://github.com/makhmurov/go-s2-solverhost/tree/dev/task/rotation)
- ["Чудные вхождения в массив"](https://github.com/makhmurov/go-s2-solverhost/tree/dev/task/unpaired)
- ["Проверка последовательности"](https://github.com/makhmurov/go-s2-solverhost/tree/dev/task/sequence)
- ["Поиск отсутствующего элемента"](https://github.com/makhmurov/go-s2-solverhost/tree/dev/task/missed)

## Часть 2. Проверка заданий

Приложение выполняет проверку корректоности выполнения заданий, обращаясь к внешнему сервису для получения наборов данных и верификации результатов.

Для работы программы требуется установить следуюшие переменные окружения:

Перменная | Значение 
-|-
USER_NAME   | Никнейм пользователя бота
DATASET_URL | URL поставщика наборов данных, без имени задачи, к примеру: http://example.com:8080/tasks/
VERIFY_URL  | URL сервиса проверки результатов, к примеру: https://example.com:8443/tasks/solution/

> Notice:
> Work In Progress.
