## 简介（Introduction）

Go语言编写轻量级网络库 (grapeNet is a lightweight and Easy Use Network Framework)

可用于游戏服务端、强网络服务器端或其他类似应用场景，每个模块单独提取并且拥有独立的使用方法，内部耦合性较轻。

其实GO语言曾经有过很多强架构的框架，比如GOWOLRD之类的，已经足够了，但是我会将库用于各种轻量级应用不需要过于复杂的内容，所以我设计了GrapeNet，目的是模块独立化。
你可以拆开只使用其中很小的模块，也可以组合成一个服务端，并且在架构中设计也较为轻松，至于热更新的问题，目前脚本数据支持热更新，并且是自动的，只要跑一下UPDATE即可，程序本身稍后测试后发布（仅支持LINUX）。

慢慢更新中，很多坑要填，目前暂不适合用于商业项目。

## 模块表（Function）

* Lua脚本绑定管理（可绑定任何类型的函数）
* 日志库（底层采用Seelog）
* 函数管理系统（可以根据任何类型参数将其与函数绑定并互相调用）
* 流处理
* Tcp网络
* Websocket网络 （基础版）
* Codec（任意类型注册对象并在其他位置动态创建该对象）

## 依赖第三方

* Seelog (github.com/cihub/seelog)
* Gopher-lua(github.com/yuin/gopher-lua)
* Gopher-luar(layeh.com/gopher-luar)

不依赖任何CGO内容，lua本身也是纯GO实现。

## 文档（Documentation）

### Lua脚本库使用

Lua库为线程安全库，可以在任意协程中并行调用脚本文件中的函数，也可以合并脚本库。

#### 新建文件脚本

```go
    lua := NewFromFile("testlua", "../_lua_tests/luascripts/call_lua_test.lua")
```

#### 新建字符串脚本（用于加密LUA文件或内存脚本）

```go
    lua := NewVM("testRegister")
    lua.DoString(sScript) // 执行字符串
```

#### 获取脚本中的结构

```go
	person := Person{}
	err := lua.GetTable("person", &person)
	if err != nil {
		t.Error(err)
		return
	}
```

#### 绑定与GO互相调用

脚本模块允许Go函数与Lua脚本无缝调用且线程安全。

> 注意：部分代码请参考LuaCall_test.go中的代码

```go
    // 绑定本地GO函数
    lua := NewVM("testRegister")
	lua.SetGlobal("TestGoFunc", bindTestFn)

    // 一定要在DO文件之前绑定，否则调用该文件时可能无效
    // 调用Lua中函数

    err := lua.CallGlobal("TestAbc", "a", 2)
	if err != nil {
		return
	}
```

### 其他模块

暂无文档 有待补全