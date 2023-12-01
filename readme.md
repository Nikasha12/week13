#Setup Oracle SQL in SQL Developer

sql.Open("goracle", "SYS AS SYSDBA/oracle@localhost:1521/ORCLPDB")


#create table
CREATE TABLE time_log (
    id NUMBER GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    timestamp TIMESTAMP
);



# install oracle driver

go get -u gopkg.in/goracle.v2
