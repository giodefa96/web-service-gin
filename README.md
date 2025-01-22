# 1️⃣ Inizializza il modulo Go
go mod init example/web-service-gin

# 2️⃣ Installa le dipendenze
go get -u github.com/gin-gonic/gin
go get -u gorm.io/gorm
go get -u gorm.io/driver/mysql
go get -u github.com/joho/godotenv
go mod tidy

# 3️⃣ Avvia Docker (costruisce e avvia Go + MySQL)
docker-compose up --build

# 4️⃣ Controlla lo stato dei container
docker ps

# 5️⃣ Test API
curl -X GET http://localhost:8080/users

# 6️⃣ Accedi a MySQL nel container
docker exec -it mysql_container mysql -u root -p

# 7️⃣ Arresta i container
docker-compose down


go install github.com/go-delve/delve/cmd/dlv@latest
