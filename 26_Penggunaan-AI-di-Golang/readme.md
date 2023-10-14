Beberapa alasan memilih Go:
- Dikembangkan oleh Google untuk pengembangan aplikasi server-side dan berbasis cloud
- Go adalah compiled language sehingga lebih cepat dibanding interpreted language (contoh: Python)
- Mendukung concurrenry melalui fitur goroutines sehingga dapat menjalankan banyak task secara simultan tanpa membebani memori
- Memiliki standard libraries yang sangat kaya, didukung oleh komunitas yang kuat
- Termasuk libraries untuk mendukung pengembangan AI

Library Go untuk machine learning adalah 
1. GoLearn:
    - Bersifat open source
    - 'batteries included' machine learning
    - Simplicity & customizability
    - Contoh penerapan: https://github.com/sjwhitworth/golearn#getting-started
2. Gorgonia:
    - Memudahkan menulis dan mengevaluasi persamaan matematika yang melibatkan array multidimensi
    - Low-level library, seperti Theano, namun lebih tinggi dibanding Tensorflow
    - Contoh penerapan: https://github.com/gorgonia/gorgonia#usage
3. gonum:
    - Sebuah set packages yang didesain untuk menulis numerical & algoritma sains menjadi produktif, memiliki performa tinggi, dan scalable
    - Mirip seperti numpy dan scipy, library yang dibuat menggunakan Python
    - Referensi: https://www.gonum.org/post/intro_to_gonum/
4. gomind:
    - Contoh penggunaan: https://github.com/surenderthakran/gomind#usage

Terdapat cara yang lebih mudah untuk mengembangkan aplikasi berbasis AI, yaitu menggunakan teknologi bernama AI as a Service (AIaaS). AIaaS adalah sebuah tools AI yang dapat segera digunakan (pre-build or off-the-shelf AI tools). AIaaS berguna bagi bisnis yang ingin menerapkan AI tanpa merekrut ahlinya dan tanpa mengeluarkan biaya yang relatif banyak. Beberapa perusahaan penyedia AIaaS, contohnya:
- Amazon Web Service (AWS)
- Microsoft Azure
- Google Cloud Platform (GCP)
- IBM
- OpenAI

Tipe-tipe AIaas:
1. Bots/Chatbots: membuat AI-Based Customer Services
    - Amazon Lex (AWS), Azure Bot Service (Microsoft Azure), DialogFlow (GCP)
2. APIs: mengintegrasikan AI milik vendor dengan aplikasi sendiri
    - Amazon Rekognition, Amazon Comprehend
    - Azure Cognitive Service, Azure Speech Services
    - Google Cloud Vision, Cloud Natural Language
3. Machine Learning: build & deploy ML model easily
    - Amazon SageMaker
    - Azure Machine Learning
    - Google Cloud AI

Keuntungan menggunaan AIaaS:
- Pengurangan Cost
- Low-code
- Proses deployment cepat
- Flexibility
- Usability
- Scalability
- Customization

Kelemahan menggunakan AIaaS:
- Cost yang berlebih dalam jangka panjang
- Privasi data
- Keamanan
- Transparansi
- Vendor lock-in

Prompt engineering: Panduan penulisan prompt:
1. Gunakan prompt yang jelas & spesifik
    - Gunakan delimiters (contoh:```, ''', <>, <tag></tag>) untuk menandakan bagian mana yang menjadi input
    - Tuliskan struktur output yang diinginkan
    - Minta model untuk memeriksa apakah input sudah sesuai
    - Few-shot prompting
2. Berikan "waktu berpikir" untuk menghindari solusi yang salah
    - Menulis langkah-langkah yang dibutuhkan untuk menyelesaikan tugas dan output yang diinginkan
    - Menginstruksikan model untuk menuliskan solusinya sendiri sebelum masuk ke kesimpulan

