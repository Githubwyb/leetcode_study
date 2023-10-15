package main

func distanceTraveled(mainTank int, additionalTank int) int {
	res := 0
	for mainTank >= 5 {
		var add int
		add, mainTank = mainTank/5, mainTank%5
		res += add * 50
		if additionalTank < add {
			add = additionalTank
		}
		additionalTank -= add
		mainTank += add
	}
	res += mainTank * 10
	return res
}
