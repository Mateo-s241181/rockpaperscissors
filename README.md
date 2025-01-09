# Schere-Stein-Papier

Dieses Repo enthält Aufgaben für eine Beispiel-Implementierung zu "Schere-Stein-Papier".
Ziel ist es, anhand dieses einfachen Spiels den Umgang mit selbstdefinierten Datentypen
zu üben.

## Empfohlene Aufgaben-Reihenfolge

Sie sollten zuerst die Aufgaben im Package `values` bearbeiten, danach die in `player`
und zuletzt die in `game`.

## Anmerkungen

* In der Datei `player.go` im Package `player` gibt es keine Aufgaben.
  Hier wird ein *Interface* definiert, eine Vorlage, die sowohl vom Datentyp `Human`
  als auch von `Bot` implementiert wird.
* In `rps.go` im Verzeichnis `rps` gibt es ebenfalls keine Aufgaben.
  Hier ist ein fertiges Spiel implementiert, das die Datentypen `Game`, `Player` und `Value`
  verwendet.
  * Beachten Sie, dass bei der Entwicklung nur die Tests verwendet werden,
    dass das eigentliche Spiel dann aber eine `main`-Funktion hat.
  * Beachten Sie auch, dass die eigentliche Benutzereingabe in `Human` nicht getestet wird.
    Diese müssen Sie durch das laufende Spiel testen oder sich eigene `main`-Programme
    dafür schreiben.
