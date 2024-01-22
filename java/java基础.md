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

* `class` 是由 JVM 在执行过程中动态加载的。JVM在第一次读取到一种 `class` 类型时，将其加载进内存。

* 每加载一种 `class`，JVM 就为其创建一个 `Class` 类型的实例，并关联起来.

    ```java
    public final class Class {
        private Class()
    }
    ```

    ```java
    // 加载 String，并创建实例
    Class cls = new Class(String);
    ```

* `Class` 实例是由 JVM 内部创建的，`Class` 类的构造方法是 `private`，只有 JVM 能创建 `Class` 实例。

* JVM 持有的每个 `Class` 实例都指向一个数据类型（`class` 或 `interface`）

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

* 通过 `Class` 实例获取 `class` 信息的方法称为反射

* 获取 `class` 的 `Class` 的三种方法：

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

* `Class` 实例在 JVM 中是唯一的。所以可以用 `==` 比较两个 `Class` 实例

    ```java
    Class cls1 = String.class;

    String s = "Hello";
    Class cls2 = s.getClass();

    boolean sameClass = cls == cls2; // true
    ```

* `Class` 实例比较与 `instanceof` 的区别

    ```java
    Integer n = new Integer(123);

    boolean b1 = n instanceof Integer; // true， 因为 n 是 Integer 类型
    boolean b2 = n instanceof Number; // true， 因为 n 是 Number 类型的子类

    boolean b3 = n.getClass() == Integer.class; // true，因为 n.getClass() 返回 Integer.class
    boolean b4 = n.getClass() == Number.class; // false，因为 Integer.class != Number.class
    ```

* 用`instanceof`不但匹配指定类型，还匹配指定类型的子类。而用`==`判断`class`实例可以精确地判断数据类型，但不能作子类型比较。

* 反射的目的是为了获得某个实例的信息。当拿到某个`Object`实例时，可以通过反射获取该`Object`的`class`信息

    ```java
    void printObjectInfo(Object obj) {
        Class cls = obj.getClass();
    }
    ```

* 获取 `Class` 实例基本信息的方法

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

* JVM 为每一种基本类型也创建了 `Class` 实例，可以通过 `基本类型.class` 访问

* 可以通过 `Class` 实例来创建对应类型的实例

    ```java
    // 获取 String 的 Class 实例
    Class cls = String.class;
    // 创建一个 String 实例
    String s = (String) cls.newInstance();
    ```

* `Class.newInstance()` 只能调用 `public` 的无参数构造方法。带参数的构造方法或者非 `public` 的构造方法都无法通过 `Class.newInstance()` 被调用

#### 1.1、动态加载

* JVM在执行Java程序的时候，并不是一次性把所有用到的class全部加载到内存，而是第一次需要用到class时才加载。

* 动态加载`class`的特性对于Java程序非常重要。利用JVM动态加载`class`的特性，才能在运行期根据条件加载不同的实现类。

#### 1.2、小结

* JVM为每个加载的`class`及`interface`创建了对应的`Class`实例来保存`class`及`interface`的所有信息；
* 获取一个`class`对应的`Class`实例后，就可以获取该`class`的所有信息；
* 通过Class实例获取`class`信息的方法称为反射（Reflection）；
* JVM总是动态加载`class`，可以在运行期根据条件来控制加载class。



### 2、访问字段

* 获取 `Field` 信息方法
  * `Field getField(name)`：根据字段名获取某个 `public` 的 `field`（包括父类）
  * `Field getDeclaredField(name)`：根据字段名获取当前类的某个 `field`（不包括父类）
  * `Field[] getFields()`：获取所有 `public` 的 `field`（包括父类）
  * `Field[] getDeclaredFields()`：获取当前类的所有 `field`（不包括父类）

    ```java
    public class Main {
        public static void main(String[] args) throws Exception {
            Class stdClass = Student.class;
            // 获取 public 字段 "score"
            System.out.println(stdClass.getField("score")); // public int Student.score
            // 获取继承的 public 字段 name
            System.out.println(stdClass.getField("name")); // public java.lang.String Person.name
            // 获取 private 字段 "grade"
            System.out.pringln(stdClass.getDeclaredField("grade")); // private int Student.grade
        }
    }
    
    class Student extends Person {
        public int score;
        private int grade;
    }
    
    class Person {
        public String name;
    }
    ```

