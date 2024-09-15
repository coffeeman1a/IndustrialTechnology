# IndustrialTechnology
# 09.09 Практика 1 Родионов ЭФМО-02-20
Интернет-магазин электроники

Описание сущностей ERD:

1. **User (Пользователь)**: Представляет зарегистрированного пользователя интернет-магазина, который может делать заказы. Содержит информацию о логине, пароле, активности, а также связи с юридическим или физическим лицом.

2. **LegalEntity (Юридическое лицо)**: Представляет юридическое лицо, содержащее ИНН, КПП, официальное название и юридический адрес.

3. **Individual (Физическое лицо)**: Представляет физическое лицо, содержащее информацию о фамилии, имени, отчества, адресе электронной почты и номере телефона.

4. **IllegalContact (Контактное лицо)**: Представляет контактное лицо, связанное с физическим или юридическим лицом, которое может быть незаконным (например, для целей сохранения анонимности).

5. **Order (Заказ)**: Представляет заказ пользователя, включая его статус, комментарии и дату создания.

6. **Payment (Оплата)**: Представляет информацию об оплате заказа, включая статус платежа, банковские реквизиты и дату оплаты.

7. **OrderItem (Товар в заказе)**: Представляет конкретные товары, включенные в заказ, с информацией о количестве и цене за единицу.

8. **Cart (Корзина)**: Представляет корзину пользователя, где он может временно хранить товары до оформления заказа.

9. **CartItem (Товар в корзине)**: Представляет товары, которые пользователь добавил в свою корзину с указанием количества и цены за единицу.

10. **Product (Товар)**: Представляет продукт в интернет-магазине, содержащий название, цену, код продукта, превью изображения, а также информацию о производителе, категории и поставщике.

11. **Category (Категория)**: Представляет категорию товаров в магазине (например, электроника), с указанием кода и названия категории.

12. **Supplier (Поставщик)**: Представляет поставщика товаров с информацией о его названии и описанием.

13. **Manufacturer (Производитель)**: Представляет производителя товаров с информацией о названии и описанием.

14. **Favorite (Избранное)**: Представляет избранные товары пользователей, где они могут сохранять понравившиеся продукты.

15. **RemainingStock (Остатки на складе)**: Представляет оставшееся количество товаров на складе, с привязкой к конкретному магазину.

16. **Shop (Магазин)**: Представляет физический магазин, где хранятся товары, включая его название, адрес и часы работы.

Эти сущности совместно описывают функциональность интернет-магазина, обеспечивая хранение и обработку информации о пользователях, заказах, товарах и их поставщиках.

## Логическая модель базы данных
![image](https://github.com/user-attachments/assets/505d0dad-d92d-4393-a184-3a2802de0a3f)

## Физическая модель базы данных
![image](https://github.com/user-attachments/assets/bc973299-095f-4349-9041-7db5df7f7fb9)



