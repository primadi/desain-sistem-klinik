# APPLICATION DESIGN DOCUMENT
## PROYEK: SISTEM INFORMASI KLINIK TERINTEGRASI SATUSEHAT
**Versi: 1.0**

### 1. STRUKTUR MENU & NAVIGASI (SITEMAP)

Sistem akan membagi menu berdasarkan Role User (RBAC).

#### A. Public / Common
*   **Login Page**: Username, Password.
*   **Dashboard**: Ringkasan data (Antrean Hari Ini, Pendapatan Hari Ini, Stok Kritis).

#### B. Front Office (Pendaftaran)
*   **Pendaftaran Pasien**:
    *   *Sub-menu: Pasien Baru (Input NIK)*
    *   *Sub-menu: Pasien Lama (Cari RM/Nama)*
*   **Daftar Kunjungan (Visit List)**: Monitor status pasien (Waiting, In-Consultation, Done).
*   **Jadwal Dokter**: View jadwal praktek.

#### C. Dokter (Medical Desk)
*   **Antrean Pasien**: List pasien yang menunggu di poli dokter tersebut.
*   **Rekam Medis (E-Medical Record)**:
    *   *History Pasien*
    *   *Input SOAP (Assessment)*
*   **E-Resep & Order**: Input resep dan permintaan lab/tindakan.

#### D. Perawat (Nurse Station)
*   **Triage / Vital Sign**: Input TTV sebelum pasien masuk ke dokter.
*   **Tindakan Keperawatan**: Input tindakan pendelegasian.


#### E. Unit Penunjang (Lab / Radiologi)
*   **Daftar Order Masuk**: List permintaan dari Dokter.
    *   *Indikator Prepaid*: "LUNAS/BELUM LUNAS". (Tidak bisa input hasil jika PREPAID & BELUM LUNAS).
*   **Input Hasil**: Form input parameter hasil pemeriksaan.

#### F. Farmasi (Apotek)
*   **E-Resep Masuk**: Queue resep dari dokter.
*   **Input Resep Manual**: Form untuk resep luar/manual.
*   **Penjualan Bebas (OTC)**: Kasir khusus obat bebas.
*   **Manajemen Stok**:
    *   *Kartu Stok*
    *   *Opname Stok*

#### G. Kasir (Billing)
*   **Tagihan Open**: List pasien yang belum lunas (Prepaid/Postpaid).
*   **Riwayat Transaksi**: Laporan pembayaran harian.

#### H. Administrator / Back Office
*   **Master Data**: User, Dokter, Poli, Jasa Medis, Data Obat/KFA.
*   **Laporan**: Kunjungan, 10 Besar Penyakit, Pendapatan, Stok Log.
*   **Settings**: Bridging SatuSehat (Client ID/Secret), Printer Config.

---

### 2. SPESIFIKASI FORM & INPUT

#### A. Form Pendaftaran Pasien Baru
*   **Input**:
    *   `NIK` (16 digit) -> *Cari SatuSehat*
    *   `Nama Lengkap` (Auto/Manual)
    *   `Tanggal Lahir` (Auto/Manual)
    *   `Jenis Kelamin` (Auto/Manual)
    *   `No HP / WhatsApp`
    *   `Alamat Domisili`
    *   `Penjamin` (Umum/BPJS/Asuransi Lain)
*   **Output**: Data Pasien Baru, Nomor Rekam Medis (Auto-generated).

#### B. Form Vital Sign (Perawat)
*   **Input**:
    *   `Sistole` / `Diastole` (mmHg)
    *   `Suhu Tubuh` (C)
    *   `Nadi` (x/menit)
    *   `Pernapasan/RR` (x/menit)
    *   `Tinggi Badan` (cm) & `Berat Badan` (kg) -> *Auto hitung IMT*
    *   `Lingkar Perut` (cm)