*  `Field` 对象包含一个字段的所有信息

  * `getName()`: 返回字段名称

  * `getType()`: 返回字段类型

  * `getModifiers()`: 返回字段的修饰符，返回一个 `int`，不同的 `bit` 表示不同的含义。

    ```java
    public final class String {
        private final byte[] value;
    }
    
    // 通过反射获取字段信息
    Field f = String.class.getDeclaredField("value");
    f.getName(); // 获取名称。value
    f.getType(); // 获取类型。class [B
    int m = f.getModifiers(); // 获取字段修饰符
    Modifier.isFinal(m); // 是否为 final。true
    Modifier.isPublic(m); // 是否为 public。false
    Modifier.isProtected(m); // 是否为 protected。false
    Modifier.isPrivate(m); // 是否为 private。true
    Modifier.isStatic(m); // 是否为 static。false
    ```

#### 2.1、访问字段值

  * 通过 `Field.get(Object)` 获取指定实例的指定字段的指

      ```java
      public class Main() {
          public static void main(String[] args) throws Exception {
              Object p = new Person("xiao ming");
              Class cls = p.getClass();
              Field f = c.getDeclaredField("name");
              Object value = f.get("p");
              System.out.println(value); // xiao ming
          }
      }
      
      class Person {
      	private String name;
          public Person(String name) {
              this.name = name;
          }
      }
      ```
  
  * 正常情况下无法访问 `Class` 类的 `private` 字段。可以通过 `Field.setAccessible(true)` 进行忽略字段修饰符。但可能会失败；
  
  * JVM运行期存在`SecurityManager`，它会根据规则进行检查，有可能阻止`setAccessible(true)`；
  
    ```java
    import java.lang.reflect.Field;
    
    public class Main() {
        public static void main(String[] args) throws Exception {
            Object p = new Person("xiao ming");
            Class cls = p.getClass();
            Field f = c.getDeclaredField("name");
            f.setAccessible(true);
            Object value = f.get("p");
            System.out.println(value); // xiao ming
        }
    }
    
    class Person {
    	private String name;
        public Person(String name) {
            this.name = name;
        }
    }
    ```
  
  * 反射是一种非常规的用法，使用反射，首先代码非常繁琐，其次，它更多地是给工具或者底层框架来使用，目的是在不知道目标实例任何信息的情况下，获取特定字段的值。

#### 2.2、设置字段值

* 通过 `Field.set(Object. Object)` 可以修改字段的值，第一个 `Object` 是指定实例，第二个 `Object` 是待修改的值

  ```java
  import java.lang.reflect.Field;
  
  public class Main {
      public static void main(String[] args) throws Exception {
          Person p = new Person("xiao ming");
          System.out.println(p.getName()); // xiao ming
          Class cls = p.getClass();
          Field f = c.getDeclaredField("name");
          f.setAccessible(true);
          f.set(p, "xiao hong");
          System.out.pringln(p.getName()); // xiao hong
      }
  }
  
  class Person {
      private String name;
      
      public Person(String name) {
          this.name = name;
      }
      
      public String getName() {
          return this.name;
      }
  }
  ```

#### 2.3、小结

* Java的反射API提供的`Field`类封装了字段的所有信息：

* 通过`Class`实例的方法可以获取`Field`实例：`getField()`，`getFields()`，`getDeclaredField()`，`getDeclaredFields()`；

* 通过Field实例可以获取字段信息：`getName()`，`getType()`，`getModifiers()`；

* 通过Field实例可以读取或设置某个对象的字段，如果存在访问限制，要首先调用`setAccessible(true)`来访问非`public`字段。

* 通过反射读写字段是一种非常规方法，它会破坏对象的封装。



### 3、调用方法

