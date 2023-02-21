FROM mysql:5.7
# copy my.cnf into image because mouting doesn't work with codespaces
COPY my.cnf /etc/mysql/conf.d/my.cnf
RUN chmod 644 /etc/mysql/conf.d/my.cnf  
COPY init.sql /docker-entrypoint-initdb.d/