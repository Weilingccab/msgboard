# msgboard

Using :  
---
    Postgresql on AWS. 
    Go: gin gorm. 
    Docker build image deploy to AWS EC2
    API document: gin-swagger 
需求：
---  

    1.完成使用者註冊、登入. 
    2.使用者登入後可留言、回覆留言. 
    3.系統管理人員有以下權限. 
      3.1.使用者停權.
      3.2.留言關閉  
      3.3.留言隱藏. 
      3.4.彈性查詢留言. 
連線測試：
---
  http://35.78.97.207:8080/msgboard/test/ping
  連進去如果看到回應pong，即成功
  
  http://35.78.97.207:8080/msgboard/userInfo/users
  會回傳使用者陣列
  
  http://35.78.97.207:8080/swagger/index.html
  可以透過該頁面看到API文件
  
詳細操作文件：
---
  https://hackmd.io/@1R4BnDrZQVa-_xVJKEo6ww/H1cWssDHq

ER Model:
---
  
  ![image](https://user-images.githubusercontent.com/99722169/163770785-2a98defd-091d-45af-90dc-a8fffbd39e24.png)

