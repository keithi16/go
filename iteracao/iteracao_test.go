package iteracao

import "testing"


func TestRepetir(t *testing.T) {
	repeticoes := Repetir("a", quantidadeRepeticoes)
	esperado := "aaaaaa"

	if repeticoes != esperado {
		t.Errorf("esperado '%s', mas obteve '%s'", esperado, repeticoes)
	}
}

func BenchMarkRepetir(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repetir("a", quantidadeRepeticoes)
	}
}