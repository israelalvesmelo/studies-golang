package tabela

import (
	"strings"
	"testing"
)

const msgIndex = "%s (parte: %s) - Índices: esperado (%d) <> encontrado (%d)."

func TestIndex(t *testing.T) {
	testes := []struct {
		texto    string
		parte    string
		esperado int
	}{
		{"Cod@r é show", "Cod@r", 0},
		{"", "", 0},
		{"Opa", "opa", -1},
		{"leonardo", "o", 2},
	}

	for _, teste := range testes {
		t.Logf("Massa %v", teste) //Para exibir os logs, basta executar assim: go test./... -v
		atual := strings.Index(teste.texto, teste.parte)

		if atual != teste.esperado {
			t.Errorf(msgIndex, teste.texto, teste.parte, teste.esperado, atual)
		}
	}
}
