# Geek Cloud Native Camp Homework
###

### M1-p1

    给定一个字符串数组
    [“I”,“am”,“stupid”,“and”,“weak”]
    用 for 循环遍历该数组并修改为
    [“I”,“am”,“smart”,“and”,“strong”]

### M1-p2

    基于 Channel 编写一个简单的单线程生产者消费者模型：

    队列：
    队列长度 10，队列元素类型为 int
    生产者：
    每 1 秒往队列中放入一个类型为 int 的元素，队列满时生产者可以阻塞 
    消费者：
    每一秒从队列中获取一个元素并打印，队列为空时消费者阻塞

### M2-p1

    将M1-p2改为多生产者和多消费者

### M2-p2

    1 接收客户端 request，并将 request 中带的 header 写入 response header
    2 读取当前系统的环境变量中的 VERSION 配置，并写入 response header
    3 Server 端记录访问日志包括客户端 IP，HTTP 返回码，输出到 server 端的标准输出
    4 当访问 localhost/healthz 时，应返回 200

