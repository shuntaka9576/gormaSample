# mysql init
MYSQLUSER:=root
MYSQLPASSWORD:=Passw0rd
CONTAINERNAME=container_gormaMysqlService_1
mysql/init:
	docker cp createdb $(CONTAINERNAME):/
	-docker exec -it $(CONTAINERNAME) mysql -u $(MYSQLUSER) -p$(MYSQLPASSWORD) -e "source createdb/createdb.sql"
	docker exec -it $(CONTAINERNAME) mysql -u $(MYSQLUSER) -p$(MYSQLPASSWORD) -e "show databases;"
mysql/show:
	docker exec -it $(CONTAINERNAME) mysql -u $(MYSQLUSER) -p$(MYSQLPASSWORD) -e "show databases;use celler;show tables;show variables like 'char%';"
mysql/login:
	docker exec -it $(CONTAINERNAME) mysql -u $(MYSQLUSER) -p$(MYSQLPASSWORD)

# migration
DBNAME:=celler
ENV:=development
migrate/up:
	sql-migrate up -env=$(ENV)
migrate/down:
	sql-migrate down -env=$(ENV)
