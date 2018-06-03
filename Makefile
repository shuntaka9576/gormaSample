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

# goa
REPO:=github.com/shuntaka9576/gormaSample
generate:
	@goagen app     -d $(REPO)/design
	@goagen swagger -d $(REPO)/design
	@goagen client -d $(REPO)/design
	@goagen js -d $(REPO)/design
	@goagen schema -d $(REPO)/design
	@go build -o build

clean:
	@rm -rf app
	@rm -rf client
	@rm -rf tool
	@rm -rf swagger
	@rm -rf schema
	@rm -rf js
	@rm -f build