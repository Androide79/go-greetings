package greetings

import (
	"math/rand"
	"testing"
)

func TestGreet(t *testing.T) {

}

func TestGetSimpleGreet(t *testing.T) {

	var langNotPresent string = "kr"
	var defaultValue string = "Hello Dude!"
	//Get random known languange
	var lang = []string{"it", "de", "fr", "es"}
	var rndLangIndex int = rand.Int() % len(lang)
	var rndLang string = lang[rndLangIndex]

	// test 'getSimpleGreet' returns default value with a language not presents in language's array
	var getSimpleGreetUnknown string = getSimpleGreet(langNotPresent)

	if getSimpleGreetUnknown != defaultValue {
		t.Errorf("getSimpleGreet(\"%v\") failed, expected %v, got %v", langNotPresent, defaultValue, getSimpleGreetUnknown)
	} else {
		// log test success only if you use verbose test argument
		t.Logf("getSimpleGreet(\"%v\") success, expected %v, got %v", langNotPresent, defaultValue, getSimpleGreetUnknown)
	}

	// test 'getSimpleGreet' returns no-default value if language passed is know
	var getSimpleGreetKnown string = getSimpleGreet(rndLang)

	if getSimpleGreetKnown == defaultValue {
		t.Errorf("getSimpleGreet(\"%v\") failed, unexpected result %v", rndLang, getSimpleGreetKnown)
	} else {
		// log test success only if you use verbose test argument
		t.Logf("getSimpleGreet(\"%v\") success", langNotPresent)
	}

	// test 'getSimpleGreet' return 'Ciao Amico!' if passed language is 'it'
	var itResult string = getSimpleGreet("it")
	if itResult != "Ciao Amico!" {
		t.Errorf("getSimpleGreet(\"it\") failed, expected \"Ciao Amico!\", got %v", itResult)
	} else {
		// log test success only if you use verbose test argument
		t.Logf("getSimpleGreet(\"it\") success, expected \"Ciao Amico!\", got %v", itResult)
	}

	// test 'getSimpleGreet' return 'Bonjour, mec!' if passed language is 'fr'
	var frResult string = getSimpleGreet("fr")
	if frResult != "Bonjour, mec!" {
		t.Errorf("getSimpleGreet(\"fr\") failed, expected \"Bonjour, mec!\", got %v", frResult)
	} else {
		// log test success only if you use verbose test argument
		t.Logf("getSimpleGreet(\"fr\") success, expected \"Bonjour, mec!\", got %v", frResult)
	}

	// test 'getSimpleGreet' return 'Hallo Kumpel!' if passed language is 'de'
	var deResult string = getSimpleGreet("de")
	if deResult != "Hallo Kumpel!" {
		t.Errorf("getSimpleGreet(\"de\") failed, expected \"Hallo Kumpel!\", got %v", deResult)
	} else {
		// log test success only if you use verbose test argument
		t.Logf("getSimpleGreet(\"de\") success, expected \"Hallo Kumpel!\", got %v", deResult)
	}

	// test 'getSimpleGreet' return 'Hola, amigo!' if passed language is 'es'
	var esResult string = getSimpleGreet("es")
	if esResult != "Hola, amigo!" {
		t.Errorf("getSimpleGreet(\"es\") failed, expected \"Hola, amigo!\", got %v", esResult)
	} else {
		// log test success only if you use verbose test argument
		t.Logf("getSimpleGreet(\"es\") success, expected \"Hola, amigo!\", got %v", esResult)
	}
}

func TestGetGreetWithName(t *testing.T) {

}
