

```
db, err := sql.Open("sqlite", "./english.db")
```

`sql.Open(...)`这是 Go 标准库 `database/sql` 包中的一个函数，用来**打开一个数据库连接**。

`"sqlite"`指定的 **驱动名称**，它告诉 `sql.Open` 要使用哪种类型的数据库驱动。

`"./english.db"`是**数据库文件的路径**，如果这个文件存在，程序会尝试连接这个 SQLite 数据库。如果这个文件不存在，SQLite 会**自动创建一个新的数据库文件**。

```
start := time.Now()
fmt.Printf("数据读取到插入完成，总用时：%.2f 秒\n", time.Since(start).Seconds())
```

`time.Since(start).Seconds()`计算自一个时间点（start）到现在经过的时间，并以秒为单位输出。

```
tx, err := db.Begin()
```

从数据库连接`db`中**开启一个数据库事务**，返回一个事务对象 `tx`（类型是 `*sql.Tx`），适合批量数据操作场景

这样是只开始一次数据库事务，在这个事务中批量操作很多条数据，再一次性提交，能极大减少磁盘写入次数。

```go
_, err := db.Exec(q)
```

再数据库`db`中按照`q`代表的语句建表。

```
res, err := tx.Exec("INSERT OR IGNORE INTO words(word) VALUES(?)", word.Word)
```

在一个数据库事务（`tx`）中**执行一条 SQL 插入语句**，如果单词已经存在就忽略插入。往 `words` 表中的 `word` 字段插入以`？`表示的`word.Word`

```go
if id, err := res.LastInsertId(); err == nil && id != 0 {
    // 新插入
} else {
    //已存在
}
```

`res.LastInsertId()` 返回本次 `INSERT` 操作影响的主键 ID：如果成功插入，返回新生成的 ID；如果被 IGNORE（因为重复），返回 0

```
err = tx.QueryRow("SELECT id FROM words WHERE word = ?", word.Word).Scan(&wordID)
```

查找数据库中某个**已存在**单词的 ID 并赋值给 `wordID`

```
Exec()
```

执行SQL语句的方法，适用于 `CREATE TABLE` `INSERT` `UPDATE` `DELETE`，以及其他不返回结果的 SQL

```
log.Fatal(err)
```

Go 语言中一种**快速终止程序并输出错误信息**的方式。把 `err` 的内容输出到控制台（标准错误输出），程序立即退出，**状态码为 1**
