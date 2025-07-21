# BFF-Sandbox (Backend For Frontend Sandbox)

`Made By Mukho`

⭐ 2025-07-21 Mon

## 📋 프로젝트 개요

BFF-Sandbox는 마이크로서비스 아키텍처를 시뮬레이션하는 샌드박스 프로젝트입니다. Backend For Frontend (BFF) 패턴을 사용하여 여러 마이크로서비스를 통합하고, 클라이언트에게 단일 API 엔드포인트를 제공합니다.

## 🎯 주요 특징

1. **마이크로서비스 아키텍처**: 각 도메인별로 독립적인 서비스 구성
2. **다양한 기술 스택**: Node.js, Python, Ruby, Go를 사용한 폴리글랏 아키텍처
3. **BFF 패턴**: 클라이언트 친화적인 단일 API 제공
4. **RESTful API**: 표준 REST API 설계
5. **데이터 통합**: 여러 서비스의 데이터를 조합하여 통합 응답 제공

## 🔧 개발 환경

- **IDE**: Visual Studio Code
- **API 테스트**: curl, Postman
- **버전 관리**: Git

### 🛠️ 기술 스택

#### 🔥 BFF Service

- **언어**: JavaScript (Node.js)
- **프레임워크**: Koa.js
- **HTTP 클라이언트**: Axios
- **런타임**: Node.js v14+

#### Customer Service

- **언어**: Python
- **프레임워크**: Django
- **데이터베이스**: JSON 파일
- **런타임**: Python v3.8+

#### Account Service

- **언어**: Ruby
- **프레임워크**: Ruby on Rails
- **데이터베이스**: JSON 파일
- **런타임**: Ruby v3.0+

#### Transaction Service

- **언어**: Go
- **프레임워크**: Gin
- **데이터베이스**: JSON 파일
- **런타임**: Go v1.19+

## 🏗️ 아키텍처

```
┌─────────────────┐      ┌──────────────────────────────────┐
│    Frontend     │─────▶│           BFF Service            │
│    (Client)     │◀─────│          (Node.js/Koa)           │
└─────────────────┘      │           Port: 4000             │
                         └──────────────────────────────────┘
                            │    ▲      │   ▲      │     ▲
                            │    │      │   │      │     │
                            ▼    │      ▼   │      ▼     │
                         ┌──────────┐ ┌───────┐ ┌───────────┐
                         │ Customer │ │Account│ │Transaction│
                         │ Service  │ │Service│ │  Service  │
                         │ (Django) │ │(Rails)│ │  (Go/Gin) │
                         │Port: 3001│ │Port:  │ │Port: 3003 │
                         │          │ │ 3002  │ │           │
                         └──────────┘ └───────┘ └───────────┘
```

## 📂 학습 문서

