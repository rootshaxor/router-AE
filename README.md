# router-AE

<p align="center">
<img width="256" height="256" src="https://raw.githubusercontent.com/rootshaxor/router-AE/master/img/router-AE.png"> 
</p>

  Menjadikan linux sebagai router , kini lebih mudah dengan menggunakan router-AE,  
router-AE dapat dijadikan sebagai shell di linux (tidak disarankan untuk v0.0.1-alpha).
Namun di versi sekarang masih dalam tahap pengembangan, di versi saat ini hanya memiliki
sedikit fitur. Untuk menjalankan router-AE harus dengan `superuser`


### Configure
Download versi binary [disini](https://github.com/rootshaxor/router-AE/releases)

- Menjadikan router-AE sebagai shell di linux (`tidak disarankan untuk v0.0.1-alpha`)
```shell
#!/bin/bash
mkdir ~/bin
wget https://github.com/rootshaxor/router-AE/releases/download/alpha/Router-AE-alpha-0.0.1-x86_64-linux.zip

unzip Router-AE-alpha-0.0.1-x86_64-linux.zip -d ~/bin

#Add `~/bin` to $PATH
export PATH=~/bin:$PATH

#Change Shell, pilih shell di ~/bin/router-AE
chsh 

## Ini Tidak Disarankan Untuk v0.0.1-alpha (saat ini) ##

```

- Manual
```shell
#!/bin/bash
mkdir ~/bin
wget https://github.com/rootshaxor/router-AE/releases/download/alpha/Router-AE-alpha-0.0.1-x86_64-linux.zip

unzip Router-AE-alpha-0.0.1-x86_64-linux.zip -d ~/bin

#Add `~/bin` to $PATH
export PATH=~/bin:$PATH

#Menjalakan router-AE
router-AE
```

### Screenshot
Lihat Screenshot router-AE [disini](https://github.com/rootshaxor/router-AE/blob/master/img/Screenshost/Screenshot.md)

### Lisensi 
[lisensi](https://github.com/rootshaxor/router-AE/blob/master/LICENSE)

### Catatan
[catatan](https://github.com/rootshaxor/router-AE/blob/master/Note.md)
