kode peserta 	: 149368582101-633
nama lengkap 	: Asep Saefuddin
Summary sesi 7
      ~ menggunakan postgresql :
      		1. Instalasi package postgresql dengan go get -u github.com/lib/pq
          2. Create database, create table
          3. connecting golang ke database
          4. CRUD data employee
      ~  Gorm adalah ORM (Object Relation Mapping) yang cukup populer untuk bahasa Go, 
      yang dimana Gorm telah menyediakan berbagai macam fitur seperti auto migration, eager loading, association, query method, dan lain-lain.
      menggunakan :
      		1. Declaring models : pembuatan model tiap tabel
          2. Association/relation : mengatur tipe data, foreign key, primary key pada tiap tabel
          3. Connecting To Database And Table Migration : materi ini membahas autoMigrate sehingga tabel akan langsung terbuat dengan sendirinya
          4. CRUD user dan product
