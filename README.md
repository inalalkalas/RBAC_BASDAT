#### cara menjalankannya 
jalankan kode database terlebih dahulu kemudian jalankan codenya
untuk menjalankan database via terminal/sql shell
jika terdapat tulisan ...JS maka masuk/connect terlebih daulu ke sql dengan cara "\sql"
kemudian connect kan ke localhostdengan cara "\connec 127.0.0.1:3306"
dan kemudian jalankan file datnases tersebut dengan command "source Path/file/database/ini"

jika lewat xampp/laragon maka aktifkan terlebih dahulu aplikasi tersebut kemudian jalankan kode sqlnya pada aplikasi tersebut.

membuka 2 terminal/cmd
terminal ke satu untuk menjalank golang tersebut dengan cara "go run main.go"

terminal ke-2 untuk menjalankan curl atau thunder clients pada vscode
jika menjalankan menggunakan curel maka "curl "http://localhost:8080/items/3?user_id=3"" atau "curl -X GET "http://localhost:8080/items/0?user_id=1""

#### Apa saja yang perlu diganti
yang perlu diganti adalah password untuk koneksi ke database diganti dengan pssword milik anda sendiri.
letak koneksi database berad di direktori config/
