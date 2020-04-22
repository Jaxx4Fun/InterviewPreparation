# 奇技淫巧记录
```go
const (
    Height = 1920
    Width = 1080
)
// 给切片的切片分配底层数组
var picture = make([][]uint8, Height)
var pixels = make([]uint8, Height*Width)
for i:=0; i< Height; i++ {
    picture[i], pixels = pixels[:Width], pixels[Width:]
}
```