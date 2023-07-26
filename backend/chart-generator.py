import psycopg2
from kerykeion import KrInstance, MakeSvgInstance
from datetime import datetime
import sys
import os

# Define the database connection details
dbUser = "starfja8_vivian"
dbPass = "PZq(tDO^0NjV"  # Your actual password
dbName = "starfja8_users"
dbHost = "localhost"
dbPort = "5432"

# Connect to the PostgreSQL database
print("Connecting to the database...")
conn = psycopg2.connect(
    dbname=dbName,
    user=dbUser,
    password=dbPass,
    host=dbHost,
    port=dbPort
)

# Create a cursor object
cur = conn.cursor()

# Get the user data from the command-line arguments
first_name, birth_date, birth_time, city, chart_type = sys.argv[1:]

# Parse the birth date and time
print("Parsing birth date and time...")
birth_date = datetime.strptime(birth_date, "%Y-%m-%d")
birth_time = datetime.strptime(birth_time, "%H:%M")

# Create a KrInstance object
print("Creating KrInstance object...")
kr_instance = KrInstance(first_name, birth_date.year, birth_date.month, birth_date.day, birth_time.hour, birth_time.minute, city)

# Use the KrInstance object to create a MakeSvgInstance object
print("Generating SVG file...")
make_svg_instance = MakeSvgInstance(kr_instance, chart_type=chart_type)

# Define the directory to save the SVG file
svg_dir = "../frontend/src/app/assets/charts/"
os.makedirs(svg_dir, exist_ok=True)

# Define the SVG file path
make_svg_instance.svg_file_path = os.path.join(svg_dir, f"{first_name}{chart_type}Chart.svg")

# Generate the SVG file
make_svg_instance.makeSVG()

# Close the cursor and connection
cur.close()
conn.close()

print("SVG file generated.")
