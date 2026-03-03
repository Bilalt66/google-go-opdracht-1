# google-go-opdracht-1

STAP 1
Het programma start met package main en importeert drie packages: fmt voor invoer en uitvoer, math/rand voor willekeurige getallen en time om de random generator te initialiseren.

STAP 2
In de main() functie wordt eerst de random generator ingesteld, daarna worden de waarden 1 en 100 opgeslagen als minimum en maximum. Met deze waarden wordt een willekeurig getal tussen 1 en 100 gegenereerd en opgeslagen in randomNum. Dit is het geheime getal. Vervolgens wordt er uitleg op het scherm geprint: welkom, het bereik van het getal en dat de speler 5 kansen heeft.

STAP 3
Daarna begint een for-loop die vijf keer loopt. In elke ronde wordt de functie askInput() aangeroepen om een getal van de gebruiker te lezen. De ingevoerde waarde wordt vergeleken met het geheime getal. Is het te laag, dan komt “HOGER!”. Is het te hoog, dan komt “LAGER!”. Is het goed, dan stopt het programma meteen. Na elke foutieve gok wordt berekend hoeveel kansen nog over zijn en dat wordt geprint. Als alle kansen op zijn, wordt het juiste nummer getoond.

STAP 4
De functie askInput() print een vraag, leest de invoer van de gebruiker in een variabele en geeft die waarde terug aan de main() functie.