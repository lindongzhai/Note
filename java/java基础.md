# java 基础

## 简介

Java最早是由SUN公司（已被Oracle收购）的[詹姆斯·高斯林](https://en.wikipedia.org/wiki/James_Gosling)（高司令，人称Java之父）在上个世纪90年代初开发的一种编程语言，最初被命名为Oak，目标是针对小型家电设备的嵌入式应用，结果市场没啥反响。谁料到互联网的崛起，让Oak重新焕发了生机，于是SUN公司改造了Oak，在1995年以Java的名称正式发布，原因是Oak已经被人注册了，因此SUN注册了Java这个商标。随着互联网的高速发展，Java逐渐成为最重要的网络编程语言。

Java介于编译型语言和解释型语言之间。编译型语言如C、C++，代码是直接编译成机器码执行，但是不同的平台（x86、ARM等）CPU的指令集不同，因此，需要编译出每一种平台的对应机器码。解释型语言如Python、Ruby没有这个问题，可以由解释器直接加载源码然后运行，代价是运行效率太低。而Java是将代码编译成一种“字节码”，它类似于抽象的CPU指令，然后，针对不同平台编写虚拟机，不同平台的虚拟机负责加载字节码并执行，这样就实现了“一次编写，到处运行”的效果。当然，这是针对Java开发者而言。对于虚拟机，需要为每个平台分别开发。为了保证不同平台、不同公司开发的虚拟机都能正确执行Java字节码，SUN公司制定了一系列的Java虚拟机规范。从实践的角度看，JVM的兼容性做得非常好，低版本的Java字节码完全可以正常运行在高版本的JVM上。

随着Java的发展，SUN给Java又分出了三个不同版本：

- Java SE：Standard Edition
- Java EE：Enterprise Edition
- Java ME：Micro Edition

三者的关系如下

```ascii
┌───────────────────────────┐
│Java EE                    │
│    ┌────────────────────┐ │
│    │Java SE             │ │
│    │    ┌─────────────┐ │ │
│    │    │   Java ME   │ │ │
│    │    └─────────────┘ │ │
│    └────────────────────┘ │
└───────────────────────────┘
```

## 反射

### 1、class 类

`class` 是由 JVM 在执行过程中动态加载的。JVM在第一次读取到一种 `class` 类型时，将其加载进内存。

每加载一种 `class`，JVM 就为其创建一个 `Class` 类型的实例，并关联起来.

```java
public final class Class {
    private Class()
}
```

```java
// 加载 String，并创建实例
Class cls = new Class(String);
```

`Class` 实例是由 JVM 内部创建的，`Class` 类的构造方法是 `private`，只有 JVM 能创建 `Class` 实例。

JVM 持有的每个 `Class` 实例都指向一个数据类型（`class` 或 `interface`）

```
┌───────────────────────────┐
│      Class Instance       │──────> String
├───────────────────────────┤
│name = "java.lang.String"  │
└───────────────────────────┘
┌───────────────────────────┐
│      Class Instance       │──────> Random
├───────────────────────────┤
│name = "java.util.Random"  │
└───────────────────────────┘
┌───────────────────────────┐
│      Class Instance       │──────> Runnable
├───────────────────────────┤
│name = "java.lang.Runnable"│
└───────────────────────────┘
┌───────────────────────────┐
│      Class Instance       │──────> String
├───────────────────────────┤
│name = "java.lang.String"  │
├───────────────────────────┤
│package = "java.lang"      │
├───────────────────────────┤
│super = "java.lang.Object" │
├───────────────────────────┤
│interface = CharSequence...│
├───────────────────────────┤
│field = value[],hash,...   │
├───────────────────────────┤
│method = indexOf()...      │
└───────────────────────────┘
```

通过 `Class` 实例获取 `class` 信息的方法称为反射

获取 `class` 的 `Class` 的三种方法：

1. 直接通过一个 `class` 的静态变量 `class` 获取

   ```java
   Class cls = String.class;
   ```

2. 通过实例变量的 `getClass()` 方法获取

   ```java
   String s = "Hello";
   Class cls = s.getClass();
   ```

3. 通过`class` 的完整类名的静态方法 `class.forName()` 方法获取

   ```java
   Class cls = Class.forName("java.lang.String")
   ```

`Class` 实例在 JVM 中是唯一的。所以可以用 `==` 比较两个 `Class` 实例

```java
Class cls1 = String.class;

String s = "Hello";
Class cls2 = s.getClass();

boolean sameClass = cls == cls2; // true
```

`Class` 实例比较与 `instanceof` 的区别

```java
Integer n = new Integer(123);

boolean b1 = n instanceof Integer; // true， 因为 n 是 Integer 类型
boolean b2 = n instanceof Number; // true， 因为 n 是 Number 类型的子类

boolean b3 = n.getClass() == Integer.class; // true，因为 n.getClass() 返回 Integer.class
boolean b4 = n.getClass() == Number.class; // false，因为 Integer.class != Number.class
```

用`instanceof`不但匹配指定类型，还匹配指定类型的子类。而用`==`判断`class`实例可以精确地判断数据类型，但不能作子类型比较。

反射的目的是为了获得某个实例的信息。当拿到某个`Object`实例时，可以通过反射获取该`Object`的`class`信息

```java
void printObjectInfo(Object obj) {
    Class cls = obj.getClass();
}
```

获取 `Class` 实例基本信息的方法

```java
Class cls = Integer.class;
cls.getName(); // 获取类的全限定名，包括包名。 java.lang.Integer
cls.getSimpleName(); // 获取类的简单名称，即类名部分，不包括包名。 Integer
cls.getPackage(); // 获取表示类所在包的 Package 对象。java.lang
cls.getPackage().getName(); // 获取类所在包的名称。java.lang
cls.isInterface(); // 判断类是否是接口。false
cls.isEnum(); // 判断类是否是枚举类型。false
cls.isArray(); // 判断类是否是数组类型。false
cls.isPrimitive(); // 判断类是否是基本数据类型。false
```

JVM 为每一种基本类型也创建了 `Class` 实例，可以通过 `基本类型.class` 访问

可以通过 `Class` 实例来创建对应类型的实例

```java
// 获取 String 的 Class 实例
Class cls = String.class;
// 创建一个 String 实例
String s = (String) cls.newInstance();
```

`Class.newInstance()` 只能调用 `public` 的无参数构造方法。带参数的构造方法或者非 `public` 的构造方法都无法通过 `Class.newInstance()` 被调用

#### 1.1、动态加载

JVM在执行Java程序的时候，并不是一次性把所有用到的class全部加载到内存，而是第一次需要用到class时才加载。

动态加载`class`的特性对于Java程序非常重要。利用JVM动态加载`class`的特性，才能在运行期根据条件加载不同的实现类。



#### 1.2、小结

JVM为每个加载的`class`及`interface`创建了对应的`Class`实例来保存`class`及`interface`的所有信息；

获取一个`class`对应的`Class`实例后，就可以获取该`class`的所有信息；

通过Class实例获取`class`信息的方法称为反射（Reflection）；

JVM总是动态加载`class`，可以在运行期根据条件来控制加载class。