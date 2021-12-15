Simulacija se pokreće naredbom: go run -tags musl ./cmd/main

Kako bi se promjenilo opterećenje sustava može se dodati parametar koji određuje koliko često putnici dolaze u lift.
Taj parametar je broj veći jedan. Ukoliko se ništa ne navede, pretpostavljena vrijednost je 2.
Primjer: go run -tags musl ./cmd/main 5
