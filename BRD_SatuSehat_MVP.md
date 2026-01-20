# BUSINESS REQUIREMENTS DOCUMENT (BRD)
## PROYEK: SISTEM INFORMASI KLINIK TERINTEGRASI SATUSEHAT (MVP)

### 1. INFORMASI UMUM

**Tujuan:**
Membangun Minimum Viable Product (MVP) Sistem Informasi Manajemen Klinik yang mengelola alur pasien dari pendaftaran hingga pelaporan, dengan fokus utama pada efisiensi operasional dan kepatuhan (compliance) terhadap standar SatuSehat Kemenkes RI.

**Target Pengguna:**
*   **Front Office (FO)/Kasir:** Pendaftaran pasien, verifikasi asuransi/pembayaran, billing.
*   **Dokter:** Anamnesa, Pemeriksaan Fisik, Diagnosa (ICD-10), Tindakan (ICD-9-CM), Resep Elektronik.
*   **Perawat/Nakes:** Vital sign, triage, tindakan keperawatan.
*   **Unit Penunjang (Lab/Radiologi):** Input hasil pemeriksaan.
*   **Farmasi/Apoteker:** Verifikasi resep, dispensing obat, update stok.
*   **Administrator:** Manajemen user, master data.

### 2. CAKUPAN SISTEM (SCOPE - MVP)

Fokus pada *Core Features* untuk operasional dasar klinik rawat jalan:
1.  **SatuSehat Bridging:** Sinkronisasi data pasien dan rekam medis.
2.  **Rawat Jalan (Outpatient):** Poli Umum/Spesialis.
3.  **Rekam Medis Elektronik (RME):** Standar SOAP & coding diagnosis/tindakan.
4.  **Farmasi Dasar:** E-Prescribing & Dispensing (Stok sederhana).
5.  **Billing & Kasir:** Invoice simple & Payment.

### 3. ALUR KERJA SISTEM (WORKFLOW) & DIAGRAM

#### 3.1 Flowchart Registrasi Pasien (Integrasi SatuSehat)
```mermaid
graph TD
    A[Pasien Datang] --> B{Pasien Baru/Lama?}
    B -- Baru --> C[Input NIK]
    C --> D[Request API SatuSehat (Patient by NIK)]
    D --> E{Data Ditemukan?}
    E -- Ya --> F[Auto-Fill Data Identitas]
    E -- Tidak --> G[Input Manual & Sync ke SatuSehat]
    B -- Lama --> H[Cari Nama/RM/NIK Lokal]
    F --> I[Pilih Poli & Dokter]
    G --> I
    H --> I
    I --> J[Generate No. Antrean]
    J --> K[Status: Menunggu di Nurse Station]
```

#### 3.2 Flowchart Pelayanan Medis & Pembayaran (Prepaid vs Postpaid)

*   **Prepaid (Bayar Dulu):** Item tindakan/lab tertentu wajib dilunasi sebelum dilakukan. Sistem akan memblokir input hasil jika belum lunas.
*   **Postpaid (Bayar Belakangan):** Tagihan ditambahkan setelah tindakan dilakukan (biasanya untuk BHP/Tindakan insidental).

```mermaid
graph TD
    A[Antrean Nurse Station] --> B[Perawat Input Vital Sign]
    B --> C[Status: Menunggu Dokter]
    C --> D[Dokter: Anamnesa & Pemfis (S & O)]
    D --> E[Dokter: Diagnosis ICD-10 (A)]
    E --> F{Order Penunjang?}
    F -- Ya -> Lab/Rad --> G{Tipe Bayar?}
    G -- Prepaid --> H[Kasir: Bayar Dulu]
    H --> I[Petugas Lab: Input Hasil]
    G -- Postpaid --> I
    F -- Tidak --> J[Dokter: Terapi/Resep & Tindakan (P)]
    J --> K[Finalisasi Resume Medis (Encounter Finish)]
    I --> K
    K --> L[Kirim Data 'Encounter' & 'Condition' ke SatuSehat]
    L --> M[Status: Menunggu Farmasi/Kasir]
```

### 4. KEBUTUHAN FUNGSIONAL (DETAILED)

#### A. Modul Pendaftaran & Antrean (Front Office)
1.  **Pencarian Pasien (KYC):**
    *   Input NIK -> GET `/Patient` (SatuSehat).
    *   Fallback input manual jika NIK tidak ditemukan/offline.
2.  **Registrasi Kunjungan:**
    *   Pilih Dokter, Poli, dan Penjamin (Umum/Asuransi/BPJS).
    *   Cetak tiket antrean.
3.  **Dashboard Antrean:**
    *   Tampilan antrean per poli untuk layar TV ruang tunggu.

#### B. Modul Rekam Medis Elektronik (RME) - Dokter & Perawat
1.  **Asesmen Awal (Perawat):**
    *   Input Tanda Vital (Tensi, Nadi, Suhu, RR, TB, BB).
    *   Auto-map ke resource `Observation` (SatuSehat).
