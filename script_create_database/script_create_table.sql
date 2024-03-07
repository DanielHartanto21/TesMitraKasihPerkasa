-- Database: TesMKP

-- DROP DATABASE IF EXISTS "TesMKP"; 

-- CREATE DATABASE "TesMKP"
--     WITH
--     OWNER = postgres
--     ENCODING = 'UTF8'
--     LC_COLLATE = 'English_United States.1252'
--     LC_CTYPE = 'English_United States.1252'
--     LOCALE_PROVIDER = 'libc'
--     TABLESPACE = pg_default
--     CONNECTION LIMIT = -1
--     IS_TEMPLATE = False;
DROP table IF EXISTS Users CASCADE;
DROP table IF EXISTS Customer CASCADE;
DROP table IF EXISTS Product CASCADE;
DROP table IF EXISTS SoldProduct;
/*creating tables*/
create table Users(
	UserID serial primary key,	
	Uuid varchar(256) not null unique,
	NamaUser varchar(256) not null unique,
	Password varchar(256) not null,
	Alamat varchar(256) not null,
	NoHp varchar(15) not null,
	Jabatan varchar(256) not null
);
create table Customer(
	CustomerID serial primary key,
	NamaCustomer varchar(256) not null unique,
	Alamat varchar(256),
	NoHp varchar(15) not null
);
create table Product(
	ProductID serial primary key,
	NamaProduk varchar(256) not null unique,
	JenisProduk varchar(256) not null,
	HargaProduk Float,
	TanggalUpdate timestamp default CURRENT_TIMESTAMP,
	UpdatedBy varchar(256),
	stock int
);
create Table SoldProduct(
	SoldProductID serial primary key,
	CustomerID serial,
	UserID serial,
	ProductID serial,
	WaktuPenjualan timestamp default CURRENT_TIMESTAMP,
	JumlahPenjualan int,
	foreign key (CustomerID) references Customer(CustomerID) on delete cascade,
	foreign key (UserID) references Users(UserID) on delete cascade,
	foreign key (ProductID) references Product(ProductID) on delete cascade 
);
/* membuat trigger agar bisa langsung stock berkurang sebelum memasukan menjual*/
create or replace function before_insert_trigger_function()
returns trigger as $$
begin
    
    if NEW.JumlahPenjualan > (select Stock from Product where Product.ProductID = NEW.ProductID) then
        raise exception 'Stok tidak mencukupi untuk transaksi ini';
    end if;
	update Product set Stock=Stock-NEW.JumlahPenjualan, TanggalUpdate=CURRENT_TIMESTAMP  where Product.ProductID = NEW.ProductID;

    return NEW;
end;
$$ language 'plpgsql';

CREATE TRIGGER before_insert_trigger
BEFORE INSERT ON SoldProduct
FOR EACH ROW
execute FUNCTION before_insert_trigger_function();
/* insert sample */ 
insert into Users (Uuid,NamaUser, Password, Alamat, NoHp, Jabatan) values
('b9645646-537f-4f4d-878b-15386088c01c','Daniel', '123', 'jl. kapten mulyadi 277', '012345678', 'Admin'),
('52614aac-06c3-43ff-9645-480a131c0c54','Daniel Hartanto', '123', 'jl. kapten mulyadi 277', '012345678', 'Admin');
insert into Customer (NamaCustomer, Alamat, NoHp) values
('Customer 1', 'jl. kapten mulyadi 277', '012345678'),
('Customer 2', 'jl. kapten mulyadi 277', '012345678');
insert into Product (NamaProduk, JenisProduk, HargaProduk, UpdatedBy, Stock) values
('tiket mobil', 'tiket parkir', 5000.00, 'Daniel', 500),
('tiket ekspo', 'tiket masuk', 75000.00, 'Daniel', 2000);
insert into SoldProduct (CustomerID, UserID, ProductID, JumlahPenjualan) values
(1, 1, 1, 1),
(1, 1, 2, 10),
(2, 2, 2, 5);

/*contoh query untuk memasukan cek stock tersisa*/
select ProductID,NamaProduk,JenisProduk,HargaProduk,stock from Product;

/*contoh query untuk mencek penjualan hari ini*/
select Customer.NamaCustomer,Users.NamaUser,Product.NamaProduk,SoldProduct.JumlahPenjualan,
SoldProduct.WaktuPenjualan,Product.HargaProduk*SoldProduct.JumlahPenjualan as pendapatan
from SoldProduct join Product on Product.ProductID=SoldProduct.ProductID 
join Customer on Customer.CustomerID=SoldProduct.CustomerID
join Users on Users.UserID= SoldProduct.UserID
where date(SoldProduct.WaktuPenjualan)=CURRENT_DATE;