#### C. Form Pemeriksaan Dokter (SOAP)
*   **Subjective (S)**: Textarea keluhan pasien.
*   **Objective (O)**: Read-only Vital Sign + Textarea Pemeriksaan Fisik.
*   **Assessment (A)**:
    *   `Diagnosis Utama` (Autocomplete ICD-10) - *Wajib*
    *   `Diagnosis Sekunder` (Autocomplete ICD-10)
*   **Plan (P)**:
    *   `Tindakan` (Autocomplete ICD-9-CM / Master Tindakan)
    *   `Instruksi Medis` (Textarea)

    *   `Instruksi Medis` (Textarea)

#### D. Form Order & Input Hasil Penunjang (Lab/Rad)
*   **Input Order (Dokter)**:
    *   `Jenis Pemeriksaan` (Darah Lengkap, Rontgen Thorax, dll).
    *   `Prioritas` (Cito/Regular).
*   **Input Hasil (Petugas Lab)**:
    *   `Parameter` (HB, Leukosit, dll).
    *   `Nilai Hasil` & `Nilai Rujukan`.
    *   `Upload Lampiran` (PDF/Image) - *Optional*.

#### E. Form E-Resep
*   **Input Item**:
    *   `Nama Obat` (Search Master Obat)
    *   `Jumlah`
    *   `Signa/Aturan Pakai` (3x1, Sesudah Makan, dll)
    *   `Rute Pemberian` (Oral, Injeksi, dll)
*   **Action**: Tambah ke list -> Simpan Resep.

#### F. Form Pembayaran (Kasir)
*   **View**: Detail Tagihan (Jasa Dokter, Tindakan, Lab, Obat).
*   **Input**:
    *   `Diskon` (Nominal/Persen) - *Optional*
    *   `Metode Bayar` (Tunai/QRIS/Debit)
    *   `Jumlah Bayar`
*   **Output**: Struk/Invoice, Status Lunas.

---

### 3. ENTITY RELATIONSHIP DIAGRAMS (ERD)

Desain database dipecah menjadi beberapa modul untuk kemudahan pembacaan.

#### 3.1. Clinical Core (Pasien & Rekam Medis)

Mengatur data utama pasien, kunjungan, dan data klinis (SOAP).

```mermaid
erDiagram
    USERS {
        uuid id PK
        uuid tenant_id FK
        string username
        string full_name
        string role "ENUM: DOCTOR, NURSE, ADMIN, ETC"
        string ihs_id "SatuSehat ID"
    }

    PATIENTS {
        uuid id PK
        uuid tenant_id FK
        string nik UK
        string no_rm UK
        string full_name
        date birth_date
        string address
        string ihs_id "SatuSehat ID"
    }

    ENCOUNTERS {
        uuid id PK
        uuid tenant_id FK
        uuid patient_id FK
        uuid practitioner_id FK
        timestamp start_time
        timestamp end_time
        string status "ARRIVED, IN_PROGRESS, FINISHED"
        string ihs_id
        text subjective_complaint "S: Patient Complaint"
        text objective_findings "O: Physical Examination Narrative"
        text plan_instructions "P: Instructions/Education"
    }

    OBSERVATIONS {
        %% Vital Signs Data & Lab Results - Mandatory for SatuSehat
        uuid id PK
        uuid tenant_id FK
        uuid encounter_id FK
        string code "LOINC Code (e.g., 85354-9 for Bio Data)"
        string value "Value (e.g., 120/80)"
        string unit "Unit (e.g., mmHg)"
    }

    CONDITIONS {
        %% Diagnosis Data (ICD-10) - Mandatory for SatuSehat
        uuid id PK
        uuid tenant_id FK
        uuid encounter_id FK
        string icd10_code "Diagnosis Code (e.g., A00.1)"
        string type "PRIMARY / SECONDARY"
    }

    USERS ||--|{ ENCOUNTERS : handles
    PATIENTS ||--|{ ENCOUNTERS : has
    ENCOUNTERS ||--|{ OBSERVATIONS : contains
    ENCOUNTERS ||--|{ CONDITIONS : diagnosed_with
```

