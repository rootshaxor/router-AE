# router-AE

### Simple Shell to use LINUX for router 

  Membuat sheel linux menjadi sebagai shell untuk mengatur jaringan
ini terinspirasi dari mikrotik OS & cisco IOS, pada tahap ini router-AE
masih pada versi alpa 0.0.1, fiture yang dimiliki masih belum lengkap.
router-AE harus dijalankan dengan superuser.
### Untuk Penggunaan 

  penggunaan bisa digunakan sebagai shell di linux (tidak di sarankan)
lebih di utamakan untuk penggunaan secara manual pada versi ini

- Shell 
```shell
mkdir ~/bin
go build -o router main.go
mv router ~/bin
export PATH=~/bin:$PATH
chsh

#[pilih shell pada direktory yang telah dibuat `~/bin/router`]

```
- Manual
```shell
mkdir ~/bin
go build -o router main.go
mv router ~/bin
export PATH=~/bin:$PATH
router
```
screenshot dapat dilihat [disini](https://github.com/rootshaxor/router-AE/blob/master/Screenshost/Screenshot.md)
