# test_nginx
测试nginx



## 准备C/S

- client，发送echo
- server，返回echo。分三个，目的是测试nginx分发协议
  - /api/s1/echo
  - /api/s2/echo
  - /api/s3/echo