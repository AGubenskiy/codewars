func Comp(array1 []int, array2 []int) (r bool,err error) {
	if err != nil {
		return false,err
	}
	fmt.Println(len(array1),len(array2))
	if (len(array1) == 0 && len(array2) == 0) || ( fmt.Sprint(array1)==fmt.Sprint(array2)){
		return true,nil
	}else if len(array1) != len(array1) || len(array1) == 0 || len(array2) == 0 || array1==nil||array2==nil{
		return false,nil
	}else
	{
		for i,x:= range array1{
			array1[i]=x*x
		}
		sort.Ints(array1)
		sort.Ints(array2)
		return reflect.DeepEqual(array1, array2),nil
	}
	}
