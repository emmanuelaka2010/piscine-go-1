package piscine

func FindNextPrime(nb int)int{
	for i:>=nb+1{
		if premier(i){
			return i
		}
		i++	
	}

}


func premier(nb int) bool{
	decision:=true
	if nb<=1{
		return false
	}
	for i:=2;i<nb;i++{
		if nb%i==0{
			decision=false
		}
	}
	return decision
}



