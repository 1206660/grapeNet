## Lua脚本库使用

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