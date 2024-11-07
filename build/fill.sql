INSERT INTO forecasts (forecast_id, title, short, descr, color, measure_type, extended_desc, img_url) VALUES
    (1, 'Прогноз температуры', 'Температура', 
    'Предскажем температуру посредством применения метода скользящего среднего','255, 195, 182, 1','градусы цельсия, °C',
    'Нахождение вероятных значений средних температур на последующие дни с учетом тренда изменения средней за скользящее окно дней температуры.', 
    'http://127.0.0.1:9000/test/image-1.png'),
    (2, 'Предсказать давление', 'Давление', 
    'Покажем в мм рт. ст. наиболее вероятного значения атмосферного давления','213, 206, 255, 1','миллиметры ртутного столба, мм рт. ст.',
    'Нахождение на последующие дни наиболее вероятного диапазона давлений с учетом тренда. Следует учитывать вероятность зацикливания предсказаний в следствие ограниченности диапазона давлений.', 
    'http://127.0.0.1:9000/test/image-2.png'),
    (3, 'Предугадать влажность', 'Влажность', 
    'Подскажем, как одеться по влажности атмосферного воздуха, в процентах','223, 229, 255, 1','проценты, %',
    'Нахождение статистически вероятных значений влажности на последующие дни с учетом тренда изменения средних значений влажности.
		Стоит отметить, что метод будет работать исправно лишь при использовании больших значений величины скользящего окна:
		в долгосрочной перспективе учитываются макроизменения, которые зачастую вбирают в себя все микроизменения, чего, очевидно, не происходит в короткосрочных измерениях.
		В последних любые экстримальные значения, обусловленные осадками или аномалиями, значительно снизят точность предсказания.', 
    'http://127.0.0.1:9000/test/image-3.png');

    INSERT INTO users (login, password, role) VALUES
    ('borizzler','123',3)