decimal 设计
```go

type decimal string

```
2.固定小数点位数
```go

type decimal struct{
    negative bool
	integer int64
	negExp uint32 // 以10为底的负指数
	frac uint32
}

```

3. 不固定小数点位数,占用存储更多
```go

type decimal struct{
	Int int64
	Frac []uint8 // 前N位是以10为底的负指数
	exp  uint8
}

```