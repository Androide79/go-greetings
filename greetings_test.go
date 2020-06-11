package greetings

import "testing"

func TestHello(t *testing.T) {
	// test for empty argument
	var emptyName string = Greet("")

	if emptyName != "Hello Dude!" {
		t.Errorf("greetings(\"\") failed, expected %v, got %v", "Hello Dude!", emptyName)
	} else {
		// log test success only if you use verbose test argument
		t.Logf("greetings(\"\") success, expected %v, got %v", "Hello Pigi!", emptyName)
	}

	// test for a valid argument
	var name string = Greet("Pigi")

	if name != "Hello Pigi!" {
		t.Errorf("greetings(\"Pigi\") failed, expected %v, got %v", "Hello Pigi!", name)
	} else {
		// log test success only if you use verbose test argument
		t.Logf("greetings(\"Pigi\") success, expected %v, got %v", "Hello Pigi!", name)
	}
}
