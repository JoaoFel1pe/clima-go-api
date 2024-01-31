from geopy.geocoders import Nominatim
import sys
import json

def get_geolocation(city):
    geolocator = Nominatim(user_agent="clima_tempo")
    location = geolocator.geocode(city)
    if location is None:
        raise Exception("It wasn't possible to convert the location to geo coordinates")
    
    return {"lat": location.latitude, "lng": location.longitude}

city = sys.argv[1]
result = get_geolocation(city)
json_result = json.dumps(result)
sys.stdout.write(json_result)