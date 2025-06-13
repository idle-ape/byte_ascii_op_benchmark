#### []byte与ASCII值加减实现63.6%性能加速以及83%内存消耗下降，基准测试如下：
```
$ /usr/local/sbin/go test -benchmem -run=^$ -bench "^(BenchmarkGenAesKeyWithMap|BenchmarkGenAesKeyWitthSlice)$" . -v
goos: linux
goarch: amd64
pkg: github.com/idle-ape/byte_ascii_op_benchmark
cpu: Intel(R) Celeron(R) J4025 CPU @ 2.00GHz
BenchmarkGenAesKeyWithMap
BenchmarkGenAesKeyWithMap-2               225337              5071 ns/op             958 B/op         27 allocs/op
BenchmarkGenAesKeyWitthSlice
BenchmarkGenAesKeyWitthSlice-2            600284              1846 ns/op             160 B/op          7 allocs/op
PASS
ok      github.com/idle-ape/byte_ascii_op_benchmark     2.339s
```

```
$ /usr/local/sbin/go test -benchmem -run=^$ -bench "^(BenchmarkGenAesKeyWithMap|BenchmarkGenAesKeyWitthSlice)$" . -v
goos: linux
goarch: amd64
pkg: github.com/idle-ape/byte_ascii_op_benchmark
cpu: Intel(R) Celeron(R) J4025 CPU @ 2.00GHz
BenchmarkGenAesKeyWithMap
BenchmarkGenAesKeyWithMap-2               220682              5062 ns/op             958 B/op         27 allocs/op
BenchmarkGenAesKeyWitthSlice
BenchmarkGenAesKeyWitthSlice-2            608526              1860 ns/op             160 B/op          7 allocs/op
PASS
ok      github.com/idle-ape/byte_ascii_op_benchmark     2.331s
```

#### 性能测试结果解释
| 指标          | GenAesKeyWithMap (基准) | GenAesKeyWithSlice | 比较结果                 |
|---------------|-----------------------|---------------------|------------------------|
| 执行耗时      | 5071 ns/op            | 1846 ns/op            | ⬇63.6% (快2.745倍)   |
| 内存分配量    | 958 B/op              | 160 B/op              | ⬇83.3%                |
| 分配次数      | 27 allocs/op          | 7 allocs/op           | ⬇74.07%               |
| 吞吐量        | 225,337 次/秒         | 600,284 次/秒         | ⬆166.39%               |