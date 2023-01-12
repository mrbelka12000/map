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

### Wiki

#### Definitions 
```
Point - это точка на карте может иметь название или же нет. Используется для прокладки roads, постройки polygons и lines

Polygon - это область, состоит из Point. Может быть чем угодно(здание, область, времянка)

Roads - это детальная схема дорог с каждым вьездом и выездом

Line - это только крупные дороги соединяющие города и области 

Column Point.WAY - это закодиравнные данные координат тип geometry(point, 3857)  
Column Geomerty.WAY - это закодиравнные данные координат тип geometry(geometry, 3857)  
```

#### Commands

Получаем центральную точку из геомитрии полигонов
```sql
select st_astext(st_transform(st_centroid(way), 4386))
from planet_osm_polygon
where name = 'Аграрный сервис';
```

Получение координат точек из geometry(sql type)
```sql
select st_astext(st_transform(way, 4326))
from planet_osm_polygon
where name = 'Аграрный сервис';
```

Получение координа точек из point(sql type)
```sql
SELECT  ST_X(ST_Transform (way, 4326)) AS "Longitude",
        ST_Y(ST_Transform (way, 4326)) AS "Latitude"
FROM planet_osm_point;
```

Получение ближайших улиц от точки 
```sql
select st_asewkt(way , 1)
from planet_osm_point
where name like 'Избушка'
limit 1;

Output: 'SRID=3857;POINT(7954839.3 6644558.6)'
    
select name, st_transform(way, 4326), way<-> 'SRID=3857;POINT(7954839.3 6644558.6)'::geometry as dist
from planet_osm_roads
order by dist
limit 5;
```