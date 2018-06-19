# go-paginator

Pagination API for golang slices [![Circle CI](https://circleci.com/gh/goibibo/gopage.png?style=badge)](https://github.com/goibibo/gopage)

#Example
```golang

import(
        "fmt"
        "github.com/shane-exley/gopage"
)

func main() {
        p, _ := gopage.NewPaginator([]int{1, 2, 3, 4, 5})
        p.SetPageSize(2)

        fmt.Println(p.GetPageCount())

        fmt.Println(p.Page(1))
        fmt.Println(p.Page(2))
        fmt.Println(p.Page(3))
}
```
