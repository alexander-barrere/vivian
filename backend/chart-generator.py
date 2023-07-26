import psycopg2
from kerykeion import KrInstance, MakeSvgInstance
from datetime import datetime

# Define the database connection details
dbUser = "starfja8_vivian"
dbPass = "PZq(tDO^0NjV"  # Your actual password
dbName = "starfja8_users"
dbHost = "localhost"
dbPort = "5432"

# Connect to the PostgreSQL database
conn = psycopg2.connect(
    dbname=dbName,
    user=dbUser,
    password=dbPass,
    host=dbHost,
    port=dbPort
)

# Create a cursor object
cur = conn.cursor()

# Execute a SQL query to fetch the user data
cur.execute("SELECT first_name, birth_date, birth_time, city FROM profile WHERE id = 42")

# Fetch the result of the query
user = cur.fetchone()

# Parse the birth date and time
birth_date = datetime.strptime(user[1], "%Y-%m-%d")
birth_time = datetime.strptime(user[2], "%H:%M")  # Changed format string

# Create a KrInstance object
kr_instance = KrInstance(user[0], birth_date.year, birth_date.month, birth_date.day, birth_time.hour, birth_time.minute, user[3])

# Use the KrInstance object to create a MakeSvgInstance object and generate the SVG file
make_svg_instance = MakeSvgInstance(kr_instance, chart_type="Natal")
make_svg_instance.makeSVG()

# Close the cursor and connection
cur.close()
conn.close()