2.  **SOAP Dokter:**
    *   **Subjective:** Keluhan utama, RPS, RPD.
    *   **Objective:** Hasil pemeriksaan fisik.
    *   **Assessment:** Searchable ICD-10 database. (Auto-map ke `Condition` SatuSehat).
    *   **Plan:**
        *   **Tindakan:** Searchable ICD-9-CM database. (Auto-map ke `Procedure` SatuSehat).
        *   **Resep:** Searchable KFA (Kamus Farmasi & Alkes).
3.  **Riwayat Medis:**
    *   View timeline kunjungan sebelumnya.

#### C. Modul Penunjang (Lab & Tindakan)
1.  **Order List:** Notifikasi order baru dari dokter.
2.  **Input Hasil:** Template hasil sederhana.
3.  **Billing Trigger:** Tindakan otomatis menambah tagihan.

#### D. Modul Farmasi
1.  **E-Resep Queue:** Daftar resep elektronik masuk dari Poli.
2.  **Input Resep Manual:** Fitur bagi Dokter/Apoteker untuk input resep manual jika sistem down atau resep luar.
3.  **Penjualan Bebas (OTC & Resep Umum):**
    *   Penjualan obat bebas (Over The Counter) tanpa resep dokter klinik.
    *   Layanan resep dari luar klinik (Resep Umum).
4.  **Verifikasi & Etiket:**
    *   Cek ketersediaan stok.
    *   Cetak label obat (Aturan pakai, Nama Pasien).
5.  **Dispensing:** Konfirmasi obat diserahkan -> Kurangi stok -> Trigger `MedicationDispense` (SatuSehat).

#### E. Modul Kasir
1.  **Invoice Generation:** Gabungan Jasa Medis + Obat + Tindakan + Admin.
2.  **Pembayaran:** Split payment (Cash/QRIS/Card).
3.  **Closing:** Laporan pendapatan harian (Shift).

### 5. PERSYARATAN INTEGRASI SATUSEHAT (TECHNICAL MAPPING)

| Resource SatuSehat | Trigger Event di Sistem Lokal | Data Wajib (Mandatory) |
| :--- | :--- | :--- |
| **Patient** | Pendaftaran Pasien Baru | NIK, Nama, Tgl Lahir, Gender |
| **Encounter** | Check-in (Start) & Check-out (End) | ID Pasien, ID Praktisi (Nakes), Lokasi, Waktu |
| **Observation** | Simpan Vital Sign | Code LOINC (TTV), Value, Unit |
| **Condition** | Simpan Diagnosis (Dokter) | Code ICD-10, Category (Encounter Diagnosis) |
| **MedicationRequest** | Simpan Resep (Dokter) | Code KFA, Dosis, Rute, Instruksi |
| **Procedure** | Simpan Tindakan (Dokter/Perawat) | Code ICD-9-CM |

*Catatan: Token API & Mapping ID Lokasi/Praktisi (IHS ID) harus dikonfigurasi di menu Settings.*

### 6. MATRIKS HAK AKSES (RBAC)

| Role | Pendaftaran | RME (SOAP) | Resep (Input) | Resep (Proses) | Kasir | Laporan |
| :--- | :---: | :---: | :---: | :---: | :---: | :---: |
| **Front Office** | ✅ | ❌ | ❌ | ❌ | ✅ | ✅ (Basic) |
| **Dokter** | 👀 (View) | ✅ | ✅ | ❌ | ❌ | ✅ (Jasa Medis) |
| **Perawat** | 👀 | ✅ (S,O) | ❌ | ❌ | ❌ | ❌ |
| **Apoteker** | ❌ | 👀 | ❌ | ✅ | ❌ | ✅ (Obat) |
| **Admin** | ✅ | ✅ | ✅ | ✅ | ✅ | ✅ (Full) |

### 7. ARSITEKTUR & TECH STACK (REKOMENDASI)

**Modern Web Application (Monorepo/Modular):**
*   **Frontend:** Next.js (React) - Performance & SEO friendly.
    *   UI Framework: Tailwind CSS + ShadcnUI (Clean, Modern, Medical aesthetic).
### 7. ARSITEKTUR & TECH STACK (REKOMENDASI)

**Modern Web Application (Monorepo/Modular):**
*   **Frontend:** Next.js (React) - Performance & SEO friendly.
    *   UI Framework: Tailwind CSS + ShadcnUI (Clean, Modern, Medical aesthetic).
*   **Backend:** Golang.
    *   Framework: **Lokstra**.
*   **Database:** PostgreSQL (Relational, JSONB support untuk FHIR data).
*   **Authentication:** Custom Auth Module (Golang).
*   **Hosting:** Docker Containerization.

### 8. RENCANA PENGEMBANGAN (ROADMAP MVP)

1.  **Sprint 1:** Setup Project, Database Schema, Master Data (Dokter, Poli, Obat/KFA, ICD-10).
2.  **Sprint 2:** Modul Pendaftaran & Integrasi SatuSehat (Patient & Encounter).
3.  **Sprint 3:** Modul RME (SOAP) & Integrasi SatuSehat (Condition & Observation).
4.  **Sprint 4:** E-Resep, Farmasi & Integrasi SatuSehat (Medication).
5.  **Sprint 5:** Kasir & Billing.
6.  **Sprint 6:** UAT & Deployment.

---
**Perserujuan:**
Dokumen ini disepakati sebagai acuan pengembangan MVP Sistem Klinik.
