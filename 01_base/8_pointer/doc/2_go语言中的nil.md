# 1.说明
在go语言中，不同类型的数据零值是不一样的。<br/>

| 类型           | 零值                                  | 备注 |
| -------------- | ------------------------------------- | ---- |
| bool           | false                                 |      |
| 数值系列       | 0                                     |      |
| string         | ""                                    |      |
| pointer (指针) | nil                                   |      |
| slice (切片)   | nil                                   |      |
| map            | nil                                   |      |
| channel        | nil                                   |      |
| interface      | nil                                   |      |
| function       | nil                                   |      |
| struct         | 默认值是具体字段的默认值（而不是nil） |      |    

