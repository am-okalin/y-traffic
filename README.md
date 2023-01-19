## sql
```sql
# drop table `test-db`.trans
select count(*) from trans;

# 1.周五_0820_观音桥_进站
select DATE_FORMAT(trans_time, '%Y-%m-%d %H:00:00') as time, count(*)
from trans
where trans_date = "210820"
  and station_name = "观音桥"
  and trans_code = "21"
group by time order by time;

# 2.周五_0820_观音桥_出站
select DATE_FORMAT(trans_time, '%Y-%m-%d %H:00:00') as time, count(*)
from trans
where trans_date = "210820"
  and station_name = "观音桥"
  and trans_code = "22"
group by time order by time;

# 3.周一_0816_光电园_进站
select DATE_FORMAT(trans_time, '%Y-%m-%d %H:00:00') as time, count(*)
from trans
where trans_date = "210816"
  and station_name = "光电园"
  and trans_code = "21"
group by time order by time;

# 4.周一_0816_光电园_出站
select DATE_FORMAT(trans_time, '%Y-%m-%d %H:00:00') as time, count(*)
from trans
where trans_date = "210816"
  and station_name = "光电园"
  and trans_code = "22"
group by time order by time;

# 5.周六_0821_重庆北站南广场_进站
select DATE_FORMAT(trans_time, '%Y-%m-%d %H:00:00') as time, count(*)
from trans
where trans_date = "210821"
  and station_name = "重庆北站南广场"
  and trans_code = "21"
group by time order by time;

# 6.周六_0821_重庆北站南广场_出站
select DATE_FORMAT(trans_time, '%Y-%m-%d %H:00:00') as time, count(*)
from trans
where trans_date = "210821"
  and station_name = "重庆北站南广场"
  and trans_code = "22"
group by time order by time;

# 7.周六_0821_04_进站
select DATE_FORMAT(trans_time, '%Y-%m-%d %H:00:00') as time, count(*)
from trans
where trans_date = "210821"
  and line = "04"
  and trans_code = "21"
group by time order by time;

# 8.周六_0821_04_出站
select DATE_FORMAT(trans_time, '%Y-%m-%d %H:00:00') as time, count(*)
from trans
where trans_date = "210821"
  and line = "04"
  and trans_code = "22"
group by time order by time;


# 1. 0816_08-10_车站分组_数量排序
select station_name, count(*) as count
from trans
where trans_time between "2021-08-16 08:00:00" and "2021-08-16 10:00:00"
group by station_name order by count desc limit 10;

# 2. 0816_17-19_车站分组_数量排序
select station_name, count(*) as count
from trans
where trans_time between "2021-08-16 17:00:00" and "2021-08-16 19:00:00"
group by station_name order by count desc limit 10;



```