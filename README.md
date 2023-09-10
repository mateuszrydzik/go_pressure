# go pressure;
#### pozyskiwanie danych o ciśnieniu

zastanwiałem się, czy moje losowe bóle głowy i zmiany w samopoczuciu są spowodowane nagłymi zmianami ciśnienia atmosferycznego. zamiast sprawdzić aktualny stan pogodowy, napisałem aplikację w go, która pozyskuje aktualną wartość ciśnienia z api imgw, a następnie zapisuje ją do bazy danych.

do bazy zapisywane są wybrane dane:

- station: wybrana stacja pomiarowa
- pressure: wartość ciśnienia
- date: dzień, miesiąc i rok pomiaru
- hour: godzina pomiaru

## instalacja
aplikacja wymaga dockera i docker compose

1. zmienne środowiskowe

naley utworzy plik `.env` w głównym katalogu, na bazie pliku `.envtemplate`, i uzupełnić go o wartości.

2. uruchomienie

przy pierwszym uruchomieniu, należy wykonać polecenie
`docker-compose up --build`
przy kolejnych uruchomieniach, wystarczy wykonać polecenie
`docker-compose up`

przy uruchomionej aplikacji, można lokalnie połączyć się z bazą danych, wykorzystując podane zmienne środowiskowe do połączenia.