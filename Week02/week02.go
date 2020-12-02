git@github.com:Wpjwwwww/Go-000.git
package main
import (
"database/sql"
"sql"
"runtime"
)
func main(){
  db,err:=sql.Open("查询的数据库名","数据库地址")
  if err!=nil{
    fmt.Println(err)
  }
  defer db.close()
  type User struct{
  Id int32
  Name string
  Age int8
  }
  var user User
  err=db.QueryRow(
  'SELECT id,name,age WHERE id=?',2
  ).Scan(&user.Id,&user.Name,&user.Age)
  switch{
    case err==sql.ErrNoRows:
    case err !=nil:
    if _,file,line,ok :=runtime.Caller(0);
    ok{
      fmt.Println(err,file,line)}
  }
}
//总结：Wraps erros时，1.错误要被日志记录；2.应用程序处理错误，要保证100%的完整性；3.之后不再报告当前错误；




