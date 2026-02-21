# Go-Redis

Bu proje, Go dili kullanılarak Redis protokolünü ve temel işleyişini anlamak amacıyla geliştirilen "mini-redis" çalışmasıdır.

## Şu Ana Kadar Neler Yaptık?

Projenin başlangıcından bu yana aşağıdaki adımları başarıyla tamamladık:

### 1. TCP Sunucu Altyapısı

Uygulamanın temelini oluşturmak için Go'nun `net` paketini kullanarak düşük seviyeli bir TCP sunucusu kurduk.

- **Port Yapılandırması:** Sunucuyu `:6380` portu üzerinden dinlemeye açtık.
- **Bağlantı Dinleme:** Sunucu, istemciden gelen bağlantıları (`ln.Accept()`) bekleyecek şekilde yapılandırıldı.

### 2. İletişim ve Veri Okuma

İstemci ile sunucu arasındaki veri alışverişini yönettik.

- **Buffered Reader:** `bufio` kullanarak bağlantı üzerinden gelen verileri daha performanslı ve rahat bir şekilde okuduk.
- **Mesaj İşleme:** İstemciden gelen satırları okuduk, etrafındaki boşlukları temizledik ve logladık.
- **Yanıt Dönme:** Her başarılı mesaj alımında istemciye Redis protokolüne benzer şekilde `+OK\r\n` yanıtını gönderdik.

### 3. Bağlantı Yönetimi

- **Eşzamanlılık (Goroutines):** Her yeni istemci bağlantısını ayrı bir `goroutine` içinde (`go handleConnection`) işleyerek sunucunun aynı anda birden fazla istemciye hizmet vermesini sağladık.
- **Güvenli Kapatma:** İstemci bağlantıyı kestiğinde (`io.EOF`), kaynakları temizlemek için bağlantıyı düzgünce kapattık.

### 4. Mimari ve Modüler Yapı

Kodun sürdürülebilir olması için projeyi parçalara ayırdık:

- **Config Katmanı:** Ayarların (Port vb.) tek merkezden yönetilmesi sağlandı.
- **Server Katmanı:** TCP sunucu mantığı `internal/server` paketine taşınarak `main.go` sadeleştirildi.

### 5. Geliştirme Ortamı ve GitHub

- **Otomatik Yenileme:** `air` kurulumu ile kod değişikliklerinin anında sunucuya yansımasını sağladık.
- **Versiyon Kontrolü:** `.gitignore` dosyamızı hazırladık ve tüm süreci GitHub reponuza pushladık.

---

## Nasıl Çalıştırılır?

1. Bağımlılıkları yükleyin:
   ```bash
   go mod tidy
   ```
2. Uygulamayı çalıştırın:
   ```bash
   go run cmd/miniredis-server/main.go
   # VEYA air yüklüyse sadece:
   air
   ```
3. Başka bir terminalden test edin:
   ```bash
   nc 127.0.0.1 6380
   # Yazı yazıp Enter'a basın
   ```
