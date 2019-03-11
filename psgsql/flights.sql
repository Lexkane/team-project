CREATE TABLE trips_flights (
flight_id uuid REFERENCES flights(id) ON DELETE CASCADE,
trip_id uuid REFERENCES trips(trip_id) ON DELETE CASCADE,
id uuid DEFAULT uuid_generate_v1(),
PRIMARY KEY (id)
);

CREATE TABLE  flights (
id uuid DEFAULT uuid_generate_v1(),
departure_city VARCHAR (30) NOT NULL,
departure_time TIME,
departure_date DATE,
arrival_city VARCHAR (30) NOT NULL,
arrival_time TIME,
arrival_date DATE,
price INT NOT NULL,
PRIMARY KEY (id)
);

INSERT INTO flights
  (id,  departure_city, departure_time, departure_date, arrival_city, arrival_time, arrival_date, price)
  VALUES (uuid_generate_v1(),'Kyiv', '09:05:10', '2019-03-02', 'Lviv','10:12:22', '2019-03-02', '400');
INSERT INTO flights
  (id,  departure_city, departure_time, departure_date, arrival_city, arrival_time, arrival_date, price)
  VALUES (uuid_generate_v1(),'Kyiv', '09:10:16', '2019-03-04', 'Warsaw','11:10:16', '2019-03-04', '800');
INSERT INTO flights
  (id,  departure_city, departure_time, departure_date, arrival_city, arrival_time, arrival_date, price)
  VALUES (uuid_generate_v1(),'Kyiv', '09:30:05', '2018-03-05', 'Berlin','13:11:12', '2018-03-05', '700');
INSERT INTO flights
  (id,  departure_city, departure_time, departure_date, arrival_city, arrival_time, arrival_date, price)
  VALUES (uuid_generate_v1(),'Kyiv', '09:40:18', '2018-03-07', 'London','15:10:10', '2018-03-07', '1100');
