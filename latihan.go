package main
import "fmt"

func aneh(n int)bool{
    var b bool
    var i, hasil int
    hasil = 0
    i = 1
    for i < n {
        if n%i == 0{
            hasil+= i
            i++
        }else{
            i++
        }
    }
    if hasil > n {
        b = true
    }else{
        b = false
    }
    return b

}
func main(){
    var i int
    fmt.Scan(&i)
    fmt.Print(aneh(i))
}