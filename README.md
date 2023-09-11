# go pressure;
#### pozyskiwanie danych o ciśnieniu

zastanawiałem się, czy moje losowe bóle głowy i zmiany w samopoczuciu są spowodowane nagłymi zmianami ciśnienia atmosferycznego. zamiast sprawdzić aktualny stan pogodowy, napisałem aplikację w go, która pozyskuje aktualną wartość ciśnienia z api imgw, a następnie zapisuje ją do bazy danych.

do bazy zapisywane są wybrane dane:

- station: wybrana stacja pomiarowa
- pressure: wartość ciśnienia
- date: dzień, miesiąc i rok pomiaru
- hour: godzina pomiaru

## instalacja
aplikacja wymaga dockera i docker compose

### 1. zmienne środowiskowe

naley utworzy plik `.env` w głównym katalogu, na bazie pliku `.envtemplate`, i uzupełnić go o wartości.

jeśli chcemy korzystać zarówno z aplikacji i bazy danych wewnątrz kontenerów, należy zostawić domyślne wartości z pliku `.envtemplate`. w pozostałych przypadkach, wypełniamy plik `.env` o odpowiednie wartości do połączenia aplikacji z docelową bazą danych. 

### 2. uruchomienie kontenerów

przy pierwszym uruchomieniu, należy wykonać polecenie
`docker-compose up --build`
przy kolejnych uruchomieniach, wystarczy wykonać polecenie
`docker-compose up`

przy uruchomionej aplikacji, można lokalnie połączyć się z bazą danych, wykorzystując podane zmienne środowiskowe do połączenia.

### 3. uruchomienie aplikacji

aplikacja jest dostępna wewnątrz kontenera

`docker exec -it {CONTAINER_NAME}_app bash`

dla domyślnej zmiennej CONTAINER_NAME, polecenie będzie wyglądało następująco:
`docker exec -it pressure_app bash`

po uzupełnieniu pliku `.env` można sprawdzić, czy aplikacja działa poprawnie, korzystając z polecenia `go run main.go`. jeśli nie zostanie zwrócony żaden błąd, aplikacja działa poprawnie, a w tabeli powinien pojawić się nowy rekord.

jeśli aplikacja działa poprawnie, możemy zbudować plik wykonywalny, przy pomocy komendy `go build main.go`. w celu zbudowania pliku pod wybrane środowisko, należy zmodyfikować polecenie o odpowiednie wartoci docelowego systemu operacyjnego oraz architektury.

na przykład, dla macbook z apple m1, polecenie będzie wyglądało następująco:
`GOOS=darwin GOARCH=arm64 go build main.go`

spis dostępnych GOOS i GOARCH można znaleźć pod adresem:
https://gist.github.com/zfarbp/121a76d5a3fde562c3955a606a9d6fcc

plik wykonywalny zostanie zbudowany w katalogu aplikacji.
## działanie aplikacji

uruchamiając aplikację, zostanie wykonane zapytanie do api imgw, a następnie zapisanie danych do bazy danych. przypadkiem użycia i głównym wykorzystaniem tej aplikacji jest dodanie zadania do crona, w celu regularnego pozyskiwania danych (np. trzy razy dziennie).