[BFF-Sandbox](https://boom-dead-1ee.notion.site/BFF-Sandbox-237ba2b7e4b580b2b019d2cc79075e31?source=copy_link) - Notion 문서에 정리했습니다.

## 🚀 서비스 구성

### 1. 🔥 BFF Service (bs-bff)

- **기술 스택**: Node.js, Koa.js
- **포트**: 4000
- **역할**: 클라이언트와 마이크로서비스 간의 중계 역할
- **주요 기능**:
  - 고객 정보, 계좌 정보, 거래 내역을 통합하여 단일 응답 제공
  - API Gateway 역할 수행

### 2. Customer Service (bs-customer)

- **기술 스택**: Python, Django
- **포트**: 3001
- **역할**: 고객 정보 관리
- **API 엔드포인트**:
  - `GET /customer/{id}` - 특정 고객 정보 조회

### 3. Account Service (bs-account)

- **기술 스택**: Ruby, Ruby on Rails
- **포트**: 3002
- **역할**: 계좌 정보 관리
- **API 엔드포인트**:
  - `GET /account/{id}` - 특정 계좌 정보 조회
  - `GET /accounts/customer/{customer_id}` - 특정 고객의 모든 계좌 조회

### 4. Transaction Service (bs-transaction)

- **기술 스택**: Go, Gin
- **포트**: 3003
- **역할**: 거래 내역 관리
- **API 엔드포인트**:
  - `GET /transactions/account/{account_id}` - 특정 계좌의 거래 내역 조회

## 📡 BFF API

### GET /bank/{customer_id}

고객의 통합 정보를 반환합니다.

**요청 예시**:

```bash
GET /bank/1
```

**응답 예시**:

```json
{
  "customer": {
    "id": 1,
    "name": "김철수",
    "birthDate": "1999-05-16",
    "address": "서울특별시 강남구",
    "created_at": "2024-01-01T10:00:00Z",
    "updated_at": "2024-06-01T12:00:00Z"
  },
  "accounts": [
    {
      "id": 1,
      "customer_id": 1,
      "name": "급여통장",
      "type": "CHECKING",
      "balance": 1800000,
      "account_no": "110-123-456789",
      "currency": "KRW",
      "created_at": "2024-01-02T10:00:00Z",
      "updated_at": "2024-07-10T09:00:00Z",
      "transactions": [
        {
          "id": 4,
          "account_id": 1,
          "type": "IN",
          "amount": 3435604,
          "balance_after": 3435604,
          "transaction_at": "2025-06-21T02:52:43.780456Z",
          "description": "회사 월급",
          "counter_account_no": "987-654-321000",
          "created_at": "2025-06-21T02:52:43.780456Z",
          "updated_at": "2025-06-21T02:52:43.780456Z"
        },
        {
          "id": 17,
          "account_id": 1,
          "type": "OUT",
          "amount": 3008105,
          "balance_after": 8116788,
          "transaction_at": "2025-06-26T05:35:43.780583Z",
          "description": "자동차 할부",
          "counter_account_no": "",
          "created_at": "2025-06-26T05:35:43.780583Z",
          "updated_at": "2025-06-26T05:35:43.780583Z"
        },
        {
          "id": 22,
          "account_id": 1,
          "type": "OUT",
          "amount": 213453,
          "balance_after": 8847901,
          "transaction_at": "2025-06-21T12:17:43.780641Z",
          "description": "렌트비 이체",
          "counter_account_no": "BANK-123-456789",
          "created_at": "2025-06-21T12:17:43.780641Z",
          "updated_at": "2025-06-21T12:17:43.780641Z"
        },
        {
          "id": 62,
          "account_id": 1,
          "type": "IN",
          "amount": 3706563,
          "balance_after": 3706563,
          "transaction_at": "2025-06-24T17:01:43.781013Z",
          "description": "급여 입금",
          "counter_account_no": "222-333-444555",
          "created_at": "2025-06-24T17:01:43.781013Z",
          "updated_at": "2025-06-24T17:01:43.781013Z"
        },
        {
          "id": 95,
          "account_id": 1,
          "type": "OUT",
          "amount": 1129705,
          "balance_after": 5353514,
          "transaction_at": "2025-07-13T05:22:43.781280Z",
          "description": "렌트비 이체",
          "counter_account_no": "BANK-123-456789",
          "created_at": "2025-07-13T05:22:43.781280Z",
          "updated_at": "2025-07-13T05:22:43.781280Z"
        }
      ]
    },
    {
      "id": 2,
      "customer_id": 1,
      "name": "저축통장",
      "type": "SAVINGS",
      "balance": 5200000,
      "account_no": "120-456-987654",
      "currency": "KRW",
      "created_at": "2024-02-01T10:00:00Z",
      "updated_at": "2024-07-10T09:00:00Z",
      "transactions": [
        {
          "id": 15,
          "account_id": 2,
          "type": "OUT",
          "amount": 4622437,
          "balance_after": 2258940,
          "transaction_at": "2025-07-02T21:58:43.780561Z",
          "description": "공과금 납부",
          "counter_account_no": "COMP-333-999111",
          "created_at": "2025-07-02T21:58:43.780561Z",
          "updated_at": "2025-07-02T21:58:43.780561Z"
        },
        {
          "id": 27,
          "account_id": 2,
          "type": "OUT",
          "amount": 660270,
          "balance_after": 9416224,
          "transaction_at": "2025-07-03T12:11:43.780675Z",
          "description": "카페 결제",
          "counter_account_no": "COMP-333-999111",
          "created_at": "2025-07-03T12:11:43.780675Z",
          "updated_at": "2025-07-03T12:11:43.780675Z"
        },
        {
          "id": 45,
          "account_id": 2,
          "type": "OUT",
          "amount": 1841732,
          "balance_after": 3876265,
          "transaction_at": "2025-07-08T14:38:43.780823Z",
          "description": "렌트비 이체",
          "counter_account_no": "111-222-333333",
          "created_at": "2025-07-08T14:38:43.780823Z",
          "updated_at": "2025-07-08T14:38:43.780823Z"
        },
        {
          "id": 83,
          "account_id": 2,
          "type": "IN",
          "amount": 2355523,
          "balance_after": 2355523,
          "transaction_at": "2025-06-24T17:59:43.781194Z",
          "description": "적립이자",
          "counter_account_no": "987-111-333444",
          "created_at": "2025-06-24T17:59:43.781194Z",
          "updated_at": "2025-06-24T17:59:43.781194Z"
        },
        {
          "id": 94,
          "account_id": 2,
          "type": "IN",
          "amount": 4587597,
          "balance_after": 4587597,
          "transaction_at": "2025-07-03T06:25:43.781273Z",
          "description": "급여 입금",
          "counter_account_no": "BANK-123-456789",
          "created_at": "2025-07-03T06:25:43.781273Z",
          "updated_at": "2025-07-03T06:25:43.781273Z"
        }
      ]
    },
    {
      "id": 3,
      "customer_id": 1,
      "name": "CMA통장",
      "type": "CMA",
      "balance": 3500000,
      "account_no": "130-000-111111",
      "currency": "KRW",
      "created_at": "2024-03-01T10:00:00Z",
      "updated_at": "2024-07-10T09:00:00Z",
      "transactions": [
        {
          "id": 19,
          "account_id": 3,
          "type": "OUT",
          "amount": 2361577,
          "balance_after": 3363723,
          "transaction_at": "2025-06-25T03:14:43.780619Z",
          "description": "공과금 납부",
          "counter_account_no": "COMP-333-999111",
          "created_at": "2025-06-25T03:14:43.780619Z",
          "updated_at": "2025-06-25T03:14:43.780619Z"
        },
        {
          "id": 37,
          "account_id": 3,
          "type": "IN",
          "amount": 319709,
          "balance_after": 319709,
          "transaction_at": "2025-07-13T21:51:43.780744Z",
          "description": "저축 자동이체",
          "counter_account_no": "111-222-333333",
          "created_at": "2025-07-13T21:51:43.780744Z",
          "updated_at": "2025-07-13T21:51:43.780744Z"
        },
        {
          "id": 56,
          "account_id": 3,
          "type": "OUT",
          "amount": 1520420,
          "balance_after": 7919111,
          "transaction_at": "2025-06-23T06:10:43.780946Z",
          "description": "공과금 납부",
          "counter_account_no": "BANK-123-456789",
          "created_at": "2025-06-23T06:10:43.780946Z",
          "updated_at": "2025-06-23T06:10:43.780946Z"
        },
        {
          "id": 71,
          "account_id": 3,
          "type": "IN",
          "amount": 268000,
          "balance_after": 268000,
          "transaction_at": "2025-06-30T17:31:43.781081Z",
          "description": "이체입금",
          "counter_account_no": "COMP-333-999111",
          "created_at": "2025-06-30T17:31:43.781081Z",
          "updated_at": "2025-06-30T17:31:43.781081Z"
        },
        {
          "id": 73,
          "account_id": 3,
          "type": "IN",
          "amount": 1876597,
          "balance_after": 1876597,
          "transaction_at": "2025-07-16T21:37:43.781106Z",
          "description": "이체입금",
          "counter_account_no": "222-333-444555",
          "created_at": "2025-07-16T21:37:43.781106Z",
          "updated_at": "2025-07-16T21:37:43.781106Z"
        }
      ]
    },
    {
      "id": 4,
      "customer_id": 1,
      "name": "비상금통장",
      "type": "SAVINGS",
      "balance": 500000,
      "account_no": "140-222-333333",
      "currency": "KRW",
      "created_at": "2024-04-01T10:00:00Z",
      "updated_at": "2024-07-10T09:00:00Z",
      "transactions": [
        {
          "id": 11,
          "account_id": 4,
          "type": "OUT",
          "amount": 1828762,
          "balance_after": 2698802,
          "transaction_at": "2025-07-09T15:46:43.780532Z",
          "description": "카페 결제",
          "counter_account_no": "987-654-321000",
          "created_at": "2025-07-09T15:46:43.780532Z",
          "updated_at": "2025-07-09T15:46:43.780532Z"
        },
        {
          "id": 25,
          "account_id": 4,
          "type": "IN",
          "amount": 371341,
          "balance_after": 371341,
          "transaction_at": "2025-07-03T06:52:43.780661Z",
          "description": "저축 자동이체",
          "counter_account_no": "222-333-444555",
          "created_at": "2025-07-03T06:52:43.780661Z",
          "updated_at": "2025-07-03T06:52:43.780661Z"
        },
        {
          "id": 46,
          "account_id": 4,
          "type": "IN",
          "amount": 3465129,
          "balance_after": 3465129,
          "transaction_at": "2025-07-05T05:55:43.780829Z",
          "description": "적립이자",
          "counter_account_no": "COMP-333-999111",
          "created_at": "2025-07-05T05:55:43.780829Z",
          "updated_at": "2025-07-05T05:55:43.780829Z"
        },
        {
          "id": 50,
          "account_id": 4,
          "type": "OUT",
          "amount": 1165064,
          "balance_after": 3170582,
          "transaction_at": "2025-07-02T06:48:43.780885Z",
          "description": "자동차 할부",
          "counter_account_no": "",
          "created_at": "2025-07-02T06:48:43.780885Z",
          "updated_at": "2025-07-02T06:48:43.780885Z"
        },
        {
          "id": 72,
          "account_id": 4,
          "type": "OUT",
          "amount": 2400648,
          "balance_after": 1344203,
          "transaction_at": "2025-06-25T06:17:43.781095Z",
          "description": "공과금 납부",
          "counter_account_no": "222-333-444555",
          "created_at": "2025-06-25T06:17:43.781095Z",
          "updated_at": "2025-06-25T06:17:43.781095Z"
        },
        {
          "id": 80,
          "account_id": 4,
          "type": "OUT",
          "amount": 2119558,
          "balance_after": 7503453,
          "transaction_at": "2025-07-11T00:56:43.781174Z",
          "description": "공과금 납부",
          "counter_account_no": "COMP-333-999111",
          "created_at": "2025-07-11T00:56:43.781174Z",
          "updated_at": "2025-07-11T00:56:43.781174Z"
        },
        {
          "id": 96,
          "account_id": 4,
          "type": "OUT",
          "amount": 2577050,
          "balance_after": 2565831,
          "transaction_at": "2025-07-04T10:08:43.781287Z",
          "description": "카페 결제",
          "counter_account_no": "987-654-321000",
          "created_at": "2025-07-04T10:08:43.781287Z",
          "updated_at": "2025-07-04T10:08:43.781287Z"
        }
      ]
    }
  ]
}
```

## 📌 각 서비스 API

### Customer Service API

#### GET /customer/{id}

특정 고객의 정보를 조회합니다.

**요청 예시**:

```bash
GET http://localhost:3001/customer/1
```

**응답 예시**:

```json
{
  "id": 1,
  "name": "김철수",
  "birthDate": "1999-05-16",
  "address": "서울특별시 강남구",
  "created_at": "2024-01-01T10:00:00Z",
  "updated_at": "2024-06-01T12:00:00Z"
}
```

### Account Service API

#### GET /account/{id}

특정 계좌의 정보를 조회합니다.

**요청 예시**:

```bash
GET http://localhost:3002/account/1
```

**응답 예시**:

```json
{
  "id": 1,
  "customer_id": 1,
  "name": "급여통장",
  "type": "CHECKING",
  "balance": 1800000,
  "account_no": "110-123-456789",
  "currency": "KRW",
  "created_at": "2024-01-02T10:00:00Z",
  "updated_at": "2024-07-10T09:00:00Z"
}
```

#### GET /accounts/customer/{customer_id}

특정 고객의 모든 계좌를 조회합니다.

**요청 예시**:

```bash
GET http://localhost:3002/accounts/customer/1
```

**응답 예시**:

```json
[
  {
    "id": 1,
    "customer_id": 1,
    "name": "급여통장",
    "type": "CHECKING",
    "balance": 1800000,
    "account_no": "110-123-456789",
    "currency": "KRW",
    "created_at": "2024-01-02T10:00:00Z",
    "updated_at": "2024-07-10T09:00:00Z"
  },
  {
    "id": 2,
    "customer_id": 1,
    "name": "저축통장",
    "type": "SAVINGS",
    "balance": 5200000,
    "account_no": "120-456-987654",
    "currency": "KRW",
    "created_at": "2024-02-01T10:00:00Z",
    "updated_at": "2024-07-10T09:00:00Z"
  }
]
```

### Transaction Service API

#### GET /transactions/account/{account_id}

특정 계좌의 거래 내역을 조회합니다.

**요청 예시**:

```bash
GET http://localhost:3003/transactions/account/1
```

**응답 예시**:

```json
[
  {
    "id": 4,
    "account_id": 1,
    "type": "IN",
    "amount": 3435604,
    "balance_after": 3435604,
    "transaction_at": "2025-06-21T02:52:43.780456Z",
    "description": "회사 월급",
    "counter_account_no": "987-654-321000",
    "created_at": "2025-06-21T02:52:43.780456Z",
    "updated_at": "2025-06-21T02:52:43.780456Z"
  },
  {
    "id": 17,
    "account_id": 1,
    "type": "OUT",
    "amount": 3008105,
    "balance_after": 8116788,
    "transaction_at": "2025-06-26T05:35:43.780583Z",
    "description": "자동차 할부",
    "counter_account_no": "",
    "created_at": "2025-06-26T05:35:43.780583Z",
    "updated_at": "2025-06-26T05:35:43.780583Z"
  }
]
```

## �🛠️ 설치 및 실행

### 필수 요구사항

- Node.js (v14 이상)
- Python (v3.8 이상)
- Ruby (v3.0 이상)
- Go (v1.19 이상)

### 1. Customer Service 실행

```bash
cd bs-customer
pip install django
python manage.py runserver 3001
```

### 2. Account Service 실행

```bash
cd bs-account
bundle install
rails server -p 3002
```

### 3. Transaction Service 실행

```bash
cd bs-transaction
go mod tidy
go run main.go
```

### 4. 🔥 BFF Service 실행

```bash
cd bs-bff
npm install
npm start
```

## 🧪 테스트

BFF 서비스 테스트:

```bash
curl http://localhost:4000/bank/1
```

## 📊 데이터 구조

각 서비스는 JSON 파일을 데이터베이스로 사용합니다:

- `bs-customer/db.json` - 고객 정보
- `bs-account/db.json` - 계좌 정보
- `bs-transaction/db.json` - 거래 내역

## 📝 라이선스

MIT License

---

**Note**: 이 프로젝트는 학습 목적의 샌드박스 프로젝트입니다.
