状态机go语言版本。QQ交流：505637659@qq.com。

状态机支持父状态与子状态，一个父状态可以注册多个子状态，子状态之间可以互相跳转，但不同父状态之间的子状态是隔离的，不能跳转。

simpleTest和fullTest目录中的statemachine是状态机的核心代码，simpleTest目录演示了一个简单例子，该示例只包含一个父状态，fullTest目录演示了复杂一点的例子，包含了两个父状态。

**1.simpleTest**

编译运行方法：

 cd simpleTest && go build .
 ./fsm
 
 该示例的状态机框架如下图所示：

<img width="944" height="497" alt="image" src="https://github.com/user-attachments/assets/dd6f2928-8cb4-49ef-b42d-e757de9caa80" />

外部通过调用StateMachine对象的Start和Stop方法来启停状态机，调用Start方法时，状态机会运行work state中的子状态，子状态都有自己的标签，StateGetJob的标签为"get job",
StateDoJob的标签为"do job",StateFinish的标签为"finish job"，状态机通过判断下一个状态的标签来实现状态之间的转换。调用Stop方法时，不论状态机处于何种状态，
都会立刻进入StateStop状态，该状态的OnEntry函数可以做一些异常处理操作。

simpleTest目录下main.go中有完整的代码注释，下面是补充说明：

**用户的子状态需要覆盖Stater接口中的3个方法：**

OnEntry(para interface{})   //在进入该状态时被调用，即从其他状态跳转到该状态，可以做一些初始化操作，para是从上一个状态传递过来的信息

OnLoop()                    //这是该状态的自循环函数，可以执行多次，用于等待某个事件，如果不需要自循环，该方法可以不覆盖

OnExit()                    //在退出该状态时被调用，可以做一些清理操作

想从一个子状态跳转到另一个子状态，需要调用接口的PostEvent方法，该方法原型如下：

PostEvent(event string, paras ...interface{})

其中event参数是跳转到目标状态的“标签”, para参数是需要给下一个状态传递的自定义信息，默认为空

如果想在某个子状态中自循环，可以在OnEntry函数中调用PostEvent方法，方法中event参数填写自身的标签，状态机发现下个运行的状态还是该状态，则调用该状态的OnLoop方法，
此时在该状态的OnLoop函数中需要继续调用PostEvent函数并填写自身的标签，以实现循环调用OnLoop函数，如果某个条件成立时想退出该状态，则调用PostEvent函数并填写其他状态的标签

**状态机注意要点：**

(1)至少有一个父状态，子状态需要向父状态注册，父状态是子状态的载体

(2)子状态的OnEntry方法必须调用一次PostEvent函数，告诉状态机下一个运行的状态，如果是自己，那么在OnLoop方法中要继续调用PostEvent函数，如果忘记调用PostEvent，状态机不知道下一个运行的状态，会提示错误

(3)不要在子状态中死循环或者阻塞，会占用状态机内部线程，导致外部无法启停状态机或进行父状态切换，可以用OnLoop方法替代阻塞

**2.fullTest**

编译运行方法：

 cd fullTest && go build .

 ./fsm
 
 该示例的状态机框架如下图所示：
 <img width="961" height="446" alt="image" src="https://github.com/user-attachments/assets/ef4242c6-6a8e-481a-80e9-f09f198c2081" />

 该示例展示了两个父状态的运行情况，父状态之间切换是通过调用StateMachine对象的TransferState函数实现的，该函数接收一个参数，即要跳转的父状态的标签，在该示例中work state的标签是"work state",
 charge state的标签是"charge state"。

 状态机启动时，需要确定运行于哪个父状态，这是通过调用StateMachine对象的SetInitState方法实现的，该方法接收一个参数，即父状态的标签。
 
 父状态运行时，也需要确定先运行哪个子状态，通过调用父状态的SetInitState方法，传入该父状态下的某个子状态标签实现。


