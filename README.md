# Golang Egitim Notlari

| Topic | Keywords |
| ------ | ------ |
| Variable | const, var|
| Array | array, make, len|
| Slice | slices, make, len, cap|
| Map | len|
| String | Sprintf, HasPrefix, ToUpper |
| Loops | range |
| Structs | struct, string, json |
| Pointers | referance type  |
| Functions | defer, panic, recovery |

## array
golang' da array"ler limitlidir, boyutlari bellidir. Slice"larin ise boyutlari genisleyebilir.

## slice
* append fonksiyonu ile slice'e eleman ekleriz.
* cap fonksiyonu slice'in kapasitesini verir.

## map
* memory'de key value degerler tutmamizi saglar.

## strings
* strings degerleri tutmak icin kullanilir.

## loops
* range ile dizi icindeki elemanlara dongu ile ulasilabilir.

## structs
* OOP icin nesneler ile calismamizi saglar.

## pointers
* Pointer degiskenlerin adresidir. Yani degerin saklandigi konumdur.

## functions
* bir veya birden fazla parametre alarak cesitli gorevleri yaptiktan sonra geriye deger veya degerler donderen kod bloklaridir.
* go'da func'lar **birden fazla** deger dondurebilmektedir.
* tanimli degiskenleri sadece return yazarak return edebiliyoruz.
* **defer** ile tanimladigimiz islemler fonksiyonun isi bittikten sonra calisir. (En son tanimlanan defer ilk calisir)
* **panic** ile hata durumunda program calismaz.
* **defer** tanimlari fonksiyonlarin basinda tanimlanmasi daha uygundur
* fonksiyon icindeki panic'leri go ile builtin gelen **recovery** fonksiyonu ile kontrol edebiliriz.