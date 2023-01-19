[TOC]

# 数据转换流程
- YD数据转为Trip行程
- IC数据转为Trans交通
- 对IC->Trans数据进行数据清洗
- 对YD数据进行拆分 -> 清洗 -> 过滤 -> 追加到Trans
- 将清洗后的数据导入文件
- 将文件中的数据导入数据库

## GO关于数据转换的代码
```golang
## YD 到 trip 数据的映射
type Trip struct {
	TripId      string //行程ID
	MatchMarker string //匹配标记[进站抵消 单边进 单边出 补登进 补登出 出站票]
	InTransId   string //进站ID
	OutTransId  string //出站ID
	In          Trans  //进站的 票ID 站台ID 站台名称 时间
	Out         Trans  //出站的 票ID 站台ID 站台名称 时间
}
Trip{
	TripId:      m["行程编号"][i],
	MatchMarker: m["匹配标记"][i],
	InTransId:   m["进站行程编号"][i],
	OutTransId:  m["出站行程编号"][i],
	In: Trans{
		TransCode:   In,
		TicketId:    m["虚拟卡号"][i],
		Line:        m["IN_STATION_ID"][i][0:2],
		StationId:   m["IN_STATION_ID"][i],
		StationName: m["进站车站"][i],
		TransId:     m["进站行程编号"][i],
		TransTime:   time.Time{},
		TransDate:   "",
	},
	Out: Trans{
		TransCode:   Out,
		TicketId:    m["虚拟卡号"][i],
		Line:        m["OUT_STATION_ID"][i][0:2],
		StationId:   m["OUT_STATION_ID"][i],
		StationName: m["出站车站"][i],
		TransId:     m["出站行程编号"][i],
		TransTime:   time.Time{},
		TransDate:   "",
	},
}


## IC 到 trans 数据的映射
type Trans struct {
	TransCode   string    `gorm:"index"`      //交通类型码[21进站 22出站]
	TicketId    string    `gorm:"index"`      //票ID
	Line        string    `gorm:"index"`      //地铁线路
	StationId   string    `gorm:"index"`      //站台ID
	StationName string    `gorm:"index"`      //站台名称
	TransId     string    `gorm:"primaryKey"` //进出站ID
	TransTime   time.Time `gorm:"index"`      //进出站时间
	TransDate   string    `gorm:"index"`      //进出站日期(凌晨1点前属于前一天)
}
Trans{
	transTime, err := time.Parse(TransTimeFormat, m["TXN_DATE"][i]+m["TXN_TIME"][i])
	if err != nil {
		return nil, err
	}
	list[i].TransCode = m["TRANS_CODE"][i]
	list[i].TicketId = m["TICKET_ID"][i]
	list[i].Line = m["TXN_STATION_ID"][i][0:2]
	list[i].StationId = m["TXN_STATION_ID"][i]
	list[i].StationName = station.StationNameById(m["TXN_STATION_ID"][i])
	list[i].TransTime = transTime
	list[i].TransDate = transTime.Add(-1 * time.Hour).Format("060102")
	list[i].SetTransId()
}
```

### 样例数据
```csv
TransCode,TicketId,Line,StationId,StationName,TransId,TransTime,TransDate,CreateAt
21,00004000000222953094,10,1025,中央公园西,21_1025_20210821171403_00004000000222953094,20210821171403,210821,2021-08-21T17:14:03Z
22,00004000000222953094,02,0204,大溪沟,22_0204_20210821183634_00004000000222953094,20210821183634,210821,2021-08-21T18:36:34Z
21,00004000001520344633,03,0312,六公里,21_0312_20210821111921_00004000001520344633,20210821111921,210821,2021-08-21T11:19:21Z
22,00004000001520344633,03,0321,观音桥,22_0321_20210821114951_00004000001520344633,20210821114951,210821,2021-08-21T11:49:51Z
21,00004000001520344633,03,0321,观音桥,21_0321_20210821144709_00004000001520344633,20210821144709,210821,2021-08-21T14:47:09Z
22,00004000001520344633,03,0312,六公里,22_0312_20210821151512_00004000001520344633,20210821151512,210821,2021-08-21T15:15:12Z
```

## 数据库结构
```sql
create table `test-db`.trans
(
    trans_code   varchar(191) null,
    ticket_id    varchar(191) null,
    line         varchar(191) null,
    station_id   varchar(191) null,
    station_name varchar(191) null,
    trans_id     varchar(191) not null
        primary key,
    trans_time   datetime(3)  null,
    trans_date   varchar(191) null
);

create index idx_trans_line
    on `test-db`.trans (line);

create index idx_trans_station_id
    on `test-db`.trans (station_id);

create index idx_trans_station_name
    on `test-db`.trans (station_name);

create index idx_trans_ticket_id
    on `test-db`.trans (ticket_id);

create index idx_trans_trans_code
    on `test-db`.trans (trans_code);

create index idx_trans_trans_date
    on `test-db`.trans (trans_date);

create index idx_trans_trans_time
    on `test-db`.trans (trans_time);
```

## 使用到的sql查询
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


## trans数据清洗流程
- 剔除 transCode不为 [21,22]的数据
- 剔除 0-6点的数据
- 对每一天的数据做进出站匹配。丢弃错误数据包括：连续进站，连续出站，仅出站，仅进站，同站进出

### 进出站匹配逻辑
- 一天内同一个票的所有trans数据按时间排序进行分桶(栈结构存储)
- 对桶内数据从末端按一条"出"，一条"进"的顺序取出，若不符合这个规则则丢弃(即上一步描述的错误数据)