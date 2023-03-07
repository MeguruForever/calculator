package src

import (
	"errors"
	"fmt"
	"strconv"
)

//这里是规定一个栈的结构体
type Stack struct{
	MaxSize int
	Top int
	arr [] int
}
//这里是在栈中添加元素的方法
func (stack *Stack) Push(val int) (err error){
	if stack.Top == stack.MaxSize-1{
		return errors.New("stack full")
	}
	stack.Top++
	stack.arr[stack.Top] = val
	return nil
}
//这里是在栈中取出元素的方法
func (stack *Stack) Pop() (val int, err error){
	if stack.Top == -1{
		return 0,errors.New("stack empty")
	}
	val = stack.arr[stack.Top]
	stack.arr[stack.Top]=0
	stack.Top--
	return val, nil
}
//遍历栈
func (stack *Stack) List(){
	if stack.Top == -1{
		fmt.Println("stack empty")
		return
	}
	for i:=stack.Top;i>=0;i--{
		fmt.Printf("arr[%d]=%d\n",i,stack.arr[i])
	}
}
//判断是否是运算符
func (stack *Stack) IsOper(val int) bool{
	return val == 43 || val == 45 || val == 42 || val == 47
}
//这里是计算的主要方法
func Calc(num1 int,num2 int,operator int)(int){
	res:=0
	switch operator{
		//这里相反是因为后进先出
	case 43 :
		res = num1 + num2
	case 45 :
		res = num2 - num1
	case 42 :
		res = num1 * num2
	case 47 :
		res = num2 / num1
	default:
		fmt.Println("运算符有误")
	}
	return res
}
//定义运算符优先级
func (stack *Stack) Priority(oper int) int{
	res:=0
	switch oper{
	case 43 :
		res = 1
	case 45 :
		res = 1
	case 42 :
		res = 2
	case 47 :
		res = 2
	default:
		fmt.Println("运算符有误")
	}
	return res
}
//计算器方法
func Calculate(exp string)(int){
	numStack := &Stack{
		arr: make([]int, 20),
		MaxSize: 20,
		Top: -1,
	}
	operatorStack := &Stack{
		arr: make([]int, 20),
		MaxSize: 20,
		Top: -1,
	}
	//定义index 用于扫描
	index:=0
	//定义运算所需的变量
	num1 := 0
	num2 := 0
	operator := 0
	result := 0
	keepNum := ""

	//循环遍历字符串
	for {
		charr := exp [index : index+1]
		//将这个charr变为数字ascii码
		temp := int([]byte(charr)[0])
		//如果是字符
		if operatorStack.IsOper(temp){
			//如果是一个空栈
			if operatorStack.Top == -1{
				operatorStack.Push(temp)
			}else{//如果不是空栈
				//如果栈顶的运算符优先级大于等于此时的运算符
				//将栈顶的运算符pop出 然后数字栈中也POP出两个进行运算
				//将运算结果重新存入数栈
				if operatorStack.Priority(operatorStack.arr[operatorStack.Top])>=operatorStack.Priority(temp){
					num1,_ = numStack.Pop()
					num2,_ = numStack.Pop()
					operator,_ = operatorStack.Pop()
					result = Calc(num1,num2,operator)
					//将计算结果重新放入数字栈
					numStack.Push(result)
					//将刚刚准备入符号栈的加进去
					operatorStack.Push(temp)
				}else{
					//直接入栈
					operatorStack.Push(temp)
				}
			}
		}else{//如果是数字
			//处理连续的数字使用keepNum作拼接
			keepNum+=charr
			//每次遍历都要检测下一个字符是不是符号
			//这里先判断是否到了算式的底端 如果是 把存的数堆栈
			if index == len(exp)-1{
				val,_ :=strconv.ParseInt(keepNum,10,64)
				numStack.Push(int(val))
			}else{
				if operatorStack.IsOper(int([]byte(exp[index+1 : index+2])[0])){
					val,_ := strconv.ParseInt(keepNum,10,64)
					numStack.Push(int(val))
					keepNum =""
				}
			}
		}
		//检测是否已经扫描到了最后
		if index+1 == len(exp) {
			break
		}
		index++
	}
//如果完全扫描完毕 从符号栈里取出一个符号 在从数字栈中取出一个数字
	for {
		//停止条件
		if operatorStack.Top==-1{
			break
		}
		//存入数栈
		num1,_ = numStack.Pop()
		num2,_ = numStack.Pop()
		operator,_ = operatorStack.Pop()
		result = Calc(num1,num2,operator)
		numStack.Push(result)
	}
	res,_ := numStack.Pop()
	return res
}