## UserCentre
 wuc is a user centre service. The minimum requirement of Go is **1.5**.

 wuc provides restful apis to expose services.

 To install wuc:

 	go get github.com/wcreate/wuc

### Rest API
  - Generate a captcha

    **Example request**:
    ```
    GET http://127.0.0.1:8080/captcha/new HTTP/1.1
    ```
    **Example response**:
    ```
    HTTP/1.1 200 OK
    Content-Type: application/json; charset=UTF-8
    ```

  - Add a user

  - Delete a user

  - Get a user's info

  - Modify a user's Info

  - Modify a user's password

  - Modify a user's email then send a confirm email

  - User Login

  - User Logout

  - Check a user existed by username/email/mobile
  
  - Confirm a user after register or modify email
