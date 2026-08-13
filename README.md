# Workable Parser

Di challenge pertama kamu sudah berhasil untuk melakukan parsing data pada Grab Career Page.

Cuma memang disana, semua data yang mau kamu parsing sudah dikirim ketika HTTP client kamu request data HTML-nya menggunakan `http.Get()`, jadi kamu bisa langsung saja parsing data dari HTML yang dikirimkan tersebut (data HTML ini disebut juga sebagai `Page Source`).

Nah tapi ada kalanya juga dimana data yang kamu perlukan baru akan dikirimkan melalui request dari javascript yang akan di eksekusi setelah halaman web-nya selesai di load.

Nah, untuk kasus seperti ini kamu perlu trik yang lain. Ada 2 pendekatan yang bisa kamu lakukan:

1. Analisis javascript request yang membawa data yang kamu perlukan.
2. Jalankan web page-nya menggunakan 'headless browser' lalu parse HTML yang sudah di load secara sempurna oleh browsernya.

Di challenge ini kamu akan belajar untuk mengimplementasikan kedua approach ini dengan menggunakan Go.

Jadi di challenge ini kamu harus membuat 2 program yang melakukan hal berikut:

1. Program membaca file [urls.txt](./urls.txt).
2. Program membaca setiap url yang ada di file tersebut dan melakukan parsing. Satu program melakukan parsing dengan menggunakan metode javascript request & satu lagi menggunakan metode headless browser. Data yang harus di parse adalah sebagai berikut:
	- `title` => Title dari role yang ada di vacancy.
	- `workplace` => `Remote`, `Onsite`, atau `Hybrid`.
	- `job_type` => `Full time`, `Part time`, `Contract`, atau yang lainnya.
	- `location` => Lokasi dimana target applicant berada. Jika ada lebih dari satu, ambil value yang paling pertama.
	- `description` => Seluruh deskripsi tentang vacancy-nya.
	- `page_url` => URL menuju page ini.
3. Seluruh hasil parsing dijadikan satu file `*.csv` dengan format nama `[YYYY-MM-DD].csv` misalnya saja: `2026-08-06.csv`.

Oh ya, di challenge kali ini juga kamu perlu dockerize program Go kamu.

Kalau ada pertanyaan, jangan sungkan-sungkan untuk tanya ke mentor kamu ya. 😁