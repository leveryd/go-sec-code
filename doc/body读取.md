* flask读取body
    
    ``` 
    from flask import Flask
    from flask import request
    
    app = Flask(__name__)
    
    @app.route("/", methods=['POST'])
    def hello():
        print(request.form)
        return "<p>hello world!</p>"
    ```

* gin读取body
    ```
    func ReadBody(c *gin.Context) {
      //c.Request.Body
    }
    ```
  

* 区别和联系

    flask、net/http 在读body时，都有可能因为没读到Content-Length大小的字节而阻塞。

    c.Request.Body 实现了io.Reader接口，read方法返回`(0 <= n <= len(p))`个字节

    ```
    type Reader interface {
        Read(p []byte) (n int, err error)
    }
    ```
