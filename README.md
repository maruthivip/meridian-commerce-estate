# Meridian Commerce Estate

Polyglot microservices estate for the Meridian storefront platform.

## Services
| Service | Language | Role |
|---|---|---|
| storegate | Go | Storefront HTTP gateway; calls most backend services |
| basketledger | C# | Basket state, backed by redis |
| orderforge | Go | Order placement orchestration; writes order_ledger |
| catalogvault | Go | Product catalog |
| fxrates | Node | Currency conversion |
| payguard | Node | Payment authorization |
| freightquote | Go | Shipping quotes |
| mailcourier | Python | Order confirmation email (terminal consumer) |
| suggestr | Python | Product recommendations |
| promobeam | Java | Contextual promotions |
| trafficmimic | Python/Locust | Synthetic traffic |

The storegate web layer additionally depends on the internal `@meridian/*`
SDK packages under `libs/`.

Note: the aidexk assistant service is operated by a partner team outside this
repository; storegate carries its endpoint configuration only.
