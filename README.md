# go-paginator
Pagination API for golang slices

#Example
```golang

import(
        "fmt"
        "github.com/goibibo/gopage"
)

func main() {
        p, _ := gopage.NewPaginator([]int{1, 2, 3, 4, 5})
        p.SetPageSize(2)

        fmt.Println(p.Page(1))
        fmt.Println(p.Page(2))
        fmt.Println(p.Page(3))
}
```
