# Pseudorandom distribution utility
В общем случае используется для распределения порядковых номеров(начиная с 0) 
между объектами построчно перечисленными в текстовом файле.

Данный пример - распределение билетов между студентами.
Список студентов находится в файле `students.txt`

Распоеделение можно менять изменяя параметр распределения и меняя имена студентов в списке.
Распреление является детерминированным, то есть одинаковым для одних и тех же пар список-параметр.

# Сборка и запуск:
**Примечание:** Инструкция написана под Ubuntu 18.04

# Склонируйте репозиторий:
`git clone https://github.com/AndreyBMWX6/prand-distr.git`
# Сборка:
`go build ./main.go`
# Запуск:
`./main --file <filename(string)> --numbilets <uint> --parameter <int>`

**Параметры:** имя файла с ФИО студентов, число билетов, параметр, меняющий распределение.

# Пример запуска:
`./main --file students.txt --numbilets 20 --parameter 42`

Результат работы:
```
Иванов Иван Иванович 12
Ярцев Ярослав Ярославович 10
Бутусов Валерий Викторович 7
Петров Петр Петрович 19
```