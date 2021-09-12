# test_nginx
测试nginx



## 准备C/S

- client，发送echo
- server，返回echo。分三个，目的是测试nginx分发协议
  - /api/s1/echo
  - /api/s2/echo
  - /api/s3/echo



## nginx规则

如下，得巧妙的使用rewrite

1. 譬如，server没有定义/s5的处理方式，得让nginx，将/s4的内容，转到s1下：rewrite /api/s4/(.*) /api/s1/$1 break;

```
server {
        listen 20210;
        server_name 0.0.0.0;


        location /api/s1 {
             proxy_pass  http://127.0.0.1:20211/api/s1;
        }


        location /api/s2 {
             proxy_pass  http://127.0.0.1:20212/api/s2;
        }

        location /api/s3 {
             proxy_pass  http://127.0.0.1:20213/api/s3;
        }
        
        location /api/ {
             # 将s5/s6/s7/echo?xxxx替换成s1/echo?xxxx。注意使用了$3，很简单因为定义了3个参数，所以由s3
             rewrite /api/s5/(.*)/(.*)/(.*) /api/s1/$3 break;
             # 将s4替换成s1
             rewrite /api/s4/(.*) /api/s1/$1 break;
             # 将其他的协议，全都转到/api/s1下。
             rewrite /api/(.*) /api/s1/$1 break;
             proxy_pass http://127.0.0.1:20211;
        }
    }
```

