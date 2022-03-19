# Data code-level

### Tasks

- [Create DB for](#create-db):
    - [Deposit](#deposit-repository)
    - [Spend](#spend-repository)
    - [Money-storage](#money-storage-repository)
- [Write three repository-package](#repository-package-and-interface)
- Write interface for work with DB
- Implement those interfaces and check work with DB
- [Create mocks and write auto-tests for checking the repository](#tests)

### Create DB

Need three-basic DB to store info about deposits and

#### Deposit repository

```
Fields:

id          int
title       string
amount      float
created_at  time.Time
category    depositCategoryType
```

#### Spend repository

```
Fields:

id          int
title       string
amount      float
created_at  time.Time
category    spendCategoryType 
```

#### Money-storage repository

##### money-storage list
```
Fields

storage_id      int
storage_name    string
```

##### money-transactions
```
Fields

storage_id            int
storage_amount_after  float
transaction_type      transactionType
transaction_id        int 
```

### Repository-package and interface

Package should look like
- Interface with (at least) CRUD methods
- Implementation folder with working methods
- Mocks folder 
- Tests folder

### Tests
- For each method:
  - Mocks-unit tests with random responses from db
  - Integration test on blank DB with CRUD DB fields
