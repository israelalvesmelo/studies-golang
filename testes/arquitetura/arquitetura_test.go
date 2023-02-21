package arquitetura

import (
	"runtime"
	"testing"
)

func TestDependente(t *testing.T) {
	t.Parallel() //É usado para rodar os testes em paralelo
	if runtime.GOARCH == "amd64" {
		t.Skip("Não funciona em arquitetura amd64") //Log para testes que foram pulados
	}
	//...
	t.Fail()
}
