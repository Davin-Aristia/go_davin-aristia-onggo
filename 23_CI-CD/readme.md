Continuous Integration merupakan proses otomatis yang dilakukan untuk integrasi berbagai kode dari berbagai potential sources dengan tujuan untuk di build atau di test

Continuous Deployment/Continuous Delivery merupakan proses yang melangkah lebih jauh daripada integrasi dan proses delivery. Pipeline dari continuous deployment secara otomatis deploy setiap build yang telah diverifikasi
Cycle Continuous Delivery dan Continuous Deployment hampir sama yaitu:
    Unit Test -> Platform Test -> Deliver to Staging -> Application Acceptance Tests -> Deploy to Production -> Post deploy tests
Bedanya adalah setiap step dilakukan secara otomatis kecuali pada Continuous Delivery pada step ke Application Acceptance Tests -> Deploy to Production yang merupakan manual


Kubernetes merupakan open-source system untuk deployment otomatis, scaling, dan pengelolaan aplikasi dalam container. Kubernetes bisa melakukan hal seperti:
- Service discovery dan load balancing
- Horizontal scaling
- Automated rollouts and rollbacks
- Secret and configuration management

What Kubernetes is not:
- Does not limit the types of applications supported
- Does not deploy source code and does not build your application
- Does not provide application-level services, such as middleware
- Does not dictate logging, monitoring, or alerting solutions

1. Namespace menyediakan cakupan sumber daya kubernetes, membagi cluster menjadi unit-unit yang lebih kecil. 
2. Pod merupakan koleksi dari container yang membagi network dan mount namespaces dan merupakan basic unit dari deployment di Kuberneter. 
3. Label merupakan mekanisme yang digunakan untuk mengorganisasi objek Kubernetes
4. Deployment merupakan supervisor untuk pods, memberikan fine-grained control terhadap bagaimana sebuah versi pod baru roll out maupun roll back ke keadaan sebelumnya
5. Rolling Updates mengizinkan deployment update untuk terjadi dengan zero downtime dengan cara secara bertahap update instansi pod dengan yang baru. Pod yang baru akan dijadwalkan dengan node dengan resource yang tersedia
6. Secrets memberikan mekanisme untuk menggunakan informasi secara aman dan dapat diandalkan
7. Service merupakan abstraksi dari pods, menyediakan kandang dinamakan IP(VIP) address
8. Ingress merupakan objek API yang mengelola akses external ke service di dalam cluster, khususnya HTTP
9. DNS adalah sebuah sistem yang menghubungkan Uniform Resource Locator (URL) dengan Internet Protocal Address (IP Address)

