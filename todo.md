- [] CustomErrors
  - [] test error
  - [] http errors
  - [] comparação de erro
- [] Json
  - [] ToJson
  - [] ToJsonString
  - [] FromJson?



GOPRIVATE=github.com/allanborba/* go get github.com/allanborba/go-utilitaries@main



# Feature: assert de slices

## Requisitos
- Criar um novo método Slices, pode chamar de SlicesNew por enquanto.
- Ele deve iterar pelo slice expected e verificar se o elemento existe no result. 
- Os slices de expected e result não precisam estar ordenados iguais.
- Caso exista um elemento no expected que não esteja presente no result, deve mostrar o erro e indicar os elementos faltantes
- Caso exista elemento no result que não esteja presente no expected, deve mostrar o erro e indicar os elementos extras
- Caso os slices sejam iguais, não deve mostrar erro
- O valor recebido é um slice de qualquer tipo
- Utilize o @object para fazer as comparações
- O caso o slice seja de structs, e as structs possuam campo que tenha slice, ele deve fazer a comparação pelos slices de forma recursiva utilizando o próprio SliceNew