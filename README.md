
# Name:  Daniel Hartanto

<img src="ERD.jpg" alt="database-design" >

1. **User Table**:

   - table ini untuk pengguna sistem seperti admin
   - terdiri dari id user, uuid untuk menjadi token,nama user,password, alamat,nomor hp, jabatan

2. **Product Table**:
   
   - table ini berisi produk yang dijual beserta harga dan stocknya.
   - terdiri dari id produk, nama produk, tanggal perubahan,updated by, harga produk,  stock/ jumlah produk 
   - tanggal perubahan digunakan untuk mengetahui kapan stock terakhir kali ditambahkan
   - updated by digunakan untuk penanggung jawab siapa yang telah menambahkan stock terakhir kali
   - stock/jumlah produk digunakan untuk mengetahui sisa produk agar tidak kekurangan saat penjualan
   - stock auto berkurang saat customer membeli produk tersebut menggunakan trigger
     
3. **Customer Table**:

   - tabel ini berisikan data diri customer.
   - terdiri dari id customer,nama customer, alamat, nomor hp

4. **Sold Product Table**:
   - table ini digunakan untuk menyimpan data setiap pembelian produk terjadi
   - table ini berguna untuk melihat penjualan harian maupun penjualan per user maupun penjualan per produk 
   - terdiri dari id user foreign key table user, id customer foreign key table customer,id produk foreign key table produk, waktu penjualan, jumlah penjualan



6. **Relationships**:
   - The 'user_id' field in the 'sale' table references the 'id' field in the 'user' table, linking each sale to a specific user or cashier.
   - The 'customer_id' field in the 'sale' table references the 'id' field in the 'customer' table, connecting each sale to a specific customer.
   - The 'sale_id' field in the 'sale_item' table references the 'id' field in the 'sale' table, associating each sale item with a specific sale.
   - The 'product_id' field in the 'sale_item' table references the 'id' field in the 'product' table, indicating which product is being sold.

The design ensures that information about users, customers, products, and sales is structured logically and can be efficiently retrieved and managed. The system allows for tracking sales, managing stock, and providing insights into sales and customer data.
