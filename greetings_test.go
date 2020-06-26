package greetings

import (
	"fmt"
	"math/rand"
	"testing"

	"github.com/Pallinder/go-randomdata"
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

	// test 'getGreetWithName' returns default value with a language not presents in language's array
	var langNotPresent string = "kr"
	var rndName = randomdata.FirstName(randomdata.RandomGender)
	var defaultResult = fmt.Sprintf("Hello %v!", rndName)
	var getGreetWithNameResult = getGreetWithName(rndName, langNotPresent)

	if getGreetWithNameResult != defaultResult {
		t.Errorf("getGreetWithName(\"%v\", \"%v\") failed, expected %v, got %v", rndName, langNotPresent, defaultResult, getGreetWithNameResult)
	} else {
		// log test success only if you use verbose test argument
		t.Logf("getGreetWithName(\"%v\", \"%v\") success, expected %v, got %v", rndName, langNotPresent, defaultResult, getGreetWithNameResult)
	}

	// test 'getGreetWithName' returns 'Ciao <name_passed>!' if passed language is 'it'
	var itResult string = getGreetWithName(rndName, "it")
	var itExpectedResult string = fmt.Sprintf("Ciao %v!", rndName)

	if itResult != itExpectedResult {
		t.Errorf("getGreetWithName(\"%v\", \"it\") failed, expected %v, got %v", rndName, itExpectedResult, itResult)
	} else {
		// log test success only if you use verbose test argument
		t.Logf("getGreetWithName(\"%v\", \"it\") success, expected \"%v\", got %v", rndName, itExpectedResult, itResult)
	}

	// test 'getGreetWithName' returns 'Bonjour <name_passed>!' if passed language is 'fr'
	var frResult string = getGreetWithName(rndName, "fr")
	var frExpectedResult string = fmt.Sprintf("Bonjour %v!", rndName)

	if frResult != frExpectedResult {
		t.Errorf("getGreetWithName(\"%v\", \"fr\") failed, expected %v, got %v", rndName, frExpectedResult, frResult)
	} else {
		// log test success only if you use verbose test argument
		t.Logf("getGreetWithName(\"%v\", \"fr\") success, expected \"%v\", got %v", rndName, frExpectedResult, frResult)
	}

	// test 'getGreetWithName' returns 'Hallo <name_passed>!' if passed language is 'de'
	var deResult string = getGreetWithName(rndName, "de")
	var deExpectedResult string = fmt.Sprintf("Hallo %v!", rndName)

	if deResult != deExpectedResult {
		t.Errorf("getGreetWithName(\"%v\", \"de\") failed, expected %v, got %v", rndName, deExpectedResult, deResult)
	} else {
		// log test success only if you use verbose test argument
		t.Logf("getGreetWithName(\"%v\", \"de\") success, expected \"%v\", got %v", rndName, deExpectedResult, deResult)
	}

	// test 'getGreetWithName' returns 'Hola <name_passed>!' if passed language is 'es'
	var esResult string = getGreetWithName(rndName, "es")
	var esExpectedResult string = fmt.Sprintf("Hola %v!", rndName)

	if esResult != esExpectedResult {
		t.Errorf("getGreetWithName(\"%v\", \"es\") failed, expected %v, got %v", rndName, esExpectedResult, esResult)
	} else {
		// log test success only if you use verbose test argument
		t.Logf("getGreetWithName(\"%v\", \"es\") success, expected \"%v\", got %v", rndName, esExpectedResult, esResult)
	}
}
