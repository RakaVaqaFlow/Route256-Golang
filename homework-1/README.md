# Домашнее задание

## Необходимо выполнить два задания:

### Необходимо написать Go функцию, осуществляющую примитивную распаковку строки, содержащую повторяющиеся символы/руны, например:
* "a4bc2d5e" => "aaaabccddddde"
* "abcd" => "abcd"
* "3abc" => "" (некорректная строка)
* "45" => "" (некорректная строка)
* "aaa10b" => "" (некорректная строка)
* "aaa0b" => "aab"
* "" => ""
* "d\n5abc" => "d\n\n\n\n\nabc"

Как видно из примеров, разрешено использование цифр, но не чисел.

В случае, если была передана некорректная строка, функция должна возвращать ошибку.
При необходимости можно выделять дополнительные функции / ошибки.

**(💎) Дополнительное задание: поддержка экранирования через `\`:**
* "qwe\4\5" => "qwe45"
* "qwe\45" => "qwe44444"
* "qwe\\\5a" => "qwe\\\\\\\\\\a"
* \`qw\ne\` (внимание на косые кавычки) => "" (некорректная строка)


### Необходимо написать Go функцию, принимающую на вход строку с текстом и возвращающую слайс с 10-ю наиболее часто встречаемыми в тексте словами.

Если слова имеют одинаковую частоту, то должны быть отсортированы **лексиграфически**.

* Словом считается набор символов, разделенных пробельными символами.

* Если есть более 10 самых частотных слов (например 15 разных слов встречаются ровно 133 раза,
  остальные < 100), то следует вернуть 10 лексиграфически первых слов.

* Словоформы не учитываем: "нога", "ногу", "ноги" - это разные слова.

* Слово с большой и маленькой буквы считать за разные слова. "Нога" и "нога" - это разные слова.

* Знаки препинания считать "буквами" слова или отдельными словами.
  "-" (тире) - это отдельное слово. "нога," и "нога" - это разные слова.

#### Пример
```
cat and dog, one dog,two cats and one man
```
Топ 7:
- `and`     (2)
- `one`     (2)
- `cat`     (1)
- `cats`    (1)
- `dog,`    (1)
- `dog,two` (1)
- `man`     (1)

(💎) Дополнительное задание: НЕ учитывать большие/маленькие буквы и знаки препинания:

"Нога" и "нога" - это одинаковые слова, "нога!", "нога", "нога," и " 'нога' " - это одинаковые слова;
"какой-то" и "какойто" - это разные слова, "-" (тире) - это не слово.