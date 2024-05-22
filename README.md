# Примеры запросов <a name="examples"></a>


* [Создание заказа:POST http://localhost:8000/order c телом:]
```
{
  "order_uid": "b563feb7b2b84b6test4",
  "track_number": "WBILMTESTTRACK",
  "entry": "WBIL",
  "delivery_name": "Test Testov",
  "delivery_phone": "+9720000000",
  "delivery_zip": "2639809",
  "delivery_city": "Kiryat Mozkin",
  "delivery_address": "Ploshad Mira 15",
  "delivery_region": "Kraiot",
  "delivery_email": "test@gmail.com",
  "payment_transaction": "b563feb7b2b84b6test",
  "request_id": "",
  "currency": "USD",
  "provider": "wbpay",
  "amount": 1817,
  "payment_dt": 1637907727,
  "bank": "alpha",
  "delivery_cost": 1500,
  "goods_total": 317,
  "custom_fee": 1,
  "items": [
    {
      "chrt_id": 9934930,
      "track_number": "WBILMTESTTRACK",
      "price": 453,
      "rid": "ab4219087a764ae0btest",
      "name": "Mascaras",
      "sale": 30,
      "size": "0",
      "total_price": 317,
      "nm_id": 2389212,
      "brand": "Vivienne Sabo",
      "status": 202
    }
  ],
  "locale": "en",
  "internal_signature": "",
  "customer_id": "test",
  "delivery_service": "meest",
  "shardkey": "9",
  "sm_id": 99,
  "date_created": "2021-11-26T06:22:19Z",
  "oof_shard": "1"
}
```


* [Получение заказа по ID: GET http://localhost:8000/order/b563feb7b2b84b6test4]
