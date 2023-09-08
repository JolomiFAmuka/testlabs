import psycopg2

#connect to the db
con = psycopg2.connect(
    host = "localhost",
    database = "dvdrental",
    user = "postgres",
    password = "postgres"
)

#cursor
cur = con.cursor()

#execute query
cur.execute("insert into actor(actor_id, first_name, last_name) values(%s, %s, %s)",(210, "Jolomski", "Amuka"))

cur.execute("select actor_id, first_name from actor")

rows = cur.fetchall()

for r in rows:
    print(f"id {r[0]} name {r[1]}")

#commit the transaction
con.commit()

#close cursor
cur.close()

#close the connection
con.close()