#### 3.2. Pharmacy & Inventory (Obat & Resep)

Mengelola master data obat, stok, dan resep elektronik.

```mermaid
erDiagram
    PRODUCTS {
        uuid id PK
        uuid tenant_id FK
        string name
        string kfa_code "KFA Code (Pharmacy Dictionary)"
        decimal price
        boolean is_otc
    }

    LOCATIONS {
        uuid id PK
        uuid tenant_id FK
        string name "e.g., Main Warehouse, Poly A, Poly B"
        string type "WAREHOUSE / DEPO / UNIT"
    }

    STOCK_BATCHES {
        uuid id PK
        uuid tenant_id FK
        uuid product_id FK
        uuid location_id FK
        string batch_number
        date expired_date
        int stock_qty
    }

    STOCK_MOVEMENTS {
        uuid id PK
        uuid tenant_id FK
        uuid product_id FK
        uuid from_location_id FK
        uuid to_location_id FK
        uuid reference_id "ID of Invoice/PO/Adjustment"
        string movement_type "SALE, PURCHASE, TRANSFER, ADJUSTMENT"
        int quantity
        timestamp created_at
    }

    MEDICATION_REQUESTS {
        uuid id PK
        uuid tenant_id FK
        uuid encounter_id FK
        uuid product_id FK
        int quantity
        string dosage_instruction
    }

    ENCOUNTERS ||--|{ MEDICATION_REQUESTS : prescribed
    PRODUCTS ||--|{ STOCK_BATCHES : "stored_as"
    LOCATIONS ||--|{ STOCK_BATCHES : "has_stock"
    PRODUCTS ||--|{ STOCK_MOVEMENTS : "moved"
```

#### 3.3. Billing & Services (Kasir & Penunjang)

Mengelola tagihan, pembayaran, dan order penunjang (Lab/Rad).

```mermaid
erDiagram
    INVOICES {
        uuid id PK
        uuid tenant_id FK
        uuid encounter_id FK
        decimal total_amount
        decimal discount
        decimal final_amount
        string status "DRAFT, PAID"
        timestamp created_at
    }
    
    INVOICE_ITEMS {
        uuid id PK
        uuid tenant_id FK
        uuid invoice_id FK
        string item_type "SERVICE, DRUG, LAB"
        string name
        int quantity
        decimal price
        decimal subtotal
    }

    SERVICE_REQUESTS {
        uuid id PK
        uuid tenant_id FK
        uuid encounter_id FK
        string service_code
        string status "ORDERED, PAID, IN_PROGRESS, COMPLETED"
        decimal price
    }
    
    DIAGNOSTIC_REPORTS {
        uuid id PK
        uuid tenant_id FK
        uuid request_id FK
        uuid patient_id FK "Optional: If external lab result without encounter"
        string result_value
        string attachment_url
    }

    %% Konek ke Encounter dari modul Clinical
    ENCOUNTERS {
        uuid id PK
        string payment_status
    }

    ENCOUNTERS ||--|| INVOICES : billed_as
    INVOICES ||--|{ INVOICE_ITEMS : consists_of
    ENCOUNTERS ||--|{ SERVICE_REQUESTS : orders
    SERVICE_REQUESTS ||--|{ DIAGNOSTIC_REPORTS : results_in
```

### 4. FUTURE DEVELOPMENT ROADMAP (ERP MODULES)

Sistem dirancang untuk dapat dikembangkan ke arah ERP (Enterprise Resource Planning) di masa depan:

*   **Modul Pembelian (Purchasing)**: Purchase Order (PO) obat/alkes, Penerimaan Barang.
*   **Modul Hutang (AP)**: Manajemen hutang ke supplier obat.
*   **Modul Piutang (AR)**: Manajemen tagihan asuransi/perusahaan.
*   **Modul Keuangan (Finance)**: Petty Cash, Bank Transfer.
*   **Modul Akuntansi (GL)**: Jurnal Otomatis, Buku Besar, Neraca, Laba/Rugi.

