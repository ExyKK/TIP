# 09.09 Практика №1
# Пащенко Е.М. ЭФМО-02-24
## Интернет-магазин электронных приборов и компонентов

### Логическая модель базы данных
![ERD](https://github.com/user-attachments/assets/cc522390-7384-47b6-bca3-f5179b3f0c50)

### Физическая модель базы данных
![ERD_typed](https://github.com/user-attachments/assets/78938313-577e-47b3-ab1a-ab476aeba5ce)

### Сущности ERD
1. **products (товары)** — информация о товарах, хранит его категорию, производителя и поставщика, изображения, артикулы, название, описания, стоимость и статус 
2. **reviews (отзывы)** — отзывы о товаре, включает пользователя и сам товар, заголовок, текст отзыва, дату создания и пятибалльный рейтинг
3. **categories (категории)** — категории товаров, содержит наименование категории
4. **images (изображения)** — изображения товаров, хранит 5 ссылок (путей) на изображения  
5. **suppliers (поставщики)** — информация о поставщиках, включает наименование и подробности
6. **manufacturers (производители)** — информация о производителях, включает наименование и подробности
7. **warehouses (склады)** — информация о складах, включает название, адрес, статус и подробности
8. **warehouses_products (склады-товары)** — представляет связь многие-ко-многим между складами и товарами, указывает количество товара на складе
9. **users (пользователи)** — информация о пользователе магазина, включает его логин и пароль, имя, фамилию, отчество, а также email и номер телефона
10. **carts (корзины)** — представляет корзину, в которую пользователь может добавлять товары, хранит дату её создания
11. **carts_products (корзины-товары)** — представляет связь многие-ко-многим между корзинами и товарами, указывает количество товара в корзине
12. **shipping_addresses (адреса доставок)** — адреса доставок, которые пользователь может указать в личном кабинете и которые содержатся в заказе
13. **orders (заказы)** — представляет заказ, включая пользователя, адрес доставки, код заказа, время создания, ожидаемое время доставки, статус и детали заказа
14. **orders_products (заказы-товары)** — представляет связь многие-ко-многим между заказами и товарами, указывает количество товара в заказе
15. **payments (платежи)** — информация о платежах, включает номер заказа, способ оплаты, сумму, время оплаты, статус