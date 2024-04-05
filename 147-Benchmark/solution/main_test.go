package solution

import (
	"fmt"
	"testing"
)

func TestGreet(t *testing.T) {
	s := Greet("selvin")
	if s != "Welcome to the course selvin" {
		t.Error("got", s, "Expected", "Welcome to the course selvin")
	}

}

func ExampleGreet() {
	fmt.Println(Greet("selvin"))
	//Output:
	//Welcome to the course selvin
}

func BenchmarkGreet(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Greet("selvin")
	}
}

// If you error in godoc

/*

$ export PATH=$PATH:~/go/bin

$ GOBIN=$(go env GOPATH)/pkg/mod/bin go install golang.org/x/tools/cmd/godoc

$ godoc -http=:8080

using module mode; GOMOD=C:\Users\kumaran\Desktop\Learn Go\go.mod

*/

/* commands

go test
go test -bench .
go test -cover
go test -coverprofile c.out
go tool cover -html=c.out before running this command run this "go test -coverprofile c.out"

*/
