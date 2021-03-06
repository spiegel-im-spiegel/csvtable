# Demonstration for CSV/TSV Access

[![GitHub license](https://img.shields.io/badge/license-CC0-blue.svg)](https://raw.githubusercontent.com/spiegel-im-spiegel/csvtable/master/LICENSE)

## Demonstration

See [sample.go](https://github.com/spiegel-im-spiegel/csvtable/blob/master/sample/sample.go).

```
$ cat data/hightemp.txt
city	location	temperature	date
高知県	江川崎	41	2013-08-12
埼玉県	熊谷	40.9	2007-08-16
岐阜県	多治見	40.9	2007-08-16
山形県	山形	40.8	1933-07-25
山梨県	甲府	40.7	2013-08-10
和歌山県	かつらぎ	40.6	1994-08-08
静岡県	天竜	40.6	1994-08-04
山梨県	勝沼	40.5	2013-08-10
埼玉県	越谷	40.4	2007-08-16
群馬県	館林	40.3	2007-08-16
群馬県	上里見	40.3	1998-07-04
愛知県	愛西	40.3	1994-08-05
千葉県	牛久	40.2	2004-07-20
静岡県	佐久間	40.2	2001-07-24
愛媛県	宇和島	40.2	1927-07-22
山形県	酒田	40.1	1978-08-03
岐阜県	美濃	40	2007-08-16
群馬県	前橋	40	2001-07-24
千葉県	茂原	39.9	2013-08-11
埼玉県	鳩山	39.9	1997-07-05
大阪府	豊中	39.9	1994-08-08
山梨県	大月	39.9	1990-07-19
山形県	鶴岡	39.9	1978-08-03
愛知県	名古屋	39.9	1942-08-02

$ cat data/hightemp.txt | go run sample/sample.go
city	temperature
高知県	41
埼玉県	40.9
岐阜県	40.9
山形県	40.8
山梨県	40.7
和歌山県	40.6
静岡県	40.6
山梨県	40.5
埼玉県	40.4
群馬県	40.3
群馬県	40.3
愛知県	40.3
千葉県	40.2
静岡県	40.2
愛媛県	40.2
山形県	40.1
岐阜県	40
群馬県	40
千葉県	39.9
埼玉県	39.9
大阪府	39.9
山梨県	39.9
山形県	39.9
愛知県	39.9
```
