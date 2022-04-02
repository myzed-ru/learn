package main

type agent007 struct {
	On    bool
	Ammo  int
	Power int
}

func (agent *agent007) Shoot() bool {
	if agent.Ammo > 0 && agent.On {
		agent.Ammo--
		return true
	}
	return false
}

func (agent *agent007) RideBike() bool {
	if agent.Power > 0 && agent.On {
		agent.Power--
		return true
	}
	return false
}
func main() {
	testStruct := new(agent007)

	//var testStruct = terminator{
	//						On: true,
	//						Ammo: 5,
	//						Power: 7,
	//					}

}
