# 说明

https://github.com/swaggo/swag/blob/master/README_zh-CN.md


## API操作

Example [celler/controller](https://github.com/swaggo/swag/tree/master/example/celler/controller)

| 注释                 | 描述                                                                                                    |                                                    |
| -------------------- | ------------------------------------------------------------------------------------------------------- | -------------------------------------------------- |
| description          | 操作行为的详细说明。                                                                                    |
| description.markdown | 应用程序的简短描述。该描述将从名为`endpointname.md`的文件中读取。                                       | // @description.file endpoint.description.markdown |
| id                   | 用于标识操作的唯一字符串。在所有API操作中必须唯一。                                                     |
| tags                 | 每个API操作的标签列表，以逗号分隔。                                                                     |
| summary              | 该操作的简短摘要。                                                                                      |
| accept               | API可以使用的MIME类型的列表。值必须如“[Mime类型](#mime-types)”中所述。                                  |
| produce              | API可以生成的MIME类型的列表。值必须如“[Mime类型](#mime-types)”中所述。                                  |
| param                | 用空格分隔的参数。`param name`,`param type`,`data type`,`is mandatory?`,`comment` `attribute(optional)` |
| security             | 每个API操作的[安全性](#security)。                                                                      |
| success              | 以空格分隔的成功响应。`return code`,`{param type}`,`data type`,`comment`                                |
| failure              | 以空格分隔的故障响应。`return code`,`{param type}`,`data type`,`comment`                                |
| header               | 以空格分隔的头字段。 `return code`,`{param type}`,`data type`,`comment`                                 |
| router               | 以空格分隔的路径定义。 `path`,`[httpMethod]`                                                            |
| x-name               | 扩展字段必须以`x-`开头，并且只能使用json值。                                                            |

## Mime类型

`swag` 接受所有格式正确的MIME类型, 即使匹配 `*/*`。除此之外，`swag`还接受某些MIME类型的别名，如下所示：

| Alias                 | MIME Type                         |
| --------------------- | --------------------------------- |
| json                  | application/json                  |
| xml                   | text/xml                          |
| plain                 | text/plain                        |
| html                  | text/html                         |
| mpfd                  | multipart/form-data               |
| x-www-form-urlencoded | application/x-www-form-urlencoded |
| json-api              | application/vnd.api+json          |
| json-stream           | application/x-json-stream         |
| octet-stream          | application/octet-stream          |
| png                   | image/png                         |
| jpeg                  | image/jpeg                        |
| gif                   | image/gif                         |

## 参数类型

- query
- path
- header
- body
- formData

## 数据类型

- string (string)
- integer (int, uint, uint32, uint64)
- number (float32)
- boolean (bool)
- user defined struct
