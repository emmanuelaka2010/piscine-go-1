package piscine

import "math"

func FindNextPrime(nb int)int{
	nextprime:=nb-1
	i:=nb+1
	for i>nb{
		if nextprime>math.MaxInt32{
			return nextprime
		}
		nextprime++
		if premier(nextprime){
			return nextprime
		}
		i++	
	}
	return nextprime
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



