# go-mysql-test

Eseguire

```
go run main.go
```

In una finestra separata eseguire

```
go test
```

Risultato atteso

```
=== RUN   TestMain
--- PASS: TestMain (0.00s)
    main_test.go:29: [name email phone_numbers created_at] 
    main_test.go:33: *********************** 
    main_test.go:60: name = Test1 Doe 
    main_test.go:60: email = john@doe.com 
    main_test.go:60: phone_numbers = ["555-555-555"] 
    main_test.go:60: created_at = 2019-11-17 15:52:24 
    main_test.go:33: *********************** 
    main_test.go:60: name = Test2 Doe 
    main_test.go:60: email = johnalt@doe.com 
    main_test.go:60: phone_numbers = [] 
    main_test.go:60: created_at = 2019-11-17 15:52:24 
    main_test.go:33: *********************** 
    main_test.go:60: name = Test3 Doe 
    main_test.go:60: email = jane@doe.com 
    main_test.go:60: phone_numbers = [] 
    main_test.go:60: created_at = 2019-11-17 15:52:24 
    main_test.go:33: *********************** 
    main_test.go:60: name = Test4 Bob 
    main_test.go:60: email = evilbob@gmail.com 
    main_test.go:60: phone_numbers = ["555-666-555","666-666-666"] 
    main_test.go:60: created_at = 2019-11-17 15:52:24 
PASS
ok      github.com/eturella/go-mysql-test       0.022s
```


Per provare l'accesso da altri linguaggi

```
php main_test.php
```
