##БОКСИГЕН - генератор SVG файлов раскроек коробок для лазерной резки.
- [ ] - получаем габаритные размеры - (sW, sH, sD) внешние, длину шипа (Sl - по умолчанию 20) выборку под палец (tR - (0 - не нужно, >0 - радиус)) и толщину материала (mT).
- [ ] - если габариты какого либо элемента меньше 2 х mT, сообщаем пользователю, отправляемся в 1.
- [ ] - Формируем элементы дно, длинная стенка х 2, торцевая стенка, торцевая стенка с пальцем.
- [ ] - В каждом элементе рассчитываем количество шипов по всем соединяемым сторонам. Формируем массив вершин.