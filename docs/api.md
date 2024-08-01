# API

ユースケースごとにAPIの仕様を記述する.

## 請求書データ作成

### url

```text
POST /api/invoices
```

### request body

```json
{
  "issue_date": "2023-07-01",
  "payment_amount": 1000.00,
  "fee": 50.00,
  "fee_rate": 5.00,
  "consumption_tax": 80.00,
  "consumption_tax_rate": 8.00,
  "payment_due_date": "2023-08-01",
  "status": "未処理"
}
```

### response

No Content

### error

e.g. validation error

```json
{
    "error_code": "validation_error",
}
```

## 請求書データ一覧取得

### url

```text
GET /api/invoices
```

### query params

- `start_date`: string, required
- `end_date`: string, required

### response

```json
{
  "invoices": [
    {
      "invoice_id": 1,
      "client_id": 1,
      "issue_date": "2023-07-01",
      "payment_amount": 1000.00,
      "fee": 50.00,
      "fee_rate": 5.00,
      "consumption_tax": 80.00,
      "consumption_tax_rate": 8.00,
      "payment_due_date": "2023-08-01",
      "billing_amount": 1088.00,
      "status": "未処理"
    }
  ]
}
```

### error

e.g. validation error

```json
{
    "error_code": "validation_error",
}
```
