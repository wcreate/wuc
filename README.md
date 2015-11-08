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

    {
      "id_name": "captcha_id",
      "id_value": "RYZTF7NRKl35F4x",
      "img_url": "/captcha/img/RYZTF7NRKl35F4x.png"
    }
    ```

  - Add a user

      **Example request**:
      ```
      POST http://127.0.0.1:8080/api/user/signup HTTP/1.1
      Content-Type: application/json; charset=UTF-8

      {
        "email": "1@test.com",
        "username": "test1",
        "password": "123456",
        "captcha_id": "xD586zFSptv3Qq2",
        "captcha_value": "689826"
      }
      ```
      **Example response**:
      ```
      HTTP/1.1 200 OK
      Content-Type: application/json; charset=UTF-8

      {
        "uid": 1,
        "username": "test1",
        "token": "2JzpfhH61KjbzXtaUniEI4wI5FS8WGfWxfNA9LIqybcyHDC-w58xln2BjehGhN-r3uRvXthfsqJNPJeUnqsfIP1uKUUNa8e-t4rAVUF7I-JiipLTLGi7LBF1PeIs4c6g"
      }
      ```

  - Delete a user

  - Get a user's info

  - Modify a user's Info

  - Modify a user's password

  - Modify a user's email then send a confirm email

  - User Login

  - User Logout

  - Check a user existed by username/email/mobile

  - Confirm a user after register or modify email
