### Golang register with PostgreSQL

# Register structure
1. Ro'yxatdan o'tish
2. show - userlarni ko'rish
3. get by user id
4. SignIn
5. update user f_name, parol
6. delete user - id

# Register 
- f_name
- lastName
- job
- age
- parol

# Golang orqali databasedan ma'lumot o'qish 
1 - query yoziladi M: `SELECT name, lastName ... from databaseName ...`
2 - shu queryni database bilan bog'lash uchun dbga so'rov jo'natiladi db.Query(query) orqali bu 2 ta qiymat qaytaradi birinchisi row ikkinchisi err, row bu databasedagi har bir row qatorlarning ponteri
3 - kiyin har bir rowni aylanib chiqiash kerak buladi row.Next() orqali va ma'lumotlarni row.Scan() orqali o'qib olinadi
4 - Va oxirida har diom ochilgan databaseni yopib ketish kerak buladi defer row.Close() orqali

[To'liqroq ma'lumot uchun link](http://go-database-sql.org/retrieving.html)