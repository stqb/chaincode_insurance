# chaincode_insurance
##保险保单的查询和编辑功能

#简介

该智能合约实现一个简单的商业应用案例，即通过客户端查询当前系统中保单的详细信息，并可编辑该保单的内容，包括增加保单，修改保单内容，删除保单。在这之中一共分为二种角色：
    普通用户（即投保人）
    保险从业人员（可增删改）

主要实现如下功能:
  1.通过投保人身份证号码 可以查询该用户名下有多少张保单
  2.通过查询得到的保单可以对其进行编辑，包括：增加信息，删除保单等等。

##主要函数
1. init   : 初始化保单信息，并增加一定的保单数量
1. invoke : 调用合约内部的函数
1. query  : 查询相关信息
1. insert : 增加保单
1. delete :  删除保单
1. change : 变更保单内容

##数据结构设计
###Policy 
	PolicyNo     :  保单号码
	PolicyType   :  险种
	Startdate    :  保险生效时间
	Enddate      :  保险失效时间
	Status       :  保单状态
	PolicyHolder :  投保人
	Assured      :  被保险人
	Beneficiary  :  保险受益人
	Premium      :  保费
	Amount       :  保险金
##接口设计
###init 
*request 参数:

  args[0] 

  args[1]

*Response 参数

  返回：
  
###Read
*Request参数

*Response参数


###Write 
*Request参数

*Response参数


##其它 
所有保单数据信息全部保存在区块链上。





先实现一个把保单信息保存到链的功能。
1，保存的信息包括：
保单号码，投保人，被保险人，受益人（身份证号码、名字），保单状态，连续缴费时间，保单价值等。
2，页面：
UI实现【增、删、改、查】功能。
可以根据【保单号码、身份证号码】等信息查询保存在链上的保单信息。
页面能显示被保存的保单信息，以及区块（block）信息。


