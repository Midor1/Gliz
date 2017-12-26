# Gliz, the RESTful backend for SmartChain

# API document for v0.3

# Root [/]

## Individual Users [/individual]

### Create an individual user [POST]

注册一个个人用户，请保证用户名和密码均不为空。返回的status为-1表示用户名有冲突或用户名密码有空缺。正常情况返回0和生成的用户id，请注意这个操作不会自动执行登录。

+ Request (multipart/form-data)
    + Key-Value Pairs
        + UserName (string, required)
        + Password (string, required)
        + Email   （string)
        + Tel      (int)

+ Response 200 (application/json)
    + Attributes
        + status (int)
        + id     (int)
    + Body
            
             {
                status:-1,
                id:0               
             }
    

### Get active user [GET]

取得当前登录者的uid。若当前session尚未登录则status返回-1。

+ Response 200 (application/json)
    + Attributes
        + status (int)
        + id     (int)
    + Body
        
            {
                status:0,
                id:11
            }

    

## Authenticator [/auth]

### Password Authentication [POST]

密码登录，成功返回0，失败返回-1。

+ Request (multipart/form-data)
     + Key-Value Pairs
         + UserName (string, required)
         + Password (string, required)

+ Response 200 (application/json)
    + Attribute
        + status (int)
    + Body
     
          {
              status:-1
          }
    

### User Logout [DELETE]

登出。返回登出者的UID，若未登录则返回-1。

+ Response 200 (application/json)
   + Attribute
       + status (int)
   + Body
   
          {
              status:123
          }

## Items [/item]

### Create Item [POST]

上架物品。这个操作需要当前用户拥有卖家权限。

+ Request (multipart/form-data)
     + Key-Value Pairs
         + ItemName    (string, required)
         + Category    (string, required)
         + Price       (int, required) 
         + Description (string, required)
         + Image       (string, required)

+ Response 200 (application/json)
     + Attribute
         + ItemID (int)
     + Body
     
            {
                ItemID: 5
            }

### Item Query [GET]

物品搜索。

+ Parameters
    + name     (string)
    + price_lb (number) 物品价格下界
    + price_ub (number) 物品价格上界
    + category (string)

+ Response 200 (application/json)
     + Attribute
         + ItemsCount (int)
         + Items ([]Item)
     + Body
     
            {
                "ItemsCount": 3,
                "Items": [
                    {
                        "ItemID": 1,
                        "ItemName": "Apple Pie",
                        "Description": "Umai!",
                        "Price": 2300,
                        "Category": "/Food/Bakery",
                        "Image": ""
                    },
                    {
                        "ItemID": 2,
                        "ItemName": "Apple Juice",
                        "Description": "Umai!",
                        "Price": 800,
                        "Category": "/Drink/Juice",
                        "Image": ""
                    },
                    {
                        "ItemID": 3,
                        "ItemName": "Fresh Apple",
                        "Description": "Umai!",
                        "Price": 600,
                        "Category": "/Food/Fruit",
                        "Image": ""
                    }
                ]
            }

## A Certain Item [/item/{ItemID}]

### Get Item Info [GET]

+ Response 200 (application/json)
     + Attribute
         + ItemName    (string)
         + Category    (string)
         + Price       (int) 
         + Description (string)
         + Image       (string)
     + Body
     
            {
                "ItemName": "Fresh Apple",
                "Description": "Umai!",
                "Price": 600,
                "Category": "/Food/Fruit",
                "Image": ""
            }
            
## Cart [/cart]

### Add to Cart [POST]

+ Request (multipart/form-data)
     + Key-Value Pairs
         + ItemID    (int, required)
         + Amount    (int, default = 1)

+ Response 200 (application/json)
    + Attribute
        + status (int)
    + Body
     
          {
              status:-1
          }

### Show Cart [GET]

+ Response 200 (application/json)
     + Attribute
         + ItemsCount  (int)  
         + CartItems   (CartItem[])
     + Body
     
            {
                "ItemsCount": 3,
                "CartItems": [
                    {
                        "ItemID": 1,
                        "Amount": 3
                    },
                    {
                        "ItemID": 3,
                        "Amount": 1
                    }
                ]
            }
            
### Delete from Cart [DELETE]

+ Request (multipart/form-data)
     + Key-Value Pairs
         + ItemID    (int, required)
         + Amount    (int, default = 1)
  
+ Response 200 (application/json)
    + Attribute
        + status (int)
    + Body
     
          {
              status:-1
          }

## Orders [/order]

### Check Out [POST]

+ Response 200 (application/json)
    + Attribute
         + OrderID    (int)
    + Body
     
           {
               OrderID:242545
           }
           
### Order Query [GET]

+ Parameters
    + status   (int)
    + time     (timestamp)

+ Response 200 (application/json)
     + Attribute
         + OrdersCount (int)
         + Orders ([]int)
     + Body
     
            {
                "OrdersCount": 3,
                "Orders": [
                    32465425463544,
                    32565636454631,
                    787654323456
                ]
            }
            
## A Certain Order [/order/{OrderID}]

### Get Order Info [GET]

This API will check the permission of the active user.

+ Response 200 (application/json)
     + Attribute
         + UserID     (int)
         + Subtotal   (int)
         + Time       (timestamp) 
         + OrderItems (OrderItem[])
     + Body
     
            {
                "UserID": 23,
                "Subtotal": 239800,
                "Time": 3456787320,
                "OrderItems": [
                        {
                            "ItemID": 1,
                            "Amount": 3,
                            "Price": 2399 
                        },
                        {
                            "ItemID": 3,
                            "Amount": 1,
                            "Price": 400
                        }
                    ]
            }


