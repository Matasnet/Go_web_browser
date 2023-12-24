# Prosta Tekstowa Przeglądarka w Go

## Opis

Prosta tekstowa przeglądarka internetowa napisana w języku Go. Program umożliwia pobieranie i wyświetlanie zawartości stron internetowych w formie tekstu.

## Użycie

### Uruchamianie programu

Aby uruchomić program, użyj następującego polecenia:

```bash
go run main.go  
```
  Ewentualnie możesz również podać adres
```bash
go run main.go https://example.com
```

### Dodanie protokołu
  Jeżeli nie podasz "https://" podczas wprowadzania adresu zostanie on dodany automatycznie

### Zależności
  Program korzysta z zależności, które są zarządzane za pomocą pliku go.mod. Aby zaktualizować je nalezy użyć polecenia

```bash
go mod tidy
```


