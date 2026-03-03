# google-go-opdracht-1

We maken een package main aan, dit zorgt ervoor dat het programma uitgevoerd kan worden.

Daarna importeren we fmt, math/rand en time.
fmt gebruiken we voor printen en input,
math/rand voor een willekeurig getal,
time om de random generator goed te starten.

Dan maken we de functie main, hier begint het programma.

In main starten we de random generator met de huidige tijd en daarna maken we twee variabelen: min is 1 en max is 100.

Met rand.Intn maken we een willekeurig getal tussen 1 en 100 en slaan dat op in randomNum.

Vervolgens printen we een welkom bericht en leggen we uit dat het getal tussen 1 en 100 zit.
Daarna laten we de speler een moeilijkheid kiezen.

We maken de variabele diff en lezen de keuze in met fmt.Scan, met een if controleren we of de keuze geldig is. Als dat niet zo is zetten we hem automatisch op Easy.

Daarna maken we de variabelen chances en diffString. Met een switch passen we het aantal kansen aan. Bij Medium krijgt de speler 5 kansen, bij Hard 3 kansen.

Dan printen we hoeveel kansen de speler heeft.

Daarna begint een for loop. Deze loop draait zolang de speler nog kansen heeft.

In elke ronde roepen we de functie askInpu aan. Die functie vraagt om een getal en geeft dat terug.

We vergelijken de invoer met het random getal, is het te laag dan zeggen we dat het hoger moet. Is het te hoog dan zeggen we dat het lager moet. Is het goed, dan printen we dat het juist is en stoppen we het programma. Na elke gok berekenen we hoeveel kansen nog over zijn en printen we dat.

Onder de main functie staat askInput.
Deze functie vraagt om invoer, leest het getal in en geeft het terug aan de main functie.