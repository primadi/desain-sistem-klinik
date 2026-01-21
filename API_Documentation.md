# API DOCUMENTATION
## PROYEK: SISTEM INFORMASI KLINIK TERINTEGRASI SATUSEHAT
**Versi: 1.0**

Dokumentasi ini mendefinisikan kontrak API antara Frontend (Next.js) dan Backend (Golang). Format response standar menggunakan JSON.

### 1. COMMON & AUTHENTICATION

#### Login
*   **Endpoint**: `POST /auth/login`
*   **Request**:
    ```json
    {
      "username": "dr_budi",
      "password": "secure_password"
    }
    ```
*   **Response**:
    ```json
    {
      "token": "jwt_token_string",
      "user": {
        "id": "uuid",
        "role": "DOCTOR",
        "full_name": "Dr. Budi Santoso",
        "tenant_id": "uuid"
      }
    }
    ```

#### Master Data (Lookup)
*   **Endpoint**: `GET /master/{type}`
*   **Params**: `type` (poli, doctors, products, icd10, kfa)
*   **Response**: List of master data items.

---

### 2. PATIENT MANAGEMENT

#### Search Patient
*   **Endpoint**: `GET /patients`
*   **Query**: `?q=Budi` (Search by Name/RM/NIK)
*   **Response**: `[ { "id": "...", "full_name": "...", "no_rm": "..." } ]`

#### Lookup NIK (SatuSehat Bridge)
*   **Endpoint**: `GET /patients/lookup-nik`
*   **Query**: `?nik=1234567890123456`
*   **Response**: Returns patient data from SatuSehat.

#### Register Patient
*   **Endpoint**: `POST /patients`
*   **Request**:
    ```json
    {
      "nik": "1234567890",
      "full_name": "Budi Santoso",
      "birth_date": "1990-01-01",
      "gender": "male",
      "address": "Jl. Merdeka No. 1"
    }
    ```

---

### 3. CLINICAL & ENCOUNTERS

#### Get Queue / Visit List
*   **Endpoint**: `GET /encounters`
*   **Query**: `?status=ARRIVED` or `?practitioner_id=...`
*   **Response**: List of active visits.

#### Start Encounter (Check-in)
*   **Endpoint**: `POST /encounters`
*   **Request**:
    ```json
    {
      "patient_id": "uuid",
      "practitioner_id": "uuid",
      "start_time": "2024-01-21T08:00:00Z"
    }
    ```

#### Update Clinical Data (SOAP)
*   **Endpoint**: `PUT /encounters/{id}/soap`
*   **Request**:
    ```json
    {
      "subjective_complaint": "Demam 3 hari",
      "objective_findings": "Suhu 38.5C",
      "conditions": [ { "icd10_code": "A00.1", "type": "PRIMARY" } ],
      "observations": [ { "code": "8310-5", "value": "38.5", "unit": "C" } ],
      "plan_instructions": "Istirahat cukup"
    }
    ```

#### Finish Encounter (Check-out)
*   **Endpoint**: `POST /encounters/{id}/finish`
*   **Description**: Finalizes the medical record and triggers SatuSehat sync.

---

### 4. PHARMACY & INVENTORY

#### Get Product List
*   **Endpoint**: `GET /products`
*   **Query**: `?search=Paracetamol`
*   **Response**: Product list with current stock across locations.

#### Record Stock Mutation
*   **Endpoint**: `POST /inventory/mutations`
*   **Request**:
    ```json
    {
      "product_id": "uuid",
      "from_location_id": "uuid",
      "to_location_id": "uuid",
      "type": "TRANSFER",
      "quantity": 10,
      "batch_number": "BATCH123"
    }
    ```

#### Create Prescription
*   **Endpoint**: `POST /medication-requests`
*   **Request**:
    ```json
    {
      "encounter_id": "uuid",
      "items": [
        { "product_id": "uuid", "quantity": 10, "dosage": "3x1 tab" }
      ]
    }
    ```

---

### 5. LAB & RADIOLOGY

#### Order Lab Service
*   **Endpoint**: `POST /service-requests`
*   **Request**:
    ```json
    {
      "encounter_id": "uuid",
      "service_code": "LAB_DARAH_LENGKAP",
      "notes": "Cito"
    }
    ```

#### Submit Result
*   **Endpoint**: `POST /diagnostic-reports`
*   **Request**:
    ```json
    {
      "request_id": "uuid",
      "result_value": "HB: 12.0",
      "attachment_file": "(Multipart/Form-Data Analysis)"
    }
    ```

---

### 6. BILLING

#### Get Invoice
*   **Endpoint**: `GET /invoices`
*   **Query**: `?encounter_id=...`

#### Process Payment
*   **Endpoint**: `POST /invoices/{id}/pay`
*   **Request**:
    ```json
    {
      "payment_method": "CASH",
      "amount_paid": 150000
    }
    ```
