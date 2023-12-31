CREATE TABLE IF NOT EXISTS timetable (
                                               id SERIAL PRIMARY KEY,
                                               weekday VARCHAR(20),
                                               schedule TEXT
);

INSERT INTO timetable (weekday, schedule)
VALUES
('ПН', 'Диф. ур. (лаб) 511\nТер. вер. (лаб) 511\nВеб (лаб) 609'),
('ВТ', 'Мат. анализ (лекция) 602\nМат. комп. пр. (лаб) 600\nВеб (лаб) 606'),
('СР', 'Диф. ур. (лекция) 602\nСети (лаб) 609\nМат. анализ (лаб) 609\nФиз-ра'),
('ЧТ', 'Сети (лаб) 602\nИностранный язык (практ.) 604\nКультурология (лекция) 602'),
('ПТ', 'Тер.вер (лекция) 602\nМат. комп. пр. (лаб) 600\nПрограммирование (лаб) 605\nКультурология (практ) 602');