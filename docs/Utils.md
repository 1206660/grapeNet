## Util辅助函数库

本库用于补充日常使用中需要繁琐编写代码，但可以通过简单库替代的代码，例如三元运算符、interface{}的类型转换检测等。

已在实际项目中应用，请放心使用！

### 三元逻辑运算库

Golang是不存在三元运算符的，例如 在C/C++中的 c ? a : b 的写法是没有的，所以我写了这个库来补充这类需要但总是要繁琐编写的代码。
该库有预设几个简单类型，有一个通用的检测。

```
    // 预设类型 字符串
    isOpen := utils.Ifs(a.open,"开启","关闭")

    // 数值
    battleValue := utils.Ifn(a.battle,1000,3000)

    // 嵌套运算
    battleValue = utils.Ifn(a.sp,1000,utils.Ifn(a.sp2,3000,6000))
```

所有代码均可以使用 ```utils.If``` 来做泛型类型扩展。

### interface{}类型转换

我本人是将这个库用于需要大量interface{}转换，但是在懒得每次都去写代码的情况，例如 map中具体数值的取出和转换。

这个库同时具有强制类型转换的行为，就是说 你可以传入一个int64然后转换为一个int。

并且支持传入一个string转换为一个int或float哦

这个库目前我使用的场景有2类:

* 在大量interface{}需要转换为指定类型时，场景很多，例如map取值或某些函数返回interface{}
* 在大量字符串例如 "10000" 这种，需要转换为真正的数字时，由于每次需要编写指定函数，而must库恰恰不需要并支持很多类型。

那就以map为例：

```
 check := map[string]interface{}{
     "ret":0,
     "data1":2000,
     "data2":"this is data2",
 }

  valI64 := MustInt64(check["data1"], 1000)
  if valI64 != 2000 {
	t.Fail()
	return
  }

  valStr := MustString(check["data2"],"")
  if valStr != "this is data2" {
    t.Fail()
	return
  }
```

特殊类型转换 ，例如 字符串转换为int或bool字符串分析并转换为bool等

```
    valF64 := MustFloat64("1000", 0.01)
	if valF64 == 0.01 {
		t.Fail()
		return
	}

	valF64 = MustFloat64("4.345", 0.01)
	if valF64 == 0.01 {
		t.Fail()
		return
	}

	valBool := MustBool("true", false)
	if valBool == false {
		t.Fail()
		return
	}
```

还有详细用法可以参考convert_test.go，本库线程安全，且有benchmark，性能还不错！

## Jobs轻并行任务类

它很轻很轻，不适用于重度任务分布式计算，只用于比如单函数中需要大量计算，我把它拆成子函数，丢入jobs中，并行执行。

例如，我需要同时取出玩家装备信息、战斗信息、宠物信息等，使用这个库同时发起请求并赋值某个结果合并发送。

> 注意：非重度，长时间且需要拆分的任务，不要使用这个类，因为有大量反射耗时且ALLOC很多！在我机器平均完成单一任务执行需要 4147 ns/op & 11 allocs/op，我感觉好慢！
> 你可以在你机器上跑跑benchmark

### 基本用法

```
	jobs := &SyncJob{}

	err := jobs.AppendR(func(a, rb string) string {
		fmt.Println(a, "inter call")
		return rb
	}, func(r string) {
		fmt.Println(r, "return")
	}, "args1", "return")
	if err != nil {
		fmt.Println(err)
		t.Fail()
		return
	}

	err = jobs.AppendR(func(a, rb string) string {
		fmt.Println(a, "inter call 02")
		return rb
	}, func(r string) {
		fmt.Println(r, "return 02")
	}, "args2", "return2")
	if err != nil {
		fmt.Println(err)
		t.Fail()
		return
	}

	jobs.StartWait()
```

## []byte合并函数库

一般用于大量[]byte类型用于协议传输时合并多个序列化后的[]byte，并可以将其拆开

```
	mergeBuf := MergeBinary([]byte("i am first words"), []byte("i am second words"))

	splitBuf := SplitBinary(mergeBuf)
```