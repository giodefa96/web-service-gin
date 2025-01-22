FROM mysql:latest

# Imposta le variabili d'ambiente richieste
ENV MYSQL_ROOT_PASSWORD=yourpassword
ENV MYSQL_DATABASE=mydatabase

# Espone la porta 3306 per le connessioni al database
EXPOSE 3306

# Comando di avvio
CMD ["mysqld"]
