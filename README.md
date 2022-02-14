# jvminhand
自己动手写jvm系列。
用go语言写Java虚拟机，感觉很有趣。

## ch02 加载类
整体而言就是在学习go的文件和字符串操作，感觉还挺好玩的  
犯了一个低级错误，但最后结果还不错，正常获取到了该类文件

## ch03 解析类
解析类文件。
  
## ch04 运行时数据区
  
## ch05 指令集和解释器
```java
public class GaussTest {
    public static void main(String[] args) {
        int sum = 0;
        for (int i = 1; i <= 100; i++) {
            sum += i;
        }
        System.out.println(sum);
    }
}

```
  
出现5050计算结果，正解！
![image](https://user-images.githubusercontent.com/63081109/153890107-bd61b372-3900-4d06-9872-506a3d149ecc.png)