* 获取所有 `Method` 信息方法

  * `Method getMethod(name, Class...)`：获取某个`public`的`Method`（包括父类）

  * `Method getDeclaredMethod(name, Class...)`：获取当前类的某个`Method`（不包括父类）

  * `Method[] getMethods()`：获取所有`public`的`Method`（包括父类）

  * `Method[] getDeclaredMethods()`：获取当前类的所有`Method`（不包括父类）

    ```java
    public class Main {
        public static void main(String[] args) throws Exception {
            Class stdClass = Student.class;
            // 获取 public 方法 getScore，参数为 String
            System.out.println(stdClass.getMethod("getScore"), String.class); // public int Student.getScore(java.lang.String)
            // 获取继承的 public 方法 getName，无参数
            System.out.println(stdClass.getMethod("getName")); // public java.lang.String Person.getName()
            // 获取 private 方法 getGrade，参数为 int
            System.out.println(stdClass.getDeclaredMethod("getGrade", int.class)) // private int Student.getGrade(int)
        }
    }
    
    class Student extends Person {
        public int getScore(String type) {
            return 99;
        }
        
        public int getGrade(int year) {
            return 1;
        }
    }
    
    class Person {
        public String getName() {
            return "Person";
        }
    }
    ```

* `Method` 对象包含一个方法的所有信息：
  * `getName()`: 返回方法名称
  * `getReturnType()`: 返回方法返回值类型
  * `getParameterTypes()`: 返回方法的参数类型
  * `getModifiers()`: 返回方法的修饰符

#### 3.1、调用方法

* 对`Method`实例调用`invoke`就相当于调用该方法，`invoke`的第一个参数是对象实例，即在哪个实例上调用该方法，后面的可变参数要与方法参数一致，否则将报错。

    ```java
    import java.lang.reflect.Method;
    
    public class Main {
        public static void main(String[] args) throws Exception {
            // String 对象
            String s = "Hello World";
            // 获取 String substring(int) 方法，参数为 int
            Method m = String.class.getMethod("subString", int.class);
            // 在 s 对象上调用该方法并获取结果
            String r = (String) m.invoke(s, 6);
            // 打印调用结果
            System.out.println(r);
        }
    }
    ```

#### 3.2、调用静态方法

* 对`Method`实例调用`invoke`就相当于调用该方法，`invoke` 的第一个参数为 `null` 则为调用静态方法。

  ```java
  public class Main {
      public static void main(String[] args) throws Exception {
          // 获取Integer.parseInt(String)方法，参数为String:
          Method m = Integer.class.getMethod("parseInt", String.class);
          // 调用该静态方法并获取结果:
          Integer n = (Integer) m.invoke(null, "12345");
          // 打印调用结果:
          System.out.println(n);
      }
  }
  ```

#### 3.3、调用非 `public` 方法

* 正常情况下无法访问 `Class` 类的 `private` 方法。可以通过 `Field.setAccessible(true)` 进行忽略字段修饰符。但可能会失败；

* JVM运行期存在`SecurityManager`，它会根据规则进行检查，有可能阻止`setAccessible(true)`；

  ```java
  public class Main {
      public static void main(String[] args) throws Exception {
          Person p = new Person();
          Method m = p.getClass().getDeclaredMethod("setName", String.class);
          m.setAccessible(true);
          m.invoke(p, "Bob");
          System.out.println(p.name);
      }
  }
  
  class Person {
      String name;
      private void setName(String name) {
          this.name = name;
      }
  }
  ```


#### 3.4、多态

* 使用反射调用方法时，仍然遵循多态原则

  ```java
  public class Main {
      public static void main(String[] args) throws Exception {
          // 获取Person的hello方法:
          Method h = Person.class.getMethod("hello");
          // 对Student实例调用hello方法:
          h.invoke(new Student());
      }
  }
  
  class Person {
      public void hello() {
          System.out.println("Person:hello");
      }
  }
  
  class Student extends Person {
      public void hello() {
          System.out.println("Student:hello");
      }
  }
  ```

#### 3.5、总结

* Java的反射API提供的 `Method` 对象封装了方法的所有信息

* 通过`Class`实例的方法可以获取`Method`实例：`getMethod()`，`getMethods()`，`getDeclaredMethod()`，`getDeclaredMethods()`；

* 通过`Method`实例可以获取方法信息：`getName()`，`getReturnType()`，`getParameterTypes()`，`getModifiers()`；

* 通过`Method`实例可以调用某个对象的方法：`Object invoke(Object instance, Object... parameters)`；

* 通过设置`setAccessible(true)`来访问非`public`方法；

* 通过反射调用方法时，仍然遵循多态原则。
