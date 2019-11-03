## 基本数据类型
### 3.1. Integers 51
### 3.2. Floating-Point Numbers 56
#### IEEE754
IEEE754(浮点数在内存中实际存储的内容)，用科学计数法表示小数，转为二进制数后记作1.nnnn x 2^m,存储时省去1.，nnnn为尾数，m为指数。。
##### 单精度
- 总长度32bit，即4byte。
- 1bit表示符号0正1负。
- 8bit表示指数(-127～128)，存储时为了不判断符号，加上127。
- 23bit表示尾数(小数)。

### 3.3. Complex Numbers 61
### 3.4. Booleans 63
### 3.5. Strings 64
### 3.6. Constants
