## Goland之补充基础

### 一、修改字符串

要修改字符串，需要先将其转换成`[]rune`或`[]byte`，完成后再转换为`string`。无论哪种转换，都会重新分配内存，并复制字节数组。

```go
// 英文用byte	
s := "qqqww"
	tmp := []byte(s)
	tmp[0] = 'f'
	fmt.Println(string(tmp))

// 中文用runes
	s2 := "白萝卜"
	runeS2 := []rune(s2)
	runeS2[0] = '红'
	fmt.Println(string(runeS2))
```

