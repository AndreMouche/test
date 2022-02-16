## Test [tidwall/btree](https://github.com/tidwall/btree) VS [google/btree](https://github.com/google/btree)

| Func  |[tidwall/btree](https://github.com/tidwall/btree)   | [google/btree](https://github.com/google/btree)  | 
|:---|---:|---|
| go test -bench="BenchmarkInsert"  |  231.6 ns/op |215.2 ns/op   |  
| go test -bench="BenchmarkInsert"   | 229.6 ns/op  | 198.3 ns/op  |   
| go test -bench="BenchmarkSeek" | 128.6 ns/op  |221.9 ns/op   |   
| go test -bench="BenchmarkSeek"  | 127.4 ns/op  | 202.5 ns/op  |   
| go test -bench="BenchmarkDelet"  | 205.7 ns/op  | 210.4 ns/op  |   
| go test -bench="BenchmarkDelet"  |  212.5 ns/op  | 235.9 ns/op  |   
| go test -bench="BenchmarkGet"  | 150.9 ns/op  |   173.5 ns/op|   
| go test -bench="BenchmarkGet"  | 147.5 ns/op  |172.2 ns/op   |   
| go test -bench="BenchmarkAscendGreaterOrEqual"|37669 ns/op | 57381 ns/op|
| go test -bench="BenchmarkAscendGreaterOrEqual"| 37254 ns/op| 56140 ns/op|
| go test -bench="BenchmarkDescendLessOrEqual"|34492 ns/op|81551 ns/op|
| go test -bench="BenchmarkDescendLessOrEqual"|35603 ns/op|81346 ns/op|