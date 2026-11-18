# Go: primers passos

Petits exemples per aprendre els fonaments de **Go (Golang)** a l'assignatura de Programació de Serveis i Processos (PSP) del CFGS de Desenvolupament d'Aplicacions Multiplataforma.

## Contingut

| Carpeta | Què mostra |
|---|---|
| `001_helloworld` | Servidor HTTP mínim amb `net/http` i rutes amb paràmetre (`GET /` i `GET /{name}`). |
| `002_declaracio` | Formes de declarar variables: `var`, inferència de tipus, declaració curta `:=` i blocs `var (...)`. |
| `003_estructures_control` | Estructures de control: `for` com a `while`, `for` clàssic i `for ... range` sobre un array. |
| `004_estructures` | `struct`, mètodes i composició (embedding) com a alternativa a l'herència. |

## Requisits

- [Go](https://go.dev/dl/) 1.22 o superior (les rutes amb `{name}` requereixen el `http.ServeMux` millorat).

## Execució

Cada carpeta és un programa independent:

```bash
cd 001_helloworld
go run main.go
```

L'exemple `001_helloworld` escolta al port 8080: prova `http://localhost:8080/` o `http://localhost:8080/Sara`.

Per compilar: `go build -o main .`
