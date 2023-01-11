## Map


### Download map osm.pbf
https://download.geofabrik.de/asia/kazakhstan.html

### Export osm.pbf map to Postgis
```db2_is
create extension postgis;
osm2pgsql -U postgres -W -d "dbName" -H "host" --number-processes 24 -C 20480 "path to file"
```

### Visualize map from postgres
https://qgis.org/ru/site/forusers/download.html
